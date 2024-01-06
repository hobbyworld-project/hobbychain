package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
)

func (k msgServer) PrivateData(goCtx context.Context, msg *types.MsgPrivateData) (*types.MsgPrivateDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	if len(msg.Value) > types.PrivateDataKeyMaxSize {
		return nil, fmt.Errorf("private data key size exceeded, must less than equal %v", types.PrivateDataKeyMaxSize)
	}
	if len(msg.Value) > types.PrivateDataValueMaxSize {
		return nil, fmt.Errorf("private data value size exceeded, must less than equal %v", types.PrivateDataValueMaxSize)
	}
	k.SetPrivateData(ctx, msg.Creator, msg.Key, msg.Value)
	return &types.MsgPrivateDataResponse{}, nil
}
