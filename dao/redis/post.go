package redis

import (
	"strconv"
	"time"
	"web_app/model"

	"github.com/go-redis/redis"
)

// 创建帖子
func CreatePost(postID, communityID int64) (err error) {
	//
	pipeline := rdb.Pipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: strconv.Itoa(int(postID)),
	})
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	pipeline.SAdd(getRedisKey(KeyCommunitySetPF+strconv.Itoa(int(communityID))), postID)
	_, err = pipeline.Exec()
	return
}

func getIDsFromKey(key string, page, size int64) ([]string, error) {
	// 确定查询的索引范围
	start := (page - 1) * size
	end := start + size - 1
	//ZREVRANGE 按分数从大到小的顺序查询指定数量的元素
	return rdb.ZRevRange(key, start, end).Result()
}

// GetPostIDsInOrder 根据范围返回IDs
func GetPostIDsInOrder(p *model.ParamPostList) ([]string, error) {
	// 获取key的类型
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == model.OrderByScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDsFromKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据ids查询每篇帖子投赞成票的数据
func GetPostVoteData(ids []string) (voteData []int64, err error) {
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		pipeline.ZCount(getRedisKey(KeyPostVotedZSetPF+id), "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return
	}
	voteData = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		num := cmder.(*redis.IntCmd).Val()
		voteData = append(voteData, num)
	}
	return

}

//GetCommunityPostIDsInOrder 按社区查找ids
func GetCommunityPostIDsInOrder(p *model.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == model.OrderByScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	//使用缓存key,减少zinterstore的执行次数
	//zinterzet redis联合查询 一个新key,2个需要联合的key,
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	key := orderKey + ":" + strconv.Itoa(int(p.CommunityID))
	if rdb.Exists(key).Val() < 1 { //不存在，创建key，缓存时间120s

		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, ckey, orderKey)
		pipeline.Expire(key, time.Minute*2)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFromKey(key, p.Page, p.Size)
}
