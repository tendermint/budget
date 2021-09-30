<!-- order: 4 -->

# Begin-Block

At each `BeginBlock`, get all budgets registered in `params.Budgets` and select the valid budgets to collect budgets for the block by its respective plan (defined rate, budget source address, collection address, start time and end time). Then, it distributes the collected amount of coins from `BudgetSourceAddrss` to `CollectionAddress`.

+++ https://github.com/tendermint/budget/blob/master/x/budget/abci.go#L15-L22

## Workflow

1. Get all the budgets registered in `params.Budgets` and proceed with the valid and not expired budgets. Otherwise, it exits and waits for the next block. 

2. Create a map by `BudgetSourceAddress` to handle the budgets for the same `BudgetSourceAddress` together based on the same balance when calculating rates for the same `BudgetSourceAddress`.

3. Collect budgets from `BudgetSourceAddress` and send amount of coins to `CollectionAddress` relative to the rate of each budget`.

4. Cumulate `TotalCollectedCoins` and emit events about the successful budget collection for each budget.

