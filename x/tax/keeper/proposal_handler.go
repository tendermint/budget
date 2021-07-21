package keeper

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: not implemented

// HandleUpdateTaxesProposal is a handler for updating taxes proposal.
func HandleUpdateTaxesProposal(ctx sdk.Context, k Keeper, taxesAny []*codectypes.Any) error {
	//taxes, err := types.UnpackTaxes(taxesAny)
	//if err != nil {
	//	return err
	//}
	//
	//for _, tax := range taxes {
	//
	//}

	return nil
}
