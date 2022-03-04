package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

// Keys for oracle store, with <prefix><key> -> <value>
const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

// Keys for x/oracle store
// stored as format: key -> encoding(value)
// 0x01 | claimtype bytes															-> ProtocolBuffer(VoteRound)
// 0x02 | claimtype bytes															-> ProtocolBuffer(PrevoteRound)
// 0x03 | claimtype bytes | address length byte | validator operator address bytes	-> ProtocolBuffer(Prevote)
// 0x04 | claimtype bytes | address length byte | validator operator address bytes 	-> ProtocolBuffer(Vote)
// 0x05 | address length byte | validator operator address bytes  					-> sdk.AccAddress
// 0x06 | address length byte | delegate address bytes  							-> sdk.ValAddress
var (
	VoteRoundKey    = []byte{0x01} // prefix for a key to a VoteRound stored by claim type
	PrevoteRoundKey = []byte{0x02} // prefix for a key to a PrevoteRound stored by claim type
	PrevoteKey      = []byte{0x03} // prefix for a key to a Prevote stored by claim type | validator operator address
	VoteKey         = []byte{0x04} // prefix for a key to a Vote stored by claim type | validator operator address
	DelValKey       = []byte{0x05} // prefix for a key to a Delegate address stored by validator operator address
	ValDelKey       = []byte{0x06} // prefix for a key to a validator address stored by assigned delegate address
)

// KeyPrefix helper
func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetVoteRoundKey returns a key to a VoteRound - stored by *claimType*
func GetVoteRoundKey(claimType string) []byte {
	return append(VoteRoundKey, []byte(claimType)...)
}

// GetPrevoteRoundKey returns a key to a PrevoteRound - stored by *claimType*
func GetPrevoteRoundKey(claimType string) []byte {
	return append(PrevoteRoundKey, []byte(claimType)...)
}

// GetPrevoteKey returns a key to a Prevote - stored by *claimType* | *Validator* operator address
func GetPrevoteKey(val sdk.ValAddress, claimType string) []byte {
	key := append(PrevoteKey, []byte(claimType)...)
	return append(key, address.MustLengthPrefix(val)...)
}

// GetVoteKey returns a key to a Vote - stored by *claimType*| *Validator* operator address
func GetVoteKey(val sdk.ValAddress, claimType string) []byte {
	key := append(VoteKey, []byte(claimType)...)
	return append(key, address.MustLengthPrefix(val)...)
}

// GetDelValKey returns the validator for a given delegate address - stored by *delegate* address
func GetDelValKey(del sdk.AccAddress) []byte {
	return append(DelValKey, address.MustLengthPrefix(del)...)
}

// GetValDelKey returns the delegate for a given validator address - stored by *Validator* operator address
func GetValDelKey(val sdk.ValAddress) []byte {
	return append(ValDelKey, address.MustLengthPrefix(val)...)
}
