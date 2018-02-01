package jsonrpc_client

import (
	"math/big"
	"strings"
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

// ZeroPad
func ZeroPad(s string, count int) string {
	if len(s) >= count {
		return s
	}
	diff := count - len(s)
	return strings.Repeat("0", diff) + s
}
