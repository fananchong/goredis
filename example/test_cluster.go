package main

import (
	"fmt"

	"github.com/fananchong/goredis"
	"github.com/gomodule/redigo/redis"
)

func testCluster() {
	option := goredis.NewDefaultOption()
	option.Type = goredis.Cluster
	addrs := []string{"192.168.1.4:39379", "192.168.1.4:39381", "192.168.1.4:39383"}
	db, err0 := goredis.NewClient("", addrs, option)
	if err0 != nil {
		fmt.Println(err0)
		return
	}
	_, err1 := db.Do("SET", "c", "12345")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	c, err2 := redis.Int(db.Do("GET", "c"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("c =", c)
}
