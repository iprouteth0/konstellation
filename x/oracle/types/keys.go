package types

var (
	ExchangeRateKey = []byte{0x00}
)

const (
	// ModuleName is the name of the module
	ModuleName = "oracle"

	RouterKey = ModuleName

	// StoreKey is the default store key for mint
	StoreKey = ModuleName

	QueryExchangeRate = "exchange-rate"
)
