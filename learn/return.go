package main

import (
	"encoding/json"
	"fmt"
)

var str = `{name: "ddd", age:11}`

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	resp, err := Myfunc([]byte(str))
	fmt.Println(string(resp), err)
}

func Myfunc(req []byte) (resp []byte, err error) {
	defer func() {
		fmt.Printf("resp:%s, err:%v", string(resp), err)
	}()

	u := User{
		Name: "xx",
		Age:  20,
	}
	p, err := json.Marshal(u)
	if err == nil {
		return
	}

	err = json.Unmarshal(p, &u)
	return req, err
}
