package types_test

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/encichain/enci/x/oracle/exported"
	"github.com/encichain/enci/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func testCreateMsgVote(t *testing.T, c exported.Claim, s sdk.AccAddress) exported.MsgVoteI {
	msg, err := types.NewMsgVote(s, c, "")
	require.NoError(t, err)
	return msg
}

func TestMsgVote(t *testing.T) {
	submitter := sdk.AccAddress("test________________")
	claim := types.NewTestClaim(0, "test", "test")
	claimAny, err := codectypes.NewAnyWithValue(claim)
	require.NoError(t, err)
	invalidClaimAny, err := codectypes.NewAnyWithValue(&types.Prevote{Hash: "", Validator: "", SubmitBlock: 1})
	require.NoError(t, err)

	testCases := []struct {
		desc       string
		msg        sdk.Msg
		submitter  sdk.AccAddress
		expectPass bool
	}{
		{
			"invalid claim blockheight",
			testCreateMsgVote(t, &types.TestClaim{
				BlockHeight: -1,
				Content:     "test",
				ClaimType:   "test",
			}, submitter),
			submitter,
			false,
		},
		{
			"valid",
			testCreateMsgVote(t, &types.TestClaim{
				BlockHeight: 10,
				Content:     "test",
				ClaimType:   "test",
			}, submitter),
			submitter,
			true,
		},
		{
			"invalid msg signer",
			&types.MsgVote{
				Salt:   "",
				Claim:  claimAny,
				Signer: "",
			},
			submitter,
			false,
		},
		{
			"claim does not satisfy Claim interface",
			&types.MsgVote{
				Salt:   "",
				Claim:  invalidClaimAny,
				Signer: submitter.String(),
			},
			submitter,
			false,
		},
	}

	for i, tc := range testCases {
		_, ok := tc.msg.(*types.MsgVote)
		require.Equal(t, sdk.MsgTypeURL(tc.msg), "/enci.oracle.v1beta1.MsgVote", "unexpected result for tc #%d", i)
		require.True(t, ok, "unexpected result for tc #%d", i)

		if tc.expectPass {
			require.Equal(t, tc.msg.GetSigners(), []sdk.AccAddress{tc.submitter}, "unexpected result for tc #%d", i)
			require.NoError(t, tc.msg.ValidateBasic())
			continue
		}

		require.Error(t, tc.msg.ValidateBasic(), "unexpected result for tc #%d", i)
	}
}

func TestMsgPrevote(t *testing.T) {
	submitter := sdk.AccAddress("test________________")
	claim := types.NewTestClaim(0, "test", "test")
	voteHash := types.CreateVoteHash("12", claim.Hash().String(), sdk.ValAddress(submitter))

	testCases := []struct {
		desc    string
		hash    string
		signer  sdk.AccAddress
		expPass bool
	}{
		{"valid", voteHash.String(), submitter, true},
		{"invalid signer", voteHash.String(), sdk.AccAddress([]byte{}), false},
		{"invalid hash length", "ffffffff", submitter, false},
	}

	for i, tc := range testCases {
		hash, err := types.HexStringToVoteHash(tc.hash)
		require.NoError(t, err)
		msgPrevote := types.NewMsgPrevote(claim.Type(), tc.signer, hash)
		if tc.expPass {
			require.NoError(t, msgPrevote.ValidateBasic(), "no error expected for tc #%d", i)
		} else {
			require.Error(t, msgPrevote.ValidateBasic(), "error expected for tc #%d", i)
		}
	}
}

func TestMsgDelegate(t *testing.T) {
	addrs := []sdk.AccAddress{
		sdk.AccAddress("test1_______________"),
		sdk.AccAddress("test2_______________"),
	}
	testCases := []struct {
		desc    string
		valAddr sdk.ValAddress
		accAddr sdk.AccAddress
		expPass bool
	}{
		{"valid", sdk.ValAddress(addrs[0]), addrs[1], true},
		{"invalid delegator", sdk.ValAddress([]byte{}), addrs[1], false},
		{"invalid delegate", sdk.ValAddress(addrs[0]), sdk.AccAddress{}, false},
	}

	for i, tc := range testCases {
		msgDelegate := types.NewMsgDelegate(tc.valAddr, tc.accAddr)
		if tc.expPass {
			require.NoError(t, msgDelegate.ValidateBasic(), "no error expected for tc #%d", i)
		} else {
			require.Error(t, msgDelegate.ValidateBasic(), "error expected for tc #%d", i)
		}
	}

}
