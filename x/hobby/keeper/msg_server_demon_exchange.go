package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
)

func (k msgServer) DemonExchange(goCtx context.Context, msg *types.MsgDemonExchange) (*types.MsgDemonExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	log := ctx.Logger()
	log.Info("demon exchange: creator %s amount %s", msg.Creator, msg.Amount)
	return &types.MsgDemonExchangeResponse{}, nil
}
