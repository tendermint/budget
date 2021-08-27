<!-- order: 4 -->

# Begin-Block

1. if current height % `params.EpochBlocks` != 0 or `params.EpochBlocks` == 0, exit.
1. Get all the Budgets registered in params.Budgets and select only the valid budgets. If there is no valid budget, exit.
2. Create a map by `BudgetSourceAddress` to handle the budgets for the same `BudgetSourceAddress` together based on the same balance when calculating rates for the same `BudgetSourceAddress`.
3. Collect budgets from `BudgetSourceAddress` and send them to `CollectionAddress` according to the `Rate` of each `Budget`.
4. Write to metric about successful budget collection and emit events.