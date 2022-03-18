package types

// NewVoteRound returns a VoteRound
func NewVoteRound(claimType string, votes []Vote) VoteRound {
	var totalPower uint64

	for _, vote := range votes {
		totalPower += vote.VotePower
	}
	return VoteRound{
		ClaimType:      claimType,
		Votes:          votes,
		AggregatePower: totalPower,
	}
}

// NewPrevoteRound returns a PrevoteRound
func NewPrevoteRound(claimType string, prevotes []Prevote) PrevoteRound {
	return PrevoteRound{
		ClaimType: claimType,
		Prevotes:  prevotes,
	}
}
