package web_im

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/typed"
	"github.com/ricky97gr/homeOnline/internal/webservice/database/redis"
	client_manager "github.com/ricky97gr/homeOnline/internal/webservice/im_server/client"
	message_receiver "github.com/ricky97gr/homeOnline/internal/webservice/im_server/message"
)

const offLineDuration = 30

func AddWebSocketClient(uid, userName string, c *websocket.Conn) {
	rs, err := redis.GetRedisClient()
	if err != nil {
		newlog.Logger.Errorf("failed to get redis client, err: %+v\n", err)
		// return
	}
	client := &typed.WebSocketClient{
		Client:   c,
		UserName: "",
	}
	//没找到client，不存在发送广播消息
	if rs.Get(uid).Err() != nil {
		//发送广播消息
		msg := &typed.MessageInfo{
			FromUID:     "",
			FromName:    "",
			ToUID:       "",
			Type:        typed.SystemBroadCast,
			Context:     fmt.Sprintf("用户%s上线", userName),
			GroupID:     "",
			ContextType: 0,
			Time:        time.Now().UnixMilli(),
		}
		message_receiver.SetMessage2Chan(msg)

		go func() {
			for {
				mt, message, err := c.ReadMessage()
				if err != nil {
					newlog.Logger.Errorf("read: %+v\n", err)
					break
				}

				msg := &typed.MessageInfo{}
				err = json.Unmarshal(message, msg)
				if err != nil {
					newlog.Logger.Errorf("failed to Unmarshal message, err: %+v\n", err)
					continue
				}

				err = message_receiver.SetMessage2Chan(msg)
				if err != nil {
					busymsg := &typed.MessageInfo{
						FromUID:     "",
						FromName:    "",
						ToUID:       "",
						Type:        typed.SystemIsBusy,
						Context:     "系统忙，请稍后再试",
						GroupID:     "",
						ContextType: 0,
						Time:        time.Now().UnixMilli(),
					}
					data, err := json.Marshal(busymsg)
					if err != nil {
						continue
					}
					err = c.WriteMessage(mt, data)
					if err != nil {
						newlog.Logger.Errorf("failed to send system busy message to uid: %+v\n", err)
					}
				}

			}
		}()
	}
	client_manager.AddClient(uid, client)

}
