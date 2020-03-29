package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Person struct {
	UserId   int    `gorm:"user_id"`
	UserName string `gorm:"username"`
	Sex      string `gorm:"sex"`
	Email    string `gorm:"email"`
}

func (Person) TableName() string {
	return "person"
}
func main() {
	db, err := gorm.Open("mysql", "user:passwd@tcp(IP)/databasename?charset=utf8")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	fmt.Println("sucess")
	defer db.Close()
	var person Person
	db.LogMode(true)
	db.Where("username=?", "lisi").First(&person) //SELECT * FROM `person`  WHERE (username='lisi') LIMIT 1
	fmt.Println(person)
	var people []Person
	db.Where("username<>?", "llisi").Find(&people) //SELECT * FROM `person`  WHERE (username<>'llisi')
	fmt.Println(people)
	db.Where("username like ?", "%xiao%").Find(&people) //SELECT * FROM `person`  WHERE (username like '%xiao%')
	fmt.Println(people)
	var person1 Person
	db.Where("username=? and sex=?", "xiaofang", "woman").Find(&person1) //SELECT * FROM `person`  WHERE (username='xiaofang'and sex='woman')
	fmt.Println(person1)
	var users []Person
	db.Where("username in (?)", []string{"zhangsan", "lisi"}).Find(&users) // SELECT * FROM `person`  WHERE (username in ('zhangsan','lisi'))
	fmt.Println(users)
	var use Person
	db.Where(&Person{UserName: "zhangsan", Sex: "man"}).Find(&use)
	fmt.Println(use)

}
