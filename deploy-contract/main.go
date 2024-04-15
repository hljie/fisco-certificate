package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/FISCO-BCOS/go-sdk/v3/client"
)

func main() {
	privateKey, _ := hex.DecodeString("145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58")
	// disable ssl of node rpc
	config := &client.Config{IsSMCrypto: false, GroupID: "group0", DisableSsl: false,
		PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./conf/ca.crt", TLSKeyFile: "./conf/sdk.key", TLSCertFile: "./conf/sdk.crt"}
	client, err := client.DialContext(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=================DeployHelloWorld===============")
	address, receipt, _, err := DeployCertificateManagement(client.GetTransactOpts(), client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved, will use in next example
	fmt.Println("transaction hash: ", receipt.TransactionHash)

}
