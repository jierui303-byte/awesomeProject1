package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//传参数hello?name=jierui303&age=30&classes=嵌入式
	r.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		//通过事先定义结构体来接收URL传递的参数
		var student Student
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		//打印结构体内的数据
		fmt.Println(student.Name)
		fmt.Println(student.Age)
		fmt.Println(student.Classes)

		//context.JSON(200, gin.H{
		//	"message" : "hello, jierui303! \n",
		//})

		context.Writer.Write([]byte("Hello, <h1 style='color:red;'>" + student.Name + "</h1>"))
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

type Student struct {
	Name string `form:"name"`
	Age int `form:"age"`
	Classes string `form:"classes"`
}