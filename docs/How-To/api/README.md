---
Title: REST APIs
Description: A high-level overview of gRPC-gateway REST Routes in budget module.
---

## Swagger Documentation

- [Swagger Docs v0.1.0](https://app.swaggerhub.com/apis-docs/gravity-devs/budget/0.1.0)

## gRPC-gateway REST Routes

In order to test out the following REST routes, you need to set up a local node to query from. You can refer to this [localnet tutorial](../../Tutorials/localnet) on how to build `budgetd` binary and bootstrap a local network in your local machine.

Query the address of `fee_collector` with address type 1(`AddressType20Bytes`)
http://localhost:1317/cosmos/budget/v1beta1/addresses/fee_collector?type=1 <!-- markdown-link-check-disable-line -->

```json
{
  "address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta"
}
```

Query the address of `GravityDEXFarmingBudget` on farming module with default address type 0(`AddressType32Bytes`)
http://localhost:1317/cosmos/budget/v1beta1/addresses/GravityDEXFarmingBudget?module_name=farming <!-- markdown-link-check-disable-line -->

```json
{
  "address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky"
}
```

Query the values set as budget parameters
http://localhost:1317/cosmos/budget/v1beta1/params <!-- markdown-link-check-disable-line -->

```json
{
  "params": {
    "epoch_blocks": 1,
    "budgets": [
      {
        "name": "gravity-dex-farming-20213Q-20221Q",
        "rate": "0.300000000000000000",
        "source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
        "destination_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
        "start_time": "2021-10-01T00:00:00Z",
        "end_time": "2022-04-01T00:00:00Z"
      }
    ]
  }
}
```

Query all the budget plans exist in the network
http://localhost:1317/cosmos/budget/v1beta1/budgets <!-- markdown-link-check-disable-line -->

```json
{
  "budgets": [
    {
      "budget": {
        "name": "gravity-dex-farming-20213Q-20221Q",
        "rate": "0.300000000000000000",
        "source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
        "destination_address": "cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky",
        "start_time": "2021-10-01T00:00:00Z",
        "end_time": "2022-04-01T00:00:00Z"
      },
      "total_collected_coins": [
        {
          "denom": "stake",
          "amount": "66785"
        }
      ]
    }
  ]
}
```