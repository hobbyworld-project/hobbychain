package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultGenesisNftVestingEpochs = 1126286 // about 3 months
	DefaultSettleIntervalEpochs    = 12343   // about 1 day
)

var (
	DefaultMasterVestingReward = sdk.MustNewDecFromStr("20000000000000000000000")    //vesting tokens
	DefaultSlaveVestingReward  = sdk.MustNewDecFromStr("10000000000000000000000")    //vesting token
	DefaultCommonVestingReward = sdk.MustNewDecFromStr("4973000000000000000000")     //vesting token
	DefaultMintQuota           = sdk.MustNewDecFromStr("50000000000000000000000000") //mint quota
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	//default uhby -> usby ratio
	ratio, err := sdk.NewDecFromStr("10.0")
	if err != nil {
		panic("exchange ratio invalid")
	}
	return Params{
		Exchange: DenomExchange{
			FromDenom:     "uhby",
			ToDenom:       "usby",
			ExchangeRatio: ratio,
		},
		GovErc721: GovErc721{
			EnableEvm:            true,
			Denom:                "uhby",
			MintQuota:            &DefaultMintQuota,
			MasterVestingReward:  &DefaultMasterVestingReward,
			MasterVestingEpochs:  DefaultGenesisNftVestingEpochs, // default 3month
			SlaveVestingReward:   &DefaultSlaveVestingReward,
			SlaveVestingEpochs:   DefaultGenesisNftVestingEpochs, // default 3month
			CommonVestingReward:  &DefaultCommonVestingReward,
			CommonVestingEpochs:  DefaultGenesisNftVestingEpochs, // default 3month
			SettleIntervalEpochs: DefaultSettleIntervalEpochs,    // settle interval epochs (default 1day)
			AllowDeploy:          false,                          // default close user deploy governance contract
		},
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
