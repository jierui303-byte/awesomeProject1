package controller

import (
	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

//路由请求和解析函数绑定
func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.Hello)
}

//解析hello
func (hello *HelloController) Hello(context *gin.Context) {
	context.JSON(200, map[string]interface{}{
		"code": 1,
		"msg":  "你好，我是jierui303",
	})
}
