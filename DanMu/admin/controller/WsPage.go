/*
@Title : WsPage
@Description :
@Author : 谭靖渝
@Update : 2021/8/1 16:43
*/
package controller

import (
	"DanMu/admin/handler"
	"DanMu/pkg/common"
	"DanMu/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

//升级ws服务
func WsPage(c *gin.Context) {
	ID, ok := c.Params.Get("uid")
	if !ok{
		c.JSON(http.StatusOK, common.Response(common.ParmaError,"参数错误",""))
	}
	channel, ok := c.Params.Get("vid")
	if !ok{
		c.JSON(http.StatusOK, common.Response(common.ParmaError,"参数错误",""))
	}
	//升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, common.Response(common.WsError,"ws服务错误",err))
	}
	// 创建连接
	client := &ws.Client{ID: ID, Socket: conn, Send: make(chan []byte),Channel: channel}
	//放入注册管道，相当于进入视频
	ws.Manager.Register <- client
	//开启该视频的订阅

	go handler.Subs(channel,client)
	go client.Write()

	c.JSON(http.StatusOK, common.Response(common.OK,"成功",err))
}

