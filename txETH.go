package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"strconv"
)

func txETH(to, amount string) {
	privateKey, _, fromAddress := getKeys()

	//获得帐户的随机数(nonce)。 每笔交易都需要一个nonce。 根据定义，nonce是仅使用一次的数字。
	// 如果是发送交易的新帐户，则该随机数将为“0”。来自帐户的每个新事务都必须具有前一个nonce增加1的nonce。
	// 很难对所有nonce进行手动跟踪，于是ethereum客户端提供一个帮助方法PendingNonceAt，它将返回你应该使用的下一个nonce。
	//该函数需要我们发送的帐户的公共地址 - 这个我们可以从私钥派生。
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	nonce, err := rinkbyClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//设置我们将要转移的ETH数量。 但是我们必须将ETH转换为wei

	amountF64, err := strconv.ParseFloat(amount,64)
	calAmount := int64(amountF64*1000000000000000000)
	txAmount := big.NewInt(calAmount)

	//ETH转账的燃气应设上限为“21000”单位。
	gasLimit := uint64(210000)

	//燃气价格必须以wei为单位设定
	gasPrice := big.NewInt(30000000000) //30 wei

	/*
	//对燃气价格进行硬编码有时并不理想。
	// go-ethereum客户端提供SuggestGasPrice函数，用于根据'x'个先前块来获得平均燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	*/

	//send to the address
	toAddress := common.HexToAddress(to)

	// 认证信息组装
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = txAmount     // in wei
	auth.GasLimit = gasLimit // in units
	auth.GasPrice = gasPrice
	auth.From = fromAddress

	data := "Try first ABC 12321s."

	//通过导入go-ethereumcore/types包并调用NewTransaction来生成我们的未签名以太坊事务，
	// 这个函数需要接收nonce，地址，值，燃气上限值，燃气价格和可选发的数据
	tx := types.NewTransaction(nonce, toAddress, txAmount, gasLimit, gasPrice, []byte(data))
	//tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	//使用发件人的私钥对事务进行签名。 为此，我们调用SignTx方法，该方法接受一个未签名的事务和我们之前构造的私钥。
	// SignTx方法需要EIP155签名者，这个也需要我们先从客户端拿到链ID
	/*chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}*/

	signedTx ,err:= auth.Signer(types.HomesteadSigner{}, auth.From, tx)
	//signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//调用“SendTransaction”来将已签名的事务广播到整个网络
	//err = client.SendTransaction(context.Background(), signedTx)
	err = rinkbyClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	// 等待挖矿完成
	//bind.WaitMined(context.Background(),client,signedTx)
}