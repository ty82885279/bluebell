package logic

import (
	"strconv"
	"web_app/dao/redis"
	"web_app/model"

	"go.uber.org/zap"
)

//
func VoteForPost(userID int64, p *model.ParamVoteDatas) (err error) {

	zap.L().Debug("VoteForPost", zap.Int64("userID", userID), zap.String("postID", p.PostID), zap.Int("value", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))

}
