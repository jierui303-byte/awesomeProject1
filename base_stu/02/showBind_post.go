package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

/**
	post请求 注册表单form数据提交参数接收方式
 */
func main() {
	engine := gin.Default()

	engine.POST("/register", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		var register Register
		err := context.ShouldBind(&register)
		if err != nil{
			//日志记录
			log.Fatal(err.Error())
			return
		}

		//打印输出
		fmt.Println(register.UserName)
		fmt.Println(register.Phone)
		fmt.Println(register.Password)

		//字符串拼接+
		context.Writer.Write([]byte(register.UserName + "注册成功了"))
	})

	//自定义监听端口
	engine.Run(":8081")
}

type Register struct {
	UserName string `form:"name"`
	Phone string `form:"iphone"`
	Password string `form:"password"`
}
