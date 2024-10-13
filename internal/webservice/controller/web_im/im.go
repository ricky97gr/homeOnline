package web_im

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	// "github.com/ricky97gr/homeOnline/internal/pkg/response"
	"github.com/ricky97gr/homeOnline/internal/webservice/service/web_im"
)

func ReceiveClientComm(ctx *gin.Context) {
	upgrader := websocket.Upgrader{
		HandshakeTimeout: time.Second * 10,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		// 解决跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{ctx.Request.Header.Get("Sec-Websocket-Protocol")},
	}
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		newlog.Logger.Errorf("failed to generate websocket connection, err:%+v\n", err)
		return
	}

	userID := ctx.Request.Header.Get("userID")
	userName := ctx.Request.Header.Get("userName")
	fmt.Printf("uid: %s, userName: %s\n", userID, userName)
	if userID == "" {
		newlog.Logger.Errorf("failed to get uuid from header\n")
		return
	}
	web_im.AddWebSocketClient(userID, userName, c)

}
