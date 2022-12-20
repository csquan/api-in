package main

import (
	IAllERC20 "github.com/ethereum/coin-manage/contract"
	"github.com/ethereum/coin-manage/db"
	"github.com/ethereum/coin-manage/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"strconv"
)

type Config struct {
	Data struct {
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
}

const CONTRACTLEN = 42

func main() {

	r := gin.Default()

	r.GET("/getCoinHolders/:contractAddr", getCoinHolders)
	r.POST("/addBlack", addBlack)
	r.POST("/addBlackIn", addBlackIn)
	r.POST("/addBlackOut", addBlackOut)
	r.POST("/frozen", frozen)
	r.POST("/addBlackRange", addBlackRange)
	r.POST("/mint", mint)
	r.POST("/burn", burn)

	r.Run(":8000")
}

func getCoinHolders(c *gin.Context) {
	addr := c.Param("contractAddr")
	if addr[:2] != "0x" {
		c.JSON(400, gin.H{
			"error": "start not with 0x",
		})
		return
	}
	if len(addr) != CONTRACTLEN {
		c.JSON(400, gin.H{
			"error": "wrong len",
		})
		return
	}
	dbconn := db.Createdb()
	defer dbconn.Close()

	holderInfos := db.QueryData(dbconn, addr)

	c.JSON(200, gin.H{
		"HolderInfos": holderInfos,
	})
}
func addBlack(c *gin.Context) {
	account := c.PostForm("account")

	instance, auth := util.PrepareTx()

	tx, err := instance.AddBlack(auth, common.HexToAddress(account))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"banAccount hash": tx.Hash(),
	})
}
func addBlackIn(c *gin.Context) {
	account := c.PostForm("account")

	instance, auth := util.PrepareTx()

	tx, err := instance.AddBlackIn(auth, common.HexToAddress(account))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"restriAccountIn hash": tx.Hash(),
	})
}
func addBlackOut(c *gin.Context) {
	account := c.PostForm("account")

	instance, auth := util.PrepareTx()

	tx, err := instance.AddBlackOut(auth, common.HexToAddress(account))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"restriAccountOut hash": tx.Hash(),
	})
}
func frozen(c *gin.Context) {
	account := c.PostForm("account")
	amount := c.PostForm("amonut")

	parseInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	instance, auth := util.PrepareTx()

	tx, err := instance.Frozen(auth, common.HexToAddress(account), big.NewInt(parseInt))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"freezeAmount hash": tx.Hash(),
	})
}

func addBlackRange(c *gin.Context) {
	startblock := c.PostForm("startblock")
	endblock := c.PostForm("endblock")

	parseStartPos, err := strconv.ParseInt(startblock, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	parseEndPos, err := strconv.ParseInt(endblock, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	instance, auth := util.PrepareTx()

	br := IAllERC20.IFATERC20ConfigBlockRange{
		BeginBlock: big.NewInt(parseStartPos),
		EndBlock:   big.NewInt(parseEndPos),
	}
	tx, err := instance.AddBlackBlock(auth, br)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"Mint hash": tx.Hash(),
	})
}

func mint(c *gin.Context) {
	account := c.PostForm("account")
	amount := c.PostForm("amonut")

	parseInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	instance, auth := util.PrepareTx()

	tx, err := instance.Mint(auth, common.HexToAddress(account), big.NewInt(parseInt))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"Mint hash": tx.Hash(),
	})
}

func burn(c *gin.Context) {
	amount := c.PostForm("amonut")

	parseInt, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	instance, auth := util.PrepareTx()

	tx, err := instance.Burn(auth, big.NewInt(parseInt))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"Burn hash": tx.Hash(),
	})
}
