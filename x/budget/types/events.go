package types

// Event types for the budget module.
const (
	EventTypeAddBudget       = "add_budget"
	EventTypeDeleteBudget    = "delete_budget"
	EventTypeUpdateBudget    = "update_budget"
	EventTypeBudgetCollected = "budget_collected"

	AttributeValueName                = "name"
	AttributeValueCollectionAddress   = "collection_address"
	AttributeValueBudgetSourceAddress = "budget_source_address"
	AttributeValueRate                = "rate"
	AttributeValueAmount              = "amount"
)
