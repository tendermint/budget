<!-- order: 2 -->

# State

## Tax

The Tax structure is not stored in kv, but in parameters in the tax module as taxes.

```go
type Tax struct {
	Name                  string // name defines the name of the tax
	Rate                  sdk.Dec // rate specifies the distributing amount by ratio of total tax source
	CollectionAddress     string // collection_address defines the bech32-encoded address of the tax pool to distribute
	CollectionAccountName string // collection_account_name is module account name of collection_address, fill in this field optionally if you need to define a new module account or want to verify that address matches collection_address that module name
	TaxSourceAddress      string // tax_source_address defines the bech32-encoded address that source of the tax
	TaxSourceAccountName  string // tax_source_account_name is module account name of tax_source_address, fill in this field optionally if you need to define a new module account or want to verify that address matches tax_source_address that module name
	StartTime             time.Time // start_time specifies the start time of the tax
	EndTime               time.Time // end_time specifies the end time of the tax
}
```
