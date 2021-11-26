package main

import (
	"MyGoCode/cmdb/config"
	_ "MyGoCode/cmdb/routers" // 初始化路由
	"MyGoCode/cmdb/services"
	"log"
	"net/http"
)

func main() {
	/*
		1、定义 handler/handlerFunc
		2、绑定 url 和 handler 的关系
		3、启动服务
	*/

	//addr := ":9999"
	//err := http.ListenAndServe(addr, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	err := services.InitDb(config.DbType, config.DbDsn)
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(config.ServerAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
