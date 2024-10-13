package database

import (
	"strconv"
	"time"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/pkg/groupid"
	"github.com/ricky97gr/homeOnline/pkg/userid"
)

func Start() {
	c, err := mysql.GetMysqlClient()
	if err != nil {
		newlog.Logger.Errorf("mysql.GetMysqlClient failed: %v\n", err)
		return
	}
	err = c.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Tag{},
		&model.CommunityComment{},
		&model.Topic{},
		&model.GiveLike{},
		&model.Score{},
		&model.Article{},

		&model.Group{},
		&model.GroupUserRelation{},

		&model.FriendShip{},
		&model.Plugin{},
	)
	if err != nil {
		newlog.Logger.Errorf("failed to auto migrate mysql table, err:%+v\n", err)
		return
	}
	createSuperGroup()
	createSuperAdminUser()
	createSuperGroupRelation()

}

func createSuperAdminUser() {
	u := model.User{
		UserID:     strconv.Itoa(userid.SuperAdministrator),
		UserName:   "超级管理员",
		Password:   "123456",
		NickName:   "超级管理员",
		Theme:      "",
		Avatar:     "",
		Role:       3,
		Phone:      "13888888888",
		Email:      "forgocode@163.com",
		Status:     1,
		Sex:        1,
		CreateTime: time.Now().UnixMilli(),
	}
	c, err := mysql.GetMysqlClient()
	if err != nil {
		newlog.Logger.Errorf("failed to get mysql client, err:%+v\n", err)
		return
	}
	var user model.User
	result := c.Where("userID = ?", u.UserID).Find(&user)
	if result.Error != nil {
		newlog.Logger.Errorf("failed to get user by %+v, err:%+v\n", u, err)
		return
	}
	if result.RowsAffected == 0 {
		result = c.Create(&u)
		if result.Error != nil {
			newlog.Logger.Errorf("failed to create user info: %+v, err:%+v\n", u, err)
			return
		}
	}
}

func createSuperGroup() {
	group := &model.Group{
		GroupUID:    strconv.Itoa(groupid.SuperAdministratorGroup),
		Name:        "所有用户群",
		OwnerUID:    strconv.Itoa(userid.SuperAdministrator),
		CreateTime:  time.Now().UnixMilli(),
		Level:       1,
		Description: "所有用户都会加入的群组",
	}
	c, err := mysql.GetMysqlClient()
	if err != nil {
		newlog.Logger.Errorf("failed to get mysql client, err: %+v\n", err)
		return
	}
	var g model.Group
	result := c.Where("groupUID = ?", group.GroupUID).Find(&g)
	if result.Error != nil {
		newlog.Logger.Errorf("failed to get group by %+v, err:%+v\n", group, err)
		return
	}
	if result.RowsAffected == 0 {
		result = c.Create(&group)
		if result.Error != nil {
			newlog.Logger.Errorf("failed to create group info: %+v, err: %+v\n", group, err)
			return
		}
	}
}

func createSuperGroupRelation() {
	relation := &model.GroupUserRelation{
		GroupUID:   strconv.Itoa(groupid.SuperAdministratorGroup),
		UserUID:    strconv.Itoa(userid.SuperAdministrator),
		Relation:   model.GroupOwner,
		CreateTime: time.Now().UnixMilli(),
	}
	c, err := mysql.GetMysqlClient()
	if err != nil {
		newlog.Logger.Errorf("failed to get mysql client, err: %+v\n", err)
		return
	}
	var r model.GroupUserRelation
	result := c.Where("groupUID = ? and userUID = ?", relation.GroupUID, relation.UserUID).Find(&r)
	if result.Error != nil {
		newlog.Logger.Errorf("failed to get relation by %+v, err:%+v\n", relation, err)
		return
	}
	if result.RowsAffected == 0 {
		result = c.Create(&relation)
		if result.Error != nil {
			newlog.Logger.Errorf("failed to create relation info: %+v, err: %+v\n", relation, err)
			return
		}
	}
}
