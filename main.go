package main

import (
	"blogweb_gin/config"
	"blogweb_gin/dao"
	logger "blogweb_gin/gb"
	"blogweb_gin/routers"
	"fmt"
)

func main() {
	if err:=config.Init("conf/conf.json");err!=nil{
		return
	}
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		return
	}
	if err:=dao.InitMySQL(config.Conf.MySQLConfig);err!=nil{
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	router:=routers.InitRouter()
	fmt.Println(config.Conf.ServerConfig.Port, *config.Conf.LogConfig)
	router.Run()
}

