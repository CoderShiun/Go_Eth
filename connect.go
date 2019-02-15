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

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil{
		log.Fatal(err)
	}

	//if you run the local node
	/*
	client, err := ethclient.Dial("/home/shiun/Ethereum/Pri_Air00/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}
*/
	//fmt.Println("we have a connection")
	_ = client //??
	return client
}

func connectRinkbyClient() *ethclient.Client {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil{
		log.Fatal(err)
	}

	_ = client //??
	return client
}

func connectRinkbyEthScan() *etherscan.Client {
	//// create a API client for specified ethereum net
	tokenEthScan := etherscan.New(etherscan.Rinkby, "W8M6B92HBM7CUAQINJ8IMST29RY2ZVSQH4")

	/*
	client.BeforeRequest = func(module, action string, param map[string]interface{}) error {
		// ...
	}

	client.AfterRequest = func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error) {
		// ...
	}
	*/

	return tokenEthScan
}

func connectEthScan() *etherscan.Client {
	tokenEthScan := etherscan.New(etherscan.Mainnet, "W8M6B92HBM7CUAQINJ8IMST29RY2ZVSQH4")

	return  tokenEthScan
}

func getKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey, common.Address) {
	//加载的私钥
	// 获取私钥方式一，通过keystore文件

	fromKeystore,err := ioutil.ReadFile("/home/shiun/.ethereum/rinkeby/keystore/UTC--2018-12-04T12-10-18.166247692Z--79e9ad2b2cdc815de93a02ec48c94f88a27fce86")
	if err != nil{
		log.Fatal(err)
	}
	fromKey, err := keystore.DecryptKey(fromKeystore,"MXCtest00")
	if err != nil {
		log.Fatal(err)
	}
	privateKey := fromKey.PrivateKey
	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	return privateKey, &publicKey, fromAddress
	/*
	fmt.Println(fromKey)
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	fmt.Println(fromAddress)*/

	// 获取私钥方式二，通过私钥字符串
/*
	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, publicKeyECDSA, fromAddress
	*/
}
