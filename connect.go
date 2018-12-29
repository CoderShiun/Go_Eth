package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
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
