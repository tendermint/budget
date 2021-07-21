package keeper_test

import (
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tax/app"
	"github.com/tendermint/tax/x/tax/types"
)

// createTestApp returns a tax app with custom TaxKeeper.
func createTestApp(isCheckTx bool) (*app.TaxApp, sdk.Context) {
	app := app.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.TaxKeeper.SetParams(ctx, types.DefaultParams())

	return app, ctx
}
