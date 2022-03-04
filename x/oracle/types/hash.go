package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// VoteHash is a SHA256 hash of the salt, hash of the claim, and validator address which is meant to hide vote
// Format: SHA256("{salt}:{claim}:{validator address}")
type VoteHash []byte

// VoteHash returns the SHA-256 hash for a precommit given the proper args
func CreateVoteHash(salt string, claimHash string, validator sdk.ValAddress) VoteHash {
	h := sha256.New()
	_, err := h.Write([]byte(fmt.Sprintf("%s:%s:%s", salt, claimHash, validator.String())))
	if err != nil {
		panic(err)
	}
	return h.Sum(nil)
}

// HexStringToVoteHash coverts a hex string representation of a VoteHash to a VoteHash
func HexStringToVoteHash(s string) (VoteHash, error) {
	h, err := hex.DecodeString(s)
	if err != nil {
		return h, err
	}
	return h, nil
}

// String implements Stringer interface
func (h VoteHash) String() string {
	return hex.EncodeToString(h)
}

// Size returns the len of VoteHash
func (h VoteHash) Size() int {
	return len(h)
}

// Empty checks if VoteHash is empty
func (h VoteHash) Empty() bool {
	return len(h) == 0
}

// Equal compares h with h2 Votehash
func (h VoteHash) Equal(h2 VoteHash) bool {
	return bytes.Equal(h, h2)
}
