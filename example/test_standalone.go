package main

import (
	"fmt"

	"github.com/fananchong/goredis"
	"github.com/garyburd/redigo/redis"
)

func testStandalone() {
	option := goredis.NewDefaultOption()
	option.Type = goredis.Standalone
	addrs := []string{"192.168.1.4:16379"}
	db := goredis.NewClient("", addrs, option)
	_, err1 := db.Do("SET", "a", "12345")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	a, err2 := redis.Int(db.Do("GET", "a"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("a =", a)
}
