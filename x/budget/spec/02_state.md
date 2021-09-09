<!-- order: 2 -->

# State

## Budget

The Budget structure is not stored in kv, but in parameters in the budget module as budgets.

```go
type Budget struct {
	Name                string // name defines the name of the budget
	Rate                sdk.Dec // rate specifies the distributing amount by ratio of total budget source
	BudgetSourceAddress string // budget_source_address defines the bech32-encoded address that source of the budget
	CollectionAddress   string // collection_address defines the bech32-encoded address of the budget pool to distribute
	StartTime           time.Time // start_time specifies the start time of the budget
	EndTime             time.Time // end_time specifies the end time of the budget
}
```

### Total collected coins

The total collected coins is cumulative coins collected in the budget since the bucket was created.
The total collected coins for each budget are stored as follows:
```
0x11 | BudgetName -> TotalCollectedCoins
```

Where `TotalCollectedCoins` is just a wrapper around `sdk.Coins`:
```go
type TotalCollectedCoins struct {
	TotalCollectedCoins sdk.Coins
}
```
