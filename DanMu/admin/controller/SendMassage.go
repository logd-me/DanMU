/*
@Title : SendMassage
@Description :
@Author : 谭靖渝
@Update : 2021/8/1 16:57
*/
package controller

import (
	"DanMu/admin/handler"
	"DanMu/pkg/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMassage(c *gin.Context)  {
	var s handler.MassageServer
	if err := c.ShouldBindJSON(&s);err==nil{
		c.JSON(http.StatusOK, s.Send())
	}else {
		c.JSON(http.StatusOK, common.Response(common.OperateError, "操作失败", err))
	}
}
