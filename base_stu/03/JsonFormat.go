package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
	map类型 ，struct结构体 作为json格式返回
 */
func main() {
	engine := gin.Default()
	engine.GET("/hellojson", func(context *gin.Context) {
		fullPath := "请求路径：" + context.FullPath()
		fmt.Println(fullPath)

		//map类型数据 作为json返回
		context.JSON(200, map[string]interface{}{
			"code" : 1,
			"message" : "ok",
			"data" : fullPath,
		})
	})

	engine.GET("/jsonstruct", func(context *gin.Context) {
		fullPath := "请求路径：" + context.FullPath()
		fmt.Println(fullPath)

		//struct结构体 作为json返回
		resp := Response{Code: 1, Message: "ok", Data: fullPath}
		context.JSON(200, &resp)
	})
	engine.Run(":8080")
}

type Response struct {
	Code int
	Message string
	Data interface{}
}