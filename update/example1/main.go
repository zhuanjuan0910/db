package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	defer db.Close()
	db.LogMode(true)
	//更新查询出来的那条记录的name和age字段
	var user User
	db.First(&user) //SELECT * FROM `user`   ORDER BY `user`.`id` ASC LIMIT 1
	fmt.Println(user)
	user.Name = "hello"
	user.Age = 26
	db.Save(&user) // UPDATE `user` SET `name` = 'hello', `age` = 26  WHERE `user`.`id` = 1
	fmt.Println(user)

}
