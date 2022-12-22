package config

import (
	"fmt"
	"github.com/jialanli/windward"
)

type Config struct {
	Db struct {
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
		Ip       string `yaml:"ip"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
	}
	Endpoint struct {
		Ip   string `yaml:"ip"`
		Port string `yaml:"port"`
	}
	Pri struct {
		Value string `yaml:"value"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
}

func Readconfig() (*Config, error) {
	//加载配置文件
	file := "./config/config.yaml"
	w := windward.GetWindward()
	w.InitConf([]string{file}) //初始化自定义的配置文件

	//获取数据库连接名密码等数据
	var config Config //定义结构体【注意：这里需要有两层结构，因为w.ReadConfig读取的是data以及data中的数据】

	err := w.ReadConfig(file, &config)
	if err != nil {
		fmt.Sprintln("初始化配置文件失败")
		return nil, err
	}
	return &config, nil
}
