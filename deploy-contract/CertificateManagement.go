// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// CertificateManagementCertificate is an auto generated low-level Go binding around an user-defined struct.
type CertificateManagementCertificate struct {
	Id              *big.Int
	StudentName     string
	Grade           *big.Int
	Course          string
	StudentAddress  common.Address
	ApplicationTime *big.Int
	ApprovalTime    *big.Int
	Status          string
	Approver        common.Address
}

// CertificateManagementABI is the input ABI used to generate the binding from.
const CertificateManagementABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"studentAddress\",\"type\":\"address\"}],\"name\":\"CertificateApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"CertificateApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"CertificateRejected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_studentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"}],\"name\":\"applyForCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_grade\",\"type\":\"uint256\"}],\"name\":\"approveCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"certificateCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"certificates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"studentName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"grade\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"studentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"applicationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"approvalTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getCertificateDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_student\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_studentName\",\"type\":\"string\"}],\"name\":\"getPendingCertificates\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"studentName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"grade\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"studentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"applicationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"approvalTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"internalType\":\"structCertificateManagement.Certificate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_student\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_status\",\"type\":\"string\"}],\"name\":\"getStudentCertificates\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"studentName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"grade\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"course\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"studentAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"applicationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"approvalTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"internalType\":\"structCertificateManagement.Certificate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_studentName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_course\",\"type\":\"string\"}],\"name\":\"modifyPendingCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingCertificates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"rejectCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"studentCertificates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CertificateManagementBin is the compiled bytecode used for deploying new contracts.
var CertificateManagementBin = "0x608060405234801561001057600080fd5b5061294d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063a901fe5e11610071578063a901fe5e14610158578063c3934f0e14610190578063c9eb1f58146101c0578063dc81e194146101f0578063e08b2ab614610220578063fca400f614610250576100a9565b806303ace42b146100ae5780634c9e3f2d146100ca57806352d7b89c146100e6578063663b3e22146101045780637fe9829a1461013c575b600080fd5b6100c860048036038101906100c39190611f77565b61026c565b005b6100e460048036038101906100df9190612016565b610482565b005b6100ee610743565b6040516100fb91906120b0565b60405180910390f35b61011e600480360381019061011991906120cb565b610749565b6040516101339998979695949392919061218f565b60405180910390f35b61015660048036038101906101519190612231565b61096f565b005b610172600480360381019061016d91906120cb565b610bec565b6040516101879998979695949392919061218f565b60405180910390f35b6101aa60048036038101906101a59190612284565b610ef6565b6040516101b791906120b0565b60405180910390f35b6101da60048036038101906101d59190612016565b610f27565b6040516101e791906124c1565b60405180910390f35b61020a60048036038101906102059190612016565b611402565b60405161021791906124c1565b60405180910390f35b61023a600480360381019061023591906120cb565b61191c565b60405161024791906120b0565b60405180910390f35b61026a60048036038101906102659190612284565b611940565b005b60405160200161027b9061253a565b604051602081830303815290604052805190602001206000808581526020019081526020016000206007016040516020016102b69190612644565b604051602081830303815290604052805190602001201461030c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610303906126a7565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1660008085815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146103af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a690612713565b60405180910390fd5b8160008085815260200190815260200160002060010190805190602001906103d8929190611c6e565b50806000808581526020019081526020016000206003019080519060200190610402929190611c6e565b5060008084815260200190815260200160002060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16837fbaa64ac25db7cce96af0e1f049feb7b9e6c9c462e18c47e1259f79d7174585dc60405160405180910390a350505050565b6003600081548092919061049590612762565b91905055506040518061012001604052806003548152602001838152602001600081526020018281526020018473ffffffffffffffffffffffffffffffffffffffff168152602001428152602001600081526020016040518060400160405280600781526020017f50656e64696e67000000000000000000000000000000000000000000000000008152508152602001600073ffffffffffffffffffffffffffffffffffffffff16815250600080600354815260200190815260200160002060008201518160000155602082015181600101908051906020019061057a929190611c6e565b506040820151816002015560608201518160030190805190602001906105a1929190611c6e565b5060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816005015560c0820151816006015560e0820151816007019080519060200190610619929190611c6e565b506101008201518160080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003549080600181540180825580915050600190039060005260206000200160009091909190915055600260035490806001815401808255809150506001900390600052602060002001600090919091909150558273ffffffffffffffffffffffffffffffffffffffff166003547fbaa64ac25db7cce96af0e1f049feb7b9e6c9c462e18c47e1259f79d7174585dc60405160405180910390a3505050565b60035481565b60006020528060005260406000206000915090508060000154908060010180546107729061257e565b80601f016020809104026020016040519081016040528092919081815260200182805461079e9061257e565b80156107eb5780601f106107c0576101008083540402835291602001916107eb565b820191906000526020600020905b8154815290600101906020018083116107ce57829003601f168201915b5050505050908060020154908060030180546108069061257e565b80601f01602080910402602001604051908101604052809291908181526020018280546108329061257e565b801561087f5780601f106108545761010080835404028352916020019161087f565b820191906000526020600020905b81548152906001019060200180831161086257829003601f168201915b5050505050908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154908060060154908060070180546108c69061257e565b80601f01602080910402602001604051908101604052809291908181526020018280546108f29061257e565b801561093f5780601f106109145761010080835404028352916020019161093f565b820191906000526020600020905b81548152906001019060200180831161092257829003601f168201915b5050505050908060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905089565b60405160200161097e9061253a565b604051602081830303815290604052805190602001206000808481526020019081526020016000206007016040516020016109b99190612644565b6040516020818303038152906040528051906020012014610a0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a06906126a7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008084815260200190815260200160002060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610ab3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aaa9061281d565b60405180910390fd5b80600080848152602001908152602001600020600201819055506040518060400160405280600881526020017f417070726f7665640000000000000000000000000000000000000000000000008152506000808481526020019081526020016000206007019080519060200190610b2b929190611c6e565b5042600080848152602001908152602001600020600601819055508260008084815260200190815260200160002060080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610ba382611ba2565b8273ffffffffffffffffffffffffffffffffffffffff16827f1ffd4b0c43de7f13bb9e8c45b1b35af572231837bf8a3c876e9eaf96a3cbd3d960405160405180910390a3505050565b6000606060006060600080600060606000806000808c81526020019081526020016000206040518061012001604052908160008201548152602001600182018054610c369061257e565b80601f0160208091040260200160405190810160405280929190818152602001828054610c629061257e565b8015610caf5780601f10610c8457610100808354040283529160200191610caf565b820191906000526020600020905b815481529060010190602001808311610c9257829003601f168201915b5050505050815260200160028201548152602001600382018054610cd29061257e565b80601f0160208091040260200160405190810160405280929190818152602001828054610cfe9061257e565b8015610d4b5780601f10610d2057610100808354040283529160200191610d4b565b820191906000526020600020905b815481529060010190602001808311610d2e57829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820154815260200160068201548152602001600782018054610dce9061257e565b80601f0160208091040260200160405190810160405280929190818152602001828054610dfa9061257e565b8015610e475780601f10610e1c57610100808354040283529160200191610e47565b820191906000526020600020905b815481529060010190602001808311610e2a57829003601f168201915b505050505081526020016008820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050806000015181602001518260400151836060015184608001518560a001518660c001518760e00151886101000151995099509950995099509950995099509950509193959799909294969850565b60016020528160005260406000208181548110610f1257600080fd5b90600052602060002001600091509150505481565b6060600060028054905067ffffffffffffffff811115610f4a57610f49611e4c565b5b604051908082528060200260200182016040528015610f8357816020015b610f70611cf4565b815260200190600190039081610f685790505b5090506000805b6002805490508110156113f557600080600060028481548110610fb057610faf61283d565b5b906000526020600020015481526020019081526020016000206040518061012001604052908160008201548152602001600182018054610fef9061257e565b80601f016020809104026020016040519081016040528092919081815260200182805461101b9061257e565b80156110685780601f1061103d57610100808354040283529160200191611068565b820191906000526020600020905b81548152906001019060200180831161104b57829003601f168201915b505050505081526020016002820154815260200160038201805461108b9061257e565b80601f01602080910402602001604051908101604052809291908181526020018280546110b79061257e565b80156111045780601f106110d957610100808354040283529160200191611104565b820191906000526020600020905b8154815290600101906020018083116110e757829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201548152602001600682015481526020016007820180546111879061257e565b80601f01602080910402602001604051908101604052809291908181526020018280546111b39061257e565b80156112005780601f106111d557610100808354040283529160200191611200565b820191906000526020600020905b8154815290600101906020018083116111e357829003601f168201915b505050505081526020016008820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050905060008751141580156112c357508060600151604051602001611283919061289d565b60405160208183030381529060405280519060200120876040516020016112aa919061289d565b6040516020818303038152906040528051906020012014155b156112ce57506113e2565b6000865114158015611330575080602001516040516020016112f0919061289d565b6040516020818303038152906040528051906020012086604051602001611317919061289d565b6040516020818303038152906040528051906020012014155b1561133b57506113e2565b600073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff16141580156113a85750806080015173ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614155b156113b357506113e2565b808484815181106113c7576113c661283d565b5b602002602001018190525082806113dd90612762565b935050505b80806113ed90612762565b915050610f8a565b5081925050509392505050565b60606000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905067ffffffffffffffff81111561146257611461611e4c565b5b60405190808252806020026020018201604052801561149b57816020015b611488611cf4565b8152602001906001900390816114805790505b5090506000805b600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905081101561190f576000806000600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084815481106115425761154161283d565b5b9060005260206000200154815260200190815260200160002060405180610120016040529081600082015481526020016001820180546115819061257e565b80601f01602080910402602001604051908101604052809291908181526020018280546115ad9061257e565b80156115fa5780601f106115cf576101008083540402835291602001916115fa565b820191906000526020600020905b8154815290600101906020018083116115dd57829003601f168201915b505050505081526020016002820154815260200160038201805461161d9061257e565b80601f01602080910402602001604051908101604052809291908181526020018280546116499061257e565b80156116965780601f1061166b57610100808354040283529160200191611696565b820191906000526020600020905b81548152906001019060200180831161167957829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201548152602001600682015481526020016007820180546117199061257e565b80601f01602080910402602001604051908101604052809291908181526020018280546117459061257e565b80156117925780601f1061176757610100808354040283529160200191611792565b820191906000526020600020905b81548152906001019060200180831161177557829003601f168201915b505050505081526020016008820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050600087511415801561185557508060600151604051602001611815919061289d565b604051602081830303815290604052805190602001208760405160200161183c919061289d565b6040516020818303038152906040528051906020012014155b1561186057506118fc565b60008651141580156118c257508060e00151604051602001611882919061289d565b60405160208183030381529060405280519060200120866040516020016118a9919061289d565b6040516020818303038152906040528051906020012014155b156118cd57506118fc565b808484815181106118e1576118e061283d565b5b602002602001018190525082806118f790612762565b935050505b808061190790612762565b9150506114a2565b5081925050509392505050565b6002818154811061192c57600080fd5b906000526020600020016000915090505481565b60405160200161194f9061253a565b6040516020818303038152906040528051906020012060008083815260200190815260200160002060070160405160200161198a9190612644565b60405160208183030381529060405280519060200120146119e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119d7906126a7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008083815260200190815260200160002060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611a84576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a7b9061281d565b60405180910390fd5b6040518060400160405280600881526020017f52656a65637465640000000000000000000000000000000000000000000000008152506000808381526020019081526020016000206007019080519060200190611ae2929190611c6e565b5042600080838152602001908152602001600020600601819055508160008083815260200190815260200160002060080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611b5a81611ba2565b8173ffffffffffffffffffffffffffffffffffffffff16817fd16874d1c78ac14ec0cc9028724b5f14c02545485fb5e0bdf58457d3fe3c194660405160405180910390a35050565b60005b600280549050811015611c6a578160028281548110611bc757611bc661283d565b5b90600052602060002001541415611c575760026001600280549050611bec91906128b4565b81548110611bfd57611bfc61283d565b5b906000526020600020015460028281548110611c1c57611c1b61283d565b5b90600052602060002001819055506002805480611c3c57611c3b6128e8565b5b60019003818190600052602060002001600090559055611c6a565b8080611c6290612762565b915050611ba5565b5050565b828054611c7a9061257e565b90600052602060002090601f016020900481019282611c9c5760008555611ce3565b82601f10611cb557805160ff1916838001178555611ce3565b82800160010185558215611ce3579182015b82811115611ce2578251825591602001919060010190611cc7565b5b509050611cf09190611d6c565b5090565b60405180610120016040528060008152602001606081526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b5b80821115611d85576000816000905550600101611d6d565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611dc882611d9d565b9050919050565b611dd881611dbd565b8114611de357600080fd5b50565b600081359050611df581611dcf565b92915050565b6000819050919050565b611e0e81611dfb565b8114611e1957600080fd5b50565b600081359050611e2b81611e05565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611e8482611e3b565b810181811067ffffffffffffffff82111715611ea357611ea2611e4c565b5b80604052505050565b6000611eb6611d89565b9050611ec28282611e7b565b919050565b600067ffffffffffffffff821115611ee257611ee1611e4c565b5b611eeb82611e3b565b9050602081019050919050565b82818337600083830152505050565b6000611f1a611f1584611ec7565b611eac565b905082815260208101848484011115611f3657611f35611e36565b5b611f41848285611ef8565b509392505050565b600082601f830112611f5e57611f5d611e31565b5b8135611f6e848260208601611f07565b91505092915050565b60008060008060808587031215611f9157611f90611d93565b5b6000611f9f87828801611de6565b9450506020611fb087828801611e1c565b935050604085013567ffffffffffffffff811115611fd157611fd0611d98565b5b611fdd87828801611f49565b925050606085013567ffffffffffffffff811115611ffe57611ffd611d98565b5b61200a87828801611f49565b91505092959194509250565b60008060006060848603121561202f5761202e611d93565b5b600061203d86828701611de6565b935050602084013567ffffffffffffffff81111561205e5761205d611d98565b5b61206a86828701611f49565b925050604084013567ffffffffffffffff81111561208b5761208a611d98565b5b61209786828701611f49565b9150509250925092565b6120aa81611dfb565b82525050565b60006020820190506120c560008301846120a1565b92915050565b6000602082840312156120e1576120e0611d93565b5b60006120ef84828501611e1c565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015612132578082015181840152602081019050612117565b83811115612141576000848401525b50505050565b6000612152826120f8565b61215c8185612103565b935061216c818560208601612114565b61217581611e3b565b840191505092915050565b61218981611dbd565b82525050565b6000610120820190506121a5600083018c6120a1565b81810360208301526121b7818b612147565b90506121c6604083018a6120a1565b81810360608301526121d88189612147565b90506121e76080830188612180565b6121f460a08301876120a1565b61220160c08301866120a1565b81810360e08301526122138185612147565b9050612223610100830184612180565b9a9950505050505050505050565b60008060006060848603121561224a57612249611d93565b5b600061225886828701611de6565b935050602061226986828701611e1c565b925050604061227a86828701611e1c565b9150509250925092565b6000806040838503121561229b5761229a611d93565b5b60006122a985828601611de6565b92505060206122ba85828601611e1c565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6122f981611dfb565b82525050565b600082825260208201905092915050565b600061231b826120f8565b61232581856122ff565b9350612335818560208601612114565b61233e81611e3b565b840191505092915050565b61235281611dbd565b82525050565b60006101208301600083015161237160008601826122f0565b50602083015184820360208601526123898282612310565b915050604083015161239e60408601826122f0565b50606083015184820360608601526123b68282612310565b91505060808301516123cb6080860182612349565b5060a08301516123de60a08601826122f0565b5060c08301516123f160c08601826122f0565b5060e083015184820360e08601526124098282612310565b915050610100830151612420610100860182612349565b508091505092915050565b60006124378383612358565b905092915050565b6000602082019050919050565b6000612457826122c4565b61246181856122cf565b935083602082028501612473856122e0565b8060005b858110156124af5784840389528151612490858261242b565b945061249b8361243f565b925060208a01995050600181019050612477565b50829750879550505050505092915050565b600060208201905081810360008301526124db818461244c565b905092915050565b600081905092915050565b7f50656e64696e6700000000000000000000000000000000000000000000000000600082015250565b60006125246007836124e3565b915061252f826124ee565b600782019050919050565b600061254582612517565b9150819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061259657607f821691505b602082108114156125aa576125a961254f565b5b50919050565b60008190508160005260206000209050919050565b600081546125d28161257e565b6125dc81866124e3565b945060018216600081146125f757600181146126085761263b565b60ff1983168652818601935061263b565b612611856125b0565b60005b8381101561263357815481890152600182019150602081019050612614565b838801955050505b50505092915050565b600061265082846125c5565b915081905092915050565b7f4365727469666963617465206973206e6f742070656e64696e67000000000000600082015250565b6000612691601a83612103565b915061269c8261265b565b602082019050919050565b600060208201905081810360008301526126c081612684565b9050919050565b7f53747564656e74206164647265737320646f6573206e6f74206d617463680000600082015250565b60006126fd601e83612103565b9150612708826126c7565b602082019050919050565b6000602082019050818103600083015261272c816126f0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061276d82611dfb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156127a05761279f612733565b5b600182019050919050565b7f436572746966696361746520616c726561647920617070726f766564206f722060008201527f72656a6563746564000000000000000000000000000000000000000000000000602082015250565b6000612807602883612103565b9150612812826127ab565b604082019050919050565b60006020820190508181036000830152612836816127fa565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000612877826120f8565b61288181856124e3565b9350612891818560208601612114565b80840191505092915050565b60006128a9828461286c565b915081905092915050565b60006128bf82611dfb565b91506128ca83611dfb565b9250828210156128dd576128dc612733565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220a2eb6e4836acad947b2c131007c6dab8b86c70ac7245066f1a678302318f5e4164736f6c634300080b0033"
var CertificateManagementSMBin = "0x"

// DeployCertificateManagement deploys a new contract, binding an instance of CertificateManagement to it.
func DeployCertificateManagement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *CertificateManagement, error) {
	parsed, err := abi.JSON(strings.NewReader(CertificateManagementABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(CertificateManagementSMBin)
	} else {
		bytecode = common.FromHex(CertificateManagementBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, CertificateManagementABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &CertificateManagement{CertificateManagementCaller: CertificateManagementCaller{contract: contract}, CertificateManagementTransactor: CertificateManagementTransactor{contract: contract}, CertificateManagementFilterer: CertificateManagementFilterer{contract: contract}}, nil
}

func AsyncDeployCertificateManagement(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(CertificateManagementABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(CertificateManagementSMBin)
	} else {
		bytecode = common.FromHex(CertificateManagementBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, CertificateManagementABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// CertificateManagement is an auto generated Go binding around a Solidity contract.
type CertificateManagement struct {
	CertificateManagementCaller     // Read-only binding to the contract
	CertificateManagementTransactor // Write-only binding to the contract
	CertificateManagementFilterer   // Log filterer for contract events
}

// CertificateManagementCaller is an auto generated read-only Go binding around a Solidity contract.
type CertificateManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateManagementTransactor is an auto generated write-only Go binding around a Solidity contract.
type CertificateManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateManagementFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type CertificateManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateManagementSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type CertificateManagementSession struct {
	Contract     *CertificateManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CertificateManagementCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type CertificateManagementCallerSession struct {
	Contract *CertificateManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// CertificateManagementTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type CertificateManagementTransactorSession struct {
	Contract     *CertificateManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// CertificateManagementRaw is an auto generated low-level Go binding around a Solidity contract.
type CertificateManagementRaw struct {
	Contract *CertificateManagement // Generic contract binding to access the raw methods on
}

// CertificateManagementCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type CertificateManagementCallerRaw struct {
	Contract *CertificateManagementCaller // Generic read-only contract binding to access the raw methods on
}

// CertificateManagementTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type CertificateManagementTransactorRaw struct {
	Contract *CertificateManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificateManagement creates a new instance of CertificateManagement, bound to a specific deployed contract.
func NewCertificateManagement(address common.Address, backend bind.ContractBackend) (*CertificateManagement, error) {
	contract, err := bindCertificateManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CertificateManagement{CertificateManagementCaller: CertificateManagementCaller{contract: contract}, CertificateManagementTransactor: CertificateManagementTransactor{contract: contract}, CertificateManagementFilterer: CertificateManagementFilterer{contract: contract}}, nil
}

// NewCertificateManagementCaller creates a new read-only instance of CertificateManagement, bound to a specific deployed contract.
func NewCertificateManagementCaller(address common.Address, caller bind.ContractCaller) (*CertificateManagementCaller, error) {
	contract, err := bindCertificateManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateManagementCaller{contract: contract}, nil
}

// NewCertificateManagementTransactor creates a new write-only instance of CertificateManagement, bound to a specific deployed contract.
func NewCertificateManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificateManagementTransactor, error) {
	contract, err := bindCertificateManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateManagementTransactor{contract: contract}, nil
}

// NewCertificateManagementFilterer creates a new log filterer instance of CertificateManagement, bound to a specific deployed contract.
func NewCertificateManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificateManagementFilterer, error) {
	contract, err := bindCertificateManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificateManagementFilterer{contract: contract}, nil
}

// bindCertificateManagement binds a generic wrapper to an already deployed contract.
func bindCertificateManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CertificateManagementABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateManagement *CertificateManagementRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CertificateManagement.Contract.CertificateManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateManagement *CertificateManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.CertificateManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateManagement *CertificateManagementRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.CertificateManagementTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateManagement *CertificateManagementCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CertificateManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateManagement *CertificateManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateManagement *CertificateManagementTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// CertificateCount is a free data retrieval call binding the contract method 0x52d7b89c.
//
// Solidity: function certificateCount() constant returns(uint256)
func (_CertificateManagement *CertificateManagementCaller) CertificateCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CertificateManagement.contract.Call(opts, out, "certificateCount")
	return *ret0, err
}

// CertificateCount is a free data retrieval call binding the contract method 0x52d7b89c.
//
// Solidity: function certificateCount() constant returns(uint256)
func (_CertificateManagement *CertificateManagementSession) CertificateCount() (*big.Int, error) {
	return _CertificateManagement.Contract.CertificateCount(&_CertificateManagement.CallOpts)
}

// CertificateCount is a free data retrieval call binding the contract method 0x52d7b89c.
//
// Solidity: function certificateCount() constant returns(uint256)
func (_CertificateManagement *CertificateManagementCallerSession) CertificateCount() (*big.Int, error) {
	return _CertificateManagement.Contract.CertificateCount(&_CertificateManagement.CallOpts)
}

// Certificates is a free data retrieval call binding the contract method 0x663b3e22.
//
// Solidity: function certificates(uint256 ) constant returns(uint256 id, string studentName, uint256 grade, string course, address studentAddress, uint256 applicationTime, uint256 approvalTime, string status, address approver)
func (_CertificateManagement *CertificateManagementCaller) Certificates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id              *big.Int
	StudentName     string
	Grade           *big.Int
	Course          string
	StudentAddress  common.Address
	ApplicationTime *big.Int
	ApprovalTime    *big.Int
	Status          string
	Approver        common.Address
}, error) {
	ret := new(struct {
		Id              *big.Int
		StudentName     string
		Grade           *big.Int
		Course          string
		StudentAddress  common.Address
		ApplicationTime *big.Int
		ApprovalTime    *big.Int
		Status          string
		Approver        common.Address
	})
	out := ret
	err := _CertificateManagement.contract.Call(opts, out, "certificates", arg0)
	return *ret, err
}

// Certificates is a free data retrieval call binding the contract method 0x663b3e22.
//
// Solidity: function certificates(uint256 ) constant returns(uint256 id, string studentName, uint256 grade, string course, address studentAddress, uint256 applicationTime, uint256 approvalTime, string status, address approver)
func (_CertificateManagement *CertificateManagementSession) Certificates(arg0 *big.Int) (struct {
	Id              *big.Int
	StudentName     string
	Grade           *big.Int
	Course          string
	StudentAddress  common.Address
	ApplicationTime *big.Int
	ApprovalTime    *big.Int
	Status          string
	Approver        common.Address
}, error) {
	return _CertificateManagement.Contract.Certificates(&_CertificateManagement.CallOpts, arg0)
}

// Certificates is a free data retrieval call binding the contract method 0x663b3e22.
//
// Solidity: function certificates(uint256 ) constant returns(uint256 id, string studentName, uint256 grade, string course, address studentAddress, uint256 applicationTime, uint256 approvalTime, string status, address approver)
func (_CertificateManagement *CertificateManagementCallerSession) Certificates(arg0 *big.Int) (struct {
	Id              *big.Int
	StudentName     string
	Grade           *big.Int
	Course          string
	StudentAddress  common.Address
	ApplicationTime *big.Int
	ApprovalTime    *big.Int
	Status          string
	Approver        common.Address
}, error) {
	return _CertificateManagement.Contract.Certificates(&_CertificateManagement.CallOpts, arg0)
}

// GetCertificateDetails is a free data retrieval call binding the contract method 0xa901fe5e.
//
// Solidity: function getCertificateDetails(uint256 _id) constant returns(uint256, string, uint256, string, address, uint256, uint256, string, address)
func (_CertificateManagement *CertificateManagementCaller) GetCertificateDetails(opts *bind.CallOpts, _id *big.Int) (*big.Int, string, *big.Int, string, common.Address, *big.Int, *big.Int, string, common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(common.Address)
		ret5 = new(*big.Int)
		ret6 = new(*big.Int)
		ret7 = new(string)
		ret8 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
	}
	err := _CertificateManagement.contract.Call(opts, out, "getCertificateDetails", _id)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, err
}

// GetCertificateDetails is a free data retrieval call binding the contract method 0xa901fe5e.
//
// Solidity: function getCertificateDetails(uint256 _id) constant returns(uint256, string, uint256, string, address, uint256, uint256, string, address)
func (_CertificateManagement *CertificateManagementSession) GetCertificateDetails(_id *big.Int) (*big.Int, string, *big.Int, string, common.Address, *big.Int, *big.Int, string, common.Address, error) {
	return _CertificateManagement.Contract.GetCertificateDetails(&_CertificateManagement.CallOpts, _id)
}

// GetCertificateDetails is a free data retrieval call binding the contract method 0xa901fe5e.
//
// Solidity: function getCertificateDetails(uint256 _id) constant returns(uint256, string, uint256, string, address, uint256, uint256, string, address)
func (_CertificateManagement *CertificateManagementCallerSession) GetCertificateDetails(_id *big.Int) (*big.Int, string, *big.Int, string, common.Address, *big.Int, *big.Int, string, common.Address, error) {
	return _CertificateManagement.Contract.GetCertificateDetails(&_CertificateManagement.CallOpts, _id)
}

// GetPendingCertificates is a free data retrieval call binding the contract method 0xc9eb1f58.
//
// Solidity: function getPendingCertificates(address _student, string _course, string _studentName) constant returns([]CertificateManagementCertificate)
func (_CertificateManagement *CertificateManagementCaller) GetPendingCertificates(opts *bind.CallOpts, _student common.Address, _course string, _studentName string) ([]CertificateManagementCertificate, error) {
	var (
		ret0 = new([]CertificateManagementCertificate)
	)
	out := ret0
	err := _CertificateManagement.contract.Call(opts, out, "getPendingCertificates", _student, _course, _studentName)
	return *ret0, err
}

// GetPendingCertificates is a free data retrieval call binding the contract method 0xc9eb1f58.
//
// Solidity: function getPendingCertificates(address _student, string _course, string _studentName) constant returns([]CertificateManagementCertificate)
func (_CertificateManagement *CertificateManagementSession) GetPendingCertificates(_student common.Address, _course string, _studentName string) ([]CertificateManagementCertificate, error) {
	return _CertificateManagement.Contract.GetPendingCertificates(&_CertificateManagement.CallOpts, _student, _course, _studentName)
}

// GetPendingCertificates is a free data retrieval call binding the contract method 0xc9eb1f58.
//
// Solidity: function getPendingCertificates(address _student, string _course, string _studentName) constant returns([]CertificateManagementCertificate)
func (_CertificateManagement *CertificateManagementCallerSession) GetPendingCertificates(_student common.Address, _course string, _studentName string) ([]CertificateManagementCertificate, error) {
	return _CertificateManagement.Contract.GetPendingCertificates(&_CertificateManagement.CallOpts, _student, _course, _studentName)
}

// GetStudentCertificates is a free data retrieval call binding the contract method 0xdc81e194.
//
// Solidity: function getStudentCertificates(address _student, string _course, string _status) constant returns([]CertificateManagementCertificate)
func (_CertificateManagement *CertificateManagementCaller) GetStudentCertificates(opts *bind.CallOpts, _student common.Address, _course string, _status string) ([]CertificateManagementCertificate, error) {
	var (
		ret0 = new([]CertificateManagementCertificate)
	)
	out := ret0
	err := _CertificateManagement.contract.Call(opts, out, "getStudentCertificates", _student, _course, _status)
	return *ret0, err
}

// GetStudentCertificates is a free data retrieval call binding the contract method 0xdc81e194.
//
// Solidity: function getStudentCertificates(address _student, string _course, string _status) constant returns([]CertificateManagementCertificate)
func (_CertificateManagement *CertificateManagementSession) GetStudentCertificates(_student common.Address, _course string, _status string) ([]CertificateManagementCertificate, error) {
	return _CertificateManagement.Contract.GetStudentCertificates(&_CertificateManagement.CallOpts, _student, _course, _status)
}

// GetStudentCertificates is a free data retrieval call binding the contract method 0xdc81e194.
//
// Solidity: function getStudentCertificates(address _student, string _course, string _status) constant returns([]CertificateManagementCertificate)
func (_CertificateManagement *CertificateManagementCallerSession) GetStudentCertificates(_student common.Address, _course string, _status string) ([]CertificateManagementCertificate, error) {
	return _CertificateManagement.Contract.GetStudentCertificates(&_CertificateManagement.CallOpts, _student, _course, _status)
}

// PendingCertificates is a free data retrieval call binding the contract method 0xe08b2ab6.
//
// Solidity: function pendingCertificates(uint256 ) constant returns(uint256)
func (_CertificateManagement *CertificateManagementCaller) PendingCertificates(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CertificateManagement.contract.Call(opts, out, "pendingCertificates", arg0)
	return *ret0, err
}

// PendingCertificates is a free data retrieval call binding the contract method 0xe08b2ab6.
//
// Solidity: function pendingCertificates(uint256 ) constant returns(uint256)
func (_CertificateManagement *CertificateManagementSession) PendingCertificates(arg0 *big.Int) (*big.Int, error) {
	return _CertificateManagement.Contract.PendingCertificates(&_CertificateManagement.CallOpts, arg0)
}

// PendingCertificates is a free data retrieval call binding the contract method 0xe08b2ab6.
//
// Solidity: function pendingCertificates(uint256 ) constant returns(uint256)
func (_CertificateManagement *CertificateManagementCallerSession) PendingCertificates(arg0 *big.Int) (*big.Int, error) {
	return _CertificateManagement.Contract.PendingCertificates(&_CertificateManagement.CallOpts, arg0)
}

// StudentCertificates is a free data retrieval call binding the contract method 0xc3934f0e.
//
// Solidity: function studentCertificates(address , uint256 ) constant returns(uint256)
func (_CertificateManagement *CertificateManagementCaller) StudentCertificates(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CertificateManagement.contract.Call(opts, out, "studentCertificates", arg0, arg1)
	return *ret0, err
}

// StudentCertificates is a free data retrieval call binding the contract method 0xc3934f0e.
//
// Solidity: function studentCertificates(address , uint256 ) constant returns(uint256)
func (_CertificateManagement *CertificateManagementSession) StudentCertificates(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CertificateManagement.Contract.StudentCertificates(&_CertificateManagement.CallOpts, arg0, arg1)
}

// StudentCertificates is a free data retrieval call binding the contract method 0xc3934f0e.
//
// Solidity: function studentCertificates(address , uint256 ) constant returns(uint256)
func (_CertificateManagement *CertificateManagementCallerSession) StudentCertificates(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CertificateManagement.Contract.StudentCertificates(&_CertificateManagement.CallOpts, arg0, arg1)
}

// ApplyForCertificate is a paid mutator transaction binding the contract method 0x4c9e3f2d.
//
// Solidity: function applyForCertificate(address _address, string _studentName, string _course) returns()
func (_CertificateManagement *CertificateManagementTransactor) ApplyForCertificate(opts *bind.TransactOpts, _address common.Address, _studentName string, _course string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _CertificateManagement.contract.TransactWithResult(opts, out, "applyForCertificate", _address, _studentName, _course)
	return transaction, receipt, err
}

func (_CertificateManagement *CertificateManagementTransactor) AsyncApplyForCertificate(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address, _studentName string, _course string) (*types.Transaction, error) {
	return _CertificateManagement.contract.AsyncTransact(opts, handler, "applyForCertificate", _address, _studentName, _course)
}

// ApplyForCertificate is a paid mutator transaction binding the contract method 0x4c9e3f2d.
//
// Solidity: function applyForCertificate(address _address, string _studentName, string _course) returns()
func (_CertificateManagement *CertificateManagementSession) ApplyForCertificate(_address common.Address, _studentName string, _course string) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.ApplyForCertificate(&_CertificateManagement.TransactOpts, _address, _studentName, _course)
}

func (_CertificateManagement *CertificateManagementSession) AsyncApplyForCertificate(handler func(*types.Receipt, error), _address common.Address, _studentName string, _course string) (*types.Transaction, error) {
	return _CertificateManagement.Contract.AsyncApplyForCertificate(handler, &_CertificateManagement.TransactOpts, _address, _studentName, _course)
}

// ApplyForCertificate is a paid mutator transaction binding the contract method 0x4c9e3f2d.
//
// Solidity: function applyForCertificate(address _address, string _studentName, string _course) returns()
func (_CertificateManagement *CertificateManagementTransactorSession) ApplyForCertificate(_address common.Address, _studentName string, _course string) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.ApplyForCertificate(&_CertificateManagement.TransactOpts, _address, _studentName, _course)
}

func (_CertificateManagement *CertificateManagementTransactorSession) AsyncApplyForCertificate(handler func(*types.Receipt, error), _address common.Address, _studentName string, _course string) (*types.Transaction, error) {
	return _CertificateManagement.Contract.AsyncApplyForCertificate(handler, &_CertificateManagement.TransactOpts, _address, _studentName, _course)
}

// ApproveCertificate is a paid mutator transaction binding the contract method 0x7fe9829a.
//
// Solidity: function approveCertificate(address _address, uint256 _id, uint256 _grade) returns()
func (_CertificateManagement *CertificateManagementTransactor) ApproveCertificate(opts *bind.TransactOpts, _address common.Address, _id *big.Int, _grade *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _CertificateManagement.contract.TransactWithResult(opts, out, "approveCertificate", _address, _id, _grade)
	return transaction, receipt, err
}

func (_CertificateManagement *CertificateManagementTransactor) AsyncApproveCertificate(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address, _id *big.Int, _grade *big.Int) (*types.Transaction, error) {
	return _CertificateManagement.contract.AsyncTransact(opts, handler, "approveCertificate", _address, _id, _grade)
}

// ApproveCertificate is a paid mutator transaction binding the contract method 0x7fe9829a.
//
// Solidity: function approveCertificate(address _address, uint256 _id, uint256 _grade) returns()
func (_CertificateManagement *CertificateManagementSession) ApproveCertificate(_address common.Address, _id *big.Int, _grade *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.ApproveCertificate(&_CertificateManagement.TransactOpts, _address, _id, _grade)
}

func (_CertificateManagement *CertificateManagementSession) AsyncApproveCertificate(handler func(*types.Receipt, error), _address common.Address, _id *big.Int, _grade *big.Int) (*types.Transaction, error) {
	return _CertificateManagement.Contract.AsyncApproveCertificate(handler, &_CertificateManagement.TransactOpts, _address, _id, _grade)
}

// ApproveCertificate is a paid mutator transaction binding the contract method 0x7fe9829a.
//
// Solidity: function approveCertificate(address _address, uint256 _id, uint256 _grade) returns()
func (_CertificateManagement *CertificateManagementTransactorSession) ApproveCertificate(_address common.Address, _id *big.Int, _grade *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.ApproveCertificate(&_CertificateManagement.TransactOpts, _address, _id, _grade)
}

func (_CertificateManagement *CertificateManagementTransactorSession) AsyncApproveCertificate(handler func(*types.Receipt, error), _address common.Address, _id *big.Int, _grade *big.Int) (*types.Transaction, error) {
	return _CertificateManagement.Contract.AsyncApproveCertificate(handler, &_CertificateManagement.TransactOpts, _address, _id, _grade)
}

// ModifyPendingCertificate is a paid mutator transaction binding the contract method 0x03ace42b.
//
// Solidity: function modifyPendingCertificate(address _address, uint256 _id, string _studentName, string _course) returns()
func (_CertificateManagement *CertificateManagementTransactor) ModifyPendingCertificate(opts *bind.TransactOpts, _address common.Address, _id *big.Int, _studentName string, _course string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _CertificateManagement.contract.TransactWithResult(opts, out, "modifyPendingCertificate", _address, _id, _studentName, _course)
	return transaction, receipt, err
}

func (_CertificateManagement *CertificateManagementTransactor) AsyncModifyPendingCertificate(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address, _id *big.Int, _studentName string, _course string) (*types.Transaction, error) {
	return _CertificateManagement.contract.AsyncTransact(opts, handler, "modifyPendingCertificate", _address, _id, _studentName, _course)
}

// ModifyPendingCertificate is a paid mutator transaction binding the contract method 0x03ace42b.
//
// Solidity: function modifyPendingCertificate(address _address, uint256 _id, string _studentName, string _course) returns()
func (_CertificateManagement *CertificateManagementSession) ModifyPendingCertificate(_address common.Address, _id *big.Int, _studentName string, _course string) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.ModifyPendingCertificate(&_CertificateManagement.TransactOpts, _address, _id, _studentName, _course)
}

func (_CertificateManagement *CertificateManagementSession) AsyncModifyPendingCertificate(handler func(*types.Receipt, error), _address common.Address, _id *big.Int, _studentName string, _course string) (*types.Transaction, error) {
	return _CertificateManagement.Contract.AsyncModifyPendingCertificate(handler, &_CertificateManagement.TransactOpts, _address, _id, _studentName, _course)
}

// ModifyPendingCertificate is a paid mutator transaction binding the contract method 0x03ace42b.
//
// Solidity: function modifyPendingCertificate(address _address, uint256 _id, string _studentName, string _course) returns()
func (_CertificateManagement *CertificateManagementTransactorSession) ModifyPendingCertificate(_address common.Address, _id *big.Int, _studentName string, _course string) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.ModifyPendingCertificate(&_CertificateManagement.TransactOpts, _address, _id, _studentName, _course)
}

func (_CertificateManagement *CertificateManagementTransactorSession) AsyncModifyPendingCertificate(handler func(*types.Receipt, error), _address common.Address, _id *big.Int, _studentName string, _course string) (*types.Transaction, error) {
	return _CertificateManagement.Contract.AsyncModifyPendingCertificate(handler, &_CertificateManagement.TransactOpts, _address, _id, _studentName, _course)
}

// RejectCertificate is a paid mutator transaction binding the contract method 0xfca400f6.
//
// Solidity: function rejectCertificate(address _address, uint256 _id) returns()
func (_CertificateManagement *CertificateManagementTransactor) RejectCertificate(opts *bind.TransactOpts, _address common.Address, _id *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _CertificateManagement.contract.TransactWithResult(opts, out, "rejectCertificate", _address, _id)
	return transaction, receipt, err
}

func (_CertificateManagement *CertificateManagementTransactor) AsyncRejectCertificate(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address, _id *big.Int) (*types.Transaction, error) {
	return _CertificateManagement.contract.AsyncTransact(opts, handler, "rejectCertificate", _address, _id)
}

// RejectCertificate is a paid mutator transaction binding the contract method 0xfca400f6.
//
// Solidity: function rejectCertificate(address _address, uint256 _id) returns()
func (_CertificateManagement *CertificateManagementSession) RejectCertificate(_address common.Address, _id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.RejectCertificate(&_CertificateManagement.TransactOpts, _address, _id)
}

func (_CertificateManagement *CertificateManagementSession) AsyncRejectCertificate(handler func(*types.Receipt, error), _address common.Address, _id *big.Int) (*types.Transaction, error) {
	return _CertificateManagement.Contract.AsyncRejectCertificate(handler, &_CertificateManagement.TransactOpts, _address, _id)
}

// RejectCertificate is a paid mutator transaction binding the contract method 0xfca400f6.
//
// Solidity: function rejectCertificate(address _address, uint256 _id) returns()
func (_CertificateManagement *CertificateManagementTransactorSession) RejectCertificate(_address common.Address, _id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CertificateManagement.Contract.RejectCertificate(&_CertificateManagement.TransactOpts, _address, _id)
}

func (_CertificateManagement *CertificateManagementTransactorSession) AsyncRejectCertificate(handler func(*types.Receipt, error), _address common.Address, _id *big.Int) (*types.Transaction, error) {
	return _CertificateManagement.Contract.AsyncRejectCertificate(handler, &_CertificateManagement.TransactOpts, _address, _id)
}

// CertificateManagementCertificateApplied represents a CertificateApplied event raised by the CertificateManagement contract.
type CertificateManagementCertificateApplied struct {
	Id             *big.Int
	StudentAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// WatchCertificateApplied is a free log subscription operation binding the contract event 0xbaa64ac25db7cce96af0e1f049feb7b9e6c9c462e18c47e1259f79d7174585dc.
//
// Solidity: event CertificateApplied(uint256 indexed id, address indexed studentAddress)
func (_CertificateManagement *CertificateManagementFilterer) WatchCertificateApplied(fromBlock *int64, handler func(int, []types.Log), id *big.Int, studentAddress common.Address) (string, error) {
	return _CertificateManagement.contract.WatchLogs(fromBlock, handler, "CertificateApplied", id, studentAddress)
}

func (_CertificateManagement *CertificateManagementFilterer) WatchAllCertificateApplied(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _CertificateManagement.contract.WatchLogs(fromBlock, handler, "CertificateApplied")
}

// ParseCertificateApplied is a log parse operation binding the contract event 0xbaa64ac25db7cce96af0e1f049feb7b9e6c9c462e18c47e1259f79d7174585dc.
//
// Solidity: event CertificateApplied(uint256 indexed id, address indexed studentAddress)
func (_CertificateManagement *CertificateManagementFilterer) ParseCertificateApplied(log types.Log) (*CertificateManagementCertificateApplied, error) {
	event := new(CertificateManagementCertificateApplied)
	if err := _CertificateManagement.contract.UnpackLog(event, "CertificateApplied", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchCertificateApplied is a free log subscription operation binding the contract event 0xbaa64ac25db7cce96af0e1f049feb7b9e6c9c462e18c47e1259f79d7174585dc.
//
// Solidity: event CertificateApplied(uint256 indexed id, address indexed studentAddress)
func (_CertificateManagement *CertificateManagementSession) WatchCertificateApplied(fromBlock *int64, handler func(int, []types.Log), id *big.Int, studentAddress common.Address) (string, error) {
	return _CertificateManagement.Contract.WatchCertificateApplied(fromBlock, handler, id, studentAddress)
}

func (_CertificateManagement *CertificateManagementSession) WatchAllCertificateApplied(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _CertificateManagement.Contract.WatchAllCertificateApplied(fromBlock, handler)
}

// ParseCertificateApplied is a log parse operation binding the contract event 0xbaa64ac25db7cce96af0e1f049feb7b9e6c9c462e18c47e1259f79d7174585dc.
//
// Solidity: event CertificateApplied(uint256 indexed id, address indexed studentAddress)
func (_CertificateManagement *CertificateManagementSession) ParseCertificateApplied(log types.Log) (*CertificateManagementCertificateApplied, error) {
	return _CertificateManagement.Contract.ParseCertificateApplied(log)
}

// CertificateManagementCertificateApproved represents a CertificateApproved event raised by the CertificateManagement contract.
type CertificateManagementCertificateApproved struct {
	Id       *big.Int
	Approver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchCertificateApproved is a free log subscription operation binding the contract event 0x1ffd4b0c43de7f13bb9e8c45b1b35af572231837bf8a3c876e9eaf96a3cbd3d9.
//
// Solidity: event CertificateApproved(uint256 indexed id, address indexed approver)
func (_CertificateManagement *CertificateManagementFilterer) WatchCertificateApproved(fromBlock *int64, handler func(int, []types.Log), id *big.Int, approver common.Address) (string, error) {
	return _CertificateManagement.contract.WatchLogs(fromBlock, handler, "CertificateApproved", id, approver)
}

func (_CertificateManagement *CertificateManagementFilterer) WatchAllCertificateApproved(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _CertificateManagement.contract.WatchLogs(fromBlock, handler, "CertificateApproved")
}

// ParseCertificateApproved is a log parse operation binding the contract event 0x1ffd4b0c43de7f13bb9e8c45b1b35af572231837bf8a3c876e9eaf96a3cbd3d9.
//
// Solidity: event CertificateApproved(uint256 indexed id, address indexed approver)
func (_CertificateManagement *CertificateManagementFilterer) ParseCertificateApproved(log types.Log) (*CertificateManagementCertificateApproved, error) {
	event := new(CertificateManagementCertificateApproved)
	if err := _CertificateManagement.contract.UnpackLog(event, "CertificateApproved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchCertificateApproved is a free log subscription operation binding the contract event 0x1ffd4b0c43de7f13bb9e8c45b1b35af572231837bf8a3c876e9eaf96a3cbd3d9.
//
// Solidity: event CertificateApproved(uint256 indexed id, address indexed approver)
func (_CertificateManagement *CertificateManagementSession) WatchCertificateApproved(fromBlock *int64, handler func(int, []types.Log), id *big.Int, approver common.Address) (string, error) {
	return _CertificateManagement.Contract.WatchCertificateApproved(fromBlock, handler, id, approver)
}

func (_CertificateManagement *CertificateManagementSession) WatchAllCertificateApproved(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _CertificateManagement.Contract.WatchAllCertificateApproved(fromBlock, handler)
}

// ParseCertificateApproved is a log parse operation binding the contract event 0x1ffd4b0c43de7f13bb9e8c45b1b35af572231837bf8a3c876e9eaf96a3cbd3d9.
//
// Solidity: event CertificateApproved(uint256 indexed id, address indexed approver)
func (_CertificateManagement *CertificateManagementSession) ParseCertificateApproved(log types.Log) (*CertificateManagementCertificateApproved, error) {
	return _CertificateManagement.Contract.ParseCertificateApproved(log)
}

// CertificateManagementCertificateRejected represents a CertificateRejected event raised by the CertificateManagement contract.
type CertificateManagementCertificateRejected struct {
	Id       *big.Int
	Approver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchCertificateRejected is a free log subscription operation binding the contract event 0xd16874d1c78ac14ec0cc9028724b5f14c02545485fb5e0bdf58457d3fe3c1946.
//
// Solidity: event CertificateRejected(uint256 indexed id, address indexed approver)
func (_CertificateManagement *CertificateManagementFilterer) WatchCertificateRejected(fromBlock *int64, handler func(int, []types.Log), id *big.Int, approver common.Address) (string, error) {
	return _CertificateManagement.contract.WatchLogs(fromBlock, handler, "CertificateRejected", id, approver)
}

func (_CertificateManagement *CertificateManagementFilterer) WatchAllCertificateRejected(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _CertificateManagement.contract.WatchLogs(fromBlock, handler, "CertificateRejected")
}

// ParseCertificateRejected is a log parse operation binding the contract event 0xd16874d1c78ac14ec0cc9028724b5f14c02545485fb5e0bdf58457d3fe3c1946.
//
// Solidity: event CertificateRejected(uint256 indexed id, address indexed approver)
func (_CertificateManagement *CertificateManagementFilterer) ParseCertificateRejected(log types.Log) (*CertificateManagementCertificateRejected, error) {
	event := new(CertificateManagementCertificateRejected)
	if err := _CertificateManagement.contract.UnpackLog(event, "CertificateRejected", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchCertificateRejected is a free log subscription operation binding the contract event 0xd16874d1c78ac14ec0cc9028724b5f14c02545485fb5e0bdf58457d3fe3c1946.
//
// Solidity: event CertificateRejected(uint256 indexed id, address indexed approver)
func (_CertificateManagement *CertificateManagementSession) WatchCertificateRejected(fromBlock *int64, handler func(int, []types.Log), id *big.Int, approver common.Address) (string, error) {
	return _CertificateManagement.Contract.WatchCertificateRejected(fromBlock, handler, id, approver)
}

func (_CertificateManagement *CertificateManagementSession) WatchAllCertificateRejected(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _CertificateManagement.Contract.WatchAllCertificateRejected(fromBlock, handler)
}

// ParseCertificateRejected is a log parse operation binding the contract event 0xd16874d1c78ac14ec0cc9028724b5f14c02545485fb5e0bdf58457d3fe3c1946.
//
// Solidity: event CertificateRejected(uint256 indexed id, address indexed approver)
func (_CertificateManagement *CertificateManagementSession) ParseCertificateRejected(log types.Log) (*CertificateManagementCertificateRejected, error) {
	return _CertificateManagement.Contract.ParseCertificateRejected(log)
}
