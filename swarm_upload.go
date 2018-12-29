package main

//在端口“8500”上运行了一个作为背景进程的swarm节点。 接下来导入swarm包go-ethereumswearm/api/client。
// 将包装别名为bzzclient
/*
import (
	"fmt"
	bzzclient "github.com/ethereum/go-ethereum/swarm/api/client"
	"log"
)

func swarm_upload() {
	//调用NewClient函数向它传递swarm背景程序的url
	client := bzzclient.NewClient("http://127.0.0.1:8500")

	//在我们的Go应用程序中，我们将使用Swarm客户端软件包中的“Open”打开我们刚刚创建的文件。
	// 该函数将返回一个File类型，它表示swarm清单中的文件，用于上传和下载swarm内容
	file, err := bzzclient.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	//调用Upload函数，为它提供文件对象。 第二个参数是一个可选添的现有内容清单字符串，用于添加文件，否则它将为我们创建。
	// 第三个参数是我们是否希望我们的数据被加密。
	//返回的哈希值是文件的内容清单的哈希值，其中包含hello.txt文件作为其唯一条目。
	// 默认情况下，主要内容和清单都会上传。 清单确保您可以使用正确的mime类型检索文件。
	manifestHash, err := client.Upload(file, "", false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(manifestHash)

	//然后我们就可以在这里查看上传的文件 bzz://manifestHash
}
*/