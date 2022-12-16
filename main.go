package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/getCoinHolders/:coin", getCoinHolders)
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
	coin := c.Param("coin")
	//限制是否3位字母
	fmt.Printf(coin)
	c.JSON(200, gin.H{
		"Example": "getCoinHolders",
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
