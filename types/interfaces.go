package types

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_db.go -package=mock
type IReader interface {
	QueryCoinholders(contract_addr string) ([]*Balance_Erc20, error)
	QueryCoinInfos(account_addr string) ([]*Erc20Info, error)
	QueryTxHistory(account_addr string) ([]*Tx, error)
}

type IWriter interface {
}

type IDB interface {
	IReader
	IWriter
}
