package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// tax module sentinel errors
var (
	ErrInvalidTaxEndTime = sdkerrors.Register(ModuleName, 2, "invalid Tax end time")
)
