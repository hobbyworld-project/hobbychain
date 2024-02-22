package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/hobby module sentinel errors
var (
	// evm errors
	ErrERC721Disabled        = errorsmod.Register(ModuleName, 17, "erc20 module is disabled")
	ErrInternalTokenPair     = errorsmod.Register(ModuleName, 18, "internal ethereum token mapping error")
	ErrContractNotFound      = errorsmod.Register(ModuleName, 19, "contract not found")
	ErrContractAlreadyExists = errorsmod.Register(ModuleName, 20, "contract already exists")
	ErrUndefinedOwner        = errorsmod.Register(ModuleName, 21, "undefined owner of contract pair")
	ErrBalanceInvariance     = errorsmod.Register(ModuleName, 22, "post transfer balance invariant failed")
	ErrUnexpectedEvent       = errorsmod.Register(ModuleName, 23, "unexpected event")
	ErrABIPack               = errorsmod.Register(ModuleName, 24, "contract ABI pack failed")
	ErrABIUnpack             = errorsmod.Register(ModuleName, 25, "contract ABI unpack failed")
	ErrEVMDenom              = errorsmod.Register(ModuleName, 26, "EVM denomination registration")
	ErrEVMCall               = errorsmod.Register(ModuleName, 27, "EVM call unexpected error")
	ErrAccessDenied          = errorsmod.Register(ModuleName, 28, "access denied")
)
