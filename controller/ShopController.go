package controller

import (
	"awesomeProject1/service"
	"awesomeProject1/tool"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ShopController struct {
}

//shop模块的路由解析
func (sc *ShopController) Router(app *gin.Engine) {
	app.GET("/api/shops", sc.GetShopList)
}

//获取商铺列表
func (sc *ShopController) GetShopList(context *gin.Context) {
	//接收参数
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")

	//参数条件检测
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34"
		latitude = "40.34"
	}
	fmt.Println(longitude, longitude)

	//传递参数到service层进行逻辑编写
	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude)
	//判断获取数据结果集的条数
	if len(shops) != 0 {
		tool.Success(context, shops)
		return
	}
	tool.Failed(context, "暂未获取到商户信息")
}
