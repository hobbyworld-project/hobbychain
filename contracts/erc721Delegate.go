package contracts

import (
	_ "embed"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/evmos/v15/x/evm/types"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
)

var (
	//go:embed compiled_contracts/ERC721Delegate.json
	ERC721DelegateJSON []byte //nolint: golint

	ERC721DelegateContract evmtypes.CompiledContract //evm precompiled contract

	ERC721DelegateAddress common.Address //contract module address
)

func init() {
	ERC721DelegateAddress = types.ModuleAddress
	err := json.Unmarshal(ERC721DelegateJSON, &ERC721DelegateContract)
	if err != nil {
		panic(err)
	}

	if len(ERC721DelegateContract.Bin) == 0 {
		panic("load erc721 gov contract failed")
	}
}
