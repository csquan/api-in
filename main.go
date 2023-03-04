package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/coin-manage/api"
	"github.com/ethereum/coin-manage/config"
	"github.com/ethereum/coin-manage/db"
	"github.com/sirupsen/logrus"
	"os"
)

const CONTRACTLEN = 42

var (
	yamlConfig string
)

func main() {
	flag.StringVar(&yamlConfig, "conf", "conf/config.yaml", "conf file")

	flag.Parse()

	conf, err := config.LoadConf(yamlConfig)
	if err != nil {
		logrus.Fatalf("load conf err:%v", err)
		panic(err)
	}

	dbConnections, err := db.NewMysql(conf)
	if err != nil {
		logrus.Fatalf("connect to dbConnection error:%v", err)
		panic(err)
	}

	apiService := api.NewApiService(dbConnections, conf)
	go apiService.Run()

	//listen kill signal
	closeCh := make(chan os.Signal, 1)

	for {
		select {
		case <-closeCh:
			fmt.Printf("receive os close sigal")
			return
		}
	}
}
