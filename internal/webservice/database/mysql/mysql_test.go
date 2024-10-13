package mysql

import (
	"fmt"
	"testing"

	"github.com/ricky97gr/homeOnline/internal/conf"
)

func TestInitMySql(t *testing.T) {
	config := conf.GetConfig()
	fmt.Println(config)
	_, err := InitMySql(config.Mysql.IP, config.Mysql.User, config.Mysql.Password, config.Mysql.DB, config.Mysql.Port)
	if err != nil {
		t.Logf("failed to connect to mysql, err: %+v\n", err)
		return
	}
	t.Log("successfully to connect to mysql")
}
