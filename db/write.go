package db

import (
	"fmt"
	"github.com/ethereum/api-in/types"
	"time"

	"github.com/ethereum/api-in/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"xorm.io/core"
)

type Mysql struct {
	conf   *config.Config
	engine *xorm.Engine
}

func NewMysql(cfg *config.Config) (m *Mysql, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.Db.Name, cfg.Db.Password, cfg.Db.Ip, cfg.Db.Port, cfg.Db.Database)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logrus.Errorf("create engine error: %v", err)
		return
	}
	engine.ShowSQL(false)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	location, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	engine.SetTZLocation(location)
	engine.SetTZDatabase(location)

	m = &Mysql{
		conf:   cfg,
		engine: engine,
	}

	return
}

func (m *Mysql) GetEngine() *xorm.Engine {
	return m.engine
}

func (m *Mysql) GetSession() *xorm.Session {
	return m.engine.NewSession()
}

func (m *Mysql) InsertTransfer(itf xorm.Interface, transfer *types.Transfer) (err error) {
	_, err = itf.Insert(transfer)
	if err != nil {
		logrus.Errorf("insert transfer  error:%v, transfer:%v", err, transfer)
	}
	return
}

func (m *Mysql) InsertWithdraw(itf xorm.Interface, withdraw *types.Withdraw) (err error) {
	_, err = itf.Insert(withdraw)
	if err != nil {
		logrus.Errorf("insert withdraw  error:%v, withdraw:%v", err, withdraw)
	}
	return
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
