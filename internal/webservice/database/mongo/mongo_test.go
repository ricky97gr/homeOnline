package mongo

import (
	"testing"

	"github.com/ricky97gr/homeOnline/internal/conf"
)

func TestInitMongo(t *testing.T) {
	config := conf.GetConfig()
	_, err := InitMongo(config.Mongo.User, config.Mongo.Password, config.Mongo.IP, config.Mongo.Port, config.Mongo.DB)
	if err != nil {
		t.Logf("failed to connect to mongo, err: %+v\n", err)
		return
	}
	t.Log("successfully to connect to mongo")
}
