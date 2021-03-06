package controller

import (
	"awesomeProject1/model"
	"awesomeProject1/param"
	"awesomeProject1/service"
	"awesomeProject1/tool"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	//解析接口地址
	engine.GET("/api/sendcode", mc.sendSmsCode)
	engine.OPTIONS("/api/login_sms", mc.smsLogin)     //手机验证码登录注册
	engine.GET("/api/captcha", mc.captcha)            //生成验证码图片
	engine.POST("/api/vertifycha", mc.vertifyCaptcha) //校验验证码图片是否正确
	//手机号+密码+验证码登录
	engine.POST("/api/login_pwd", mc.nameLogin)

	//头像上传[上传到服务器本地]
	engine.POST("/api/upload/avator", mc.uploadAvator)
}

//头像上传
func (mc *MemberController) uploadAvator(context *gin.Context) {
	//1.解析上传的参数：file
	userId := context.PostForm("user_id") //从postfrom接收参数
	fmt.Println(userId)
	file, err := context.FormFile("avator") //从formfile接收图片数据
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	//2.判断user_id对应的用户是否已经登录
	sess := tool.GetSession(context, "user_"+userId) //获取session
	if sess == nil {
		tool.Failed(context, "参数不合法")
		return
	}

	var member model.Member
	//解析参数到member对象里面
	json.Unmarshal(sess.([]byte), &member)

	//3.【第一种方式】file保存到本地
	fileName := "./uploadfile" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename //文件名生成规则
	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Failed(context, "头像更新失败")
		return
	}

	//3.【第二种方式】把file上传到本地服务器上的fastDFS分布式文件存储系统中【封装一个上传fastDFS工具】
	fileId := tool.UploadFileToDFS(fileName)
	if fileId != "" {
		//删除本地uplodfile下的文件
		os.Remove(fileName)

		//4.将保存后的文件本地路径 保存到用户表中的头像字段
		//http://localhost:8090/static/.../devie.png
		memberService := service.MemberService{}
		path := memberService.UploadAvator(member.Id, fileId)
		if path != "" {
			//tool.Success(context, "http://localhost:8090"+path)
			tool.Success(context, tool.FileServerAddr()+"/"+path)
			return
		}
	}

	//4.将保存后的文件本地路径 保存到用户表中的头像字段
	//http://localhost:8090/static/.../devie.png
	memberService := service.MemberService{}
	path := memberService.UploadAvator(member.Id, fileName[1:])
	if path != "" {
		//tool.Success(context, "http://localhost:8090"+path)
		tool.Success(context, tool.FileServerAddr()+"/"+path)
		return
	}
	//5.返回结果
	tool.Failed(context, "上传失败")
}

//手机号+密码+验证码登录
func (mc *MemberController) nameLogin(context *gin.Context) {
	//1.解析用户登录传递参数【去param里定义前端传递给我们参数的结构体】
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败1111")
		return
	}

	//2.验证验证码
	validate := tool.VertifyCaptcha(loginParam.Id, loginParam.Value)
	if !validate {
		tool.Failed(context, "验证码不正确，清重新验证")
		return
	}

	//3.登录
	ms := service.MemberService{}
	//去MemberService里封装一个登录的方法
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		//用户信息保存到session
		sess, _ := json.Marshal(member)
		err := tool.SetSession(context, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "登录失败")
			return
		}

		tool.Success(context, &member)
		return
	}

	tool.Failed(context, "登录失败")
}

//生成验证码
func (mc *MemberController) captcha(context *gin.Context) {
	//todo:生成验证码图片
	tool.GenerateCaptcha(context)
}

//验证验证码是否正确
func (mc *MemberController) vertifyCaptcha(context *gin.Context) {
	//接收客户端传输过来的参数
	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	result := tool.VertifyCaptcha(captcha.Id, captcha.VertifyValue)
	if result {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}
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
func (mc *MemberController) smsLogin(context *gin.Context) {
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
	if member != nil {
		//用户信息保存到session
		sess, _ := json.Marshal(member)
		err := tool.SetSession(context, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "登录失败")
			return
		}

		tool.Success(context, member)
		return
	}

	tool.Failed(context, "登录失败")
}
