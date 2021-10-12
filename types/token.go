package types

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	MicroTokenDenom = "uenci"
	MilliTokenDenom = "menci"
	TokenDenom      = "enci"
)

var TokenMetaData = banktypes.Metadata{
	Description: "The native token of EnciChain",
	DenomUnits: []*banktypes.DenomUnit{
		{Denom: MicroTokenDenom, Exponent: uint32(0), Aliases: []string{"microenci"}},
		{Denom: MilliTokenDenom, Exponent: uint32(3), Aliases: []string{"millienci"}},
		{Denom: TokenDenom, Exponent: uint32(6), Aliases: []string{}},
	},

	Base:    MicroTokenDenom,
	Display: TokenDenom,
	Name:    "ENCI",
	Symbol:  "ENCI",
}
