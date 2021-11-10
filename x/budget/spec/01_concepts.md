<!-- order: 1 -->

# Concepts

## Budget Module

`x/budget` is a simple Cosmos SDK module that implements budget functionality. It is an independent module from other SDK modules and core functionality is to enable anyone to create a budget plan through parameter change governance proposal. Once it is agreed within the community, voted, and passed, it uses `SourceAddress` to distribute amount of coins relative to the rate defined in the plan to the `DestinationAddress`. At each `BeginBlock`, collecting all budgets and distribution take place every `EpochBlocks`. `EpochBlocks` is a global parameter that can be modified by a governance proposal.

A primary use case is for Gravity DEX farming plan. A budget module can be used to create a budget plan that has `SourceAddress` for Cosmos Hub's [FeeCollector](https://github.com/cosmos/cosmos-sdk/blob/v0.44.0/x/auth/types/keys.go#L15) module account which collects transaction gas fees and part of ATOM inflation. Then, `SourceAddress` plans to distribute some amount of coins to `DestinationAddress` for farming plan.

### Budget Plan for ATOM Inflation Use Case

Cosmos SDK's current reward workflow

- In AnteHandler

    - Gas fees are collected in ante handler and they are sent to `FeeCollectorName` module account

    - Reference the following lines of code

      +++ https://github.com/cosmos/cosmos-sdk/blob/v0.44.0/x/auth/ante/fee.go#L112-L140

- In `x/mint` module

  - ATOM inflation is minted in `x/mint` module and they are sent to `FeeCollectorName` module account

  - Reference the following lines of code

    +++ https://github.com/cosmos/cosmos-sdk/blob/v0.44.0/x/mint/abci.go#L27-L40

    +++ https://github.com/cosmos/cosmos-sdk/blob/v0.44.0/x/mint/keeper/keeper.go#L108-L110

- In `x/distribution` module

  - Send all rewards in `FeeCollectorName` to distribution module account
  
  - From `distributionModuleAccount`, substitute `communityTax`

  - Rest are distributed to proposer and validator reward pools

  - Substituted amount for community budget is saved in kv store

  - Reference the following lines of code

    +++ https://github.com/cosmos/cosmos-sdk/blob/v0.44.0/x/distribution/keeper/allocation.go#L13-L102

Implementation with Budget Module

  - A budget module is 100% independent of other Cosmos SDK's existing modules

  - BeginBlock processing order is the following order

      - mint module → budget module → distribution module

  - Distribute ATOM inflation and transaction gas fees to different budget purposes

    - ATOM inflation and gas fees are accumulated in `FeeCollectorName` module account

    - Distribute budget amounts from `FeeCollectorName` module account to each budget pool module account

    - Rest amounts stay in `FeeCollectorName` so that distribution module can use it for community fund and staking rewards distribution as what it is doing now

  - Create, modify or remove budget plans via governance process
    - A budget plan can be created, modified or removed by parameter change governance proposal
