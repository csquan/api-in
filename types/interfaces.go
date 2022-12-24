package types

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_db.go -package=mock
type IReader interface {
	QueryCoinHolders(contractAddr string) ([]*Balance_Erc20, error)
	QueryCoinInfos(accountAddr string) ([]*Erc20Info, error)
	QueryTxHistory(accountAddr string) ([]*Tx, error)
	QueryTxErc20History(accountAddr string) ([]*Erc20Tx, error)
	QueryABI(contractAddr string) (*ContractAbi, error)
}

type IWriter interface {
}

type IDB interface {
	IReader
	IWriter
}
