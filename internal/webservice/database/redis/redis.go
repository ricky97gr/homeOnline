package redis

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/ricky97gr/homeOnline/internal/conf"
)

const (
	RedisAuth = iota
)

var c *redis.Client

func GetRedisClient() (*redis.Client, error) {
	if c == nil {
		config := conf.GetConfig()
		return InitRedis(fmt.Sprintf("%s:%d", config.Redis.IP, config.Redis.Port), config.Redis.Password)
	}
	_, err := c.Ping().Result()
	if err != nil {
		config := conf.GetConfig()
		c, err = InitRedis(fmt.Sprintf("%s:%d", config.Redis.IP, config.Redis.Port), config.Redis.Password)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func InitRedis(addr, passwd string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		DB:       0,
		Addr:     addr,
		Password: passwd,
		PoolSize: 100,
	})
	_, err := client.Ping().Result()
	if err == nil {
		//TODO: 清除所有的redis token
		//client.FlushDB()
	}
	return client, err
}
