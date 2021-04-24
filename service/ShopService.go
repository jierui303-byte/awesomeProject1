package service

import (
	"awesomeProject1/dao"
	"awesomeProject1/model"
	"strconv"
)

type ShopService struct {
}

func (shopService *ShopService) ShopList(long, lat string) []model.Shop {
	//类型转换：
	longitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}

	//数据库操作去dao层进行，service层负责接收结果并处理返回给控制器
	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longitude, latitude)
}
