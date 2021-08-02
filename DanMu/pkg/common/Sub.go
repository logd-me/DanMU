/*
@Title : sub
@Description :
@Author : 谭靖渝
@Update : 2021/8/2 17:40
*/
package common

import (
	Redis "github.com/gomodule/redigo/redis"
)

type Subscriber struct {
	Client Redis.PubSubConn
}

var  S  Subscriber

func (c *Subscriber) Subscribe(channel interface{}) (err error) {
	err = c.Client.Subscribe(channel)
	if err != nil {
		return
	}
	return
}
func (c *Subscriber) UnSubscribe(channel interface{}) (err error) {
	err = c.Client.Unsubscribe(channel)
	if err != nil {
		return
	}
	return
}

