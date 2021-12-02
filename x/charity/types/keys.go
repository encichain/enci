package types

import (
	"encoding/binary"
)

const (
	// ModuleName defines the module name
	ModuleName = "charity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	//CharityCollectorName defines the root string for the fee collector account address
	CharityCollectorName = "charitytax_collector"

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_charity"

	// BurnAccName defines the root string for the charity burn account address
	BurnAccName = "burner"
	// this line is used by starport scaffolding # ibc/keys/name
)

// Keys for store
// stored as format - key -> value
// 0x01 				-> ProtocolBuffer(TaxRateLimits)
// 0x02 | denom bytes	-> ProtocolBuffer(sdk.Int)
// 0x03 				-> ProtocolBuffer(TaxProceeds{TaxProceeds: sdk.Coins})
// 0x04 | period bytes  -> ProtocolBuffer(TaxProceeds{TaxProceeds: sdk.Coins})
// 0x05 | period bytes  -> ProtocolBuffer([]Payout)
var (
	TaxRateLimitsKey         = []byte{0x01} // Key for tax rate limits
	TaxCapKeyPref            = []byte{0x02} // Prefix to taxcaps key
	TaxProceedsKey           = []byte{0x03} // Key for tax proceeds
	PeriodTaxProceedsKeyPref = []byte{0x04} // Prefix to *period* TaxProceeds Key
	PayoutsKeyPref           = []byte{0x05} // Prefix to *period* Payouts Key
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetTaxCapKey - stored by *denom*
func GetTaxCapKey(denom string) []byte {
	return append(TaxCapKeyPref, []byte(denom)...)
}

// GetPeriodTaxProceedsKey - stored by *period* in CollectionPeriod
func GetPeriodTaxProceedsKey(period int64) []byte {
	return GetSubKeyForPeriod(PeriodTaxProceedsKeyPref, period)
}

// GetPayoutsKey - stored by *period*
func GetPayoutsKey(period int64) []byte {
	return GetSubKeyForPeriod(PayoutsKeyPref, period)
}

// GetSubKeyForPeriod returns a subkey stored by *period*
func GetSubKeyForPeriod(prefix []byte, period int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(period))
	return append(prefix, b...)
}
