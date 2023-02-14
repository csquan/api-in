package db

import (
	"fmt"
	"github.com/ethereum/coin-manage/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"time"
	"xorm.io/core"
)

type Mysql struct {
	conf   *config.Config
	engine *xorm.Engine
}

type Db struct {
	ChainName string `yaml:"chainName"`
	UserName  string `yaml:"userName"`
	Password  string `yaml:"password"`
	Ip        string `yaml:"ip"`
	Port      string `yaml:"port"`
	Database  string `yaml:"database"`
}

func NewMysql(cfg *config.Config) (conn map[string]*Mysql, err error) {
	conn_map := make(map[string]*Mysql)
	for _, db := range cfg.Dbs {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db.UserName, db.Password, db.Ip, db.Port, db.Database)
		engine, err := xorm.NewEngine("mysql", dsn)
		if err != nil {
			logrus.Errorf("create engine error: %v", err)
			return nil, err
		}
		engine.ShowSQL(false)
		engine.Logger().SetLevel(core.LOG_DEBUG)
		location, err := time.LoadLocation("UTC")
		if err != nil {
			return nil, err
		}
		engine.SetTZLocation(location)
		engine.SetTZDatabase(location)

		m := &Mysql{
			conf:   cfg,
			engine: engine,
		}
		conn_map[db.ChainName] = m
	}
	return conn_map, nil
}

func (m *Mysql) GetEngine() *xorm.Engine {
	return m.engine
}

func (m *Mysql) GetSession() *xorm.Session {
	return m.engine.NewSession()
}

//
//func (m *Mysql) UpdateTransactionTask(itf xorm.Interface, task *types.TransactionTask) error {
//	_, err := itf.Table("t_transaction_task").Where("f_id = ?", task.ID).Update(task)
//	return err
//}
//func (m *Mysql) UpdateTransactionTaskMessage(taskID uint64, message string) error {
//	_, err := m.engine.Exec("update t_transaction_task set f_message = ? where f_id = ?", message, taskID)
//	return err
//}
