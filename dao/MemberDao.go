package dao

import (
	"awesomeProject1/model"
	"awesomeProject1/tool"
	"fmt"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/util/log"
)

//结构体绑定orm对象
type MemberDao struct {
	*tool.Orm
}

//根据用户名和密码查询用户信息
func (md *MemberDao) Query(name string, password string) *model.Member {
	var member model.Member

	password = tool.EncoderSha256(password)

	_, err := md.Where("user_name = ? and password = ? ", name, password).Get(&member)
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	return &member
}

//验证手机号和验证码是否存在
func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode

	//查询是否存在这条记录
	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms); err != nil {
		fmt.Println(err.Error())
	}

	return &sms
}

//根据手机号查询会员信息
func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member

	if _, err := md.Where("mobile = ?", phone).Get(&member); err != nil {
		fmt.Println(err.Error())
	}

	return &member
}

//会员新用户新增操作
func (md *MemberDao) InsertMember(member model.Member) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		logger.Error(err.Error())
	}
	return result
}

//dao层 -- 插入code方法
func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		logger.Error(err.Error())
	}
	return result
}
