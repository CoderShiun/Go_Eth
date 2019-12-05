package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math"
	"math/big"
)

/*func connectClient() *ethclient.Client {
	client, err := ethclient.Dial("/home/shiun/Ethereum/Pri_Air00/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}


	//fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
	return client
}*/

func getBalance() {
	/*
	//传区块号能让您读取该区块时的账户余额。区块号必须是big.Int类型
	blockNumber := big.NewInt(5532993)
*/
	//nil = the last block
	account := common.HexToAddress("0xa58b752d895c8365cda6a5e43586ef4661f7a9c1")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	//convert wei to ETH
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("Balance is: ", ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fpendingBalance := new(big.Float)
	fpendingBalance.SetString(pendingBalance.String())
	ethPendingValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Print("Pending balance is: ", ethPendingValue)
}

func getTokenBalance(contractAddress, holderAddress string) {
	tokenBalance, err := rinkbyClient.TokenBalance(contractAddress, holderAddress)
	if err != nil{
		log.Panic(err)
	}

	fmt.Println(tokenBalance.Int())
}