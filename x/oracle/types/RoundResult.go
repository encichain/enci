package types

import tmbytes "github.com/tendermint/tendermint/libs/bytes"

// ClaimVoteResult is a record of votes for each claim containing the hash of a claim and its VotePower.
type ClaimVoteResult struct {
	ClaimHash tmbytes.HexBytes
	VotePower int64
}

// RoundResult is a record of vote tallies for a given round and claim type
type RoundResult struct {
	VotePower  int64
	TotalPower int64
	ClaimType  string
	Claims     []*ClaimVoteResult
}

// UpsertClaim upserts a claim to a RoundResult{}
func (r *RoundResult) UpsertClaim(claimHash tmbytes.HexBytes, votePower int64) {
	var existingClaim *ClaimVoteResult
	// Check for existing claim in RoundResult object
	for _, claim := range r.Claims {
		if claim.ClaimHash.String() == claimHash.String() {
			existingClaim = claim
		}
	}
	// If existing claim exists, combine by adding to votePower
	if existingClaim != nil {
		existingClaim.VotePower += votePower
		return
	}
	newClaim := newClaimVoteResult(claimHash, votePower)
	r.Claims = append(r.Claims, newClaim)
}

func newClaimVoteResult(claimHash tmbytes.HexBytes, votePower int64) *ClaimVoteResult {
	return &ClaimVoteResult{
		ClaimHash: claimHash,
		VotePower: votePower,
	}
}
