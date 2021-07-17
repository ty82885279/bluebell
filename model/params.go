package model

const (
	OrderByTime  = "time"
	OrderByScore = "score"
)

type ParamSignUp struct {
	UserName   string `json:"username" form:"username" binding:"required"`                      // 用户名 必填
	Password   string `json:"password" form:"password" binding:"required"`                      // 密码   必填
	RePassword string `json:"repassword" form:"repassword" binding:"required,eqfield=Password"` //确认密码 必填
}

type ParamLogin struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
type ParamVoteDatas struct {
	// UserID 从context中获取
	PostID    string `json:"post_id" binding:"required"`
	Direction int    `json:"direction" binding:"oneof= -1 0 1"`
}

type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"` // 社区ID  必填
	Page        int64  `json:"page" form:"page"`                 // 页数	   选填 默认:1
	Size        int64  `json:"size" form:"size"`                 // 条数    选填 默认:10
	Order       string `json:"page" form:"order" example:"time"` // 排序    选填 默认:time
}

// 社区帖子列表参数
//type ParamCommunityIDPostList struct {
//	*ParamPostList
//	CommunityID int64 `json:"community_id" form:"community_id" binding:"required"`
//}
