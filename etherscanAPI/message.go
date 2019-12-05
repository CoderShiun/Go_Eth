package main

type Message struct {
	Status  string
	Message string
	Result  string
}

type Messages struct {
	Status  string
	Message string
	Result  []result
}

type result struct {
	BlockNumber       string
	TimeStamp         string
	Hash              string
	Nonce             string
	BlockHash         string
	TransactionIndex  string
	From              string
	To                string
	Value             string
	Gas               string
	GasPrice          string
	IsError           string
	Txreceipt_status  string
	Input             string
	ContractAddress   string
	CumulativeGasUsed string
	GasUsed           string
	Confirmations     string
	TokenName         string
	TokenSymbol       string
	TokenDecimal      string
}

//type Messages []Message
