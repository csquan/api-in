package util

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	IERC20 "github.com/ethereum/coin-manage/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type HolderInfo struct {
	Addr          string `yaml:"addr"`
	Balance       string `yaml:"balance"`
	Contract_addr string `yaml:"contract_addr"`
}

type BlockRange struct {
	BeginBlock *big.Int
	EndBlock   *big.Int
}

func PrepareTx() (*IERC20.IAllERC20, *bind.TransactOpts) {
	client, err := ethclient.Dial("43.198.66.226:8545")
	if err != nil {
		log.Fatal(err)
	}
	//私钥
	privateKey, err := crypto.HexToECDSA("2da57ce0e3f9f53c0f8004d791c220a667c3d13ccc2db76a381fd7be98d5b6ea")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xab542fd2d1f6cb46e02bdf19f9f9c6922cf9b270")
	instance, err := IERC20.NewIAllERC20(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance, auth
}
