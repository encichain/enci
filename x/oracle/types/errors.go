package types

// DONTCOVER

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// x/oracle module errors
var (
	ErrInvalidClaim        = sdkerrors.Register(ModuleName, 2, "invalid claim")
	ErrNoClaimExists       = sdkerrors.Register(ModuleName, 3, "no claim exits")
	ErrNoClaimTypeExists   = sdkerrors.Register(ModuleName, 4, "claim type is not registered as part of the oracle params")
	ErrNoPrevote           = sdkerrors.Register(ModuleName, 5, "no prevote from this validator exists for this claim")
	ErrIncorrectClaimRound = sdkerrors.Register(ModuleName, 6, "claim must be submitted after the prevote round is over")
	ErrNoVoteDelegate      = sdkerrors.Register(ModuleName, 7, "no vote delegate exists for the validator")
	ErrNoVoteDelegator     = sdkerrors.Register(ModuleName, 8, "address is not listed as vote delegate for any validator")
	ErrInvalidHash         = sdkerrors.Register(ModuleName, 9, "invalid hash")
	ErrInvalidHashLength   = sdkerrors.Register(ModuleName, 10, fmt.Sprintf("invalid truncated hash length. should be : %d", 2*tmhash.TruncatedSize))
	ErrNotPrevotePeriod    = sdkerrors.Register(ModuleName, 11, "current block is not part of a prevote period")
	ErrInvalidPrevoteBlock = sdkerrors.Register(ModuleName, 12, "prevote submit block is invalid for this vote")
	ErrNotVotePeriod       = sdkerrors.Register(ModuleName, 13, "current block is not part of a voting period")
	ErrNoSigner            = sdkerrors.Register(ModuleName, 14, "could not get signer account address")
	ErrVerificationFailed  = sdkerrors.Register(ModuleName, 15, "prevote hash does not match hashed vote")
)
