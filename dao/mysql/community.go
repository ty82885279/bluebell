package mysql

import (
	"database/sql"
	"errors"
	"web_app/model"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*model.Community, err error) {

	sqlStr := "select community_id, community_name from community"

	if err = DB.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("There is no community in db")
			err = nil
		}
	}

	return
}
func GetCommunityDetailByID(id int64) (community *model.CommunityDetail, err error) {
	community = new(model.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction,create_time from community where community_id =?`
	if err = DB.Get(community, sqlStr, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("无效的id")
			err = ErrorinvalidID
		}
	}
	return
}
