---
Title: Budgetd
Description: A high-level overview of how the command-line (CLI) interfaces work for the budget module.
---

# Budgetd

This document provides a high-level overview of how the command-line (CLI) interfaces work for the budget module.

## Command-Line Interfaces

In order to test out the following command-line interfaces, you need to set up a local node to either send transaction or query from. You can refer to this [localnet tutorial](./Tutorials/localnet) on how to build `budgetd` binary and bootstrap a local network in your local machine.

- [Transaction N/A](#Transaction)
- [Query](#Query)
    * [Params](#Params)
    * [Budgets](#Budgets)

## Transaction

There is no command-line interface for the Budget module. However, in order to query budget parameters and plans we are going to submit a governance proposal to create a budget plan in this documentation.

### Create a Budget Plan

Let's create `proposal.json` file. Depending on what budget plan you plan to create, change the following values of the fields for your need. In this case, we plan to create a budget plan that distributes partial amount of coins from the Cosmos Hub's gas fees and ATOM inflation accrued in [FeeCollector](https://github.com/cosmos/cosmos-sdk/blob/master/x/auth/types/keys.go#L15) module account for Gravity DEX farming plan to GravityDEXFarmingBudget account (see the code below)

```go
sdk.AccAddress(address.Module(ModuleName, []byte("GravityDEXFarmingBudget")))
// cosmos1228ryjucdpdv3t87rxle0ew76a56ulvnfst0hq0sscd3nafgjpqqkcxcky
```

- `name`: is the name of the budget plan used for display
- `description`: is the description of the budget plan used for display
- `rate`: is the distributing amount by ratio of the total budget source
- `budget_source_address`: is the address where the source of budget comes from
- `collection_address`: is the address that collects budget from the budget source address 
- `start_time`: is start time of the budget plan 
- `end_time`: is end time of the budget plan

```json
{
  "title": "Create a Budget Plan",
  "description": "Here is an example of how to add a budget plan by using ParameterChangeProposal",
  "changes": [
    {
      "subspace": "budget",
      "key": "Budgets",
      "value": [
        {
          "name": "gravity-dex-farming-20213Q-20221Q",
          "rate": "0.300000000000000000",
          "budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
          "collection_address": "cosmos10pg34xts7pztyu9n63vsydujjayge7gergyzavl4dhpq36hgmkts880rwl",
          "start_time": "2021-10-01T00:00:00Z",
          "end_time": "2022-04-01T00:00:00Z"
        }
      ]
    }
  ],
  "deposit": "10000000stake"
}
```

```bash
# Submit a parameter changes proposal to create a budget plan
budgetd tx gov submit-proposal param-change proposal.json \
--chain-id localnet \
--from user1 \
--keyring-backend test \
--broadcast-mode block \
--yes

# Query the proposal to check the status PROPOSAL_STATUS_VOTING_PERIOD
budgetd q gov proposals --output json | jq

# Vote
budgetd tx gov vote 1 yes \
--chain-id localnet \
--from val1 \
--keyring-backend test \
--broadcast-mode block \
--yes

#
# Wait a while (30s) for the proposal to pass
#

# Query the proposal again to check the status PROPOSAL_STATUS_PASSED
budgetd q gov proposals --output json | jq
 
# Query the balances of collection_address for a couple times 
# the balances should increrase over time as gas fees and part of ATOM inflation flow in
budgetd q bank balances cosmos10pg34xts7pztyu9n63vsydujjayge7gergyzavl4dhpq36hgmkts880rwl --output json | jq
```

## Query

https://github.com/tendermint/budget/blob/master/proto/tendermint/budget/v1beta1/query.proto

### Params 

```bash
# Query the values set as budget parameters
# Note that default params are empty. You need to submit governance proposal to create budget plan
# Reference the Transaction section in thid documentation
budgetd q budget params --output json | jq
```

```json
{
  "epoch_blocks": 1,
  "budgets": [
    {
      "name": "gravity-dex-farming-20213Q-20221Q",
      "rate": "0.300000000000000000",
      "budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
      "collection_address": "cosmos10pg34xts7pztyu9n63vsydujjayge7gergyzavl4dhpq36hgmkts880rwl",
      "start_time": "2021-10-01T00:00:00Z",
      "end_time": "2022-04-01T00:00:00Z"
    }
  ]
}
```

### Budgets

```bash
# Query all the budget plans exist in the network
budgetd q budget budgets --output json | jq
```

```json
{
  "budgets": [
    {
      "budget": {
        "name": "gravity-dex-farming-20213Q-20221Q",
        "rate": "0.300000000000000000",
        "budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta",
        "collection_address": "cosmos10pg34xts7pztyu9n63vsydujjayge7gergyzavl4dhpq36hgmkts880rwl",
        "start_time": "2021-10-01T00:00:00Z",
        "end_time": "2022-04-01T00:00:00Z"
      },
      "total_collected_coins": [
        {
          "denom": "stake",
          "amount": "2220"
        }
      ]
    }
  ]
}
```
