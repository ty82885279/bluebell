package logic

import (
	"web_app/dao/mysql"
	"web_app/model"
)

//获取社区列表
func GetCommunityList() ([]*model.Community, error) {
	//

	return mysql.GetCommunityList()

}

//获取社区详情
func CommunityDetail(id int64) (*model.CommunityDetail, error) {

	return mysql.GetCommunityDetailByID(id)
}
