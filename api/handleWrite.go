package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/api-in/enc"
	"github.com/ethereum/api-in/types"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

//func (a *ApiService) init(c *gin.Context) {
//	buf := make([]byte, 2048)
//	n, _ := c.Request.Body.Read(buf)
//	data1 := string(buf[0:n])
//	res := types.HttpRes{}
//
//	isValid := gjson.Valid(data1)
//	if isValid == false {
//		logrus.Error("Not valid json")
//		res.Code = http.StatusBadRequest
//		res.Message = "Not valid json"
//		c.SecureJSON(http.StatusBadRequest, res)
//		return
//	}
//	name := gjson.Get(data1, "name")
//	apiKey := gjson.Get(data1, "apiKey")
//	apiSecret := gjson.Get(data1, "apiSecret")
//
//	mechanismData := types.Mechanism{
//		Name:      name.String(),
//		ApiKey:    apiKey.String(),
//		ApiSecret: apiSecret.String(),
//	}
//
//	err := a.db.CommitWithSession(a.db, func(s *xorm.Session) error {
//		if err := a.db.InsertMechanism(s, &mechanismData); err != nil {
//			logrus.Errorf("insert  InsertMechanism task error:%v tasks:[%v]", err, mechanismData)
//			return err
//		}
//		return nil
//	})
//	if err != nil {
//		logrus.Error(err)
//	}
//
//	res.Code = 0
//	res.Message = err.Error()
//	res.Data = ""
//
//	c.SecureJSON(http.StatusOK, res)
//}

func (a *ApiService) transfer(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])
	res := types.HttpRes{}

	isValid := gjson.Valid(data1)
	if isValid == false {
		logrus.Error("Not valid json")
		res.Code = http.StatusBadRequest
		res.Message = "Not valid json"
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	fromAccount := gjson.Get(data1, "fromAccount")
	toAccount := gjson.Get(data1, "toAccount")
	thirdId := gjson.Get(data1, "thirdId")
	token := gjson.Get(data1, "token")
	amount := gjson.Get(data1, "amount")
	callBack := gjson.Get(data1, "callBack")
	ext := gjson.Get(data1, "ext")

	mechainismName := gjson.Get(data1, "mechainismName")

	mechainism, err := a.db.GetMechanismInfo(mechainismName.String())

	if err != nil {
		logrus.Error(err)
	}

	handleShake := enc.LiveHandshake(mechainism.ApiKey, mechainism.ApiSecret)
	logrus.Info(handleShake)

	transferData := types.Transfer{
		FromAccount: fromAccount.String(),
		ToAccount:   toAccount.String(),
		ThirdId:     thirdId.String(),
		Token:       token.String(),
		Amount:      amount.String(),
		CallBack:    callBack.String(),
		Ext:         ext.String(),
	}
	transfers := make([]types.Transfer, 0)

	transfers = append(transfers, transferData)

	data := types.InternalTransfer{
		Transfers:     transfers,
		IsSync:        true,
		IsTransaction: true,
		Handshake:     handleShake,
	}

	jsonValue, _ := json.Marshal(data)
	r, err := http.Post("https://api.huiwang.io/api/v1/assets/internalTransfer", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(r.Body)
	fmt.Println("Post result:", string(body))

	code := gjson.Get(string(body), "code")
	message := gjson.Get(string(body), "message")
	dataRes := gjson.Get(string(body), "data")

	TransferRecord := types.TransferRecord{
		FromAccount: fromAccount.String(),
		ToAccount:   toAccount.String(),
		ThirdId:     thirdId.String(),
		Token:       token.String(),
		Amount:      amount.String(),
		CallBack:    callBack.String(),
		Ext:         ext.String(),
	}
	err = a.db.CommitWithSession(a.db, func(s *xorm.Session) error {
		if err := a.db.InsertTransfer(s, &TransferRecord); err != nil {
			logrus.Errorf("insert transferRecord transaction task error:%v tasks:[%v]", err, transferData)
			return err
		}
		return nil
	})

	res.Code = int(code.Int())
	res.Message = message.String()
	res.Data = dataRes.String()

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) withdraw(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])
	res := types.HttpRes{}

	isValid := gjson.Valid(data1)
	if isValid == false {
		logrus.Error("Not valid json")
		res.Code = http.StatusBadRequest
		res.Message = "Not valid json"
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	thirdId := gjson.Get(data1, "thirdId")
	account := gjson.Get(data1, "accout")
	symbol := gjson.Get(data1, "symbol")
	amount := gjson.Get(data1, "amount")
	chain := gjson.Get(data1, "chain")
	addr := gjson.Get(data1, "addr")
	isSync := gjson.Get(data1, "isSync")

	mechainismName := gjson.Get(data1, "mechainismName")

	mechainism, err := a.db.GetMechanismInfo(mechainismName.String())
	if err != nil {
		logrus.Error(err)
	}

	handleShake := enc.LiveHandshake(mechainism.ApiKey, mechainism.ApiSecret)
	logrus.Info(handleShake)

	withdrawData := types.Withdraw{
		Handshake: handleShake,
		Account:   account.String(),
		ThirdId:   thirdId.String(),
		Symbol:    symbol.String(),
		Amount:    amount.String(),
		Chain:     chain.String(),
		Addr:      addr.String(),
		IsSync:    isSync.String(),
	}

	jsonValue, _ := json.Marshal(withdrawData)
	r, err := http.Post("https://api.huiwang.io/api/v1/assets/withdraw", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(r.Body)
	fmt.Println("Post result:", string(body))

	err = a.db.CommitWithSession(a.db, func(s *xorm.Session) error {
		if err := a.db.InsertWithdraw(s, &withdrawData); err != nil {
			logrus.Errorf("insert withdrawData transaction task error:%v tasks:[%v]", err, withdrawData)
			return err
		}
		return nil
	})

	res.Code = 0
	res.Message = "success"
	res.Data = "null"

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) exchange(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	data1 := string(buf[0:n])
	res := types.HttpRes{}

	isValid := gjson.Valid(data1)
	if isValid == false {
		logrus.Error("Not valid json")
		res.Code = http.StatusBadRequest
		res.Message = "Not valid json"
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	//下面将信息存入db
	res.Code = 0
	res.Message = "success"
	res.Data = "null"

	c.SecureJSON(http.StatusOK, res)
}
