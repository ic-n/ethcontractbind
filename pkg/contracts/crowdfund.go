// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// CrowdFundMetaData contains all meta data concerning the CrowdFund contract.
var CrowdFundMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startAt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"endAt\",\"type\":\"uint32\"}],\"name\":\"Launch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Pledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unpledge\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaigns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pledged\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endAt\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_goal\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_startAt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_endAt\",\"type\":\"uint32\"}],\"name\":\"launch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pledgedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unpledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001cf138038062001cf18339818101604052810190620000379190620000dc565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506200010e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000a48262000077565b9050919050565b620000b68162000097565b8114620000c257600080fd5b50565b600081519050620000d681620000ab565b92915050565b600060208284031215620000f557620000f462000072565b5b60006200010584828501620000c5565b91505092915050565b608051611bab62000146600039600081816104b1015281816109fa01528181610e6d01528181610f8501526110fc0152611bab6000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806340e58ee51161006657806340e58ee51461014a578063711853ab14610166578063aa4fb63a14610182578063fc0c546a146101b2578063fde327be146101d05761009e565b806306661abd146100a3578063141961bc146100c1578063278ecde1146100f65780632c63f14614610112578063379607f51461012e575b600080fd5b6100ab6101ec565b6040516100b89190611208565b60405180910390f35b6100db60048036038101906100d69190611254565b6101f2565b6040516100ed969594939291906112fc565b60405180910390f35b610110600480360381019061010b9190611254565b61027b565b005b61012c60048036038101906101279190611389565b6105a3565b005b61014860048036038101906101439190611254565b610839565b005b610164600480360381019061015f9190611254565b610afa565b005b610180600480360381019061017b91906113dc565b610d75565b005b61019c60048036038101906101979190611448565b610f5e565b6040516101a99190611208565b60405180910390f35b6101ba610f83565b6040516101c791906114e7565b60405180910390f35b6101ea60048036038101906101e591906113dc565b610fa7565b005b60005481565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030160009054906101000a900463ffffffff16908060030160049054906101000a900463ffffffff16908060030160089054906101000a900460ff16905086565b6000600160008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481526020016003820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016003820160049054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016003820160089054906101000a900460ff1615151515815250509050806080015163ffffffff1642116103ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b19061155f565b60405180910390fd5b8060200151816040015110610404576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103fb906115cb565b60405180910390fd5b60006002600084815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060006002600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b815260040161050a9291906115eb565b6020604051808303816000875af1158015610529573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054d9190611640565b503373ffffffffffffffffffffffffffffffffffffffff167f21e12a7cad0da5928167e1084ea4d5fdf8d9af66657a2543a9ac76a0ca081477848360405161059692919061166d565b60405180910390a2505050565b428263ffffffff1610156105ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e3906116e2565b60405180910390fd5b8163ffffffff168163ffffffff16101561063b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106329061174e565b60405180910390fd5b6276a7004261064a919061179d565b8163ffffffff161115610692576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106899061181d565b60405180910390fd5b60016000808282546106a4919061179d565b925050819055506040518060c001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001848152602001600081526020018363ffffffff1681526020018263ffffffff16815260200160001515815250600160008054815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030160006101000a81548163ffffffff021916908363ffffffff16021790555060808201518160030160046101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160030160086101000a81548160ff0219169083151502179055509050503373ffffffffffffffffffffffffffffffffffffffff167f0601cd0d650b473037e838a2696d41e654433d065b3f56b28d1d3302e44a304f60005485858560405161082c949392919061183d565b60405180910390a2505050565b60006001600083815260200190815260200160002090503373ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146108e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d9906118ce565b60405180910390fd5b8060030160049054906101000a900463ffffffff1663ffffffff16421161093e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109359061155f565b60405180910390fd5b806001015481600201541015610989576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109809061193a565b60405180910390fd5b8060030160089054906101000a900460ff16156109db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d2906119a6565b60405180910390fd5b60018160030160086101000a81548160ff0219169083151502179055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600201546040518363ffffffff1660e01b8152600401610a7b9291906115eb565b6020604051808303816000875af1158015610a9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610abe9190611640565b507f7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f782604051610aee9190611208565b60405180910390a15050565b6000600160008381526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282015481526020016003820160009054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016003820160049054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016003820160089054906101000a900460ff16151515158152505090503373ffffffffffffffffffffffffffffffffffffffff16816000015173ffffffffffffffffffffffffffffffffffffffff1614610c5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c56906118ce565b60405180910390fd5b806060015163ffffffff164210610cab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ca290611a12565b60405180910390fd5b60016000838152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182016000905560028201600090556003820160006101000a81549063ffffffff02191690556003820160046101000a81549063ffffffff02191690556003820160086101000a81549060ff021916905550507f8bf30e7ff26833413be5f69e1d373744864d600b664204b4a2f9844a8eedb9ed82604051610d699190611208565b60405180910390a15050565b60006001600084815260200190815260200160002090508060030160049054906101000a900463ffffffff1663ffffffff16421115610de9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de090611a7e565b60405180910390fd5b81816002016000828254610dfd9190611a9e565b92505081905550816002600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e649190611a9e565b925050819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b8152600401610ec69291906115eb565b6020604051808303816000875af1158015610ee5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f099190611640565b503373ffffffffffffffffffffffffffffffffffffffff16837f2eeeab891b26a214d1b25749f88a406bdea852bd8c9bfda977e0ef8114c180ba84604051610f519190611208565b60405180910390a3505050565b6002602052816000526040600020602052806000526040600020600091509150505481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006001600084815260200190815260200160002090508060030160009054906101000a900463ffffffff1663ffffffff1642101561101b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101290611b1e565b60405180910390fd5b8060030160049054906101000a900463ffffffff1663ffffffff16421115611078576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161106f90611a7e565b60405180910390fd5b8181600201600082825461108c919061179d565b92505081905550816002600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546110f3919061179d565b925050819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161115793929190611b3e565b6020604051808303816000875af1158015611176573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061119a9190611640565b503373ffffffffffffffffffffffffffffffffffffffff16837f06bdb975df800a73232998e71ed585d536222f1dfeaa622d7f62a23ada686c82846040516111e29190611208565b60405180910390a3505050565b6000819050919050565b611202816111ef565b82525050565b600060208201905061121d60008301846111f9565b92915050565b600080fd5b611231816111ef565b811461123c57600080fd5b50565b60008135905061124e81611228565b92915050565b60006020828403121561126a57611269611223565b5b60006112788482850161123f565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006112ac82611281565b9050919050565b6112bc816112a1565b82525050565b600063ffffffff82169050919050565b6112db816112c2565b82525050565b60008115159050919050565b6112f6816112e1565b82525050565b600060c08201905061131160008301896112b3565b61131e60208301886111f9565b61132b60408301876111f9565b61133860608301866112d2565b61134560808301856112d2565b61135260a08301846112ed565b979650505050505050565b611366816112c2565b811461137157600080fd5b50565b6000813590506113838161135d565b92915050565b6000806000606084860312156113a2576113a1611223565b5b60006113b08682870161123f565b93505060206113c186828701611374565b92505060406113d286828701611374565b9150509250925092565b600080604083850312156113f3576113f2611223565b5b60006114018582860161123f565b92505060206114128582860161123f565b9150509250929050565b611425816112a1565b811461143057600080fd5b50565b6000813590506114428161141c565b92915050565b6000806040838503121561145f5761145e611223565b5b600061146d8582860161123f565b925050602061147e85828601611433565b9150509250929050565b6000819050919050565b60006114ad6114a86114a384611281565b611488565b611281565b9050919050565b60006114bf82611492565b9050919050565b60006114d1826114b4565b9050919050565b6114e1816114c6565b82525050565b60006020820190506114fc60008301846114d8565b92915050565b600082825260208201905092915050565b7f6e6f7420656e6465640000000000000000000000000000000000000000000000600082015250565b6000611549600983611502565b915061155482611513565b602082019050919050565b600060208201905081810360008301526115788161153c565b9050919050565b7f706c6564676564203e3d20676f616c0000000000000000000000000000000000600082015250565b60006115b5600f83611502565b91506115c08261157f565b602082019050919050565b600060208201905081810360008301526115e4816115a8565b9050919050565b600060408201905061160060008301856112b3565b61160d60208301846111f9565b9392505050565b61161d816112e1565b811461162857600080fd5b50565b60008151905061163a81611614565b92915050565b60006020828403121561165657611655611223565b5b60006116648482850161162b565b91505092915050565b600060408201905061168260008301856111f9565b61168f60208301846111f9565b9392505050565b7f7374617274206174203c206e6f77000000000000000000000000000000000000600082015250565b60006116cc600e83611502565b91506116d782611696565b602082019050919050565b600060208201905081810360008301526116fb816116bf565b9050919050565b7f656e64206174203c207374617274206174000000000000000000000000000000600082015250565b6000611738601183611502565b915061174382611702565b602082019050919050565b600060208201905081810360008301526117678161172b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006117a8826111ef565b91506117b3836111ef565b92508282019050808211156117cb576117ca61176e565b5b92915050565b7f656e64206174203e206d6178206475726174696f6e0000000000000000000000600082015250565b6000611807601583611502565b9150611812826117d1565b602082019050919050565b60006020820190508181036000830152611836816117fa565b9050919050565b600060808201905061185260008301876111f9565b61185f60208301866111f9565b61186c60408301856112d2565b61187960608301846112d2565b95945050505050565b7f6e6f742063726561746f72000000000000000000000000000000000000000000600082015250565b60006118b8600b83611502565b91506118c382611882565b602082019050919050565b600060208201905081810360008301526118e7816118ab565b9050919050565b7f706c6564676564203c20676f616c000000000000000000000000000000000000600082015250565b6000611924600e83611502565b915061192f826118ee565b602082019050919050565b6000602082019050818103600083015261195381611917565b9050919050565b7f636c61696d656400000000000000000000000000000000000000000000000000600082015250565b6000611990600783611502565b915061199b8261195a565b602082019050919050565b600060208201905081810360008301526119bf81611983565b9050919050565b7f7374617274656400000000000000000000000000000000000000000000000000600082015250565b60006119fc600783611502565b9150611a07826119c6565b602082019050919050565b60006020820190508181036000830152611a2b816119ef565b9050919050565b7f656e646564000000000000000000000000000000000000000000000000000000600082015250565b6000611a68600583611502565b9150611a7382611a32565b602082019050919050565b60006020820190508181036000830152611a9781611a5b565b9050919050565b6000611aa9826111ef565b9150611ab4836111ef565b9250828203905081811115611acc57611acb61176e565b5b92915050565b7f6e6f742073746172746564000000000000000000000000000000000000000000600082015250565b6000611b08600b83611502565b9150611b1382611ad2565b602082019050919050565b60006020820190508181036000830152611b3781611afb565b9050919050565b6000606082019050611b5360008301866112b3565b611b6060208301856112b3565b611b6d60408301846111f9565b94935050505056fea26469706673582212201d32f89bea74053a80e132181f99ba71b78d81b177e92890e5d988b7d0c7877664736f6c63430008170033",
}

// CrowdFundABI is the input ABI used to generate the binding from.
// Deprecated: Use CrowdFundMetaData.ABI instead.
var CrowdFundABI = CrowdFundMetaData.ABI

// CrowdFundBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrowdFundMetaData.Bin instead.
var CrowdFundBin = CrowdFundMetaData.Bin

// DeployCrowdFund deploys a new Ethereum contract, binding an instance of CrowdFund to it.
func DeployCrowdFund(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *CrowdFund, error) {
	parsed, err := CrowdFundMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrowdFundBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrowdFund{CrowdFundCaller: CrowdFundCaller{contract: contract}, CrowdFundTransactor: CrowdFundTransactor{contract: contract}, CrowdFundFilterer: CrowdFundFilterer{contract: contract}}, nil
}

// CrowdFund is an auto generated Go binding around an Ethereum contract.
type CrowdFund struct {
	CrowdFundCaller     // Read-only binding to the contract
	CrowdFundTransactor // Write-only binding to the contract
	CrowdFundFilterer   // Log filterer for contract events
}

// CrowdFundCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdFundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdFundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdFundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdFundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrowdFundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdFundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdFundSession struct {
	Contract     *CrowdFund        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdFundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdFundCallerSession struct {
	Contract *CrowdFundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrowdFundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdFundTransactorSession struct {
	Contract     *CrowdFundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrowdFundRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdFundRaw struct {
	Contract *CrowdFund // Generic contract binding to access the raw methods on
}

// CrowdFundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdFundCallerRaw struct {
	Contract *CrowdFundCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdFundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdFundTransactorRaw struct {
	Contract *CrowdFundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdFund creates a new instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFund(address common.Address, backend bind.ContractBackend) (*CrowdFund, error) {
	contract, err := bindCrowdFund(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrowdFund{CrowdFundCaller: CrowdFundCaller{contract: contract}, CrowdFundTransactor: CrowdFundTransactor{contract: contract}, CrowdFundFilterer: CrowdFundFilterer{contract: contract}}, nil
}

// NewCrowdFundCaller creates a new read-only instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFundCaller(address common.Address, caller bind.ContractCaller) (*CrowdFundCaller, error) {
	contract, err := bindCrowdFund(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdFundCaller{contract: contract}, nil
}

// NewCrowdFundTransactor creates a new write-only instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFundTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdFundTransactor, error) {
	contract, err := bindCrowdFund(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdFundTransactor{contract: contract}, nil
}

// NewCrowdFundFilterer creates a new log filterer instance of CrowdFund, bound to a specific deployed contract.
func NewCrowdFundFilterer(address common.Address, filterer bind.ContractFilterer) (*CrowdFundFilterer, error) {
	contract, err := bindCrowdFund(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrowdFundFilterer{contract: contract}, nil
}

// bindCrowdFund binds a generic wrapper to an already deployed contract.
func bindCrowdFund(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrowdFundMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdFund *CrowdFundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdFund.Contract.CrowdFundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdFund *CrowdFundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdFund.Contract.CrowdFundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdFund *CrowdFundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdFund.Contract.CrowdFundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdFund *CrowdFundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdFund.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdFund *CrowdFundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdFund.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdFund *CrowdFundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdFund.Contract.contract.Transact(opts, method, params...)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address creator, uint256 goal, uint256 pledged, uint32 startAt, uint32 endAt, bool claimed)
func (_CrowdFund *CrowdFundCaller) Campaigns(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Creator common.Address
	Goal    *big.Int
	Pledged *big.Int
	StartAt uint32
	EndAt   uint32
	Claimed bool
}, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "campaigns", arg0)

	outstruct := new(struct {
		Creator common.Address
		Goal    *big.Int
		Pledged *big.Int
		StartAt uint32
		EndAt   uint32
		Claimed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Creator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Goal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Pledged = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartAt = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.EndAt = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.Claimed = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address creator, uint256 goal, uint256 pledged, uint32 startAt, uint32 endAt, bool claimed)
func (_CrowdFund *CrowdFundSession) Campaigns(arg0 *big.Int) (struct {
	Creator common.Address
	Goal    *big.Int
	Pledged *big.Int
	StartAt uint32
	EndAt   uint32
	Claimed bool
}, error) {
	return _CrowdFund.Contract.Campaigns(&_CrowdFund.CallOpts, arg0)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address creator, uint256 goal, uint256 pledged, uint32 startAt, uint32 endAt, bool claimed)
func (_CrowdFund *CrowdFundCallerSession) Campaigns(arg0 *big.Int) (struct {
	Creator common.Address
	Goal    *big.Int
	Pledged *big.Int
	StartAt uint32
	EndAt   uint32
	Claimed bool
}, error) {
	return _CrowdFund.Contract.Campaigns(&_CrowdFund.CallOpts, arg0)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CrowdFund *CrowdFundCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CrowdFund *CrowdFundSession) Count() (*big.Int, error) {
	return _CrowdFund.Contract.Count(&_CrowdFund.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CrowdFund *CrowdFundCallerSession) Count() (*big.Int, error) {
	return _CrowdFund.Contract.Count(&_CrowdFund.CallOpts)
}

// PledgedAmount is a free data retrieval call binding the contract method 0xaa4fb63a.
//
// Solidity: function pledgedAmount(uint256 , address ) view returns(uint256)
func (_CrowdFund *CrowdFundCaller) PledgedAmount(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "pledgedAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgedAmount is a free data retrieval call binding the contract method 0xaa4fb63a.
//
// Solidity: function pledgedAmount(uint256 , address ) view returns(uint256)
func (_CrowdFund *CrowdFundSession) PledgedAmount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _CrowdFund.Contract.PledgedAmount(&_CrowdFund.CallOpts, arg0, arg1)
}

// PledgedAmount is a free data retrieval call binding the contract method 0xaa4fb63a.
//
// Solidity: function pledgedAmount(uint256 , address ) view returns(uint256)
func (_CrowdFund *CrowdFundCallerSession) PledgedAmount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _CrowdFund.Contract.PledgedAmount(&_CrowdFund.CallOpts, arg0, arg1)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdFund *CrowdFundCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrowdFund.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdFund *CrowdFundSession) Token() (common.Address, error) {
	return _CrowdFund.Contract.Token(&_CrowdFund.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdFund *CrowdFundCallerSession) Token() (common.Address, error) {
	return _CrowdFund.Contract.Token(&_CrowdFund.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactor) Cancel(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "cancel", _id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _id) returns()
func (_CrowdFund *CrowdFundSession) Cancel(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Cancel(&_CrowdFund.TransactOpts, _id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactorSession) Cancel(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Cancel(&_CrowdFund.TransactOpts, _id)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactor) Claim(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "claim", _id)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _id) returns()
func (_CrowdFund *CrowdFundSession) Claim(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Claim(&_CrowdFund.TransactOpts, _id)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactorSession) Claim(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Claim(&_CrowdFund.TransactOpts, _id)
}

// Launch is a paid mutator transaction binding the contract method 0x2c63f146.
//
// Solidity: function launch(uint256 _goal, uint32 _startAt, uint32 _endAt) returns()
func (_CrowdFund *CrowdFundTransactor) Launch(opts *bind.TransactOpts, _goal *big.Int, _startAt uint32, _endAt uint32) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "launch", _goal, _startAt, _endAt)
}

// Launch is a paid mutator transaction binding the contract method 0x2c63f146.
//
// Solidity: function launch(uint256 _goal, uint32 _startAt, uint32 _endAt) returns()
func (_CrowdFund *CrowdFundSession) Launch(_goal *big.Int, _startAt uint32, _endAt uint32) (*types.Transaction, error) {
	return _CrowdFund.Contract.Launch(&_CrowdFund.TransactOpts, _goal, _startAt, _endAt)
}

// Launch is a paid mutator transaction binding the contract method 0x2c63f146.
//
// Solidity: function launch(uint256 _goal, uint32 _startAt, uint32 _endAt) returns()
func (_CrowdFund *CrowdFundTransactorSession) Launch(_goal *big.Int, _startAt uint32, _endAt uint32) (*types.Transaction, error) {
	return _CrowdFund.Contract.Launch(&_CrowdFund.TransactOpts, _goal, _startAt, _endAt)
}

// Pledge is a paid mutator transaction binding the contract method 0xfde327be.
//
// Solidity: function pledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundTransactor) Pledge(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "pledge", _id, _amount)
}

// Pledge is a paid mutator transaction binding the contract method 0xfde327be.
//
// Solidity: function pledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundSession) Pledge(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Pledge(&_CrowdFund.TransactOpts, _id, _amount)
}

// Pledge is a paid mutator transaction binding the contract method 0xfde327be.
//
// Solidity: function pledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundTransactorSession) Pledge(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Pledge(&_CrowdFund.TransactOpts, _id, _amount)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactor) Refund(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "refund", _id)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 _id) returns()
func (_CrowdFund *CrowdFundSession) Refund(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Refund(&_CrowdFund.TransactOpts, _id)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 _id) returns()
func (_CrowdFund *CrowdFundTransactorSession) Refund(_id *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Refund(&_CrowdFund.TransactOpts, _id)
}

// Unpledge is a paid mutator transaction binding the contract method 0x711853ab.
//
// Solidity: function unpledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundTransactor) Unpledge(opts *bind.TransactOpts, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.contract.Transact(opts, "unpledge", _id, _amount)
}

// Unpledge is a paid mutator transaction binding the contract method 0x711853ab.
//
// Solidity: function unpledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundSession) Unpledge(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Unpledge(&_CrowdFund.TransactOpts, _id, _amount)
}

// Unpledge is a paid mutator transaction binding the contract method 0x711853ab.
//
// Solidity: function unpledge(uint256 _id, uint256 _amount) returns()
func (_CrowdFund *CrowdFundTransactorSession) Unpledge(_id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _CrowdFund.Contract.Unpledge(&_CrowdFund.TransactOpts, _id, _amount)
}

// CrowdFundCancelIterator is returned from FilterCancel and is used to iterate over the raw logs and unpacked data for Cancel events raised by the CrowdFund contract.
type CrowdFundCancelIterator struct {
	Event *CrowdFundCancel // Event containing the contract specifics and raw log

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
func (it *CrowdFundCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundCancel)
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
		it.Event = new(CrowdFundCancel)
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
func (it *CrowdFundCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundCancel represents a Cancel event raised by the CrowdFund contract.
type CrowdFundCancel struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancel is a free log retrieval operation binding the contract event 0x8bf30e7ff26833413be5f69e1d373744864d600b664204b4a2f9844a8eedb9ed.
//
// Solidity: event Cancel(uint256 id)
func (_CrowdFund *CrowdFundFilterer) FilterCancel(opts *bind.FilterOpts) (*CrowdFundCancelIterator, error) {

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Cancel")
	if err != nil {
		return nil, err
	}
	return &CrowdFundCancelIterator{contract: _CrowdFund.contract, event: "Cancel", logs: logs, sub: sub}, nil
}

// WatchCancel is a free log subscription operation binding the contract event 0x8bf30e7ff26833413be5f69e1d373744864d600b664204b4a2f9844a8eedb9ed.
//
// Solidity: event Cancel(uint256 id)
func (_CrowdFund *CrowdFundFilterer) WatchCancel(opts *bind.WatchOpts, sink chan<- *CrowdFundCancel) (event.Subscription, error) {

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Cancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundCancel)
				if err := _CrowdFund.contract.UnpackLog(event, "Cancel", log); err != nil {
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

// ParseCancel is a log parse operation binding the contract event 0x8bf30e7ff26833413be5f69e1d373744864d600b664204b4a2f9844a8eedb9ed.
//
// Solidity: event Cancel(uint256 id)
func (_CrowdFund *CrowdFundFilterer) ParseCancel(log types.Log) (*CrowdFundCancel, error) {
	event := new(CrowdFundCancel)
	if err := _CrowdFund.contract.UnpackLog(event, "Cancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the CrowdFund contract.
type CrowdFundClaimIterator struct {
	Event *CrowdFundClaim // Event containing the contract specifics and raw log

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
func (it *CrowdFundClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundClaim)
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
		it.Event = new(CrowdFundClaim)
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
func (it *CrowdFundClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundClaim represents a Claim event raised by the CrowdFund contract.
type CrowdFundClaim struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7.
//
// Solidity: event Claim(uint256 id)
func (_CrowdFund *CrowdFundFilterer) FilterClaim(opts *bind.FilterOpts) (*CrowdFundClaimIterator, error) {

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &CrowdFundClaimIterator{contract: _CrowdFund.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7.
//
// Solidity: event Claim(uint256 id)
func (_CrowdFund *CrowdFundFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *CrowdFundClaim) (event.Subscription, error) {

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundClaim)
				if err := _CrowdFund.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7.
//
// Solidity: event Claim(uint256 id)
func (_CrowdFund *CrowdFundFilterer) ParseClaim(log types.Log) (*CrowdFundClaim, error) {
	event := new(CrowdFundClaim)
	if err := _CrowdFund.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundLaunchIterator is returned from FilterLaunch and is used to iterate over the raw logs and unpacked data for Launch events raised by the CrowdFund contract.
type CrowdFundLaunchIterator struct {
	Event *CrowdFundLaunch // Event containing the contract specifics and raw log

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
func (it *CrowdFundLaunchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundLaunch)
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
		it.Event = new(CrowdFundLaunch)
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
func (it *CrowdFundLaunchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundLaunchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundLaunch represents a Launch event raised by the CrowdFund contract.
type CrowdFundLaunch struct {
	Id      *big.Int
	Creator common.Address
	Goal    *big.Int
	StartAt uint32
	EndAt   uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLaunch is a free log retrieval operation binding the contract event 0x0601cd0d650b473037e838a2696d41e654433d065b3f56b28d1d3302e44a304f.
//
// Solidity: event Launch(uint256 id, address indexed creator, uint256 goal, uint32 startAt, uint32 endAt)
func (_CrowdFund *CrowdFundFilterer) FilterLaunch(opts *bind.FilterOpts, creator []common.Address) (*CrowdFundLaunchIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Launch", creatorRule)
	if err != nil {
		return nil, err
	}
	return &CrowdFundLaunchIterator{contract: _CrowdFund.contract, event: "Launch", logs: logs, sub: sub}, nil
}

// WatchLaunch is a free log subscription operation binding the contract event 0x0601cd0d650b473037e838a2696d41e654433d065b3f56b28d1d3302e44a304f.
//
// Solidity: event Launch(uint256 id, address indexed creator, uint256 goal, uint32 startAt, uint32 endAt)
func (_CrowdFund *CrowdFundFilterer) WatchLaunch(opts *bind.WatchOpts, sink chan<- *CrowdFundLaunch, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Launch", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundLaunch)
				if err := _CrowdFund.contract.UnpackLog(event, "Launch", log); err != nil {
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

// ParseLaunch is a log parse operation binding the contract event 0x0601cd0d650b473037e838a2696d41e654433d065b3f56b28d1d3302e44a304f.
//
// Solidity: event Launch(uint256 id, address indexed creator, uint256 goal, uint32 startAt, uint32 endAt)
func (_CrowdFund *CrowdFundFilterer) ParseLaunch(log types.Log) (*CrowdFundLaunch, error) {
	event := new(CrowdFundLaunch)
	if err := _CrowdFund.contract.UnpackLog(event, "Launch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundPledgeIterator is returned from FilterPledge and is used to iterate over the raw logs and unpacked data for Pledge events raised by the CrowdFund contract.
type CrowdFundPledgeIterator struct {
	Event *CrowdFundPledge // Event containing the contract specifics and raw log

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
func (it *CrowdFundPledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundPledge)
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
		it.Event = new(CrowdFundPledge)
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
func (it *CrowdFundPledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundPledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundPledge represents a Pledge event raised by the CrowdFund contract.
type CrowdFundPledge struct {
	Id     *big.Int
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPledge is a free log retrieval operation binding the contract event 0x06bdb975df800a73232998e71ed585d536222f1dfeaa622d7f62a23ada686c82.
//
// Solidity: event Pledge(uint256 indexed id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) FilterPledge(opts *bind.FilterOpts, id []*big.Int, caller []common.Address) (*CrowdFundPledgeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Pledge", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdFundPledgeIterator{contract: _CrowdFund.contract, event: "Pledge", logs: logs, sub: sub}, nil
}

// WatchPledge is a free log subscription operation binding the contract event 0x06bdb975df800a73232998e71ed585d536222f1dfeaa622d7f62a23ada686c82.
//
// Solidity: event Pledge(uint256 indexed id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) WatchPledge(opts *bind.WatchOpts, sink chan<- *CrowdFundPledge, id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Pledge", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundPledge)
				if err := _CrowdFund.contract.UnpackLog(event, "Pledge", log); err != nil {
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

// ParsePledge is a log parse operation binding the contract event 0x06bdb975df800a73232998e71ed585d536222f1dfeaa622d7f62a23ada686c82.
//
// Solidity: event Pledge(uint256 indexed id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) ParsePledge(log types.Log) (*CrowdFundPledge, error) {
	event := new(CrowdFundPledge)
	if err := _CrowdFund.contract.UnpackLog(event, "Pledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the CrowdFund contract.
type CrowdFundRefundIterator struct {
	Event *CrowdFundRefund // Event containing the contract specifics and raw log

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
func (it *CrowdFundRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundRefund)
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
		it.Event = new(CrowdFundRefund)
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
func (it *CrowdFundRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundRefund represents a Refund event raised by the CrowdFund contract.
type CrowdFundRefund struct {
	Id     *big.Int
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x21e12a7cad0da5928167e1084ea4d5fdf8d9af66657a2543a9ac76a0ca081477.
//
// Solidity: event Refund(uint256 id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) FilterRefund(opts *bind.FilterOpts, caller []common.Address) (*CrowdFundRefundIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Refund", callerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdFundRefundIterator{contract: _CrowdFund.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x21e12a7cad0da5928167e1084ea4d5fdf8d9af66657a2543a9ac76a0ca081477.
//
// Solidity: event Refund(uint256 id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *CrowdFundRefund, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Refund", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundRefund)
				if err := _CrowdFund.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x21e12a7cad0da5928167e1084ea4d5fdf8d9af66657a2543a9ac76a0ca081477.
//
// Solidity: event Refund(uint256 id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) ParseRefund(log types.Log) (*CrowdFundRefund, error) {
	event := new(CrowdFundRefund)
	if err := _CrowdFund.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdFundUnpledgeIterator is returned from FilterUnpledge and is used to iterate over the raw logs and unpacked data for Unpledge events raised by the CrowdFund contract.
type CrowdFundUnpledgeIterator struct {
	Event *CrowdFundUnpledge // Event containing the contract specifics and raw log

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
func (it *CrowdFundUnpledgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdFundUnpledge)
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
		it.Event = new(CrowdFundUnpledge)
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
func (it *CrowdFundUnpledgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdFundUnpledgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdFundUnpledge represents a Unpledge event raised by the CrowdFund contract.
type CrowdFundUnpledge struct {
	Id     *big.Int
	Caller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnpledge is a free log retrieval operation binding the contract event 0x2eeeab891b26a214d1b25749f88a406bdea852bd8c9bfda977e0ef8114c180ba.
//
// Solidity: event Unpledge(uint256 indexed id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) FilterUnpledge(opts *bind.FilterOpts, id []*big.Int, caller []common.Address) (*CrowdFundUnpledgeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.FilterLogs(opts, "Unpledge", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdFundUnpledgeIterator{contract: _CrowdFund.contract, event: "Unpledge", logs: logs, sub: sub}, nil
}

// WatchUnpledge is a free log subscription operation binding the contract event 0x2eeeab891b26a214d1b25749f88a406bdea852bd8c9bfda977e0ef8114c180ba.
//
// Solidity: event Unpledge(uint256 indexed id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) WatchUnpledge(opts *bind.WatchOpts, sink chan<- *CrowdFundUnpledge, id []*big.Int, caller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _CrowdFund.contract.WatchLogs(opts, "Unpledge", idRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdFundUnpledge)
				if err := _CrowdFund.contract.UnpackLog(event, "Unpledge", log); err != nil {
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

// ParseUnpledge is a log parse operation binding the contract event 0x2eeeab891b26a214d1b25749f88a406bdea852bd8c9bfda977e0ef8114c180ba.
//
// Solidity: event Unpledge(uint256 indexed id, address indexed caller, uint256 amount)
func (_CrowdFund *CrowdFundFilterer) ParseUnpledge(log types.Log) (*CrowdFundUnpledge, error) {
	event := new(CrowdFundUnpledge)
	if err := _CrowdFund.contract.UnpackLog(event, "Unpledge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
