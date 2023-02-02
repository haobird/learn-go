package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now().Unix()
	var nInt64 int64 = 10 //这样就可以了
	ttl := time.Duration(nInt64) * time.Second
	fmt.Println("ttl:", ttl)
	time.Sleep(ttl)
	endTime := time.Now().Unix()
	fmt.Println("start:", startTime, "end:", endTime)

}
