package main

import (
	"fmt"

	"github.com/fananchong/goredis"
	"github.com/gomodule/redigo/redis"
)

func testSentinel() {
	option := goredis.NewDefaultOption()
	option.Type = goredis.Sentinel
	addrs := []string{"192.168.1.4:46379", "192.168.1.4:46380", "192.168.1.4:46381"}
	dbName := "mysentinel"
	db, err0 := goredis.NewClient(dbName, addrs, option)
	if err0 != nil {
		fmt.Println(err0)
		return
	}
	_, err1 := db.Do("SET", "b", "12345")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	b, err2 := redis.Int(db.Do("GET", "b"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("b =", b)
}
