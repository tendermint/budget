package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// tax module sentinel errors
var (
	ErrInvalidTaxEndTime   = sdkerrors.Register(ModuleName, 2, "invalid Tax end time")
	ErrInvalidTaxName      = sdkerrors.Register(ModuleName, 3, "tax name should be connected to - without spaces and must be within 30 characters")
	ErrInvalidNameOfAddr   = sdkerrors.Register(ModuleName, 4, "address must be generated in accordance with Rule adr-028 using the name and matched")
	ErrInvalidStartEndTime = sdkerrors.Register(ModuleName, 5, "tax endtime should not be earlier than tax starttime")
	ErrZeroTaxRate         = sdkerrors.Register(ModuleName, 6, "tax rate can't be zero")
)
