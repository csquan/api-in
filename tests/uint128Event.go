// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tests

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

// Uint128EventMetaData contains all meta data concerning the Uint128Event contract.
var Uint128EventMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_beginBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_endBlock\",\"type\":\"uint128\"}],\"name\":\"addBlack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_beginBlock\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_endBlock\",\"type\":\"uint128\"}],\"name\":\"AddBlackBlock\",\"type\":\"event\"}]",
}

// Uint128EventABI is the input ABI used to generate the binding from.
// Deprecated: Use Uint128EventMetaData.ABI instead.
var Uint128EventABI = Uint128EventMetaData.ABI

// Uint128Event is an auto generated Go binding around an Ethereum contract.
type Uint128Event struct {
	Uint128EventCaller     // Read-only binding to the contract
	Uint128EventTransactor // Write-only binding to the contract
	Uint128EventFilterer   // Log filterer for contract events
}

// Uint128EventCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uint128EventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uint128EventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uint128EventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uint128EventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uint128EventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uint128EventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uint128EventSession struct {
	Contract     *Uint128Event     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uint128EventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uint128EventCallerSession struct {
	Contract *Uint128EventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Uint128EventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uint128EventTransactorSession struct {
	Contract     *Uint128EventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Uint128EventRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uint128EventRaw struct {
	Contract *Uint128Event // Generic contract binding to access the raw methods on
}

// Uint128EventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uint128EventCallerRaw struct {
	Contract *Uint128EventCaller // Generic read-only contract binding to access the raw methods on
}

// Uint128EventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uint128EventTransactorRaw struct {
	Contract *Uint128EventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUint128Event creates a new instance of Uint128Event, bound to a specific deployed contract.
func NewUint128Event(address common.Address, backend bind.ContractBackend) (*Uint128Event, error) {
	contract, err := bindUint128Event(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uint128Event{Uint128EventCaller: Uint128EventCaller{contract: contract}, Uint128EventTransactor: Uint128EventTransactor{contract: contract}, Uint128EventFilterer: Uint128EventFilterer{contract: contract}}, nil
}

// NewUint128EventCaller creates a new read-only instance of Uint128Event, bound to a specific deployed contract.
func NewUint128EventCaller(address common.Address, caller bind.ContractCaller) (*Uint128EventCaller, error) {
	contract, err := bindUint128Event(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uint128EventCaller{contract: contract}, nil
}

// NewUint128EventTransactor creates a new write-only instance of Uint128Event, bound to a specific deployed contract.
func NewUint128EventTransactor(address common.Address, transactor bind.ContractTransactor) (*Uint128EventTransactor, error) {
	contract, err := bindUint128Event(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uint128EventTransactor{contract: contract}, nil
}

// NewUint128EventFilterer creates a new log filterer instance of Uint128Event, bound to a specific deployed contract.
func NewUint128EventFilterer(address common.Address, filterer bind.ContractFilterer) (*Uint128EventFilterer, error) {
	contract, err := bindUint128Event(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uint128EventFilterer{contract: contract}, nil
}

// bindUint128Event binds a generic wrapper to an already deployed contract.
func bindUint128Event(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Uint128EventABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uint128Event *Uint128EventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uint128Event.Contract.Uint128EventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uint128Event *Uint128EventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uint128Event.Contract.Uint128EventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uint128Event *Uint128EventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uint128Event.Contract.Uint128EventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uint128Event *Uint128EventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uint128Event.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uint128Event *Uint128EventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uint128Event.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uint128Event *Uint128EventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uint128Event.Contract.contract.Transact(opts, method, params...)
}

// AddBlack is a paid mutator transaction binding the contract method 0xe77795c4.
//
// Solidity: function addBlack(uint256 _i, uint128 _beginBlock, uint128 _endBlock) returns()
func (_Uint128Event *Uint128EventTransactor) AddBlack(opts *bind.TransactOpts, _i *big.Int, _beginBlock *big.Int, _endBlock *big.Int) (*types.Transaction, error) {
	return _Uint128Event.contract.Transact(opts, "addBlack", _i, _beginBlock, _endBlock)
}

// AddBlack is a paid mutator transaction binding the contract method 0xe77795c4.
//
// Solidity: function addBlack(uint256 _i, uint128 _beginBlock, uint128 _endBlock) returns()
func (_Uint128Event *Uint128EventSession) AddBlack(_i *big.Int, _beginBlock *big.Int, _endBlock *big.Int) (*types.Transaction, error) {
	return _Uint128Event.Contract.AddBlack(&_Uint128Event.TransactOpts, _i, _beginBlock, _endBlock)
}

// AddBlack is a paid mutator transaction binding the contract method 0xe77795c4.
//
// Solidity: function addBlack(uint256 _i, uint128 _beginBlock, uint128 _endBlock) returns()
func (_Uint128Event *Uint128EventTransactorSession) AddBlack(_i *big.Int, _beginBlock *big.Int, _endBlock *big.Int) (*types.Transaction, error) {
	return _Uint128Event.Contract.AddBlack(&_Uint128Event.TransactOpts, _i, _beginBlock, _endBlock)
}

// Uint128EventAddBlackBlockIterator is returned from FilterAddBlackBlock and is used to iterate over the raw logs and unpacked data for AddBlackBlock events raised by the Uint128Event contract.
type Uint128EventAddBlackBlockIterator struct {
	Event *Uint128EventAddBlackBlock // Event containing the contract specifics and raw log

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
func (it *Uint128EventAddBlackBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Uint128EventAddBlackBlock)
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
		it.Event = new(Uint128EventAddBlackBlock)
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
func (it *Uint128EventAddBlackBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Uint128EventAddBlackBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Uint128EventAddBlackBlock represents a AddBlackBlock event raised by the Uint128Event contract.
type Uint128EventAddBlackBlock struct {
	I          *big.Int
	BeginBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddBlackBlock is a free log retrieval operation binding the contract event 0x13aaaa4feed46599d583b7da15a15ceba52f8313e21dd345ccf2d839e253daca.
//
// Solidity: event AddBlackBlock(uint256 i, uint128 _beginBlock, uint128 _endBlock)
func (_Uint128Event *Uint128EventFilterer) FilterAddBlackBlock(opts *bind.FilterOpts) (*Uint128EventAddBlackBlockIterator, error) {

	logs, sub, err := _Uint128Event.contract.FilterLogs(opts, "AddBlackBlock")
	if err != nil {
		return nil, err
	}
	return &Uint128EventAddBlackBlockIterator{contract: _Uint128Event.contract, event: "AddBlackBlock", logs: logs, sub: sub}, nil
}

// WatchAddBlackBlock is a free log subscription operation binding the contract event 0x13aaaa4feed46599d583b7da15a15ceba52f8313e21dd345ccf2d839e253daca.
//
// Solidity: event AddBlackBlock(uint256 i, uint128 _beginBlock, uint128 _endBlock)
func (_Uint128Event *Uint128EventFilterer) WatchAddBlackBlock(opts *bind.WatchOpts, sink chan<- *Uint128EventAddBlackBlock) (event.Subscription, error) {

	logs, sub, err := _Uint128Event.contract.WatchLogs(opts, "AddBlackBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Uint128EventAddBlackBlock)
				if err := _Uint128Event.contract.UnpackLog(event, "AddBlackBlock", log); err != nil {
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

// ParseAddBlackBlock is a log parse operation binding the contract event 0x13aaaa4feed46599d583b7da15a15ceba52f8313e21dd345ccf2d839e253daca.
//
// Solidity: event AddBlackBlock(uint256 i, uint128 _beginBlock, uint128 _endBlock)
func (_Uint128Event *Uint128EventFilterer) ParseAddBlackBlock(log types.Log) (*Uint128EventAddBlackBlock, error) {
	event := new(Uint128EventAddBlackBlock)
	if err := _Uint128Event.contract.UnpackLog(event, "AddBlackBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
