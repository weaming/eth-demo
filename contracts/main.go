package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	keystore       = `{"address":"ec13ab5ea2a8e1c1289a3ca590ec25c6beac9198","crypto":{"cipher":"aes-128-ctr","ciphertext":"27484bc62a122c3ff359659dbcb6a2b78095ba5d86a0ba604e5371b2303a635e","cipherparams":{"iv":"f427af83a394e5890f35a41b9527863f"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ddd5c66a13c4aa00e565c15cec6728cc5acf7a5362f091da2c3eb3c3e0019905"},"mac":"c2c338d074ba495132bf3148edf5d8e4ff4b159e78655a0ab7742b304bf82d36"},"id":"f3802075-9beb-4890-9433-fc63f56c2734","version":3}`
	testAddr       = "0xEc13aB5ea2a8e1C1289A3cA590ec25c6Beac9198"
	ipcAddr        = "/Users/weaming/Library/Ethereum/goerli/geth.ipc"
	gatewayMainnet = "https://mainnet.infura.io/v3/202838a2ecd94c61b78c66ee72d82958"
	gatewayGoerli  = "https://3597fa2ca50f434da95f8133ce35432a.goerli.rpc.rivet.cloud/"
	gatewayGoerli2 = "https://eth-goerli.gateway.pokt.network/v1/lb/619f72f64df71a00392de6d9"
)

func main() {
	client, err := ethclient.Dial(gatewayGoerli)
	// client, err := ethclient.Dial(ipcAddr)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("connect success!")
	account := common.HexToAddress(testAddr)
	log.Println(account.Hex())

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(balance)

	auth, err := bind.NewTransactor(strings.NewReader(keystore), os.Getenv("PASSWORD"))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	log.Printf("Suggested gas price: %s", gasPrice)
	auth.GasPrice = gasPrice

	address, tx, instance, err := DeployStorage(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy new contract: %v", err)
	}
	log.Printf("Contract pending deploy: 0x%x\n", address)
	log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
}
