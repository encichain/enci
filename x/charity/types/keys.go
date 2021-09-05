package types

const (
	// ModuleName defines the module name
	ModuleName = "charity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	//CharityCollectorName the root string for the fee collector account address
	CharityCollectorName = "charity_collector"

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_charity"

	// this line is used by starport scaffolding # ibc/keys/name
)

// Keys for store
// stored as format - key: value
// 0x01: sdk.Dec

var (
	TaxRateKey = []byte{0x01} // Key for tax rate
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}
