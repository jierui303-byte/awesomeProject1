package main

import (
	"awesomeProject1/controller"
	"awesomeProject1/tool"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
)

func main() {

	//自动加载config配置文件
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	//实例化数据库 初始化orm对象
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	//实例化redis配置
	tool.InitRedisStore()

	app := gin.Default()

	//设置全局跨域访问，调用中间件
	app.Use(Cors())

	//集成session-调用session初始化
	tool.InitSession(app)

	//路由调用
	registerRouter(app)

	app.Run(cfg.APPHost + ":" + cfg.APPPort)
}

//路由设置-注册路由
func registerRouter(router *gin.Engine) {
	//hello控制器
	new(controller.HelloController).Router(router)

	//member控制器
	new(controller.MemberController).Router(router)

	//foodCategory食品类控制器
	new(controller.FoodCategoryController).Router(router)
}

//跨域访问中间件：cross origin resource share
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		//拿到http请求头里的origin值
		origin := context.Request.Header.Get("origin")

		//创建一个切片数组
		var headerKeys []string
		for key, _ := range context.Request.Header {
			//往切片里面增加数据
			headerKeys = append(headerKeys, key)
		}

		//切片数组拼接成字符串
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") //设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}

		context.Next()
	}
}
