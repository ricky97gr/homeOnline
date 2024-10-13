package group

import (
	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/internal/webservice/service/system"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
)

type UIGroup struct{}

func GetAllGroupByUID(uid string) ([]model.Group, error) {
	groups, err := normalGetAllGroup()
	if err != nil {
		return nil, err
	}
	memberGroup, err := normalGetAllGroupByUID(uid)
	if err != nil {
		return nil, err
	}
	allGroups := make(map[string]model.Group)
	for _, v := range groups {
		allGroups[v.GroupUID] = v
	}
	var result []model.Group
	for _, v := range memberGroup {
		if g, ok := allGroups[v.GroupUID]; ok {
			result = append(result, g)
		}
	}
	return result, nil
}

func GetAllMemeberByGroupUID(groupUID string, p paginate.PageQuery) ([]model.User, error) {
	members, err := normalGetAllMemberByGroupUID(groupUID)
	if err != nil {
		return nil, err
	}

	allUser, _, err := system.GetAllUser(&p)
	if err != nil {
		return nil, err
	}
	users := make(map[string]model.User)
	for _, v := range allUser {
		users[v.UserID] = v
	}
	var result []model.User

	for _, v := range members {
		if u, ok := users[v.UserUID]; ok {
			result = append(result, u)
		}
	}
	return result, nil
}

func normalGetSelfGroupByUID(uid string) ([]model.Group, error) {
	c := mysql.GetClient()
	var groups []model.Group
	result := c.C.Model(&model.Group{}).Where("ownerUID = ?", uid).Find(&groups)
	return groups, result.Error
}

func normalGetAllGroupByUID(uid string) ([]model.GroupUserRelation, error) {
	c := mysql.GetClient()
	var relations []model.GroupUserRelation
	result := c.C.Model(&model.GroupUserRelation{}).Where("userUID = ?", uid).Find(&relations)
	return relations, result.Error
}

func normalGetAllGroup() ([]model.Group, error) {
	c := mysql.GetClient()
	var groups []model.Group
	result := c.C.Model(&model.Group{}).Find(&groups)
	return groups, result.Error
}

func normalGetAllMemberByGroupUID(groupUID string) ([]model.GroupUserRelation, error) {
	c := mysql.GetClient()
	var relations []model.GroupUserRelation
	result := c.C.Model(&model.GroupUserRelation{}).Where("groupUID = ?", groupUID).Find(&relations)
	return relations, result.Error
}
