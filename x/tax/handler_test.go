package tax_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	taxapp "github.com/tendermint/tax/app"
	"github.com/tendermint/tax/x/tax/keeper"
	"github.com/tendermint/tax/x/tax/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// createTestInput returns a simapp with custom TaxKeeper
// to avoid messing with the hooks.
func createTestInput() (*taxapp.TaxApp, sdk.Context, []sdk.AccAddress) {
	app := taxapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.TaxKeeper = keeper.NewKeeper(
		app.AppCodec(),
		app.GetKey(types.StoreKey),
		app.GetSubspace(types.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.DistrKeeper,
		map[string]bool{},
	)

	addrs := taxapp.AddTestAddrs(app, ctx, 1, sdk.NewInt(100000))

	return app, ctx, addrs
}
