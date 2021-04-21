package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
	系统默认提供的中间件：logger中间件 + Recovery中间件
	自定义中间件：

	context.Next函数可以将中间件代码的执行顺序一分为二， next函数调用之前的代码在请求处理之前，
	当程序执行到context.next时，会中断向下执行，转而先去执行具体的业务逻辑，执行完业务逻辑处理函数之后，
	程序会再次回到context.next处，继续执行中间件后续的代码。
	即：前置中间件，后置中间件

*/
func main() {
	engine := gin.Default()

	//一：使用自定义中间件：全局路由使用方式
	engine.Use(RequestInfos())

	//路由
	engine.GET("/query", func(context *gin.Context) {

		fmt.Println("执行query方法。。。")

		context.JSON(404, map[string]interface{}{
			"code" : 1,
			"msg" : context.FullPath(),
		})
	})

	//二：单独对某一个路由使用中间件的方式
	engine.GET("/hellos", RequestHellos(), func(context *gin.Context) {
		//todo
	})

	engine.Run()
}

//自定义打印请求信息的中间件
func RequestInfos() gin.HandlerFunc{
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method

		fmt.Println("请求path：" + path)
		fmt.Println("请求method：" + method)
		fmt.Println("状态码-前：", context.Writer.Status())

		context.Next() // 此代码之前代码是前置中间件代码   代码之后是后置中间件代码

		fmt.Println("状态码-后：", context.Writer.Status())
	}
}

//自定义打印请求信息的中间件
func RequestHellos() gin.HandlerFunc{
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method

		fmt.Println("请求path：" + path)
		fmt.Println("请求method：" + method)
	}
}