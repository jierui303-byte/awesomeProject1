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


	//查询
	// 获取第一条记录，按主键排序
	db.First(&user)
	//SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取最后一条记录，按主键排序
	db.Last(&user)
	//SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// 获取所有记录
	db.Find(&users)
	//SELECT * FROM users;

	// 按主键获取
	db.First(&user, 23)
	//SELECT * FROM users WHERE id = 23 LIMIT 1;

	// 简单SQL
	db.Find(&user, "name = ?", "jinzhu")
	//SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	//SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// Struct
	db.Find(&users, User{Age: 20})
	//SELECT * FROM users WHERE age = 20;

	// Map
	db.Find(&users, map[string]interface{}{"age": 20})
	//SELECT * FROM users WHERE age = 20;

	db.Modal(&User{}).Find(&users)


	//where进行sql条件查询
	whereQuery()

}

func whereQuery()  {
	// 获取第一个匹配记录
	db.Where("name = ?", "jinzhu").First(&user)
	//SELECT * FROM users WHERE name = 'jinzhu' limit 1;

	// 获取所有匹配记录
	db.Where("name = ?", "jinzhu").Find(&users)
	//SELECT * FROM users WHERE name = 'jinzhu';

	db.Where("name <> ?", "jinzhu").Find(&users)

	// IN
	db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)

	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&users)

	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

	// Time
	db.Where("updated_at > ?", lastWeek).Find(&users)

	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)

}

func mapOrStractQuery()  {
	// Struct
	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	//SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

	// Map
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	//SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 主键的Slice
	db.Where([]int64{20, 21, 22}).Find(&users)
	//SELECT * FROM users WHERE id IN (20, 21, 22);
}

func notQuery()  {
	db.Not("name", "jinzhu").First(&user)
	SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1;

	// Not In
	db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Not In slice of primary keys
	db.Not([]int64{1,2,3}).First(&user)
	SELECT * FROM users WHERE id NOT IN (1,2,3);

	db.Not([]int64{}).First(&user)
	SELECT * FROM users;

	// Plain SQL
	db.Not("name = ?", "jinzhu").First(&user)
	SELECT * FROM users WHERE NOT(name = "jinzhu");

	// Struct
	db.Not(User{Name: "jinzhu"}).First(&user)
	SELECT * FROM users WHERE name <> "jinzhu";
}

func orQuery()  {
	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
	SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	// Struct
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2"}).Find(&users)
	SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';

	// Map
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2"}).Find(&users)
}

//查询链:多个查询条件可以直接拼接构建复合条件
func lianQuery()  {
	db.Where("name <> ?","jinzhu").Where("age >= ? and role <> ?",20,"admin").Find(&users)
	SELECT * FROM users WHERE name <> 'jinzhu' AND age >= 20 AND role <> 'admin';

	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Not("name = ?", "jinzhu").Find(&users)
	1
	2
	3
	4
}

//Select字段:通过Select方法进行部分字段的查询
func selectQuery()  {
	db.Select("name, age").Find(&users)
	//SELECT name, age FROM users;

	db.Select([]string{"name", "age"}).Find(&users)
	//SELECT name, age FROM users;

	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	//SELECT COALESCE(age,'42') FROM users;
}

//Order排序:通过Order方法对返回结果进行排序
func orderQuery()  {
	db.Order("age desc, name").Find(&users)
	SELECT * FROM users ORDER BY age desc, name;

	// Multiple orders
	db.Order("age desc").Order("name").Find(&users)
	SELECT * FROM users ORDER BY age desc, name;

	// ReOrder
	db.Order("age desc").Find(&users1).Order("age", true).Find(&users2)
	SELECT * FROM users ORDER BY age desc; (users1)
	SELECT * FROM users ORDER BY age; (users2)
}

//Limit
func limitQuery()  {
	db.Limit(3).Find(&users)
	SELECT * FROM users LIMIT 3;

	// Cancel limit condition with -1
	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	SELECT * FROM users LIMIT 10; (users1)
	SELECT * FROM users; (users2)
}

//Offset
func OffsetQuery()  {
	db.Offset(3).Find(&users)
	SELECT * FROM users OFFSET 3;

	// Cancel offset condition with -1
	db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	SELECT * FROM users OFFSET 10; (users1)
	SELECT * FROM users; (users2)
}

//Count:count方法返回结果条数
func countQuery()  {
	db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
	SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)

	db.Table("deleted_users").Count(&count)
	SELECT count(*) FROM deleted_users;
}

//Join:通过Join方法进行多表查询
func joinQuery()  {
	db.Table("users")
	.Select("users.name, emails.email")
	.Joins("left join emails on emails.user_id = users.id")
}

//Scan:Scan方法将结果扫描到另一个结构中。比如
func scanQuery()  {
	type User struct{}
	type Email struct{}
	type result struct{
		User
		Email
	}

	func main(){
		...
		user := User{}
		db.Modal(&User{}).Where("1 = 1").Scan(&user)
		...
		res := make([]Result,1)
		db.Table("users")
		.Select("users.name, emails.email")
		.Joins("left join emails on emails.user_id = users.id")
		.Scan(&res)
	}
}

//Scopes:通过Scopes可以将Where语句封装为方法来使用，动态添加参数
func scopesQuery()  {
	func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
		return func (db *gorm.DB) *gorm.DB {
			return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
		}
	}
	db.Scopes(OrderStatus([]string{"paid", "shipped"})).Find(&orders)
	// 查找所有付费，发货订单
}

//关联结构:一对一  默认使用主键作为外键，外键默认命名为 (关联结构体类型名称+关联结构体主键属性名称)
func guanlianQuery()  {
	// `User`属于`Profile`, `ProfileID`为外键
	type User struct {
		gorm.Model
		Profile   Profile
		ProfileID int
	}

	type Profile struct {
		gorm.Model
		Name string
	}

	db.Model(&user).Related(&profile)
	SELECT * FROM profiles WHERE id = 111; // 111是user的外键ProfileID


	/**
	通过配置ForeignKey指定该关联属性对应在本结构体的外键
	通过配置AssociationForeignKey指定该关联属性在其关联结构体的外键属性
	 */
	type Profile struct {
		gorm.Model
		Refer string
		Name  string
	}

	type User struct {
		gorm.Model
		Profile   Profile `gorm:"ForeignKey:ProfileID;AssociationForeignKey:Refer"`
		ProfileID int
	}
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