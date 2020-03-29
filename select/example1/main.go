package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Person struct {
	UserId   int    `gorm:"primary_key"`
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
	fmt.Println("success")
	defer db.Close()
	db.LogMode(true)
	var person []Person //注意
	//db.Where("username=?", "zhangsan").Find(&person)
	//fmt.Println(person)
	// db.Last(&person)
	// fmt.Println(person)
	db.Find(&person) //SELECT * FROM `person`
	fmt.Println(person)
	var person1 Person
	db.First(&person1, 2) // SELECT * FROM `person`  WHERE (`person`.`user_id` =2)
	fmt.Println(person1)
	var person2 Person
	db.Last(&person2)
	fmt.Println(person2) //SELECT * FROM `person`   ORDER BY `person`.`user_id` DESC  LIMIT 1

}
