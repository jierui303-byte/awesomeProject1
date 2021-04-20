package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

/**
	post请求 传递json格式的数据接收方式
	{
	  "name":"jierui303",
	  "age":23,
	  "sex":"女"
	}
 */
func main() {
	engine := gin.Default()
	engine.POST("/addStudent", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var person Person
		err := context.BindJSON(&person)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println("姓名：" , person.Name)
		//fmt.Println("年龄：" , person.Age)
		fmt.Println("性别：" , person.Sex)

		context.Writer.Write([]byte("添加记录：" + person.Name))
	})

	engine.Run(":8082")
}

type Person struct {
	Name string `form:"name"`
	Sex string `form:"sex"`
	Age string `form:"age"`
}
