package keeper

import (
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
)

var _ types.QueryServer = Keeper{}
