package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
	使用GORM操作数据库：go get -u github.com/jinzhu/gorm
*/
func main() {
	connStr := "root:12345678@tcp(127.0.0.1:3306)/ginsql"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//创建数据库表person:id,name,age
	db.Exec("create table person(" +
		"id int auto_increment primary key," +
		"name varchar(12) not null," +
		"age int default 1" +
		")")

}
