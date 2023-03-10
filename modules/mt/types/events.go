package types

// MT module event types
var (
	EventTypeIssueDenom    = "issue_denom"
	EventTypeTransfer      = "transfer_mt"
	EventTypeEditMT        = "edit_mt"
	EventTypeMintMT        = "mint_mt"
	EventTypeBurnMT        = "burn_mt"
	EventTypeTransferDenom = "transfer_denom"

	AttributeValueCategory = ModuleName

	AttributeKeySender    = "sender"
	AttributeKeyCreator   = "creator"
	AttributeKeyRecipient = "recipient"
	AttributeKeyOwner     = "owner"
	AttributeKeyMTID      = "mt_id"
	AttributeKeySupply    = "supply"
	AttributeKeyAmount    = "amount"
	AttributeKeyDenomID   = "denom_id"
	AttributeKeyDenomName = "denom_name"
)
