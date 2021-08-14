<!-- order: 2 -->

# State

## Tax

The Tax structure is not stored in kv, but in parameters in the tax module as taxes.

```go
type Tax struct {
	Name                  string // name defines the name of the tax
	Rate                  sdk.Dec // rate specifies the distributing amount by ratio of total tax source
	TaxSourceAddress      string // tax_source_address defines the bech32-encoded address that source of the tax
	CollectionAddress     string // collection_address defines the bech32-encoded address of the tax pool to distribute
	StartTime             time.Time // start_time specifies the start time of the tax
	EndTime               time.Time // end_time specifies the end time of the tax
}
```
