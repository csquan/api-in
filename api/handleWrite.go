package api

import (
	"encoding/hex"
	"encoding/json"
	IAllERC20 "github.com/ethereum/coin-manage/contract"
	"github.com/ethereum/coin-manage/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func addBlackData(method string, accountAddr common.Address) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		logrus.Error(err)
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		logrus.Error(err)
	}
	return contractAbi.Pack(method, accountAddr)
}

func forzenData(method string, accountAddr common.Address, amount *big.Int) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		logrus.Error(err)
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		logrus.Error(err)
	}
	return contractAbi.Pack(method, accountAddr, amount)
}

func removeblackRangeData(method string, index *big.Int) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		logrus.Error(err)
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		logrus.Error(err)
	}
	return contractAbi.Pack(method, index)
}

func addblackRangeData(method string, blockRange IAllERC20.IFATERC20ConfigBlockRange) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		logrus.Error(err)
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		logrus.Error(err)
	}
	return contractAbi.Pack(method, blockRange)
}

func mintData(method string, receiverAddr common.Address, amount *big.Int) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		logrus.Error(err)
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		logrus.Error(err)
	}
	return contractAbi.Pack(method, receiverAddr, amount)
}

func burnData(method string, amount *big.Int) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		logrus.Error(err)
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		logrus.Error(err)
	}
	return contractAbi.Pack(method, amount)
}

func burnFromData(method string, targetAddr common.Address, amount *big.Int) ([]byte, error) {
	data, err := ioutil.ReadFile("./contract/IAllERC20.abi")
	if err != nil {
		logrus.Error(err)
	}

	abiStr := string(data)

	r := strings.NewReader(abiStr)
	contractAbi, err := abi.JSON(r)
	if err != nil {
		logrus.Error(err)
	}
	return contractAbi.Pack(method, targetAddr, amount)
}

func (a *ApiService) addBlack(c *gin.Context) {
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

	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	err := checkAddr(contractAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(targetAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := addBlackData("addBlack", common.HexToAddress(targetAddr.String()))

	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}
	cli := resty.New()

	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	data := types.TxData{
		RequestID: strconv.Itoa(int(time.Now().Unix())),
		UID:       uid.String(),
		UUID:      strconv.Itoa(int(time.Now().Unix())),
		From:      operatorAddr.String(),
		To:        contractAddr.String(),
		Data:      "0x" + hex.EncodeToString(inputData),
		Value:     "0x0",
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) removeBlack(c *gin.Context) {
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

	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	err := checkAddr(contractAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(targetAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := addBlackData("removeBlack", common.HexToAddress(targetAddr.String()))

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) removeBlackIn(c *gin.Context) {
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

	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	err := checkAddr(targetAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := addBlackData("removeBlackIn", common.HexToAddress(targetAddr.String()))

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) addBlackIn(c *gin.Context) {
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

	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	err := checkAddr(targetAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := addBlackData("addBlackIn", common.HexToAddress(targetAddr.String()))

	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}
func (a *ApiService) addBlackOut(c *gin.Context) {
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

	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	err := checkAddr(targetAddr.String())
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := addBlackData("addBlackOut", common.HexToAddress(targetAddr.String()))

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}
func (a *ApiService) removeBlackOut(c *gin.Context) {
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

	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	err := checkAddr(targetAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := addBlackData("removeBlackOut", common.HexToAddress(targetAddr.String()))

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) unfrozen(c *gin.Context) {
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
	amount := gjson.Get(data1, "amount")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	parseInt, err := strconv.ParseInt(amount.String(), 10, 64)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil || parseInt <= 0 {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(targetAddr.String())
	if err != nil || parseInt <= 0 {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	Amount := &big.Int{}
	Amount.SetInt64(parseInt)

	inputData, err := forzenData("unfrozen", common.HexToAddress(targetAddr.String()), Amount)

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(err)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) frozen(c *gin.Context) {
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
	amount := gjson.Get(data1, "amount")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	parseInt, err := strconv.ParseInt(amount.String(), 10, 64)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil || parseInt <= 0 {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(targetAddr.String())
	if err != nil || parseInt <= 0 {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	Amount := &big.Int{}
	Amount.SetInt64(parseInt)

	inputData, err := forzenData("frozen", common.HexToAddress(targetAddr.String()), Amount)

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(err)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) addBlackRange(c *gin.Context) {
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
	startblock := gjson.Get(data1, "startblock")
	endblock := gjson.Get(data1, "endblock")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	uid := gjson.Get(data1, "uid")

	parseStartPos, err := strconv.ParseInt(startblock.String(), 10, 64)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	parseEndPos, err := strconv.ParseInt(endblock.String(), 10, 64)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	br := IAllERC20.IFATERC20ConfigBlockRange{
		BeginBlock: big.NewInt(parseStartPos),
		EndBlock:   big.NewInt(parseEndPos),
	}
	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	inputData, err := addblackRangeData("addBlackBlock", br)

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(result)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(result)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) removeBlackRange(c *gin.Context) {
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
	index := gjson.Get(data1, "index")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	uid := gjson.Get(data1, "uid")

	indexPos, err := strconv.ParseInt(index.String(), 10, 64)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	Amount := &big.Int{}
	Amount.SetInt64(indexPos)

	inputData, err := removeblackRangeData("removeBlackBlock", Amount)

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) mint(c *gin.Context) {
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
	amount := gjson.Get(data1, "amount")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	uid := gjson.Get(data1, "uid")

	parseInt, err := strconv.ParseInt(amount.String(), 10, 64)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	err = checkAddr(operatorAddr.String())
	if err != nil || parseInt <= 0 {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	db := *a.conns["hui"]

	coinInfo, err := db.QuerySpecifyCoinInfo(strings.ToLower(contractAddr.String()))
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		c.SecureJSON(http.StatusInternalServerError, res)
		return
	}

	decimalInt, err := strconv.ParseInt(coinInfo.Decimals, 10, 64)
	if err != nil {
		logrus.Error(err)
	}

	Amount := &big.Int{}
	Amount.SetInt64(parseInt)
	big10 := &big.Int{}
	big10.SetInt64(10)

	for i := 0; i < int(decimalInt); i++ {
		Amount = Amount.Mul(Amount, big10)
	}
	logrus.Info(Amount.String())
	inputData, err := mintData("mint", common.HexToAddress(operatorAddr.String()), Amount)

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) burnFrom(c *gin.Context) {
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
	amount := gjson.Get(data1, "amount")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	targetAddr := gjson.Get(data1, "targetAddr")
	uid := gjson.Get(data1, "uid")

	parseInt, err := strconv.ParseInt(amount.String(), 10, 64)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	if parseInt <= 0 {
		logrus.Error(parseInt)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	err = checkAddr(operatorAddr.String())
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	err = checkAddr(contractAddr.String())
	if err != nil || parseInt <= 0 {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	err = checkAddr(targetAddr.String())
	if err != nil || parseInt <= 0 {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	Amount := &big.Int{}
	Amount.SetInt64(parseInt)

	inputData, err := burnFromData("burnFrom", common.HexToAddress(targetAddr.String()), Amount)

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}

func (a *ApiService) burn(c *gin.Context) {
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
	amount := gjson.Get(data1, "amount")
	contractAddr := gjson.Get(data1, "contractAddr")
	operatorAddr := gjson.Get(data1, "operatorAddr")
	uid := gjson.Get(data1, "uid")

	parseInt, err := strconv.ParseInt(amount.String(), 10, 64)
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	if parseInt <= 0 {
		logrus.Error(parseInt)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}
	chainInfo, err := getChainInfo(a.config.ChainInfos, "hui")
	if err != nil {
		logrus.Error(err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()
		c.SecureJSON(http.StatusBadRequest, res)
		return
	}

	Amount := &big.Int{}
	Amount.SetInt64(parseInt)

	inputData, err := burnData("burn", Amount)

	if err != nil {
		logrus.Error(err)
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
		ChainId:   chainInfo.ChainId,
	}

	var result types.HttpRes
	resp, err := cli.R().SetBody(data).SetResult(&result).Post(a.config.TxState.EndPoint + "/tx/create")
	if err != nil {
		logrus.Error(err)
	}
	if resp.StatusCode() != http.StatusOK {
		logrus.Error(resp)
	}
	if result.Code != 0 {
		logrus.Error(result)
	}

	if err != nil {
		logrus.Error(err)
	}

	d, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Data = string(d)

	c.SecureJSON(http.StatusOK, res)
}
