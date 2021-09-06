package types

const (
	// ModuleName is the name of the budget module
	ModuleName = "budget"

	// RouterKey is the message router key for the budget module
	RouterKey = ModuleName

	// StoreKey is the default store key for the budget module
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the budget module
	QuerierRoute = ModuleName
)

var (
	TotalCollectedCoinsKeyPrefix = []byte{0x11}
)

func GetTotalCollectedCoinsKey(budgetName string) []byte {
	return append(TotalCollectedCoinsKeyPrefix, []byte(budgetName)...)
}
