package jsonrpc_client

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

type BlockResult struct {
	Author           string              `json:"author"` // Parity only
	Difficulty       string              `json:"difficulty"`
	ExtraData        string              `json:"extraData"`
	GasLimit         string              `json:"gasLimit"`
	GasUsed          string              `json:"gasUsed"`
	Hash             string              `json:"hash"`
	LogsBloom        string              `json:"logsBloom"`
	Miner            string              `json:"miner"`
	MixHash          string              `json:"mixHash"`
	Nonce            string              `json:"nonce"`
	Number           string              `json:"number"`
	ParentHash       string              `json:"parentHash"`
	ReceiptsRoot     string              `json:"receiptsRoot"`
	SealFields       []string            `json:"sealFields"` // Parity only
	SHA3Uncles       string              `json:"sha3Uncles"`
	Size             string              `json:"size"`
	StateRoot        string              `json:"stateRoot"`
	Timestamp        string              `json:"timestamp"`
	TotalDifficulty  string              `json:"totalDifficulty"`
	Transactions     []TransactionResult `json:"transactions"`
	TransactionsRoot string              `json:"transactionsRoot"`
	Uncles           []string            `json:"uncles"`
}

// ToBlock converts a BlockResult to a Block
func (blockResult *BlockResult) ToBlock() (*Block, error) {
	// string-to-integer conversions
	difficulty, err := strconv.ParseInt(blockResult.Difficulty, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("ToBlock Difficulty: %v", err)
	}

	gasLimit, err := strconv.ParseInt(blockResult.GasLimit, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToBlock GasLimit: %v", err)
	}

	gasUsed, err := strconv.ParseInt(blockResult.GasUsed, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToBlock GasUsed: %v", err)
	}

	nonce := new(big.Int)
	nonce.SetString(blockResult.Nonce, 0)

	number, err := strconv.ParseInt(blockResult.Number, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToBlock Number: %v", err)
	}

	size, err := strconv.ParseInt(blockResult.Size, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToBlock Size: %v", err)
	}

	timestamp, err := strconv.ParseInt(blockResult.Timestamp, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToBlock Timestamp: %v", err)
	}

	totalDifficulty := new(big.Int)
	totalDifficulty.SetString(blockResult.TotalDifficulty, 0)

	block := Block{
		Author:          blockResult.Author,
		Difficulty:      difficulty,
		ExtraData:       blockResult.ExtraData,
		GasLimit:        int(gasLimit),
		GasUsed:         int(gasUsed),
		Hash:            blockResult.Hash,
		LogsBloom:       blockResult.LogsBloom,
		Miner:           blockResult.Miner,
		MixHash:         blockResult.MixHash,
		Nonce:           nonce,
		Number:          int(number),
		ParentHash:      blockResult.ParentHash,
		ReceiptsRoot:    blockResult.ReceiptsRoot,
		SealFields:      blockResult.SealFields,
		SHA3Uncles:      blockResult.SHA3Uncles,
		Size:            int(size),
		StateRoot:       blockResult.StateRoot,
		Timestamp:       int(timestamp),
		TotalDifficulty: totalDifficulty,
		// Transactions
		TransactionsRoot: blockResult.TransactionsRoot,
		Uncles:           blockResult.Uncles,
	}

	// populate the transactions in the block
	for _, resultTx := range blockResult.Transactions {
		tx, err := resultTx.ToTransaction()
		if err != nil {
			return nil, err
		}
		block.Transactions = append(block.Transactions, *tx)
	}

	return &block, nil
}

// ToJSON marshals a BlockResult into JSON
func (blockResult *BlockResult) ToJSON() ([]byte, error) {
	s, err := json.Marshal(blockResult)
	if err != nil {
		return nil, err
	}
	return s, nil
}
