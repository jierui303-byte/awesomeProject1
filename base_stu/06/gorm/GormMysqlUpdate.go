package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

/**
	使用GORM操作数据库：go get -u github.com/jinzhu/gorm



*/
func main() {
	db, err := gorm.Open("mysql", "root:root@/gorm?charset=utf8&parseTime=True&loc=local")
	defer db.Close()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//创建表默认使用结构体类型名称的驼峰命名复数形式作为表名，比如User就是users
	//设定不使用复数形式，则User对应的表为user
	db.SingularTable(true)

	//创建表名users
	db.CreateTable(&User{})

	//更新表
	db.CreateTable(&User{})
	db.AutoMigrate(&User{})

	//插入/更新记录
	user := User{Name: "aa"}
	db.Save(&user)
	//新增成功后，如果主键是由数据库生成，会将主键回显到实体对象的属性
	println(user.ID)
	user.Name = "bb"
	db.Save(&user)//更新

	//插入记录
	user := User{Name: "aa"}
	db.Create(&user)	//没有设定主键，默认由数据库自增
	println(user.ID)

	//NewRecord方法用于判断某个对象是否可以作为新纪录插入，
	//如果该对象主键为空或者0，或者数据库表中不存在该主键记录，返回true，否则返回false，
	//所以可以用于辅助Create方法
	if(db.NewRecord(&user)){
		db.Create(&user)
	}


	/**
		修改记录
		Update和Updates方法提供对记录进行更新操作，可以通过Map或者struct传递更新属性，建议通过Map
	因为通过struct更新时，FORM将仅更新具有非空值的字段
	*/
	// 使用`map`更新多个属性，只会更新这些更改的字段
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})

	// 使用组合条件批量更新单个属性
	db.Model(&user).Where("name= ?", "aa").Update("name", "hello")
	UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND name='aa';

	// 使用`struct`更新多个属性，只会更新这些更改的和非空白字段
	db.Model(&user).Updates(User{Name: "hello", Age: 18})
	UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// 对于下面的更新，什么都不会更新为""，0，false是其类型的空白值
	db.Model(&user).Updates(User{Name: "", Age: 0, Actived: false})


	/**
		删除
	通过Delete方法删除记录，如果记录中包含了DeletedAt字段，那么将不会真正删除该记录，
	只是设置了该记录的该字段为当前时间（软删除），
	通过Unscoped方法的返回对象调用Find、Delete可以执行到被软删除的对象，进行查询或者永久删除
	*/
	db.Delete(&user)
	UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

	// 批量删除
	db.Where("name = ?", "aa").Delete(&User{})
	UPDATE users SET deleted_at="2013-10-29 10:23" WHERE name = 'aa';

	// 软删除的记录将在查询时被忽略
	db.Where("name = 'aa'").Find(&user)
	SELECT * FROM users WHERE name = 'aa' AND deleted_at IS NULL;

	// 使用Unscoped查找软删除的记录
	db.Unscoped().Where("name = 'aa'").Find(&users)
	SELECT * FROM users WHERE name = 'aa';

	// 使用Unscoped永久删除记录
	db.Unscoped().Delete(&order)
	DELETE FROM orders WHERE id=10;


	//查询


}

//定义嵌套gorm.Model这个结构体的类型-定义数据库表中的一些常用基本字段
type Model struct {
	ID   uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//定义User实体
type User struct{
	gorm.Model
	Name string
}