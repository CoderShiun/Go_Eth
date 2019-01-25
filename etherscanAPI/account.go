package main

import (
	"fmt"
	"log"
)

func getBalance_oneAddress(address, apiKey string) {
	url := "https://api-rinkeby.etherscan.io/api?module=account&action=balance&address=" + address + "&tag=latest&apikey=" + apiKey
	msg, err := getMessage(url)
	if err != nil{
		log.Panic(err)
	}

	fmt.Println("Status", msg.Status)
	fmt.Println("Message", msg.Message)
	fmt.Println("Result", msg.Result)
}

func getTxList(address, apiKey string) {
	url := "https://api-rinkeby.etherscan.io/api?module=account&action=txlist&address=" + address + "&startblock=0&endblock=99999999&page=1&offset=10&sort=asc&apikey=" + apiKey
	msgs, err := getMessages(url)
	if err != nil{
		log.Panic(err)
	}

	fmt.Println("Status", msgs.Status)
	fmt.Println("Message", msgs.Message)

	for i := 0; i < len(msgs.Result) ; i++ {
		fmt.Println("Result: ", msgs.Result[i])
	}
}

func getInternalTx(address, apiKey string) {
	url := "https://api-rinkeby.etherscan.io/api?module=account&action=txlistinternal&address="+address+"&startblock=0&endblock=2702578&page=1&offset=10&sort=asc&apikey="+apiKey
	msgs, err := getMessages(url)
	if err != nil{
		log.Panic(err)
	}

	fmt.Println("Status", msgs.Status)
	fmt.Println("Message", msgs.Message)

	for i := 0; i < len(msgs.Result) ; i++ {
		fmt.Println("Result: ", msgs.Result[i])
	}
}

func getInternalTxByHash(hash, apiKey string) {
	url := "https://api-rinkeby.etherscan.io/api?module=account&action=txlistinternal&txhash="+hash+"&apikey="+apiKey
	msgs, err := getMessages(url)
	if err != nil{
		log.Panic(err)
	}

	fmt.Println("Status", msgs.Status)
	fmt.Println("Message", msgs.Message)

	for i := 0; i < len(msgs.Result) ; i++ {
		fmt.Println("Result: ", msgs.Result[i])
	}
}

func getERC20Tx(contractAddress, apiKey string) {
	url := "https://api-rinkeby.etherscan.io/api?module=account&action=tokentx&contractaddress="+contractAddress+"&page=1&offset=100&sort=asc&apikey="+ apiKey
	msgs, err := getMessages(url)
	if err != nil{
		log.Panic(err)
	}

	fmt.Println("Status", msgs.Status)
	fmt.Println("Message", msgs.Message)

	for i := 0; i < len(msgs.Result) ; i++ {
		fmt.Println("Result: ", msgs.Result[i])
	}
}

func getBlocksMinedList(address, apiKey string) {
	url := "https://api-rinkeby.etherscan.io/api?module=account&action=getminedblocks&address="+address+"&blocktype=blocks&page=1&offset=10&apikey="+apiKey
	msgs, err := getMessages(url)
	if err != nil{
		log.Panic(err)
	}

	fmt.Println("Status", msgs.Status)
	fmt.Println("Message", msgs.Message)

	for i := 0; i < len(msgs.Result) ; i++ {
		fmt.Println("Result: ", msgs.Result[i])
	}
}
