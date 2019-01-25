package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"log"
	"math/big"
)

func txToken() {
	//代币传输不需要传输ETH，因此将交易“值”设置为“0”。
	value := big.NewInt(0)

	//将要发送代币的地址存储在变量中。
	toAddress := common.HexToAddress("0x5f36247e4f1e5160d6980c4828bafb57ae450d2d")

	tokenAddress := common.HexToAddress("")

	//函数名将是传递函数的名称，即ERC-20规范中的transfer和参数类型。 第一个参数类型是address（令牌的接收者），
	// 第二个类型是uint256（要发送的代币数量）。 不需要没有空格和参数名称。 我们还需要用字节切片格式
	transferFnSignature := []byte("transfer(address,uint256)")

	//从go-ethereum导入crypto/sha3包以生成函数签名的Keccak256哈希。 然后我们只使用前4个字节来获取方法ID
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	//将给我们发送代币的地址左填充到32字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	//发送多少个代币，在这里是1,000个，并且我们需要在big.Int中格式化为wei
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens

	//代币量也需要左填充到32个字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	//将方法ID，填充后的地址和填后的转账量，接到将成为我们数据字段的字节片
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//燃气上限制将取决于交易数据的大小和智能合约必须执行的计算步骤。客户端提供了EstimateGas方法，
	// 它可以为我们估算所需的燃气量。 这个函数从ethereum包中获取CallMsg结构，我们在其中指定数据和地址。
	// 它将返回我们估算的完成交易所需的估计燃气上限

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit)

	privateKey, _, fromAddress := getKeys()
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//构建交易事务类型，这类似于您在ETH转账部分中看到的，除了to字段将是代币智能合约地址。
	// 这个常让人困惑。我们还必须在调用中包含0 ETH的值字段和刚刚生成的数据字节
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	//使用发件人的私钥对事务进行签名。 SignTx方法需要EIP155签名器(EIP155 signer)，这需要我们从客户端拿到链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//广播交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}