<!-- order: 6 -->

# Parameters

The budget module contains the following parameters:


| Key         | Type     | Example                                                                              |
| ----------- | -------- | ------------------------------------------------------------------------------------ |
| EpochBlocks | uint32   | {"epoch_blocks":1}                                                                   |
| Budgets     | []Budget | {"budgets":[{"name":"liquidity-farming-20213Q-20221Q","rate":"0.300000000000000000","budget_source_address":"cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta","collection_address":"cosmos10pg34xts7pztyu9n63vsydujjayge7gergyzavl4dhpq36hgmkts880rwl","start_time":"2021-10-01T00:00:00Z","end_time":"2022-04-01T00:00:00Z"}]}                                        |

## EpochBlocks

The universal epoch length in number of blocks.
Every process for budget collecting is executed with this `epoch_blocks` frequency.

The default value is 1 and all budget collections are disabled if the value is 0. Budget collection logic is executed with the following condition. 

```
params.EpochBlocks > 0 && Current Block Height % params.EpochBlocks == 0
```

You can reference [the line of the code](https://github.com/tendermint/budget/blob/main/x/budget/keeper/budget.go#L78).

## Budgets

The Budget structure can be found in [02_state.md](02_state.md).

Parameter of a budget can be added, modified, and deleted through [parameter change governance proposal](https://docs.cosmos.network/master/modules/gov/01_concepts.html#proposal-submission).

An example of how to add a budget plan can be found in this [docs/How-To/cli](../../../docs/How-To/cli) guide. 

### Validity Checks

- Budget name only allows letters(`A-Z, a-z`), digits(`0-9`), and `-` without spaces. Also, it has the maximum length of 50 and it should not be duplicate with the existing budget names.

- Validate `CollectionAddress` address.

- Validate `BudgetSourceAddress` address.

- EndTime should not be earlier than StartTime.

- The total rate of budgets with the same `BudgetSourceAddress` value should not exceed 1 (100%).

+++ https://github.com/tendermint/budget/blob/main/x/budget/types/budget.go#L33-L63