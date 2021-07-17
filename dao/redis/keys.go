package redis

const (
	Prefix             = "bluebell:"   //项目前缀
	KeyPostTimeZSet    = "post:time"   //zset:帖子发布时间
	KeyPostScoreZSet   = "post:score"  //zset:帖子分数
	KeyPostVotedZSetPF = "post:voted:" //zset:记录用户投票及投票类型  参数：post_id
	KeyCommunitySetPF  = "community:"  //set:记录每个社区下的帖子    参数：community_id
)

func getRedisKey(key string) string {
	return Prefix + key
}
