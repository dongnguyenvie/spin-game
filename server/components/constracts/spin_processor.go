// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package constracts

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

// SpincontractMetaData contains all meta data concerning the Spincontract contract.
var SpincontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"totalDepositOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositUsd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawUSdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SpincontractABI is the input ABI used to generate the binding from.
// Deprecated: Use SpincontractMetaData.ABI instead.
var SpincontractABI = SpincontractMetaData.ABI

// Spincontract is an auto generated Go binding around an Ethereum contract.
type Spincontract struct {
	SpincontractCaller     // Read-only binding to the contract
	SpincontractTransactor // Write-only binding to the contract
	SpincontractFilterer   // Log filterer for contract events
}

// SpincontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpincontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpincontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpincontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpincontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpincontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpincontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpincontractSession struct {
	Contract     *Spincontract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpincontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpincontractCallerSession struct {
	Contract *SpincontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SpincontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpincontractTransactorSession struct {
	Contract     *SpincontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SpincontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpincontractRaw struct {
	Contract *Spincontract // Generic contract binding to access the raw methods on
}

// SpincontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpincontractCallerRaw struct {
	Contract *SpincontractCaller // Generic read-only contract binding to access the raw methods on
}

// SpincontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpincontractTransactorRaw struct {
	Contract *SpincontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpincontract creates a new instance of Spincontract, bound to a specific deployed contract.
func NewSpincontract(address common.Address, backend bind.ContractBackend) (*Spincontract, error) {
	contract, err := bindSpincontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Spincontract{SpincontractCaller: SpincontractCaller{contract: contract}, SpincontractTransactor: SpincontractTransactor{contract: contract}, SpincontractFilterer: SpincontractFilterer{contract: contract}}, nil
}

// NewSpincontractCaller creates a new read-only instance of Spincontract, bound to a specific deployed contract.
func NewSpincontractCaller(address common.Address, caller bind.ContractCaller) (*SpincontractCaller, error) {
	contract, err := bindSpincontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpincontractCaller{contract: contract}, nil
}

// NewSpincontractTransactor creates a new write-only instance of Spincontract, bound to a specific deployed contract.
func NewSpincontractTransactor(address common.Address, transactor bind.ContractTransactor) (*SpincontractTransactor, error) {
	contract, err := bindSpincontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpincontractTransactor{contract: contract}, nil
}

// NewSpincontractFilterer creates a new log filterer instance of Spincontract, bound to a specific deployed contract.
func NewSpincontractFilterer(address common.Address, filterer bind.ContractFilterer) (*SpincontractFilterer, error) {
	contract, err := bindSpincontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpincontractFilterer{contract: contract}, nil
}

// bindSpincontract binds a generic wrapper to an already deployed contract.
func bindSpincontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SpincontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spincontract *SpincontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Spincontract.Contract.SpincontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spincontract *SpincontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spincontract.Contract.SpincontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spincontract *SpincontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spincontract.Contract.SpincontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Spincontract *SpincontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Spincontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Spincontract *SpincontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spincontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Spincontract *SpincontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Spincontract.Contract.contract.Transact(opts, method, params...)
}

// TotalDepositOf is a free data retrieval call binding the contract method 0xc5d511e1.
//
// Solidity: function totalDepositOf(address from) view returns(uint256)
func (_Spincontract *SpincontractCaller) TotalDepositOf(opts *bind.CallOpts, from common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Spincontract.contract.Call(opts, &out, "totalDepositOf", from)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositOf is a free data retrieval call binding the contract method 0xc5d511e1.
//
// Solidity: function totalDepositOf(address from) view returns(uint256)
func (_Spincontract *SpincontractSession) TotalDepositOf(from common.Address) (*big.Int, error) {
	return _Spincontract.Contract.TotalDepositOf(&_Spincontract.CallOpts, from)
}

// TotalDepositOf is a free data retrieval call binding the contract method 0xc5d511e1.
//
// Solidity: function totalDepositOf(address from) view returns(uint256)
func (_Spincontract *SpincontractCallerSession) TotalDepositOf(from common.Address) (*big.Int, error) {
	return _Spincontract.Contract.TotalDepositOf(&_Spincontract.CallOpts, from)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Spincontract *SpincontractTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spincontract.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Spincontract *SpincontractSession) Deposit() (*types.Transaction, error) {
	return _Spincontract.Contract.Deposit(&_Spincontract.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Spincontract *SpincontractTransactorSession) Deposit() (*types.Transaction, error) {
	return _Spincontract.Contract.Deposit(&_Spincontract.TransactOpts)
}

// DepositUsd is a paid mutator transaction binding the contract method 0x4cc6fd15.
//
// Solidity: function depositUsd(uint256 amount) returns()
func (_Spincontract *SpincontractTransactor) DepositUsd(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Spincontract.contract.Transact(opts, "depositUsd", amount)
}

// DepositUsd is a paid mutator transaction binding the contract method 0x4cc6fd15.
//
// Solidity: function depositUsd(uint256 amount) returns()
func (_Spincontract *SpincontractSession) DepositUsd(amount *big.Int) (*types.Transaction, error) {
	return _Spincontract.Contract.DepositUsd(&_Spincontract.TransactOpts, amount)
}

// DepositUsd is a paid mutator transaction binding the contract method 0x4cc6fd15.
//
// Solidity: function depositUsd(uint256 amount) returns()
func (_Spincontract *SpincontractTransactorSession) DepositUsd(amount *big.Int) (*types.Transaction, error) {
	return _Spincontract.Contract.DepositUsd(&_Spincontract.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Spincontract *SpincontractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spincontract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Spincontract *SpincontractSession) Withdraw() (*types.Transaction, error) {
	return _Spincontract.Contract.Withdraw(&_Spincontract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Spincontract *SpincontractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Spincontract.Contract.Withdraw(&_Spincontract.TransactOpts)
}

// WithdrawUSdt is a paid mutator transaction binding the contract method 0xf2195ccd.
//
// Solidity: function withdrawUSdt() returns()
func (_Spincontract *SpincontractTransactor) WithdrawUSdt(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Spincontract.contract.Transact(opts, "withdrawUSdt")
}

// WithdrawUSdt is a paid mutator transaction binding the contract method 0xf2195ccd.
//
// Solidity: function withdrawUSdt() returns()
func (_Spincontract *SpincontractSession) WithdrawUSdt() (*types.Transaction, error) {
	return _Spincontract.Contract.WithdrawUSdt(&_Spincontract.TransactOpts)
}

// WithdrawUSdt is a paid mutator transaction binding the contract method 0xf2195ccd.
//
// Solidity: function withdrawUSdt() returns()
func (_Spincontract *SpincontractTransactorSession) WithdrawUSdt() (*types.Transaction, error) {
	return _Spincontract.Contract.WithdrawUSdt(&_Spincontract.TransactOpts)
}

// SpincontractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Spincontract contract.
type SpincontractDepositIterator struct {
	Event *SpincontractDeposit // Event containing the contract specifics and raw log

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
func (it *SpincontractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpincontractDeposit)
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
		it.Event = new(SpincontractDeposit)
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
func (it *SpincontractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpincontractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpincontractDeposit represents a Deposit event raised by the Spincontract contract.
type SpincontractDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 amount)
func (_Spincontract *SpincontractFilterer) FilterDeposit(opts *bind.FilterOpts) (*SpincontractDepositIterator, error) {

	logs, sub, err := _Spincontract.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &SpincontractDepositIterator{contract: _Spincontract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 amount)
func (_Spincontract *SpincontractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SpincontractDeposit) (event.Subscription, error) {

	logs, sub, err := _Spincontract.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpincontractDeposit)
				if err := _Spincontract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 amount)
func (_Spincontract *SpincontractFilterer) ParseDeposit(log types.Log) (*SpincontractDeposit, error) {
	event := new(SpincontractDeposit)
	if err := _Spincontract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
