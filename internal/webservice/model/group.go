package model

type Group struct {
	GroupUID    string `json:"groupUID" gorm:"column:groupUID"`
	Name        string `json:"name" gorm:"column: name"`
	OwnerUID    string `json:"ownerUID" gorm:"column: ownerUID"`
	CreateTime  int64  `json:"createTime" gorm:"column: createTime"`
	Level       int    `json:"level" gorm:"column: level"`
	Description string `json:"description" gorm:"column: description"`
	// ManagerUUID []string `json:"managerUUID" gorm:"column:managerUUID"`
}

func (Group) TableName() string {
	return "group"
}
