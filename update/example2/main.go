package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id   int    `gorm:"id" "primary key"`
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
	var user User
	user.Id = 2
	db.Model(&user).Update("name", "hello")
	fmt.Println(user)
	db.Model(&user).Where("age=?", 16).Update("name", "xiaoming") // UPDATE `user` SET `name` = 'xiaoming'  WHERE `user`.`id` = 2 AND ((age=16))
	fmt.Println(user)
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 37}) //UPDATE `user` SET `age` = 37, `name` = 'hello'  WHERE `user`.`id` = 2
	fmt.Println(user)
	//	警告:当使用struct更新时，FORM将仅更新具有非空值的字段
	// 对于下面的更新，什么都不会更新为""，0，false是其类型的空白值
	db.Model(&user).Updates(User{Name: "", Age: 0})
	fmt.Println(user)
}
