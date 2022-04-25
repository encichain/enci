package keeper

import (
	"github.com/encichain/enci/x/enciprice/types"
)

var _ types.QueryServer = Keeper{}
