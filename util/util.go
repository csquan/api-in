package util

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/coin-manage/config"
	IERC20 "github.com/ethereum/coin-manage/contract"
	"github.com/ethereum/coin-manage/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"unsafe"
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

func PrepareTx(info *config.ChainInfo, cfg *config.Config, contractAddr string) (*IERC20.IAllERC20, *bind.TransactOpts) {
	client, err := ethclient.Dial(info.Rpc)
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

	//chainId := info.ChainId
	////这里先取出对应的chainId十六进制字符串，如果有则去除0x前缀，然后
	//if chainId[:1] == "0x" {
	//	chainId = chainId[1:]
	//}
	//parseUint, err := strconv.ParseInt(chainId, 16, 32)
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println(parseUint)
	//}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(8888))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(8000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contractAddr)
	instance, err := IERC20.NewIAllERC20(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance, auth
}

//func HttpGet(url string) (string, error) {
//	response, err := http.Get(url)
//	if err != nil {
//		logrus.Error(err)
//	}
//	defer response.Body.Close()
//	body, err2 := ioutil.ReadAll(response.Body)
//	if err2 != nil {
//		logrus.Error(err2)
//	}
//	return string(body), err
//}

func HttpGet(url string, authorization string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Error(err)
	}
	req.Header.Set("authorization", authorization)
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func HttpPost(url string, data []byte, authorization string) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorization)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Println("ioutil read error")
	}
	return string(body), err
}

func Post(requestUrl string, bytesData []byte) (ret string, err error) {
	res, err := http.Post(requestUrl,
		"application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		logrus.Error(err)
		return "", err
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	str := (*string)(unsafe.Pointer(&content)) //转化为string,优化内存
	return *str, nil
}

func GetAccountId(url string, accountID string) (ret string, err error) {
	param := types.AccountParam{
		AccountId: accountID,
	}
	msg, err := json.Marshal(param)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	str, err := Post(url, msg)
	if err != nil {
		return "", err
	}
	accountAddr := gjson.Get(str, "eth")
	return accountAddr.String(), nil
}
