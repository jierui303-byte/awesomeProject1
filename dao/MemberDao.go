package dao

import (
	"awesomeProject1/model"
	"awesomeProject1/tool"

	"github.com/micro/go-micro/v2/logger"
)

//结构体绑定orm对象
type MemberDao struct {
	*tool.Orm
}

//dao层 -- 插入code方法
func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	result, err := md.InsertOne(&sms)
	if err != nil {
		logger.Error(err.Error())
	}
	return result
}
