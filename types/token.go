package types

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	MicroTokenDenom = "utoken"
	MilliTokenDenom = "mtoken"
	TokenDenom      = "token"
)

var TokenMetaData = banktypes.Metadata{
	Description: "The native token of the Charity Chain",
	DenomUnits: []*banktypes.DenomUnit{
		{Denom: MicroTokenDenom, Exponent: uint32(0), Aliases: []string{"microtoken"}},
		{Denom: MilliTokenDenom, Exponent: uint32(3), Aliases: []string{"millitoken"}},
		{Denom: TokenDenom, Exponent: uint32(6), Aliases: []string{}},
	},

	Base:    MicroTokenDenom,
	Display: TokenDenom,
	//Name:    "TOKE",
	//Symbol:  "TOKE",
}
