// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package score

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

// ScoreMetaData contains all meta data concerning the Score contract.
var ScoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"ScoreUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"GetEpochRangeScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"}],\"name\":\"GetScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"}],\"name\":\"GetScoreData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"}],\"name\":\"IsValidatorRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"}],\"name\":\"RegisterValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"SetTargetValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"name\":\"UpdateScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockAggregates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"scores\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d1e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063640c65dd1161008c578063b7e3707f11610066578063b7e3707f14610227578063c4c8e16f14610257578063e580b56714610287578063ed612f8c146102b7576100cf565b8063640c65dd146101be5780637484182d146101ef5780639845e3b81461020b576100cf565b8063235d16ac146100d457806327258b22146100f25780634f0e1b081461012257806351351d531461013e5780635152c0581461015c5780635b36ac1c1461018d575b600080fd5b6100dc6102d5565b6040516100e991906107f7565b60405180910390f35b61010c6004803603810190610107919061084d565b6102db565b6040516101199190610895565b60405180910390f35b61013c600480360381019061013791906108dc565b6102fb565b005b610146610374565b604051610153919061094a565b60405180910390f35b610176600480360381019061017191906108dc565b610379565b604051610184929190610965565b60405180910390f35b6101a760048036038101906101a2919061084d565b61039d565b6040516101b5929190610965565b60405180910390f35b6101d860048036038101906101d3919061084d565b6103c1565b6040516101e6929190610965565b60405180910390f35b6102096004803603810190610204919061098e565b61040e565b005b6102256004803603810190610220919061084d565b610540565b005b610241600480360381019061023c91906109ce565b610680565b60405161024e91906107f7565b60405180910390f35b610271600480360381019061026c919061084d565b61072e565b60405161027e91906107f7565b60405180910390f35b6102a1600480360381019061029c919061084d565b6107ae565b6040516102ae9190610895565b60405180910390f35b6102bf6107d8565b6040516102cc91906107f7565b60405180910390f35b60005481565b60046020528060005260406000206000915054906101000a900460ff1681565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461036a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036190610a6b565b60405180910390fd5b8060008190555050565b600081565b60036020528060005260406000206000915090508060000154908060010154905082565b60026020528060005260406000206000915090508060000154908060010154905082565b600080600060026000858152602001908152602001600020604051806040016040529081600082015481526020016001820154815250509050806000015181602001519250925050915091565b60006002600084815260200190815260200160002090506000816000015490506000826001015490506000811461049d57600060036000838152602001908152602001600020905060008160010154111561049b57828160000160008282546104779190610aba565b9250508190555060018160010160008282546104939190610aba565b925050819055505b505b8383600001819055504383600101819055506000600360004381526020019081526020016000209050848160000160008282546104da9190610aee565b9250508190555060018160010160008282546104f69190610aee565b925050819055507f472b3d132ca0291b211798f7f566c888183f4fe746fa3170f08ec2213293e2a486864360405161053093929190610b31565b60405180910390a1505050505050565b600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a690610a6b565b60405180910390fd5b6004600082815260200190815260200160002060009054906101000a900460ff1615610610576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060790610bb4565b60405180910390fd5b60016004600083815260200190815260200160002060006101000a81548160ff021916908315150217905550604051806040016040528061019081526020014381525060026000838152602001908152602001600020600082015181600001556020820151816001015590505050565b600080600090506000808590505b848111610703576000600360008381526020019081526020016000206040518060400160405290816000820154815260200160018201548152505090508060000151846106db9190610aee565b93508060200151836106ed9190610aee565b92505080806106fb90610bd4565b91505061068e565b506000810361071757600092505050610728565b80826107239190610c4b565b925050505b92915050565b60006004600083815260200190815260200160002060009054906101000a900460ff16610790576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078790610cc8565b60405180910390fd5b60026000838152602001908152602001600020600001549050919050565b60006004600083815260200190815260200160002060009054906101000a900460ff169050919050565b60015481565b6000819050919050565b6107f1816107de565b82525050565b600060208201905061080c60008301846107e8565b92915050565b600080fd5b6000819050919050565b61082a81610817565b811461083557600080fd5b50565b60008135905061084781610821565b92915050565b60006020828403121561086357610862610812565b5b600061087184828501610838565b91505092915050565b60008115159050919050565b61088f8161087a565b82525050565b60006020820190506108aa6000830184610886565b92915050565b6108b9816107de565b81146108c457600080fd5b50565b6000813590506108d6816108b0565b92915050565b6000602082840312156108f2576108f1610812565b5b6000610900848285016108c7565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061093482610909565b9050919050565b61094481610929565b82525050565b600060208201905061095f600083018461093b565b92915050565b600060408201905061097a60008301856107e8565b61098760208301846107e8565b9392505050565b600080604083850312156109a5576109a4610812565b5b60006109b385828601610838565b92505060206109c4858286016108c7565b9150509250929050565b600080604083850312156109e5576109e4610812565b5b60006109f3858286016108c7565b9250506020610a04858286016108c7565b9150509250929050565b600082825260208201905092915050565b7f4f6e6c792073797374656d2063616e2063616c6c000000000000000000000000600082015250565b6000610a55601483610a0e565b9150610a6082610a1f565b602082019050919050565b60006020820190508181036000830152610a8481610a48565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ac5826107de565b9150610ad0836107de565b9250828203905081811115610ae857610ae7610a8b565b5b92915050565b6000610af9826107de565b9150610b04836107de565b9250828201905080821115610b1c57610b1b610a8b565b5b92915050565b610b2b81610817565b82525050565b6000606082019050610b466000830186610b22565b610b5360208301856107e8565b610b6060408301846107e8565b949350505050565b7f416c726561647920726567697374657265640000000000000000000000000000600082015250565b6000610b9e601283610a0e565b9150610ba982610b68565b602082019050919050565b60006020820190508181036000830152610bcd81610b91565b9050919050565b6000610bdf826107de565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c1157610c10610a8b565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610c56826107de565b9150610c61836107de565b925082610c7157610c70610c1c565b5b828204905092915050565b7f56616c696461746f72206e6f7420726567697374657265640000000000000000600082015250565b6000610cb2601883610a0e565b9150610cbd82610c7c565b602082019050919050565b60006020820190508181036000830152610ce181610ca5565b905091905056fea2646970667358221220efbe160bac875961ba6d95cd69294c64192cbdd12d335a6b9720331c5f596e8964736f6c63430008120033",
}

// ScoreABI is the input ABI used to generate the binding from.
// Deprecated: Use ScoreMetaData.ABI instead.
var ScoreABI = ScoreMetaData.ABI

// ScoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ScoreMetaData.Bin instead.
var ScoreBin = ScoreMetaData.Bin

// DeployScore deploys a new Ethereum contract, binding an instance of Score to it.
func DeployScore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Score, error) {
	parsed, err := ScoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Score{ScoreCaller: ScoreCaller{contract: contract}, ScoreTransactor: ScoreTransactor{contract: contract}, ScoreFilterer: ScoreFilterer{contract: contract}}, nil
}

// Score is an auto generated Go binding around an Ethereum contract.
type Score struct {
	ScoreCaller     // Read-only binding to the contract
	ScoreTransactor // Write-only binding to the contract
	ScoreFilterer   // Log filterer for contract events
}

// ScoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScoreSession struct {
	Contract     *Score            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScoreCallerSession struct {
	Contract *ScoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ScoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScoreTransactorSession struct {
	Contract     *ScoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScoreRaw struct {
	Contract *Score // Generic contract binding to access the raw methods on
}

// ScoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScoreCallerRaw struct {
	Contract *ScoreCaller // Generic read-only contract binding to access the raw methods on
}

// ScoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScoreTransactorRaw struct {
	Contract *ScoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScore creates a new instance of Score, bound to a specific deployed contract.
func NewScore(address common.Address, backend bind.ContractBackend) (*Score, error) {
	contract, err := bindScore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Score{ScoreCaller: ScoreCaller{contract: contract}, ScoreTransactor: ScoreTransactor{contract: contract}, ScoreFilterer: ScoreFilterer{contract: contract}}, nil
}

// NewScoreCaller creates a new read-only instance of Score, bound to a specific deployed contract.
func NewScoreCaller(address common.Address, caller bind.ContractCaller) (*ScoreCaller, error) {
	contract, err := bindScore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScoreCaller{contract: contract}, nil
}

// NewScoreTransactor creates a new write-only instance of Score, bound to a specific deployed contract.
func NewScoreTransactor(address common.Address, transactor bind.ContractTransactor) (*ScoreTransactor, error) {
	contract, err := bindScore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScoreTransactor{contract: contract}, nil
}

// NewScoreFilterer creates a new log filterer instance of Score, bound to a specific deployed contract.
func NewScoreFilterer(address common.Address, filterer bind.ContractFilterer) (*ScoreFilterer, error) {
	contract, err := bindScore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScoreFilterer{contract: contract}, nil
}

// bindScore binds a generic wrapper to an already deployed contract.
func bindScore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Score *ScoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Score.Contract.ScoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Score *ScoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Score.Contract.ScoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Score *ScoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Score.Contract.ScoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Score *ScoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Score.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Score *ScoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Score.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Score *ScoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Score.Contract.contract.Transact(opts, method, params...)
}

// GetEpochRangeScore is a free data retrieval call binding the contract method 0xb7e3707f.
//
// Solidity: function GetEpochRangeScore(uint256 startBlock, uint256 endBlock) view returns(uint256)
func (_Score *ScoreCaller) GetEpochRangeScore(opts *bind.CallOpts, startBlock *big.Int, endBlock *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Score.contract.Call(opts, &out, "GetEpochRangeScore", startBlock, endBlock)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochRangeScore is a free data retrieval call binding the contract method 0xb7e3707f.
//
// Solidity: function GetEpochRangeScore(uint256 startBlock, uint256 endBlock) view returns(uint256)
func (_Score *ScoreSession) GetEpochRangeScore(startBlock *big.Int, endBlock *big.Int) (*big.Int, error) {
	return _Score.Contract.GetEpochRangeScore(&_Score.CallOpts, startBlock, endBlock)
}

// GetEpochRangeScore is a free data retrieval call binding the contract method 0xb7e3707f.
//
// Solidity: function GetEpochRangeScore(uint256 startBlock, uint256 endBlock) view returns(uint256)
func (_Score *ScoreCallerSession) GetEpochRangeScore(startBlock *big.Int, endBlock *big.Int) (*big.Int, error) {
	return _Score.Contract.GetEpochRangeScore(&_Score.CallOpts, startBlock, endBlock)
}

// GetScore is a free data retrieval call binding the contract method 0xc4c8e16f.
//
// Solidity: function GetScore(bytes32 pubKey) view returns(uint256)
func (_Score *ScoreCaller) GetScore(opts *bind.CallOpts, pubKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Score.contract.Call(opts, &out, "GetScore", pubKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetScore is a free data retrieval call binding the contract method 0xc4c8e16f.
//
// Solidity: function GetScore(bytes32 pubKey) view returns(uint256)
func (_Score *ScoreSession) GetScore(pubKey [32]byte) (*big.Int, error) {
	return _Score.Contract.GetScore(&_Score.CallOpts, pubKey)
}

// GetScore is a free data retrieval call binding the contract method 0xc4c8e16f.
//
// Solidity: function GetScore(bytes32 pubKey) view returns(uint256)
func (_Score *ScoreCallerSession) GetScore(pubKey [32]byte) (*big.Int, error) {
	return _Score.Contract.GetScore(&_Score.CallOpts, pubKey)
}

// GetScoreData is a free data retrieval call binding the contract method 0x640c65dd.
//
// Solidity: function GetScoreData(bytes32 pubKey) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_Score *ScoreCaller) GetScoreData(opts *bind.CallOpts, pubKey [32]byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	var out []interface{}
	err := _Score.contract.Call(opts, &out, "GetScoreData", pubKey)

	outstruct := new(struct {
		Score            *big.Int
		LastUpdatedBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Score = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdatedBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetScoreData is a free data retrieval call binding the contract method 0x640c65dd.
//
// Solidity: function GetScoreData(bytes32 pubKey) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_Score *ScoreSession) GetScoreData(pubKey [32]byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	return _Score.Contract.GetScoreData(&_Score.CallOpts, pubKey)
}

// GetScoreData is a free data retrieval call binding the contract method 0x640c65dd.
//
// Solidity: function GetScoreData(bytes32 pubKey) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_Score *ScoreCallerSession) GetScoreData(pubKey [32]byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	return _Score.Contract.GetScoreData(&_Score.CallOpts, pubKey)
}

// IsValidatorRegistered is a free data retrieval call binding the contract method 0xe580b567.
//
// Solidity: function IsValidatorRegistered(bytes32 pubKey) view returns(bool)
func (_Score *ScoreCaller) IsValidatorRegistered(opts *bind.CallOpts, pubKey [32]byte) (bool, error) {
	var out []interface{}
	err := _Score.contract.Call(opts, &out, "IsValidatorRegistered", pubKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorRegistered is a free data retrieval call binding the contract method 0xe580b567.
//
// Solidity: function IsValidatorRegistered(bytes32 pubKey) view returns(bool)
func (_Score *ScoreSession) IsValidatorRegistered(pubKey [32]byte) (bool, error) {
	return _Score.Contract.IsValidatorRegistered(&_Score.CallOpts, pubKey)
}

// IsValidatorRegistered is a free data retrieval call binding the contract method 0xe580b567.
//
// Solidity: function IsValidatorRegistered(bytes32 pubKey) view returns(bool)
func (_Score *ScoreCallerSession) IsValidatorRegistered(pubKey [32]byte) (bool, error) {
	return _Score.Contract.IsValidatorRegistered(&_Score.CallOpts, pubKey)
}

// SYSTEM is a free data retrieval call binding the contract method 0x51351d53.
//
// Solidity: function SYSTEM() view returns(address)
func (_Score *ScoreCaller) SYSTEM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Score.contract.Call(opts, &out, "SYSTEM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEM is a free data retrieval call binding the contract method 0x51351d53.
//
// Solidity: function SYSTEM() view returns(address)
func (_Score *ScoreSession) SYSTEM() (common.Address, error) {
	return _Score.Contract.SYSTEM(&_Score.CallOpts)
}

// SYSTEM is a free data retrieval call binding the contract method 0x51351d53.
//
// Solidity: function SYSTEM() view returns(address)
func (_Score *ScoreCallerSession) SYSTEM() (common.Address, error) {
	return _Score.Contract.SYSTEM(&_Score.CallOpts)
}

// BlockAggregates is a free data retrieval call binding the contract method 0x5152c058.
//
// Solidity: function blockAggregates(uint256 ) view returns(uint256 totalScore, uint256 count)
func (_Score *ScoreCaller) BlockAggregates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TotalScore *big.Int
	Count      *big.Int
}, error) {
	var out []interface{}
	err := _Score.contract.Call(opts, &out, "blockAggregates", arg0)

	outstruct := new(struct {
		TotalScore *big.Int
		Count      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalScore = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Count = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BlockAggregates is a free data retrieval call binding the contract method 0x5152c058.
//
// Solidity: function blockAggregates(uint256 ) view returns(uint256 totalScore, uint256 count)
func (_Score *ScoreSession) BlockAggregates(arg0 *big.Int) (struct {
	TotalScore *big.Int
	Count      *big.Int
}, error) {
	return _Score.Contract.BlockAggregates(&_Score.CallOpts, arg0)
}

// BlockAggregates is a free data retrieval call binding the contract method 0x5152c058.
//
// Solidity: function blockAggregates(uint256 ) view returns(uint256 totalScore, uint256 count)
func (_Score *ScoreCallerSession) BlockAggregates(arg0 *big.Int) (struct {
	TotalScore *big.Int
	Count      *big.Int
}, error) {
	return _Score.Contract.BlockAggregates(&_Score.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 ) view returns(bool)
func (_Score *ScoreCaller) IsRegistered(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Score.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 ) view returns(bool)
func (_Score *ScoreSession) IsRegistered(arg0 [32]byte) (bool, error) {
	return _Score.Contract.IsRegistered(&_Score.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0x27258b22.
//
// Solidity: function isRegistered(bytes32 ) view returns(bool)
func (_Score *ScoreCallerSession) IsRegistered(arg0 [32]byte) (bool, error) {
	return _Score.Contract.IsRegistered(&_Score.CallOpts, arg0)
}

// Scores is a free data retrieval call binding the contract method 0x5b36ac1c.
//
// Solidity: function scores(bytes32 ) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_Score *ScoreCaller) Scores(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	var out []interface{}
	err := _Score.contract.Call(opts, &out, "scores", arg0)

	outstruct := new(struct {
		Score            *big.Int
		LastUpdatedBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Score = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdatedBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Scores is a free data retrieval call binding the contract method 0x5b36ac1c.
//
// Solidity: function scores(bytes32 ) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_Score *ScoreSession) Scores(arg0 [32]byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	return _Score.Contract.Scores(&_Score.CallOpts, arg0)
}

// Scores is a free data retrieval call binding the contract method 0x5b36ac1c.
//
// Solidity: function scores(bytes32 ) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_Score *ScoreCallerSession) Scores(arg0 [32]byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	return _Score.Contract.Scores(&_Score.CallOpts, arg0)
}

// TargetValidatorsCount is a free data retrieval call binding the contract method 0x235d16ac.
//
// Solidity: function targetValidatorsCount() view returns(uint256)
func (_Score *ScoreCaller) TargetValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Score.contract.Call(opts, &out, "targetValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetValidatorsCount is a free data retrieval call binding the contract method 0x235d16ac.
//
// Solidity: function targetValidatorsCount() view returns(uint256)
func (_Score *ScoreSession) TargetValidatorsCount() (*big.Int, error) {
	return _Score.Contract.TargetValidatorsCount(&_Score.CallOpts)
}

// TargetValidatorsCount is a free data retrieval call binding the contract method 0x235d16ac.
//
// Solidity: function targetValidatorsCount() view returns(uint256)
func (_Score *ScoreCallerSession) TargetValidatorsCount() (*big.Int, error) {
	return _Score.Contract.TargetValidatorsCount(&_Score.CallOpts)
}

// ValidatorsCount is a free data retrieval call binding the contract method 0xed612f8c.
//
// Solidity: function validatorsCount() view returns(uint256)
func (_Score *ScoreCaller) ValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Score.contract.Call(opts, &out, "validatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorsCount is a free data retrieval call binding the contract method 0xed612f8c.
//
// Solidity: function validatorsCount() view returns(uint256)
func (_Score *ScoreSession) ValidatorsCount() (*big.Int, error) {
	return _Score.Contract.ValidatorsCount(&_Score.CallOpts)
}

// ValidatorsCount is a free data retrieval call binding the contract method 0xed612f8c.
//
// Solidity: function validatorsCount() view returns(uint256)
func (_Score *ScoreCallerSession) ValidatorsCount() (*big.Int, error) {
	return _Score.Contract.ValidatorsCount(&_Score.CallOpts)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x9845e3b8.
//
// Solidity: function RegisterValidator(bytes32 pubKey) returns()
func (_Score *ScoreTransactor) RegisterValidator(opts *bind.TransactOpts, pubKey [32]byte) (*types.Transaction, error) {
	return _Score.contract.Transact(opts, "RegisterValidator", pubKey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x9845e3b8.
//
// Solidity: function RegisterValidator(bytes32 pubKey) returns()
func (_Score *ScoreSession) RegisterValidator(pubKey [32]byte) (*types.Transaction, error) {
	return _Score.Contract.RegisterValidator(&_Score.TransactOpts, pubKey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x9845e3b8.
//
// Solidity: function RegisterValidator(bytes32 pubKey) returns()
func (_Score *ScoreTransactorSession) RegisterValidator(pubKey [32]byte) (*types.Transaction, error) {
	return _Score.Contract.RegisterValidator(&_Score.TransactOpts, pubKey)
}

// SetTargetValidatorsCount is a paid mutator transaction binding the contract method 0x4f0e1b08.
//
// Solidity: function SetTargetValidatorsCount(uint256 _count) returns()
func (_Score *ScoreTransactor) SetTargetValidatorsCount(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _Score.contract.Transact(opts, "SetTargetValidatorsCount", _count)
}

// SetTargetValidatorsCount is a paid mutator transaction binding the contract method 0x4f0e1b08.
//
// Solidity: function SetTargetValidatorsCount(uint256 _count) returns()
func (_Score *ScoreSession) SetTargetValidatorsCount(_count *big.Int) (*types.Transaction, error) {
	return _Score.Contract.SetTargetValidatorsCount(&_Score.TransactOpts, _count)
}

// SetTargetValidatorsCount is a paid mutator transaction binding the contract method 0x4f0e1b08.
//
// Solidity: function SetTargetValidatorsCount(uint256 _count) returns()
func (_Score *ScoreTransactorSession) SetTargetValidatorsCount(_count *big.Int) (*types.Transaction, error) {
	return _Score.Contract.SetTargetValidatorsCount(&_Score.TransactOpts, _count)
}

// UpdateScore is a paid mutator transaction binding the contract method 0x7484182d.
//
// Solidity: function UpdateScore(bytes32 pubKey, uint256 score) returns()
func (_Score *ScoreTransactor) UpdateScore(opts *bind.TransactOpts, pubKey [32]byte, score *big.Int) (*types.Transaction, error) {
	return _Score.contract.Transact(opts, "UpdateScore", pubKey, score)
}

// UpdateScore is a paid mutator transaction binding the contract method 0x7484182d.
//
// Solidity: function UpdateScore(bytes32 pubKey, uint256 score) returns()
func (_Score *ScoreSession) UpdateScore(pubKey [32]byte, score *big.Int) (*types.Transaction, error) {
	return _Score.Contract.UpdateScore(&_Score.TransactOpts, pubKey, score)
}

// UpdateScore is a paid mutator transaction binding the contract method 0x7484182d.
//
// Solidity: function UpdateScore(bytes32 pubKey, uint256 score) returns()
func (_Score *ScoreTransactorSession) UpdateScore(pubKey [32]byte, score *big.Int) (*types.Transaction, error) {
	return _Score.Contract.UpdateScore(&_Score.TransactOpts, pubKey, score)
}

// ScoreScoreUpdatedIterator is returned from FilterScoreUpdated and is used to iterate over the raw logs and unpacked data for ScoreUpdated events raised by the Score contract.
type ScoreScoreUpdatedIterator struct {
	Event *ScoreScoreUpdated // Event containing the contract specifics and raw log

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
func (it *ScoreScoreUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoreScoreUpdated)
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
		it.Event = new(ScoreScoreUpdated)
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
func (it *ScoreScoreUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoreScoreUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoreScoreUpdated represents a ScoreUpdated event raised by the Score contract.
type ScoreScoreUpdated struct {
	PubKey      [32]byte
	Score       *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterScoreUpdated is a free log retrieval operation binding the contract event 0x472b3d132ca0291b211798f7f566c888183f4fe746fa3170f08ec2213293e2a4.
//
// Solidity: event ScoreUpdated(bytes32 pubKey, uint256 score, uint256 blockNumber)
func (_Score *ScoreFilterer) FilterScoreUpdated(opts *bind.FilterOpts) (*ScoreScoreUpdatedIterator, error) {

	logs, sub, err := _Score.contract.FilterLogs(opts, "ScoreUpdated")
	if err != nil {
		return nil, err
	}
	return &ScoreScoreUpdatedIterator{contract: _Score.contract, event: "ScoreUpdated", logs: logs, sub: sub}, nil
}

// WatchScoreUpdated is a free log subscription operation binding the contract event 0x472b3d132ca0291b211798f7f566c888183f4fe746fa3170f08ec2213293e2a4.
//
// Solidity: event ScoreUpdated(bytes32 pubKey, uint256 score, uint256 blockNumber)
func (_Score *ScoreFilterer) WatchScoreUpdated(opts *bind.WatchOpts, sink chan<- *ScoreScoreUpdated) (event.Subscription, error) {

	logs, sub, err := _Score.contract.WatchLogs(opts, "ScoreUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoreScoreUpdated)
				if err := _Score.contract.UnpackLog(event, "ScoreUpdated", log); err != nil {
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

// ParseScoreUpdated is a log parse operation binding the contract event 0x472b3d132ca0291b211798f7f566c888183f4fe746fa3170f08ec2213293e2a4.
//
// Solidity: event ScoreUpdated(bytes32 pubKey, uint256 score, uint256 blockNumber)
func (_Score *ScoreFilterer) ParseScoreUpdated(log types.Log) (*ScoreScoreUpdated, error) {
	event := new(ScoreScoreUpdated)
	if err := _Score.contract.UnpackLog(event, "ScoreUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScoreValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the Score contract.
type ScoreValidatorRegisteredIterator struct {
	Event *ScoreValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *ScoreValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoreValidatorRegistered)
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
		it.Event = new(ScoreValidatorRegistered)
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
func (it *ScoreValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoreValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoreValidatorRegistered represents a ValidatorRegistered event raised by the Score contract.
type ScoreValidatorRegistered struct {
	PubKey [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0xcb2f5bc9a8ab981f7104d49f5ff0c41a092b6479a6ae764a389bab934a177510.
//
// Solidity: event ValidatorRegistered(bytes32 pubKey)
func (_Score *ScoreFilterer) FilterValidatorRegistered(opts *bind.FilterOpts) (*ScoreValidatorRegisteredIterator, error) {

	logs, sub, err := _Score.contract.FilterLogs(opts, "ValidatorRegistered")
	if err != nil {
		return nil, err
	}
	return &ScoreValidatorRegisteredIterator{contract: _Score.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0xcb2f5bc9a8ab981f7104d49f5ff0c41a092b6479a6ae764a389bab934a177510.
//
// Solidity: event ValidatorRegistered(bytes32 pubKey)
func (_Score *ScoreFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *ScoreValidatorRegistered) (event.Subscription, error) {

	logs, sub, err := _Score.contract.WatchLogs(opts, "ValidatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoreValidatorRegistered)
				if err := _Score.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0xcb2f5bc9a8ab981f7104d49f5ff0c41a092b6479a6ae764a389bab934a177510.
//
// Solidity: event ValidatorRegistered(bytes32 pubKey)
func (_Score *ScoreFilterer) ParseValidatorRegistered(log types.Log) (*ScoreValidatorRegistered, error) {
	event := new(ScoreValidatorRegistered)
	if err := _Score.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
