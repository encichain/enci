package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	charitytypes "github.com/encichain/enci/x/charity/types"
)

// AccountKeeper defines the contract needed for AccountKeeper related APIs.
// Interface provides support to use non-sdk AccountKeeper for AnteHandler's decorators.
type AccountKeeper interface {
	GetParams(ctx sdk.Context) (params types.Params)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	SetAccount(ctx sdk.Context, acc types.AccountI)
	GetModuleAddress(moduleName string) sdk.AccAddress
}

// FeegrantKeeper defines the expected feegrant keeper.
type FeegrantKeeper interface {
	UseGrantedFees(ctx sdk.Context, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) error
}

// CharityKeeper defines the expected charity keeper
type CharityKeeper interface {
	RecordTaxProceeds(ctx sdk.Context, proceeds sdk.Coins)
	GetTaxCap(ctx sdk.Context, denom string) sdk.Int
	GetTaxRate(ctx sdk.Context) sdk.Dec
	GetTaxRateLimits(ctx sdk.Context) charitytypes.TaxRateLimits
}
