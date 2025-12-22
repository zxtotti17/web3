// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package class5

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
	_ = abi.ConvertType
)

// Class5MetaData contains all meta data concerning the Class5 contract.
var Class5MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"by\",\"type\":\"uint256\"}],\"name\":\"Increment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"by\",\"type\":\"uint256\"}],\"name\":\"incBy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506104008061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80630c55699c1461004e578063371303c01461006c5780636d4ce63c1461007657806370119d0614610094575b5f5ffd5b6100566100b0565b60405161006391906101b8565b60405180910390f35b6100746100b5565b005b61007e610105565b60405161008b91906101b8565b60405180910390f35b6100ae60048036038101906100a991906101ff565b61010d565b005b5f5481565b5f5f8154809291906100c690610257565b91905055507f51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a8160016040516100fb91906102e0565b60405180910390a1565b5f5f54905090565b5f811161014f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161014690610379565b60405180910390fd5b805f5f82825461015f9190610397565b925050819055507f51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a818160405161019591906101b8565b60405180910390a150565b5f819050919050565b6101b2816101a0565b82525050565b5f6020820190506101cb5f8301846101a9565b92915050565b5f5ffd5b6101de816101a0565b81146101e8575f5ffd5b50565b5f813590506101f9816101d5565b92915050565b5f60208284031215610214576102136101d1565b5b5f610221848285016101eb565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610261826101a0565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036102935761029261022a565b5b600182019050919050565b5f819050919050565b5f819050919050565b5f6102ca6102c56102c08461029e565b6102a7565b6101a0565b9050919050565b6102da816102b0565b82525050565b5f6020820190506102f35f8301846102d1565b92915050565b5f82825260208201905092915050565b7f696e6342793a20696e6372656d656e742073686f756c6420626520706f7369745f8201527f6976650000000000000000000000000000000000000000000000000000000000602082015250565b5f6103636023836102f9565b915061036e82610309565b604082019050919050565b5f6020820190508181035f83015261039081610357565b9050919050565b5f6103a1826101a0565b91506103ac836101a0565b92508282019050808211156103c4576103c361022a565b5b9291505056fea26469706673582212205bce60fa00c7978600a265f43bd089482dc7c80f8b782cb29021fcafde0c504564736f6c63430008210033",
}

// Class5ABI is the input ABI used to generate the binding from.
// Deprecated: Use Class5MetaData.ABI instead.
var Class5ABI = Class5MetaData.ABI

// Class5Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Class5MetaData.Bin instead.
var Class5Bin = Class5MetaData.Bin

// DeployClass5 deploys a new Ethereum contract, binding an instance of Class5 to it.
func DeployClass5(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Class5, error) {
	parsed, err := Class5MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Class5Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Class5{Class5Caller: Class5Caller{contract: contract}, Class5Transactor: Class5Transactor{contract: contract}, Class5Filterer: Class5Filterer{contract: contract}}, nil
}

// Class5 is an auto generated Go binding around an Ethereum contract.
type Class5 struct {
	Class5Caller     // Read-only binding to the contract
	Class5Transactor // Write-only binding to the contract
	Class5Filterer   // Log filterer for contract events
}

// Class5Caller is an auto generated read-only Go binding around an Ethereum contract.
type Class5Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Class5Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Class5Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Class5Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Class5Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Class5Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Class5Session struct {
	Contract     *Class5           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Class5CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Class5CallerSession struct {
	Contract *Class5Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Class5TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Class5TransactorSession struct {
	Contract     *Class5Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Class5Raw is an auto generated low-level Go binding around an Ethereum contract.
type Class5Raw struct {
	Contract *Class5 // Generic contract binding to access the raw methods on
}

// Class5CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Class5CallerRaw struct {
	Contract *Class5Caller // Generic read-only contract binding to access the raw methods on
}

// Class5TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Class5TransactorRaw struct {
	Contract *Class5Transactor // Generic write-only contract binding to access the raw methods on
}

// NewClass5 creates a new instance of Class5, bound to a specific deployed contract.
func NewClass5(address common.Address, backend bind.ContractBackend) (*Class5, error) {
	contract, err := bindClass5(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Class5{Class5Caller: Class5Caller{contract: contract}, Class5Transactor: Class5Transactor{contract: contract}, Class5Filterer: Class5Filterer{contract: contract}}, nil
}

// NewClass5Caller creates a new read-only instance of Class5, bound to a specific deployed contract.
func NewClass5Caller(address common.Address, caller bind.ContractCaller) (*Class5Caller, error) {
	contract, err := bindClass5(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Class5Caller{contract: contract}, nil
}

// NewClass5Transactor creates a new write-only instance of Class5, bound to a specific deployed contract.
func NewClass5Transactor(address common.Address, transactor bind.ContractTransactor) (*Class5Transactor, error) {
	contract, err := bindClass5(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Class5Transactor{contract: contract}, nil
}

// NewClass5Filterer creates a new log filterer instance of Class5, bound to a specific deployed contract.
func NewClass5Filterer(address common.Address, filterer bind.ContractFilterer) (*Class5Filterer, error) {
	contract, err := bindClass5(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Class5Filterer{contract: contract}, nil
}

// bindClass5 binds a generic wrapper to an already deployed contract.
func bindClass5(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Class5MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Class5 *Class5Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Class5.Contract.Class5Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Class5 *Class5Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Class5.Contract.Class5Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Class5 *Class5Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Class5.Contract.Class5Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Class5 *Class5CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Class5.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Class5 *Class5TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Class5.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Class5 *Class5TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Class5.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Class5 *Class5Caller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Class5.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Class5 *Class5Session) Get() (*big.Int, error) {
	return _Class5.Contract.Get(&_Class5.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Class5 *Class5CallerSession) Get() (*big.Int, error) {
	return _Class5.Contract.Get(&_Class5.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Class5 *Class5Caller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Class5.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Class5 *Class5Session) X() (*big.Int, error) {
	return _Class5.Contract.X(&_Class5.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Class5 *Class5CallerSession) X() (*big.Int, error) {
	return _Class5.Contract.X(&_Class5.CallOpts)
}

// Inc is a paid mutator transaction binding the contract method 0x371303c0.
//
// Solidity: function inc() returns()
func (_Class5 *Class5Transactor) Inc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Class5.contract.Transact(opts, "inc")
}

// Inc is a paid mutator transaction binding the contract method 0x371303c0.
//
// Solidity: function inc() returns()
func (_Class5 *Class5Session) Inc() (*types.Transaction, error) {
	return _Class5.Contract.Inc(&_Class5.TransactOpts)
}

// Inc is a paid mutator transaction binding the contract method 0x371303c0.
//
// Solidity: function inc() returns()
func (_Class5 *Class5TransactorSession) Inc() (*types.Transaction, error) {
	return _Class5.Contract.Inc(&_Class5.TransactOpts)
}

// IncBy is a paid mutator transaction binding the contract method 0x70119d06.
//
// Solidity: function incBy(uint256 by) returns()
func (_Class5 *Class5Transactor) IncBy(opts *bind.TransactOpts, by *big.Int) (*types.Transaction, error) {
	return _Class5.contract.Transact(opts, "incBy", by)
}

// IncBy is a paid mutator transaction binding the contract method 0x70119d06.
//
// Solidity: function incBy(uint256 by) returns()
func (_Class5 *Class5Session) IncBy(by *big.Int) (*types.Transaction, error) {
	return _Class5.Contract.IncBy(&_Class5.TransactOpts, by)
}

// IncBy is a paid mutator transaction binding the contract method 0x70119d06.
//
// Solidity: function incBy(uint256 by) returns()
func (_Class5 *Class5TransactorSession) IncBy(by *big.Int) (*types.Transaction, error) {
	return _Class5.Contract.IncBy(&_Class5.TransactOpts, by)
}

// Class5IncrementIterator is returned from FilterIncrement and is used to iterate over the raw logs and unpacked data for Increment events raised by the Class5 contract.
type Class5IncrementIterator struct {
	Event *Class5Increment // Event containing the contract specifics and raw log

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
func (it *Class5IncrementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Class5Increment)
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
		it.Event = new(Class5Increment)
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
func (it *Class5IncrementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Class5IncrementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Class5Increment represents a Increment event raised by the Class5 contract.
type Class5Increment struct {
	By  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIncrement is a free log retrieval operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 by)
func (_Class5 *Class5Filterer) FilterIncrement(opts *bind.FilterOpts) (*Class5IncrementIterator, error) {

	logs, sub, err := _Class5.contract.FilterLogs(opts, "Increment")
	if err != nil {
		return nil, err
	}
	return &Class5IncrementIterator{contract: _Class5.contract, event: "Increment", logs: logs, sub: sub}, nil
}

// WatchIncrement is a free log subscription operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 by)
func (_Class5 *Class5Filterer) WatchIncrement(opts *bind.WatchOpts, sink chan<- *Class5Increment) (event.Subscription, error) {

	logs, sub, err := _Class5.contract.WatchLogs(opts, "Increment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Class5Increment)
				if err := _Class5.contract.UnpackLog(event, "Increment", log); err != nil {
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

// ParseIncrement is a log parse operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 by)
func (_Class5 *Class5Filterer) ParseIncrement(log types.Log) (*Class5Increment, error) {
	event := new(Class5Increment)
	if err := _Class5.contract.UnpackLog(event, "Increment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
