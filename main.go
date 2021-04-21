package main

import (
	"awesomeProject1/controller"
	"awesomeProject1/tool"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
)

func main() {

	//自动加载config配置文件
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	//初始化orm对象
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		logger.Error(err.Error())
		return
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
