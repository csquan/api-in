package types

import "math/big"

type Balance_Erc20 struct {
	Id             string `xorm:"id"`
	Addr           string `xorm:"addr"`
	ContractAddr   string `xorm:"contract_addr"`
	Balance        string `xorm:"balance"`
	Height         string `xorm:"height"`
	Balance_Origin string `xorm:"balance_origin"`
}

type Tx struct {
	Id                   string `xorm:"id"`
	TxType               string `xorm:"tx_type"`
	From                 string `xorm:"addr_from"`
	To                   string `xorm:"addr_to"`
	Hash                 string `xorm:"tx_hash"`
	Index                string `xorm:"tx_index"`
	Value                string `xorm:"tx_value"`
	Input                string `xorm:"input"`
	Nonce                string `xorm:"nonce"`
	GasPrice             string `xorm:"gas_price"`
	GasLimit             string `xorm:"gas_limit"`
	GasUsed              string `xorm:"gas_used"`
	IsContract           string `xorm:"is_contract"`
	IsContractCreate     string `xorm:"is_contract_create"`
	BlockTime            string `xorm:"block_time"`
	BlockNum             string `xorm:"block_num"`
	BlockHash            string `xorm:"block_hash"`
	ExecStatus           string `xorm:"exec_status"`
	CreateTime           string `xorm:"create_time"`
	BlockState           string `xorm:"block_state"`
	MaxFeePerGas         string `xorm:"max_fee_per_gas"`
	BaseFee              string `xorm:"base_fee"`
	MaxPriorityFeePerGas string `xorm:"max_priority_fee_per_gas"`
	BurntFees            string `xorm:"burnt_fees"`
}

type Erc20Tx struct {
	Id                string `xorm:"id"`
	TxHash            string `xorm:"tx_hash"`
	Addr              string `xorm:"addr"`
	Sender            string `xorm:"sender"`
	Receiver          string `xorm:"receiver"`
	Tokens_Cnt        string `xorm:"token_cnt"`
	Log_Index         string `xorm:"log_index"`
	Tokens_Cnt_Origin string `xorm:"token_cnt_origin"`
	Create_Time       string `xorm:"create_time"`
	Block_State       string `xorm:"block_state"`
	Block_Num         string `xorm:"block_num"`
	Block_Time        string `xorm:"block_time"`
}

type Erc20Info struct {
	Id                   string `xorm:"id"`
	Addr                 string `xorm:"addr"`
	Name                 string `xorm:"name"`
	Symbol               string `xorm:"symbol"`
	Decimals             string `xorm:"decimals"`
	Totoal_Supply        string `xorm:"total_supply"`
	Totoal_Supply_Origin string `xorm:"total_supply_origin"`
	Create_Time          string `xorm:"create_time"`
}

type CoinInfo struct {
	BaseInfo    Erc20Info
	HolderCount int
	Status      uint8
}

type StatusInfo struct {
	IsBlack          bool
	IsBlackIn        bool
	IsBlackOut       bool
	NowFrozenAmount  *big.Int
	WaitFrozenAmount *big.Int
}

type TxRes struct {
	Hash      string
	OpParams  *OpParam
	Amount    uint64
	TxGeneral *Tx
}

//type ContractAbi struct {
//	Contract_addr string
//	Abi_data      string
//}
//
//type ContractReceiver struct {
//	Contract_Addr string `xorm:"contract_addr"`
//	Receiver_Addr string `xorm:"reciver_addr"`
//}

type HttpRes struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CoinData struct {
	InitCoinSupply string `json:"init_coin_supply"`
	AddCoinHistory string `json:"add_coin_history"`
}

type TxData struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Data      string `json:"data"`
	ChainId   string `json:"chainId"`
	Value     string `json:"value"`
	RequestID string `json:"requestId"`
	UID       string `json:"uid"`
	UUID      string `json:"uuid"`
}

type TxLog struct {
	Id         uint64 `xorm:"id"`
	Hash       string `xorm:"tx_hash"`
	Addr       string `xorm:"addr"`
	AddrFrom   string `xorm:"addr_from"`
	AddrTo     string `xorm:"addr_to"`
	Topic0     string `xorm:"topic0"`
	Topic1     string `xorm:"topic1"`
	Topic2     string `xorm:"topic2"`
	Topic3     string `xorm:"topic3"`
	Data       string `xorm:"log_data"`
	Index      uint   `xorm:"log_index"`
	BlockState uint8  `xorm:"block_state"`
	BlockNum   uint64 `xorm:"block_num"`
	BlockTime  uint64 `xorm:"block_time"`
}

type EventHash struct {
	Op        string `xorm:"op"`
	EventHash string `xorm:"eventhash"`
}

type OpParam struct {
	Op     string `json:"op"`
	Addr1  string `json:"addr1"`
	Addr2  string `json:"addr2"`
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
	Value3 string `json:"value3"`
}

type SignTask struct {
	SafeId     int64  `json:"safeId"`
	TaskName   string `json:"taskName"`
	Params     string `json:"params"`
	Method     int    `json:"method"`
	URL        string `json:"url"`
	CustomData string `json:"customData"`
}

type Response struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	Data    string `json:"data"`
	Success string `json:"success"`
}
