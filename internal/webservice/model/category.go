package model

type Category struct {
	CreateTime int64  `json:"createTime" gorm:"column:createTime"`
	Creator    string `json:"creator" gorm:"column:creator"`
	Name       string `json:"name" gorm:"column:name"`
	IsShow     bool   `json:"isShow" gorm:"column:isShow"`
	UUID       string `json:"uuid" gorm:"column:uuid"`
	UsedCount  int64  `json:"usedCount" gorm:"column:usedCount"`
}

func (t Category) TableName() string {
	return "category"
}
