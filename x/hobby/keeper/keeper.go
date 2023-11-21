package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper    types.AccountKeeper
		authzKeeper      types.AuthzKeeper
		bankKeeper       types.BankKeeper
		mintKeeper       types.MintKeeper
		stakingKeeper    types.StakingKeeper
		slashingKeeper   types.SlashingKeeper
		crisisKeeper     types.CrisisKeeper
		evidenceKeeper   types.EvidenceKeeper
		upgradeKeeper    types.UpgradeKeeper
		capabilityKeeper types.CapabilityKeeper
		paramsKeeper     types.ParamsKeeper
		groupKeeper      types.GroupKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	authzKeeper types.AuthzKeeper,
	bankKeeper types.BankKeeper,
	mintKeeper types.MintKeeper,
	stakingKeeper types.StakingKeeper,
	slashingKeeper types.SlashingKeeper,
	crisisKeeper types.CrisisKeeper,
	evidenceKeeper types.EvidenceKeeper,
	upgradeKeeper types.UpgradeKeeper,
	capabilityKeeper types.CapabilityKeeper,
	paramsKeeper types.ParamsKeeper,
	groupKeeper types.GroupKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		accountKeeper:    accountKeeper,
		authzKeeper:      authzKeeper,
		bankKeeper:       bankKeeper,
		mintKeeper:       mintKeeper,
		stakingKeeper:    stakingKeeper,
		slashingKeeper:   slashingKeeper,
		crisisKeeper:     crisisKeeper,
		evidenceKeeper:   evidenceKeeper,
		upgradeKeeper:    upgradeKeeper,
		capabilityKeeper: capabilityKeeper,
		paramsKeeper:     paramsKeeper,
		groupKeeper:      groupKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
