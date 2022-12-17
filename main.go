package main

import (
	"fmt"
	"github.com/ethereum/coin-manage/db"
	"github.com/gin-gonic/gin"
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
	r.POST("/banAccount", banAccount)
	r.POST("/restriAccountIn", restriAccountIn)
	r.POST("/restriAccountOut", restriAccountOut)
	r.POST("/freezeAmount", freezeAmount)
	r.POST("/setAccountTransfer", setAccountTransfer)
	r.POST("/setTransferTime", setTransferTime)
	r.POST("/increaseCoin", increaseCoin)
	r.POST("/destoryCoin", destoryCoin)

	r.Run(":8000")
}

func getCoinHolders(c *gin.Context) {
	addr := c.Param("contractAddr")
	if addr[:2] != "0x" {
		c.JSON(400, gin.H{
			"error": "start not with 0x",
		})
	}
	if len(addr) != CONTRACTLEN {
		c.JSON(400, gin.H{
			"error": "wrong len",
		})
	}
	dbconn := db.Createdb()
	defer dbconn.Close()

	holderInfos := db.QueryData(dbconn, addr)

	c.JSON(200, gin.H{
		"HolderInfos": holderInfos,
	})
}
func banAccount(c *gin.Context) {
	account := c.PostForm("account")
	fmt.Printf(account)
	c.JSON(200, gin.H{
		"Example": "banAccount",
	})
}
func restriAccountIn(c *gin.Context) {
	c.JSON(200, gin.H{
		"Example": "restriAccountIn",
	})
}
func restriAccountOut(c *gin.Context) {
	c.JSON(200, gin.H{
		"Example": "restriAccountOut",
	})
}
func freezeAmount(c *gin.Context) {
	c.JSON(200, gin.H{
		"Example": "freezeAmount",
	})
}
func setAccountTransfer(c *gin.Context) {
	c.JSON(200, gin.H{
		"Example": "setAccountTransfer",
	})
}

func setTransferTime(c *gin.Context) {
	c.JSON(200, gin.H{
		"Example": "setTransferTime",
	})
}

func increaseCoin(c *gin.Context) {
	c.JSON(200, gin.H{
		"Example": "increaseCoin",
	})
}

func destoryCoin(c *gin.Context) {
	c.JSON(200, gin.H{
		"Example": "destoryCoin",
	})
}
