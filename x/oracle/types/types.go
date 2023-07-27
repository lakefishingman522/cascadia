package types

// Custom call data that has uint8 as parameter
type Calldata struct {
	Symbols            []string
	MinimumSourceCount uint8
}

type OraclePriceData struct {
	Symbol       string
	ResponseCode uint8
	Rate         uint64
}

// Oracle response
type OraclePriceResults_UInt8Version struct {
	Responses []OraclePriceData
}