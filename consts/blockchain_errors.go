package Consts

import "fmt"

type BlockError int

const (
	MismatchedIndex BlockError = iota
	InvalidHash
	AchronologicalTimeStamp
	MismatchedPreviousHash
	InvalidGenesisBlockFormat
	InvalidInput
	InsufficientInputValue
	InvalidCoinBaseTransaction
)

func (e BlockError) Error() string {
	switch e {
	case MismatchedIndex:
		return "mismatched index"
	case InvalidHash:
		return "invalid hash"
	case AchronologicalTimeStamp:
		return "achronological timestamp"
	case MismatchedPreviousHash:
		return "mismatched previous hash"
	case InvalidGenesisBlockFormat:
		return "invalid genesis block format"
	case InvalidInput:
		return "invalid input"
	case InsufficientInputValue:
		return "insufficient input value"
	case InvalidCoinBaseTransaction:
		return "invalid coinbase transaction"
	default:
		return fmt.Sprintf("unknown error code: %d", e)
	}
}
