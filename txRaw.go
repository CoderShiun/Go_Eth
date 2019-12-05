package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"io/ioutil"
	"log"
	"math/big"
)

func txRaw() []byte {
	fromKeystore,err := ioutil.ReadFile("/home/shiun/Ethereum/Pri_Air00/keystore/UTC--2018-11-18T00-01-44.834373565Z--a58b752d895c8365cda6a5e43586ef4661f7a9c1")
	if err != nil{
		log.Fatal(err)
	}
	fromKey,err := keystore.DecryptKey(fromKeystore,"air00")
	privateKey := fromKey.PrivateKey
	publicKey := privateKey.PublicKey
	fromAddress := crypto.PubkeyToAddress(publicKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//设置我们将要转移的ETH数量。 但是我们必须将ETH转换为wei
	amount := big.NewInt(1000000000000000000)

	//ETH转账的燃气应设上限为“21000”单位。
	gasLimit := uint64(21000)

	//燃气价格必须以wei为单位设定
	gasPrice := big.NewInt(30000000000) //30 wei

	toAddress := common.HexToAddress("0x5f36247e4f1e5160d6980c4828bafb57ae450d2d")
	var data []byte

	//首先构造事务对象并对其进行签名
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//以原始字节格式获取事务之前，我们需要初始化一个types.Transactions类型，并将签名后的交易作为第一个值。
	ts := types.Transactions{signedTx}

	//原因是因为Transactions类型提供了一个GetRlp方法，用于以RLP编码格式返回事务。
	// RLP是以太坊用于序列化对象的特殊编码方法。 结果是原始字节。
	rawTxBytes := ts.GetRlp(0)

	//将原始字节转换为十六进制字符串。
	rawTxHex := hex.EncodeToString(rawTxBytes)

	fmt.Printf(rawTxHex)

	return rawTxBytes
}

/**
将其广播到以太坊网络，以便最终被处理和被矿工打包到区块。
 */
func sendRaw() {
	//首先将原始事务十六进制解码为字节格式。
	//rawTx := "f86d8202b28477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ca05924bde7ef10aa88db9c66dd4f5fb16b46dff2319b9968be983118b57bb50562a001b24b31010004f13d9a26b320845257a6cfc2bf819a3d55e3fc86263c5f0772"

	//rawTxBytes, err := hex.DecodeString(rawTx)

	rawTxBytes := txRaw()

	//初始化一个新的types.Transaction指针并从go-ethereumrlp包中调用DecodeBytes，
	// 将原始事务字节和指针传递给以太坊事务类型。 RLP是以太坊用于序列化和反序列化数据的编码方法
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	//使用的以太坊客户端轻松地广播交易
	err := client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n tx sent: %s", tx.Hash().Hex())
}
