package db

import (
	_ "github.com/go-sql-driver/mysql"
)

//func QueryCoinholders(Db *sqlx.DB, contract_addr string) []util.HolderInfo {
//	holderInfos := make([]util.HolderInfo, 0)
//	str := fmt.Sprintf("select addr,balance from balance_erc20 where contract_addr = \"%s\"", contract_addr)
//	rows, err := Db.Query(str)
//	if err != nil {
//		fmt.Printf("query faied, error:[%v]", err.Error())
//		return nil
//	}
//	for rows.Next() {
//		//定义变量接收查询数据
//		var addr, balance string
//
//		err := rows.Scan(&addr, &balance)
//		if err != nil {
//			fmt.Println("get data failed, error:[%v]", err.Error())
//			return nil
//		}
//		holderInfo := util.HolderInfo{
//			Addr:          addr,
//			Balance:       balance,
//			Contract_addr: contract_addr,
//		}
//		holderInfos = append(holderInfos, holderInfo)
//	}
//
//	//关闭结果集（释放连接）
//	rows.Close()
//	return holderInfos
//}
//
//func QueryTxHistory(Db *sqlx.DB, account_addr string) []util.HolderInfo {
//	HistoryInfos := make([]util.HistoryInfo, 0)
//	str := fmt.Sprintf("SELECT* FROM tx where addr_from = \"%s\" or addr_to = \"%s\"", account_addr, account_addr)
//	rows, err := Db.Query(str)
//	if err != nil {
//		fmt.Printf("query faied, error:[%v]", err.Error())
//		return nil
//	}
//	for rows.Next() {
//		//定义变量接收查询数据
//		var id, params, op string
//		tx := util.Tx{}
//		//type Tx struct {
//		//	TxType           uint8
//		//	From             string
//		//	To               string
//		//	Hash             string
//		//	Index            int
//		//	Value            *big.Int
//		//	Input            string
//		//	Nonce            uint64
//		//	GasPrice         *big.Int
//		//	GasLimit         uint64
//		//	GasUsed          uint64
//		//	IsContract       bool
//		//	IsContractCreate bool
//		//	BlockTime        int64
//		//	BlockNum         uint64
//		//	BlockHash        string
//		//	ExecStatus       uint64
//		//}
//
//		err := rows.Scan(&id, &tx.To, &tx.From, &tx.Hash, &tx.Index, &tx.Value, &tx.Nonce, &tx.GasPrice, &tx.GasLimit, &tx.GasUsed, &tx.IsContract, &tx.IsContractCreate,&tx.BlockTime,&tx.BlockNum,&tx.BlockHash.&tx.ExecStatus,
//			&tx.CreateTime,&tx.BlockState,&tx.MaxFeePerGas,&tx.MaxPriorityFeePerGas,&tx.BurntFees,tx.TxType)
//		)
//		if err != nil {
//			fmt.Println("get data failed, error:[%v]", err.Error())
//			return nil
//		}
//
//		if is_contract == "0" { //普通转账
//			op = "transfer"
//		} else {
//			if is_contract_create == "1" { //创建合约交易
//				op = "create contract"
//			} else { //调用合约交易
//				op = "call contract"
//			}
//		}
//
//		HistoryInfo := util.HistoryInfo{
//			Symbol:  "HUI",
//			Time:    block_time,
//			Balance: 1000,
//			Op:      op,
//			Amount:  amount,
//			Params:  params,
//		}
//		HistoryInfos = append(HistoryInfos, HistoryInfo)
//	}
//
//	//关闭结果集（释放连接）
//	rows.Close()
//	return HistoryInfos
//}

//func QueryCoinInfos(Db *sqlx.DB, account_addr string) []util.HolderInfo {
//	coinInfos := make([]util.CoinInfo, 0)
//	str := fmt.Sprintf("SELECT addr, name, symbol, decimals, total_supply FROM erc20_info where addr in (SELECT contract_addr FROM balance_erc20 where addr =\"%s\") ", account_addr)
//	rows, err := Db.Query(str)
//	if err != nil {
//		fmt.Printf("query faied, error:[%v]", err.Error())
//		return nil
//	}
//	for rows.Next() {
//		//定义变量接收查询数据
//		var contract_addr, name, symbol, decimal, total_supply string
//
//		err := rows.Scan(&contract_addr, &name, &symbol, &decimal, &total_supply)
//		if err != nil {
//			fmt.Println("get data failed, error:[%v]", err.Error())
//			return nil
//		}
//		coinInfo := util.CoinInfo{
//			ContractAddr: contract_addr,
//			Name:         name,
//			Symbol:       symbol,
//			Decimal:      decimal,
//			Total_Supply: total_supply,
//		}
//		coinInfos = append(coinInfos, coinInfo)
//	}
//
//	//关闭结果集（释放连接）
//	rows.Close()
//	return coinInfos
//}
