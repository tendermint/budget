<!-- order: 6 -->

# Parameters

The tax module contains the following parameters:

| Key         | Type   | Example                                                                                                                                                                                                                                                                                                                  |
| ----------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| EpochBlocks | uint32 | {"epoch_blocks":1}                                                                                                                                                                                                                                                                                                       |
| Taxes       | []Tax  | {"taxes":[{"name":"liquidity-farming-20213Q-20221Q","rate":"0.300000000000000000","tax_source_address":"cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta","collection_address":"cosmos10pg34xts7pztyu9n63vsydujjayge7gergyzavl4dhpq36hgmkts880rwl","start_time":"2021-10-01T00:00:00Z","end_time":"2022-04-01T00:00:00Z"}]} |

## EpochBlocks

The universal epoch length in number of blocks.
Every process for tax collecting is executed with this epoch_blocks frequency.

the default value is 1 and when the value is 0, all tax collections are disabled.
when current height % `params.EpochBlocks` == 0  the tax collection logic is executed.

## Taxes

The structure of the Tax can be found at [State](02_state.md).

Taxes parameter can be added, deleted, and modified through [gov.ParameterChangeProposal](https://docs.cosmos.network/master/modules/gov/01_concepts.html#proposal-submission), and for each purpose, the changes in the existing tax list should be applied and set.

Below is an example of adding tax.

`taxd tx gov submit-proposal param-change proposal.json`

proposal.json

```json
{
  "title": "Add a example tax",
  "description": "This proposition is an example of adding Taxes using ParameterChangeProposal.",
  "changes": [
    {
      "subspace": "tax",
      "key": "Taxes",
      "value": [
        {
          "name": "liquidity-farming-20213Q-20221Q",
          "rate": "0.300000000000000000",
          "tax_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
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
- The total rate of Taxes with the same `TaxSourceAddress` value should not exceed 1.
- EndTime should not be earlier than StartTime.
- Check that the `CollectionAddress` is a valid address.
- Check that the `TaxSourceAddress` is a valid address.
