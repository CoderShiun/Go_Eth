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
	fromKeystore,err := ioutil.ReadFile("/home/shiun/.ethereum/testnet/keystore/UTC--2018-11-21T22-08-22.991819776Z--18d9052e5191527d1dfab77dc6fa108c62d8f232")
	if err != nil{
		log.Fatal(err)
	}
	fromKey, err := keystore.DecryptKey(fromKeystore,"mxc01")
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