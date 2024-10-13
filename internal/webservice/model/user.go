package model

type User struct {
	UserID   string `json:"userID" gorm:"column:userID; index"`
	UserName string `json:"userName" gorm:"column:userName;comment: the name of user"`
	Password string `json:"-" gorm:"column:password; comment: the password of user"`
	NickName string `json:"nickName" gorm:"column:nickName; comment: the nickName of user"`
	Theme    string `json:"theme" gorm:"column:theme; comment: the theme of user"`
	Avatar   string `json:"avatar" gorm:"column:avatar; comment: the avatar of user"`
	//1:normal user 2: normal admin 3: super admin
	Role          int    `json:"role" gorm:"column:role; comment: the role of user"`
	Phone         string `json:"phone" gorm:"column:phone; comment: the phone of user"`
	Email         string `json:"email" gorm:"column:email; comment: the email of user"`
	Status        int    `json:"status" gorm:"column:status; comment: the status of user"`
	CreateTime    int64  `json:"createTime" bson:"createTime" gorm:"column:createTime"`
	LastLoginTime int64  ` json:"lastLoginTime" bson:"lastLoginTime" gorm:"column:lastLoginTime"`
	Score         int64  `json:"score" bson:"score" gorm:"column:score"`
	Description   string `json:"description" bson:"description" gorm:"column:description"`
	//1: man 2: women
	Sex int `json:"sex" gorm:"column:sex"`
}

const (
	UserIsNormal = 1
	UserIsBaned  = 2
)

const (
	NormalUser = 1
	Admin      = 2
	SuperAdmin = 3
)

func (User) TableName() string {
	return "user"
}
