[![codecov](https://codecov.io/gh/tendermint/budget/branch/main/graph/badge.svg)](https://codecov.io/gh/tendermint/budget?branch=main)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tendermint/budget)](https://pkg.go.dev/github.com/tendermint/budget)

# Budget Module

The budget module is a Cosmos SDK module that implements budget functionality. It is an independent module from other SDK modules and core functionality is to enable anyone to create a budget plan through governance param change proposal. Once it is agreed within the community, voted, and passed, it uses the source address to distribute amount of coins by the rate defined in the plan to the destination address. Collecting all budgets and distribution take place every epoch blocks that can be modified by a governance proposal.

A primary use case is for Gravity DEX farming plan. The budget module can be used to create a budget plan that defines Cosmos Hub's FeeCollector module account where transaction gas fees and part of ATOM inflation are collected as source address and uses a custom module account (created by budget creator) as destination address. Read [spec docs](./x/budget/spec/01_concepts.md) to get to know more about the module.

## Versions

See the [main](https://github.com/tendermint/budget/tree/main) branch for the latest, and see [releases](https://github.com/tendermint/budget/releases) for the latest release

## Dependencies

If you haven't already, install Golang by following the [official docs](https://golang.org/doc/install). Make sure that your `GOPATH` and `GOBIN` environment variables are properly set up.

Requirement | Notes
----------- | -----------------
Go version  | Go1.16 or higher
Cosmos SDK  | v0.44.3 or higher

### Installation

```bash
# Use git to clone budget module source code and install `budgetd`
git clone https://github.com/tendermint/budget.git
cd budget
make install
```

## Getting Started

To get started to the project, visit the [TECHNICAL-SETUP.md](./TECHNICAL-SETUP.md) docs.

## Documentation

The following documents are available to help you quickly get onboard with the budget module:

- [Technical specification](./x/budget/spec)
- [How to bootstrap a local network with budget module](./docs/Tutorials/localnet)
- [How to use Command Line Interfaces](./docs/How-To/cli)
- [How to use gRPC-gateway REST Routes](./docs/How-To)
- [Demo for how to budget and budget modules](https://github.com/tendermint/farming/blob/main/docs/Tutorials/demo/budget_with_farming.md)

## Contributing

We welcome contributions from everyone. The [main](https://github.com/tendermint/budget/tree/main) branch contains the development version of the code. You can branch of from main and create a pull request, or maintain your own fork and submit a cross-repository pull request. If you're not sure where to start check out [CONTRIBUTING.md](./CONTRIBUTING.md) for our guidelines & policies for how we develop budget module. Thank you to all those who have contributed to budget module!
