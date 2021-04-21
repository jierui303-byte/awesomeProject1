package controller

import (
	"awesomeProject1/dao"
	"awesomeProject1/model"
	"awesomeProject1/tool"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

//路由请求和解析函数绑定
func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.Hello)
}

//解析hello
func (hello *HelloController) Hello(context *gin.Context) {

	//尝试写入数据库
	smsCode := model.SmsCode{
		Phone:      "122222",
		Code:       "1234",
		BizId:      "78",
		CreateTime: time.Now().Unix(),
	}

	//实例化数据库dao层进行数据库数据写入操作InsertCode
	//tool.DbEngine实际上是orm实例化对象
	memberDao := dao.MemberDao{tool.DbEngine}
	result := memberDao.InsertCode(smsCode)
	if result < 0 {
		fmt.Println("测试写入失败")
	}

	fmt.Println("写入成功, result: ", result)

	context.JSON(200, map[string]interface{}{
		"code": 1,
		"msg":  "你好，我是jierui303",
	})
}
