package model

//定义数据库字段结构体
type SmsCode struct {
	Id         int64  `xorm:"pk autoincr" json:"id"`
	Phone      string `xorm:"varchar(11)" json:"phone"`
	BizId      string `xorm:"varchar(20)" json:"biz_id"`
	Code       string `xorm:"varchar(6)" json:"code"`
	CreateTime int64  `xorm:"bigint" json:"create_time"`
}
