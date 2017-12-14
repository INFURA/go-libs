package jsonrpc_client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type JSONRPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	ID      int64         `json:"id"`
	Params  []interface{} `json:"params"`
}

// ToJSON marshals a JSONRPCRequest into JSON
func (req *JSONRPCRequest) ToJSON() ([]byte, error) {
	s, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type EthereumClient struct {
	URL string
}

// issueRequest issues the JSON-RPC request
func (client *EthereumClient) issueRequest(reqBody *JSONRPCRequest) ([]byte, error) {

	payload, err := reqBody.ToJSON()
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(string(payload))
	resp, err := http.Post(client.URL, JSON_MEDIA_TYPE, reader)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Eth_newBlockFilter calls the eth_newBlockFilter JSON-RPC method
func (client *EthereumClient) Eth_newBlockFilter() (string, error) {

	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "eth_newBlockFilter",
		Params:  nil,
	}

	body, err := client.issueRequest(&reqBody)
	if err != nil {
		return "", err
	}

	var clientResp NewFilterResponse
	err = json.Unmarshal(body, &clientResp)
	if err != nil {
		return "", err
	}

	return clientResp.Result, nil
}

// Eth_newPendingTransactionFilter calls the eth_newPendingTransactionFilter JSON-RPC method
func (client *EthereumClient) Eth_newPendingTransactionFilter() (string, error) {

	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "eth_newPendingTransactionFilter",
		Params:  nil,
	}

	body, err := client.issueRequest(&reqBody)
	if err != nil {
		return "", err
	}

	var clientResp NewFilterResponse
	err = json.Unmarshal(body, &clientResp)
	if err != nil {
		return "", err
	}

	return clientResp.Result, nil
}

// Eth_getFilterChanges calls the eth_getFilterChanges JSON-RPC method
func (client *EthereumClient) Eth_getFilterChanges(filterID string) ([]string, error) {

	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "eth_getFilterChanges",
		Params:  []interface{}{filterID},
	}

	body, err := client.issueRequest(&reqBody)
	if err != nil {
		return nil, err
	}

	var clientResp GetFilterChangesResponse
	err = json.Unmarshal(body, &clientResp)
	if err != nil {
		return nil, err
	}

	return clientResp.Result, nil
}

// Eth_getBlockByHash calls the eth_getBlockByHash JSON-RPC method
func (client *EthereumClient) Eth_getBlockByHash(blockHash string, full bool) (*Block, error) {

	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "eth_getBlockByHash",
		Params:  []interface{}{blockHash, full},
	}

	body, err := client.issueRequest(&reqBody)
	if err != nil {
		return nil, err
	}

	var clientResp BlockResponse
	err = json.Unmarshal(body, &clientResp)
	if err != nil {
		return nil, err
	}

	block, err := clientResp.Result.ToBlock()
	if err != nil {
		return nil, err
	}

	return block, nil
}

// Eth_getTransactionByHash calls the eth_getTransactionByHash JSON-RPC method
func (client *EthereumClient) Eth_getTransactionByHash(txHash string) (*Transaction, error) {

	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "eth_getTransactionByHash",
		Params:  []interface{}{txHash},
	}

	body, err := client.issueRequest(&reqBody)
	if err != nil {
		return nil, err
	}

	var clientResp TransactionResponse
	err = json.Unmarshal(body, &clientResp)
	if err != nil {
		return nil, err
	}

	tx, err := clientResp.Result.ToTransaction()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// Eth_getBlockByNumber calls the eth_getBlockByNumber JSON-RPC method
func (client *EthereumClient) Eth_getBlockByNumber(blockNumber int, full bool) (*Block, error) {

	blockNumberHex := "0x" + strconv.FormatInt(int64(blockNumber), 16)

	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{blockNumberHex, full},
	}

	body, err := client.issueRequest(&reqBody)
	if err != nil {
		return nil, err
	}

	var clientResp BlockResponse
	err = json.Unmarshal(body, &clientResp)
	if err != nil {
		return nil, err
	}

	block, err := clientResp.Result.ToBlock()
	if err != nil {
		return nil, err
	}

	return block, nil
}

// Eth_blockNumber calls the eth_blockNumber JSON-RPC method
func (client *EthereumClient) Eth_blockNumber() (int, error) {

	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
	}

	body, err := client.issueRequest(&reqBody)
	if err != nil {
		return 0, err
	}

	var clientResp BlockNumberResponse
	err = json.Unmarshal(body, &clientResp)
	if err != nil {
		return 0, err
	}

	blockNumber, err := strconv.ParseInt(clientResp.Result, 0, 32)
	if err != nil {
		return 0, err
	}

	return int(blockNumber), nil
}

// Web3_clientVersion calls the web3_clientVersion JSON-RPC method
func (client *EthereumClient) Web3_clientVersion() (string, error) {

	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "web3_clientVersion",
		Params:  []interface{}{},
	}

	body, err := client.issueRequest(&reqBody)
	if err != nil {
		return "", err
	}

	var clientResp StringResponse
	err = json.Unmarshal(body, &clientResp)
	if err != nil {
		return "", err
	}

	return clientResp.Result, nil
}

// Eth_syncing calls the eth_syncing JSON-RPC method
func (client *EthereumClient) Eth_syncing() (bool, error) {

	reqBody := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "eth_syncing",
		Params:  []interface{}{},
	}

	body, err := client.issueRequest(&reqBody)
	if err != nil {
		return false, err
	}

	var clientResp BoolResponse
	err = json.Unmarshal(body, &clientResp)
	if err != nil {
		return false, err
	}

	return clientResp.Result, nil
}
