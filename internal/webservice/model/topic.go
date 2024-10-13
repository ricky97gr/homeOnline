package model

type Topic struct {
	CreateTime int64  `json:"createTime" gorm:"column:createTime"`
	Creator    string `json:"creator" gorm:"column:creator"`
	Name       string `json:"name" gorm:"column:name"`
	IsShow     bool   `json:"isShow" gorm:"column:isShow"`
	UsedCount  int64  `json:"usedCount" gorm:"column:usedCount"`
	UUID       string `json:"uuid" gorm:"column:uuid"`
}

func (t Topic) TableName() string {
	return "topic"
}
