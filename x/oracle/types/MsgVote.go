package types

import (
	fmt "fmt"

	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/encichain/enci/x/oracle/exported"
	proto "github.com/gogo/protobuf/proto"
)

// Message types for the oracle module
const (
	TypeMsgVote = "vote"
)

var (
	_ sdk.Msg                       = &MsgVote{}
	_ types.UnpackInterfacesMessage = MsgVote{}
	_ exported.MsgVoteI             = &MsgVote{}
)

// NewMsgVote returns a new MsgVote with a signer/submitter.
func NewMsgVote(s sdk.AccAddress, claim exported.Claim, salt string) (*MsgVote, error) {
	msg, ok := claim.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("cannot proto marshal %T", claim)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}
	return &MsgVote{
		Signer: s.String(),
		Claim:  any,
		Salt:   salt,
	}, nil
}

// GetSigners get msg signers. The signer can be either the validator or its delegate
func (msg MsgVote) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method
func (msg MsgVote) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validation
func (msg MsgVote) ValidateBasic() error {
	if msg.Signer == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "voter can't be empty")
	}
	claim := msg.GetClaim()
	if claim == nil {
		return sdkerrors.Wrap(ErrInvalidClaim, "missing claim")
	}
	if err := claim.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

// GetClaim get the claim
func (msg MsgVote) GetClaim() exported.Claim {
	claim, ok := msg.Claim.GetCachedValue().(exported.Claim)
	if !ok {
		return nil
	}
	return claim
}

// GetSigner gets the submitter account address
func (msg MsgVote) GetSigner() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil
	}
	return accAddr
}

// MustGetSigner returns the account address, panics if nil
func (msg MsgVote) MustGetSigner() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return accAddr
}

// UnpackInterfaces unpack
func (msg MsgVote) UnpackInterfaces(ctx types.AnyUnpacker) error {
	var claim exported.Claim
	return ctx.UnpackAny(msg.Claim, &claim)
}

// ===== Implements legacytx.LegacyMsg interface =====

// Route get msg route
func (msg MsgVote) Route() string {
	return RouterKey
}

// Type get msg type
func (msg MsgVote) Type() string {
	return TypeMsgVote
}
