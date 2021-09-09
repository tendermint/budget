package types

// NewGenesisState returns new GenesisState.
func NewGenesisState(params Params, records []BudgetRecord) *GenesisState {
	return &GenesisState{
		Params:        params,
		BudgetRecords: records,
	}
}

// DefaultGenesisState returns the default genesis state.
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		DefaultParams(),
		[]BudgetRecord{},
	)
}

// ValidateGenesis validates GenesisState.
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}
	for _, record := range data.BudgetRecords {
		if err := record.TotalCollectedCoins.Validate(); err != nil {
			return err
		}
		if err := ValidateName(record.Name); err != nil {
			return err
		}
	}
	return nil
}
