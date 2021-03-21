package main

import (
	"flag"
	"log"
	"my_go/global"
	"my_go/internal/routers"
	"my_go/pkg/setting"
	"strings"
)

var (
	runmode string
	port string
	config string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
}

func main()  {
	//gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	router.Run(`:8081`)
}

func setupFlag() error {
	flag.StringVar(&port,"port","","端口")
	flag.StringVar(&runmode,"mode","","启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")

	return nil
}

func setupSetting() error {
	s,err := setting.NewSetting(strings.Split(config,`,`)...)
	if err != nil {
		return err
	}
	err = s.ReadSection(`Server`,&global.ServerSetting)
	if err != nil {
		return err
	}

	return nil
}
