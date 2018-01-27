package jsonrpc_client

import (
	"math/big"
)

// AreEqualString
func AreEqualString(a, b *string) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return *a == *b
}

// AreEqualInt
func AreEqualInt(a, b *int) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return *a == *b
}

// AreEqualBigInt
func AreEqualBigInt(a, b *big.Int) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return a.Text(16) == b.Text(16)
}
