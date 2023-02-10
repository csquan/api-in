package api

import (
	"fmt"
	"github.com/ethereum/coin-manage/config"
	"github.com/ethereum/coin-manage/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

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
	r.GET("/getTxHistory/:accountAddr/:contractAddr/:beginTime/:endTime", a.getTxHistory)
	r.GET("/hasBurnAmount/:accountAddr/:contractAddr", a.hasBurnAmount)

	r.GET("/getBlockHeight", a.getBlockHeight)

	r.GET("/getCoinHistory/:contractAddr", a.getCoinHistory)

	r.POST("/addMultiSign", a.addmultisign)
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
	r.POST("/flashFee", a.getFlashFee)

	r.POST("/model", a.model)
	r.POST("/tx/get", a.GetTask)

	logrus.Info("coin-manage run at " + a.config.Server.Port)

	err := r.Run(fmt.Sprintf(":%s", a.config.Server.Port))
	if err != nil {
		logrus.Fatalf("start http server err:%v", err)
	}
}
