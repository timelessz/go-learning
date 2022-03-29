package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)
// 定义一个数据模型(user表)// 列名是字段名的蛇形小写(PassWd->pass_word)

type User struct {
	Id       uint   `gorm:"AUTO_INCREMENT"`
	Name     string `gorm:"size:50"`
	Age      int    `gorm:"size:3"`
	Birthday *time.Time
	Email    string `gorm:"type:varchar(50);unique_index"`
	PassWord string `gorm:"type:varchar(25)"`
}

var db *gorm.DB

func main() {
	db, err := gorm.Open("mysql", "root:bgbiao.top@(127.0.0.1:13306)/test_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Errorf("创建数据库连接失败:%v", err)
	}
	defer db.Close()
	// 自动迁移数据结构(table schema)
	//注意:在gorm中，默认的表名都是结构体名称的复数形式，比如User结构体默认创建的表为users
	//db.SingularTable(true) 可以取消表名的复数形式，使得表名和结构体名称一致    db.AutoMigrate(&User{})
	//添加唯一索引
	db.Model(&User{}).AddUniqueIndex("name_email", "id", "name", "email")
	// 插入记录
	db.Create(&User{Name: "bgbiao", Age: 18, Email: "bgbiao@bgbiao.top"})
	db.Create(&User{Name: "xxb", Age: 18, Email: "xxb@bgbiao.top"})
	var user User
	var users []User
	// 查看插入后的全部元素
	fmt.Printf("插入后元素:\n")
	db.Find(&users)
	fmt.Println(users)
	// 查询一条记录
	db.First(&user, "name = ?", "bgbiao")
	fmt.Println("查看查询记录:", user)
	// 更新记录(基于查出来的数据进行更新)
	db.Model(&user).Update("name", "biaoge")
	fmt.Println("更新后的记录:", user)
	// 删除记录
	db.Delete(&user)
	// 查看全部记录
	fmt.Println("查看全部记录:")
	db.Find(&users)
	fmt.Println(users)
}
