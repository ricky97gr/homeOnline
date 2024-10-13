package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/webservice/database/mongo"
	"github.com/nsqio/go-nsq"
	"time"
)

type MongoHandler struct {
	Title string
	DB    string
}

// HandleMessage 是需要实现的处理消息的方法
func (m *MongoHandler) HandleMessage(msg *nsq.Message) (err error) {
	newlog.Logger.Infof("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	c, err := mongo.GetMongoClient(m.DB)
	if err != nil {
		newlog.Logger.Errorf("failed to get mongo client, err: %+v\n", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	var tmp map[string]interface{}
	err = json.Unmarshal(msg.Body, &tmp)
	if err != nil {
		newlog.Logger.Errorf("failed to unmarshal json, err: %+v\n", err)
		return
	}
	_, err = c.InsertOne(ctx, tmp)
	if err != nil {
		newlog.Logger.Errorf("failed to insert log info: %+v\n", err)
	}
	cancel()
	return
}

// 初始化消费者
func InitConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	consumer := &MongoHandler{
		Title: "mongodb",
		DB:    "operation_log",
	}
	c.AddHandler(consumer)

	// if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
	if err := c.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询
		return err
	}
	return nil

}
