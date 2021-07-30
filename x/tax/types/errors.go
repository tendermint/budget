package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// tax module sentinel errors
var (
	ErrInvalidTaxEndTime   = sdkerrors.Register(ModuleName, 2, "invalid Tax end time")
	ErrInvalidTaxName      = sdkerrors.Register(ModuleName, 3, "tax name only allowed letters, digits, and dash without spaces and the maximum length is 50")
	ErrInvalidNameOfAddr   = sdkerrors.Register(ModuleName, 4, "address must be generated in accordance with Rule adr-028 using the name and matched")
	ErrInvalidStartEndTime = sdkerrors.Register(ModuleName, 5, "tax endtime should not be earlier than tax starttime")
	ErrZeroTaxRate         = sdkerrors.Register(ModuleName, 6, "tax rate can't be zero")
	ErrOverflowedTaxRate   = sdkerrors.Register(ModuleName, 7, "the total rate of taxes with the same tax source address value should not exceed 1")
	ErrDuplicatedTaxName   = sdkerrors.Register(ModuleName, 8, "tax name can be duplicated")
)
