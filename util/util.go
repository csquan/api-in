package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/coin-manage/config"
	"log"
	"math/big"

	IERC20 "github.com/ethereum/coin-manage/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CoinInfo struct {
	Name         string `yaml:"name"`
	Symbol       string `yaml:"symbol"`
	ContractAddr string `yaml:"contract_addr"`
	Decimal      string `yaml:"decimal"`
	Total_Supply string `yaml:"total_supply"`
}

type HolderInfo struct {
	Addr          string `yaml:"addr"`
	Balance       string `yaml:"balance"`
	Contract_addr string `yaml:"contract_addr"`
}

type HistoryInfo struct {
	Symbol  string `yaml:"symbol"`
	Time    string `yaml:"time"`
	Balance string `yaml:"balance"`
	Op      string `yaml:"op"`
	Amount  string `yaml:"amount"`
	Params  string `yaml:"params"`
}

type BlockRange struct {
	BeginBlock *big.Int
	EndBlock   *big.Int
}

func PrepareTx(cfg *config.Config) (*IERC20.IAllERC20, *bind.TransactOpts) {
	endpoint := fmt.Sprintf("http://%s:%s", cfg.Endpoint.Ip, cfg.Endpoint.Port)

	client, err := ethclient.Dial(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	//私钥
	privateKey, err := crypto.HexToECDSA(cfg.Pri.Value)
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(8888))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(8000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x8356d3280bfffdd4a1c0c71b5e2cfb68b1b93df8")
	instance, err := IERC20.NewIAllERC20(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance, auth
}
