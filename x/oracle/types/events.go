package types

// claim module events
const (
	AttributeValueCategory = "oracle"

	EventTypeOraclePeriod = "oracle_start"
	EventTypeVoteBegin    = "vote_period_begin"
	EventTypePrevoteBegin = "prevote_period_begin"

	EventTypeVoterDelegation = "voter_delegation"
	AttributeKeyDelegate     = "delegate"
	AttributeKeyValidator    = "validator"

	EventTypeVote           = "vote"
	EventTypePrevote        = "prevote"
	AttributeKeyPrevoteHash = "prevote_hash"
	AttributeKeyVoter       = "voter"
)
