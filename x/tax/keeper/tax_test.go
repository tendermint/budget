package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/tax/x/tax/types"
)

func (suite *KeeperTestSuite) TestTaxCollection() {
	taxes1 := []types.Tax{
		{
			Name:              "tax1",
			Rate:              sdk.MustNewDecFromStr("0.5"),
			TaxSourceAddress:  suite.taxSourceAddrs[0].String(),
			CollectionAddress: suite.collectionAddrs[0].String(),
			StartTime:         mustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:           mustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:              "tax2",
			Rate:              sdk.MustNewDecFromStr("0.5"),
			TaxSourceAddress:  suite.taxSourceAddrs[0].String(),
			CollectionAddress: suite.collectionAddrs[1].String(),
			StartTime:         mustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:           mustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:              "tax3",
			Rate:              sdk.MustNewDecFromStr("1.0"),
			TaxSourceAddress:  suite.taxSourceAddrs[1].String(),
			CollectionAddress: suite.collectionAddrs[2].String(),
			StartTime:         mustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:           mustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:              "tax4",
			Rate:              sdk.MustNewDecFromStr("1"),
			TaxSourceAddress:  suite.taxSourceAddrs[2].String(),
			CollectionAddress: suite.collectionAddrs[3].String(),
			StartTime:         mustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:           mustParseRFC3339("0000-01-01T00:00:00Z"),
		},
	}
	taxes3 := []types.Tax{
		{
			Name:              "tax5",
			Rate:              sdk.MustNewDecFromStr("0.5"),
			TaxSourceAddress:  suite.taxSourceAddrs[3].String(),
			CollectionAddress: suite.collectionAddrs[0].String(),
			StartTime:         mustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:           mustParseRFC3339("9999-12-31T00:00:00Z"),
		},
		{
			Name:              "tax6",
			Rate:              sdk.MustNewDecFromStr("0.5"),
			TaxSourceAddress:  suite.taxSourceAddrs[3].String(),
			CollectionAddress: suite.collectionAddrs[1].String(),
			StartTime:         mustParseRFC3339("0000-01-01T00:00:00Z"),
			EndTime:           mustParseRFC3339("9999-12-31T00:00:00Z"),
		},
	}

	for _, tc := range []struct {
		name           string
		taxes          []types.Tax
		epochBlocks    uint32
		accAsserts     []sdk.AccAddress
		balanceAsserts []sdk.Coins
		expectErr      bool
	}{
		{
			"basic taxes case",
			taxes1,
			types.DefaultEpochBlocks,
			[]sdk.AccAddress{
				suite.collectionAddrs[0],
				suite.collectionAddrs[1],
				suite.collectionAddrs[2],
				suite.collectionAddrs[3],
				suite.taxSourceAddrs[0],
				suite.taxSourceAddrs[1],
				suite.taxSourceAddrs[2],
			},
			[]sdk.Coins{
				mustParseCoinsNormalized("500000000denom1,500000000denom2,500000000denom3,500000000stake"),
				mustParseCoinsNormalized("500000000denom1,500000000denom2,500000000denom3,500000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				sdk.Coins{},
				sdk.Coins{},
				sdk.Coins{},
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
			},
			false,
		},
		{
			"only expired tax case",
			[]types.Tax{taxes1[3]},
			types.DefaultEpochBlocks,
			[]sdk.AccAddress{
				suite.collectionAddrs[3],
				suite.taxSourceAddrs[2],
			},
			[]sdk.Coins{
				sdk.Coins{},
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
			},
			false,
		},
		{
			"tax source has small balances case",
			taxes3,
			types.DefaultEpochBlocks,
			[]sdk.AccAddress{
				suite.collectionAddrs[0],
				suite.collectionAddrs[1],
				suite.taxSourceAddrs[3],
			},
			[]sdk.Coins{
				mustParseCoinsNormalized("1denom2,1denom3,500000000stake"),
				mustParseCoinsNormalized("1denom2,1denom3,500000000stake"),
				mustParseCoinsNormalized("1denom1,1denom3"),
			},
			false,
		},
		{
			"none taxes case",
			nil,
			types.DefaultEpochBlocks,
			[]sdk.AccAddress{
				suite.collectionAddrs[0],
				suite.collectionAddrs[1],
				suite.collectionAddrs[2],
				suite.collectionAddrs[3],
				suite.taxSourceAddrs[0],
				suite.taxSourceAddrs[1],
				suite.taxSourceAddrs[2],
				suite.taxSourceAddrs[3],
			},
			[]sdk.Coins{
				sdk.Coins{},
				sdk.Coins{},
				sdk.Coins{},
				sdk.Coins{},
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1denom1,2denom2,3denom3,1000000000stake"),
			},
			false,
		},
		{
			"disabled tax epoch",
			nil,
			0,
			[]sdk.AccAddress{
				suite.collectionAddrs[0],
				suite.collectionAddrs[1],
				suite.collectionAddrs[2],
				suite.collectionAddrs[3],
				suite.taxSourceAddrs[0],
				suite.taxSourceAddrs[1],
				suite.taxSourceAddrs[2],
				suite.taxSourceAddrs[3],
			},
			[]sdk.Coins{
				sdk.Coins{},
				sdk.Coins{},
				sdk.Coins{},
				sdk.Coins{},
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1denom1,2denom2,3denom3,1000000000stake"),
			},
			false,
		},
		{
			"disabled tax epoch with taxes",
			taxes1,
			0,
			[]sdk.AccAddress{
				suite.collectionAddrs[0],
				suite.collectionAddrs[1],
				suite.collectionAddrs[2],
				suite.collectionAddrs[3],
				suite.taxSourceAddrs[0],
				suite.taxSourceAddrs[1],
				suite.taxSourceAddrs[2],
				suite.taxSourceAddrs[3],
			},
			[]sdk.Coins{
				sdk.Coins{},
				sdk.Coins{},
				sdk.Coins{},
				sdk.Coins{},
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1000000000denom1,1000000000denom2,1000000000denom3,1000000000stake"),
				mustParseCoinsNormalized("1denom1,2denom2,3denom3,1000000000stake"),
			},
			false,
		},
	} {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			params := suite.keeper.GetParams(suite.ctx)
			params.Taxes = tc.taxes
			params.EpochBlocks = tc.epochBlocks
			suite.keeper.SetParams(suite.ctx, params)

			err := suite.keeper.TaxCollection(suite.ctx)
			if tc.expectErr {
				suite.Error(err)
			} else {
				suite.NoError(err)

				for i, acc := range tc.accAsserts {
					suite.True(suite.app.BankKeeper.GetAllBalances(suite.ctx, acc).IsEqual(tc.balanceAsserts[i]))
				}
			}
		})
	}
}
