package main

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
使用GORM操作数据库：go get -u github.com/jinzhu/gorm
*/
func main() {

	//账号：root 密码：12345678 数据库名：ginsql
	db, err := gorm.Open("mysql", "root:12345678@/gorm?charset=utf8&parseTime=True&loc=local")
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

	//一：插入/更新记录
	user1 := User{Name: "aa"}
	db.Save(&user1)
	//新增成功后，如果主键是由数据库生成，会将主键回显到实体对象的属性
	println(user1.ID)
	user1.Name = "bb"
	db.Save(&user1) //更新

	//二：插入记录
	user2 := User{Name: "aa"}
	db.Create(&user2) //没有设定主键，默认由数据库自增
	println(user2.ID)

	//NewRecord方法用于判断某个对象是否可以作为新纪录插入，
	//如果该对象主键为空或者0，或者数据库表中不存在该主键记录，返回true，否则返回false，
	//所以可以用于辅助Create方法
	if db.NewRecord(&user2) {
		db.Create(&user2)
	}

}

//定义嵌套gorm.Model这个结构体的类型-定义数据库表中的一些常用基本字段
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

//定义User实体
type User struct {
	gorm.Model
	Name string
}
