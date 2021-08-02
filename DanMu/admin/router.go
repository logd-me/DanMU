/*
@Title : router
@Description :
@Author : 谭靖渝
@Update : 2021/8/1 16:42
*/
package admin

import (
	"DanMu/admin/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	//设置默认引擎
	r := gin.Default()
	r.GET("/videos/:vid/users/:uid/ws",controller.WsPage)
	r.POST("/send",controller.SendMassage)
	return r
}
