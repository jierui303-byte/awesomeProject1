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

	/**
		删除
	通过Delete方法删除记录，如果记录中包含了DeletedAt字段，那么将不会真正删除该记录，
	只是设置了该记录的该字段为当前时间（软删除），
	通过Unscoped方法的返回对象调用Find、Delete可以执行到被软删除的对象，进行查询或者永久删除
	*/
	db.Delete(&user)
	//UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

	// 批量删除
	db.Where("name = ?", "aa").Delete(&User{})
	//UPDATE users SET deleted_at="2013-10-29 10:23" WHERE name = 'aa';

	// 软删除的记录将在查询时被忽略
	db.Where("name = 'aa'").Find(&user)
	//SELECT * FROM users WHERE name = 'aa' AND deleted_at IS NULL;

	// 使用Unscoped查找软删除的记录
	db.Unscoped().Where("name = 'aa'").Find(&users)
	//SELECT * FROM users WHERE name = 'aa';

	// 使用Unscoped永久删除记录
	db.Unscoped().Delete(&order)
	//DELETE FROM orders WHERE id=10;

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