package hobby_test

import (
	"testing"

	keepertest "github.com/hobbyworld-project/hobbychain/testutil/keeper"
	"github.com/hobbyworld-project/hobbychain/testutil/nullify"
	"github.com/hobbyworld-project/hobbychain/x/hobby"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HobbyKeeper(t)
	hobby.InitGenesis(ctx, *k, genesisState)
	got := hobby.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
