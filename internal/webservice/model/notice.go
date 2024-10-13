package model

type NoticeInfo struct {
	UID        string `json:"uid" bson:"uid"`
	Type       int    `json:"type" bson:"type"`
	CreateTime int64  `json:"createTime" bson:"createTime"`
	TargetUID  string `json:"targetUID" bson:"targetUID"`
	Status     int    `json:"status" bson:"status"`
}
