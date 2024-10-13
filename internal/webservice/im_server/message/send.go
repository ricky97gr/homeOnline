package message

import (
	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"

	client_manager "github.com/ricky97gr/homeOnline/internal/webservice/im_server/client"
)

func broadCastMessage(data []byte) {
	clients := client_manager.ListClient()
	for _, c := range clients {
		err := c.Client.WriteMessage(1, data)
		newlog.Logger.Debugf("server write message info: %+v\n", err)
		if err != nil {
			newlog.Logger.Errorf("write: %+v\n", err)
		}
	}
}

func sendMsgToUser(uid string, data []byte) error {
	toC, err := client_manager.FindClientByUid(uid)
	if err != nil {
		return err
	}
	err = toC.Client.WriteMessage(1, data)
	if err != nil {
		newlog.Logger.Errorf("write: %+v\n", err)
		return err
	}
	return nil
}

func sendMsgToGroupMember(groupUID, fromUID string, data []byte) {
	users := client_manager.GetGroupMemberByGroupUID(groupUID)
	for _, u := range users {
		if fromUID == u {
			continue
		}
		toC, _ := client_manager.FindClientByUid(u)
		err := toC.Client.WriteMessage(1, data)
		if err != nil {
			newlog.Logger.Errorf("write: %+v\n", err)
			return
		}
	}
}
