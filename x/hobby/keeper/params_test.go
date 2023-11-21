package keeper_test

import (
	"testing"

	testkeeper "github.com/hobbyworld-project/hobbychain/testutil/keeper"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HobbyKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
