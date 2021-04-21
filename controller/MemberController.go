package controller

import (
	"awesomeProject1/service"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
}

//http://127.0.0.1:8090/api/sendcode?phone=13523419148
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//发送验证码
	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "参数解析失败",
		})
		return
	}

	//实例化服务层MemberService
	//service是上面import引入的service包
	ms := service.MemberService{}
	//调用发送短信的服务层方法
	isSend := ms.Sendcode(phone)
	if isSend {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  "发送成功",
		})
		return
	}
	context.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  "发送失败",
	})

}
