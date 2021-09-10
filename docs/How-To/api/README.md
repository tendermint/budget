---
Title: REST APIs
Description: A high-level overview of how the REST API interfaces work for the budget module.
---

## Table of Contetns

- [Prerequisite](#Prerequisite)
- [REST APIs](#REST-APIs)
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

# Start
$BINARY start
```

## REST APIs

Query the values set as budget parameters
http://localhost:1317/cosmos/budget/v1beta1/params

```json
{
  "params": {
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
}
```

Query all the budget plans exist in the network
http://localhost:1317/cosmos/budget/v1beta1/budgets

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
          "amount": "66785"
        }
      ]
    }
  ]
}
```