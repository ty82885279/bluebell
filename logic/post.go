package logic

import (
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/model"
	"web_app/pkg/snowflake"

	"go.uber.org/zap"
)

// CreatePost  创建帖子
func CreatePost(p *model.Post) (err error) {
	//
	p.ID = snowflake.GetID()
	err = mysql.CreatePost(p)
	if err != nil {
		return
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return

}

// GetPostDetailByID  获取帖子详情
func GetPostDetailByID(id int64) (postDetail *model.ApiPostDetail, err error) {

	return mysql.GetPostDetailByID(id)
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*model.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return
	}
	data = make([]*model.ApiPostDetail, 0, len(posts))
	for _, post := range posts {

		//查询作者
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			continue
		}

		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			continue
		}
		postDetail := &model.ApiPostDetail{
			AuthorName:      user.UserName,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}

	return
}

//根据排序(时间or分数)返回帖子列表
func GetPostList2(p *model.ParamPostList) (data []*model.ApiPostDetail, err error) {

	//去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))

	//提前查好帖子的投赞成票的数据
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	//根据ids查询帖子列表
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	//组合数据
	for idx, post := range posts {
		//查询作者
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			continue
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			continue
		}
		postDetail := &model.ApiPostDetail{
			AuthorName:      user.UserName,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

//GetCommunityPostList 根据社区返回帖子列表
func GetCommunityPostList(p *model.ParamPostList) (data []*model.ApiPostDetail, err error) {
	//去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetCommunityPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetCommunityPostList", zap.Any("ids", ids))

	//提前查好帖子的投赞成票的数据
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	//根据ids查询帖子列表
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	//组合数据
	for idx, post := range posts {
		//查询作者
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			continue
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			continue
		}
		postDetail := &model.ApiPostDetail{
			AuthorName:      user.UserName,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostListNew(p *model.ParamPostList) (data []*model.ApiPostDetail, err error) {
	if p.CommunityID == 0 {
		data, err = GetPostList2(p)
		zap.L().Debug("GetPostListNew(p *model.ParamPostList)", zap.Any("p.CommunityID", p.CommunityID))
	} else {
		data, err = GetCommunityPostList(p)
		zap.L().Debug("GetPostListNew(p *model.ParamPostList)", zap.Any("p.CommunityID", p.CommunityID))
	}
	if err != nil {
		zap.L().Error("GetPostListNew(p *model.ParamPostList)", zap.Error(err))
		return nil, err
	}

	return
}
