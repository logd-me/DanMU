/*
@Title : Massage
@Description :
@Author : 谭靖渝
@Update : 2021/8/2 14:56
*/
package handler

import (
	"DanMu/history"
	"DanMu/pkg/common"
	"DanMu/pkg/db/redis"
)

type MassageServer struct {
	Channel string `json:"channel"`
	Content string `json:"content"`
}

func (service *MassageServer) Send() common.Result {
	//推送消息
	if err := redis.Instance().Publish(service.Channel, service.Content); err != nil {
		return common.Response(common.PushError, "推送失败", err)
	}
	msg, err := history.GetInfo(service.Channel)
	if err != nil {
		return common.Response(common.CacheError, "缓存错误", err)
	}
	//装配弹幕。
	msg.DanMu = append(msg.DanMu, service.Content)
	//将弹幕压入redis缓存
	err  = history.SetInfo(service.Channel, msg)
	if err != nil {
		return common.Response(common.CacheError, "缓存错误", err)
	}
	return common.Response(common.OK,"成功","")
}
