package api

import (
	"fmt"
	"github.com/ethereum/coin-manage/config"
	IAllERC20 "github.com/ethereum/coin-manage/contract"
	"github.com/ethereum/coin-manage/types"
	"github.com/ethereum/coin-manage/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"strconv"
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
	r.GET("/getCinInfos/:accountAddr", a.getCoinInfos)
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

	err := r.Run(fmt.Sprintf(":%d", a.config.Server.Port))
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

func (a *ApiService) getCoinInfos(c *gin.Context) {
	addr := c.Param("accountAddr")

	err := checkAddr(addr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}
	coinInfos, err := a.db.QueryCoinInfos(addr)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	}
	c.JSON(200, gin.H{
		"success": coinInfos,
	})
}

func (a *ApiService) getCoinHolders(c *gin.Context) {
	addr := c.Param("contractAddr")

	err := checkAddr(addr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	holderInfos, err := a.db.QueryCoinholders(addr)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	}
	c.JSON(200, gin.H{
		"success": holderInfos,
	})
}

func (a *ApiService) getTxHistory(c *gin.Context) {
	addr := c.Param("account")

	err := checkAddr(addr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	holderInfos, err := a.db.QueryTxHistory(addr)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
	}
	c.JSON(200, gin.H{
		"success": holderInfos,
	})
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
