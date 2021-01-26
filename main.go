package main

import (
	"leopard/conf"
	"leopard/server"
)

func main() {
	// 初始化环境
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	_ = r.Run(":3000")

}
