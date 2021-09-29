<!-- order: 3 -->

# State Transitions

This document describes the state transaction operations pertaining to the budget module.

- BudgetCollection
- Accumulate `TotalCollectedCoins`
## BudgetCollection

Get all budgets registered in `params.Budgets` and select the valid budgets to collect budgets for the block by its respective rate.
`BudgetCollection` state transition occurs at each `BeginBlock`. You can read more about it in [04_begin_block.md](04_begin_block.md).
## Accumulate TotalCollectedCoins

`TotalCollectedCoins` are cumulative coins collected in the budget since the bucket was created.
This state transition also occurs at each `BeginBlock` with `BudgetCollection`. You can read more about it in [04_begin_block.md](04_begin_block.md).
