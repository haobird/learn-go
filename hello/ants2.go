package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/panjf2000/ants/v2"
)

// 无限增加协程
// 控制协程数量
// 统计协程数据

func main() {
	runTimes := 10000
	p, _ := ants.NewPoolWithFunc(5, demoFunc2)

	go func() {
		for i := 0; i < runTimes; i++ {
			p.Invoke(i)
			// go demoFunc2(i)
			// fmt.Println(time.Now().UTC())
		}
	}()

	go func() {
		for i := 0; i < runTimes; i++ {
			p.Invoke(i)
			// demoFunc2(i)
			// fmt.Println(time.Now().UTC())
		}
	}()
	// 启动pprof
	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}

func demoFunc2(i interface{}) {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World", i)
}
