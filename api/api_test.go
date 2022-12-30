package api

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

var expected = "361737640000000000000000000000006f4a1faf7b9ef87a5213389be3fee3d160f26a8600000000000000000000000000000000000000000000000000000000000003e8"

/*TODO(keep), 利用abi.Pack方式packed，需要注意的是在abigen 生成的uintXX 都是用*big.Int 来表示的"(_IAllERC20 *IAllERC20Transactor) Frozen(opts *bind.TransactOpts, account common.Address, amount *big.Int)
所以我这里也用big.Int
*/
func TestPackforzenData(t *testing.T) {
	amount := big.NewInt(1000)
	encoded, err := forzenData1("frozen", common.HexToAddress("0x6f4A1faF7B9EF87A5213389Be3Fee3D160f26a86"), amount)
	require.NoError(t, err)
	require.Equal(t, expected, hex.EncodeToString(encoded))
}

//TODO(keep),手工packed, 需要人工计算每个输入参数所占的byte的长度
func TestPackforzenData2(t *testing.T) {
	buf := make([]byte, 0)
	selector := common.FromHex("0x36173764")

	buf = append(buf, selector...)

	buf1 := make([]byte, 32)
	big.NewInt(0).SetBytes(common.FromHex("0x6f4A1faF7B9EF87A5213389Be3Fee3D160f26a86")).FillBytes(buf1)
	buf = append(buf, buf1...)

	buf2 := make([]byte, 32)
	big.NewInt(1000).FillBytes(buf2)
	buf = append(buf, buf2...)

	require.Equal(t, expected, hex.EncodeToString(buf))
}
