package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//定义常量枚举
const (
	SUCCESS int = 0 //操作成功
	FAILED int = 1 //操作失败
)

//封装公共 普通的成功响应方法
func Success(ctx *gin.Context, v interface{}){
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code" : SUCCESS,
		"msg" : "成功",
		"data" : v,
	})
}

//封装公共 普通的失败响应方法
func Failed(ctx *gin.Context, v interface{}){
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code" : FAILED,
		"msg" : v,
	})
}
