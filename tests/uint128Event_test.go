package tests

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseUint128Event(t *testing.T) {
	l2url := "https://prealpha.scroll.io/l2"
	//uint128Addr := common.HexToAddress("0xec3aacF9cD28A304C547E8b7b995fd0A5fd7F2a0")
	hash := common.HexToHash("0x14392c497dcc67cd372e1f6c6e910c732bd51b87cedf38c0c0b3185c5849a08a")

	l2client, err := ethclient.Dial(l2url)
	require.NoError(t, err)

	tx, _, err := l2client.TransactionByHash(context.Background(), hash)
	require.NoError(t, err)
	fmt.Printf("tx:%v\n", tx)

	receipts, err := l2client.TransactionReceipt(context.Background(), hash)
	require.NoError(t, err)
	fmt.Printf("receipts:%+v\n", receipts)
	fmt.Printf("log0:%v\n", receipts.Logs[0])
	fmt.Printf("data:%v\n", hex.EncodeToString(receipts.Logs[0].Data))

}

0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 200
0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 44