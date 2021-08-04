package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/tendermint/tax/x/tax/types"
)

func (k Keeper) TaxCollection(ctx sdk.Context) error {
	// Get all the Taxes registered in params.Taxes and select only the valid taxes. If there is no valid tax, exit.
	taxes := k.CollectibleTaxes(ctx)
	if len(taxes) == 0 {
		return nil
	}

	var inputs []banktypes.Input
	var outputs []banktypes.Output
	sendCoins := func(from, to sdk.AccAddress, coins sdk.Coins) {
		if !coins.Empty() && coins.IsValid() {
			inputs = append(inputs, banktypes.NewInput(from, coins))
			outputs = append(outputs, banktypes.NewOutput(to, coins))
		}
	}
	// Create a map by TaxSourceAddress to handle the taxes for the same TaxSourceAddress together based on the
	// same balance when calculating rates for the same TaxSourceAddress.
	taxesByTaxSourceMap := types.GetTaxesByTaxSourceMap(taxes)
	for taxSource, taxesByTaxSource := range taxesByTaxSourceMap {
		taxSourceAcc, err := sdk.AccAddressFromBech32(taxSource)
		if err != nil {
			continue
		}
		taxSourceBalances := sdk.NewDecCoinsFromCoins(k.bankKeeper.GetAllBalances(ctx, taxSourceAcc)...)
		expectedCollectionCoins, expectedChangeCoins := taxSourceBalances.MulDecTruncate(taxesByTaxSource.TotalRate).TruncateDecimal()
		var totalCollectionCoins sdk.Coins
		var totalChangeCoins sdk.DecCoins
		for _, tax := range taxesByTaxSource.Taxes {
			collectionAcc, err := sdk.AccAddressFromBech32(tax.CollectionAddress)
			if err != nil {
				continue
			}
			collectionCoins, changeCoins := taxSourceBalances.MulDecTruncate(tax.Rate).TruncateDecimal()
			totalCollectionCoins = totalCollectionCoins.Add(collectionCoins...)
			totalChangeCoins = totalChangeCoins.Add(changeCoins...)
			sendCoins(taxSourceAcc, collectionAcc, collectionCoins)
		}
		// temporary validation logic
		if totalCollectionCoins.IsAnyGT(expectedCollectionCoins) {
			panic("totalCollectionCoins.IsAnyGT(expectedCollectionCoins)")

		}
		if expectedChangeCoins.Sub(totalChangeCoins).IsAnyNegative() {
			panic("expectedChangeCoins.Sub(totalChangeCoins).IsAnyNegative()")
		}
	}
	if err := k.bankKeeper.InputOutputCoins(ctx, inputs, outputs); err != nil {
		return err
	}
	// TODO: add metric or record total collection coins each tax
	return nil
}

// Get all the Taxes registered in params.Taxes and return only the valid and not expired taxes
func (k Keeper) CollectibleTaxes(ctx sdk.Context) (taxes []types.Tax) {
	params := k.GetParams(ctx)
	for _, tax := range params.Taxes {
		err := tax.Validate()
		if err == nil && !tax.Expired(ctx.BlockTime()) {
			taxes = append(taxes, tax)
		}
	}
	return
}
