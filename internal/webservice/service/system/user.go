package system

import "C"
import (
	"errors"
	"time"

	"github.com/ricky97gr/homeOnline/internal/webservice/database/mysql"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/ricky97gr/homeOnline/pkg/paginate"
)

type UIUser struct {
	NickName      string `json:"nickName"`
	Sex           int    `json:"sex"`
	UID           string `json:"uid"`
	Role          int    `json:"role"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	IsShow        string `json:"isShow"`
	CreateTime    int64  `json:"createTime"`
	LastLoginTime int64  `json:"lastLoginTime"`
}

func (u *UIUser) Convert() *model.User {
	return &model.User{
		UserID:        u.UID,
		NickName:      u.NickName,
		Password:      u.Password,
		Theme:         "",
		Avatar:        "",
		Role:          u.Role,
		Phone:         u.Phone,
		Email:         u.Email,
		Status:        model.UserIsNormal,
		CreateTime:    time.Now().UnixMilli(),
		LastLoginTime: 0,
		Score:         0,
		Sex:           u.Sex,
	}
}

func UpdateUserScore(uid string, score int64) error {
	c := mysql.GetClient()
	return c.C.Model(&model.User{}).Where("userID = ?", uid).Update("score", score).Error

}

func DeleteUser(userID string) error {
	return deleteUser(userID)
}

func UpdateUser(userID string, isShow bool) error {
	if isShow {
		return unBanUser(userID)
	}
	return banUser(userID)
}

func GetUserScore(uid string) (int64, error) {
	c := mysql.GetClient()
	var u model.User
	result := c.C.Model(&model.User{}).Where("userID = ?", uid).Find(&u)
	if result.Error != nil {
		return 0, result.Error
	}
	return u.Score, nil
}

func GetUserByPhone(phone string, passwd string) (*model.User, error) {

	user, err := getUserByPhone(phone, passwd)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func AdminCreateUser(user *UIUser) error {
	return createUser(user.Convert())
}

func GetAllUser(q *paginate.PageQuery) ([]model.User, int64, error) {

	users, err := getAllUser(q)
	if err != nil {
		return nil, 0, err
	}
	count, err := getUserCount()
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}

func UpdateUserLastLogin(userID string) error {
	u := model.User{UserID: userID, LastLoginTime: time.Now().UnixMilli()}
	return updateUserInfo(u)
}

func UserAddTrend() (interface{}, error) {
	type resultInfo struct {
		Date  time.Time `json:"date"`
		Count int       `json:"count"`
	}
	type trendInfo struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}
	c := mysql.GetClient()
	var info []resultInfo
	var trend []trendInfo

	result := c.C.Raw("select a.created_at as date,ifnull (b.count, 0) as count\nfrom(\n    SELECT curdate() as created_at\n    union all\n    SELECT date_sub(curdate(), interval 1 day) as created_at\n    union all\n    SELECT date_sub(curdate(), interval 2 day) as created_at\n    union all\n    SELECT date_sub(curdate(), interval 3 day) as created_at\n    union all\n    SELECT date_sub(curdate(), interval 4 day) as created_at\n    union all\n    SELECT date_sub(curdate(), interval 5 day) as created_at\n    union all\n    SELECT date_sub(curdate(), interval 6 day) as created_at\n) a left join (\nSELECT DATE_FORMAT( from_unixtime(createTime/1000),'%Y-%m-%d')  as date,count(*) as count FROM user GROUP BY date\n) b on a.created_at = b.date order by a.created_at asc;").Scan(&info)
	for _, i := range info {
		trend = append(trend, trendInfo{Date: i.Date.Format("2006/01/02"), Count: i.Count})
	}
	return trend, result.Error
}

func UserActiveCount() (interface{}, error) {
	type resultInfo struct {
		LoginCount   int `json:"loginCount"`
		UnLoginCount int `json:"unLoginCount"`
	}
	c := mysql.GetClient()
	var info resultInfo
	result := c.C.Raw("SELECT SUM(CASE WHEN FROM_UNIXTIME(lastLoginTime / 1000) >= NOW() - INTERVAL 30 DAY THEN 1 ELSE 0 END) AS LoginCount, SUM(CASE WHEN FROM_UNIXTIME(lastLoginTime / 1000) < NOW() - INTERVAL 30 DAY OR lastLoginTime IS NULL THEN 1 ELSE 0 END) AS UnLoginCount FROM user;").Scan(&info)
	return info, result.Error
}

func UserRoleCount() (interface{}, error) {
	type RoleStats struct {
		Role      int `json:"role"`      // 角色 (1: 普通用户, 2: 管理员, 3: 超级管理员)
		RoleCount int `json:"roleCount"` // 每个角色的用户数量
	}

	var roleStats []RoleStats
	c := mysql.GetClient()
	result := c.C.Raw(`
    SELECT role, COUNT(*) AS RoleCount
    FROM user
    GROUP BY role
`).Scan(&roleStats)
	return roleStats, result.Error
}

func GetUserByUserID(userID string) (*model.User, error) {
	c := mysql.GetClient()
	var user model.User
	result := c.C.Where("userID = ?", userID).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func UserScoreTop() ([]model.User, error) {
	c := mysql.GetClient()
	var users []model.User
	result := c.C.Model(&model.User{}).Select("userID", "nickName", "score").Order("score desc").Limit(10).Find(&users)
	return users, result.Error
}

func UserActive30() {
	type resultInfo struct {
		Role  int `json:"role"`
		Count int `json:"count"`
	}
	// var info []resultInfo
	// c := mysql.GetClient()
	// c.C.Raw("select role, count(*) as count from user group by ac")

}

func createUser(user *model.User) error {
	c := mysql.GetClient()

	if ok, err := isPhoneExist(user.Phone); err != nil || ok {
		return errors.New("phone is exist")
	}
	result := c.C.Create(user)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}

func isUserIDExist(userID string) (bool, error) {
	c := mysql.GetClient()
	var user model.User
	result := c.C.Where("userID = ?", userID).Find(&user)
	if result.Error != nil {
		return true, result.Error
	}
	if result.RowsAffected != 0 {
		return true, nil
	}
	return false, nil
}

func isPhoneExist(phone string) (bool, error) {
	c := mysql.GetClient()
	var user model.User
	result := c.C.Where("phone = ?", phone).Find(&user)
	if result.Error != nil {
		return true, result.Error
	}
	if result.RowsAffected != 0 {
		return true, nil
	}
	return false, nil
}

func getUserCount() (int64, error) {
	c := mysql.GetClient()
	var count int64
	result := c.C.Model(&model.User{}).Count(&count)
	return count, result.Error
}

func getAllUser(q *paginate.PageQuery) ([]model.User, error) {
	c := mysql.GetClient()
	var users []model.User
	result := c.C.Model(&model.User{}).Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&users)
	return users, result.Error
}

func getUserByPhone(phone, passwd string) (*model.User, error) {
	c := mysql.GetClient()
	var user model.User
	result := c.C.Where("phone = ? AND password = ?", phone, passwd).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func banUser(userID string) error {
	c := mysql.GetClient()
	result := c.C.Model(&model.User{}).Where("userID = ?", userID).Update("status", model.UserIsBaned)
	return result.Error
}

func unBanUser(userID string) error {
	c := mysql.GetClient()
	result := c.C.Model(&model.User{}).Where("userID = ?", userID).Update("status", model.UserIsNormal)
	return result.Error
}

func deleteUser(userID string) error {
	c := mysql.GetClient()
	result := c.C.Where("userID = ?", userID).Delete(&model.User{})
	return result.Error
}

// 需要更新哪个就赋值哪个字段
func updateUserInfo(user model.User) error {
	c := mysql.GetClient()
	result := c.C.Where("userID = ?", user.UserID).Updates(user)
	return result.Error

}

func getNewUserInfo(newobj, oldobj model.User) model.User {
	if newobj.NickName != "" {
		oldobj.NickName = newobj.NickName
	}
	if newobj.Email != "" {
		oldobj.Email = newobj.Email
	}
	if newobj.Password != "" {
		oldobj.Password = newobj.Password
	}
	if newobj.Phone != "" {
		oldobj.Phone = newobj.Phone
	}
	if newobj.Avatar != "" {
		oldobj.Avatar = newobj.Avatar
	}
	return oldobj

}
