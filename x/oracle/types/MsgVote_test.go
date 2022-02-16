package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/exported"
	"github.com/encichain/enci/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func testMsgCreateClaim(t *testing.T, c exported.Claim, s sdk.AccAddress) exported.MsgVoteI {
	msg, err := types.NewMsgVote(s, c, "")
	require.NoError(t, err)
	return msg
}

func TestMsgCreateClaim(t *testing.T) {
	submitter := sdk.AccAddress("test________________")

	testCases := []struct {
		msg       sdk.Msg
		submitter sdk.AccAddress
		expectErr bool
	}{
		{
			testMsgCreateClaim(t, &types.TestClaim{
				BlockHeight: 0,
				Content:     "test",
				ClaimType:   "test",
			}, submitter),
			submitter,
			true,
		},
		{
			testMsgCreateClaim(t, &types.TestClaim{
				BlockHeight: 10,
				Content:     "test",
				ClaimType:   "test",
			}, submitter),
			submitter,
			false,
		},
	}

	for i, tc := range testCases {
		//require.Equal(t, sdk.MsgTypeURL(tc.msg), types.RouterKey, "unexpected result for tc #%d", i)
		_, ok := tc.msg.(*types.MsgVote)
		require.True(t, ok, "unexpected result for tc #%d", i)
		require.Equal(t, tc.expectErr, tc.msg.ValidateBasic() != nil, "unexpected result for tc #%d", i)

		if !tc.expectErr {
			require.Equal(t, tc.msg.GetSigners(), []sdk.AccAddress{tc.submitter}, "unexpected result for tc #%d", i)
		}
	}
}
