package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PrivateData(goCtx context.Context, req *types.QueryPrivateDataRequest) (*types.QueryPrivateDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if req.Address == "" || req.Key == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request address or key")
	}
	ctx.Logger().Info("query private data request", "height", ctx.BlockHeight(), "address", req.Address, "key", req.Key)
	value := k.GetPrivateData(ctx, req.Address, req.Key)
	return &types.QueryPrivateDataResponse{
		Value: value,
	}, nil
}
