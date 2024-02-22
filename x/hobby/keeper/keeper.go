package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

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
		evmKeeper        types.EVMKeeper
		feeCollectorName string
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
	evmKeeper types.EVMKeeper,
	feeCollectorName string,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		memKey:           memKey,
		paramstore:       ps,
		evmKeeper:        evmKeeper,
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
		feeCollectorName: feeCollectorName,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetPrivateData(ctx sdk.Context, creator, key string) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrivateDataPrefix)
	bz := store.Get(k.GetCreatorKeyPrefix(creator, key))
	return string(bz)
}

func (k Keeper) SetPrivateData(ctx sdk.Context, creator, key, value string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.PrivateDataPrefix)
	store.Set(k.GetCreatorKeyPrefix(creator, key), []byte(value))
}

func (k Keeper) GetCreatorKeyPrefix(creator, key string) []byte {
	return []byte(fmt.Sprintf("%s/%s", creator, key))
}

func (k Keeper) GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k.bankKeeper.GetAllBalances(ctx, addr)
}

// GetModuleAccount returns the module account
func (keeper Keeper) GetModuleAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	return keeper.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
}

func (keeper Keeper) SetModuleAccount(ctx sdk.Context, acc authtypes.ModuleAccountI) {
	keeper.accountKeeper.SetModuleAccount(ctx, acc)
}
