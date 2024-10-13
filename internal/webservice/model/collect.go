package model

// 收藏
type Collect struct {
	ID       string `json:"ID" gorm:"column:ID"`
	UserID   string `json:"UserID" gorm:"column:userID"`
	UserName string `json:"userName" gorm:"column:userName"`
	Type     int8   `json:"type" gorm:"column:type"`
}

func (Collect) TableName() string {
	return "collect"
}
