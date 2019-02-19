// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethertoken

import (
	"math/big"
	"strings"

	ethereum "github.com/Onther-Tech/plasma-evm"
	"github.com/Onther-Tech/plasma-evm/accounts/abi"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/core/types"
	"github.com/Onther-Tech/plasma-evm/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x608060405234801561001057600080fd5b5061027a806100206000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005b57806370a0823114610082578063a9059cbb146100b0575b600080fd5b34801561006757600080fd5b506100706100f5565b60408051918252519081900360200190f35b34801561008e57600080fd5b5061007073ffffffffffffffffffffffffffffffffffffffff600435166100fb565b3480156100bc57600080fd5b506100e173ffffffffffffffffffffffffffffffffffffffff60043516602435610123565b604080519115158252519081900360200190f35b60015490565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b3360009081526020819052604081205482111561013f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316151561016157600080fd5b33600090815260208190526040902054610181908363ffffffff61022916565b336000908152602081905260408082209290925573ffffffffffffffffffffffffffffffffffffffff8516815220546101c0908363ffffffff61023b16565b73ffffffffffffffffffffffffffffffffffffffff8416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b60008282111561023557fe5b50900390565b8181018281101561024857fe5b929150505600a165627a7a723058205f95e97f3507015fb923be4e259ec86e0bb15cc1125ddb305716e7f8d3c3db7e0029`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, _who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, _who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, _to, _value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// EtherTokenABI is the input ABI used to generate the binding from.
const EtherTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rootchain\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"swapFromEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"swapEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"swapToEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"getBalanceTrieKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootchain\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"applyRequestInRootChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"applyRequestInChildChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_development\",\"type\":\"bool\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_swapEnabled\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Depositted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// EtherTokenBin is the compiled bytecode used for deploying new contracts.
const EtherTokenBin = `0x608060405234801561001057600080fd5b50604051606080610f3b8339810160409081528151602083015191909201516003805461ff0019166101008515150217905560048054600160a060020a031916600160a060020a03841690811790915515838061008457508080156100725750815b80610084575080158015610084575081155b151561008f57600080fd5b5060048054911515740100000000000000000000000000000000000000000260a060020a60ff02199092169190911790555050610e6a806100d16000396000f30060806040526004361061011c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b38114610126578063158ef93e1461015e57806318160ddd1461017357806319ab453c1461019a57806323b872dd146101bb5780632e1a7d4d146101e557806359310da71461011c57806366188463146101fd5780636ddd17131461022157806370a08231146102365780637b929c2714610257578063a9059cbb1461026c578063b06c672914610290578063b18fcfdf146102a8578063b6b55f25146102c9578063bcc7874d146102e1578063d73dd62314610312578063d9afd3a914610336578063dd62ed3e14610365578063e904e3d91461038c578063fc0c546a146103bb575b6101246103d0565b005b34801561013257600080fd5b5061014a600160a060020a0360043516602435610405565b604080519115158252519081900360200190f35b34801561016a57600080fd5b5061014a61046b565b34801561017f57600080fd5b50610188610474565b60408051918252519081900360200190f35b3480156101a657600080fd5b5061014a600160a060020a036004351661047a565b3480156101c757600080fd5b5061014a600160a060020a03600435811690602435166044356104cd565b3480156101f157600080fd5b5061014a600435610630565b34801561020957600080fd5b5061014a600160a060020a0360043516602435610736565b34801561022d57600080fd5b5061014a610825565b34801561024257600080fd5b50610188600160a060020a0360043516610846565b34801561026357600080fd5b5061014a610861565b34801561027857600080fd5b5061014a600160a060020a036004351660243561086f565b34801561029c57600080fd5b5061012460043561093c565b3480156102b457600080fd5b50610188600160a060020a0360043516610988565b3480156102d557600080fd5b5061014a600435610a25565b3480156102ed57600080fd5b506102f6610af2565b60408051600160a060020a039092168252519081900360200190f35b34801561031e57600080fd5b5061014a600160a060020a0360043516602435610b07565b34801561034257600080fd5b5061014a6004351515602435600160a060020a0360443516606435608435610ba0565b34801561037157600080fd5b50610188600160a060020a0360043581169060243516610c61565b34801561039857600080fd5b5061014a6004351515602435600160a060020a0360443516606435608435610c8c565b3480156103c757600080fd5b506102f6610cdc565b60045474010000000000000000000000000000000000000000900460ff1615156103f957600080fd5b6104033334610ceb565b565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60035460ff1681565b60015490565b60035460009060ff161561048d57600080fd5b6003805460ff19600160a060020a03909416620100000275ffffffffffffffffffffffffffffffffffffffff000019909116179290921660011790915590565b600160a060020a0383166000908152602081905260408120548211156104f257600080fd5b600160a060020a038416600090815260026020908152604080832033845290915290205482111561052257600080fd5b600160a060020a038316151561053757600080fd5b600160a060020a038416600090815260208190526040902054610560908363ffffffff610d6e16565b600160a060020a038086166000908152602081905260408082209390935590851681522054610595908363ffffffff610d8016565b600160a060020a038085166000908152602081815260408083209490945591871681526002825282812033825290915220546105d7908363ffffffff610d6e16565b600160a060020a0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610e1f833981519152929181900390910190a35060019392505050565b60035460009060ff16151561064457600080fd5b61064e3383610d93565b604080513381526020810184905281517f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5929181900390910190a160048054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815233938101939093526024830185905251600160a060020a039091169163a9059cbb9160448083019260209291908290030181600087803b1580156106f757600080fd5b505af115801561070b573d6000803e3d6000fd5b505050506040513d602081101561072157600080fd5b5051151561072e57600080fd5b506001919050565b336000908152600260209081526040808320600160a060020a038616845290915281205480831061078a57336000908152600260209081526040808320600160a060020a03881684529091528120556107bf565b61079a818463ffffffff610d6e16565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60045474010000000000000000000000000000000000000000900460ff1681565b600160a060020a031660009081526020819052604090205490565b600354610100900460ff1681565b3360009081526020819052604081205482111561088b57600080fd5b600160a060020a03831615156108a057600080fd5b336000908152602081905260409020546108c0908363ffffffff610d6e16565b3360009081526020819052604080822092909255600160a060020a038516815220546108f2908363ffffffff610d8016565b600160a060020a03841660008181526020818152604091829020939093558051858152905191923392600080516020610e1f8339815191529281900390910190a350600192915050565b60045474010000000000000000000000000000000000000000900460ff16151561096557600080fd5b6109703330836104cd565b151561097b57600080fd5b6109853382610d93565b50565b60408051600060208083018290526c01000000000000000000000000600160a060020a038616028385015283516034818503018152605490930193849052825191939182918401908083835b602083106109f35780518252601f1990920191602091820191016109d4565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b60035460009060ff161515610a3957600080fd5b610a433383610ceb565b604080513381526020810184905281517f46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4929181900390910190a160048054604080517f23b872dd00000000000000000000000000000000000000000000000000000000815233938101939093523060248401526044830185905251600160a060020a03909116916323b872dd9160648083019260209291908290030181600087803b1580156106f757600080fd5b600354620100009004600160a060020a031681565b336000908152600260209081526040808320600160a060020a0386168452909152812054610b3b908363ffffffff610d8016565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b60035460009060ff161515610bb457600080fd5b600354620100009004600160a060020a03163314610bd157600080fd5b610bda84610988565b8314610be557600080fd5b8515610bfa57610bf58483610ceb565b610c04565b610c048483610d93565b604080518715158152600160a060020a03861660208201528082018590526060810184905290517f72507fe0ffa88ff97b08e473da2cc7f83aecabbf60819dd1972f8736c309474f9181900360800190a150600195945050505050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600090610100900460ff1680610ca3575033155b1515610cae57600080fd5b610cb784610988565b8314610cc257600080fd5b8515610cd257610bf58483610d93565b610c048483610ceb565b600454600160a060020a031681565b600154610cfe908263ffffffff610d8016565b600155600160a060020a038216600090815260208190526040902054610d2a908263ffffffff610d8016565b600160a060020a038316600081815260208181526040808320949094558351858152935192939192600080516020610e1f8339815191529281900390910190a35050565b600082821115610d7a57fe5b50900390565b81810182811015610d8d57fe5b92915050565b600160a060020a038216600090815260208190526040902054610dbc908263ffffffff610d6e16565b600160a060020a038316600090815260208190526040902055600154610de8908263ffffffff610d6e16565b600155604080518281529051600091600160a060020a03851691600080516020610e1f8339815191529181900360200190a350505600ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa165627a7a723058200ddca94f21e8128d293ecfd4a349987d453fe63166bbdc8beba4f9669847cdce0029`

// DeployEtherToken deploys a new Ethereum contract, binding an instance of EtherToken to it.
func DeployEtherToken(auth *bind.TransactOpts, backend bind.ContractBackend, _development bool, _token common.Address, _swapEnabled bool) (common.Address, *types.Transaction, *EtherToken, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EtherTokenBin), backend, _development, _token, _swapEnabled)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherToken{EtherTokenCaller: EtherTokenCaller{contract: contract}, EtherTokenTransactor: EtherTokenTransactor{contract: contract}, EtherTokenFilterer: EtherTokenFilterer{contract: contract}}, nil
}

// EtherToken is an auto generated Go binding around an Ethereum contract.
type EtherToken struct {
	EtherTokenCaller     // Read-only binding to the contract
	EtherTokenTransactor // Write-only binding to the contract
	EtherTokenFilterer   // Log filterer for contract events
}

// EtherTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherTokenSession struct {
	Contract     *EtherToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherTokenCallerSession struct {
	Contract *EtherTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EtherTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherTokenTransactorSession struct {
	Contract     *EtherTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EtherTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherTokenRaw struct {
	Contract *EtherToken // Generic contract binding to access the raw methods on
}

// EtherTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherTokenCallerRaw struct {
	Contract *EtherTokenCaller // Generic read-only contract binding to access the raw methods on
}

// EtherTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherTokenTransactorRaw struct {
	Contract *EtherTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherToken creates a new instance of EtherToken, bound to a specific deployed contract.
func NewEtherToken(address common.Address, backend bind.ContractBackend) (*EtherToken, error) {
	contract, err := bindEtherToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherToken{EtherTokenCaller: EtherTokenCaller{contract: contract}, EtherTokenTransactor: EtherTokenTransactor{contract: contract}, EtherTokenFilterer: EtherTokenFilterer{contract: contract}}, nil
}

// NewEtherTokenCaller creates a new read-only instance of EtherToken, bound to a specific deployed contract.
func NewEtherTokenCaller(address common.Address, caller bind.ContractCaller) (*EtherTokenCaller, error) {
	contract, err := bindEtherToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherTokenCaller{contract: contract}, nil
}

// NewEtherTokenTransactor creates a new write-only instance of EtherToken, bound to a specific deployed contract.
func NewEtherTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherTokenTransactor, error) {
	contract, err := bindEtherToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherTokenTransactor{contract: contract}, nil
}

// NewEtherTokenFilterer creates a new log filterer instance of EtherToken, bound to a specific deployed contract.
func NewEtherTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherTokenFilterer, error) {
	contract, err := bindEtherToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherTokenFilterer{contract: contract}, nil
}

// bindEtherToken binds a generic wrapper to an already deployed contract.
func bindEtherToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherToken *EtherTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherToken.Contract.EtherTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherToken *EtherTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherToken.Contract.EtherTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherToken *EtherTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherToken.Contract.EtherTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherToken *EtherTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherToken *EtherTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherToken *EtherTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_EtherToken *EtherTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_EtherToken *EtherTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EtherToken.Contract.Allowance(&_EtherToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_EtherToken *EtherTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EtherToken.Contract.Allowance(&_EtherToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_EtherToken *EtherTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_EtherToken *EtherTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _EtherToken.Contract.BalanceOf(&_EtherToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_EtherToken *EtherTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _EtherToken.Contract.BalanceOf(&_EtherToken.CallOpts, _owner)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_EtherToken *EtherTokenCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_EtherToken *EtherTokenSession) Development() (bool, error) {
	return _EtherToken.Contract.Development(&_EtherToken.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_EtherToken *EtherTokenCallerSession) Development() (bool, error) {
	return _EtherToken.Contract.Development(&_EtherToken.CallOpts)
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(_who address) constant returns(bytes32)
func (_EtherToken *EtherTokenCaller) GetBalanceTrieKey(opts *bind.CallOpts, _who common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "getBalanceTrieKey", _who)
	return *ret0, err
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(_who address) constant returns(bytes32)
func (_EtherToken *EtherTokenSession) GetBalanceTrieKey(_who common.Address) ([32]byte, error) {
	return _EtherToken.Contract.GetBalanceTrieKey(&_EtherToken.CallOpts, _who)
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(_who address) constant returns(bytes32)
func (_EtherToken *EtherTokenCallerSession) GetBalanceTrieKey(_who common.Address) ([32]byte, error) {
	return _EtherToken.Contract.GetBalanceTrieKey(&_EtherToken.CallOpts, _who)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_EtherToken *EtherTokenCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_EtherToken *EtherTokenSession) Initialized() (bool, error) {
	return _EtherToken.Contract.Initialized(&_EtherToken.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_EtherToken *EtherTokenCallerSession) Initialized() (bool, error) {
	return _EtherToken.Contract.Initialized(&_EtherToken.CallOpts)
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_EtherToken *EtherTokenCaller) Rootchain(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "rootchain")
	return *ret0, err
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_EtherToken *EtherTokenSession) Rootchain() (common.Address, error) {
	return _EtherToken.Contract.Rootchain(&_EtherToken.CallOpts)
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_EtherToken *EtherTokenCallerSession) Rootchain() (common.Address, error) {
	return _EtherToken.Contract.Rootchain(&_EtherToken.CallOpts)
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() constant returns(bool)
func (_EtherToken *EtherTokenCaller) SwapEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "swapEnabled")
	return *ret0, err
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() constant returns(bool)
func (_EtherToken *EtherTokenSession) SwapEnabled() (bool, error) {
	return _EtherToken.Contract.SwapEnabled(&_EtherToken.CallOpts)
}

// SwapEnabled is a free data retrieval call binding the contract method 0x6ddd1713.
//
// Solidity: function swapEnabled() constant returns(bool)
func (_EtherToken *EtherTokenCallerSession) SwapEnabled() (bool, error) {
	return _EtherToken.Contract.SwapEnabled(&_EtherToken.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_EtherToken *EtherTokenCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_EtherToken *EtherTokenSession) Token() (common.Address, error) {
	return _EtherToken.Contract.Token(&_EtherToken.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_EtherToken *EtherTokenCallerSession) Token() (common.Address, error) {
	return _EtherToken.Contract.Token(&_EtherToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EtherToken *EtherTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EtherToken *EtherTokenSession) TotalSupply() (*big.Int, error) {
	return _EtherToken.Contract.TotalSupply(&_EtherToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EtherToken *EtherTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _EtherToken.Contract.TotalSupply(&_EtherToken.CallOpts)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_EtherToken *EtherTokenTransactor) ApplyRequestInChildChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "applyRequestInChildChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_EtherToken *EtherTokenSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _EtherToken.Contract.ApplyRequestInChildChain(&_EtherToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_EtherToken *EtherTokenTransactorSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _EtherToken.Contract.ApplyRequestInChildChain(&_EtherToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_EtherToken *EtherTokenTransactor) ApplyRequestInRootChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "applyRequestInRootChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_EtherToken *EtherTokenSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _EtherToken.Contract.ApplyRequestInRootChain(&_EtherToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_EtherToken *EtherTokenTransactorSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _EtherToken.Contract.ApplyRequestInRootChain(&_EtherToken.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_EtherToken *EtherTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_EtherToken *EtherTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Approve(&_EtherToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Approve(&_EtherToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_EtherToken *EtherTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_EtherToken *EtherTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.DecreaseApproval(&_EtherToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.DecreaseApproval(&_EtherToken.TransactOpts, _spender, _subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_EtherToken *EtherTokenTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_EtherToken *EtherTokenSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Deposit(&_EtherToken.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Deposit(&_EtherToken.TransactOpts, _amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_EtherToken *EtherTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_EtherToken *EtherTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.IncreaseApproval(&_EtherToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.IncreaseApproval(&_EtherToken.TransactOpts, _spender, _addedValue)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(_rootchain address) returns(bool)
func (_EtherToken *EtherTokenTransactor) Init(opts *bind.TransactOpts, _rootchain common.Address) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "init", _rootchain)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(_rootchain address) returns(bool)
func (_EtherToken *EtherTokenSession) Init(_rootchain common.Address) (*types.Transaction, error) {
	return _EtherToken.Contract.Init(&_EtherToken.TransactOpts, _rootchain)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(_rootchain address) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) Init(_rootchain common.Address) (*types.Transaction, error) {
	return _EtherToken.Contract.Init(&_EtherToken.TransactOpts, _rootchain)
}

// SwapFromEth is a paid mutator transaction binding the contract method 0x59310da7.
//
// Solidity: function swapFromEth() returns()
func (_EtherToken *EtherTokenTransactor) SwapFromEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "swapFromEth")
}

// SwapFromEth is a paid mutator transaction binding the contract method 0x59310da7.
//
// Solidity: function swapFromEth() returns()
func (_EtherToken *EtherTokenSession) SwapFromEth() (*types.Transaction, error) {
	return _EtherToken.Contract.SwapFromEth(&_EtherToken.TransactOpts)
}

// SwapFromEth is a paid mutator transaction binding the contract method 0x59310da7.
//
// Solidity: function swapFromEth() returns()
func (_EtherToken *EtherTokenTransactorSession) SwapFromEth() (*types.Transaction, error) {
	return _EtherToken.Contract.SwapFromEth(&_EtherToken.TransactOpts)
}

// SwapToEth is a paid mutator transaction binding the contract method 0xb06c6729.
//
// Solidity: function swapToEth(_amount uint256) returns()
func (_EtherToken *EtherTokenTransactor) SwapToEth(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "swapToEth", _amount)
}

// SwapToEth is a paid mutator transaction binding the contract method 0xb06c6729.
//
// Solidity: function swapToEth(_amount uint256) returns()
func (_EtherToken *EtherTokenSession) SwapToEth(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.SwapToEth(&_EtherToken.TransactOpts, _amount)
}

// SwapToEth is a paid mutator transaction binding the contract method 0xb06c6729.
//
// Solidity: function swapToEth(_amount uint256) returns()
func (_EtherToken *EtherTokenTransactorSession) SwapToEth(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.SwapToEth(&_EtherToken.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_EtherToken *EtherTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_EtherToken *EtherTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Transfer(&_EtherToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Transfer(&_EtherToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_EtherToken *EtherTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_EtherToken *EtherTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.TransferFrom(&_EtherToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.TransferFrom(&_EtherToken.TransactOpts, _from, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns(bool)
func (_EtherToken *EtherTokenTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns(bool)
func (_EtherToken *EtherTokenSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Withdraw(&_EtherToken.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns(bool)
func (_EtherToken *EtherTokenTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _EtherToken.Contract.Withdraw(&_EtherToken.TransactOpts, _amount)
}

// EtherTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EtherToken contract.
type EtherTokenApprovalIterator struct {
	Event *EtherTokenApproval // Event containing the contract specifics and raw log

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
func (it *EtherTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenApproval)
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
		it.Event = new(EtherTokenApproval)
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
func (it *EtherTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenApproval represents a Approval event raised by the EtherToken contract.
type EtherTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_EtherToken *EtherTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EtherTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EtherTokenApprovalIterator{contract: _EtherToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_EtherToken *EtherTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EtherTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenApproval)
				if err := _EtherToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// EtherTokenDeposittedIterator is returned from FilterDepositted and is used to iterate over the raw logs and unpacked data for Depositted events raised by the EtherToken contract.
type EtherTokenDeposittedIterator struct {
	Event *EtherTokenDepositted // Event containing the contract specifics and raw log

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
func (it *EtherTokenDeposittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenDepositted)
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
		it.Event = new(EtherTokenDepositted)
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
func (it *EtherTokenDeposittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenDeposittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenDepositted represents a Depositted event raised by the EtherToken contract.
type EtherTokenDepositted struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDepositted is a free log retrieval operation binding the contract event 0x46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4.
//
// Solidity: e Depositted(_from address, _value uint256)
func (_EtherToken *EtherTokenFilterer) FilterDepositted(opts *bind.FilterOpts) (*EtherTokenDeposittedIterator, error) {

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Depositted")
	if err != nil {
		return nil, err
	}
	return &EtherTokenDeposittedIterator{contract: _EtherToken.contract, event: "Depositted", logs: logs, sub: sub}, nil
}

// WatchDepositted is a free log subscription operation binding the contract event 0x46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4.
//
// Solidity: e Depositted(_from address, _value uint256)
func (_EtherToken *EtherTokenFilterer) WatchDepositted(opts *bind.WatchOpts, sink chan<- *EtherTokenDepositted) (event.Subscription, error) {

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Depositted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenDepositted)
				if err := _EtherToken.contract.UnpackLog(event, "Depositted", log); err != nil {
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

// EtherTokenRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the EtherToken contract.
type EtherTokenRequestCreatedIterator struct {
	Event *EtherTokenRequestCreated // Event containing the contract specifics and raw log

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
func (it *EtherTokenRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenRequestCreated)
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
		it.Event = new(EtherTokenRequestCreated)
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
func (it *EtherTokenRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenRequestCreated represents a RequestCreated event raised by the EtherToken contract.
type EtherTokenRequestCreated struct {
	IsExit    bool
	Requestor common.Address
	TrieKey   [32]byte
	TrieValue [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x72507fe0ffa88ff97b08e473da2cc7f83aecabbf60819dd1972f8736c309474f.
//
// Solidity: e RequestCreated(_isExit bool, _requestor address, _trieKey bytes32, _trieValue bytes32)
func (_EtherToken *EtherTokenFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*EtherTokenRequestCreatedIterator, error) {

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &EtherTokenRequestCreatedIterator{contract: _EtherToken.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x72507fe0ffa88ff97b08e473da2cc7f83aecabbf60819dd1972f8736c309474f.
//
// Solidity: e RequestCreated(_isExit bool, _requestor address, _trieKey bytes32, _trieValue bytes32)
func (_EtherToken *EtherTokenFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *EtherTokenRequestCreated) (event.Subscription, error) {

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenRequestCreated)
				if err := _EtherToken.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// EtherTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EtherToken contract.
type EtherTokenTransferIterator struct {
	Event *EtherTokenTransfer // Event containing the contract specifics and raw log

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
func (it *EtherTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenTransfer)
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
		it.Event = new(EtherTokenTransfer)
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
func (it *EtherTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenTransfer represents a Transfer event raised by the EtherToken contract.
type EtherTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_EtherToken *EtherTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EtherTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EtherTokenTransferIterator{contract: _EtherToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_EtherToken *EtherTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EtherTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenTransfer)
				if err := _EtherToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// EtherTokenWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the EtherToken contract.
type EtherTokenWithdrawnIterator struct {
	Event *EtherTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *EtherTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherTokenWithdrawn)
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
		it.Event = new(EtherTokenWithdrawn)
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
func (it *EtherTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherTokenWithdrawn represents a Withdrawn event raised by the EtherToken contract.
type EtherTokenWithdrawn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(_from address, _value uint256)
func (_EtherToken *EtherTokenFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*EtherTokenWithdrawnIterator, error) {

	logs, sub, err := _EtherToken.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &EtherTokenWithdrawnIterator{contract: _EtherToken.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(_from address, _value uint256)
func (_EtherToken *EtherTokenFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *EtherTokenWithdrawn) (event.Subscription, error) {

	logs, sub, err := _EtherToken.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherTokenWithdrawn)
				if err := _EtherToken.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// RequestableERC20WrapperABI is the input ABI used to generate the binding from.
const RequestableERC20WrapperABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rootchain\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"development\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"getBalanceTrieKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootchain\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"applyRequestInRootChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"requestId\",\"type\":\"uint256\"},{\"name\":\"requestor\",\"type\":\"address\"},{\"name\":\"trieKey\",\"type\":\"bytes32\"},{\"name\":\"trieValue\",\"type\":\"bytes32\"}],\"name\":\"applyRequestInChildChain\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_development\",\"type\":\"bool\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Depositted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_isExit\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_requestor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_trieKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_trieValue\",\"type\":\"bytes32\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// RequestableERC20WrapperBin is the compiled bytecode used for deploying new contracts.
const RequestableERC20WrapperBin = `0x608060405234801561001057600080fd5b50604051604080610de68339810160405280516020909101516003805461ff001916610100931515939093029290921790915560048054600160a060020a031916600160a060020a03909216919091179055610d75806100716000396000f3006080604052600436106100fb5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b38114610100578063158ef93e1461013857806318160ddd1461014d57806319ab453c1461017457806323b872dd146101955780632e1a7d4d146101bf57806366188463146101d757806370a08231146101fb5780637b929c271461021c578063a9059cbb14610231578063b18fcfdf14610255578063b6b55f2514610276578063bcc7874d1461028e578063d73dd623146102bf578063d9afd3a9146102e3578063dd62ed3e14610312578063e904e3d914610339578063fc0c546a14610368575b600080fd5b34801561010c57600080fd5b50610124600160a060020a036004351660243561037d565b604080519115158252519081900360200190f35b34801561014457600080fd5b506101246103e3565b34801561015957600080fd5b506101626103ec565b60408051918252519081900360200190f35b34801561018057600080fd5b50610124600160a060020a03600435166103f2565b3480156101a157600080fd5b50610124600160a060020a0360043581169060243516604435610445565b3480156101cb57600080fd5b506101246004356105a8565b3480156101e357600080fd5b50610124600160a060020a03600435166024356106ae565b34801561020757600080fd5b50610162600160a060020a036004351661079d565b34801561022857600080fd5b506101246107b8565b34801561023d57600080fd5b50610124600160a060020a03600435166024356107c6565b34801561026157600080fd5b50610162600160a060020a0360043516610893565b34801561028257600080fd5b50610124600435610930565b34801561029a57600080fd5b506102a36109fd565b60408051600160a060020a039092168252519081900360200190f35b3480156102cb57600080fd5b50610124600160a060020a0360043516602435610a12565b3480156102ef57600080fd5b506101246004351515602435600160a060020a0360443516606435608435610aab565b34801561031e57600080fd5b50610162600160a060020a0360043581169060243516610b6c565b34801561034557600080fd5b506101246004351515602435600160a060020a0360443516606435608435610b97565b34801561037457600080fd5b506102a3610be7565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60035460ff1681565b60015490565b60035460009060ff161561040557600080fd5b6003805460ff19600160a060020a03909416620100000275ffffffffffffffffffffffffffffffffffffffff000019909116179290921660011790915590565b600160a060020a03831660009081526020819052604081205482111561046a57600080fd5b600160a060020a038416600090815260026020908152604080832033845290915290205482111561049a57600080fd5b600160a060020a03831615156104af57600080fd5b600160a060020a0384166000908152602081905260409020546104d8908363ffffffff610bf616565b600160a060020a03808616600090815260208190526040808220939093559085168152205461050d908363ffffffff610c0816565b600160a060020a0380851660009081526020818152604080832094909455918716815260028252828120338252909152205461054f908363ffffffff610bf616565b600160a060020a0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610d2a833981519152929181900390910190a35060019392505050565b60035460009060ff1615156105bc57600080fd5b6105c63383610c1b565b604080513381526020810184905281517f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5929181900390910190a160048054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815233938101939093526024830185905251600160a060020a039091169163a9059cbb9160448083019260209291908290030181600087803b15801561066f57600080fd5b505af1158015610683573d6000803e3d6000fd5b505050506040513d602081101561069957600080fd5b505115156106a657600080fd5b506001919050565b336000908152600260209081526040808320600160a060020a038616845290915281205480831061070257336000908152600260209081526040808320600160a060020a0388168452909152812055610737565b610712818463ffffffff610bf616565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526020819052604090205490565b600354610100900460ff1681565b336000908152602081905260408120548211156107e257600080fd5b600160a060020a03831615156107f757600080fd5b33600090815260208190526040902054610817908363ffffffff610bf616565b3360009081526020819052604080822092909255600160a060020a03851681522054610849908363ffffffff610c0816565b600160a060020a03841660008181526020818152604091829020939093558051858152905191923392600080516020610d2a8339815191529281900390910190a350600192915050565b60408051600060208083018290526c01000000000000000000000000600160a060020a038616028385015283516034818503018152605490930193849052825191939182918401908083835b602083106108fe5780518252601f1990920191602091820191016108df565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912095945050505050565b60035460009060ff16151561094457600080fd5b61094e3383610ca6565b604080513381526020810184905281517f46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4929181900390910190a160048054604080517f23b872dd00000000000000000000000000000000000000000000000000000000815233938101939093523060248401526044830185905251600160a060020a03909116916323b872dd9160648083019260209291908290030181600087803b15801561066f57600080fd5b600354620100009004600160a060020a031681565b336000908152600260209081526040808320600160a060020a0386168452909152812054610a46908363ffffffff610c0816565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b60035460009060ff161515610abf57600080fd5b600354620100009004600160a060020a03163314610adc57600080fd5b610ae584610893565b8314610af057600080fd5b8515610b0557610b008483610ca6565b610b0f565b610b0f8483610c1b565b604080518715158152600160a060020a03861660208201528082018590526060810184905290517f72507fe0ffa88ff97b08e473da2cc7f83aecabbf60819dd1972f8736c309474f9181900360800190a150600195945050505050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600090610100900460ff1680610bae575033155b1515610bb957600080fd5b610bc284610893565b8314610bcd57600080fd5b8515610bdd57610b008483610c1b565b610b0f8483610ca6565b600454600160a060020a031681565b600082821115610c0257fe5b50900390565b81810182811015610c1557fe5b92915050565b600160a060020a038216600090815260208190526040902054610c44908263ffffffff610bf616565b600160a060020a038316600090815260208190526040902055600154610c70908263ffffffff610bf616565b600155604080518281529051600091600160a060020a03851691600080516020610d2a8339815191529181900360200190a35050565b600154610cb9908263ffffffff610c0816565b600155600160a060020a038216600090815260208190526040902054610ce5908263ffffffff610c0816565b600160a060020a038316600081815260208181526040808320949094558351858152935192939192600080516020610d2a8339815191529281900390910190a350505600ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa165627a7a72305820078930fa8d69f83e2bff1ff7d78327142534babc36576147f07b980591e458290029`

// DeployRequestableERC20Wrapper deploys a new Ethereum contract, binding an instance of RequestableERC20Wrapper to it.
func DeployRequestableERC20Wrapper(auth *bind.TransactOpts, backend bind.ContractBackend, _development bool, _token common.Address) (common.Address, *types.Transaction, *RequestableERC20Wrapper, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableERC20WrapperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RequestableERC20WrapperBin), backend, _development, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestableERC20Wrapper{RequestableERC20WrapperCaller: RequestableERC20WrapperCaller{contract: contract}, RequestableERC20WrapperTransactor: RequestableERC20WrapperTransactor{contract: contract}, RequestableERC20WrapperFilterer: RequestableERC20WrapperFilterer{contract: contract}}, nil
}

// RequestableERC20Wrapper is an auto generated Go binding around an Ethereum contract.
type RequestableERC20Wrapper struct {
	RequestableERC20WrapperCaller     // Read-only binding to the contract
	RequestableERC20WrapperTransactor // Write-only binding to the contract
	RequestableERC20WrapperFilterer   // Log filterer for contract events
}

// RequestableERC20WrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestableERC20WrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableERC20WrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestableERC20WrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableERC20WrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestableERC20WrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestableERC20WrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestableERC20WrapperSession struct {
	Contract     *RequestableERC20Wrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RequestableERC20WrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestableERC20WrapperCallerSession struct {
	Contract *RequestableERC20WrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// RequestableERC20WrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestableERC20WrapperTransactorSession struct {
	Contract     *RequestableERC20WrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// RequestableERC20WrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestableERC20WrapperRaw struct {
	Contract *RequestableERC20Wrapper // Generic contract binding to access the raw methods on
}

// RequestableERC20WrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestableERC20WrapperCallerRaw struct {
	Contract *RequestableERC20WrapperCaller // Generic read-only contract binding to access the raw methods on
}

// RequestableERC20WrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestableERC20WrapperTransactorRaw struct {
	Contract *RequestableERC20WrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestableERC20Wrapper creates a new instance of RequestableERC20Wrapper, bound to a specific deployed contract.
func NewRequestableERC20Wrapper(address common.Address, backend bind.ContractBackend) (*RequestableERC20Wrapper, error) {
	contract, err := bindRequestableERC20Wrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20Wrapper{RequestableERC20WrapperCaller: RequestableERC20WrapperCaller{contract: contract}, RequestableERC20WrapperTransactor: RequestableERC20WrapperTransactor{contract: contract}, RequestableERC20WrapperFilterer: RequestableERC20WrapperFilterer{contract: contract}}, nil
}

// NewRequestableERC20WrapperCaller creates a new read-only instance of RequestableERC20Wrapper, bound to a specific deployed contract.
func NewRequestableERC20WrapperCaller(address common.Address, caller bind.ContractCaller) (*RequestableERC20WrapperCaller, error) {
	contract, err := bindRequestableERC20Wrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperCaller{contract: contract}, nil
}

// NewRequestableERC20WrapperTransactor creates a new write-only instance of RequestableERC20Wrapper, bound to a specific deployed contract.
func NewRequestableERC20WrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestableERC20WrapperTransactor, error) {
	contract, err := bindRequestableERC20Wrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperTransactor{contract: contract}, nil
}

// NewRequestableERC20WrapperFilterer creates a new log filterer instance of RequestableERC20Wrapper, bound to a specific deployed contract.
func NewRequestableERC20WrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestableERC20WrapperFilterer, error) {
	contract, err := bindRequestableERC20Wrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperFilterer{contract: contract}, nil
}

// bindRequestableERC20Wrapper binds a generic wrapper to an already deployed contract.
func bindRequestableERC20Wrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestableERC20WrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableERC20Wrapper *RequestableERC20WrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableERC20Wrapper.Contract.RequestableERC20WrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableERC20Wrapper *RequestableERC20WrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.RequestableERC20WrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableERC20Wrapper *RequestableERC20WrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.RequestableERC20WrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RequestableERC20Wrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.Allowance(&_RequestableERC20Wrapper.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.Allowance(&_RequestableERC20Wrapper.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.BalanceOf(&_RequestableERC20Wrapper.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.BalanceOf(&_RequestableERC20Wrapper.CallOpts, _owner)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) Development(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "development")
	return *ret0, err
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Development() (bool, error) {
	return _RequestableERC20Wrapper.Contract.Development(&_RequestableERC20Wrapper.CallOpts)
}

// Development is a free data retrieval call binding the contract method 0x7b929c27.
//
// Solidity: function development() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) Development() (bool, error) {
	return _RequestableERC20Wrapper.Contract.Development(&_RequestableERC20Wrapper.CallOpts)
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(_who address) constant returns(bytes32)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) GetBalanceTrieKey(opts *bind.CallOpts, _who common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "getBalanceTrieKey", _who)
	return *ret0, err
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(_who address) constant returns(bytes32)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) GetBalanceTrieKey(_who common.Address) ([32]byte, error) {
	return _RequestableERC20Wrapper.Contract.GetBalanceTrieKey(&_RequestableERC20Wrapper.CallOpts, _who)
}

// GetBalanceTrieKey is a free data retrieval call binding the contract method 0xb18fcfdf.
//
// Solidity: function getBalanceTrieKey(_who address) constant returns(bytes32)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) GetBalanceTrieKey(_who common.Address) ([32]byte, error) {
	return _RequestableERC20Wrapper.Contract.GetBalanceTrieKey(&_RequestableERC20Wrapper.CallOpts, _who)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Initialized() (bool, error) {
	return _RequestableERC20Wrapper.Contract.Initialized(&_RequestableERC20Wrapper.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) Initialized() (bool, error) {
	return _RequestableERC20Wrapper.Contract.Initialized(&_RequestableERC20Wrapper.CallOpts)
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) Rootchain(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "rootchain")
	return *ret0, err
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Rootchain() (common.Address, error) {
	return _RequestableERC20Wrapper.Contract.Rootchain(&_RequestableERC20Wrapper.CallOpts)
}

// Rootchain is a free data retrieval call binding the contract method 0xbcc7874d.
//
// Solidity: function rootchain() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) Rootchain() (common.Address, error) {
	return _RequestableERC20Wrapper.Contract.Rootchain(&_RequestableERC20Wrapper.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Token() (common.Address, error) {
	return _RequestableERC20Wrapper.Contract.Token(&_RequestableERC20Wrapper.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) Token() (common.Address, error) {
	return _RequestableERC20Wrapper.Contract.Token(&_RequestableERC20Wrapper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RequestableERC20Wrapper.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) TotalSupply() (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.TotalSupply(&_RequestableERC20Wrapper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperCallerSession) TotalSupply() (*big.Int, error) {
	return _RequestableERC20Wrapper.Contract.TotalSupply(&_RequestableERC20Wrapper.CallOpts)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) ApplyRequestInChildChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "applyRequestInChildChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.ApplyRequestInChildChain(&_RequestableERC20Wrapper.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInChildChain is a paid mutator transaction binding the contract method 0xe904e3d9.
//
// Solidity: function applyRequestInChildChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) ApplyRequestInChildChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.ApplyRequestInChildChain(&_RequestableERC20Wrapper.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) ApplyRequestInRootChain(opts *bind.TransactOpts, isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "applyRequestInRootChain", isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.ApplyRequestInRootChain(&_RequestableERC20Wrapper.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// ApplyRequestInRootChain is a paid mutator transaction binding the contract method 0xd9afd3a9.
//
// Solidity: function applyRequestInRootChain(isExit bool, requestId uint256, requestor address, trieKey bytes32, trieValue bytes32) returns(success bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) ApplyRequestInRootChain(isExit bool, requestId *big.Int, requestor common.Address, trieKey [32]byte, trieValue [32]byte) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.ApplyRequestInRootChain(&_RequestableERC20Wrapper.TransactOpts, isExit, requestId, requestor, trieKey, trieValue)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Approve(&_RequestableERC20Wrapper.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Approve(&_RequestableERC20Wrapper.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.DecreaseApproval(&_RequestableERC20Wrapper.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.DecreaseApproval(&_RequestableERC20Wrapper.TransactOpts, _spender, _subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Deposit(&_RequestableERC20Wrapper.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(_amount uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Deposit(&_RequestableERC20Wrapper.TransactOpts, _amount)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.IncreaseApproval(&_RequestableERC20Wrapper.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.IncreaseApproval(&_RequestableERC20Wrapper.TransactOpts, _spender, _addedValue)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(_rootchain address) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) Init(opts *bind.TransactOpts, _rootchain common.Address) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "init", _rootchain)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(_rootchain address) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Init(_rootchain common.Address) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Init(&_RequestableERC20Wrapper.TransactOpts, _rootchain)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(_rootchain address) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) Init(_rootchain common.Address) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Init(&_RequestableERC20Wrapper.TransactOpts, _rootchain)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Transfer(&_RequestableERC20Wrapper.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Transfer(&_RequestableERC20Wrapper.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.TransferFrom(&_RequestableERC20Wrapper.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.TransferFrom(&_RequestableERC20Wrapper.TransactOpts, _from, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Withdraw(&_RequestableERC20Wrapper.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns(bool)
func (_RequestableERC20Wrapper *RequestableERC20WrapperTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _RequestableERC20Wrapper.Contract.Withdraw(&_RequestableERC20Wrapper.TransactOpts, _amount)
}

// RequestableERC20WrapperApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperApprovalIterator struct {
	Event *RequestableERC20WrapperApproval // Event containing the contract specifics and raw log

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
func (it *RequestableERC20WrapperApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableERC20WrapperApproval)
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
		it.Event = new(RequestableERC20WrapperApproval)
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
func (it *RequestableERC20WrapperApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableERC20WrapperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableERC20WrapperApproval represents a Approval event raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RequestableERC20WrapperApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RequestableERC20Wrapper.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperApprovalIterator{contract: _RequestableERC20Wrapper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RequestableERC20WrapperApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RequestableERC20Wrapper.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableERC20WrapperApproval)
				if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Approval", log); err != nil {
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

// RequestableERC20WrapperDeposittedIterator is returned from FilterDepositted and is used to iterate over the raw logs and unpacked data for Depositted events raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperDeposittedIterator struct {
	Event *RequestableERC20WrapperDepositted // Event containing the contract specifics and raw log

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
func (it *RequestableERC20WrapperDeposittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableERC20WrapperDepositted)
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
		it.Event = new(RequestableERC20WrapperDepositted)
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
func (it *RequestableERC20WrapperDeposittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableERC20WrapperDeposittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableERC20WrapperDepositted represents a Depositted event raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperDepositted struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDepositted is a free log retrieval operation binding the contract event 0x46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4.
//
// Solidity: e Depositted(_from address, _value uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) FilterDepositted(opts *bind.FilterOpts) (*RequestableERC20WrapperDeposittedIterator, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.FilterLogs(opts, "Depositted")
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperDeposittedIterator{contract: _RequestableERC20Wrapper.contract, event: "Depositted", logs: logs, sub: sub}, nil
}

// WatchDepositted is a free log subscription operation binding the contract event 0x46fc0825df492a7a7f1fb1b690f6912ecb69575484fbb24be0f398ae70d97db4.
//
// Solidity: e Depositted(_from address, _value uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) WatchDepositted(opts *bind.WatchOpts, sink chan<- *RequestableERC20WrapperDepositted) (event.Subscription, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.WatchLogs(opts, "Depositted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableERC20WrapperDepositted)
				if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Depositted", log); err != nil {
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

// RequestableERC20WrapperRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperRequestCreatedIterator struct {
	Event *RequestableERC20WrapperRequestCreated // Event containing the contract specifics and raw log

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
func (it *RequestableERC20WrapperRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableERC20WrapperRequestCreated)
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
		it.Event = new(RequestableERC20WrapperRequestCreated)
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
func (it *RequestableERC20WrapperRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableERC20WrapperRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableERC20WrapperRequestCreated represents a RequestCreated event raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperRequestCreated struct {
	IsExit    bool
	Requestor common.Address
	TrieKey   [32]byte
	TrieValue [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x72507fe0ffa88ff97b08e473da2cc7f83aecabbf60819dd1972f8736c309474f.
//
// Solidity: e RequestCreated(_isExit bool, _requestor address, _trieKey bytes32, _trieValue bytes32)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*RequestableERC20WrapperRequestCreatedIterator, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperRequestCreatedIterator{contract: _RequestableERC20Wrapper.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x72507fe0ffa88ff97b08e473da2cc7f83aecabbf60819dd1972f8736c309474f.
//
// Solidity: e RequestCreated(_isExit bool, _requestor address, _trieKey bytes32, _trieValue bytes32)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *RequestableERC20WrapperRequestCreated) (event.Subscription, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableERC20WrapperRequestCreated)
				if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// RequestableERC20WrapperTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperTransferIterator struct {
	Event *RequestableERC20WrapperTransfer // Event containing the contract specifics and raw log

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
func (it *RequestableERC20WrapperTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableERC20WrapperTransfer)
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
		it.Event = new(RequestableERC20WrapperTransfer)
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
func (it *RequestableERC20WrapperTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableERC20WrapperTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableERC20WrapperTransfer represents a Transfer event raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RequestableERC20WrapperTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RequestableERC20Wrapper.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperTransferIterator{contract: _RequestableERC20Wrapper.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RequestableERC20WrapperTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RequestableERC20Wrapper.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableERC20WrapperTransfer)
				if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// RequestableERC20WrapperWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperWithdrawnIterator struct {
	Event *RequestableERC20WrapperWithdrawn // Event containing the contract specifics and raw log

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
func (it *RequestableERC20WrapperWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestableERC20WrapperWithdrawn)
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
		it.Event = new(RequestableERC20WrapperWithdrawn)
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
func (it *RequestableERC20WrapperWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestableERC20WrapperWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestableERC20WrapperWithdrawn represents a Withdrawn event raised by the RequestableERC20Wrapper contract.
type RequestableERC20WrapperWithdrawn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(_from address, _value uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*RequestableERC20WrapperWithdrawnIterator, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &RequestableERC20WrapperWithdrawnIterator{contract: _RequestableERC20Wrapper.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: e Withdrawn(_from address, _value uint256)
func (_RequestableERC20Wrapper *RequestableERC20WrapperFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *RequestableERC20WrapperWithdrawn) (event.Subscription, error) {

	logs, sub, err := _RequestableERC20Wrapper.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestableERC20WrapperWithdrawn)
				if err := _RequestableERC20Wrapper.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e168f936d37dcdf0a7ba8910ee64b8753af6aa6a636c1994872d38f04e8242810029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x608060405234801561001057600080fd5b506106ae806100206000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009257806318160ddd146100ca57806323b872dd146100f1578063661884631461011b57806370a082311461013f578063a9059cbb14610160578063d73dd62314610184578063dd62ed3e146101a8575b600080fd5b34801561009e57600080fd5b506100b6600160a060020a03600435166024356101cf565b604080519115158252519081900360200190f35b3480156100d657600080fd5b506100df610235565b60408051918252519081900360200190f35b3480156100fd57600080fd5b506100b6600160a060020a036004358116906024351660443561023b565b34801561012757600080fd5b506100b6600160a060020a03600435166024356103b0565b34801561014b57600080fd5b506100df600160a060020a036004351661049f565b34801561016c57600080fd5b506100b6600160a060020a03600435166024356104ba565b34801561019057600080fd5b506100b6600160a060020a0360043516602435610599565b3480156101b457600080fd5b506100df600160a060020a0360043581169060243516610632565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60015490565b600160a060020a03831660009081526020819052604081205482111561026057600080fd5b600160a060020a038416600090815260026020908152604080832033845290915290205482111561029057600080fd5b600160a060020a03831615156102a557600080fd5b600160a060020a0384166000908152602081905260409020546102ce908363ffffffff61065d16565b600160a060020a038086166000908152602081905260408082209390935590851681522054610303908363ffffffff61066f16565b600160a060020a03808516600090815260208181526040808320949094559187168152600282528281203382529091522054610345908363ffffffff61065d16565b600160a060020a03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b336000908152600260209081526040808320600160a060020a038616845290915281205480831061040457336000908152600260209081526040808320600160a060020a0388168452909152812055610439565b610414818463ffffffff61065d16565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526020819052604090205490565b336000908152602081905260408120548211156104d657600080fd5b600160a060020a03831615156104eb57600080fd5b3360009081526020819052604090205461050b908363ffffffff61065d16565b3360009081526020819052604080822092909255600160a060020a0385168152205461053d908363ffffffff61066f16565b600160a060020a038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120546105cd908363ffffffff61066f16565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561066957fe5b50900390565b8181018281101561067c57fe5b929150505600a165627a7a723058202ac26c0035a6c73bec3b6e8b444e960ee5cf2434b5738f7bdbda5ea3a77e3aed0029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
