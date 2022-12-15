package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	util "github.com/ethereum/coin-manage/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//util "./contract" // for demo
)

func main() {
	// 一个通知退出的chan
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x148933B2A02F51bA249b323FEdAd9409124Ab4c4")
	instance, err := util.NewUtil(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	ret, err := instance.Zzz(nil)
	fmt.Println(ret)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/getCoinHolders", getCoinHolders)
	mux.HandleFunc("/banAccount", banAccount)
	mux.HandleFunc("/restriAccountIn", restriAccountIn)
	mux.HandleFunc("/restriAccountOut", restriAccountOut)
	mux.HandleFunc("/freezeAmount", freezeAmount)
	mux.HandleFunc("/setAccountTransfer", setAccountTransfer)
	mux.HandleFunc("/setTransferTime", setTransferTime)
	mux.HandleFunc("/increaseCoin", increaseCoin)
	mux.HandleFunc("/destoryCoin", destoryCoin)

	server := &http.Server{
		Addr:         ":80",
		WriteTimeout: time.Second * 3, //设置3秒的写超时
		Handler:      mux,
	}

	go func() {
		// 接收退出信号
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:", err)
		}
	}()

	log.Println("Starting coin-manage httpserver")
	err = server.ListenAndServe()
	if err != nil {
		// 正常退出
		if err == http.ErrServerClosed {
			log.Fatal("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", err)
		}
	}
	log.Fatal("Server exited")
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is coin-manage server"))
}
func getCoinHolders(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	coin := query.Get("coin")

	fmt.Printf(coin)
	w.Write([]byte("--getCoinHolders--"))
}
func banAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	account := r.FormValue("account")
	fmt.Printf(account)
	w.Write([]byte("--banAccount--"))
}
func restriAccountIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("-iAccountIn--"))
}
func restriAccountOut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("--restriAccountOut--"))
}
func freezeAmount(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("--freezeAmount--"))
}

func setAccountTransfer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("--setAccountTransfer--"))
}
func setTransferTime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("--setTransferTime--"))
}
func increaseCoin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("--increaseCoin--"))
}
func destoryCoin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("--destoryCoin--"))
}
