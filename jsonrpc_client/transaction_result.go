package jsonrpc_client

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

type TransactionResult struct {
	BlockHash        *string `json:"blockHash"`   // null for pending tx
	BlockNumber      *string `json:"blockNumber"` // null for pending tx
	Creates          *string `json:"creates"`     // Parity only; null when not creating contract
	From             string  `json:"from"`
	Gas              string  `json:"gas"`
	GasPrice         string  `json:"gasPrice"`
	Hash             string  `json:"hash"`
	Input            string  `json:"input"`
	NetworkId        *int    `json:"networkId"` // Parity only
	Nonce            string  `json:"nonce"`
	PublicKey        string  `json:"publicKey"` // Parity only
	R                string  `json:"r"`
	Raw              string  `json:"raw"` // Parity only
	S                string  `json:"s"`
	StandardV        string  `json:"standardV"`        // Parity only
	To               *string `json:"to"`               // null when creating contract
	TransactionIndex *string `json:"transactionIndex"` // null for pending tx
	V                string  `json:"v"`
	Value            string  `json:"value"`
}

// ToJSON marshals a TransactionResult into JSON
func (txResult *TransactionResult) ToJSON() ([]byte, error) {
	s, err := json.Marshal(txResult)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// ToTransaction converts a TransactionResult to a Transaction
func (txResult *TransactionResult) ToTransaction() (*Transaction, error) {
	blockNumber, err := strconv.ParseInt(*txResult.BlockNumber, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToTransaction BlockNumber: %v", err)
	}
	blockNumberInt := int(blockNumber)

	gas, err := strconv.ParseInt(txResult.Gas, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToTransaction Gas: %v", err)
	}

	gasPrice := new(big.Int)
	gasPrice.SetString(txResult.GasPrice, 0)

	nonce, err := strconv.ParseInt(txResult.Nonce, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToTransaction Nonce: %v", err)
	}

	standardV, err := strconv.ParseInt(txResult.StandardV, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToTransaction StandardV: %v", err)
	}

	transactionIndex, err := strconv.ParseInt(*txResult.TransactionIndex, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToTransaction TransactionIndex: %v", err)
	}
	transactionIndexInt := int(transactionIndex)

	v, err := strconv.ParseInt(txResult.V, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToTransaction V: %v", err)
	}

	value := new(big.Int)
	value.SetString(txResult.Value, 0)

	tx := Transaction{
		BlockHash:        txResult.BlockHash,
		BlockNumber:      &blockNumberInt,
		Creates:          txResult.Creates,
		From:             txResult.From,
		Gas:              int(gas),
		GasPrice:         gasPrice,
		Hash:             txResult.Hash,
		Input:            txResult.Input,
		NetworkId:        txResult.NetworkId,
		Nonce:            int(nonce),
		PublicKey:        txResult.PublicKey,
		R:                txResult.R,
		Raw:              txResult.Raw,
		S:                txResult.S,
		StandardV:        int(standardV),
		To:               txResult.To,
		TransactionIndex: &transactionIndexInt,
		V:                int(v),
		Value:            value,
	}
	return &tx, nil
}
