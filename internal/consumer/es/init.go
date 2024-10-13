package es

import (
	"fmt"
	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/nsqio/go-nsq"
	"time"
)

type EsHandler struct {
	Title string
}

// HandleMessage 是需要实现的处理消息的方法
func (m *EsHandler) HandleMessage(msg *nsq.Message) (err error) {
	newlog.Logger.Infof("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return nil
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
	consumer := &EsHandler{
		Title: "es",
	}
	c.AddHandler(consumer)

	// if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
	if err := c.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询
		return err
	}
	return nil

}
