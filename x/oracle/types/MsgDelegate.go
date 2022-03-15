package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDelegate{}

// Message types for the oracle module
const (
	TypeMsgDelegate = "delegate"
)

// NewMsgDelegate returns a new MsgDelegateFeedConsent
func NewMsgDelegate(val, del sdk.AccAddress) *MsgDelegate {
	return &MsgDelegate{
		Validator: val.String(),
		Delegate:  del.String(),
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgDelegate) ValidateBasic() error {
	if _, err := sdk.ValAddressFromBech32(msg.Validator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	if _, err := sdk.AccAddressFromBech32(msg.Delegate); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}
	return nil
}

// GetSignBytes implements the LegacyMsg.GetSignBytes method
func (msg MsgDelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDelegate) GetSigners() []sdk.AccAddress {
	val, err := sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{sdk.AccAddress(val)}
}

// MustGetValidator returns the sdk.AccAddress for the validator
func (msg MsgDelegate) GetValidatorAddress() (sdk.ValAddress, error) {
	val, err := sdk.ValAddressFromBech32(msg.Validator)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// MustGetDelegate returns the sdk.AccAddress for the delegate
func (msg MsgDelegate) GetDelegateAddress() (sdk.AccAddress, error) {
	val, err := sdk.AccAddressFromBech32(msg.Delegate)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// ===== Implements legacytx.LegacyMsg interface =====

// Route implements sdk.Msg
func (msg MsgDelegate) Route() string { return ModuleName }

// Type implements sdk.Msg
func (msg MsgDelegate) Type() string { return TypeMsgDelegate }
