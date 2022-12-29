package types

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_db.go -package=mock
type IReader interface {
	QueryCoinHolderCount(contractAddr string) (int, error)
	QueryCoinHolders(contractAddr string) ([]*Balance_Erc20, error)
	QueryCoinInfos(accountAddr string) ([]*Erc20Info, error)
	QueryTxHistory(accountAddr string) ([]*Tx, error)
	QueryTxErc20History(accountAddr string) ([]*Erc20Tx, error)
	QueryAllCoinAllHolders(accountAddr string) (int, error)
	QuerySpecifyCoinInfo(contractAddr string) (*Erc20Info, error)
	QueryTxlogByHash(hash string) (*TxLog, error)
	GetEventHash() ([]*EventHash, error)

	GetCoinBalance(accountAdr string, contractAddr string) (string, error)

	GetBlockHeight() (int, error)
	QueryBurnTxs(accountAddr string, contractAddr string) ([]*Tx, error)
}

type IWriter interface {
}

type IDB interface {
	IReader
	IWriter
}
