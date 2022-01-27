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
	BlocksPerEpoch  = uint64(100)
)

// IsLastBlockEpoch returns bool representing if current block is the last block of a CollectionEpoch
func IsLastBlockEpoch(ctx sdk.Context) bool {
	return (ctx.BlockHeight()+1)%int64(BlocksPerEpoch) == 0
}
