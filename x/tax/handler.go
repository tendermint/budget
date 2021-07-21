package tax

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/tendermint/tax/x/tax/keeper"
)

func NewUpdateTaxProposal(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		//case *types.UpdateTaxesProposal:
		//	return keeper.HandleUpdateTaxesProposal(ctx, k, c.Taxes)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized tax proposal content type: %T", c)
		}
	}
}
