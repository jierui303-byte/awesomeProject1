package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/**
gin中操作数据库的驱动包有多个：gorm, sql
	使用GORM操作数据库：go get -u github.com/jinzhu/gorm
	sql: go get -u github.com/go-sql-driver/mysql
*/
func main() {

	//账号：root 密码：12345678 数据库名：ginsql
	connStr := "root:12345678@tcp(127.0.0.1:3306)/ginsql"

	//连接数据库ginsql
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//创建数据库表person:id,name,age
	// _, err = db.Exec("create table person(" +
	// 	"id int auto_increment primary key," +
	// 	"name varchar(12) not null," +
	// 	"age int default 1" +
	// 	")")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	return
	// } else {
	// 	fmt.Println("数据库表创建成功")
	// }

	//插入数据到数据库表
	_, err = db.Exec("insert into person(name, age) values(?, ?);", "jierui303", 23)
	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("数据插入成功")
	}

	//查询数据库
	rows, err := db.Query("select id,name,age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//scan定义代码块[代码块这个相当于是轮询拿到结果集中的数据]
scan:
	//提取查询结果里的数据
	if rows.Next() {
		//实例化结构体Person
		person := new(Person)
		//scan函数读取结果分别赋值给person变量中
		err := rows.Scan(&person.Id, &person.Name, &person.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println("读取数据结果：", person.Id, person.Name, person.Age)
		goto scan
	}

}

type Person struct {
	Id   int
	Name string
	Age  int
}
