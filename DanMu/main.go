/*
@Title : main
@Description :
@Author : 谭靖渝
@Update : 2021/8/1 23:06
*/
package main

import (
	"DanMu/admin"
	"DanMu/ws"
	"KFCChat/pkg/db/redis"
	"fmt"
)

func main() {
	if err := redis.Instance().Initial(redis.RedisOptions{
		Addr: "127.0.0.1:6379",
	}); err != nil {
		panic(fmt.Errorf("Fatal error redis initial:%s \n", err))
	}
	r:=admin.NewRouter()
	r.Run("127.0.0.1:9090")
	go ws.Manager.Start()
	defer redis.Instance().Close()
}
