package hobby

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hobbyworld-project/hobbychain/x/hobby/keeper"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	// if account has zero balance it probably means it's not set, so we set it
	moduleAcc := k.GetModuleAccount(ctx)
	balance := k.GetAllBalances(ctx, moduleAcc.GetAddress())
	if balance.IsZero() {
		k.SetModuleAccount(ctx, moduleAcc)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
