<!-- order: 6 -->

# Parameters

The budget module contains the following parameters:

| Key         | Type     | Example                                                                                                                                                                                                                                                                                                                       |
| ----------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| EpochBlocks | uint32   | {"epoch_blocks":1}                                                                                                                                                                                                                                                                                                            |
| Budgets     | []Budget | {"budgets":[{"name":"liquidity-farming-20213Q-20221Q","rate":"0.300000000000000000","budget_source_address":"cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta","collection_address":"cosmos10pg34xts7pztyu9n63vsydujjayge7gergyzavl4dhpq36hgmkts880rwl","start_time":"2021-10-01T00:00:00Z","end_time":"2022-04-01T00:00:00Z"}]} |

## EpochBlocks

The universal epoch length in number of blocks.
Every process for budget collecting is executed with this epoch_blocks frequency.

the default value is 1 and when the value is 0, all budget collections are disabled.
when current height % `params.EpochBlocks` == 0  the budget collection logic is executed.

## Budgets

The structure of the Budget can be found at [State](02_state.md).

Budgets parameter can be added, deleted, and modified through [gov.ParameterChangeProposal](https://docs.cosmos.network/master/modules/gov/01_concepts.html#proposal-submission), and for each purpose, the changes in the existing budget list should be applied and set.

Below is an example of adding budget.

`budgetd tx gov submit-proposal param-change proposal.json`

proposal.json

```json
{
  "title": "Add a example budget",
  "description": "This proposition is an example of adding Budgets using ParameterChangeProposal.",
  "changes": [
    {
      "subspace": "budget",
      "key": "Budgets",
      "value": [
        {
          "name": "liquidity-farming-20213Q-20221Q",
          "rate": "0.300000000000000000",
          "budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
          "collection_address": "cosmos10pg34xts7pztyu9n63vsydujjayge7gergyzavl4dhpq36hgmkts880rwl",
          "start_time": "2021-10-01T00:00:00Z",
          "end_time": "2022-04-01T00:00:00Z"
        }
      ]
    }
  ],
  "deposit": "10000000uatom"
}
```

### Validity Checks

- Name only allowed letters(`A-Z, a-z`), digits(`0-9`), and `-` without spaces and the maximum length is 50 and not duplicated.
- The total rate of Budgets with the same `BudgetSourceAddress` value should not exceed 1.
- EndTime should not be earlier than StartTime.
- Check that the `CollectionAddress` is a valid address.
- Check that the `BudgetSourceAddress` is a valid address.
