package db

import (
	"fmt"
	"github.com/ethereum/coin-manage/config"
	"github.com/ethereum/coin-manage/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Createdb() *sqlx.DB {
	dsn := config.Readconfig()
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return db
}

func QueryData(Db *sqlx.DB, contract_addr string) []util.HolderInfo {

	holderInfos := make([]util.HolderInfo, 0)

	str := fmt.Sprintf("select addr,balance from balance_erc20 where contract_addr = \"%s\"", contract_addr)
	rows, err := Db.Query(str)
	if err != nil {
		fmt.Printf("query faied, error:[%v]", err.Error())
		return nil
	}
	for rows.Next() {
		//定义变量接收查询数据
		var addr, balance string

		err := rows.Scan(&addr, &balance)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
			return nil
		}
		holderInfo := util.HolderInfo{
			Addr:          addr,
			Balance:       balance,
			Contract_addr: contract_addr,
		}
		holderInfos = append(holderInfos, holderInfo)

		fmt.Println(addr, balance)
	}

	//关闭结果集（释放连接）
	rows.Close()
	return holderInfos
}
