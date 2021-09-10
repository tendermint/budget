# Budget Module

A budget module is a Cosmos SDK module that implements budget functionality. It is an independent module from other SDK modules and core functionality is to distribute inflation and gas fees to different budget plans. Read [spec docs](./x/budget/spec/01_concepts.md) to get to know more about the module.

⚠ **Budger module v1 is in active development - see "master" branch for the latest update** ⚠

## Installation
### Requirements

If you haven't already, install Golang by following the [official docs](https://golang.org/doc/install). Make sure that your `GOPATH` and `GOBIN` environment variables are properly set up.

Requirement | Notes
----------- | -----------------
Go version  | Go1.15 or higher
Cosmos SDK  | v0.44.0 or higher

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

...