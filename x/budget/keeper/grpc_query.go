package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/budget/x/budget/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Params queries the parameters of the budget module.
func (k Querier) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)
	return &types.QueryParamsResponse{Params: params}, nil
}

// Budgets queries all budgets.
func (k Querier) Budgets(c context.Context, req *types.QueryBudgetsRequest) (*types.QueryBudgetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.SourceAddress != "" {
		if _, err := sdk.AccAddressFromBech32(req.SourceAddress); err != nil {
			return nil, err
		}
	}

	if req.DestinationAddress != "" {
		if _, err := sdk.AccAddressFromBech32(req.DestinationAddress); err != nil {
			return nil, err
		}
	}

	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)

	var budgets []types.BudgetResponse
	for _, b := range params.Budgets {
		if req.Name != "" && b.Name != req.Name ||
			req.SourceAddress != "" && b.SourceAddress != req.SourceAddress ||
			req.DestinationAddress != "" && b.DestinationAddress != req.DestinationAddress {
			continue
		}

		collectedCoins := k.GetTotalCollectedCoins(ctx, b.Name)
		budgets = append(budgets, types.BudgetResponse{
			Budget:              b,
			TotalCollectedCoins: collectedCoins,
		})
	}

	return &types.QueryBudgetsResponse{Budgets: budgets}, nil
}
