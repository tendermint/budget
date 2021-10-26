<!-- order: 2 -->
# State
## Budget

Budget object is not stored in KVStore. 

```go
// Budget contains budget information
type Budget struct {
	Name                string    // name of the budget 
	Rate                sdk.Dec   // distributing amount by ratio of total budget source
	BudgetSourceAddress string    // bech32-encoded address that source of the budget
	CollectionAddress   string    // bech32-encoded address that collects budget from the budget source address
	StartTime           time.Time // start time of the budget plan
	EndTime             time.Time // end time of the budget plan
}
```

+++ https://github.com/tendermint/budget/blob/main/proto/tendermint/budget/v1beta1/budget.proto#L25-L53

## TotalCollectedCoins

```go
// TotalCollectedCoins are cumulative coins collected in the budget since the bucket was created.
type TotalCollectedCoins struct {
	TotalCollectedCoins sdk.Coins
}
```

+++ https://github.com/tendermint/budget/blob/main/proto/tendermint/budget/v1beta1/budget.proto#L55-L64


For the purpose of tracking total collected coins for a budget, budget name is used as key to find it in store.

- TotalCollectedCoins: `0x11 | BudgetName -> TotalCollectedCoins`
