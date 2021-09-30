<!-- order: 3 -->

# State Transitions

This document describes the state transaction operations pertaining to the budget module.

- BudgetCollection
- TotalCollectedCoins
## BudgetCollection

Get all budgets registered in `params.Budgets` and select the valid budgets to collect budgets for the block by its respective plan.
This state transition occurs at each `BeginBlock`. You can read more about it in [04_begin_block.md](04_begin_block.md).

## TotalCollectedCoins

`TotalCollectedCoins` are accumulated coins in a budget since the creation of the budget.
This state transition occurs at each `BeginBlock` with `BudgetCollection`. You can read more about it in [04_begin_block.md](04_begin_block.md).
