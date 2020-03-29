package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id   int    `gorm:"id"`
	Name string `gorm:"name"`
	Age  int    `gorm:"age"`
}

func (User) TableName() string {
	return "user"
}
func main() {
	db, err := gorm.Open("mysql", "user:passwd@tcp(IP)/databasename?charset=utf8")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	fmt.Println("success")
	db.LogMode(true)
	var user User
	user.Id = 3
	db.Model(&user).Select("age").Updates(map[string]interface{}{"name": "hello", "age": 15})
	fmt.Println(user)
	db.Model(&user).Omit("age").Updates(map[string]interface{}{"name": "zhangsan", "age": 16})
	fmt.Println(user)

}
