package mysql

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ricky97gr/homeOnline/internal/conf"
)

type MysqlClient struct {
	C *gorm.DB
}

var c *gorm.DB

// TODO: 有问题
func GetClient() *MysqlClient {
	if c == nil {
		c, _ = GetMysqlClient()
	}
	return &MysqlClient{c}
}

func GetMysqlClient() (*gorm.DB, error) {
	if c == nil {
		config := conf.GetConfig()
		return InitMySql(config.Mysql.IP, config.Mysql.User, config.Mysql.Password, config.Mysql.DB, config.Mysql.Port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := c.DB()
	if err != nil {
		config := conf.GetConfig()
		c, err = InitMySql(config.Mysql.IP, config.Mysql.User, config.Mysql.Password, config.Mysql.DB, config.Mysql.Port)
		if err != nil {
			return nil, err
		}

	}
	err = db.PingContext(ctx)
	if err != nil {
		config := conf.GetConfig()
		c, err = InitMySql(config.Mysql.IP, config.Mysql.User, config.Mysql.Password, config.Mysql.DB, config.Mysql.Port)
		if err != nil {
			return nil, err
		}
	}
	return c, nil

}

func InitMySql(url, user, passwd, dbName string, port uint16) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, url, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
