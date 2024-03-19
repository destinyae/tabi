package types

const (
	EventTypeMintNode = "mint_node"

	AttributeValueCategory = ModuleName

	AttributeKeyNodeID     = "node_id"
	AttributeKeyDivisionID = "division_id"
	AttributeKeyReceiver   = "receiver"

	EventTypeUpdatePowerOnPeriod = "update_power_on_period"
	AttributeKeyOldPowerOnPeriod = "old_power_on_period"
	AttributeKeyNewPowerOnPeriod = "new_power_on_period"

	EventTypeReceiveExperience = "receive_experience"
	AttributeKeyExperience     = "experience"

	EventTypeUpdateUserExperience = "update_user_experience"
	AttributeKeyOwner             = "owner"
	AttributeKeyOldExperience     = "old_experience"
	AttributeKeyNewExperience     = "new_experience"
)
