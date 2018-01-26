package jsonrpc_client

import (
	"encoding/json"
	"math/big"
	"strconv"
)

type Transaction struct {
	BlockHash        *string  `json:"block_hash"`
	BlockNumber      *int     `json:"block_number"`
	Creates          *string  `json:"creates"`
	From             string   `json:"from"`
	Gas              int      `json:"gas"`
	GasPrice         *big.Int `json:"gas_price"`
	Hash             string   `json:"hash"`
	Input            string   `json:"input"`
	NetworkId        *int     `json:"network_id"`
	Nonce            int      `json:"nonce"`
	PublicKey        string   `json:"public_key"`
	R                string   `json:"r"`
	Raw              string   `json:"raw"`
	S                string   `json:"s"`
	StandardV        *int     `json:"standard_v"`
	To               *string  `json:"to"`
	TransactionIndex *int     `json:"transaction_index"`
	V                int      `json:"v"`
	Value            *big.Int `json:"value"`
}

func NewTransactionFromJSON(b []byte) (*Transaction, error) {
	tx := Transaction{}
	err := json.Unmarshal(b, &tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

// ToTransactionResult converts a Transaction to a TransactionResult
func (tx *Transaction) ToTransactionResult() (*TransactionResult, error) {

	blockNumber := "0x" + strconv.FormatInt(int64(*tx.BlockNumber), 16)
	gas := "0x" + strconv.FormatInt(int64(tx.Gas), 16)
	gasPrice := "0x" + tx.GasPrice.Text(16)
	nonce := "0x" + strconv.FormatInt(int64(tx.Nonce), 16)

	var standardV *string
	if tx.StandardV != nil {
		*standardV = "0x" + strconv.FormatInt(int64(*tx.StandardV), 16)
	} else {
		standardV = nil
	}
	transactionIndex := "0x" + strconv.FormatInt(int64(*tx.TransactionIndex), 16)
	v := "0x" + strconv.FormatInt(int64(tx.V), 16)
	value := "0x" + tx.Value.Text(16)

	txResult := TransactionResult{
		BlockHash:        tx.BlockHash,
		BlockNumber:      &blockNumber,
		Creates:          tx.Creates,
		From:             tx.From,
		Gas:              gas,
		GasPrice:         gasPrice,
		Hash:             tx.Hash,
		Input:            tx.Input,
		NetworkId:        tx.NetworkId,
		Nonce:            nonce,
		PublicKey:        tx.PublicKey,
		R:                tx.R,
		Raw:              tx.Raw,
		S:                tx.S,
		StandardV:        standardV,
		To:               tx.To,
		TransactionIndex: &transactionIndex,
		V:                v,
		Value:            value,
	}
	return &txResult, nil
}

// ToJSON marshals a Transaction into JSON
func (tx *Transaction) ToJSON() ([]byte, error) {
	s, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}
	return s, nil
}
