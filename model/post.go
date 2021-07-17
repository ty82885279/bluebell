package model

import "time"

type Post struct {
	ID          int64     `json:"id" db:"post_id"`                                   //帖子id
	AuthorID    int64     `json:"author_id" db:"author_id"`                          //作者id
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` //社区id
	Status      int32     `json:"status" db:"status"`                                //帖子状态
	Title       string    `json:"title" db:"title" bingding:"required"`              //标题
	Content     string    `json:"content" db:"content" binding:"required"`           //内容
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      //帖子创建时间
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum    int64  `json:"vote_num"`
	*Post
	*CommunityDetail `json:"community"`
}
