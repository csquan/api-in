package types

import "github.com/go-xorm/xorm"

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_db.go -package=mock
type IReader interface {
	GetMechanismInfo(name string) (*Mechanism, error)
}

type IWriter interface {
	GetSession() *xorm.Session
	GetEngine() *xorm.Engine
	CommitWithSession(db IDB, executeFunc func(*xorm.Session) error) (err error)

	InsertTransfer(itf xorm.Interface, transfer *Transfer) (err error)
	InsertWithdraw(itf xorm.Interface, withdraw *Withdraw) (err error)
	InsertMechanism(itf xorm.Interface, mechanism *Mechanism) (err error)
}

type IDB interface {
	IReader
	IWriter
}
