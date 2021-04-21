package tool

import (
	"awesomeProject1/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//定义一个全局变量
var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	//获取数据库配置文件内容
	//databaseInfo := cfg.database
	databaseInfo := cfg.Database

	//拼装数据库信息字符
	connStr := databaseInfo.User + ":" + databaseInfo.Password + "@tcp(" + databaseInfo.Host + ":" + databaseInfo.Port + ")/" + databaseInfo.DbName + "?charset=" + databaseInfo.Charset

	//xorm实例化
	engine, err := xorm.NewEngine(databaseInfo.Driver, connStr)
	if err != nil {
		return nil, err
	}

	//sql打印调试【动态修改是否打印，修改配置文件即可】
	engine.ShowSQL(databaseInfo.ShowSql)

	//把数据库字段结构体model.SmsCode映射转变成数据库里面的一张表
	err = engine.Sync2(new(model.SmsCode))
	if err != nil {
		return nil, err
	}

	//把ormEngine对象包装在orm结构体中返回
	orm := new(Orm)
	orm.Engine = engine

	//给全局变量赋值,相当于是把ormEngine对象专门存起来，让dao层调用
	DbEngine = orm

	return orm, nil
}
