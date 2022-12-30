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
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	r.GET("/getCoinBalance/:accountAddr/:contractAddr", a.getCoinBalance)

	r.GET("/getCoinHoldersCount/:contractAddr", a.getCoinHoldersCount)
	r.GET("/getTxHistory/:accountAddr/:contractAddr", a.getTxHistory)
	r.GET("/hasBurnAmount/:accountAddr/:contractAddr", a.hasBurnAmount)

	r.GET("/getBlockHeight", a.getBlockHeight)

	//写合约
	r.POST("/addBlack", a.addBlack)
	r.POST("/removeBlack", a.removeBlack)
	r.POST("/addBlackIn", a.addBlackIn)
	r.POST("/removeBlackIn", a.removeBlackIn)
	r.POST("/addBlackOut", a.addBlackOut)
	r.POST("/removeBlackOut", a.removeBlackOut)
	r.POST("/frozen", a.frozen)
	r.POST("/unfrozen", a.unfrozen)

	r.POST("/addBlackRange", a.addBlackRange)
	r.POST("/removeBlackRange", a.removeBlackRange)

	r.POST("/mint", a.mint)
	r.POST("/burn", a.burn)
	r.POST("/burnFrom", a.burnFrom)

	//读取合约
	r.POST("/status", a.status)
	r.POST("/blackRange", a.blackRange)

	r.POST("/hasForzenAmount", a.hasForzenAmount)

	r.POST("/cap", a.cap)
	r.POST("/taxFee", a.GetTaxFee)
	r.POST("/bonusFee", a.GetBonusFee)

	r.POST("/model", a.model)
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
func (a *ApiService) getCoinBalance(c *gin.Context) {
	accountAddr := c.Param("accountAddr")
	contractAddr := c.Param("contractAddr")

	res := types.HttpRes{}

	err := checkAddr(accountAddr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(contractAddr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	balance, err := a.db.GetCoinBalance(strings.ToLower(accountAddr), strings.ToLower(contractAddr))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = balance
	c.SecureJSON(http.StatusOK, res)
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

	count, err := a.db.QueryAllCoinAllHolders(strings.ToLower(addr))
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

	info, err := a.db.QuerySpecifyCoinInfo(strings.ToLower(addr))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	//处理下 info的精度
	info.Totoal_Supply = HandleAmountDecimals(info.Totoal_Supply, info.Decimals)
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

func HandleAmountDecimals(amount string, decimal string) string {
	decimalInt, err := strconv.ParseInt(decimal, 10, 64)
	if err != nil {
	}
	pos := decimalInt - 2
	endpos := len(amount) - int(pos)

	str := amount[:endpos-2] + "." + "00"
	return str
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

	baseInfos, err := a.db.QueryCoinInfos(strings.ToLower(addr))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	coinInfos := make([]*types.CoinInfo, 0)

	height, err := a.db.GetBlockHeight()
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	//这里拼装每个代币的持币地址数 状态
	for _, info := range baseInfos {
		coinInfo := types.CoinInfo{
			BaseInfo: *info,
			Status:   1, //正常交易
		}

		instance, _ := util.PrepareTx(a.config, info.Addr)

		blackRange, err := instance.BlackBlocks(nil)
		if err != nil && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" {
			fmt.Println(err)
			continue
		}
		for _, rangeValue := range blackRange {
			if height >= int(rangeValue.BeginBlock.Int64()) || height <= int(rangeValue.EndBlock.Int64()) {
				coinInfo.Status = 0 //暂停交易
			}
		}

		count, err := a.db.QueryCoinHolderCount(strings.ToLower(info.Addr))
		if err != nil {
			fmt.Println(err)
		}
		coinInfo.HolderCount = count

		coinInfo.BaseInfo.Totoal_Supply = HandleAmountDecimals(coinInfo.BaseInfo.Totoal_Supply, coinInfo.BaseInfo.Decimals)
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

	holderInfos, err := a.db.QueryCoinHolderCount(strings.ToLower(addr))
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

	infos := make([]*types.Balance_Erc20, 0)

	holderInfos, err := a.db.QueryCoinHolders(strings.ToLower(addr))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	//过滤Addr空地址
	for _, holderInfo := range holderInfos {
		if holderInfo.Balance != "0" {
			infos = append(infos, holderInfo)
		}
	}

	b, err := json.Marshal(infos)
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

func forzenData(method string, accountAddr common.Address, amount *big.Int) ([]byte, error) {
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

func removeblackRangeData(method string, index *big.Int) ([]byte, error) {
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
	return contractAbi.Pack(method, index)
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

func mintData(method string, receiverAddr common.Address, amount *big.Int) ([]byte, error) {
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

func burnData(method string, amount *big.Int) ([]byte, error) {
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

func burnFromData(method string, targetAddr common.Address, amount *big.Int) ([]byte, error) {
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
	return contractAbi.Pack(method, targetAddr, amount)
}

func parse(db types.IDB, txhash string) (*types.OpParam, error) {
	op := ""
	//首先在tx_log中找到这笔hash对应的交易，比对op表中的hash，看是哪个动作，取出对应的参数个数和参数格式
	tx_log, err := db.QueryTxlogByHash(txhash)
	if err != nil {
		fmt.Println(err)
	}
	eventHashs, err := db.GetEventHash()
	if err != nil {
		fmt.Println(err)
	}
	if tx_log == nil {
		fmt.Println("null")
		return nil, nil
	}
	opparam := types.OpParam{}

	for _, value := range eventHashs {
		if tx_log.Topic0 == "0x"+value.EventHash { //找到动作,然后依据格式解析参数
			op = value.Op
			opparam.Op = op
			switch op {
			case "AddBlack": //event AddBlack(address account);
				opparam.Addr1 = formatHex(tx_log.Data)
				break
			case "RemoveBlack": // event RemoveBlack(address account);
				opparam.Addr1 = formatHex(tx_log.Data)
				break
			case "AddBlackIn": // event AddBlackIn(address account);
				opparam.Addr1 = formatHex(tx_log.Data)
				break
			case "RemoveBlackIn": // event RemoveBlackIn(address account);
				opparam.Addr1 = formatHex(tx_log.Data)
				break
			case "AddBlackOut": //event AddBlackOut(address account);
				opparam.Addr1 = formatHex(tx_log.Data)
				break
			case "RemoveBlackOut": //event RemoveBlackOut(address account);
				opparam.Addr1 = formatHex(tx_log.Data)
				break
			case "AddBlackBlock": //这里tx_log.Data 含有2个uint128参数- event AddBlackBlock(uint128 _beginBlock, uint128 _endBlock);
				valueStr1 := formatHex(tx_log.Data[:64])
				if valueStr1 == "0x" {
					opparam.Value1 = "0"
				} else {
					value1, err := hexutil.DecodeBig(valueStr1)
					if err != nil {
						fmt.Println(err)
					}
					opparam.Value1 = value1.String()
				}

				valueStr2 := formatHex(tx_log.Data[64:])
				if valueStr2 == "0x" {
					opparam.Value2 = "0"
				} else {
					value2, err := hexutil.DecodeBig(valueStr2)
					if err != nil {
						fmt.Println(err)
					}
					opparam.Value2 = value2.String()
				}
				break
			case "RemoveBlackBlock": //这里tx_log.Data 含有3个uint参数- event RemoveBlackBlock(uint256 i, uint128 _beginBlock, uint128 _endBlock);
				valueStr1 := formatHex(tx_log.Data[:64])
				if valueStr1 == "0x" {
					opparam.Value1 = "0"
				} else {
					value1, err := hexutil.DecodeBig(valueStr1)
					if err != nil {
						fmt.Println(err)
					}
					opparam.Value1 = value1.String()
				}

				valueStr2 := formatHex(tx_log.Data[64:128])
				if valueStr2 == "0x" {
					opparam.Value2 = "0"
				} else {
					value2, err := hexutil.DecodeBig(valueStr2)
					if err != nil {
						fmt.Println(err)
					}
					opparam.Value2 = value2.String()
				}

				valueStr3 := formatHex(tx_log.Data[128:])
				if valueStr3 == "0x" {
					opparam.Value3 = "0"
				} else {
					value3, err := hexutil.DecodeBig(valueStr3)
					if err != nil {
						fmt.Println(err)
					}
					opparam.Value3 = value3.String()
				}
				break
			case "Frozen": //这里tx_log.Data 含有后2个uint128参数- event Frozen(address indexed account, uint256 frozen, uint256 waitFrozen);
				param1 := common.HexToAddress(tx_log.Topic1)
				opparam.Addr1 = param1.Hex()

				valueStr1 := formatHex(tx_log.Data[:64])
				if valueStr1 == "0x" {
					opparam.Value1 = "0"
				} else {
					value1, err := hexutil.DecodeBig(valueStr1)
					if err != nil {
						fmt.Println(err)
					}
					opparam.Value1 = value1.String()
				}

				valueStr2 := formatHex(tx_log.Data[64:128])
				if valueStr2 == "0x" {
					opparam.Value2 = "0"
				} else {
					value2, err := hexutil.DecodeBig(valueStr2)
					if err != nil {
						fmt.Println(err)
					}
					opparam.Value2 = value2.String()
				}
				break
			case "Transfer": // event Transfer(address indexed from, address indexed to, uint256 value);
				param1 := common.HexToAddress(tx_log.Topic1)
				opparam.Addr1 = param1.String()

				param2 := common.HexToAddress(tx_log.Topic2)
				opparam.Addr2 = param2.String()

				valueStr := formatHex(tx_log.Data)
				value, err := hexutil.DecodeBig(valueStr)
				if err != nil {
					fmt.Println(err)
				}
				opparam.Value1 = value.String()
				break
			case "UnFrozen": //这里tx_log.Data 含有后2个uint128参数- event Frozen(address indexed account, uint256 frozen, uint256 waitFrozen);
				param1 := common.HexToAddress(tx_log.Topic1)
				opparam.Addr1 = param1.Hex()

				valueStr1 := formatHex(tx_log.Data[:64])
				if valueStr1 == "0x" {
					opparam.Value1 = "0"
				} else {
					value1, err := hexutil.DecodeBig(valueStr1)
					if err != nil {
						fmt.Println(err)
					}
					opparam.Value1 = value1.String()
				}

				valueStr2 := formatHex(tx_log.Data[64:128])
				if valueStr2 == "0x" {
					opparam.Value2 = "0"
				} else {
					value2, err := hexutil.DecodeBig(valueStr2)
					if err != nil {
						fmt.Println(err)
					}
					opparam.Value2 = value2.String()
				}

				break
			case "Paused": //event Paused(address account);
				param1 := common.HexToAddress(tx_log.Data)
				opparam.Addr1 = param1.String()
				break
			case "Unpaused": //event Unpaused(address account);
				param1 := common.HexToAddress(tx_log.Data)
				opparam.Addr1 = param1.String()
				break
			}
		}
	}

	return &opparam, nil
}

func formatHex(hexstr string) string {
	res := strings.TrimLeft(hexstr[2:], "0")
	return "0x" + res
}

func (a *ApiService) hasBurnAmount(c *gin.Context) {
	accountAddr := c.Param("accountAddr")
	contractAddr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(accountAddr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	err = checkAddr(contractAddr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	Txs, err := a.db.QueryBurnTxs(strings.ToLower(accountAddr), strings.ToLower(contractAddr))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	var sum int64
	for _, tx := range Txs {
		parseInt, err := strconv.ParseInt(tx.Value, 10, 64)
		if err != nil {
			res.Code = http.StatusBadRequest
			res.Message = err.Error()
			c.SecureJSON(http.StatusBadRequest, res)
			return
		}
		sum += parseInt
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = fmt.Sprintf("%d", sum)
	c.SecureJSON(http.StatusOK, res)
}

func copyStruct(paramDest *types.OpParam, paramSrc *types.OpParam) {
	paramDest.Op = paramSrc.Op
	paramDest.Value1 = paramSrc.Value1
	paramDest.Value2 = paramSrc.Value2
	paramDest.Value3 = paramSrc.Value3
	paramDest.Addr1 = paramSrc.Addr1
	paramDest.Addr2 = paramSrc.Addr2
}

func (a *ApiService) getTxHistory(c *gin.Context) {
	accountAddr := c.Param("accountAddr")
	contractAddr := c.Param("contractAddr")
	res := types.HttpRes{}

	err := checkAddr(accountAddr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	err = checkAddr(contractAddr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	TxInfos, err := a.db.QueryTxHistory(strings.ToLower(accountAddr), strings.ToLower(contractAddr))
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	//Erc20TxInfos, err := a.db.QueryTxErc20History(strings.ToLower(addr))
	//if err != nil {
	//	res.Code = http.StatusInternalServerError
	//	res.Message = err.Error()
	//	c.SecureJSON(http.StatusInternalServerError, res)
	//	return
	//}

	//动作-TxInfos从input中解析，Erc20TxInfos是属于内部交易，动作为转账
	txArray := make([]types.TxRes, 0)

	for _, tx := range TxInfos {
		txRes := types.TxRes{}

		parseUInt, err := strconv.ParseUint(tx.Value, 10, 64)
		if err != nil {
			continue
		}
		txRes.Amount = parseUInt

		txRes.Hash = tx.Hash
		txRes.TxGeneral = tx
		opparam := types.OpParam{}

		if tx.IsContractCreate == "1" {
			opparam.Op = "ContractCreate"
		} else {
			if tx.IsContract == "1" { //需要解析input
				param, err := parse(a.db, tx.Hash)
				if err != nil {
					res.Code = http.StatusInternalServerError
					res.Message = err.Error()
					c.SecureJSON(http.StatusInternalServerError, res)
					return
				}
				if param != nil {
					copyStruct(&opparam, param)
				}

				if opparam.Op == "Transfer" {
					if tx.From == accountAddr {
						opparam.Op = "TransferOut"

						if tx.To == "" {
							opparam.Op = "Destroy"
						}
					} else {
						opparam.Op = "TransferIn"

						if tx.From == "" {
							opparam.Op = "Increase"
						}
					}
				}
			} else {
				if tx.From == accountAddr {
					opparam.Op = "TransferOut"

					if tx.To == "" {
						opparam.Op = "Destroy"
					}
				} else {
					opparam.Op = "TransferIn"

					if tx.From == "" {
						opparam.Op = "Increase"
					}
				}
			}
		}
		txRes.OpParams = &opparam
		txArray = append(txArray, txRes)
	}

	//for _, tx := range Erc20TxInfos {
	//	txRes := types.TxRes{}
	//	opparam := types.OpParam{}
	//
	//	txRes.Hash = tx.TxHash
	//	txRes.TxErc20 = tx
	//
	//	if tx.Sender == addr {
	//		opparam.Op = "TransferOut"
	//	} else {
	//		opparam.Op = "TransferIn"
	//	}
	//	b, err := json.Marshal(opparam)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	txRes.OpParams = string(b)
	//	txArray = append(txArray, txRes)
	//}

	b, err := json.Marshal(txArray)

	res.Code = Ok
	res.Message = "success"
	res.Data = json.RawMessage(b)
	c.JSON(http.StatusOK, res)
}

func (a *ApiService) getBlockHeight(c *gin.Context) {

	res := types.HttpRes{}

	count, err := a.db.GetBlockHeight()
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
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UUID:      uuid.String(),
		From:      accountAddr.String(),
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/get")
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

	inputData, err := addBlackData("removeBlack", common.HexToAddress(targetAddr.String()))

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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

func (a *ApiService) removeBlack(c *gin.Context) {
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
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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

func (a *ApiService) removeBlackIn(c *gin.Context) {
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

	inputData, err := addBlackData("removeBlackIn", common.HexToAddress(targetAddr.String()))

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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
func (a *ApiService) removeBlackOut(c *gin.Context) {
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

	inputData, err := addBlackData("removeBlackOut", common.HexToAddress(targetAddr.String()))

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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

func (a *ApiService) unfrozen(c *gin.Context) {
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

	Amount := &big.Int{}
	Amount.SetInt64(parseInt)

	inputData, err := forzenData("unfrozen", common.HexToAddress(targetAddr.String()), Amount)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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

	//b := make([]byte, 32)
	//binary.BigEndian.PutUint64(b, uint64(parseInt))

	Amount := &big.Int{}
	Amount.SetInt64(parseInt)

	inputData, err := forzenData("frozen", common.HexToAddress(targetAddr.String()), Amount)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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

	inputData, err := addblackRangeData("addBlackBlock", br)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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

func (a *ApiService) removeBlackRange(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	index := gjson.Get(data1, "index")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	uid := gjson.Get(data1, "uid")

	res := types.HttpRes{}

	indexPos, err := strconv.ParseInt(index.String(), 10, 64)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	Amount := &big.Int{}
	Amount.SetInt64(indexPos)

	inputData, err := removeblackRangeData("removeBlackBlock", Amount)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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

	Amount := &big.Int{}
	Amount.SetInt64(parseInt)

	inputData, err := mintData("mint", common.HexToAddress(operatorAddr.String()), Amount)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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

func (a *ApiService) burnFrom(c *gin.Context) {
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

	if parseInt <= 0 {
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
	err = checkAddr(contractAddr.String())
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

	Amount := &big.Int{}
	Amount.SetInt64(parseInt)

	inputData, err := burnFromData("mint", common.HexToAddress(targetAddr.String()), Amount)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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

	Amount := &big.Int{}
	Amount.SetInt64(parseInt)

	inputData, err := burnData("burn", Amount)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   "0x22B8",
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
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

func (a *ApiService) getStatus(contractAddr string, accountAddr string) (*types.StatusInfo, error) {
	instance, _ := util.PrepareTx(a.config, contractAddr)

	isBlack, err := instance.BlackOf(nil, common.HexToAddress(accountAddr))
	if err != nil {
		return nil, err
	}
	isBlackIn, err := instance.BlackInOf(nil, common.HexToAddress(accountAddr))
	if err != nil {
		return nil, err
	}
	isBlackOut, err := instance.BlackOutOf(nil, common.HexToAddress(accountAddr))
	if err != nil {
		return nil, err
	}
	nowFrozenAmount, err := instance.FrozenOf(nil, common.HexToAddress(accountAddr))
	if err != nil {
		return nil, err
	}
	waitFrozenAmount, err := instance.WaitFrozenOf(nil, common.HexToAddress(accountAddr))
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

func (a *ApiService) cap(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	contractAddr := gjson.Get(data1, "contractAddr")

	res := types.HttpRes{}

	err := checkAddr(contractAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	instance, _ := util.PrepareTx(a.config, contractAddr.String())

	capValue, err := instance.Cap(nil)
	if err != nil && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" && err.Error() != "execution reverted" {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = Ok
	res.Message = "success"

	if capValue == nil {
		res.Data = "0"
	} else {
		res.Data = capValue.String()
	}

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) hasForzenAmount(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	contractAddr := gjson.Get(data1, "contractAddr")
	accountAddr := gjson.Get(data1, "accountAddr")

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

	instance, _ := util.PrepareTx(a.config, contractAddr.String())

	FrozenAmount, err := instance.FrozenOf(nil, common.HexToAddress(accountAddr.String()))
	if err != nil && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = Ok
	res.Message = "success"
	res.Data = FrozenAmount.String()

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) blackRange(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	contractAddr := gjson.Get(data1, "contractAddr")

	res := types.HttpRes{}

	err := checkAddr(contractAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	instance, _ := util.PrepareTx(a.config, contractAddr.String())

	blackRange, err := instance.BlackBlocks(nil)
	if err != nil && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	b, err := json.Marshal(blackRange)
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

func (a *ApiService) status(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	contractAddr := gjson.Get(data1, "contractAddr")
	accountAddr := gjson.Get(data1, "accountAddr")

	res := types.HttpRes{}

	err := checkAddr(accountAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	instance, _ := util.PrepareTx(a.config, contractAddr.String())

	isBlack, err := instance.BlackOf(nil, common.HexToAddress(accountAddr.String()))
	if err != nil && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	isBlackIn, err := instance.BlackInOf(nil, common.HexToAddress(accountAddr.String()))
	if err != nil && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	isBlackOut, err := instance.BlackOutOf(nil, common.HexToAddress(accountAddr.String()))
	if err != nil && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	nowFrozenAmount, err := instance.FrozenOf(nil, common.HexToAddress(accountAddr.String()))
	if err != nil && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	waitFrozenAmount, err := instance.WaitFrozenOf(nil, common.HexToAddress(accountAddr.String()))
	if err != nil && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" {
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
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	contractAddr := gjson.Get(data1, "contractAddr")

	res := types.HttpRes{}

	instance, _ := util.PrepareTx(a.config, contractAddr.String())

	modelValue, err := instance.Model(nil)
	if err != nil && err.Error() != "no contract code at given address" && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" && err.Error() != "execution reverted" {
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
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	contractAddr := gjson.Get(data1, "contractAddr")
	res := types.HttpRes{}

	instance, _ := util.PrepareTx(a.config, contractAddr.String())

	taxFee, err := instance.GetTaxFee(nil)
	if err != nil && err.Error() != "no contract code at given address" && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" && err.Error() != "execution reverted" {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = Ok
	res.Message = "success"
	if taxFee == nil {
		res.Data = "-1"
	} else {
		res.Data = taxFee.String()
	}

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) GetBonusFee(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])

	isValid := gjson.Valid(data1)
	if isValid == false {
		fmt.Println("Not valid json")
	}
	contractAddr := gjson.Get(data1, "contractAddr")

	res := types.HttpRes{}

	instance, _ := util.PrepareTx(a.config, contractAddr.String())

	bonusFee, err := instance.GetBonusFee(nil)
	if err != nil && err.Error() != "no contract code at given address" && err.Error() != "abi: attempting to unmarshall an empty string while arguments are expected" && err.Error() != "execution reverted" {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	if bonusFee == nil {
		res.Data = "-1"
	} else {
		res.Data = bonusFee.String()
	}
	res.Code = Ok
	res.Message = "success"

	c.SecureJSON(http.StatusOK, res)
}
