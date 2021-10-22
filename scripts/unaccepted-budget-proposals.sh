#!/bin/sh

# Submit a param-change governance proposal
budgetd tx gov submit-proposal param-change ./scripts/configs/budget1.json \
--chain-id localnet \
--from user1 \
--keyring-backend test \
--broadcast-mode block \
--yes \
--output json | jq

# Vote
budgetd tx gov vote 1 yes \
--chain-id localnet \
--from val1 \
--keyring-backend test \
--broadcast-mode block \
--yes \
--output json | jq

# Submit a param-change governance proposal
budgetd tx gov submit-proposal param-change ./scripts/configs/budget2.json \
--chain-id localnet \
--from user1 \
--keyring-backend test \
--broadcast-mode block \
--yes \
--output json | jq

# Vote
budgetd tx gov vote 2 yes \
--chain-id localnet \
--from val1 \
--keyring-backend test \
--broadcast-mode block \
--yes \
--output json | jq

sleep 30s

# This shouldn't be accepted as a proposal since total rate > 1 already
budgetd tx gov submit-proposal param-change ./scripts/configs/budget3.json \
--chain-id localnet \
--from user1 \
--keyring-backend test \
--broadcast-mode block \
--yes \
--output json | jq

# Vote
budgetd tx gov vote 3 yes \
--chain-id localnet \
--from val1 \
--keyring-backend test \
--broadcast-mode block \
--yes \
--output json | jq