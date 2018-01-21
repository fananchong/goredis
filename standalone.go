package goredis

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
)

type StandaloneClient struct {
	cli    *redis.Pool
	dbName string
}

func NewStandaloneClient(dbName string, addr string, option *Option) *StandaloneClient {
	cli := &StandaloneClient{}
	cli.Init(dbName, addr, option)
	return cli
}

func (this *StandaloneClient) Init(dbName string, addr string, option *Option) {
	this.dbName = dbName
	this.cli = &redis.Pool{
		MaxIdle:     option.PoolMaxIdle,
		MaxActive:   option.PoolMaxActive,
		Wait:        true,
		IdleTimeout: option.PoolIdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if option.Password != "" {
				if _, err := c.Do("AUTH", option.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", option.DBIndex); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func (this *StandaloneClient) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := this.cli.Get()
	defer conn.Close()
	if conn != nil {
		return conn.Do(commandName, args...)
	} else {
		return nil, errors.New("[redis] Can't get redis client!")
	}
}
