package types

import (
	"context"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v15/x/evm/statedb"
	evmtypes "github.com/evmos/evmos/v15/x/evm/types"
	"time"
)

type AuthzKeeper interface {
	// Methods imported from authz should be defined here
}

type MintKeeper interface {
	// Methods imported from mint should be defined here
}

type StakingKeeper interface {
	// Methods imported from staking should be defined here
	// iterate through bonded validators by operator address, execute func for each validator
	IterateBondedValidatorsByPower(
		sdk.Context, func(index int64, validator stakingtypes.ValidatorI) (stop bool),
	)

	TotalBondedTokens(sdk.Context) math.Int // total bonded tokens within the validator set
	IterateDelegations(
		ctx sdk.Context, delegator sdk.AccAddress,
		fn func(index int64, delegation stakingtypes.DelegationI) (stop bool),
	)
	SetEvmCallback(cb sdk.GovEventCallback)
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	Undelegate(
		ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, sharesAmount sdk.Dec,
	) (time.Time, error)
	ValidateUnbondAmount(
		ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, amt math.Int,
	) (shares sdk.Dec, err error)
	BondDenom(ctx sdk.Context) string
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI // get a particular validator by consensus address
	CreateEvmValidator(ctx sdk.Context, valAddr sdk.ValAddress) (*stakingtypes.MsgCreateValidatorResponse, error)
}

type SlashingKeeper interface {
	// Methods imported from slashing should be defined here
}

type CrisisKeeper interface {
	// Methods imported from crisis should be defined here
}

type EvidenceKeeper interface {
	// Methods imported from evidence should be defined here
}

type UpgradeKeeper interface {
	// Methods imported from upgrade should be defined here
}

type CapabilityKeeper interface {
	// Methods imported from capability should be defined here
}

type ParamsKeeper interface {
	// Methods imported from params should be defined here
}

type GroupKeeper interface {
	// Methods imported from group should be defined here
}

type DistributionKeeper interface {
	IsBurnValidator(ctx sdk.Context, validator stakingtypes.ValidatorI) bool
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) types.ModuleAccountI

	// TODO remove with genesis 2-phases refactor https://github.com/cosmos/cosmos-sdk/issues/2862
	SetModuleAccount(sdk.Context, types.ModuleAccountI)
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)
}

type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// EVMKeeper defines the expected EVM keeper interface used on erc20
type EVMKeeper interface {
	GetParams(ctx sdk.Context) evmtypes.Params
	GetAccountWithoutBalance(ctx sdk.Context, addr common.Address) *statedb.Account
	EstimateGas(c context.Context, req *evmtypes.EthCallRequest) (*evmtypes.EstimateGasResponse, error)
	ApplyMessage(ctx sdk.Context, msg core.Message, tracer vm.EVMLogger, commit bool) (*evmtypes.MsgEthereumTxResponse, error)
}
