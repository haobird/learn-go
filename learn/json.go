package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Book struct {
	Title string   `json:"title"`
	Num   int      `json:"num"`
	SN    []string `json:"sn"`
}

type User struct {
	Name  string   `json:"name"`
	Other []string `json:"other"`
	List  []Book   `json:"list"`
}

func main() {

	var q float64 = 16844.42
	var h1 float64 = 1.142350
	var h2 float64 = 1.1460
	m1 := q * h1
	m2 := q * h2
	fmt.Println(m1)
	fmt.Println(m2)

	var u User
	// u.Other = make([]string, 0)
	fmt.Printf("%+v, Other:%+v, \n", u, u.Other)
	query(u)
	buf, err := json.Marshal(u)
	fmt.Println(string(buf), err)

}

func query(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	va := v.FieldByName("Other")
	fmt.Printf("Type: %T \n", va)
	fmt.Printf("Value : %v\n", va)
	fmt.Println(va.IsNil())
}
