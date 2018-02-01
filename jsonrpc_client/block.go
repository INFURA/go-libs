package jsonrpc_client

import (
	"encoding/json"
	"math/big"
	"strconv"
)

type Block struct {
	Author           string        `json:"author"`
	Difficulty       int64         `json:"difficulty"`
	ExtraData        string        `json:"extra_data"`
	GasLimit         int           `json:"gas_limit"`
	GasUsed          int           `json:"gas_used"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logs_bloom"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mix_hash"`
	Nonce            *big.Int      `json:"nonce"`
	Number           int           `json:"number"`
	ParentHash       string        `json:"parent_hash"`
	ReceiptsRoot     string        `json:"receipts_root"`
	SealFields       []string      `json:"seal_fields"`
	SHA3Uncles       string        `json:"sha3_uncles"`
	Size             int           `json:"size"`
	StateRoot        string        `json:"state_root"`
	Timestamp        int           `json:"timestamp"`
	TotalDifficulty  *big.Int      `json:"total_difficulty"`
	Transactions     []Transaction `json:"transactions"`
	TransactionsRoot string        `json:"transactions_root"`
	Uncles           []string      `json:"uncles"`
}

func NewBlockFromJSON(b []byte) (*Block, error) {
	block := Block{}
	err := json.Unmarshal(b, &block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

// ToBlockResult
func (block *Block) ToBlockResult() (*BlockResult, error) {

	difficulty := "0x" + strconv.FormatInt(block.Difficulty, 16)
	gasLimit := "0x" + strconv.FormatInt(int64(block.GasLimit), 16)
	gasUsed := "0x" + strconv.FormatInt(int64(block.GasUsed), 16)
	// the nonce must display the full 64 bits (16 hex characters)
	nonce := "0x" + ZeroPad(block.Nonce.Text(16), 16)
	number := "0x" + strconv.FormatInt(int64(block.Number), 16)
	size := "0x" + strconv.FormatInt(int64(block.Size), 16)
	timestamp := "0x" + strconv.FormatInt(int64(block.Timestamp), 16)
	totalDifficulty := "0x" + block.TotalDifficulty.Text(16)

	blockResult := BlockResult{
		Author:          block.Author,
		Difficulty:      difficulty,
		ExtraData:       block.ExtraData,
		GasLimit:        gasLimit,
		GasUsed:         gasUsed,
		Hash:            block.Hash,
		LogsBloom:       block.LogsBloom,
		Miner:           block.Miner,
		MixHash:         block.MixHash,
		Nonce:           nonce,
		Number:          number,
		ParentHash:      block.ParentHash,
		ReceiptsRoot:    block.ReceiptsRoot,
		SealFields:      block.SealFields,
		SHA3Uncles:      block.SHA3Uncles,
		Size:            size,
		StateRoot:       block.StateRoot,
		Timestamp:       timestamp,
		TotalDifficulty: totalDifficulty,
		// Transactions
		TransactionsRoot: block.TransactionsRoot,
		Uncles:           block.Uncles,
	}

	// populate the transactions in the block
	numTxs := len(block.Transactions)
	blockResult.Transactions = make([]TransactionResult, numTxs, numTxs)
	for i, tx := range block.Transactions {
		txResult, err := tx.ToTransactionResult()
		if err != nil {
			return nil, err
		}
		blockResult.Transactions[i] = *txResult
	}

	return &blockResult, nil
}

// ToJSON marshals a Block into JSON
func (block *Block) ToJSON() ([]byte, error) {
	s, err := json.Marshal(block)
	if err != nil {
		return nil, err
	}
	return s, nil
}
