package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

/**
设置订阅以便在新区块被开采时获取事件
 */
func subscribe() {
	//创建一个新的通道，用于接收最新的区块头
	headers := make(chan *types.Header)

	//调用客户端的SubscribeNewHead方法，
	// 它接收我们刚创建的区块头通道，该方法将返回一个订阅对象
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	//订阅将推送新的区块头事件到我们的通道，因此可以使用一个select语句来监听新消息。
	// 订阅对象还包括一个error通道，该通道将在订阅失败时发送消息
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			//fmt.Println(header.Hash().Hex())

			//要获得该区块的完整内容，我们可以将区块头的摘要传递给客户端的BlockByHash函数
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Block hash: ",block.Hash().Hex())
			fmt.Println("Block Number: ", block.Number().Uint64())
			fmt.Println("Timestamp: ", block.Time().Uint64())
			fmt.Println("Block nonce: ", block.Nonce())
			fmt.Println("Tx: ", len(block.Transactions()))
		}
	}
}
