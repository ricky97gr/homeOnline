package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/ricky97gr/homeOnline/internal/conf"
)

func GetMongoClient(collection string) (*mongo.Collection, error) {
	var err error
	if c == nil {
		config := conf.GetConfig()
		fmt.Printf("%+v\n", config)
		dbName = config.Mongo.DB
		c, err = InitMongo(config.Mongo.User, config.Mongo.Password, config.Mongo.IP, config.Mongo.Port, config.Mongo.DB)
		return c.Database(config.Mongo.DB).Collection(collection), err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		config := conf.GetConfig()
		c, err = InitMongo(config.Mongo.User, config.Mongo.Password, config.Mongo.IP, config.Mongo.Port, config.Mongo.DB)
		if err != nil {
			return nil, err
		}

	}
	return c.Database(dbName).Collection(collection), nil

}

var dbName string

var c *mongo.Client

func InitMongo(user string, passwd string, ip string, port uint16, db string) (*mongo.Client, error) {
	auth := options.Credential{
		AuthMechanism:           "",
		AuthMechanismProperties: nil,
		AuthSource:              "",
		Username:                user,
		Password:                passwd,
		PasswordSet:             false,
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, passwd, ip, port)).SetConnectTimeout(5*time.Second).SetAuth(auth))
	if err != nil {
		return nil, err
	}
	return client, nil
}
