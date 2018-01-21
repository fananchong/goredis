package goredis

import "time"

type RedisType int

const (
	Unknow RedisType = iota
	Standalone
	Sentinel
	Cluster
)

type Option struct {
	redisType       RedisType
	password        string
	dbIndex         int
	poolMaxIdle     int
	poolMaxActive   int
	poolIdleTimeout time.Duration
}

func NewDefaultOption() *Option {
	return &Option{
		redisType:       Unknow,
		password:        "",
		dbIndex:         0,
		poolMaxIdle:     10,
		poolMaxActive:   0,
		poolIdleTimeout: 240 * time.Second,
	}
}
