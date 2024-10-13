package model

import proto "github.com/ricky97gr/homeOnline/internal/grpcserver/proto/log"

type OperationLogInfo struct {
	CreateTime int64  `json:"createTime" bson:"createTime" gorm:"column:createTime"`
	Msg        string `json:"msg" bson:"msg" gorm:"column:msg"`
	Status     string `json:"status" bson:"status" gorm:"column:status"`
	Module     string `json:"module" bson:"module" gorm:"module"`
	UserID     string `json:"userID" bson:"userID" gorm:"userID"`
	UUID       string `json:"uuid" bson:"uuid" gorm:"uuid"`
}

func (l *OperationLogInfo) Convert(info *proto.OperationLogInfo) {
	l.Msg = info.Msg
	l.CreateTime = info.CreateTime
	l.UserID = info.User
	l.Module = info.Module
}
