package jsonrpc_client

import (
	"encoding/json"
	"math/big"
	"strconv"
)

type Transaction struct {
	BlockHash        *string  `json:"block_hash"`
	BlockNumber      *int     `json:"block_number"`
	From             string   `json:"from"`
	Gas              int      `json:"gas"`
	GasPrice         *big.Int `json:"gas_price"`
	Hash             string   `json:"hash"`
	Input            string   `json:"input"`
	Nonce            int      `json:"nonce"`
	R                string   `json:"r"`
	S                string   `json:"s"`
	To               *string  `json:"to"`
	TransactionIndex *int     `json:"transaction_index"`
	V                int      `json:"v"`
	Value            *big.Int `json:"value"`

	// Parity only
	Creates   *string `json:"creates"`
	NetworkId *int    `json:"network_id"`
	PublicKey *string `json:"public_key"`
	Raw       *string `json:"raw"`
	StandardV *int    `json:"standard_v"`
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

	// pointers
	var blockHash, blockNumber, to, transactionIndex *string
	if tx.BlockHash != nil {
		blockHashString := *tx.BlockHash // store our own copy
		blockHash = &blockHashString
	}
	if tx.BlockNumber != nil {
		blockNumberString := "0x" + strconv.FormatInt(int64(*tx.BlockNumber), 16)
		blockNumber = &blockNumberString
	}
	if tx.To != nil {
		toString := *tx.To // store our own copy
		to = &toString
	}
	if tx.TransactionIndex != nil {
		transactionIndexString := "0x" + strconv.FormatInt(int64(*tx.TransactionIndex), 16)
		transactionIndex = &transactionIndexString
	}

	gas := "0x" + strconv.FormatInt(int64(tx.Gas), 16)
	gasPrice := "0x" + tx.GasPrice.Text(16)
	nonce := "0x" + strconv.FormatInt(int64(tx.Nonce), 16)
	v := "0x" + strconv.FormatInt(int64(tx.V), 16)
	value := "0x" + tx.Value.Text(16)

	// Parity only
	var creates, publicKey, standardV, raw *string
	var networkId *int
	if tx.Creates != nil {
		createsString := *tx.Creates
		creates = &createsString
	}
	if tx.NetworkId != nil {
		networkIdString := *tx.NetworkId
		networkId = &networkIdString
	}
	if tx.PublicKey != nil {
		publicKeyString := *tx.PublicKey
		publicKey = &publicKeyString
	}
	if tx.Raw != nil {
		rawString := *tx.Raw
		raw = &rawString
	}
	if tx.StandardV != nil {
		standardVString := "0x" + strconv.FormatInt(int64(*tx.StandardV), 16)
		standardV = &standardVString
	}

	txResult := TransactionResult{
		BlockHash:        blockHash,
		BlockNumber:      blockNumber,
		From:             tx.From,
		Gas:              gas,
		GasPrice:         gasPrice,
		Hash:             tx.Hash,
		Input:            tx.Input,
		Nonce:            nonce,
		R:                tx.R,
		S:                tx.S,
		To:               to,
		TransactionIndex: transactionIndex,
		V:                v,
		Value:            value,

		// Parity only
		Creates:   creates,
		NetworkId: networkId,
		PublicKey: publicKey,
		Raw:       raw,
		StandardV: standardV,
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

// Equals determines whether two Transactions are equal
func (tx *Transaction) Equals(tx2 *Transaction) bool {

	if tx.From != tx2.From ||
		tx.Gas != tx2.Gas ||
		tx.Hash != tx2.Hash ||
		tx.Input != tx2.Input ||
		tx.Nonce != tx2.Nonce ||
		tx.R != tx2.R ||
		tx.S != tx2.S ||
		tx.V != tx2.V {
		return false
	}

	// big integers
	if !AreEqualBigInt(tx.GasPrice, tx2.GasPrice) ||
		!AreEqualBigInt(tx.Value, tx2.Value) {
		return false
	}

	// confirmed tx
	if !AreEqualString(tx.BlockHash, tx2.BlockHash) ||
		!AreEqualInt(tx.BlockNumber, tx2.BlockNumber) ||
		!AreEqualInt(tx.TransactionIndex, tx2.TransactionIndex) {
		return false
	}

	// null for contract creation
	if !AreEqualString(tx.To, tx2.To) {
		return false
	}

	// Parity only
	if !AreEqualString(tx.Creates, tx2.Creates) ||
		!AreEqualInt(tx.NetworkId, tx2.NetworkId) ||
		!AreEqualString(tx.PublicKey, tx2.PublicKey) ||
		!AreEqualString(tx.Raw, tx2.Raw) ||
		!AreEqualInt(tx.StandardV, tx2.StandardV) {
		return false
	}

	return true
}
