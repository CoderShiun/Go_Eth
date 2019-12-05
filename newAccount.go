package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func createKeys() {
	//the key will store in the path below
	ks := keystore.NewKeyStore("/home/shiun/Download", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "default"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

//调用Import方法，该方法接收keystore的JSON数据作为字节。第二个参数是用于加密私钥的口令。第三个参数是指定一个新的加密口令，
// 但在示例中使用一样的口令。导入账户将允许您按期访问该账户，但它将生成新keystore文件, 所以删除旧的
func importKeys() {
	file := "/home/shiun/Ethereum/Pri_Air00/keystore/UTC--2018-12-18T12-14-12.950275484Z--5216b4ffa892300ae930419e6ba26e9b091ffe9b"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "default"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}