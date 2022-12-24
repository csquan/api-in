package api

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/coin-manage/config"
	IAllERC20 "github.com/ethereum/coin-manage/contract"
	"github.com/ethereum/coin-manage/types"
	"github.com/ethereum/coin-manage/util"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

const ADDRLEN = 42

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

	//查询mysql数据库
	r.GET("/getCoinInfos/:accountAddr", a.getCoinInfos)
	r.GET("/getCoinHolders/:contractAddr", a.getCoinHolders)
	r.GET("/getTxHistory/:accountAddr", a.getTxHistory)

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
func (a *ApiService) getCoinInfos(c *gin.Context) {
	addr := c.Param("accountAddr")
	res := types.HttpRes{}

	err := checkAddr(addr)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
	}

	coinInfos, err := a.db.QueryCoinInfos(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
	}

	b, err := json.Marshal(coinInfos)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
	}
	res.Code = http.StatusOK
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
	}

	holderInfos, err := a.db.QueryCoinHolders(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
	}

	b, err := json.Marshal(holderInfos)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
	}
	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(b)
	c.SecureJSON(http.StatusOK, res)
}

func parse(db types.IDB, input string, contractAddr string) (string, error) {
	contractABI, err := db.QueryABI(contractAddr)
	if err != nil {
		return "", err
	}
	contractAbi := GetABI(contractABI.Abi_data)
	method, err := contractAbi.MethodById([]byte(input))
	if err != nil {
		return "", err
	}

	return method.Name, nil
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
	}

	TxInfos, err := a.db.QueryTxHistory(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
	}

	Erc20TxInfos, err := a.db.QueryTxErc20History(addr)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
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
	}
	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(b)
	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) addBlack(c *gin.Context) {
	account := c.PostForm("account")

	err := checkAddr(account)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	instance, auth := util.PrepareTx(a.config)

	tx, err := instance.AddBlack(auth, common.HexToAddress(account))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}
	log.Printf("addBlack Hash %s", tx.Hash())

	c.JSON(200, gin.H{
		"success": tx.Hash(),
	})
}
func (a *ApiService) addBlackIn(c *gin.Context) {
	account := c.PostForm("account")

	err := checkAddr(account)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	instance, auth := util.PrepareTx(a.config)

	tx, err := instance.AddBlackIn(auth, common.HexToAddress(account))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	log.Printf("addBlackIn Hash %s", tx.Hash())

	c.JSON(200, gin.H{
		"success": tx.Hash(),
	})
}
func (a *ApiService) addBlackOut(c *gin.Context) {
	account := c.PostForm("account")

	err := checkAddr(account)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	instance, auth := util.PrepareTx(a.config)

	tx, err := instance.AddBlackOut(auth, common.HexToAddress(account))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	log.Printf("addBlackOut Hash %s", tx.Hash())

	c.JSON(200, gin.H{
		"success": tx.Hash(),
	})
}
func (a *ApiService) frozen(c *gin.Context) {
	account := c.PostForm("account")
	amount := c.PostForm("amount")

	log.Println("in frozen ")

	parseInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	err = checkAddr(account)
	if err != nil || parseInt <= 0 {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	instance, auth := util.PrepareTx(a.config)

	tx, err := instance.Frozen(auth, common.HexToAddress(account), big.NewInt(parseInt))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	log.Printf("Frozen Hash %s", tx.Hash())

	c.JSON(200, gin.H{
		"success": tx.Hash(),
	})
}

func (a *ApiService) addBlackRange(c *gin.Context) {
	startblock := c.PostForm("startblock")
	endblock := c.PostForm("endblock")

	parseStartPos, err := strconv.ParseInt(startblock, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	parseEndPos, err := strconv.ParseInt(endblock, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	instance, auth := util.PrepareTx(a.config)

	br := IAllERC20.IFATERC20ConfigBlockRange{
		BeginBlock: big.NewInt(parseStartPos),
		EndBlock:   big.NewInt(parseEndPos),
	}
	tx, err := instance.AddBlackBlock(auth, br)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	log.Printf("addBlackRange Hash %s", tx.Hash())

	c.JSON(200, gin.H{
		"success": tx.Hash(),
	})
}

func (a *ApiService) mint(c *gin.Context) {
	account := c.PostForm("account")
	amount := c.PostForm("amount")

	parseInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	err = checkAddr(account)
	if err != nil || parseInt <= 0 {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	instance, auth := util.PrepareTx(a.config)

	tx, err := instance.Mint(auth, common.HexToAddress(account), big.NewInt(parseInt))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	log.Printf("mint Hash %s", tx.Hash())

	c.JSON(200, gin.H{
		"success": tx.Hash(),
	})
}

func (a *ApiService) burn(c *gin.Context) {
	amount := c.PostForm("amount")

	parseInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	if parseInt <= 0 {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	instance, auth := util.PrepareTx(a.config)

	tx, err := instance.Burn(auth, big.NewInt(parseInt))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	log.Printf("burn Hash %s", tx.Hash())

	c.JSON(200, gin.H{
		"success": tx.Hash(),
	})
}

func (a *ApiService) status(c *gin.Context) {
	account := c.Param("accountAddr")

	instance, _ := util.PrepareTx(a.config)

	isBlack, err := instance.BlackOf(nil, common.HexToAddress(account))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}
	isBlackIn, err := instance.BlackInOf(nil, common.HexToAddress(account))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}
	isBlackOut, err := instance.BlackOutOf(nil, common.HexToAddress(account))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}
	nowFrozenAmount, err := instance.FrozenOf(nil, common.HexToAddress(account))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}
	waitFrozenAmount, err := instance.WaitFrozenOf(nil, common.HexToAddress(account))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}

	status := types.StatusInfo{
		IsBlack:          isBlack,
		IsBlackIn:        isBlackIn,
		IsBlackOut:       isBlackOut,
		NowFrozenAmount:  nowFrozenAmount,
		WaitFrozenAmount: waitFrozenAmount,
	}

	log.Printf("status %v", status)

	c.JSON(200, gin.H{
		"success": status,
	})
}
