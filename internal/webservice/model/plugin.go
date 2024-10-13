package model

type Plugin struct {
	Name        string `json:"name" gorm:"column:name"`
	Md5         string `json:"md5" gorm:"column:md5"`
	Version     string `json:"version" gorm:"column:version"`
	Author      string `json:"author" gorm:"author"`
	Description string `json:"description" gorm:"description"`
	Status      int    `json:"status" gorm:"column:status"`
}

const (
	Running = iota + 1
	Stopped
	Upgrading
	Checking
)

func (Plugin) TableName() string {
	return "plugin"
}
