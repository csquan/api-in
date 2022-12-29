package types

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_db.go -package=mock
type IReader interface {
	QueryCoinHolderCount(contractAddr string) (int, error)
	QueryCoinHolders(contractAddr string) ([]*Balance_Erc20, error)
	QueryCoinInfos(accountAddr string) ([]*Erc20Info, error)
	QueryTxHistory(accountAddr string) ([]*Tx, error)
	QueryTxErc20History(accountAddr string) ([]*Erc20Tx, error)
	QueryABI(contractAddr string) (*ContractAbi, error)
	QueryAllCoinAllHolders(accountAddr string) (int, error)
	QueryReceiver(contractAddr string) (ContractReceiver, error)
	QuerySpecifyCoinInfo(contractAddr string) (*Erc20Info, error)
	QueryTxlogByHash(hash string) (*TxLog, error)
	GetEventHash() ([]*EventHash, error)

	GetCoinBalance(accountAdr string, contractAddr string) (string, error)

	GetBlockHeight() (int, error)
}

type IWriter interface {
	InsertReceiver(receiver *ContractReceiver) error
}

type IDB interface {
	IReader
	IWriter
}
