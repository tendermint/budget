<!-- order: 6 -->

# Parameters

The tax module contains the following parameters:

| Key   | Type  | Example                                                                                                                                                                                                                                                                                                                                                                      |
| ----- | ----- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Taxes | []Tax | {"taxes":[{"name":"liquidity_farming_20213Q-20221Q","rate":"0.300000000000000000","collection_address":"cosmos1...","collection_account_name":"targetModule/targetModuleAccountName","tax_source_address":"cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta","tax_source_account_name":"fee_collector","start_time":"2021-10-01T00:00:00Z","end_time":"2022-04-01T00:00:00Z"}]} |

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
          "name": "testTax",
          "rate": "0.100000000000000000",
          "collection_address": "cosmos1w3jhymtfdeshg6t0deqkgerjrhy0ma",
          "tax_source_address": "cosmos1w3jhymtfdeshg6t0deqkgerjrhy0ma",
          "start_time": "2021-07-27T19:35:52.392809Z",
          "end_time": "2021-07-27T19:35:52.39281Z"
        }
      ]
    }
  ],
  "deposit": "10000000uatom"
}
```


### Validity Checks

- Name is connected to `-` without spaces and must be within 30 characters
- The total rate of Taxes with the same `TaxSourceAddress` value should not exceed 1.
- EndTime should not be earlier than StartTime.
- Check that the `CollectionAddress` is a valid address, if there is a `CollectionAccountName` value, it must be generated in accordance with Rule [ADR-028(moudule-account)](https://github.com/cosmos/cosmos-sdk/blob/master/docs/architecture/adr-028-public-key-addresses.md#module-account-addresses) and matched to `CollectionAddress`. 
- Check that the `TaxSourceAddress` is a valid address, if there is a `TaxSourceAccountName` value, it must be generated in accordance with Rule [ADR-028(moudule-account)](https://github.com/cosmos/cosmos-sdk/blob/master/docs/architecture/adr-028-public-key-addresses.md#module-account-addresses) and matched to `TaxSourceAddress`.
- Only need to fill CollectionAccountName or TaxSourceAccountName if you need to create and define a new Module Account. All fields except this are required