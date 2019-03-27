package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/crypto/sha3"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

func txToken(to, amount string) {
	//代币传输不需要传输ETH，因此将交易“值”设置为“0”。
	value := big.NewInt(0)

	//将要发送代币的地址存储在变量中。
	toAddress := common.HexToAddress(to)

	//Contract address
	tokenAddress := common.HexToAddress("0xD4b4AE9EB383F194fc7b8c428a4c47D36A6d2540")

	//函数名将是传递函数的名称，即ERC-20规范中的transfer和参数类型。 第一个参数类型是address（令牌的接收者），
	// 第二个类型是uint256（要发送的代币数量）。 不需要没有空格和参数名称。 我们还需要用字节切片格式
	transferFnSignature := []byte("transfer(address,uint256)")

	//从go-ethereum导入crypto/sha3包以生成函数签名的Keccak256哈希。 然后我们只使用前4个字节来获取方法ID
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	//将给我们发送代币的地址左填充到32字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	//发送多少个代币，在这里是1个，并且我们需要在big.Int中格式化为wei (+18 zero)
	txAmount := new(big.Int)
	txAmount.SetString(amount + "000000000000000000", 10)


	//代币量也需要左填充到32个字节
	paddedAmount := common.LeftPadBytes(txAmount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	//将方法ID，填充后的地址和填后的转账量，接到将成为我们数据字段的字节片
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//燃气上限制将取决于交易数据的大小和智能合约必须执行的计算步骤。客户端提供了EstimateGas方法，
	// 它可以为我们估算所需的燃气量。 这个函数从ethereum包中获取CallMsg结构，我们在其中指定数据和地址。
	// 它将返回我们估算的完成交易所需的估计燃气上限

	gasLimit, err := rinkbyClient.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gaslimit: ", gasLimit)

	gasPrice, err := rinkbyClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasprice", gasPrice)

	privateKey, _, fromAddress := getKeys()
	nonce, err := rinkbyClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce: ", nonce)

	//构建交易事务类型，这类似于您在ETH转账部分中看到的，除了to字段将是代币智能合约地址。
	// 这个常让人困惑。我们还必须在调用中包含0 ETH的值字段和刚刚生成的数据字节
	tx := types.NewTransaction(nonce, tokenAddress, value, 137005, gasPrice, data)

	//使用发件人的私钥对事务进行签名。 SignTx方法需要EIP155签名器(EIP155 signer)，这需要我们从客户端拿到链ID
	chainID, err := rinkbyClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//广播交易
	err = rinkbyClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

}