package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "./logcollect.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	port, err := conf.Int("Server::port")
	if err != nil {
		fmt.Println("read server:port failed, err:", err)
		return
	}

	fmt.Println("Port:", port)
	log_level, err := conf.Int("Log::log_level")
	if err != nil {
		fmt.Println("read log_level failed, ", err)
		return
	}
	fmt.Println("log_level:", log_level)

	log_path := conf.String("Log::log_path") //返回值无错误处理
	fmt.Println("log_path:", log_path)
}
