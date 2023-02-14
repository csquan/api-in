package config

import (
	"fmt"
	"github.com/jialanli/windward"
	"os"
	"reflect"
)

type stdout struct {
	Enable bool `mapstructure:"enable"`
	Level  int  `mapstructure:"level"`
}

type file struct {
	Enable bool   `mapstructure:"enable"`
	Path   string `mapstructure:"path"`
	Level  int    `mapstructure:"level"`
}

type kafka struct {
	Enable  bool     `mapstructure:"enable"`
	Level   int      `mapstructure:"level"`
	Brokers []string `mapstructure:"kafka_servers"`
	Topic   string   `mapstructure:"topic"`
}

type Db struct {
	ChainName string `yaml:"chainName"`
	UserName  string `yaml:"userName"`
	Password  string `yaml:"password"`
	Ip        string `yaml:"ip"`
	Port      string `yaml:"port"`
	Database  string `yaml:"database"`
}

type Config struct {
	Chains   []string
	Db       interface{}
	Endpoint struct {
		Ip   string `yaml:"ip"`
		Port string `yaml:"port"`
	}
	TxState struct {
		EndPoint string `yaml:"endpoint"`
	}
	Pri struct {
		Value string `yaml:"value"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
	Log struct {
		Stdout stdout `mapstructure:"stdout"`
		File   file   `mapstructure:"file"`
		Kafka  kafka  `mapstructure:"kafka"`
	}
}

func Readconfig(filename string) (*Config, error) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Errorf("++++++err ++++++++++: %v", err)
		return nil, fmt.Errorf("err : %v", err)
	}
	//加载配置文件
	file := path + "/" + filename
	w := windward.GetWindward()
	w.InitConf([]string{file}) //初始化自定义的配置文件

	dbs := w.GetVal(file, "dbs.db")

	//获取数据库连接名密码等数据
	var config Config //定义结构体【注意：这里需要有两层结构，因为w.ReadConfig读取的是data以及data中的数据】

	//arr := dbs
	//for _, value := range arr {
	//	config.Chains = append(config.Chains, value.ChainName)
	//}

	config.Db = dbs

	switch reflect.TypeOf(dbs).Kind() {
	case reflect.Array:
		s := reflect.ValueOf(dbs)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}
	}

	return &config, nil
}
