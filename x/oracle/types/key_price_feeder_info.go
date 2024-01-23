package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PriceFeederInfoKeyPrefix is the prefix to retrieve all PriceFeederInfo
	PriceFeederInfoKeyPrefix = "PriceFeederInfo/value/"
)

// PriceFeederInfoKey returns the store key to retrieve a PriceFeederInfo from the index fields
func PriceFeederInfoKey(
	name string,
) []byte {
	var key []byte = []byte(name)
	return key
}
