/*
@Title : common
@Description :
@Author : 谭靖渝
@Update : 2021/8/1 16:46
*/
package common



type Result struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	OK           = 200  //成功
	WsError      = 1000 //ws服务初始化失败
	ParmaError   = 1001 //参数错误
	PushError    = 1002 //推送失败
	CacheError   = 1003 //获取缓存失败
	OperateError = 1004 //操作失败
)

//通用返回方法
func Response(code uint32, message string, data interface{}) Result {
	result := generateResult(code, message, data)
	return result
}

func generateResult(code uint32, message string, data interface{}) Result {
	result := Result{
		Code: code,
		Msg:  message,
		Data: data,
	}
	return result
}
