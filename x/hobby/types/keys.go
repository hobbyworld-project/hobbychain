package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	// ModuleName defines the module name
	ModuleName = "hobby"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_hobby"

	// SwapPoolName is the name of the token swap
	SwapPoolName = "swap-pool"
)

const (
	PrivateDataKeyMaxSize   = 128
	PrivateDataValueMaxSize = 1024
)

var (
	// ParamsKey is the key to query module params
	ParamsKey = []byte{0x30}
)

var (
	PrivateDataPrefix      = []byte{0x01}
	GovContractAddrPrefix  = []byte{0x40}
	GenesisNftPrefix       = []byte{0x41}
	GovContractAddrKey     = []byte{0x42}
	CandidatePrefix        = []byte{0x43}
	CandidateVoterPrefix   = []byte{0x44}
	SwapContractAddrPrefix = []byte{0x45}
	SwapContractAddrKey    = []byte{0x46}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// ModuleAddress is the native module address for EVM
var ModuleAddress common.Address

var SwapPoolAddress sdk.AccAddress

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
	SwapPoolAddress = authtypes.NewModuleAddress(SwapPoolName)
}

// GetGenesisNftKey gets a specific token id key
func GetGenesisNftKey(tokenId *big.Int) []byte {
	return append(GenesisNftPrefix, tokenId.Bytes()...)
}

func GetCandidatePrefix(valAddr sdk.ValAddress) []byte {
	return append(CandidateVoterPrefix, valAddr.Bytes()...)
}
