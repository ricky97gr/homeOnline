package sendlog

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ricky97gr/homeOnline/internal/webservice/model"
	"github.com/nsqio/go-nsq"
	"time"

	"google.golang.org/grpc"

	proto "github.com/ricky97gr/homeOnline/internal/grpcserver/proto/log"
	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
)

type msgStruct struct {
	ModuleCN string
	ModuleEN string
	MsgCN    string
	MsgEN    string
}

const (
	LoginCode = iota + 100000
	AddUser
	DeleteUser
	NewCategory
	DeleteCategory

	NewTag
	DeleteTag

	NewTopic
	DeleteTopic
)

const (
	SystemModuleEN = "system"
	SystemModuleCN = "系统模块"

	TagEN      = "Tags"
	TagCN      = "标签模块"
	CategoryEN = "Category"
	CategoryCN = "类别模块"
	TopicEN    = "Topic"
	TopicCN    = "话题模块"
)

// 前三位模块，后三位递增
var msgMap = map[int32]msgStruct{
	LoginCode: {
		ModuleCN: SystemModuleCN,
		ModuleEN: SystemModuleEN,
		MsgCN:    "用户登录成功",
		MsgEN:    "user login successfully!",
	},

	AddUser: {
		ModuleCN: SystemModuleCN,
		ModuleEN: SystemModuleEN,
		MsgCN:    "新建用户 %s",
		MsgEN:    "create new user %s",
	},

	DeleteUser: {
		ModuleCN: SystemModuleCN,
		ModuleEN: SystemModuleEN,
		MsgCN:    "用户登录成功",
		MsgEN:    "user login successfully!",
	},
	NewCategory: {
		ModuleCN: CategoryCN,
		ModuleEN: CategoryEN,
		MsgCN:    "新建分类 %s",
		MsgEN:    "new category %s",
	},
	DeleteCategory: {
		ModuleCN: CategoryCN,
		ModuleEN: CategoryEN,
		MsgCN:    "删除分类 %s",
		MsgEN:    "delete category %s",
	},
	NewTag: {
		ModuleCN: TagCN,
		ModuleEN: TagEN,
		MsgCN:    "新建标签 %s",
		MsgEN:    "new tag %s",
	},
	DeleteTag: {
		ModuleCN: TagCN,
		ModuleEN: TagEN,
		MsgCN:    "删除标签 %s",
		MsgEN:    "delete tag %s",
	},
	NewTopic: {
		ModuleCN: TopicCN,
		ModuleEN: TopicEN,
		MsgCN:    "新建话题 %s",
		MsgEN:    "new topic %s",
	},
	DeleteTopic: {
		ModuleCN: TopicCN,
		ModuleEN: TopicEN,
		MsgCN:    "删除话题 %s",
		MsgEN:    "delete topic %s",
	},
}

func getModuleByLangAndCode(lang string, msgCode int32) string {
	if lang == "en" {
		return msgMap[msgCode].ModuleEN
	}
	return msgMap[msgCode].ModuleCN
}

func getMessageByLangAndCode(lang string, msgCode int32, detail ...string) string {
	if lang == "en" {
		if len(detail) != 0 {
			return fmt.Sprintf(msgMap[msgCode].MsgEN, detail)
		}
		return msgMap[msgCode].MsgEN
	}
	if len(detail) != 0 {
		return fmt.Sprintf(msgMap[msgCode].MsgCN, detail)
	}
	return msgMap[msgCode].MsgCN
}

func SendOperationLog(userID, lang string, msgCode int32, detail ...string) error {
	logInfo := &proto.OperationLogInfo{
		User:       userID,
		Module:     getModuleByLangAndCode(lang, msgCode),
		CreateTime: time.Now().UnixMilli(),
		Msg:        getMessageByLangAndCode(lang, msgCode, detail...),
	}
	info := new(model.OperationLogInfo)
	info.Convert(logInfo)
	return sendToMQ(info)
	//return sendToGrpcServer(logInfo)

}

func sendToGrpcServer(info *proto.OperationLogInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithInsecure())
	if err != nil {
		newlog.Logger.Errorf("failed to get conn from endpoint[%s], err info: %+v\n", "127.0.0.1:10000", err)
		return err
	}
	defer conn.Close()
	client := proto.NewOperationLogClient(conn)
	_, err = client.AddOperationLog(ctx, info)
	if err != nil {
		return err
	}
	return nil
}

func sendToMQ(info *model.OperationLogInfo) error {
	data, _ := json.Marshal(info)

	return producer.Publish("operation", data)
}

func init() {
	initProducer()
}

var producer *nsq.Producer

func initProducer() error {
	var err error
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		fmt.Printf("create producer failed, err:%v\n", err)
		return err
	}

	return nil
}
