package typed

import (
	"github.com/gorilla/websocket"
)

type MessageInfo struct {
	FromUID     string      `json:"fromUID"`
	FromName    string      `json:"fromName"`
	ToUID       string      `json:"toUID"`
	Type        MessageType `json:"type"`
	Context     string      `json:"context"`
	GroupID     string      `json:"groupID"`
	ContextType int         `json:"contextType"`
	Time        int64       `json:"time"`
}

type MessageType int

const (
	//系统广播消息不存mongo
	SystemBroadCast = iota + 1
	// 系统忙 不存mongo
	SystemIsBusy
	// 组消息
	GroupMessage
	// 私聊消息
	ToPeopleMessage
)

type WebSocketClient struct {
	Client   *websocket.Conn
	UserName string
}
