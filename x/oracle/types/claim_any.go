package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/encichain/enci/x/oracle/exported"
)

func (v Vote) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var c exported.Claim
	return unpacker.UnpackAny(v.Claim, &c)
}

func (voteRound VoteRound) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, v := range voteRound.Votes {
		err := v.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}
