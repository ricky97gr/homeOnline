package model

type FriendShip struct {
	UID        string `json:"uid" gorm:"column:uid"`
	FriendUID  string `json:"friendUID" gorm:"column:friendUID"`
	CreateTime int64  `json:"createTime" gorm:"column:createTime"`
	Status     int    `json:"status" gorm:"column:status"`
}

const (
	FriendShipSend = iota + 1
	FriendShipAgree
	FriendShipBlack
)

func (FriendShip) TableName() string {
	return "firend_ship"
}
