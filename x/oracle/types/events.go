package types

// claim module events
const (
	AttributeValueCategory = "oracle"

	EventTypeVoterDelegation = "voter_delegation"
	AttributeKeyDelegate     = "delegate"
	AttributeKeyValidator    = "validator"

	EventTypeVote           = "vote"
	EventTypePrevote        = "prevote"
	AttributeKeyPrevoteHash = "prevote_hash"
	AttributeKeyVoter       = "voter"
)
