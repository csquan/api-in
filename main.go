package main

import (
	"fmt"
	"github.com/ethereum/coin-manage/api"
	"github.com/ethereum/coin-manage/config"
	"github.com/ethereum/coin-manage/db"
	"github.com/sirupsen/logrus"
	"os"
)

const CONTRACTLEN = 42

func main() {
	cfg, err := config.Readconfig()
	if err != nil {
		logrus.Fatalf("read config error:%v", err)
	}

	dbConnection, err := db.NewMysql(cfg)
	if err != nil {
		logrus.Fatalf("connect to dbConnection error:%v", err)
	}

	apiservice := api.NewApiService(dbConnection, cfg)
	go apiservice.Run()

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
