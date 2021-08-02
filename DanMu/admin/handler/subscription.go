/*
@Title : subscription
@Description :
@Author : 谭靖渝
@Update : 2021/8/2 14:29
*/
package handler

import (
	"DanMu/pkg/common"
	"DanMu/pkg/db/redis"
	"DanMu/ws"
	"encoding/json"
	"fmt"
	Redis "github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"time"
)

func KeepReceiving(con *ws.Client) {
	common.S.Client = redis.Instance().PubSubConn()
	go func() {
		for {
			switch res := common.S.Client.Receive().(type) {
			case Redis.Message:
				jsonMessage, _ := json.Marshal(&ws.Message{Content: string(res.Data)})
				//将消息放入该ws客户端的send管道中
				con.Send <- jsonMessage
			case Redis.Subscription:
				fmt.Printf("%s: %s %d\n", res.Channel, res.Kind, res.Count)
			case error:
				break
			}
		}
	}()
}

func Subs(channel interface{}, con *ws.Client) {
	KeepReceiving(con)
	err := common.S.Subscribe(channel)
	if err != nil {
		logrus.Debug(err)
	}
	for {
		time.Sleep(1 * time.Second)
	}
}
