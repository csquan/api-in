package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
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

type ChainInfo struct {
	ChainName string `yaml:"chainName"`
	Db        string `yaml:"db"`
	Rpc       string `yaml:"rpc"`
	ChainId   string `yaml:"chainId"`
}

type Config struct {
	ChainInfos []ChainInfo `mapstructure:"chainInfos"`
	Endpoint   struct {
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

// LocalConfig build viper from the local disk
func LocalConfig(filename string, v *viper.Viper) error {
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("err : %v", err)
	}

	v.AddConfigPath(path) //设置读取的文件路径

	v.SetConfigName(filename) //设置读取的文件名

	err = v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read conf file err : %v", err)
	}

	return err
}

func LoadConf(fpath string, env string) (*Config, error) {
	if fpath == "" {
		return nil, fmt.Errorf("fpath empty")
	}

	if !strings.HasSuffix(strings.ToLower(fpath), ".yaml") {
		return nil, fmt.Errorf("fpath must has suffix of .yaml")
	}

	conf := &Config{}

	vip := viper.New()
	vip.SetConfigType("yaml")

	fmt.Println("read configuration from local yaml file :", fpath)
	err := LocalConfig(fpath, vip)
	if err != nil {
		return nil, err
	}
	err = vip.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
