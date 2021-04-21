package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
	四个不同的路由地址，但是前缀/user一致，可分为一个路由组
	/user/register   用户注册
	/user/login   用户登录
	/user/info  用户信息
	/user/1001  删除1001用户

	把路由中的匿名函数单独摘出来重命名后，直接调用函数，这样，main函数中代码就会缩减很多
 */
func main() {
	engine := gin.Default()

	//路由前缀定义
	userRouterGroup := engine.Group("/user")

	//用户注册
	userRouterGroup.POST("/register", registerHandle)

	//用户登录
	userRouterGroup.POST("/login", loginHandle)

	//用户信息
	userRouterGroup.GET("/info", infoHandle)

	//用户删除
	userRouterGroup.DELETE("/:id/:new_id", deleteUserHandle)

	engine.Run()
}

func registerHandle(context *gin.Context) {
	fullPath := "用户注册 请求地址：" + context.FullPath()
	fmt.Println(fullPath)

	context.Writer.WriteString(fullPath)
}

func loginHandle(context *gin.Context) {
	fullPath := "用户登录 请求地址：" + context.FullPath()
	fmt.Println(fullPath)

	context.Writer.WriteString(fullPath)
}

func infoHandle(context *gin.Context) {
	fullPath := "用户信息 请求地址：" + context.FullPath()
	fmt.Println(fullPath)

	context.Writer.WriteString(fullPath)
}

func deleteUserHandle(context *gin.Context) {
	fullPath := "用户删除 请求地址：" + context.FullPath()
	fmt.Println(fullPath)

	//获取参数
	userId := context.Param("id")
	newId := context.Param("new_id")
	fmt.Println("删除用户ID：" + userId)

	context.Writer.WriteString(fullPath + " " + userId + "--" + newId)
}