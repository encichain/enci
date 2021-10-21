package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	BlockTime       = uint64(5)
	BlocksPerMinute = 60 / BlockTime
	BlocksPerHour   = BlocksPerMinute * 60
	BlocksPerDay    = BlocksPerHour * 24
	BlocksPerWeek   = BlocksPerDay * 7
	BlocksPerMonth  = BlocksPerDay * 30
	BlocksPerYear   = BlocksPerDay * 365
	BlocksPerPeriod = BlocksPerWeek
)

// IsLastBlockPeriod returns bool representing if current block is the last block of a CollectionPeriod
func IsLastBlockPeriod(ctx sdk.Context) bool {
	return (ctx.BlockHeight()+1)%int64(BlocksPerPeriod) == 0
}
