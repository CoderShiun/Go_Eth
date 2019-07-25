package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

func checkTxByBlockNo(blockNo int64) {
	//block := getBlock(int64(getLastBlockNo()))
	block := getBlock(blockNo)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//通过调用Transactions方法来读取块中的事务，
	// 该方法返回一个Transaction类型的列表。 然后，重复遍历集合并获取有关事务的任何信息
	for _, tx := range block.Transactions() {
		fmt.Println("Tx hash: ", tx.Hash().Hex())
		fmt.Println("Tx value: ", tx.Value().String())
		fmt.Println("Tx gas: ", tx.Gas())
		fmt.Println("Tx gas price: ", tx.GasPrice().Uint64())
		fmt.Println("Tx nonce: ", tx.Nonce())
		fmt.Println("Tx data: ", string(tx.Data()))
		fmt.Println("Tx to: ", tx.To().Hex())

		//为了读取发送方的地址，我们需要在事务上调用AsMessage，它返回一个Message类型，
		// 其中包含一个返回sender（from）地址的函数。 AsMessage方法需要EIP155签名者，这个我们从客户端拿到链ID。
		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err != nil {
			fmt.Println("Sender: ", msg.From().Hex())
		}

		//每个事务都有一个收据，其中包含执行事务的结果，例如任何返回值和日志，以及为“1”（成功）或“0”（失败）的事件结果状态
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Success: ", receipt.Status) // 0 or 1
		fmt.Println(receipt.Logs)
	}
}

//在不获取块的情况下遍历事务的另一种方法是调用客户端的TransactionInBlock方法。 此方法仅接受块哈希和块内事务的索引值。
// 可以调用TransactionCount来了解块中有多少个事务
func checkTxByBlockHash() {
	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex())
	}
}

//可以使用TransactionByHash在给定具体事务哈希值的情况下直接查询单个事务。
func checkSingleTxByBlockHash() {
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println(isPending)
}

func checkTokenTx(contractAddress, address string) {
	// check ERC20 transactions from/to a specified address
	//transfers, err := ethScan.ERC20Transfers(&contractAddress, &address, nil, nil, 0, 0)
	bNo := 0
	transfers, err := ropstenEthScan.ERC20Transfers(&contractAddress, &address, &bNo, nil, 0, 0)
	if err != nil {
		log.Panic(err)
	}

	for _, tx := range transfers{
		fmt.Println("Value: ",tx.Value.Int())
		fmt.Println(tx.BlockNumber)
		fmt.Println(tx.To)
		fmt.Println(tx.From)
		fmt.Println(tx.TokenName)
	}

	fmt.Println(transfers)
}

/*
check if the tx is failed
 */
func checkTokenInternalTx(address string){
	internalTx, err := ethScan.InternalTxByAddress(address, nil,nil, 1, 100, true)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(internalTx)
}