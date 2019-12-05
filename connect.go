package main

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"
	"io/ioutil"
	"log"
)

func connectClient() *ethclient.Client {
	client, err := ethclient.Dial("/home/shiun/Ethereum/Pri_Air00/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
	return client
}

func connectRinkbyClient() *etherscan.Client {
	//// create a API client for specified ethereum net
	tokenClient := etherscan.New(etherscan.Rinkby, "W8M6B92HBM7CUAQINJ8IMST29RY2ZVSQH4")

	/*
	client.BeforeRequest = func(module, action string, param map[string]interface{}) error {
		// ...
	}

	client.AfterRequest = func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error) {
		// ...
	}
	*/

	return tokenClient
}

func getKeys() (*ecdsa.PrivateKey, ecdsa.PublicKey, common.Address) {
	//加载的私钥
	// 获取私钥方式一，通过keystore文件
	fromKeystore,err := ioutil.ReadFile("/home/shiun/Ethereum/Pri_Air00/keystore/UTC--2018-11-18T00-01-44.834373565Z--a58b752d895c8365cda6a5e43586ef4661f7a9c1")
	if err != nil{
		log.Fatal(err)
	}
	fromKey, err := keystore.DecryptKey(fromKeystore,"air00")
	if err != nil {
		log.Fatal(err)
	}
	privateKey := fromKey.PrivateKey
	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	/*
	fmt.Println(fromKey)
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	fmt.Println(fromAddress)*/

	// 获取私钥方式二，通过私钥字符串
	/*
	privateKey, err := crypto.HexToECDSA("64b6fdc385cb673a3105f648baaf7eeee5a63f56dd111715d67dff1cd591df4e")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	*/
	return privateKey, publicKey, fromAddress
}
