package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
	响应结果
	write
	writerString
 */
func main() {
	engine := gin.Default()

	engine.GET("/hellobyte", func(context *gin.Context) {
		fullPath := "请求路径：" + context.FullPath()
		fmt.Println(fullPath)

		//write
		context.Writer.Write([]byte(fullPath))
	})
	
	engine.GET("/hellostring", func(context *gin.Context) {
		fullPath := "请求路径：" + context.FullPath()
		fmt.Println(fullPath)

		//writerString
		context.Writer.WriteString(fullPath)
	})

	engine.Run(":8090")
}
