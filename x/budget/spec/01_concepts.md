<!-- order: 1 -->

# Concepts

## Budget Module

`x/budget` is a simple Cosmos SDK module that implements budget functionality. It is an independent module from other SDK modules and core functionality is to enable anyone to create a budget plan through parameter change governance proposal. Once it is agreed within the community, voted, and passed, it uses `BudgetSourceAddress` to distribute amount of coins relative to the rate defined in the plan to the `CollectionAddress`. At each `BeginBlock`, collecting all budgets and distribution take place every `EpochBlocks`. `EpochBlocks` is a global parameter that can be modified by a governance proposal.

A primary use case is for Gravity DEX farming plan. A budget module can be used to create a budget plan that has `BudgetSourceAddress` for Cosmos Hub's [FeeCollector](https://github.com/cosmos/cosmos-sdk/blob/master/x/auth/types/keys.go#L15) module account which collects transaction gas fees and part of ATOM inflation. Then, `BudgetSourceAddress` plans to distribute some amount of coins to `CollectionAddress` for farming plan. 

### Budget Plan for ATOM Inflation Use Case

Cosmos SDK's `x/distribution` module's current distribution workflow

- In Ante Handler
    - Gas fees are collected in ante handler and they are sent to `FeeCollectorName` module account
  
    - Reference the following lines of code

      +++ https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/auth/ante/fee.go#L112-L135

- In Mint module

  - ATOM inflation is minted in `x/mint` module and they are sent to `FeeCollectorName` module account

  - Reference the following lines of code

    +++ https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/mint/abci.go#L27-L40
    
    +++ https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/mint/keeper/keeper.go#L108-L110

- In distribution module

  - Send all rewards in `FeeCollectorName` to distribution module account
  - From `distributionModuleAccount`, substitute `communityTax`
  - Rest are distributed to proposer and validator reward pools
  - Substituted amount for community budget is saved in kv store

Implementation with Budget Module

  - Implementation Independency
    - A budget module is 100% independent from other existing modules
    - BeginBlock processing order is the following order
      - mint module → budget module → d istribution module
  - Functionalities
    - Distribute ATOM inflation and transaction gas fees to different budget purposes
      - ATOM inflation and gas fees are accumulated in `feeCollectorName` module account
      - Distribute budget amounts from `FeeCollectorName` module account to each budget pool module account
      - Rest amounts stay in `FeeCollectorName` so that distribution module can use it for community fund and staking rewards distribution as what it is doing now
    - Create, modify or remove budget plans via governance
      - Budget plans can be created, modified or removed by **usual parameter governance**




  - Coin Flow with Budget module
    - In Mint Module
      - Atom inflation to `FeeCollectorName` module account
    - In Ante Handler
      - Transaction gas fees to `FeeCollectorName` module account
    - In Budget Module
      - Part of ATOM inflation and gas fees go to different budgets
      - Rest stays in `FeeCollectorName` module account
    - In Distribution Module
      - All amounts in `FeeCollectorName` module account go to community fund and staking rewards

      +++ https://github.com/cosmos/cosmos-sdk/blob/v0.43.0-rc0/x/distribution/keeper/allocation.go#L82-L101