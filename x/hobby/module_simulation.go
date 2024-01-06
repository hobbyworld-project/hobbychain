package hobby

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/hobbyworld-project/hobbychain/testutil/sample"
	hobbysimulation "github.com/hobbyworld-project/hobbychain/x/hobby/simulation"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = hobbysimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgDemonExchange = "op_weight_msg_demon_exchange"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDemonExchange int = 100

	opWeightMsgPrivateData = "op_weight_msg_private_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPrivateData int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	hobbyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&hobbyGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgDemonExchange int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDemonExchange, &weightMsgDemonExchange, nil,
		func(_ *rand.Rand) {
			weightMsgDemonExchange = defaultWeightMsgDemonExchange
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDemonExchange,
		hobbysimulation.SimulateMsgDemonExchange(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPrivateData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPrivateData, &weightMsgPrivateData, nil,
		func(_ *rand.Rand) {
			weightMsgPrivateData = defaultWeightMsgPrivateData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPrivateData,
		hobbysimulation.SimulateMsgPrivateData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgDemonExchange,
			defaultWeightMsgDemonExchange,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				hobbysimulation.SimulateMsgDemonExchange(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
	opWeightMsgPrivateData,
	defaultWeightMsgPrivateData,
	func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
		hobbysimulation.SimulateMsgPrivateData(am.accountKeeper, am.bankKeeper, am.keeper)
		return nil
	},
),
// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
