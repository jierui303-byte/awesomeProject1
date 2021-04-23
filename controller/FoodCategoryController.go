package controller

import (
	"awesomeProject1/service"
	"awesomeProject1/tool"

	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	//获取食品种类路由
	engine.GET("/api/food_category", fcc.foodCategory)
}

//获取食品种类路由地址对应的控制器方法
func (fcc *FoodCategoryController) foodCategory(ctx *gin.Context) {
	//调用service层的方法获取食品种类信息
	foodCategoryService := &service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Failed(ctx, "食品种类数据获取失败")
		return
	}

	//转换格式-对图片地址进行服务器地址和端口进行拼接
	for _, category := range categories {
		if category.ImageUrl != "" {
			category.ImageUrl = tool.FileServerAddr() + "/" + category.ImageUrl
		}
	}

	tool.Success(ctx, categories)
}
