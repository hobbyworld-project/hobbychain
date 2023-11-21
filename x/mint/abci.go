package mint

import (
	"github.com/hobbyworld-project/hobbychain/x/mint/keeper"
	"github.com/hobbyworld-project/hobbychain/x/mint/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, ic types.InflationCalculationFn) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)
	enableReduction := params.Reduction.Enable
	if !enableReduction {
		// recalculate inflation rate
		totalStakingSupply := k.StakingTokenSupply(ctx)
		bondedRatio := k.BondedRatio(ctx)
		minter.Inflation = ic(ctx, minter, params, bondedRatio)
		minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply)
		k.SetMinter(ctx, minter)

		// mint coins, update supply
		mintedCoin := minter.BlockProvision(params)
		mintedCoins := sdk.NewCoins(mintedCoin)

		err := k.MintCoins(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}
		logger := k.Logger(ctx)

		logger.Info("[MINT inflation] minted coins", "amount", mintedCoin.String())
		// send the minted coins to the fee collector account
		err = k.AddCollectedFees(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}

		if mintedCoin.Amount.IsInt64() {
			defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeMint,
				sdk.NewAttribute(types.AttributeKeyBondedRatio, bondedRatio.String()),
				sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
				sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
				sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
			),
		)
	} else {
		k.SetMinter(ctx, minter)
		// mint coins, update minted tokens
		mintedCoin := minter.BlockProvisionReduction(ctx.BlockHeight(), params)

		mintedCoins := sdk.NewCoins(mintedCoin)

		err := k.MintCoins(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}
		logger := k.Logger(ctx)

		// send the minted coins to the fee collector account
		err = k.AddCollectedFees(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}
		logger.Info("[MINT reduction] minted coins", "amount", mintedCoin.String())
		if mintedCoin.Amount.IsInt64() {
			defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
		}
	}
}
