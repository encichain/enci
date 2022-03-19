package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// Message types for the oracle module
const (
	TypeMsgPrevote = "prevote"
)

var (
	_ sdk.Msg = &MsgPrevote{}
)

// NewMsgPrevote returns a new MsgPrevotePrevote with a signer.
func NewMsgPrevote(claimType string, accountAddr sdk.AccAddress, hash VoteHash) *MsgPrevote {
	return &MsgPrevote{
		ClaimType: claimType,
		Signer:    accountAddr.String(),
		Hash:      hash.String()}
}

// GetSigners implements sdk.Msg
func (msg MsgPrevote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.MustGetSigner()}
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method
func (msg MsgPrevote) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic validation
func (msg MsgPrevote) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	_, err := HexStringToVoteHash(msg.Hash)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidHash, err.Error())
	}

	// Hex encoded hash is double the size of hash bytes
	if len(msg.Hash) != tmhash.TruncatedSize*2 {
		return ErrInvalidHashLength
	}

	if len(msg.ClaimType) == 0 {
		return fmt.Errorf("claim type cannot be empty: %s", msg.ClaimType)
	}

	return nil
}

// GetSigner gets the submitter account address
func (msg MsgPrevote) GetSigner() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil
	}
	return accAddr
}

// MustGetSigner returns submitter
func (msg MsgPrevote) MustGetSigner() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return accAddr
}

// ===== Implements legacytx.LegacyMsg interface =====

// Route get msg route
func (msg MsgPrevote) Route() string {
	return RouterKey
}

// Type get msg type
func (msg MsgPrevote) Type() string {
	return TypeMsgPrevote
}
