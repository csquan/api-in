package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/coin-manage/config"
	IAllERC20 "github.com/ethereum/coin-manage/contract"
	"github.com/ethereum/coin-manage/types"
	"github.com/ethereum/coin-manage/util"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const ADDRLEN = 42

const Ok = 0

type ApiService struct {
	db     types.IDB
	config *config.Config
}

func NewApiService(db types.IDB, cfg *config.Config) *ApiService {
	return &ApiService{
		db:     db,
		config: cfg,
	}
}

func (a *ApiService) Run() {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = []string{"*"}
	r.Use(func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")
		// ctx.Header("Access-Control-Allow-Headers", "Content-Type,addr,GoogleAuth,AccessToken,X-CSRF-Token,Authorization,Token,token,auth,x-token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	})

	//读mysql数据库
	r.GET("/getSpecifyCoinInfo/:contractAddr", a.getSpecifyCoinInfo)
	r.GET("/getCoinInfos/:accountAddr", a.getCoinInfos)
	r.GET("/getAllCoinAllCount/:accountAddr", a.getAllCoinAllCount)
	r.GET("/getCoinHolders/:contractAddr", a.getCoinHolders)
	r.GET("/getCoinHoldersCount/:contractAddr", a.getCoinHoldersCount)
	r.GET("/getTxHistory/:accountAddr", a.getTxHistory)
	r.GET("/getReceiver/:contractAddr", a.getReceiver)
	//写mysql数据库
	r.GET("/setReceiver/:contractAddr/:receiveAddr", a.setReceiver)

	//写合约
	r.POST("/addBlack", a.addBlack)
	r.POST("/addBlackIn", a.addBlackIn)
	r.POST("/addBlackOut", a.addBlackOut)
	r.POST("/frozen", a.frozen)
	r.POST("/addBlackRange", a.addBlackRange)
	r.POST("/mint", a.mint)
	r.POST("/burn", a.burn)

	//读取合约
	r.GET("/status/:accountAddr", a.status)
	r.GET("/model/:accountAddr", a.model)

	r.GET("/taxFee", a.GetTaxFee)
	r.GET("/bonusFee", a.GetBonusFee)

	r.POST("/tx/get", a.GetTask)

	err := r.Run(fmt.Sprintf(":%s", a.config.Server.Port))
	if err != nil {
		logrus.Fatalf("start http server err:%v", err)
	}
}

func checkAddr(addr string) error {
	if addr[:2] != "0x" {
		return errors.New("addr must start with 0x")
	}
	if len(addr) != ADDRLEN {
		return errors.New("addr len wrong ,must 40")
	}
	return nil
}

// 首先查询balance_erc20表，得到地址持有的代币合约地址，然后根据代币合约地址查erc20_info表
func (a *ApiService) getAllCoinAllCount(c *gin.Context) {
	addr := c.Param("accountAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	count, err := a.db.QueryAllCoinAllHolders(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = fmt.Sprintf("%d", count)
	c.SecureJSON(http.StatusOK, res)
}

// 首先查询balance_erc20表，得到地址持有的代币合约地址，然后根据代币合约地址查erc20_info表
func (a *ApiService) getSpecifyCoinInfo(c *gin.Context) {
	addr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	info, err := a.db.QuerySpecifyCoinInfo(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	b, err := json.Marshal(info)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = string(b)
	c.SecureJSON(http.StatusOK, res)
}

// 首先查询balance_erc20表，得到地址持有的代币合约地址，然后根据代币合约地址查erc20_info表
func (a *ApiService) getCoinInfos(c *gin.Context) {
	addr := c.Param("accountAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	baseInfos, err := a.db.QueryCoinInfos(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	coinInfos := make([]*types.CoinInfo, 0)

	//这里拼装每个代币的持币地址数 状态
	for _, info := range baseInfos {
		coinInfo := types.CoinInfo{
			BaseInfo: *info,
		}
		status, err := a.getStatus(info.Addr)
		if err != nil {
			fmt.Println(err)
		}
		coinInfo.Status = *status

		count, err := a.db.QueryCoinHolderCount(addr)
		if err != nil {
			fmt.Println(err)
		}
		coinInfo.HolderCount = count

		coinInfos = append(coinInfos, &coinInfo)
	}

	b, err := json.Marshal(coinInfos)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = string(b)
	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) getCoinHoldersCount(c *gin.Context) {
	addr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)

	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	holderInfos, err := a.db.QueryCoinHolderCount(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	b, err := json.Marshal(holderInfos)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	res.Code = Ok
	res.Message = "success"
	res.Data = string(b)
	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) getCoinHolders(c *gin.Context) {
	addr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)

	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	holderInfos, err := a.db.QueryCoinHolders(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	b, err := json.Marshal(holderInfos)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	res.Code = Ok
	res.Message = "success"
	res.Data = string(b)
	c.SecureJSON(http.StatusOK, res)
}

func addBlackData(method string, accountAddr common.Address) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		fmt.Println("read file err:", err.Error())
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	return contractAbi.Pack(method, accountAddr)
}

func forzenData(method string, accountAddr common.Address, amount int64) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		fmt.Println("read file err:", err.Error())
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	return contractAbi.Pack(method, accountAddr, amount)
}

func addblackRangeData(method string, blockRange IAllERC20.IFATERC20ConfigBlockRange) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		fmt.Println("read file err:", err.Error())
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	return contractAbi.Pack(method, blockRange)
}

func mintData(method string, receiverAddr string, amount string) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		fmt.Println("read file err:", err.Error())
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	return contractAbi.Pack(method, receiverAddr, amount)
}

func burnData(method string, amount string) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		fmt.Println("read file err:", err.Error())
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	return contractAbi.Pack(method, amount)
}

func parse(db types.IDB, input string, contractAddr string) (string, error) {
	//contractABI, err := db.QueryABI(contractAddr)
	//if err != nil {
	//	return "", err
	//}
	//contractAbi := GetABI(contractABI.Abi_data)
	//method, err := contractAbi.MethodById([]byte(input))
	//if err != nil {
	//	return "", err
	//}

	//return method.Name, nil
	return "testOp", nil
}

// 获取ABI对象
func GetABI(abiJSON string) abi.ABI {
	wrapABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic(err)
	}
	return wrapABI
}

func (a *ApiService) getTxHistory(c *gin.Context) {
	addr := c.Param("accountAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	TxInfos, err := a.db.QueryTxHistory(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	Erc20TxInfos, err := a.db.QueryTxErc20History(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	//动作-TxInfos从input中解析，Erc20TxInfos是属于内部交易，动作为转账
	txArray := make([]types.TxRes, 0)

	for _, tx := range TxInfos {
		txRes := types.TxRes{}

		txRes.Hash = tx.Hash
		txRes.TxGeneral = tx

		if tx.IsContractCreate == "1" {
			txRes.Op = "ContractCreate"
		} else {
			if tx.IsContract == "1" { //需要解析input
				//先由contractAddr去合约abi表中取到对应的abi
				//todo:parse OpAddr
				txRes.OpAddr = "need to do"
				txRes.Op, err = parse(a.db, tx.Input, tx.To)
				if err != nil {
					res.Code = http.StatusInternalServerError
					res.Message = err.Error()
					c.SecureJSON(http.StatusInternalServerError, res)
					return
				}
			} else {
				if tx.From == addr {
					txRes.Op = "TransferOut"
				} else {
					txRes.Op = "TransferIn"
				}
			}
		}
		txArray = append(txArray, txRes)
	}

	for _, tx := range Erc20TxInfos {
		txRes := types.TxRes{}

		txRes.Hash = tx.TxHash
		txRes.TxErc20 = tx

		if tx.Sender == addr {
			txRes.Op = "TransferOut"
		} else {
			txRes.Op = "TransferIn"
		}
		txArray = append(txArray, txRes)
	}

	b, err := json.Marshal(txArray)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	res.Code = Ok
	res.Message = "success"
	res.Data = string(b)
	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) getReceiver(c *gin.Context) {
	contract_addr := c.Param("contractAddr")

	res := types.HttpRes{}

	err := checkAddr(contract_addr)

	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	contract_receiver, err := a.db.QueryReceiver(contract_addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = contract_receiver.Receiver_Addr

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) setReceiver(c *gin.Context) {
	contract_addr := c.Param("contractAddr")
	receive_addr := c.Param("receiveAddr")

	res := types.HttpRes{}

	err := checkAddr(contract_addr)

	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(receive_addr)

	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	receiver := types.ContractReceiver{
		Contract_Addr: contract_addr,
		Receiver_Addr: receive_addr,
	}

	err = a.db.InsertReceiver(&receiver)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = Ok
	res.Message = "success"

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) GetTask(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}

	contractAddr := gjson.Get(data1, "contractAddr")
	accountAddr := gjson.Get(data1, "accountAddr")
	uuid := gjson.Get(data1, "uuid")

	res := types.HttpRes{}

	err := checkAddr(contractAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(accountAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	cli := resty.New()

	data := types.TxData{
		RequestID: "Hui-TxState",
		UUID:      uuid.String(),
		From:      accountAddr.String(),
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post("http://127.0.0.1:8080/tx/get")
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println(err)
	}
	if result.Code != 0 {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = result.Message

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) addBlack(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}

	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	res := types.HttpRes{}

	err := checkAddr(contractAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(targetAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := addBlackData("addBlack", common.HexToAddress(targetAddr.String()))

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: "Hui-TxState",
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post("http://127.0.0.1:8080/tx/create")
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println(err)
	}
	if result.Code != 0 {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}
func (a *ApiService) addBlackIn(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}

	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	res := types.HttpRes{}

	err := checkAddr(targetAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := addBlackData("addBlackIn", common.HexToAddress(targetAddr.String()))

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: "Hui-TxState",
		UID:       uid.String(),
		UUID:      time.Now().String(),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post("http://127.0.0.1:8080/tx/create")
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println(err)
	}
	if result.Code != 0 {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}
func (a *ApiService) addBlackOut(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}

	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	res := types.HttpRes{}

	err := checkAddr(targetAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := addBlackData("addBlackOut", common.HexToAddress(targetAddr.String()))

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: "Hui-TxState",
		UID:       uid.String(),
		UUID:      time.Now().String(),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post("http://127.0.0.1:8080/tx/create")
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println(err)
	}
	if result.Code != 0 {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}
func (a *ApiService) frozen(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	amount := gjson.Get(data1, "amount")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	res := types.HttpRes{}

	parseInt, err := strconv.ParseInt(amount.String(), 10, 64)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil || parseInt <= 0 {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(targetAddr.String())
	if err != nil || parseInt <= 0 {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := forzenData("frozen", common.HexToAddress(targetAddr.String()), parseInt)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: "Hui-TxState",
		UID:       uid.String(),
		UUID:      time.Now().String(),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post("http://127.0.0.1:8080/tx/create")
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println(err)
	}
	if result.Code != 0 {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) addBlackRange(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	startblock := gjson.Get(data1, "startblock")
	endblock := gjson.Get(data1, "endblock")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	uid := gjson.Get(data1, "uid")

	res := types.HttpRes{}

	parseStartPos, err := strconv.ParseInt(startblock.String(), 10, 64)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	parseEndPos, err := strconv.ParseInt(endblock.String(), 10, 64)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	br := IAllERC20.IFATERC20ConfigBlockRange{
		BeginBlock: big.NewInt(parseStartPos),
		EndBlock:   big.NewInt(parseEndPos),
	}

	inputData, err := addblackRangeData("addBlackRange", br)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: "Hui-TxState",
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post("http://127.0.0.1:8080/tx/create")
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println(err)
	}
	if result.Code != 0 {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) mint(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	amount := gjson.Get(data1, "amount")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	receiverAddr := gjson.Get(data1, "receiverAddr")
	uid := gjson.Get(data1, "uid")

	res := types.HttpRes{}

	parseInt, err := strconv.ParseInt(amount.String(), 10, 64)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil || parseInt <= 0 {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := mintData("mint", receiverAddr.String(), amount.String())

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: "Hui-TxState",
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post("http://127.0.0.1:8080/tx/create")
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println(err)
	}
	if result.Code != 0 {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) burn(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	amount := gjson.Get(data1, "amount")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	uid := gjson.Get(data1, "uid")

	res := types.HttpRes{}

	parseInt, err := strconv.ParseInt(amount.String(), 10, 64)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	if parseInt <= 0 {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := burnData("mint", amount.String())

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: "Hui-TxState",
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post("http://127.0.0.1:8080/tx/create")
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println(err)
	}
	if result.Code != 0 {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) getStatus(account string) (*types.StatusInfo, error) {
	instance, _ := util.PrepareTx(a.config)

	isBlack, err := instance.BlackOf(nil, common.HexToAddress(account))
	if err != nil {
		return nil, err
	}
	isBlackIn, err := instance.BlackInOf(nil, common.HexToAddress(account))
	if err != nil {
		return nil, err
	}
	isBlackOut, err := instance.BlackOutOf(nil, common.HexToAddress(account))
	if err != nil {
		return nil, err
	}
	nowFrozenAmount, err := instance.FrozenOf(nil, common.HexToAddress(account))
	if err != nil {
		return nil, err
	}
	waitFrozenAmount, err := instance.WaitFrozenOf(nil, common.HexToAddress(account))
	if err != nil {
		return nil, err
	}

	status := types.StatusInfo{
		IsBlack:          isBlack,
		IsBlackIn:        isBlackIn,
		IsBlackOut:       isBlackOut,
		NowFrozenAmount:  nowFrozenAmount,
		WaitFrozenAmount: waitFrozenAmount,
	}
	return &status, nil
}

func (a *ApiService) status(c *gin.Context) {
	account := c.Param("accountAddr")

	res := types.HttpRes{}

	err := checkAddr(account)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	instance, _ := util.PrepareTx(a.config)

	isBlack, err := instance.BlackOf(nil, common.HexToAddress(account))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	isBlackIn, err := instance.BlackInOf(nil, common.HexToAddress(account))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	isBlackOut, err := instance.BlackOutOf(nil, common.HexToAddress(account))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	nowFrozenAmount, err := instance.FrozenOf(nil, common.HexToAddress(account))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	waitFrozenAmount, err := instance.WaitFrozenOf(nil, common.HexToAddress(account))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	status := types.StatusInfo{
		IsBlack:          isBlack,
		IsBlackIn:        isBlackIn,
		IsBlackOut:       isBlackOut,
		NowFrozenAmount:  nowFrozenAmount,
		WaitFrozenAmount: waitFrozenAmount,
	}

	log.Printf("status %v", status)

	b, err := json.Marshal(status)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = string(b)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) model(c *gin.Context) {

	res := types.HttpRes{}

	instance, _ := util.PrepareTx(a.config)

	modelValue, err := instance.Model(nil)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	log.Printf("modelValue %d", modelValue)

	res.Code = Ok
	res.Message = "success"
	res.Data = fmt.Sprintf("%d", modelValue)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) GetTaxFee(c *gin.Context) {

	res := types.HttpRes{}

	instance, _ := util.PrepareTx(a.config)

	taxFee, err := instance.GetTaxFee(nil)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	log.Printf("taxFee %d", taxFee)

	res.Code = Ok
	res.Message = "success"
	res.Data = taxFee.String()

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) GetBonusFee(c *gin.Context) {

	res := types.HttpRes{}

	instance, _ := util.PrepareTx(a.config)

	bonusFee, err := instance.GetBonusFee(nil)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	log.Printf("bonusFee %d", bonusFee)

	res.Code = Ok
	res.Message = "success"
	res.Data = bonusFee.String()

	c.SecureJSON(http.StatusOK, res)
}
