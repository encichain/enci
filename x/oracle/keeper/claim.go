package keeper

import (
	"fmt"

	"github.com/encichain/enci/x/oracle/exported"
)

// MustUnmarshalClaim attempts to decode and return an Claim object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalClaim(bz []byte) exported.Claim {
	Claim, err := k.UnmarshalClaim(bz)
	if err != nil {
		panic(fmt.Errorf("failed to decode Claim: %w", err))
	}

	return Claim
}

// MustMarshalClaim attempts to encode an Claim object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalClaim(Claim exported.Claim) []byte {
	bz, err := k.MarshalClaim(Claim)
	if err != nil {
		panic(fmt.Errorf("failed to encode Claim: %w", err))
	}

	return bz
}

// MarshalClaim marshals a Claim interface. If the given type implements
// the Marshaler interface, it is treated as a Proto-defined message and
// serialized that way. Otherwise, it falls back on the internal Amino codec.
func (k Keeper) MarshalClaim(claimI exported.Claim) ([]byte, error) {
	return k.cdc.MarshalInterface(claimI)
}

// UnmarshalClaim returns a Claim interface from raw encoded Claim
// bytes of a Proto-based Claim type. An error is returned upon decoding
// failure.
func (k Keeper) UnmarshalClaim(bz []byte) (exported.Claim, error) {
	var claim exported.Claim
	if err := k.cdc.UnmarshalInterface(bz, &claim); err != nil {
		return nil, err
	}

	return claim, nil
}
