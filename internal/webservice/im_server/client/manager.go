package client

import (
	"sync"
	"time"

	"github.com/ricky97gr/homeOnline/internal/pkg/newlog"
	"github.com/ricky97gr/homeOnline/internal/pkg/typed"
	"github.com/ricky97gr/homeOnline/internal/webservice/database/redis"
)

type clientManager struct {
	clients   sync.Map
	group2UID sync.Map
}

const offLineDuration = 30

var manager = &clientManager{
	clients:   sync.Map{},
	group2UID: sync.Map{},
}

func AddClient(uid string, c *typed.WebSocketClient) {
	manager.addClient(uid, c)
	err := addClientToRedis(uid, c)
	if err != nil {
		newlog.Logger.Errorf("failed to store client to redis, err: %+v\n", err)
	}
}

func DeleteClient(uid string) {
	manager.delClient(uid)
	err := deleteClientFromRedis(uid)
	if err != nil {
		newlog.Logger.Errorf("failed to delete client from redis, err: %+v\n", err)
	}
}

func SetGroupMember(groupUID string, users []string) {
	manager.group2UID.Store(groupUID, users)
}

func GetGroupMemberByGroupUID(groupUID string) []string {
	if v, ok := manager.group2UID.Load(groupUID); ok {
		return v.([]string)
	}
	return []string{}
}

func FindClientByUid(uid string) (*typed.WebSocketClient, error) {
	return manager.getClient(uid)
}

func ListClient() []*typed.WebSocketClient {
	return manager.listClient()
}

func getClientFromRedis(uid string) (*typed.WebSocketClient, error) {
	rs, err := redis.GetRedisClient()
	if err != nil {
		return nil, err
	}
	c := &typed.WebSocketClient{}
	err = rs.Get(uid).Scan(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func addClientToRedis(uid string, c *typed.WebSocketClient) error {
	rs, err := redis.GetRedisClient()
	if err != nil {
		return err
	}
	return rs.Set(uid, c, offLineDuration*time.Minute).Err()
}

func deleteClientFromRedis(uid string) error {
	rs, err := redis.GetRedisClient()
	if err != nil {
		return err
	}
	return rs.Del(uid).Err()
}

func (m *clientManager) addClient(uid string, c *typed.WebSocketClient) {
	m.clients.Store(uid, c)
}

func (m *clientManager) delClient(uid string) {
	m.clients.Delete(uid)
}

func (m *clientManager) listClient() []*typed.WebSocketClient {
	var clients []*typed.WebSocketClient
	m.clients.Range(func(key, value any) bool {
		clients = append(clients, value.(*typed.WebSocketClient))
		return true
	})
	return clients

}

func (m *clientManager) getClient(uid string) (*typed.WebSocketClient, error) {
	c, ok := m.clients.Load(uid)
	if ok {
		return c.(*typed.WebSocketClient), nil

	}
	cli, err := getClientFromRedis(uid)
	if err != nil {
		return nil, err
	}
	m.addClient(uid, cli)
	return cli, nil

}
