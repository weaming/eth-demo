package main

import (
	"flag"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	keystoreFile := flag.String("keystore", "", "keystore file path")
	password := flag.String("password", "", "password")
	flag.Parse()

	keyjson, e := ioutil.ReadFile(*keystoreFile)
	if e != nil {
		panic(e)
	}

	key, e := keystore.DecryptKey(keyjson, *password)
	if e != nil {
		panic(e)
	}

	e = crypto.SaveECDSA("/dev/stdout", key.PrivateKey)
	if e != nil {
		panic(e)
	}
}
