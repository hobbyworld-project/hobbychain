package contracts

import (
	_ "embed"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/evmos/evmos/v15/x/evm/types"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
)

var (
	//go:embed compiled_contracts/WHBY.json
	WHBYJSON []byte //nolint: golint

	WHBYContract evmtypes.CompiledContract //evm precompiled contract

	WHBYAddress common.Address //contract module address
)

func init() {
	WHBYAddress = types.ModuleAddress
	err := json.Unmarshal(WHBYJSON, &WHBYContract)
	if err != nil {
		panic(err)
	}

	if len(WHBYContract.Bin) == 0 {
		panic("load swap contract abi/bin failed")
	}
}
