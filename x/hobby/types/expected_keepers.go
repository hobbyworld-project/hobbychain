package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AuthzKeeper interface {
	// Methods imported from authz should be defined here
}

type MintKeeper interface {
	// Methods imported from mint should be defined here
}

type StakingKeeper interface {
	// Methods imported from staking should be defined here
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

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
