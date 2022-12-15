package main

import (
	"fmt"
	"log"

	util "github.com/ethereum/coin-manage/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//util "./contract" // for demo
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x148933B2A02F51bA249b323FEdAd9409124Ab4c4")
	instance, err := util.NewUtil(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	ret, err := instance.Zzz(nil)
	fmt.Println(ret)
}
