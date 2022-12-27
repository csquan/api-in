// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAllERC20

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FATERC20VotesCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type FATERC20VotesCheckpoint struct {
	FromBlock uint32
	Votes     *big.Int
}

// IFATERC20ConfigBlockRange is an auto generated low-level Go binding around an user-defined struct.
type IFATERC20ConfigBlockRange struct {
	BeginBlock *big.Int
	EndBlock   *big.Int
}

// IAllERC20MetaData contains all meta data concerning the IAllERC20 contract.
var IAllERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddBlack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_beginBlock\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_endBlock\",\"type\":\"uint128\"}],\"name\":\"AddBlackBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddBlackIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddBlackOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozen\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"waitFrozen\",\"type\":\"uint256\"}],\"name\":\"Frozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RemoveBlack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_beginBlock\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_endBlock\",\"type\":\"uint128\"}],\"name\":\"RemoveBlackBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RemoveBlackIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RemoveBlackOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozen\",\"type\":\"uint256\"}],\"name\":\"UnFrozen\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addBlack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"beginBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"endBlock\",\"type\":\"uint128\"}],\"internalType\":\"structIFATERC20Config.BlockRange\",\"name\":\"_blockRanges\",\"type\":\"tuple\"}],\"name\":\"addBlackBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addBlackIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addBlackOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blackBlocks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"beginBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"endBlock\",\"type\":\"uint128\"}],\"internalType\":\"structIFATERC20Config.BlockRange[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"blackBlocksOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"beginBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"endBlock\",\"type\":\"uint128\"}],\"internalType\":\"structIFATERC20Config.BlockRange\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"blackInOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"blackOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"blackOutOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_fee\",\"type\":\"uint24\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"pos\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint224\",\"name\":\"votes\",\"type\":\"uint224\"}],\"internalType\":\"structFATERC20Votes.Checkpoint\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSnapshotId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"frozen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"frozenAmt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"waitFrozenAmt\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"frozenOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBonusFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPastTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPastVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTaxFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"model\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeBlack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeBlackBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeBlackIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeBlackOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"shareOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"snapshotId\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unfrozen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"waitFrozenOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAllERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IAllERC20MetaData.ABI instead.
var IAllERC20ABI = IAllERC20MetaData.ABI

// IAllERC20 is an auto generated Go binding around an Ethereum contract.
type IAllERC20 struct {
	IAllERC20Caller     // Read-only binding to the contract
	IAllERC20Transactor // Write-only binding to the contract
	IAllERC20Filterer   // Log filterer for contract events
}

// IAllERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IAllERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IAllERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAllERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAllERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAllERC20Session struct {
	Contract     *IAllERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAllERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAllERC20CallerSession struct {
	Contract *IAllERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IAllERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAllERC20TransactorSession struct {
	Contract     *IAllERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IAllERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IAllERC20Raw struct {
	Contract *IAllERC20 // Generic contract binding to access the raw methods on
}

// IAllERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAllERC20CallerRaw struct {
	Contract *IAllERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IAllERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAllERC20TransactorRaw struct {
	Contract *IAllERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIAllERC20 creates a new instance of IAllERC20, bound to a specific deployed contract.
func NewIAllERC20(address common.Address, backend bind.ContractBackend) (*IAllERC20, error) {
	contract, err := bindIAllERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAllERC20{IAllERC20Caller: IAllERC20Caller{contract: contract}, IAllERC20Transactor: IAllERC20Transactor{contract: contract}, IAllERC20Filterer: IAllERC20Filterer{contract: contract}}, nil
}

// NewIAllERC20Caller creates a new read-only instance of IAllERC20, bound to a specific deployed contract.
func NewIAllERC20Caller(address common.Address, caller bind.ContractCaller) (*IAllERC20Caller, error) {
	contract, err := bindIAllERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAllERC20Caller{contract: contract}, nil
}

// NewIAllERC20Transactor creates a new write-only instance of IAllERC20, bound to a specific deployed contract.
func NewIAllERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IAllERC20Transactor, error) {
	contract, err := bindIAllERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAllERC20Transactor{contract: contract}, nil
}

// NewIAllERC20Filterer creates a new log filterer instance of IAllERC20, bound to a specific deployed contract.
func NewIAllERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IAllERC20Filterer, error) {
	contract, err := bindIAllERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAllERC20Filterer{contract: contract}, nil
}

// bindIAllERC20 binds a generic wrapper to an already deployed contract.
func bindIAllERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAllERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAllERC20 *IAllERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAllERC20.Contract.IAllERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAllERC20 *IAllERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAllERC20.Contract.IAllERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAllERC20 *IAllERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAllERC20.Contract.IAllERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAllERC20 *IAllERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAllERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAllERC20 *IAllERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAllERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAllERC20 *IAllERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAllERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.Allowance(&_IAllERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.Allowance(&_IAllERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.BalanceOf(&_IAllERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.BalanceOf(&_IAllERC20.CallOpts, account)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) BalanceOfAt(opts *bind.CallOpts, account common.Address, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "balanceOfAt", account, snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _IAllERC20.Contract.BalanceOfAt(&_IAllERC20.CallOpts, account, snapshotId)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address account, uint256 snapshotId) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) BalanceOfAt(account common.Address, snapshotId *big.Int) (*big.Int, error) {
	return _IAllERC20.Contract.BalanceOfAt(&_IAllERC20.CallOpts, account, snapshotId)
}

// BlackBlocks is a free data retrieval call binding the contract method 0x7398e4aa.
//
// Solidity: function blackBlocks() view returns((uint128,uint128)[])
func (_IAllERC20 *IAllERC20Caller) BlackBlocks(opts *bind.CallOpts) ([]IFATERC20ConfigBlockRange, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "blackBlocks")

	if err != nil {
		return *new([]IFATERC20ConfigBlockRange), err
	}

	out0 := *abi.ConvertType(out[0], new([]IFATERC20ConfigBlockRange)).(*[]IFATERC20ConfigBlockRange)

	return out0, err

}

// BlackBlocks is a free data retrieval call binding the contract method 0x7398e4aa.
//
// Solidity: function blackBlocks() view returns((uint128,uint128)[])
func (_IAllERC20 *IAllERC20Session) BlackBlocks() ([]IFATERC20ConfigBlockRange, error) {
	return _IAllERC20.Contract.BlackBlocks(&_IAllERC20.CallOpts)
}

// BlackBlocks is a free data retrieval call binding the contract method 0x7398e4aa.
//
// Solidity: function blackBlocks() view returns((uint128,uint128)[])
func (_IAllERC20 *IAllERC20CallerSession) BlackBlocks() ([]IFATERC20ConfigBlockRange, error) {
	return _IAllERC20.Contract.BlackBlocks(&_IAllERC20.CallOpts)
}

// BlackBlocksOf is a free data retrieval call binding the contract method 0x634f5855.
//
// Solidity: function blackBlocksOf(uint256 _index) view returns((uint128,uint128))
func (_IAllERC20 *IAllERC20Caller) BlackBlocksOf(opts *bind.CallOpts, _index *big.Int) (IFATERC20ConfigBlockRange, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "blackBlocksOf", _index)

	if err != nil {
		return *new(IFATERC20ConfigBlockRange), err
	}

	out0 := *abi.ConvertType(out[0], new(IFATERC20ConfigBlockRange)).(*IFATERC20ConfigBlockRange)

	return out0, err

}

// BlackBlocksOf is a free data retrieval call binding the contract method 0x634f5855.
//
// Solidity: function blackBlocksOf(uint256 _index) view returns((uint128,uint128))
func (_IAllERC20 *IAllERC20Session) BlackBlocksOf(_index *big.Int) (IFATERC20ConfigBlockRange, error) {
	return _IAllERC20.Contract.BlackBlocksOf(&_IAllERC20.CallOpts, _index)
}

// BlackBlocksOf is a free data retrieval call binding the contract method 0x634f5855.
//
// Solidity: function blackBlocksOf(uint256 _index) view returns((uint128,uint128))
func (_IAllERC20 *IAllERC20CallerSession) BlackBlocksOf(_index *big.Int) (IFATERC20ConfigBlockRange, error) {
	return _IAllERC20.Contract.BlackBlocksOf(&_IAllERC20.CallOpts, _index)
}

// BlackInOf is a free data retrieval call binding the contract method 0xa9d06a8f.
//
// Solidity: function blackInOf(address account) view returns(bool)
func (_IAllERC20 *IAllERC20Caller) BlackInOf(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "blackInOf", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackInOf is a free data retrieval call binding the contract method 0xa9d06a8f.
//
// Solidity: function blackInOf(address account) view returns(bool)
func (_IAllERC20 *IAllERC20Session) BlackInOf(account common.Address) (bool, error) {
	return _IAllERC20.Contract.BlackInOf(&_IAllERC20.CallOpts, account)
}

// BlackInOf is a free data retrieval call binding the contract method 0xa9d06a8f.
//
// Solidity: function blackInOf(address account) view returns(bool)
func (_IAllERC20 *IAllERC20CallerSession) BlackInOf(account common.Address) (bool, error) {
	return _IAllERC20.Contract.BlackInOf(&_IAllERC20.CallOpts, account)
}

// BlackOf is a free data retrieval call binding the contract method 0x267b9066.
//
// Solidity: function blackOf(address account) view returns(bool)
func (_IAllERC20 *IAllERC20Caller) BlackOf(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "blackOf", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackOf is a free data retrieval call binding the contract method 0x267b9066.
//
// Solidity: function blackOf(address account) view returns(bool)
func (_IAllERC20 *IAllERC20Session) BlackOf(account common.Address) (bool, error) {
	return _IAllERC20.Contract.BlackOf(&_IAllERC20.CallOpts, account)
}

// BlackOf is a free data retrieval call binding the contract method 0x267b9066.
//
// Solidity: function blackOf(address account) view returns(bool)
func (_IAllERC20 *IAllERC20CallerSession) BlackOf(account common.Address) (bool, error) {
	return _IAllERC20.Contract.BlackOf(&_IAllERC20.CallOpts, account)
}

// BlackOutOf is a free data retrieval call binding the contract method 0xafd6e591.
//
// Solidity: function blackOutOf(address account) view returns(bool)
func (_IAllERC20 *IAllERC20Caller) BlackOutOf(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "blackOutOf", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackOutOf is a free data retrieval call binding the contract method 0xafd6e591.
//
// Solidity: function blackOutOf(address account) view returns(bool)
func (_IAllERC20 *IAllERC20Session) BlackOutOf(account common.Address) (bool, error) {
	return _IAllERC20.Contract.BlackOutOf(&_IAllERC20.CallOpts, account)
}

// BlackOutOf is a free data retrieval call binding the contract method 0xafd6e591.
//
// Solidity: function blackOutOf(address account) view returns(bool)
func (_IAllERC20 *IAllERC20CallerSession) BlackOutOf(account common.Address) (bool, error) {
	return _IAllERC20.Contract.BlackOutOf(&_IAllERC20.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_IAllERC20 *IAllERC20Session) Cap() (*big.Int, error) {
	return _IAllERC20.Contract.Cap(&_IAllERC20.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) Cap() (*big.Int, error) {
	return _IAllERC20.Contract.Cap(&_IAllERC20.CallOpts)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_IAllERC20 *IAllERC20Caller) Checkpoints(opts *bind.CallOpts, account common.Address, pos uint32) (FATERC20VotesCheckpoint, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "checkpoints", account, pos)

	if err != nil {
		return *new(FATERC20VotesCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(FATERC20VotesCheckpoint)).(*FATERC20VotesCheckpoint)

	return out0, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_IAllERC20 *IAllERC20Session) Checkpoints(account common.Address, pos uint32) (FATERC20VotesCheckpoint, error) {
	return _IAllERC20.Contract.Checkpoints(&_IAllERC20.CallOpts, account, pos)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint32,uint224))
func (_IAllERC20 *IAllERC20CallerSession) Checkpoints(account common.Address, pos uint32) (FATERC20VotesCheckpoint, error) {
	return _IAllERC20.Contract.Checkpoints(&_IAllERC20.CallOpts, account, pos)
}

// CurrentSnapshotId is a free data retrieval call binding the contract method 0x970875ce.
//
// Solidity: function currentSnapshotId() view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) CurrentSnapshotId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "currentSnapshotId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSnapshotId is a free data retrieval call binding the contract method 0x970875ce.
//
// Solidity: function currentSnapshotId() view returns(uint256)
func (_IAllERC20 *IAllERC20Session) CurrentSnapshotId() (*big.Int, error) {
	return _IAllERC20.Contract.CurrentSnapshotId(&_IAllERC20.CallOpts)
}

// CurrentSnapshotId is a free data retrieval call binding the contract method 0x970875ce.
//
// Solidity: function currentSnapshotId() view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) CurrentSnapshotId() (*big.Int, error) {
	return _IAllERC20.Contract.CurrentSnapshotId(&_IAllERC20.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_IAllERC20 *IAllERC20Caller) Delegates(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "delegates", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_IAllERC20 *IAllERC20Session) Delegates(account common.Address) (common.Address, error) {
	return _IAllERC20.Contract.Delegates(&_IAllERC20.CallOpts, account)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_IAllERC20 *IAllERC20CallerSession) Delegates(account common.Address) (common.Address, error) {
	return _IAllERC20.Contract.Delegates(&_IAllERC20.CallOpts, account)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) FlashFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "flashFee", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IAllERC20.Contract.FlashFee(&_IAllERC20.CallOpts, token, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _IAllERC20.Contract.FlashFee(&_IAllERC20.CallOpts, token, amount)
}

// FrozenOf is a free data retrieval call binding the contract method 0x1bf6e00d.
//
// Solidity: function frozenOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) FrozenOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "frozenOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenOf is a free data retrieval call binding the contract method 0x1bf6e00d.
//
// Solidity: function frozenOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) FrozenOf(account common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.FrozenOf(&_IAllERC20.CallOpts, account)
}

// FrozenOf is a free data retrieval call binding the contract method 0x1bf6e00d.
//
// Solidity: function frozenOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) FrozenOf(account common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.FrozenOf(&_IAllERC20.CallOpts, account)
}

// GetBonusFee is a free data retrieval call binding the contract method 0x37c84060.
//
// Solidity: function getBonusFee() view returns(uint24)
func (_IAllERC20 *IAllERC20Caller) GetBonusFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "getBonusFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBonusFee is a free data retrieval call binding the contract method 0x37c84060.
//
// Solidity: function getBonusFee() view returns(uint24)
func (_IAllERC20 *IAllERC20Session) GetBonusFee() (*big.Int, error) {
	return _IAllERC20.Contract.GetBonusFee(&_IAllERC20.CallOpts)
}

// GetBonusFee is a free data retrieval call binding the contract method 0x37c84060.
//
// Solidity: function getBonusFee() view returns(uint24)
func (_IAllERC20 *IAllERC20CallerSession) GetBonusFee() (*big.Int, error) {
	return _IAllERC20.Contract.GetBonusFee(&_IAllERC20.CallOpts)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) GetPastTotalSupply(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "getPastTotalSupply", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	return _IAllERC20.Contract.GetPastTotalSupply(&_IAllERC20.CallOpts, blockNumber)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 blockNumber) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) GetPastTotalSupply(blockNumber *big.Int) (*big.Int, error) {
	return _IAllERC20.Contract.GetPastTotalSupply(&_IAllERC20.CallOpts, blockNumber)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) GetPastVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "getPastVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _IAllERC20.Contract.GetPastVotes(&_IAllERC20.CallOpts, account, blockNumber)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 blockNumber) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) GetPastVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _IAllERC20.Contract.GetPastVotes(&_IAllERC20.CallOpts, account, blockNumber)
}

// GetTaxFee is a free data retrieval call binding the contract method 0xf66608fe.
//
// Solidity: function getTaxFee() view returns(uint24)
func (_IAllERC20 *IAllERC20Caller) GetTaxFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "getTaxFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTaxFee is a free data retrieval call binding the contract method 0xf66608fe.
//
// Solidity: function getTaxFee() view returns(uint24)
func (_IAllERC20 *IAllERC20Session) GetTaxFee() (*big.Int, error) {
	return _IAllERC20.Contract.GetTaxFee(&_IAllERC20.CallOpts)
}

// GetTaxFee is a free data retrieval call binding the contract method 0xf66608fe.
//
// Solidity: function getTaxFee() view returns(uint24)
func (_IAllERC20 *IAllERC20CallerSession) GetTaxFee() (*big.Int, error) {
	return _IAllERC20.Contract.GetTaxFee(&_IAllERC20.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) GetVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) GetVotes(account common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.GetVotes(&_IAllERC20.CallOpts, account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) GetVotes(account common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.GetVotes(&_IAllERC20.CallOpts, account)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.MaxFlashLoan(&_IAllERC20.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.MaxFlashLoan(&_IAllERC20.CallOpts, token)
}

// Model is a free data retrieval call binding the contract method 0x0ad9d052.
//
// Solidity: function model() view returns(uint8)
func (_IAllERC20 *IAllERC20Caller) Model(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "model")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Model is a free data retrieval call binding the contract method 0x0ad9d052.
//
// Solidity: function model() view returns(uint8)
func (_IAllERC20 *IAllERC20Session) Model() (uint8, error) {
	return _IAllERC20.Contract.Model(&_IAllERC20.CallOpts)
}

// Model is a free data retrieval call binding the contract method 0x0ad9d052.
//
// Solidity: function model() view returns(uint8)
func (_IAllERC20 *IAllERC20CallerSession) Model() (uint8, error) {
	return _IAllERC20.Contract.Model(&_IAllERC20.CallOpts)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_IAllERC20 *IAllERC20Caller) NumCheckpoints(opts *bind.CallOpts, account common.Address) (uint32, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "numCheckpoints", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_IAllERC20 *IAllERC20Session) NumCheckpoints(account common.Address) (uint32, error) {
	return _IAllERC20.Contract.NumCheckpoints(&_IAllERC20.CallOpts, account)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_IAllERC20 *IAllERC20CallerSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _IAllERC20.Contract.NumCheckpoints(&_IAllERC20.CallOpts, account)
}

// ShareOf is a free data retrieval call binding the contract method 0x21e5e2c4.
//
// Solidity: function shareOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) ShareOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "shareOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShareOf is a free data retrieval call binding the contract method 0x21e5e2c4.
//
// Solidity: function shareOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) ShareOf(account common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.ShareOf(&_IAllERC20.CallOpts, account)
}

// ShareOf is a free data retrieval call binding the contract method 0x21e5e2c4.
//
// Solidity: function shareOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) ShareOf(account common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.ShareOf(&_IAllERC20.CallOpts, account)
}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) TotalShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "totalShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_IAllERC20 *IAllERC20Session) TotalShare() (*big.Int, error) {
	return _IAllERC20.Contract.TotalShare(&_IAllERC20.CallOpts)
}

// TotalShare is a free data retrieval call binding the contract method 0x026c4207.
//
// Solidity: function totalShare() view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) TotalShare() (*big.Int, error) {
	return _IAllERC20.Contract.TotalShare(&_IAllERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IAllERC20 *IAllERC20Session) TotalSupply() (*big.Int, error) {
	return _IAllERC20.Contract.TotalSupply(&_IAllERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IAllERC20.Contract.TotalSupply(&_IAllERC20.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) TotalSupplyAt(opts *bind.CallOpts, snapshotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "totalSupplyAt", snapshotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _IAllERC20.Contract.TotalSupplyAt(&_IAllERC20.CallOpts, snapshotId)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 snapshotId) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) TotalSupplyAt(snapshotId *big.Int) (*big.Int, error) {
	return _IAllERC20.Contract.TotalSupplyAt(&_IAllERC20.CallOpts, snapshotId)
}

// WaitFrozenOf is a free data retrieval call binding the contract method 0x071ce00f.
//
// Solidity: function waitFrozenOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20Caller) WaitFrozenOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IAllERC20.contract.Call(opts, &out, "waitFrozenOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WaitFrozenOf is a free data retrieval call binding the contract method 0x071ce00f.
//
// Solidity: function waitFrozenOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20Session) WaitFrozenOf(account common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.WaitFrozenOf(&_IAllERC20.CallOpts, account)
}

// WaitFrozenOf is a free data retrieval call binding the contract method 0x071ce00f.
//
// Solidity: function waitFrozenOf(address account) view returns(uint256)
func (_IAllERC20 *IAllERC20CallerSession) WaitFrozenOf(account common.Address) (*big.Int, error) {
	return _IAllERC20.Contract.WaitFrozenOf(&_IAllERC20.CallOpts, account)
}

// AddBlack is a paid mutator transaction binding the contract method 0xeb794dd7.
//
// Solidity: function addBlack(address account) returns()
func (_IAllERC20 *IAllERC20Transactor) AddBlack(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "addBlack", account)
}

// AddBlack is a paid mutator transaction binding the contract method 0xeb794dd7.
//
// Solidity: function addBlack(address account) returns()
func (_IAllERC20 *IAllERC20Session) AddBlack(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.AddBlack(&_IAllERC20.TransactOpts, account)
}

// AddBlack is a paid mutator transaction binding the contract method 0xeb794dd7.
//
// Solidity: function addBlack(address account) returns()
func (_IAllERC20 *IAllERC20TransactorSession) AddBlack(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.AddBlack(&_IAllERC20.TransactOpts, account)
}

// AddBlackBlock is a paid mutator transaction binding the contract method 0x46f903c6.
//
// Solidity: function addBlackBlock((uint128,uint128) _blockRanges) returns()
func (_IAllERC20 *IAllERC20Transactor) AddBlackBlock(opts *bind.TransactOpts, _blockRanges IFATERC20ConfigBlockRange) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "addBlackBlock", _blockRanges)
}

// AddBlackBlock is a paid mutator transaction binding the contract method 0x46f903c6.
//
// Solidity: function addBlackBlock((uint128,uint128) _blockRanges) returns()
func (_IAllERC20 *IAllERC20Session) AddBlackBlock(_blockRanges IFATERC20ConfigBlockRange) (*types.Transaction, error) {
	return _IAllERC20.Contract.AddBlackBlock(&_IAllERC20.TransactOpts, _blockRanges)
}

// AddBlackBlock is a paid mutator transaction binding the contract method 0x46f903c6.
//
// Solidity: function addBlackBlock((uint128,uint128) _blockRanges) returns()
func (_IAllERC20 *IAllERC20TransactorSession) AddBlackBlock(_blockRanges IFATERC20ConfigBlockRange) (*types.Transaction, error) {
	return _IAllERC20.Contract.AddBlackBlock(&_IAllERC20.TransactOpts, _blockRanges)
}

// AddBlackIn is a paid mutator transaction binding the contract method 0xf6255435.
//
// Solidity: function addBlackIn(address account) returns()
func (_IAllERC20 *IAllERC20Transactor) AddBlackIn(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "addBlackIn", account)
}

// AddBlackIn is a paid mutator transaction binding the contract method 0xf6255435.
//
// Solidity: function addBlackIn(address account) returns()
func (_IAllERC20 *IAllERC20Session) AddBlackIn(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.AddBlackIn(&_IAllERC20.TransactOpts, account)
}

// AddBlackIn is a paid mutator transaction binding the contract method 0xf6255435.
//
// Solidity: function addBlackIn(address account) returns()
func (_IAllERC20 *IAllERC20TransactorSession) AddBlackIn(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.AddBlackIn(&_IAllERC20.TransactOpts, account)
}

// AddBlackOut is a paid mutator transaction binding the contract method 0xd0dba944.
//
// Solidity: function addBlackOut(address account) returns()
func (_IAllERC20 *IAllERC20Transactor) AddBlackOut(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "addBlackOut", account)
}

// AddBlackOut is a paid mutator transaction binding the contract method 0xd0dba944.
//
// Solidity: function addBlackOut(address account) returns()
func (_IAllERC20 *IAllERC20Session) AddBlackOut(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.AddBlackOut(&_IAllERC20.TransactOpts, account)
}

// AddBlackOut is a paid mutator transaction binding the contract method 0xd0dba944.
//
// Solidity: function addBlackOut(address account) returns()
func (_IAllERC20 *IAllERC20TransactorSession) AddBlackOut(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.AddBlackOut(&_IAllERC20.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IAllERC20 *IAllERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IAllERC20 *IAllERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Approve(&_IAllERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IAllERC20 *IAllERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Approve(&_IAllERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IAllERC20 *IAllERC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IAllERC20 *IAllERC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Burn(&_IAllERC20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IAllERC20 *IAllERC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Burn(&_IAllERC20.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IAllERC20 *IAllERC20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IAllERC20 *IAllERC20Session) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.BurnFrom(&_IAllERC20.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IAllERC20 *IAllERC20TransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.BurnFrom(&_IAllERC20.TransactOpts, account, amount)
}

// ChangeFee is a paid mutator transaction binding the contract method 0xbf76b42f.
//
// Solidity: function changeFee(uint24 _fee) returns()
func (_IAllERC20 *IAllERC20Transactor) ChangeFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "changeFee", _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0xbf76b42f.
//
// Solidity: function changeFee(uint24 _fee) returns()
func (_IAllERC20 *IAllERC20Session) ChangeFee(_fee *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.ChangeFee(&_IAllERC20.TransactOpts, _fee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0xbf76b42f.
//
// Solidity: function changeFee(uint24 _fee) returns()
func (_IAllERC20 *IAllERC20TransactorSession) ChangeFee(_fee *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.ChangeFee(&_IAllERC20.TransactOpts, _fee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_IAllERC20 *IAllERC20Transactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_IAllERC20 *IAllERC20Session) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.Delegate(&_IAllERC20.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_IAllERC20 *IAllERC20TransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.Delegate(&_IAllERC20.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_IAllERC20 *IAllERC20Transactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_IAllERC20 *IAllERC20Session) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IAllERC20.Contract.DelegateBySig(&_IAllERC20.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_IAllERC20 *IAllERC20TransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IAllERC20.Contract.DelegateBySig(&_IAllERC20.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IAllERC20 *IAllERC20Transactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "flashLoan", receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IAllERC20 *IAllERC20Session) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IAllERC20.Contract.FlashLoan(&_IAllERC20.TransactOpts, receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_IAllERC20 *IAllERC20TransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IAllERC20.Contract.FlashLoan(&_IAllERC20.TransactOpts, receiver, token, amount, data)
}

// Frozen is a paid mutator transaction binding the contract method 0x36173764.
//
// Solidity: function frozen(address account, uint256 amount) returns(uint256 frozenAmt, uint256 waitFrozenAmt)
func (_IAllERC20 *IAllERC20Transactor) Frozen(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "frozen", account, amount)
}

// Frozen is a paid mutator transaction binding the contract method 0x36173764.
//
// Solidity: function frozen(address account, uint256 amount) returns(uint256 frozenAmt, uint256 waitFrozenAmt)
func (_IAllERC20 *IAllERC20Session) Frozen(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Frozen(&_IAllERC20.TransactOpts, account, amount)
}

// Frozen is a paid mutator transaction binding the contract method 0x36173764.
//
// Solidity: function frozen(address account, uint256 amount) returns(uint256 frozenAmt, uint256 waitFrozenAmt)
func (_IAllERC20 *IAllERC20TransactorSession) Frozen(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Frozen(&_IAllERC20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IAllERC20 *IAllERC20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IAllERC20 *IAllERC20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Mint(&_IAllERC20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_IAllERC20 *IAllERC20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Mint(&_IAllERC20.TransactOpts, account, amount)
}

// RemoveBlack is a paid mutator transaction binding the contract method 0xd283859d.
//
// Solidity: function removeBlack(address account) returns()
func (_IAllERC20 *IAllERC20Transactor) RemoveBlack(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "removeBlack", account)
}

// RemoveBlack is a paid mutator transaction binding the contract method 0xd283859d.
//
// Solidity: function removeBlack(address account) returns()
func (_IAllERC20 *IAllERC20Session) RemoveBlack(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.RemoveBlack(&_IAllERC20.TransactOpts, account)
}

// RemoveBlack is a paid mutator transaction binding the contract method 0xd283859d.
//
// Solidity: function removeBlack(address account) returns()
func (_IAllERC20 *IAllERC20TransactorSession) RemoveBlack(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.RemoveBlack(&_IAllERC20.TransactOpts, account)
}

// RemoveBlackBlock is a paid mutator transaction binding the contract method 0x254cf795.
//
// Solidity: function removeBlackBlock(uint256 _index) returns()
func (_IAllERC20 *IAllERC20Transactor) RemoveBlackBlock(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "removeBlackBlock", _index)
}

// RemoveBlackBlock is a paid mutator transaction binding the contract method 0x254cf795.
//
// Solidity: function removeBlackBlock(uint256 _index) returns()
func (_IAllERC20 *IAllERC20Session) RemoveBlackBlock(_index *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.RemoveBlackBlock(&_IAllERC20.TransactOpts, _index)
}

// RemoveBlackBlock is a paid mutator transaction binding the contract method 0x254cf795.
//
// Solidity: function removeBlackBlock(uint256 _index) returns()
func (_IAllERC20 *IAllERC20TransactorSession) RemoveBlackBlock(_index *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.RemoveBlackBlock(&_IAllERC20.TransactOpts, _index)
}

// RemoveBlackIn is a paid mutator transaction binding the contract method 0x409de891.
//
// Solidity: function removeBlackIn(address account) returns()
func (_IAllERC20 *IAllERC20Transactor) RemoveBlackIn(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "removeBlackIn", account)
}

// RemoveBlackIn is a paid mutator transaction binding the contract method 0x409de891.
//
// Solidity: function removeBlackIn(address account) returns()
func (_IAllERC20 *IAllERC20Session) RemoveBlackIn(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.RemoveBlackIn(&_IAllERC20.TransactOpts, account)
}

// RemoveBlackIn is a paid mutator transaction binding the contract method 0x409de891.
//
// Solidity: function removeBlackIn(address account) returns()
func (_IAllERC20 *IAllERC20TransactorSession) RemoveBlackIn(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.RemoveBlackIn(&_IAllERC20.TransactOpts, account)
}

// RemoveBlackOut is a paid mutator transaction binding the contract method 0x8f9abdf9.
//
// Solidity: function removeBlackOut(address account) returns()
func (_IAllERC20 *IAllERC20Transactor) RemoveBlackOut(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "removeBlackOut", account)
}

// RemoveBlackOut is a paid mutator transaction binding the contract method 0x8f9abdf9.
//
// Solidity: function removeBlackOut(address account) returns()
func (_IAllERC20 *IAllERC20Session) RemoveBlackOut(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.RemoveBlackOut(&_IAllERC20.TransactOpts, account)
}

// RemoveBlackOut is a paid mutator transaction binding the contract method 0x8f9abdf9.
//
// Solidity: function removeBlackOut(address account) returns()
func (_IAllERC20 *IAllERC20TransactorSession) RemoveBlackOut(account common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.RemoveBlackOut(&_IAllERC20.TransactOpts, account)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_IAllERC20 *IAllERC20Transactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_IAllERC20 *IAllERC20Session) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.SetMinter(&_IAllERC20.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_IAllERC20 *IAllERC20TransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.SetMinter(&_IAllERC20.TransactOpts, _minter)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_IAllERC20 *IAllERC20Transactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_IAllERC20 *IAllERC20Session) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.SetOperator(&_IAllERC20.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_IAllERC20 *IAllERC20TransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _IAllERC20.Contract.SetOperator(&_IAllERC20.TransactOpts, _operator)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_IAllERC20 *IAllERC20Transactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_IAllERC20 *IAllERC20Session) Snapshot() (*types.Transaction, error) {
	return _IAllERC20.Contract.Snapshot(&_IAllERC20.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256)
func (_IAllERC20 *IAllERC20TransactorSession) Snapshot() (*types.Transaction, error) {
	return _IAllERC20.Contract.Snapshot(&_IAllERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IAllERC20 *IAllERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IAllERC20 *IAllERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Transfer(&_IAllERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IAllERC20 *IAllERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Transfer(&_IAllERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IAllERC20 *IAllERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IAllERC20 *IAllERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.TransferFrom(&_IAllERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IAllERC20 *IAllERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.TransferFrom(&_IAllERC20.TransactOpts, from, to, amount)
}

// Unfrozen is a paid mutator transaction binding the contract method 0xe5df3dd0.
//
// Solidity: function unfrozen(address account, uint256 amount) returns()
func (_IAllERC20 *IAllERC20Transactor) Unfrozen(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.contract.Transact(opts, "unfrozen", account, amount)
}

// Unfrozen is a paid mutator transaction binding the contract method 0xe5df3dd0.
//
// Solidity: function unfrozen(address account, uint256 amount) returns()
func (_IAllERC20 *IAllERC20Session) Unfrozen(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Unfrozen(&_IAllERC20.TransactOpts, account, amount)
}

// Unfrozen is a paid mutator transaction binding the contract method 0xe5df3dd0.
//
// Solidity: function unfrozen(address account, uint256 amount) returns()
func (_IAllERC20 *IAllERC20TransactorSession) Unfrozen(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IAllERC20.Contract.Unfrozen(&_IAllERC20.TransactOpts, account, amount)
}

// IAllERC20AddBlackIterator is returned from FilterAddBlack and is used to iterate over the raw logs and unpacked data for AddBlack events raised by the IAllERC20 contract.
type IAllERC20AddBlackIterator struct {
	Event *IAllERC20AddBlack // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20AddBlackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20AddBlack)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20AddBlack)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20AddBlackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20AddBlackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20AddBlack represents a AddBlack event raised by the IAllERC20 contract.
type IAllERC20AddBlack struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddBlack is a free log retrieval operation binding the contract event 0x5489f822c4be2c44efd51dbcbd5cf3ac98aa95d657ceb577118f6abbfbc53b96.
//
// Solidity: event AddBlack(address account)
func (_IAllERC20 *IAllERC20Filterer) FilterAddBlack(opts *bind.FilterOpts) (*IAllERC20AddBlackIterator, error) {

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "AddBlack")
	if err != nil {
		return nil, err
	}
	return &IAllERC20AddBlackIterator{contract: _IAllERC20.contract, event: "AddBlack", logs: logs, sub: sub}, nil
}

// WatchAddBlack is a free log subscription operation binding the contract event 0x5489f822c4be2c44efd51dbcbd5cf3ac98aa95d657ceb577118f6abbfbc53b96.
//
// Solidity: event AddBlack(address account)
func (_IAllERC20 *IAllERC20Filterer) WatchAddBlack(opts *bind.WatchOpts, sink chan<- *IAllERC20AddBlack) (event.Subscription, error) {

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "AddBlack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20AddBlack)
				if err := _IAllERC20.contract.UnpackLog(event, "AddBlack", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddBlack is a log parse operation binding the contract event 0x5489f822c4be2c44efd51dbcbd5cf3ac98aa95d657ceb577118f6abbfbc53b96.
//
// Solidity: event AddBlack(address account)
func (_IAllERC20 *IAllERC20Filterer) ParseAddBlack(log types.Log) (*IAllERC20AddBlack, error) {
	event := new(IAllERC20AddBlack)
	if err := _IAllERC20.contract.UnpackLog(event, "AddBlack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20AddBlackBlockIterator is returned from FilterAddBlackBlock and is used to iterate over the raw logs and unpacked data for AddBlackBlock events raised by the IAllERC20 contract.
type IAllERC20AddBlackBlockIterator struct {
	Event *IAllERC20AddBlackBlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20AddBlackBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20AddBlackBlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20AddBlackBlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20AddBlackBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20AddBlackBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20AddBlackBlock represents a AddBlackBlock event raised by the IAllERC20 contract.
type IAllERC20AddBlackBlock struct {
	BeginBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddBlackBlock is a free log retrieval operation binding the contract event 0x0cf1571993e607251d8c84fabf6d554da5bc040010cb9e1e0ff7f0753f2c7ccb.
//
// Solidity: event AddBlackBlock(uint128 _beginBlock, uint128 _endBlock)
func (_IAllERC20 *IAllERC20Filterer) FilterAddBlackBlock(opts *bind.FilterOpts) (*IAllERC20AddBlackBlockIterator, error) {

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "AddBlackBlock")
	if err != nil {
		return nil, err
	}
	return &IAllERC20AddBlackBlockIterator{contract: _IAllERC20.contract, event: "AddBlackBlock", logs: logs, sub: sub}, nil
}

// WatchAddBlackBlock is a free log subscription operation binding the contract event 0x0cf1571993e607251d8c84fabf6d554da5bc040010cb9e1e0ff7f0753f2c7ccb.
//
// Solidity: event AddBlackBlock(uint128 _beginBlock, uint128 _endBlock)
func (_IAllERC20 *IAllERC20Filterer) WatchAddBlackBlock(opts *bind.WatchOpts, sink chan<- *IAllERC20AddBlackBlock) (event.Subscription, error) {

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "AddBlackBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20AddBlackBlock)
				if err := _IAllERC20.contract.UnpackLog(event, "AddBlackBlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddBlackBlock is a log parse operation binding the contract event 0x0cf1571993e607251d8c84fabf6d554da5bc040010cb9e1e0ff7f0753f2c7ccb.
//
// Solidity: event AddBlackBlock(uint128 _beginBlock, uint128 _endBlock)
func (_IAllERC20 *IAllERC20Filterer) ParseAddBlackBlock(log types.Log) (*IAllERC20AddBlackBlock, error) {
	event := new(IAllERC20AddBlackBlock)
	if err := _IAllERC20.contract.UnpackLog(event, "AddBlackBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20AddBlackInIterator is returned from FilterAddBlackIn and is used to iterate over the raw logs and unpacked data for AddBlackIn events raised by the IAllERC20 contract.
type IAllERC20AddBlackInIterator struct {
	Event *IAllERC20AddBlackIn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20AddBlackInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20AddBlackIn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20AddBlackIn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20AddBlackInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20AddBlackInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20AddBlackIn represents a AddBlackIn event raised by the IAllERC20 contract.
type IAllERC20AddBlackIn struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddBlackIn is a free log retrieval operation binding the contract event 0x362809633ae771a1e765e7fb97462fb07d1a7eeb081190efffcea78ca70e32f5.
//
// Solidity: event AddBlackIn(address account)
func (_IAllERC20 *IAllERC20Filterer) FilterAddBlackIn(opts *bind.FilterOpts) (*IAllERC20AddBlackInIterator, error) {

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "AddBlackIn")
	if err != nil {
		return nil, err
	}
	return &IAllERC20AddBlackInIterator{contract: _IAllERC20.contract, event: "AddBlackIn", logs: logs, sub: sub}, nil
}

// WatchAddBlackIn is a free log subscription operation binding the contract event 0x362809633ae771a1e765e7fb97462fb07d1a7eeb081190efffcea78ca70e32f5.
//
// Solidity: event AddBlackIn(address account)
func (_IAllERC20 *IAllERC20Filterer) WatchAddBlackIn(opts *bind.WatchOpts, sink chan<- *IAllERC20AddBlackIn) (event.Subscription, error) {

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "AddBlackIn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20AddBlackIn)
				if err := _IAllERC20.contract.UnpackLog(event, "AddBlackIn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddBlackIn is a log parse operation binding the contract event 0x362809633ae771a1e765e7fb97462fb07d1a7eeb081190efffcea78ca70e32f5.
//
// Solidity: event AddBlackIn(address account)
func (_IAllERC20 *IAllERC20Filterer) ParseAddBlackIn(log types.Log) (*IAllERC20AddBlackIn, error) {
	event := new(IAllERC20AddBlackIn)
	if err := _IAllERC20.contract.UnpackLog(event, "AddBlackIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20AddBlackOutIterator is returned from FilterAddBlackOut and is used to iterate over the raw logs and unpacked data for AddBlackOut events raised by the IAllERC20 contract.
type IAllERC20AddBlackOutIterator struct {
	Event *IAllERC20AddBlackOut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20AddBlackOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20AddBlackOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20AddBlackOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20AddBlackOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20AddBlackOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20AddBlackOut represents a AddBlackOut event raised by the IAllERC20 contract.
type IAllERC20AddBlackOut struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddBlackOut is a free log retrieval operation binding the contract event 0x71caaa40e79637c707c879b4c8f30f163eab89084138078b6f6e9f1f7dd58d9d.
//
// Solidity: event AddBlackOut(address account)
func (_IAllERC20 *IAllERC20Filterer) FilterAddBlackOut(opts *bind.FilterOpts) (*IAllERC20AddBlackOutIterator, error) {

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "AddBlackOut")
	if err != nil {
		return nil, err
	}
	return &IAllERC20AddBlackOutIterator{contract: _IAllERC20.contract, event: "AddBlackOut", logs: logs, sub: sub}, nil
}

// WatchAddBlackOut is a free log subscription operation binding the contract event 0x71caaa40e79637c707c879b4c8f30f163eab89084138078b6f6e9f1f7dd58d9d.
//
// Solidity: event AddBlackOut(address account)
func (_IAllERC20 *IAllERC20Filterer) WatchAddBlackOut(opts *bind.WatchOpts, sink chan<- *IAllERC20AddBlackOut) (event.Subscription, error) {

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "AddBlackOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20AddBlackOut)
				if err := _IAllERC20.contract.UnpackLog(event, "AddBlackOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddBlackOut is a log parse operation binding the contract event 0x71caaa40e79637c707c879b4c8f30f163eab89084138078b6f6e9f1f7dd58d9d.
//
// Solidity: event AddBlackOut(address account)
func (_IAllERC20 *IAllERC20Filterer) ParseAddBlackOut(log types.Log) (*IAllERC20AddBlackOut, error) {
	event := new(IAllERC20AddBlackOut)
	if err := _IAllERC20.contract.UnpackLog(event, "AddBlackOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IAllERC20 contract.
type IAllERC20ApprovalIterator struct {
	Event *IAllERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20Approval represents a Approval event raised by the IAllERC20 contract.
type IAllERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IAllERC20 *IAllERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IAllERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IAllERC20ApprovalIterator{contract: _IAllERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IAllERC20 *IAllERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IAllERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20Approval)
				if err := _IAllERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IAllERC20 *IAllERC20Filterer) ParseApproval(log types.Log) (*IAllERC20Approval, error) {
	event := new(IAllERC20Approval)
	if err := _IAllERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20DelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the IAllERC20 contract.
type IAllERC20DelegateChangedIterator struct {
	Event *IAllERC20DelegateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20DelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20DelegateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20DelegateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20DelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20DelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20DelegateChanged represents a DelegateChanged event raised by the IAllERC20 contract.
type IAllERC20DelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_IAllERC20 *IAllERC20Filterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*IAllERC20DelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &IAllERC20DelegateChangedIterator{contract: _IAllERC20.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_IAllERC20 *IAllERC20Filterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *IAllERC20DelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20DelegateChanged)
				if err := _IAllERC20.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_IAllERC20 *IAllERC20Filterer) ParseDelegateChanged(log types.Log) (*IAllERC20DelegateChanged, error) {
	event := new(IAllERC20DelegateChanged)
	if err := _IAllERC20.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20DelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the IAllERC20 contract.
type IAllERC20DelegateVotesChangedIterator struct {
	Event *IAllERC20DelegateVotesChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20DelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20DelegateVotesChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20DelegateVotesChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20DelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20DelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20DelegateVotesChanged represents a DelegateVotesChanged event raised by the IAllERC20 contract.
type IAllERC20DelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_IAllERC20 *IAllERC20Filterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*IAllERC20DelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &IAllERC20DelegateVotesChangedIterator{contract: _IAllERC20.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_IAllERC20 *IAllERC20Filterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *IAllERC20DelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20DelegateVotesChanged)
				if err := _IAllERC20.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_IAllERC20 *IAllERC20Filterer) ParseDelegateVotesChanged(log types.Log) (*IAllERC20DelegateVotesChanged, error) {
	event := new(IAllERC20DelegateVotesChanged)
	if err := _IAllERC20.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20FrozenIterator is returned from FilterFrozen and is used to iterate over the raw logs and unpacked data for Frozen events raised by the IAllERC20 contract.
type IAllERC20FrozenIterator struct {
	Event *IAllERC20Frozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20FrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20Frozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20Frozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20FrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20FrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20Frozen represents a Frozen event raised by the IAllERC20 contract.
type IAllERC20Frozen struct {
	Account    common.Address
	Frozen     *big.Int
	WaitFrozen *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFrozen is a free log retrieval operation binding the contract event 0x7b0c29e799ed468266b4d070c03070137a95cb869c5eb7c18063fb7b1a7a09c5.
//
// Solidity: event Frozen(address indexed account, uint256 frozen, uint256 waitFrozen)
func (_IAllERC20 *IAllERC20Filterer) FilterFrozen(opts *bind.FilterOpts, account []common.Address) (*IAllERC20FrozenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "Frozen", accountRule)
	if err != nil {
		return nil, err
	}
	return &IAllERC20FrozenIterator{contract: _IAllERC20.contract, event: "Frozen", logs: logs, sub: sub}, nil
}

// WatchFrozen is a free log subscription operation binding the contract event 0x7b0c29e799ed468266b4d070c03070137a95cb869c5eb7c18063fb7b1a7a09c5.
//
// Solidity: event Frozen(address indexed account, uint256 frozen, uint256 waitFrozen)
func (_IAllERC20 *IAllERC20Filterer) WatchFrozen(opts *bind.WatchOpts, sink chan<- *IAllERC20Frozen, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "Frozen", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20Frozen)
				if err := _IAllERC20.contract.UnpackLog(event, "Frozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFrozen is a log parse operation binding the contract event 0x7b0c29e799ed468266b4d070c03070137a95cb869c5eb7c18063fb7b1a7a09c5.
//
// Solidity: event Frozen(address indexed account, uint256 frozen, uint256 waitFrozen)
func (_IAllERC20 *IAllERC20Filterer) ParseFrozen(log types.Log) (*IAllERC20Frozen, error) {
	event := new(IAllERC20Frozen)
	if err := _IAllERC20.contract.UnpackLog(event, "Frozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20RemoveBlackIterator is returned from FilterRemoveBlack and is used to iterate over the raw logs and unpacked data for RemoveBlack events raised by the IAllERC20 contract.
type IAllERC20RemoveBlackIterator struct {
	Event *IAllERC20RemoveBlack // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20RemoveBlackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20RemoveBlack)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20RemoveBlack)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20RemoveBlackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20RemoveBlackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20RemoveBlack represents a RemoveBlack event raised by the IAllERC20 contract.
type IAllERC20RemoveBlack struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveBlack is a free log retrieval operation binding the contract event 0xba1e86f1f4570d1feaa3b978dbe6a7dd03c7004df88f50a9eba6ed258e4d025b.
//
// Solidity: event RemoveBlack(address account)
func (_IAllERC20 *IAllERC20Filterer) FilterRemoveBlack(opts *bind.FilterOpts) (*IAllERC20RemoveBlackIterator, error) {

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "RemoveBlack")
	if err != nil {
		return nil, err
	}
	return &IAllERC20RemoveBlackIterator{contract: _IAllERC20.contract, event: "RemoveBlack", logs: logs, sub: sub}, nil
}

// WatchRemoveBlack is a free log subscription operation binding the contract event 0xba1e86f1f4570d1feaa3b978dbe6a7dd03c7004df88f50a9eba6ed258e4d025b.
//
// Solidity: event RemoveBlack(address account)
func (_IAllERC20 *IAllERC20Filterer) WatchRemoveBlack(opts *bind.WatchOpts, sink chan<- *IAllERC20RemoveBlack) (event.Subscription, error) {

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "RemoveBlack")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20RemoveBlack)
				if err := _IAllERC20.contract.UnpackLog(event, "RemoveBlack", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveBlack is a log parse operation binding the contract event 0xba1e86f1f4570d1feaa3b978dbe6a7dd03c7004df88f50a9eba6ed258e4d025b.
//
// Solidity: event RemoveBlack(address account)
func (_IAllERC20 *IAllERC20Filterer) ParseRemoveBlack(log types.Log) (*IAllERC20RemoveBlack, error) {
	event := new(IAllERC20RemoveBlack)
	if err := _IAllERC20.contract.UnpackLog(event, "RemoveBlack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20RemoveBlackBlockIterator is returned from FilterRemoveBlackBlock and is used to iterate over the raw logs and unpacked data for RemoveBlackBlock events raised by the IAllERC20 contract.
type IAllERC20RemoveBlackBlockIterator struct {
	Event *IAllERC20RemoveBlackBlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20RemoveBlackBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20RemoveBlackBlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20RemoveBlackBlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20RemoveBlackBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20RemoveBlackBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20RemoveBlackBlock represents a RemoveBlackBlock event raised by the IAllERC20 contract.
type IAllERC20RemoveBlackBlock struct {
	I          *big.Int
	BeginBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveBlackBlock is a free log retrieval operation binding the contract event 0x81a93c04f4d489ee80b0929b36576922bd096cfd416205515e1b7f028f2e0ba2.
//
// Solidity: event RemoveBlackBlock(uint256 i, uint128 _beginBlock, uint128 _endBlock)
func (_IAllERC20 *IAllERC20Filterer) FilterRemoveBlackBlock(opts *bind.FilterOpts) (*IAllERC20RemoveBlackBlockIterator, error) {

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "RemoveBlackBlock")
	if err != nil {
		return nil, err
	}
	return &IAllERC20RemoveBlackBlockIterator{contract: _IAllERC20.contract, event: "RemoveBlackBlock", logs: logs, sub: sub}, nil
}

// WatchRemoveBlackBlock is a free log subscription operation binding the contract event 0x81a93c04f4d489ee80b0929b36576922bd096cfd416205515e1b7f028f2e0ba2.
//
// Solidity: event RemoveBlackBlock(uint256 i, uint128 _beginBlock, uint128 _endBlock)
func (_IAllERC20 *IAllERC20Filterer) WatchRemoveBlackBlock(opts *bind.WatchOpts, sink chan<- *IAllERC20RemoveBlackBlock) (event.Subscription, error) {

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "RemoveBlackBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20RemoveBlackBlock)
				if err := _IAllERC20.contract.UnpackLog(event, "RemoveBlackBlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveBlackBlock is a log parse operation binding the contract event 0x81a93c04f4d489ee80b0929b36576922bd096cfd416205515e1b7f028f2e0ba2.
//
// Solidity: event RemoveBlackBlock(uint256 i, uint128 _beginBlock, uint128 _endBlock)
func (_IAllERC20 *IAllERC20Filterer) ParseRemoveBlackBlock(log types.Log) (*IAllERC20RemoveBlackBlock, error) {
	event := new(IAllERC20RemoveBlackBlock)
	if err := _IAllERC20.contract.UnpackLog(event, "RemoveBlackBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20RemoveBlackInIterator is returned from FilterRemoveBlackIn and is used to iterate over the raw logs and unpacked data for RemoveBlackIn events raised by the IAllERC20 contract.
type IAllERC20RemoveBlackInIterator struct {
	Event *IAllERC20RemoveBlackIn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20RemoveBlackInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20RemoveBlackIn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20RemoveBlackIn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20RemoveBlackInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20RemoveBlackInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20RemoveBlackIn represents a RemoveBlackIn event raised by the IAllERC20 contract.
type IAllERC20RemoveBlackIn struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveBlackIn is a free log retrieval operation binding the contract event 0x9ef10adabfdfa212070d43541902c7fd6f7ceba6cedb9dc5ccd53f6a3dd09c36.
//
// Solidity: event RemoveBlackIn(address account)
func (_IAllERC20 *IAllERC20Filterer) FilterRemoveBlackIn(opts *bind.FilterOpts) (*IAllERC20RemoveBlackInIterator, error) {

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "RemoveBlackIn")
	if err != nil {
		return nil, err
	}
	return &IAllERC20RemoveBlackInIterator{contract: _IAllERC20.contract, event: "RemoveBlackIn", logs: logs, sub: sub}, nil
}

// WatchRemoveBlackIn is a free log subscription operation binding the contract event 0x9ef10adabfdfa212070d43541902c7fd6f7ceba6cedb9dc5ccd53f6a3dd09c36.
//
// Solidity: event RemoveBlackIn(address account)
func (_IAllERC20 *IAllERC20Filterer) WatchRemoveBlackIn(opts *bind.WatchOpts, sink chan<- *IAllERC20RemoveBlackIn) (event.Subscription, error) {

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "RemoveBlackIn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20RemoveBlackIn)
				if err := _IAllERC20.contract.UnpackLog(event, "RemoveBlackIn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveBlackIn is a log parse operation binding the contract event 0x9ef10adabfdfa212070d43541902c7fd6f7ceba6cedb9dc5ccd53f6a3dd09c36.
//
// Solidity: event RemoveBlackIn(address account)
func (_IAllERC20 *IAllERC20Filterer) ParseRemoveBlackIn(log types.Log) (*IAllERC20RemoveBlackIn, error) {
	event := new(IAllERC20RemoveBlackIn)
	if err := _IAllERC20.contract.UnpackLog(event, "RemoveBlackIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20RemoveBlackOutIterator is returned from FilterRemoveBlackOut and is used to iterate over the raw logs and unpacked data for RemoveBlackOut events raised by the IAllERC20 contract.
type IAllERC20RemoveBlackOutIterator struct {
	Event *IAllERC20RemoveBlackOut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20RemoveBlackOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20RemoveBlackOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20RemoveBlackOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20RemoveBlackOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20RemoveBlackOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20RemoveBlackOut represents a RemoveBlackOut event raised by the IAllERC20 contract.
type IAllERC20RemoveBlackOut struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveBlackOut is a free log retrieval operation binding the contract event 0x1354eb903ea104cbd751eec45307639ce124585b0e8172dd33c4d0ecef3a5fe2.
//
// Solidity: event RemoveBlackOut(address account)
func (_IAllERC20 *IAllERC20Filterer) FilterRemoveBlackOut(opts *bind.FilterOpts) (*IAllERC20RemoveBlackOutIterator, error) {

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "RemoveBlackOut")
	if err != nil {
		return nil, err
	}
	return &IAllERC20RemoveBlackOutIterator{contract: _IAllERC20.contract, event: "RemoveBlackOut", logs: logs, sub: sub}, nil
}

// WatchRemoveBlackOut is a free log subscription operation binding the contract event 0x1354eb903ea104cbd751eec45307639ce124585b0e8172dd33c4d0ecef3a5fe2.
//
// Solidity: event RemoveBlackOut(address account)
func (_IAllERC20 *IAllERC20Filterer) WatchRemoveBlackOut(opts *bind.WatchOpts, sink chan<- *IAllERC20RemoveBlackOut) (event.Subscription, error) {

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "RemoveBlackOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20RemoveBlackOut)
				if err := _IAllERC20.contract.UnpackLog(event, "RemoveBlackOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveBlackOut is a log parse operation binding the contract event 0x1354eb903ea104cbd751eec45307639ce124585b0e8172dd33c4d0ecef3a5fe2.
//
// Solidity: event RemoveBlackOut(address account)
func (_IAllERC20 *IAllERC20Filterer) ParseRemoveBlackOut(log types.Log) (*IAllERC20RemoveBlackOut, error) {
	event := new(IAllERC20RemoveBlackOut)
	if err := _IAllERC20.contract.UnpackLog(event, "RemoveBlackOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IAllERC20 contract.
type IAllERC20TransferIterator struct {
	Event *IAllERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20Transfer represents a Transfer event raised by the IAllERC20 contract.
type IAllERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IAllERC20 *IAllERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IAllERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IAllERC20TransferIterator{contract: _IAllERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IAllERC20 *IAllERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IAllERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20Transfer)
				if err := _IAllERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IAllERC20 *IAllERC20Filterer) ParseTransfer(log types.Log) (*IAllERC20Transfer, error) {
	event := new(IAllERC20Transfer)
	if err := _IAllERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAllERC20UnFrozenIterator is returned from FilterUnFrozen and is used to iterate over the raw logs and unpacked data for UnFrozen events raised by the IAllERC20 contract.
type IAllERC20UnFrozenIterator struct {
	Event *IAllERC20UnFrozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAllERC20UnFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAllERC20UnFrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAllERC20UnFrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAllERC20UnFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAllERC20UnFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAllERC20UnFrozen represents a UnFrozen event raised by the IAllERC20 contract.
type IAllERC20UnFrozen struct {
	Account common.Address
	Amount  *big.Int
	Frozen  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnFrozen is a free log retrieval operation binding the contract event 0xcfaa1553132d0157579084fe1ca4a05fa35561f3b05281ed6daf9bda1ad9c32b.
//
// Solidity: event UnFrozen(address indexed account, uint256 amount, uint256 frozen)
func (_IAllERC20 *IAllERC20Filterer) FilterUnFrozen(opts *bind.FilterOpts, account []common.Address) (*IAllERC20UnFrozenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAllERC20.contract.FilterLogs(opts, "UnFrozen", accountRule)
	if err != nil {
		return nil, err
	}
	return &IAllERC20UnFrozenIterator{contract: _IAllERC20.contract, event: "UnFrozen", logs: logs, sub: sub}, nil
}

// WatchUnFrozen is a free log subscription operation binding the contract event 0xcfaa1553132d0157579084fe1ca4a05fa35561f3b05281ed6daf9bda1ad9c32b.
//
// Solidity: event UnFrozen(address indexed account, uint256 amount, uint256 frozen)
func (_IAllERC20 *IAllERC20Filterer) WatchUnFrozen(opts *bind.WatchOpts, sink chan<- *IAllERC20UnFrozen, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IAllERC20.contract.WatchLogs(opts, "UnFrozen", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAllERC20UnFrozen)
				if err := _IAllERC20.contract.UnpackLog(event, "UnFrozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnFrozen is a log parse operation binding the contract event 0xcfaa1553132d0157579084fe1ca4a05fa35561f3b05281ed6daf9bda1ad9c32b.
//
// Solidity: event UnFrozen(address indexed account, uint256 amount, uint256 frozen)
func (_IAllERC20 *IAllERC20Filterer) ParseUnFrozen(log types.Log) (*IAllERC20UnFrozen, error) {
	event := new(IAllERC20UnFrozen)
	if err := _IAllERC20.contract.UnpackLog(event, "UnFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
