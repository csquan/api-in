package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/api-in/api"
	"github.com/ethereum/api-in/config"
	"github.com/ethereum/api-in/db"
	"github.com/ethereum/api-in/log"
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

	cfg, err := config.Readconfig(conffile)
	if err != nil {
		logrus.Fatalf("read config error:%v", err)
	}

	err = log.Init("coin-manage", cfg)
	if err != nil {
		log.Fatal(err)
	}
	dbConnection, err := db.NewMysql(cfg)
	if err != nil {
		logrus.Fatalf("connect to dbConnection error:%v", err)
	}

	apiService := api.NewApiService(dbConnection, cfg)
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
