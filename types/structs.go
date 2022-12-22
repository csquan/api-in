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
	TxType               string `xorm:"txType"`
	From                 string `xorm:"from"`
	To                   string `xorm:"to"`
	Hash                 string `xorm:"hash"`
	Index                string `xorm:"index"`
	Value                string `xorm:"value"`
	Input                string `xorm:"input"`
	Nonce                string `xorm:"nonce"`
	GasPrice             string `xorm:"gasPrice"`
	GasLimit             string `xorm:"gasLimit"`
	GasUsed              string `xorm:"gasUsed"`
	IsContract           string `xorm:"isContract"`
	IsContractCreate     string `xorm:"isContractCreate"`
	BlockTime            string `xorm:"blockTime"`
	BlockNum             string `xorm:"blockNum"`
	BlockHash            string `xorm:"blockHash"`
	ExecStatus           string `xorm:"execStatus"`
	CreateTime           string `xorm:"createTime"`
	BlockState           string `xorm:"blockState"`
	MaxFeePerGas         string `xorm:"maxFeePerGas"`
	BaseFee              string `xorm:"baseFee"`
	MaxPriorityFeePerGas string `xorm:"maxPriorityFeePerGas"`
	BurntFees            string `xorm:"burntFees"`
}

type Erc20Tx struct {
	Id                string `xorm:"id"`
	TxHash            string `xorm:"tx_hash"`
	Addr              string `xorm:"addr"`
	Sender            string `xorm:"sender"`
	Receiver          string `xorm:"receiver"`
	Tokens_Cnt        string `xorm:"tokens_cnt"`
	Log_Index         string `xorm:"log_index"`
	Tokens_Cnt_Origin string `xorm:"tokens_cnt_origin"`
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
