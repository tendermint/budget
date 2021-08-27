package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// budget module sentinel errors
var (
	ErrInvalidBudgetEndTime = sdkerrors.Register(ModuleName, 2, "invalid Budget end time")
	ErrInvalidBudgetName    = sdkerrors.Register(ModuleName, 3, "budget name only allowed letters, digits, and dash without spaces and the maximum length is 50")
	ErrInvalidNameOfAddr    = sdkerrors.Register(ModuleName, 4, "address must be generated in accordance with Rule adr-028 using the name and matched")
	ErrInvalidStartEndTime  = sdkerrors.Register(ModuleName, 5, "budget endtime should not be earlier than budget starttime")
	ErrZeroBudgetRate       = sdkerrors.Register(ModuleName, 6, "budget rate can't be zero")
	ErrOverflowedBudgetRate = sdkerrors.Register(ModuleName, 7, "the total rate of budgets with the same budget source address value should not exceed 1")
	ErrDuplicatedBudgetName = sdkerrors.Register(ModuleName, 8, "budget name can be duplicated")
)
