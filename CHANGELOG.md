<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking Protobuf, gRPC and REST routes used by end-users.
"CLI Breaking" for breaking CLI commands.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.
Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## v1.0.0 - 2021-11-26

* [\#64](https://github.com/tendermint/budget/pull/64) docs: improve documentation for audit release
* [\#74](https://github.com/tendermint/budget/pull/74) fix: update docs and workflow for default main branch
* [\#66](https://github.com/tendermint/budget/pull/66) fix: Add govHandler on testcode for budget proposal and Fix Expired rule
* [\#79](https://github.com/tendermint/budget/pull/79) fix: panic instead of ignoring errors
* [\#83](https://github.com/tendermint/budget/pull/83) fix: validation totalRate to check date overlapped budgets
* [\#85](https://github.com/tendermint/budget/pull/85) fix: rename some fields of Budget
* [\#81](https://github.com/tendermint/budget/pull/81) build: bump cosmos-sdk version to 0.44.3
* [\#90](https://github.com/tendermint/budget/pull/90) feat: add address endpoint and release swagger v1.0.0
* [\#91](https://github.com/tendermint/budget/pull/91) test: add test codes using gov handler

## [v0.1.1](https://github.com/tendermint/budget/releases/tag/v0.1.1) - 2021-10-15

* [\#46](https://github.com/tendermint/budget/pull/46) feat: add emit events and update spec docs
* [\#54](https://github.com/tendermint/budget/pull/54) test: update simulation tests
* [\#55](https://github.com/tendermint/budget/pull/55) build: bump Cosmos SDK version to v0.44.2
* [\#57](https://github.com/tendermint/budget/pull/57) fix: refine sentinel errors
* [\#62](https://github.com/tendermint/budget/pull/62) feat: refactor collectbudget

## [v0.1.0](https://github.com/tendermint/budget/releases/tag/v0.1.0) - 2021-09-17