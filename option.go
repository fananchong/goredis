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
	poolMaxIdle     int
	poolMaxActive   int
	poolIdleTimeout time.Duration
}

var DefaultOption = &Option{
	redisType:       Unknow,
	poolMaxIdle:     10,
	poolMaxActive:   0,
	poolIdleTimeout: 240 * time.Second,
}
