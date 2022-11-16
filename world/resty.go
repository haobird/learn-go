package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Library struct {
	Name   string
	Latest string
}

type Libraries struct {
	Results []*Library
}

type Resp struct {
	Code int
	Msg  string
	Data interface{} `json:"data"`
}

func main() {
	client := resty.New()

	libraries := &Libraries{}
	client.R().SetResult(libraries).Get("https://api.cdnjs.com/libraries")
	fmt.Printf("%d libraries\n", len(libraries.Results))

	result := &Resp{}
	client.R().SetResult(result).ForceContentType("application/json").Get("http://127.0.0.1:8081/rtc/third/bindKey/query")
	fmt.Println("libraries", result)
}
