package dao

import (
	"awesomeProject1/model"
	"awesomeProject1/tool"
)

type ShopDao struct {
	*tool.Orm
}

func NewShopDao() *ShopDao {
	return &ShopDao{tool.DbEngine}
}

//定义常量
const DEFAULT_RANGE = 5

func (shopDao *ShopDao) QueryShops(longitude, latitude float64) []model.Shop {
	var shops []model.Shop
	err := shopDao.Engine.Where("longitude > ? and longitude < ? and latitude < ? and latitude < ?", longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)
	if err != nil {
		return nil
	}

	return shops
}
