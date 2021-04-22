package controller

import (
	"awesomeProject1/param"
	"awesomeProject1/service"
	"awesomeProject1/tool"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	//解析接口地址
	engine.GET("/api/sendcode", mc.sendSmsCode)
	engine.OPTIONS("/api/login_sms", mc.smsLogin)
}

//发送短信验证码方法 http://127.0.0.1:8090/api/sendcode?phone=13523419148
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//发送验证码
	phone, exist := context.GetQuery("phone")
	if !exist {
		//context.JSON(200, map[string]interface{}{
		//	"code": 0,
		//	"msg":  "参数解析失败",
		//})
		tool.Failed(context, "参数解析失败")
		return
	}

	//实例化服务层MemberService
	//service是上面import引入的service包
	ms := service.MemberService{}
	//调用发送短信的服务层方法
	isSend := ms.Sendcode(phone)
	if isSend {
		//context.JSON(200, map[string]interface{}{
		//	"code": 1,
		//	"msg":  "发送成功",
		//})
		tool.Success(context, "发送成功")
		return
	}
	//context.JSON(200, map[string]interface{}{
	//	"code": 0,
	//	"msg":  "发送失败",
	//})
	tool.Failed(context, "发送失败")
}

//手机号+短信 登录的方法
func (mc *MemberController) smsLogin(context *gin.Context)  {
	var smsLoginParam param.SmsLoginParam
	//参数解析-调用tool.Decode函数进行body参数解析
	err := tool.Decode(context.Request.Body, &smsLoginParam)
	if err != nil {
		//context.JSON(200, map[string]interface{}{
		//	"code" : 0,
		//	"msg" : "参数解析失败",
		//})
		tool.Failed(context, "参数解析失败")
		return
	}

	//完成手机+验证码登录
	us := service.MemberService{}
	//调用服务下的登录方法
	member := us.SmsLogin(smsLoginParam)
	if member != nil{
		tool.Success(context, member)
		return
	}

	tool.Failed(context, "登录失败")
}