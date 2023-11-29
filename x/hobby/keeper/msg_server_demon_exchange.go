package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
)

func (k msgServer) DemonExchange(goCtx context.Context, msg *types.MsgDemonExchange) (*types.MsgDemonExchangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	var err error

	logger := ctx.Logger()
	params := k.GetParams(ctx)
	ex := params.Exchange

	if len(ex.AllowList) != 0 {
		var ok bool
		for _, allow := range ex.AllowList {
			if allow == msg.Creator {
				ok = true
			}
		}
		if !ok {
			err = fmt.Errorf("demon exchange error: not allowed")
			return nil, err
		}
	}

	var amt sdk.Dec
	amt, err = sdk.NewDecFromStr(msg.Amount)
	if err != nil {
		logger.Error("demon exchange", "error", err.Error(), "creator", msg.Creator)
		return nil, err
	}
	var sender sdk.AccAddress
	sender, err = sdk.AccAddressFromBech32(msg.Creator)
	burnCoin := sdk.NewCoin(ex.FromDenom, amt.TruncateInt())
	burnCoins := sdk.NewCoins(burnCoin)
	// send coins to module and burn
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, burnCoins)
	if err != nil {
		logger.Error("demon exchange", "account send coins error", err.Error(), "sender", msg.Creator, "receiver", types.ModuleName)
		return nil, err
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if err != nil {
		logger.Error("demon exchange", "module burn coins error", err.Error())
		return nil, err
	}

	// mint amount of to denom for creator
	mintedAmt := amt.MulInt64(ex.ExchangeRatio.TruncateInt64())
	mintedCoin := sdk.NewCoin(ex.ToDenom, mintedAmt.TruncateInt())
	mintedCoins := sdk.NewCoins(mintedCoin)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, mintedCoins)
	if err != nil {
		logger.Error("demon exchange", "module mint coins error", err.Error())
		return nil, err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, mintedCoins)
	if err != nil {
		logger.Error("demon exchange", "module send coins error", err.Error())
		return nil, err
	}
	return &types.MsgDemonExchangeResponse{}, nil
}
