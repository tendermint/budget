<!-- order: 4 -->

# Begin-Block

At each `BeginBlock`, get all budgets registered in `params.Budgets` and select the valid budgets to collect budgets for the block by its respective plan (defined rate, source address, destination address, start time and end time). Then, it distributes the collected amount of coins from `SourceAddrss` to `DestinationAddress`.

+++ https://github.com/tendermint/budget/blob/main/x/budget/abci.go#L15-L22

## Workflow

1. Get all the budgets registered in `params.Budgets` and proceed with the valid and not expired budgets. Otherwise, it exits and waits for the next block. 

2. Create a map by `SourceAddress` to handle the budgets for the same `SourceAddress` together based on the same balance when calculating rates for the same `SourceAddress`.

3. Collect budgets from `SourceAddress` and send amount of coins to `DestinationAddress` relative to the rate of each budget`.

4. Cumulate `TotalCollectedCoins` and emit events about the successful budget collection for each budget.

