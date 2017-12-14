package jsonrpc_client

import (
	"encoding/json"
)

type ResponseBase struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int64  `json:"id"`
}

type BlockNumberResponse struct {
	ResponseBase
	Result string `json:"result"`
}

type NewFilterResponse struct {
	ResponseBase
	Result string `json:"result"`
}

type GetFilterChangesResponse struct {
	ResponseBase
	Result []string `json:"result"`
}

type BlockResponse struct {
	ResponseBase
	Result BlockResult `json:"result"`
}

// ToJSON marshals a BlockResponse into JSON
func (blockResp *BlockResponse) ToJSON() ([]byte, error) {
	s, err := json.Marshal(blockResp)
	if err != nil {
		return nil, err
	}
	return s, nil
}

type TransactionResponse struct {
	ResponseBase
	Result TransactionResult `json:"result"`
}

type StringResponse struct {
	ResponseBase
	Result string `json:"result"`
}

type BoolResponse struct {
	ResponseBase
	Result bool `json:"result"`
}
