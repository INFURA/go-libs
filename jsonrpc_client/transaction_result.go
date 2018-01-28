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
	From             string  `json:"from"`
	Gas              string  `json:"gas"`
	GasPrice         string  `json:"gasPrice"`
	Hash             string  `json:"hash"`
	Input            string  `json:"input"`
	Nonce            string  `json:"nonce"`
	R                string  `json:"r"`
	S                string  `json:"s"`
	To               *string `json:"to"`               // null when creating contract
	TransactionIndex *string `json:"transactionIndex"` // null for pending tx
	V                string  `json:"v"`
	Value            string  `json:"value"`

	// Parity only
	Creates   *string `json:"creates"`   // null when not creating contract
	NetworkId *int    `json:"networkId"` // null for some txs
	PublicKey *string `json:"publicKey"`
	Raw       *string `json:"raw"`
	StandardV *string `json:"standardV"`
}

func NewTransactionResultFromJSON(b []byte) (*TransactionResult, error) {
	txResult := TransactionResult{}
	err := json.Unmarshal(b, &txResult)
	if err != nil {
		return nil, err
	}
	return &txResult, nil
}

// ToTransaction converts a TransactionResult to a Transaction
func (txResult *TransactionResult) ToTransaction() (*Transaction, error) {

	// pointers
	var blockHash, to *string
	var blockNumber, transactionIndex *int
	if txResult.BlockHash != nil {
		blockHashString := *txResult.BlockHash
		blockHash = &blockHashString
	}
	if txResult.BlockNumber != nil {
		blockNumberInt64, err := strconv.ParseInt(*txResult.BlockNumber, 0, 32)
		if err != nil {
			return nil, fmt.Errorf("ToTransaction BlockNumber: %v", err)
		}
		blockNumberInt := int(blockNumberInt64)
		blockNumber = &blockNumberInt
	}
	if txResult.To != nil {
		toString := *txResult.To
		to = &toString
	}
	if txResult.TransactionIndex != nil {
		transactionIndexInt64, err := strconv.ParseInt(*txResult.TransactionIndex, 0, 32)
		if err != nil {
			return nil, fmt.Errorf("ToTransaction TransactionIndex: %v", err)
		}
		transactionIndexInt := int(transactionIndexInt64)
		transactionIndex = &transactionIndexInt
	}

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

	v, err := strconv.ParseInt(txResult.V, 0, 32)
	if err != nil {
		return nil, fmt.Errorf("ToTransaction V: %v", err)
	}

	value := new(big.Int)
	value.SetString(txResult.Value, 0)

	// Parity only
	var creates, publicKey, raw *string
	var networkId, standardV *int
	if txResult.Creates != nil {
		createsString := *txResult.Creates
		creates = &createsString
	}
	if txResult.NetworkId != nil {
		networkIdString := *txResult.NetworkId
		networkId = &networkIdString
	}
	if txResult.PublicKey != nil {
		publicKeyString := *txResult.PublicKey
		publicKey = &publicKeyString
	}
	if txResult.Raw != nil {
		rawString := *txResult.Raw
		raw = &rawString
	}
	if txResult.StandardV != nil {
		standardVInt64, err := strconv.ParseInt(*txResult.StandardV, 0, 32)
		if err != nil {
			return nil, fmt.Errorf("ToTransaction StandardV: %v", err)
		}
		standardVInt := int(standardVInt64)
		standardV = &standardVInt
	}

	tx := Transaction{
		BlockHash:        blockHash,
		BlockNumber:      blockNumber,
		From:             txResult.From,
		Gas:              int(gas),
		GasPrice:         gasPrice,
		Hash:             txResult.Hash,
		Input:            txResult.Input,
		Nonce:            int(nonce),
		R:                txResult.R,
		S:                txResult.S,
		To:               to,
		TransactionIndex: transactionIndex,
		V:                int(v),
		Value:            value,

		// Parity only
		Creates:   creates,
		NetworkId: networkId,
		PublicKey: publicKey,
		Raw:       raw,
		StandardV: standardV,
	}
	return &tx, nil
}

// ToJSON marshals a TransactionResult into JSON
func (txResult *TransactionResult) ToJSON() ([]byte, error) {
	s, err := json.Marshal(txResult)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Equals determines whether two TransactionResults are equal
func (txResult *TransactionResult) Equals(txResult2 *TransactionResult) bool {

	if txResult.From != txResult2.From ||
		txResult.Gas != txResult2.Gas ||
		txResult.GasPrice != txResult2.GasPrice ||
		txResult.Hash != txResult2.Hash ||
		txResult.Input != txResult2.Input ||
		txResult.Nonce != txResult2.Nonce ||
		txResult.R != txResult2.R ||
		txResult.S != txResult2.S ||
		txResult.V != txResult2.V ||
		txResult.Value != txResult2.Value {
		return false
	}

	// confirmed tx
	if !AreEqualString(txResult.BlockHash, txResult2.BlockHash) ||
		!AreEqualString(txResult.BlockNumber, txResult2.BlockNumber) ||
		!AreEqualString(txResult.TransactionIndex, txResult2.TransactionIndex) {
		return false
	}

	// null for contract creation
	if !AreEqualString(txResult.To, txResult2.To) {
		return false
	}

	// Parity only
	if !AreEqualString(txResult.Creates, txResult2.Creates) ||
		!AreEqualInt(txResult.NetworkId, txResult2.NetworkId) ||
		!AreEqualString(txResult.PublicKey, txResult2.PublicKey) ||
		!AreEqualString(txResult.Raw, txResult2.Raw) ||
		!AreEqualString(txResult.StandardV, txResult2.StandardV) {
		return false
	}

	return true
}
