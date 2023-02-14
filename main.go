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
	conffile string
	env      string
)

func init() {
	flag.StringVar(&conffile, "conf", "config.yaml", "conf file")
	flag.StringVar(&env, "env", "prod", "Deploy environment: [ prod | test ]. Default value: prod")
}

func main() {
	flag.Parse()

	conf, err := config.LoadConf("config.yaml", "test")
	if err != nil {
		logrus.Fatalf("load conf err:%v", err)
	}

	dbConnections, err := db.NewMysql(conf)
	if err != nil {
		logrus.Fatalf("connect to dbConnection error:%v", err)
	}

	apiservice := api.NewApiService(dbConnections, conf)
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
