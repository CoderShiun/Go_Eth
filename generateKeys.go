package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"log"
)

func generateKeys() {
	//要首先生成一个新的钱包，我们需要导入go-ethereumcrypto包，
	// 该包提供用于生成随机私钥的GenerateKey方法。
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//然后我们可以通过导入golangcrypto/ecdsa包并使用FromECDSA方法将其转换为字节。
	privateKeyBytes := crypto.FromECDSA(privateKey)

	//我们现在可以使用go-ethereumhexutil包将它转换为十六进制字符串，
	// 该包提供了一个带有字节切片的Encode方法。 然后我们在十六进制编码之后删除“0x”。
	fmt.Println("Private Key:\n", hexutil.Encode(privateKeyBytes)[2:])

	//由于公钥是从私钥派生的，因此go-ethereum的加密私钥具有一个返回公钥的Public方法。
	publicKey := privateKey.Public()

	//将其转换为十六进制的过程与我们使用转化私钥的过程类似。
	//我们剥离了0x和前2个字符04，它始终是EC前缀，不是必需的。
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public Key:\n", hexutil.Encode(publicKeyBytes)[4:])

	//现在我们拥有公钥，就可以轻松生成你经常看到的公共地址。
	// 为了做到这一点，go-ethereum加密包有一个PubkeyToAddress方法，
	// 它接受一个ECDSA公钥，并返回公共地址。
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	//公共地址其实就是公钥的Keccak-256哈希，
	// 然后我们取最后40个字符（20个字节）并用“0x”作为前缀。
	// 以下是使用go-ethereum的crypto/sha3 Keccak256函数手动完成的方法。
	hash := sha3.NewKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("Public Key:\n", hexutil.Encode(hash.Sum(nil)[12:]))
}
