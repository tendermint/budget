[![codecov](https://codecov.io/gh/tendermint/budget/branch/master/graph/badge.svg)](https://codecov.io/gh/tendermint/budget?branch=master)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tendermint/budget)](https://pkg.go.dev/github.com/tendermint/budget)

# Budget Module

A budget module is a Cosmos SDK module that implements budget functionality. It is an independent module from other SDK modules and core functionality is to enable anyone to create a budget plan through governance param change proposal. Once it is agreed within the community, voted, and passed, it uses the budget source address to distribute amount of coins by the rate defined in the plan to the collection address. Collecting all budgets and distribution take place every epoch blocks that can be modified by a governance proposal.

One use case is for Gravity DEX farming plan. Budget module can be used to create a budget plan that defines Cosmos Hub's FeeCollector module account where transaction gas fees and part of ATOM inflation are collected as budget source address and uses custom module account (created by budget creator) as collection address. Read [spec docs](./x/budget/spec/01_concepts.md) to get to know more about the module.

⚠ **Budget module v1 is in active development - see "master" branch for the latest update** ⚠

## Installation
### Requirements

If you haven't already, install Golang by following the [official docs](https://golang.org/doc/install). Make sure that your `GOPATH` and `GOBIN` environment variables are properly set up.

Requirement | Notes
----------- | -----------------
Go version  | Go1.16 or higher
Cosmos SDK  | v0.44.2 or higher

### Get Budget Module source code

```bash
git clone https://github.com/tendermint/budget.git
cd budget
make install
```

## Development

### Test

```bash
make test-all
```

### Setup local testnet using script

```bash
# This script bootstraps a single local testnet.
# Note that config, data, and keys are created in the ./data/localnet folder and
# RPC, GRPC, and REST ports are all open.
$ make localnet
```

## Resources

- [Documentation about REST APIs](./docs/How-To/api)
- [Documentation about Command-line Interfaces](./docs/How-To/cli)
- [Demo](./docs/Tutorials/demo/README.md)