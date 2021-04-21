package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
	1：渲染html页面及变量渲染
	2：图片等静态资源文件加载方式

	配置正确才能正常访问：
		Files: /Users/mac/go/awesomeProject1/base_stu/03/HtmlFormat.go
		Working Directory: /Users/mac/go/awesomeProject1/base_stu/03
 */
func main() {
	engine := gin.Default()

	//设置html目录
	engine.LoadHTMLGlob("./html/*")

	//设置静态资源文件路径 浏览器url路径和文件存储目录进行绑定
	engine.Static("/img", "./img")

	engine.GET("/hellohtml", func(context *gin.Context) {
		fullPath := "请求路径：" + context.FullPath()
		fmt.Println(fullPath)

		//一：加载渲染html页面
		//context.HTML(http.StatusOK, "index.html", nil)

		//二：加载渲染html页面并传递变量
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title" : "Gin学习html渲染",
			"fullPath" : fullPath,
		})
	})
	engine.Run()
}
