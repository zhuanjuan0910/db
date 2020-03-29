package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Place struct {
	Country string `gorm:"country"`
	City    string `gorm:"city"`
	TelCode int    `gorm:"telcode" `
	gorm.Model
}

func (Place) TableName() string {
	return "place"
}
func main() {
	db, err := gorm.Open("mysql", "user:passwd@tcp(IP)/databasename?charset=utf8")
	if err != nil {
		fmt.Println("连接失败", err)
	}
	fmt.Println("success")
	db.LogMode(true)
	var place Place
	place.ID = 1

	db.Delete(&place)
	var place1 Place
	place1.ID = 2
	db.Where("country like ?", "%er%").Delete(&place1)
	var place2 Place
	place2.ID = 3
	db.Delete(&place2, "country like ?", "%pa%")
	fmt.Println(place2)
	var place3 Place
	place3.ID = 4
	db.Delete(&place3)
	fmt.Println(place3)
	var place4 Place
	db.Unscoped().Where("country=?", "china").Find(&place4)
	db.Unscoped().Delete(&place3)

}
