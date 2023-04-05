package db

import (
	"fmt"
	"github.com/ethereum/api-in/types"
)

func (m *Mysql) GetMechanismInfo(name string) (*types.Mechanism, error) {
	mechanism := types.Mechanism{}
	sql := fmt.Sprintf("select * from t_mechanism where name = \"%s\" ;", name)
	ok, err := m.engine.SQL(sql).Limit(1).Get(&mechanism)
	if err != nil {
		return &mechanism, err
	}
	if !ok {
		return &mechanism, nil
	}

	return &mechanism, err
}

func (m *Mysql) QueryCoinHolderCount(contractAddr string) (int, error) {
	count := 0
	sql := fmt.Sprintf("select count(*) from balance_erc20 where contract_addr = \"%s\" and balance != 0;", contractAddr)
	ok, err := m.engine.SQL(sql).Limit(1).Get(&count)
	if err != nil {
		return count, err
	}
	if !ok {
		return count, nil
	}

	return count, err
}

func (m *Mysql) GetBlockHeight() (int, error) {
	count := 0
	sql := fmt.Sprintf("select num from block order by num desc")
	ok, err := m.engine.SQL(sql).Limit(1).Get(&count)
	if err != nil {
		return count, err
	}
	if !ok {
		return count, nil
	}
	return count, err
}

func (m *Mysql) QueryCoinHolders(contractAddr string) ([]*types.Balance_Erc20, error) {
	balances := make([]*types.Balance_Erc20, 0)
	err := m.engine.Table("balance_erc20").Where("contract_addr = ?", contractAddr).Find(&balances)
	if err != nil {
		return nil, err
	}
	return balances, err
}

func (m *Mysql) GetCoinBalance(accountAdr string, contractAddr string) (string, error) {
	balacne := ""
	sql := fmt.Sprintf("select  balance from balance_erc20 where addr = \"%s\" and contract_addr = \"%s\"", accountAdr, contractAddr)
	ok, err := m.engine.SQL(sql).Limit(1).Get(&balacne)
	if err != nil {
		return balacne, err
	}
	if !ok {
		return balacne, nil
	}
	return balacne, err
}

func (m *Mysql) QueryBurnTxs(accountAddr string, contractAddr string) ([]*types.Tx, error) {
	txs := make([]*types.Tx, 0)
	sql := fmt.Sprintf("SELECT a.* FROM block_data_test.tx a,tx_log b where a.addr_from = \"%s\" and a.addr_to = \"\" and a.tx_hash = b.tx_hash and b.addr = \"%s\";", accountAddr, contractAddr)
	err := m.engine.SQL(sql).Find(&txs)
	if err != nil {
		return txs, nil
	}
	return txs, err
}

func (m *Mysql) QueryTxHistory(accountAddr string, contractAddr string) ([]*types.Tx, error) {
	txs := make([]*types.Tx, 0)
	sql := fmt.Sprintf("SELECT * FROM block_data_test.tx  where (addr_from = \"%s\" or addr_to = \"%s\" ) and  tx_hash in (select DISTINCT tx_hash   from tx_log  where addr = \"%s\") order by id desc;", accountAddr, accountAddr, contractAddr)
	err := m.engine.SQL(sql).Find(&txs)
	if err != nil {
		return txs, nil
	}
	return txs, err
}

func (m *Mysql) QueryTxErc20History(accountAddr string) ([]*types.Erc20Tx, error) {
	txs := make([]*types.Erc20Tx, 0)
	err := m.engine.Table("tx_erc20").Where("sender = ? or receiver = ? ", accountAddr, accountAddr).Find(&txs)
	if err != nil {
		return nil, err
	}
	return txs, err
}

func (m *Mysql) QuerySpecifyCoinInfo(contractAddr string) (*types.Erc20Info, error) {
	info := &types.Erc20Info{}
	ok, err := m.engine.Table("erc20_info").Where("addr = ?", contractAddr).Get(info)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return info, err
}

func (m *Mysql) QueryCoinInfos(accountAddr string) ([]*types.Erc20Info, error) {
	infos := make([]*types.Erc20Info, 0)
	err := m.engine.Join("INNER", "balance_erc20", "balance_erc20.contract_addr = erc20_info.addr").Where("balance_erc20.addr = ? ", accountAddr).OrderBy("balance_erc20.id desc").Find(&infos)
	if err != nil {
		return nil, err
	}
	return infos, err
}

func (m *Mysql) QueryAllCoinAllHolders(accountAddr string) (int, error) {
	count := 0
	str := fmt.Sprintf("SELECT count(*) FROM balance_erc20 where balance!=0 and contract_addr  in (SELECT contract_addr FROM balance_erc20 where addr = \"%s\")", accountAddr)
	ok, err := m.engine.SQL(str).Get(&count)
	if err != nil {
		return -1, err
	}
	if !ok {
		return -1, nil
	}
	return count, err
}

func (m *Mysql) QueryTxlogByHash(hash string) (*types.TxLog, error) {
	txLog := &types.TxLog{}
	ok, err := m.engine.Table("tx_log").Where("tx_hash = ? ", hash).Get(txLog)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return txLog, err
}

func (m *Mysql) GetEventHash() ([]*types.EventHash, error) {
	hashs := make([]*types.EventHash, 0)
	err := m.engine.Table("t_eventhash").Find(&hashs)
	if err != nil {
		return nil, err
	}
	return hashs, nil
}

func (m *Mysql) GetCoinInfo(contractAddr string) ([]*types.Erc20Tx, error) {
	tasks := make([]*types.Erc20Tx, 0)
	err := m.engine.Table("tx_erc20").Where("addr = ?", contractAddr).OrderBy("block_num").Find(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, err
}
