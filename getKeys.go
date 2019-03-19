package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
)

func getPriKeyFromKeystore() (*ecdsa.PrivateKey, *ecdsa.PublicKey, common.Address){
	fromKeystore,err := ioutil.ReadFile("/home/shiun/.ethereum/rinkeby/keystore/UTC--2019-01-17T15-05-49.414272439Z--9f8cfcab0f63a06c455c848cc617912a35e8806e")
	if err != nil{
		log.Fatal(err)
	}
	fromKey, err := keystore.DecryptKey(fromKeystore,"mxctest00")
	if err != nil {
		log.Fatal(err)
	}
	privateKey := fromKey.PrivateKey
	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	fmt.Println("Pri :", privateKey)
	fmt.Println("Pub :", publicKey)
	fmt.Println("Address :", fromAddress)

	fmt.Println()
	fmt.Printf("%x", privateKey.D.String())// key.PrivateKey.D.Bytes())
	fmt.Println()
	fmt.Println("Pri is: ", privateKey.D.Bytes())
	fmt.Printf("%x", privateKey.D.Bytes())
	fmt.Println()
	fmt.Println("Pri is:", crypto.FromECDSA(privateKey))
	fmt.Printf("%x", crypto.FromECDSA(privateKey))

	return privateKey, &publicKey, fromAddress
}