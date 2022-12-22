package db

import (
	"github.com/ethereum/coin-manage/types"
)

func (m *Mysql) QueryCoinholders(contract_addr string) ([]*types.Balance_Erc20, error) {
	balances := make([]*types.Balance_Erc20, 0)
	err := m.engine.Table("balance_erc20").Where("contract_addr = ?", contract_addr).Find(&balances)
	if err != nil {
		return nil, err
	}
	return balances, err
}

func (m *Mysql) QueryTxHistory(account_addr string) ([]*types.Tx, error) {
	txs := make([]*types.Tx, 0)
	err := m.engine.Table("tx").Where("adrr_from = ? or addr_to = ? ", account_addr, account_addr).Find(&txs)
	if err != nil {
		return nil, err
	}
	return txs, err
}

func (m *Mysql) QueryCoinInfos(account_addr string) ([]*types.Erc20Info, error) {
	infos := make([]*types.Erc20Info, 0)
	err := m.engine.Table("Erc20Info").Where("addr = ? ", account_addr).Find(&infos)
	if err != nil {
		return nil, err
	}
	return infos, err
}
