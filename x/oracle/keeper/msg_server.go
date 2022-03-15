package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/encichain/enci/x/oracle/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the oracle MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	claim := msg.GetClaim()
	if claim == nil {
		return nil, sdkerrors.Wrap(types.ErrNoClaimExists, msg.Claim.String())
	}
	signer := msg.GetSigner()
	if signer == nil {
		return nil, sdkerrors.Wrap(types.ErrNoSigner, msg.Signer)
	}
	// Check if vote submitted during voting period
	if isVotePeriod := k.IsVotePeriod(ctx); !isVotePeriod {
		return nil, sdkerrors.Wrap(types.ErrNotVotePeriod, fmt.Sprint(ctx.BlockHeight()))
	}

	// Get validator address. If no delegator, signer of msg is the validator or invalid
	valAddr := getDelegatorAddr(ctx, k, signer)
	val, found := k.StakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return nil, sdkerrors.Wrap(stakingtypes.ErrNoValidatorFound, valAddr.String())
	}

	// Check if there is a Prevote in store for specific vote
	prevote, err := k.GetPrevote(ctx, valAddr, claim.Type())
	if err != nil {
		return nil, err
	}

	// Verify prevote was submitted during proper Prevote period
	if prevote.SubmitBlock < k.PreviousPrevotePeriod(ctx) {
		return nil, sdkerrors.Wrap(types.ErrInvalidPrevoteBlock, fmt.Sprint(prevote.SubmitBlock))
	}

	//Verify prevote hash matches Vote msg data
	hash := types.CreateVoteHash(msg.Salt, claim.Hash().String(), valAddr)
	if prevote.Hash != hash.String() {
		return nil, sdkerrors.Wrapf(types.ErrVerificationFailed, "prevote hash: %s, vote hash: %s", prevote.Hash, hash.String())
	}

	// create vote object
	vote, err := types.NewVote(claim, valAddr, uint64(val.GetConsensusPower(sdk.DefaultPowerReduction)))
	if err != nil {
		return nil, fmt.Errorf("could not create vote object for claim: %s", claim.String())
	}

	// Set vote to store and delete prevote
	k.SetVote(ctx, valAddr, vote, claim.Type())
	k.DeletePrevote(ctx, valAddr, claim.Type())

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeVote),
			sdk.NewAttribute(sdk.AttributeKeySender, signer.String()),
		),
	)

	return &types.MsgVoteResponse{}, nil
}

// Delegate implements types.MsgServer
func (k msgServer) Delegate(c context.Context, msg *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	val, err := msg.GetValidatorAddress()
	if err != nil {
		return nil, err
	}
	del, err := msg.GetDelegateAddress()
	if err != nil {
		return nil, err
	}

	// Check if validator account exists
	if _, found := k.Keeper.StakingKeeper.GetValidator(ctx, sdk.ValAddress(val)); !found {
		return nil, sdkerrors.Wrap(stakingtypes.ErrNoValidatorFound, val.String())
	}
	// Set new delegation to store
	k.SetVoterDelegation(ctx, del, val)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeVoterDelegation),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Validator),
			sdk.NewAttribute(types.AttributeKeyDelegate, msg.Delegate),
		),
	)

	return &types.MsgDelegateResponse{}, nil
}

func (k msgServer) Prevote(goCtx context.Context, msg *types.MsgPrevote) (*types.MsgPrevoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	signer := msg.GetSigner()
	if signer == nil {
		return nil, sdkerrors.Wrap(types.ErrNoSigner, msg.Signer)
	}
	// Verify prevote is submitted during Prevote Period
	if isPrevotePeriod := k.IsPrevotePeriod(ctx); !isPrevotePeriod {
		return nil, sdkerrors.Wrap(types.ErrNotPrevotePeriod, fmt.Sprint(ctx.BlockHeight()))
	}

	// Check if address has a delegator. If no delegator, it is either the validator or invalid
	valAddr := getDelegatorAddr(ctx, k, signer)
	// Validate returned validator address. This also catches prevotes submitted by unauthorized address
	_, found := k.StakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return nil, sdkerrors.Wrap(stakingtypes.ErrNoValidatorFound, valAddr.String())
	}

	voteHash, err := types.HexStringToVoteHash(msg.Hash)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidHash, msg.Hash)
	}

	prevote := types.NewPrevote(voteHash, valAddr, uint64(ctx.BlockHeight()))
	// Set prevote to store
	k.SetPrevote(ctx, valAddr, prevote, msg.ClaimType)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypePrevote),
			sdk.NewAttribute(sdk.AttributeKeySender, signer.String()),
		),
	})

	return &types.MsgPrevoteResponse{}, nil
}

func getDelegatorAddr(ctx sdk.Context, k msgServer, signer sdk.AccAddress) sdk.ValAddress {
	// get delegate's validator
	valAddr, err := k.GetVoterDelegator(ctx, signer)

	// if there is no delegation it must be the validator or invalid
	if err != nil {
		valAddr = sdk.ValAddress(signer)
	}

	return valAddr
}
