package store

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
	"time"
)

func NewPool(redisHost string, redisPort string, redisUser string, redisPassword string) *redis.Pool {
	address := fmt.Sprintf("%s:%s", redisHost, redisPort)
	user := redis.DialUsername(redisUser)
	password := redis.DialPassword(redisPassword)
	connectionPool := &redis.Pool{
		MaxIdle:     5,
		MaxActive: 10,
		Wait: true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address, user, password)
			if err != nil {
				log.Panic().Err(err).Msg(err.Error())
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return connectionPool
}