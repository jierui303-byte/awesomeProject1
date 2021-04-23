package service

import (
	"awesomeProject1/dao"
	"awesomeProject1/model"
	"awesomeProject1/param"
	"awesomeProject1/tool"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/micro/go-micro/v2/logger"
)

//定义服务层
type MemberService struct {
}

//头像上传操作
func (ms *MemberService) UploadAvator(userId int64, fileName string) string {
	memberDao := dao.MemberDao{tool.DbEngine}
	result := memberDao.UpdateMemberAvator(userId, fileName)
	//更新失败
	if result == 0 {
		return ""
	}

	return fileName
}

//定义 手机+密码+验证码登录的方法
func (ms *MemberService) Login(name string, password string) *model.Member {
	//两种情况：用户存在/用户为新用户
	//1.使用用户名+密码查询用户信息 如果存在，直接返回
	md := dao.MemberDao{tool.DbEngine}
	member := md.Query(name, password)
	if member.Id != 0 {
		return member
	}

	//2.不存在，用户作为新用户插入新增，再返回
	user := model.Member{}
	user.UserName = name
	user.Password = tool.EncoderSha256(password) //密码加密
	user.RegisterTime = time.Now().Unix()

	//执行插入新用户
	result := md.InsertMember(user)
	user.Id = result

	return &user
}

//定义手机+验证码实现登录的方法
func (ms *MemberService) SmsLogin(loginParam param.SmsLoginParam) *model.Member {
	//完成用户登录成功状态修改过程

	//1.获取到手机号和验证码

	//2.验证手机号+验证码是否正确
	md := dao.MemberDao{}
	sms := md.ValidateSmsCode(loginParam.Phone, loginParam.Code)
	if sms.Id == 0 {
		return nil
	}

	//3.根据手机号member表中查询记录
	member := md.QueryByPhone(loginParam.Phone)
	if member.Id != 0 {
		//代表会员存在
		return member
	}

	//4.新建一个member记录，并保存
	user := model.Member{}
	user.UserName = "wxxxxx"
	user.Mobile = "13523419148"
	user.RegisterTime = time.Now().Unix()

	user.Id = md.InsertMember(user)

	return &user
}

//定义发送短信的方法
func (ms *MemberService) Sendcode(phone string) bool {
	//获取全局config数据
	config := tool.GetConfig()

	//1. 产生一个验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))

	//2. 调用阿里云sdk, 完成发送
	client, err := dysmsapi.NewClientWithAccessKey(config.Sms.RegionId, config.Sms.AppKey, config.Sms.AppSecret)
	if err != nil {
		//logger记录错误日志
		logger.Error(err.Error())
		return false
	}

	//拼装数据
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.Sms.SignName
	request.TemplateCode = config.Sms.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)

	//3. 接收返回结果，并判断发送状态
	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	//短信验证码发送成功
	if response.Code == "ok" {
		//将验证码保存到数据库中
		smsCode := model.SmsCode{
			Phone:      phone,
			Code:       code,
			BizId:      response.BizId,
			CreateTime: time.Now().Unix(),
		}

		//实例化数据库dao层进行数据库数据写入操作InsertCode
		//tool.DbEngine实际上是orm实例化对象
		memberDao := dao.MemberDao{tool.DbEngine}
		result := memberDao.InsertCode(smsCode)
		if result < 0 {
			return false
		}

		return true
	}

	return false
}
