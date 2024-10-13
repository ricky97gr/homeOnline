package friend

import (
	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"time"
)

func CreateFriendShip(uid, friendUID string) error {
	return createFriendShip(model.FriendShip{
		UID:        uid,
		FriendUID:  friendUID,
		CreateTime: time.Now().UnixMilli(),
		Status:     0,
	})
}

func RemoveFriendShip(uid, friendUID string) error {
	return removeFriendShip(uid, friendUID)
}

func GetFollowFriendList(uid string) ([]model.FriendShip, error) {
	c := mysql.GetClient()
	var ships []model.FriendShip
	result := c.C.Model(&model.FriendShip{}).Where("uid = ?", uid).Find(&ships)
	return ships, result.Error
}

func GetFollowedFriendList(uid string) ([]model.FriendShip, error) {
	c := mysql.GetClient()
	var ships []model.FriendShip
	result := c.C.Model(&model.FriendShip{}).Where("friendUID = ?", uid).Find(&ships)
	return ships, result.Error
}
func createFriendShip(ship model.FriendShip) error {
	c := mysql.GetClient()
	return c.C.Model(&model.FriendShip{}).Create(&ship).Error
}

func removeFriendShip(uid, friendUID string) error {
	c := mysql.GetClient()
	return c.C.Where("uid = ? and friendUID = ?", uid, friendUID).Delete(&model.FriendShip{}).Error
}
