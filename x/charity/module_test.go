package charity_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	coreapp "github.com/user/encichain/app"
	"github.com/user/encichain/x/charity/types"
)

func TestCreatesModuleAccountOnInitBlock(t *testing.T) {
	app := coreapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	cAcc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.CharityCollectorName))
	bAcc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.BurnAccName))
	require.NotNil(t, cAcc)
	require.NotNil(t, bAcc)
}
