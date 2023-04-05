package api

import (
	"github.com/ethereum/api-in/enc"
	"github.com/ethereum/api-in/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
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
	mechainismName := gjson.Get(data1, "mechainismName")

	mechainism, err := a.db.GetMechanismInfo(mechainismName.String())
	if err != nil {
		logrus.Error(err)
	}

	handleshake := enc.LiveHandshake(mechainism.Key, mechainism.Secret)
	logrus.Info(handleshake)

	// data := ReqAccount{
	//		Handshake: handshake,
	//		AccountId: "b1JuaPlPaImVOSD",
	//	}
	//	jsonValue, _ := json.Marshal(data)
	//	r, err := http.Post("https://api.huiwang.io/api/v1/account/query", "application/json", bytes.NewBuffer(jsonValue))
	//	if err != nil {
	//		panic(err)
	//	}
	//	body, _ := io.ReadAll(r.Body)
	//	fmt.Println("Post result:", string(body))

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
	//amount := gjson.Get(data1, "amount")

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
	//amount := gjson.Get(data1, "amount")

	res.Code = 0
	res.Message = "success"
	res.Data = "null"

	c.SecureJSON(http.StatusOK, res)
}
