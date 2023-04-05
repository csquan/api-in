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

	handleShake := enc.LiveHandshake(mechainism.Key, mechainism.Secret)
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

	data := types.InternalTransfer{
		Transfer:      transferData,
		IsSync:        true,
		IsTransaction: false,
		Handshake:     handleShake,
	}

	jsonValue, _ := json.Marshal(data)
	r, err := http.Post("https://api.alpha.huiwang.io/api/v1/assets/internalTransfer", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(r.Body)
	fmt.Println("Post result:", string(body))

	err = a.db.CommitWithSession(a.db, func(s *xorm.Session) error {
		if err := a.db.InsertTransfer(s, &transferData); err != nil {
			logrus.Errorf("insert transferRecord transaction task error:%v tasks:[%v]", err, transferData)
			return err
		}
		return nil
	})

	res.Code = 0
	res.Message = "success"
	res.Data = "null"

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
	account := gjson.Get(data1, "account")
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

	handleShake := enc.LiveHandshake(mechainism.Key, mechainism.Secret)
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
	r, err := http.Post("https://api.alpha.huiwang.io/api/v1/assets/withdraw", "application/json", bytes.NewBuffer(jsonValue))
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

func (a *ApiService) exchanges(c *gin.Context) {
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

	handleShake := enc.LiveHandshake(mechainism.Key, mechainism.Secret)
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

	data := types.InternalTransfer{
		Transfer:      transferData,
		IsSync:        true,
		IsTransaction: false,
		Handshake:     handleShake,
	}

	jsonValue, _ := json.Marshal(data)
	r, err := http.Post("https://api.alpha.huiwang.io/api/v1/assets/internalTransfer", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(r.Body)
	fmt.Println("Post result:", string(body))

	transferRecord := types.Transfer{
		FromAccount: fromAccount.String(),
		ToAccount:   toAccount.String(),
		ThirdId:     thirdId.String(),
		Token:       token.String(),
		Amount:      amount.String(),
		CallBack:    callBack.String(),
		Ext:         ext.String(),
	}

	err = a.db.CommitWithSession(a.db, func(s *xorm.Session) error {
		if err := a.db.InsertTransfer(s, &transferRecord); err != nil {
			logrus.Errorf("insert transferRecord transaction task error:%v tasks:[%v]", err, transferRecord)
			return err
		}
		return nil
	})

	//下面将信息存入db
	res.Code = 0
	res.Message = "success"
	res.Data = "null"

	c.SecureJSON(http.StatusOK, res)
}
