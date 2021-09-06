package types

import (
	"fmt"
	"regexp"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	reBudgetNameString = fmt.Sprintf(`[a-zA-Z][a-zA-Z0-9-]{0,%d}`, MaxBudgetNameLength-1)
	reBudgetName       = regexp.MustCompile(fmt.Sprintf(`^%s$`, reBudgetNameString))
)

func (budget Budget) String() string {
	out, _ := budget.MarshalYAML()
	return out.(string)
}

// MarshalYAML returns the YAML representation of a Budget.
func (budget Budget) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &budget)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

func (budget Budget) Validate() error {
	//Name only allowed letters(`A-Z, a-z`), digits(`0-9`), and `-` without spaces and the maximum length is 50.
	err := ValidateName(budget.Name)
	if err != nil {
		return err
	}

	// Check that the CollectionAddress is a valid address.
	_, err = ValidityAddr(budget.CollectionAddress)
	if err != nil {
		// TODO: return error with Wrapping
		return err
	}

	// Check that the BudgetSourceAddress is a valid address.
	_, err = ValidityAddr(budget.BudgetSourceAddress)
	if err != nil {
		// TODO: return error with Wrapping
		return err
	}

	// EndTime should not be earlier than StartTime.
	if budget.EndTime.Before(budget.StartTime) {
		return ErrInvalidStartEndTime
	}

	if budget.Rate == sdk.ZeroDec() {
		return ErrZeroBudgetRate
	}
	return nil
}

func (budget Budget) Expired(blockTime time.Time) bool {
	return !budget.EndTime.After(blockTime)
}

func ValidityAddr(bech32 string) (sdk.AccAddress, error) {
	acc, err := sdk.AccAddressFromBech32(bech32)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

// ValidateName is the default validation function for Budget.Name.
func ValidateName(name string) error {
	if !reBudgetName.MatchString(name) {
		return sdkerrors.Wrap(ErrInvalidBudgetName, name)
	}
	return nil
}

type BudgetsBySource struct {
	Budgets   []Budget
	TotalRate sdk.Dec
}

type BudgetsBySourceMap map[string]BudgetsBySource

// GetBudgetsBySourceMap returns a map by BudgetSourceAddress to handle the budgets for the same BudgetSourceAddress together based on the
// same balance when calculating rates for the same BudgetSourceAddress.
func GetBudgetsBySourceMap(budgets []Budget) BudgetsBySourceMap {
	budgetsMap := make(BudgetsBySourceMap)
	for _, budget := range budgets {
		if budgetsBySource, ok := budgetsMap[budget.BudgetSourceAddress]; ok {
			budgetsBySource.TotalRate = budgetsBySource.TotalRate.Add(budget.Rate)
			budgetsBySource.Budgets = append(budgetsBySource.Budgets, budget)
			budgetsMap[budget.BudgetSourceAddress] = budgetsBySource
		} else {
			budgetsMap[budget.BudgetSourceAddress] = BudgetsBySource{
				Budgets:   []Budget{budget},
				TotalRate: budget.Rate,
			}
		}
	}
	return budgetsMap
}
