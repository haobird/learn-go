package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID    int64  `json:"id" gorm:"id"`
	Name  string `json:"name" gorm:"name"`
	Ctime int64  `json:"ctime" gorm:"ctime"`
}

var (
	dsn = "mi:4ZKk6SaCtLlGyIZWOBdRr1yAF1HoxfLGzZ@tcp(tj1-owt-mitob-staging-db-01.kscn:3306)/mitob_platform?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(err)

	var device platform.Device
	db.First(&device)
	fmt.Println(device)
}
