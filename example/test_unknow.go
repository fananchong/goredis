package main

import (
	"fmt"

	"github.com/fananchong/goredis"
	"github.com/gomodule/redigo/redis"
)

func testUnknow() {

	fmt.Println("test cluster =========================================")
	{
		option := goredis.NewDefaultOption()
		addrs := []string{"192.168.1.4:39379", "192.168.1.4:39381", "192.168.1.4:39383"}
		db, err0 := goredis.NewClient("", addrs, option)
		if err0 != nil {
			fmt.Println(err0)
			return
		}
		_, err1 := db.Do("SET", "cc", "cc")
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		cc, err2 := redis.String(db.Do("GET", "cc"))
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		fmt.Println("cc =", cc)
	}

	fmt.Println("test sentinel =========================================")
	{
		option := goredis.NewDefaultOption()
		addrs := []string{"192.168.1.4:46379", "192.168.1.4:46380", "192.168.1.4:46381"}
		db, err0 := goredis.NewClient("mysentinel", addrs, option) //dbname: mysentinel
		if err0 != nil {
			fmt.Println(err0)
			return
		}
		_, err1 := db.Do("SET", "bb", "bb")
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		bb, err2 := redis.String(db.Do("GET", "bb"))
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		fmt.Println("bb =", bb)
	}

	fmt.Println("test standalone =========================================")
	{
		option := goredis.NewDefaultOption()
		addrs := []string{"192.168.1.4:16379"}
		db, err0 := goredis.NewClient("", addrs, option)
		if err0 != nil {
			fmt.Println(err0)
			return
		}
		_, err1 := db.Do("SET", "aa", "aa")
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		aa, err2 := redis.String(db.Do("GET", "aa"))
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		fmt.Println("aa =", aa)
	}
}
