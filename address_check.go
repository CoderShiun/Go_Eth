package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"regexp"
)

func checkAddress() {
	//We can use a simple regular expression to check if the ethereum address is valid
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	//检查地址是否为账户或智能合约
	//我们可以确定，若在该地址存储了字节码，该地址是智能合约。
	//当地址上没有字节码时，我们知道它不是一个智能合约，它是一个标准的以太坊账户。
	// 在例子中，我们获取一个代币智能合约的字节码并检查其长度以验证它是一个智能合约：
	address := common.HexToAddress("0xa58b752d895c8365cda6a5e43586ef4661f7a9c1")
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract)
}
