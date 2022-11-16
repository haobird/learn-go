package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var confPath string

func main() {
	flag.StringVar(&confPath, "c", "", "config file")
	flag.Parse()

	InitConfig()

	addr := viper.GetStringSlice("kafka.addr")
	fmt.Println(addr)

	var addrs map[string][]string
	err := viper.UnmarshalKey("kafka.addrs", &addrs)
	fmt.Println(addrs)
	cluster := "c3"
	if val, ok := addrs[cluster]; ok {
		fmt.Println(val)
	}
	// for _, val := range addrs {
	// 	fmt.Println(val)
	// }
	fmt.Println(err)

}

// 读取配置文件路径
func InitConfig() {
	// 如果没有单独指定配置文件路径，读取环境变量
	if confPath == "" {
		// 直接获取环境变量中的 配置文件路径
		confPath = GetEnvDefault("APP_CONFIG_PATH", "")
	}

	// 加载日志文件
	viper.SetConfigFile(confPath) // 如果指定了配置文件，则解析指定的配置文件
	viper.AutomaticEnv()          // 读取匹配的环境变量
	viper.SetEnvPrefix("APP")     // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

// 获取环境变量信息
func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}
