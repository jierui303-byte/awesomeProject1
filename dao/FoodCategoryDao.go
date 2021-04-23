package dao

import (
	"awesomeProject1/model"
	"awesomeProject1/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}

//实例化dao对象
func NewFoodCategoryDao() *FoodCategoryDao {
	return &FoodCategoryDao{tool.DbEngine}
}

/**
从数据库当中查询所有的食品种类并返回
返回值：数组切片或者报错
*/
func (fcd *FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {
	var categories []model.FoodCategory
	if err := fcd.Engine.Find(&categories); err != nil {
		return nil, err
	}
	return categories, nil
}
