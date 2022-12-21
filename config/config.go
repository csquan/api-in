package config

import (
	"fmt"
	"github.com/jialanli/windward"
)

type Config struct {
	DB struct {
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
		Ip       string `yaml:"ip"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
	}
	ENDPOINT struct {
		Ip   string `yaml:"ip"`
		Port string `yaml:"port"`
	}
	PRI struct {
		Value string `yaml:"value"`
	}
}

func Readconfig() (string, string, string) {
	//加载配置文件
	file := "./config/config.yaml"
	w := windward.GetWindward()
	w.InitConf([]string{file}) //初始化自定义的配置文件

	//获取数据库连接名密码等数据
	var config Config //定义结构体【注意：这里需要有两层结构，因为w.ReadConfig读取的是data以及data中的数据】

	err := w.ReadConfig(file, &config)
	if err != nil {
		fmt.Sprintln("初始化配置文件失败")
		return "", "", ""
	}
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", config.DB.Name, config.DB.Password, config.DB.Ip, config.DB.Port, config.DB.Database)
	endpoint := fmt.Sprintf("http://%s:%s", config.ENDPOINT.Ip, config.ENDPOINT.Port)
	return connect, endpoint, config.PRI.Value
}
