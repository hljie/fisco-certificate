package sdk

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"

	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var Contract *CertificateManagementSession

func InitContract() {
	// 连接fisco 并初始化合约
	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	// disable ssl of node rpc
	config := &client.Config{IsSMCrypto: false, GroupID: "group0", DisableSsl: false,
		PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./sdk/ca.crt", TLSKeyFile: "./sdk/sdk.key", TLSCertFile: "./sdk/sdk.crt"}
	client, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xC8eAd4B26b2c6Ac14c9fD90d9684c9Bc2cC40085")
	instance, err := NewCertificateManagement(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	Contract = &CertificateManagementSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
}

func GenAddress() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
		return "", ""
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return "", ""
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return hexutil.Encode(privateKeyBytes)[2:], address
}
