package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-zookeeper/zk"
	"go.uber.org/zap"
)

type Mylog struct {
	logger *zap.Logger
}

func (m Mylog) Printf(msg string, params ...interface{}) {
	m.logger.Sugar().Info(msg, params)
}

func main() {
	// logger, _ := zap.NewProduction()
	log.Printf("dd")
	// mylog := Mylog{logger}
	// 创建zk连接地址
	hosts := []string{"tjwqstaging.zk.hadoop.srv:2181"}
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5, zk.WithLogInfo(true))
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(11)
	println(conn.Server())

	hostPro := new(zk.DNSHostProvider)
}

// 注册

// 删除

// 获取

// 获取一次则直接监听
