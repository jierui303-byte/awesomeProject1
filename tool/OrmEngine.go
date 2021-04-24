package tool

import (
	"awesomeProject1/model"
	"fmt"

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
	err = engine.Sync2(new(model.SmsCode), new(model.Member), new(model.FoodCategory), new(model.Shop))
	if err != nil {
		return nil, err
	}

	//把ormEngine对象包装在orm结构体中返回
	orm := new(Orm)
	orm.Engine = engine

	//给全局变量赋值,相当于是把ormEngine对象专门存起来，让dao层调用
	DbEngine = orm

	//初始化插入shop数据
	InitShopData()

	return orm, nil
}

//向shop表中插入初始化数据
func InitShopData() {
	//新建一个数组切片
	shops := []model.Shop{
		model.Shop{Id: 1, Name: "嘉禾一品（温都水城）", Address: "北京市昌平区宏福苑温都水城F1", Longitude: 116.36868, Latitude: 40.10039,
			Phone: "13437850035", Status: "1", RecentOrderNum: 106, RatingCount: 961, Rating: 4.7, PromotionInfo: "欢迎光临，用餐高峰请提前下单，谢谢",
			OpeningHours: "8:30/20:30", IsNew: true, IsPremium: true, ImagePath: "", MinimumOrderAmount: 20, DeliveryFee: 5},
		model.Shop{Id: 479, Name: "杨国福麻辣烫", Address: "北京市市蜀山区南二环路天鹅湖万达广场8号楼1705室", Longitude: 117.22124, Latitude: 31.81948, Phone: "13167583411",
			Status: "1", RecentOrderNum: 755, RatingCount: 167, Rating: 4.2, PromotionInfo: "欢迎光临，用餐高峰请提前下单，谢谢", OpeningHours: "8:30/20:30",
			IsNew: true, IsPremium: true, ImagePath: "", MinimumOrderAmount: 20, DeliveryFee: 5},
		model.Shop{Id: 485, Name: "好适口", Address: "北京市海淀区西二旗大街58号", Longitude: 120.65355, Latitude: 31.26578, Phone: "12345678901",
			Status: "1", RecentOrderNum: 58, RatingCount: 576, Rating: 4.6, PromotionInfo: "欢迎光临，用餐高峰请提前下单，谢谢", OpeningHours: "8:30/20:30",
			IsNew: true, IsPremium: true, ImagePath: "", MinimumOrderAmount: 20, DeliveryFee: 5},
		model.Shop{Id: 486, Name: "东来顺旗舰店", Address: "北京市天河区东圃镇汇彩路38号1领汇创展商务中心401", Longitude: 113.41724, Latitude: 23.1127, Status: "1",
			Phone: "13544323775", RecentOrderNum: 542, RatingCount: 372, Rating: 4.2, PromotionInfo: "老北京正宗涮羊肉,非物质文化遗产",
			OpeningHours: "09:00/21:30", IsNew: true, IsPremium: true, ImagePath: "", MinimumOrderAmount: 20, DeliveryFee: 5},
		model.Shop{Id: 487, Name: "北京酒家", Address: "北京市海淀区上下九商业步行街内", Longitude: 113.24826, Latitude: 23.11488, Phone: "13257482341", Status: "0",
			RecentOrderNum: 923, RatingCount: 871, Rating: 4.2, PromotionInfo: "北京第一家传承300年酒家", OpeningHours: "8:30/20:30", IsNew: true, IsPremium: true, ImagePath: "",
			MinimumOrderAmount: 20, DeliveryFee: 5},
		model.Shop{Id: 488, Name: "和平鸽饺子馆", Address: "北京市越秀区德政中路171", Longitude: 113.27521, Latitude: 23.12092,
			Phone: "17098764762", Status: "1", RecentOrderNum: 483, RatingCount: 273, Rating: 4.2, PromotionInfo: "吃饺子就来和平鸽饺子馆", OpeningHours: "8:30/20:30",
			IsNew: true, IsPremium: true, ImagePath: "", MinimumOrderAmount: 20, DeliveryFee: 5},
	}

	//多条数据插入：事务
	session := DbEngine.NewSession() //新建事务
	defer session.Close()

	//事务操作：事务开始, 执行操作（回滚），提交事务
	err := session.Begin() //事务开始
	if err != nil {
		fmt.Println(err.Error())
	}

	//轮询插入
	for _, shop := range shops {
		_, err := session.Insert(&shop) //事务执行
		if err != nil {
			session.Rollback() //事务回滚
			return
		}
	}

	//事务提交
	err = session.Commit()
	if err != nil {
		fmt.Println(err.Error())
	}
}
