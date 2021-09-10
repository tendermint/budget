---
Title: Budgetd
Description: A high-level overview of how the command-line (CLI) interfaces work for the budget module.
---

# Budgetd

This document provides a high-level overview of how the command-line (CLI) interfaces work for the budget module.

## Table of Contetns

- [Prerequisite](#Prerequisite)
- [Command-line Interfaces](#Command-Line-Interfaces)

## Prerequisite 

### Build

```bash
git clone https://github.com/tendermint/budget.git
cd budget
make install
```

### Boostrap

In order to test out the command-line interface, you need to boostrap local network by using the commands below.

```bash
# Configure variables
export BINARY=budgetd
export HOME_BUDGETAPP=$HOME/.budgetapp
export CHAIN_ID=localnet
export VALIDATOR_1="struggle panic room apology luggage game screen wing want lazy famous eight robot picture wrap act uphold grab away proud music danger naive opinion"
export USER_1="guard cream sadness conduct invite crumble clock pudding hole grit liar hotel maid produce squeeze return argue turtle know drive eight casino maze host"
export USER_2="fuel obscure melt april direct second usual hair leave hobby beef bacon solid drum used law mercy worry fat super must ritual bring faculty"
export VALIDATOR_1_GENESIS_COINS=10000000000stake,10000000000uatom,10000000000uusd
export USER_1_GENESIS_COINS=10000000000stake,10000000000uatom,10000000000uusd
export USER_2_GENESIS_COINS=10000000000stake,10000000000poolD35A0CC16EE598F90B044CE296A405BA9C381E38837599D96F2F70C2F02A23A4

# Bootstrap
$BINARY init $CHAIN_ID --chain-id $CHAIN_ID
echo $VALIDATOR_1 | $BINARY keys add val1 --keyring-backend test --recover
echo $USER_1 | $BINARY keys add user1 --keyring-backend test --recover
echo $USER_2 | $BINARY keys add user2 --keyring-backend test --recover
$BINARY add-genesis-account $($BINARY keys show val1 --keyring-backend test -a) $VALIDATOR_1_GENESIS_COINS
$BINARY add-genesis-account $($BINARY keys show user1 --keyring-backend test -a) $USER_1_GENESIS_COINS
$BINARY add-genesis-account $($BINARY keys show user2 --keyring-backend test -a) $USER_2_GENESIS_COINS
$BINARY gentx val1 100000000stake --chain-id $CHAIN_ID --keyring-backend test
$BINARY collect-gentxs

# Modify app.toml
sed -i '' 's/enable = false/enable = true/g' $HOME_BUDGETAPP/config/app.toml
sed -i '' 's/swagger = false/swagger = true/g' $HOME_BUDGETAPP/config/app.toml

# Modify parameters for the governance proposal
sed -i '' 's%"amount": "10000000"%"amount": "1"%g' $HOME_BUDGETAPP/config/genesis.json
sed -i '' 's%"quorum": "0.334000000000000000",%"quorum": "0.000000000000000001",%g' $HOME_BUDGETAPP/config/genesis.json
sed -i '' 's%"threshold": "0.500000000000000000",%"threshold": "0.000000000000000001",%g' $HOME_BUDGETAPP/config/genesis.json
sed -i '' 's%"voting_period": "172800s"%"voting_period": "30s"%g' $HOME_BUDGETAPP/config/genesis.json

# Start
$BINARY start
```
## Command-Line Interfaces

- [Transaction](#Transaction)
- [Query](#Query)
    * [Params](#Params)
    * [Budgets](#Budgets)

## Transaction

There is no command-line interface for the Budget module. However, in order to query budget parameters and plans we are going to submit a governance proposal to create a budget plan in this documentation.

### Create a Budget Plan

Let's create `proposal.json` file. Depending on what budget plan you plan to create, change the following values of the fields for your need. In this case, we plan to create a budget plan that distributes partial amount from the ATOM inflation for Gravity DEX farming plan. 

- `name`: is the name of the budget plan used for display
- `description`: is the budget plan's description used for display
- `rate`: is the distributing amount by ratio of the total budget source
- `budget_source_address`: is the farming module's `farmingPoolAddr` address
- `collection_address`: is the distribution module's `feeCollector` module account address
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
# Submit a governance proposal
budgetd tx gov submit-proposal param-change proposal.json \
--chain-id localnet \
--from user1 \
--keyring-backend test \
--broadcast-mode block \
--yes

# Query the proposal to check the status
# the status should be PROPOSAL_STATUS_VOTING_PERIOD
budgetd q gov proposals --output json | jq

# Vote
budgetd tx gov vote 1 yes \
--chain-id localnet \
--from val1 \
--keyring-backend test \
--broadcast-mode block \
--yes

# Wait a while (30s) for the proposal to pass
# Query the proposal again to check the status
# the status should be PROPOSAL_STATUS_PASSED
budgetd q gov proposals --output json | jq
 
# Query the balances of collection_address for a few times to see
# if its balances increase over time
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
