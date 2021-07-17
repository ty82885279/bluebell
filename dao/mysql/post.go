package mysql

import (
	"strings"
	"web_app/model"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

func CreatePost(p *model.Post) (err error) {

	sqlStr := `insert into post(post_id,title,content,author_id,community_id) values(?,?,?,?,?)`
	_, err = DB.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostDetailByID(id int64) (*model.ApiPostDetail, error) {
	// 查询帖子详情

	post := new(model.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,status,create_time from post where post_id = ?`
	err := DB.Get(post, sqlStr, id)
	if err != nil {
		zap.L().Error("GetPostDetailByID", zap.Int64("id", id), zap.Error(err))
		return nil, err

	}
	// 查询社区详情
	communityDetail, err := GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("GetPostDetailByID->GetCommunityDetailByID", zap.Int64("id", id), zap.Error(err))
		return nil, err

	}
	// 查询作者
	user, err := GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("GetPostDetailByID->GetAuthor failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return nil, err
	}
	var data = &model.ApiPostDetail{
		AuthorName:      user.UserName,
		Post:            post,
		CommunityDetail: communityDetail,
	}

	return data, err
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (posts []*model.Post, err error) {
	//posts = make([]*model.Post, 0, 5)
	sqlStr := `select post_id,title,content,author_id,community_id,status,create_time from post limit ?,?`
	err = DB.Select(&posts, sqlStr, (page-1)*size, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList err", zap.Error(err))
		return
	}
	return
}

// GetPostListByIDs 根据IDs查询帖子列表
func GetPostListByIDs(ids []string) (posts []*model.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,status,create_time from post 
    	       where post_id in (?)
			   order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = DB.Rebind(query)
	err = DB.Select(&posts, query, args...)
	return

}
