package main

var client = connectClient()
var rinkbyClient = connectRinkbyClient()

type rinkby struct {
	ContractAddress string
	From string
	To string
}

var(
	rink_contractAddress = "0xD4b4AE9EB383F194fc7b8c428a4c47D36A6d2540"
	rink_from = "0x79e9AD2b2cDC815dE93a02EC48c94F88a27FCE86"
	rink_to = "0x9f8CFcAb0f63A06c455C848Cc617912a35e8806E"
)

func main() {
	//generateKeys()
	/*for i:=0; i<=10; i++ {
		txETH()
		fmt.Println()
	}*/
	//txETH()
	//sendRaw()
	//swarm_upload()
	//fmt.Println()
	//checkBlock()
	//subscribe()


	//checkAllBlock()

	//checkTxByBlockNo(1936)

	//readCSV2()

	//txCSV()

	//getTokenBalance("0xD4b4AE9EB383F194fc7b8c428a4c47D36A6d2540", "0x79e9AD2b2cDC815dE93a02EC48c94F88a27FCE86")

	//checkTokenTx("0xD4b4AE9EB383F194fc7b8c428a4c47D36A6d2540", "0x79e9AD2b2cDC815dE93a02EC48c94F88a27FCE86")

	//checkTokenInternalTx("0x79e9AD2b2cDC815dE93a02EC48c94F88a27FCE86")

}