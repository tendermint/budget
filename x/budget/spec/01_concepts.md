<!-- order: 1 -->

# Concepts

## Budget Module

`x/budget` is a Cosmos SDK module that implements budget functionality.

### 1. Atom Inflation Distribution case

- Current : distribution module reward flow

  1. Gas fees collected in ante handler, sent to `feeCollectorName` module account

  2. Atom inflation minted in mint module, sent to `feeCollectorName` module account

  3. In distribution module

  a) Send all rewards in `feeCollectorName` to distribution module account

  b) From `distributionModuleAccount`, substitute `communityTax`

  c) Rest are distributed to proposer and validator reward pools

  d) Substituted amount for community budget is saved in kv store

- **Implementation with Budget Module**

  - Implementation Independency
    - Budget Module is **100% independent** from other existing modules
      - We don’t need to change **any module** at all to adopt Budget Module!
      - Budget Module even works **without** Distribution Module or Mint Module!
    - Begin Block Processing Order : Mint → **(Budget)** → Distribution
  - Functionalities
    - Distribute Atom inflation and gas fees to different budget purposes
      - Atom inflation and gas fees are accumulated in `feeCollectorName` module account
      - Distribute budget amounts from `feeCollectorName` module account to **each budget pool** module account
      - Rest amounts **stay** in `feeCollectorName` so that “Distribution Module” can use it for community fund and staking rewards distribution
    - Create, modify or remove budget plans via governance
      - Budget plans can be created, modified or removed by **usual parameter governance**
  - Coin Flow with Budget module
    - In Mint Module
      - Atom inflation to `feeCollectorName` module account
      - [https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/mint/abci.go#L27-L40](https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/mint/abci.go#L27-L40)
      - [https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/mint/keeper/keeper.go#L108-L110](https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/mint/keeper/keeper.go#L108-L110)
    - In Ante Handler
      - Gas fees to `feeCollectorName` module account
      - [https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/auth/ante/fee.go#L112-L135](https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/auth/ante/fee.go#L112-L135)
    - In Budget Module
      - Part of Atom inflation and gas fees go to different budget pools
      - Rest stays within `feeCollectorName` module account
    - In Distribution Module
      - All amounts in `feeCollectorName` module account go to community fund and staking rewards
      - [https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/distribution/keeper/allocation.go#L82-L101](https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/distribution/keeper/allocation.go#L82-L101)

- budget module parameter example

  ```json
  {
    "budget": {
      "params": {
        "epoch_blocks": 1,
        "budgets": [
          {
            "name": "liquidity-farming-20213Q-20221Q",
            "rate": "0.300000000000000000",
            "budget_source_address": "cosmos17xpfvakm2amg962yls6f84z3kell8c5lserqta", // Address corresponding to fee_collector module account in cosmoshub case
            "collection_address": "cosmos10pg34xts7pztyu9n63vsydujjayge7gergyzavl4dhpq36hgmkts880rwl", // 32-bytes address case, sdk.AccAddress(address.Module("farming", []byte("FarmingBudget"))).String()
            "start_time": "2021-10-01T00:00:00Z",
            "end_time": "2022-04-01T00:00:00Z"
          }
        ]
      }
    }
  }
  ```
