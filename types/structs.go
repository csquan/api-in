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

type StatusInfo struct {
	IsBlack          bool
	IsBlackIn        bool
	IsBlackOut       bool
	NowFrozenAmount  *big.Int
	WaitFrozenAmount *big.Int
}

type TxRes struct {
	Hash      string
	Op        string
	OpAddr    string
	TxGeneral *Tx
	TxErc20   *Erc20Tx
}

type ContractAbi struct {
	Contract_data string
	Abi_data      string
}

type HttpRes struct {
	Code    int
	Message string
	Data    string
}
