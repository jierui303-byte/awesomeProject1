package main

import (
	"awesomeProject1/controller"
	"awesomeProject1/tool"

	"github.com/gin-gonic/gin"
)

func main() {

	//自动加载config配置文件
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	app := gin.Default()

	//路由调用
	registerRouter(app)

	app.Run(cfg.APPHost + ":" + cfg.APPPort)
}

//路由设置
func registerRouter(router *gin.Engine) {
	//hello控制器
	new(controller.HelloController).Router(router)

	//member控制器
	new(controller.MemberController).Router(router)
}
