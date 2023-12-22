package bindings

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type TraceResult struct {
	Action          Action      `json:"action"`
	BlockHash       common.Hash `json:"blockHash"`
	BlockNumber     uint64      `json:"blockNumber"`
	Result          Result      `json:"result"`
	Subtraces       int         `json:"subtraces"`
	TraceAddress    []int       `json:"traceAddress"`
	TransactionHash common.Hash `json:"transactionHash"`
	TransactionPos  int         `json:"transactionPosition"`
	Type            string      `json:"type"`
}

type Action struct {
	CallType string         `json:"callType"`
	From     common.Address `json:"from"`
	Gas      hexutil.Uint64 `json:"gas"`
	Input    string         `json:"input"`
	To       common.Address `json:"to"`
	Value    hexutil.Big    `json:"value"`
}

type Result struct {
	GasUsed hexutil.Uint64 `json:"gasUsed"`
	Output  string         `json:"output"`
}
