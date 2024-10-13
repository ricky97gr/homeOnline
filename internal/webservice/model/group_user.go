package model

type GroupUserRelation struct {
	GroupUID   string `json:"groupUID" gorm:"column:groupUID"`
	UserUID    string `json:"userUID" gorm:"column:userUID"`
	Relation   int    `json:"relation" grom:"column:relation"`
	CreateTime int64  `json:"createTime" gorm:"column:createTime"`
}

const (
	GroupOwner        = 1
	GroupManager      = 2
	GroupNormalMember = 3
)

func (GroupUserRelation) TableName() string {
	return "group_user_relation"
}
