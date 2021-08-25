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
		{Denom: "utoken", Exponent: uint32(0), Aliases: []string{"microtoken"}},
		{Denom: "mtoken", Exponent: uint32(3), Aliases: []string{"millitoken"}},
		{Denom: "token", Exponent: uint32(6), Aliases: []string{}},
	},

	Base:    "utoken",
	Display: "token",
	//Name:    "TOKE",
	//Symbol:  "TOKE",
}
