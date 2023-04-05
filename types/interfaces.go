package types

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_db.go -package=mock
type IReader interface {
	GetMechanismInfo(name string) (*Mechanism, error)
}

type IWriter interface {
}

type IDB interface {
	IReader
	IWriter
}
