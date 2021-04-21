package service

import (
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

//定义
func (ms *MemberService) Sendcode(phone string) bool {
	//获取全局config数据
	config := tool.GetConfig()

	//1. 产生一个验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))

	//2. 调用阿里云sdk, 完成发送
	client, err := dysmsapi.NewClientWithAccessKey(config.Sms.RegionId, config.Sms.AppKey, config.Sms.AppSecret)
	if err != nil {
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
	if response.Code == "ok" {
		return true
	}

	return false
}
