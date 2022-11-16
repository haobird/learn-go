package main

import "fmt"

type Account struct {
	ID    int64  `json:"id" gorm:"id"`
	Name  string `json:"name" gorm:"name"`
	Ctime int64  `json:"ctime" gorm:"ctime"`
}

func main() {
	// var structArr []Account
	// fmt.Println(structArr)

	var s []string

	fmt.Println(s)
	fmt.Println(s == nil)

	s = append(s, "dd")
	fmt.Println(s)
}
