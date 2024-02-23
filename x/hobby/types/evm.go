package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	GenesisNftTypeDefault = 0
	GenesisNftTypeMaster  = 1
	GenesisNftTypeSlave   = 2
	GenesisNftTypeCommon  = 3
	GenesisNftTypeHuman   = 4
)

const (
	NftTokenTypeMin   = 1
	NftTokenTypeMax   = 4
	NftTokenWeightMin = 1
	NftTokenWeightMax = 30
)

const (
	ValidatorStatusInvalid   = 0
	ValidatorStatusCandidate = 1
	ValidatorStatusVoting    = 2
	ValidatorStatusOK        = 3
	ValidatorStatusExpired   = 4
	ValidatorStatusCanceled  = 5
)

const (
	ContractMethodValidatorStatus    = "validatorStatus"
	ContractMethodSetValidatorStatus = "setValidatorStatus"
	ContractMethodSetReleasedAmount  = "setReleasedAmount"
	ContractMethodSetRewardAmount    = "setRewardAmount"
	ContractMethodMaxVote            = "maxVote"
	ContractMethodMinVote            = "minVote"
	ContractMethodHeightDiff         = "heightDiff"
)

const (
	// governance events
	ContractEventNameDeploy          = "Deploy"
	ContractEventNameTransfer        = "Transfer"
	ContractEventNameVoteFinish      = "VoteFinish"
	ContractEventNameVote            = "Vote"
	ContractEventNameUnvote          = "Unvote"
	ContractEventNameUnbond          = "Unbond"
	ContractEventNameActiveToken     = "ActiveToken"
	ContractEventNameCreateCandidate = "CreateCandidate"
	ContractEventNameDeposit         = "Deposit"
	ContractEventNameWithdrawal      = "Withdrawal"
	ContractEventNameStake           = "Stake"
)

type Voter struct {
	Address         common.Address
	NativeAddress   sdk.AccAddress
	Weight          *big.Int
	ClaimableAmount sdk.Int
	RewardedAmount  sdk.Int
}

type Candidate struct {
	AccAddr          sdk.AccAddress
	ValAddr          sdk.ValAddress
	Status           int32
	StakingCoin      sdk.Coin
	ActiveHeight     int64
	LastSettleHeight int64
	CheckPoint       int64
}

type EventDeploy struct {
	Addr common.Address
}

type EventTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
}

type EventVoteFinished struct {
	ValAddr common.Address
}

type EventVote struct {
	ValAddr  common.Address
	From     common.Address
	Weight   *big.Int
	TokenIds []*big.Int
}

type EventUnvote struct {
	ValAddr  common.Address
	From     common.Address
	TokenIds []*big.Int
}

type EventUnbond struct {
	ValAddr common.Address
}

type EventActiveToken struct {
	TokenId    *big.Int
	Owner      common.Address
	TokenType  uint8
	VoteWeight *big.Int
	VestAmount *big.Int
}

type EventCreateCandidate struct {
	ValAddr common.Address
	Amount  *big.Int
}

type EventSwapDeposit struct {
	Dst common.Address
	Wad *big.Int
}

type EventSwapWithdraw struct {
	Src common.Address
	Wad *big.Int
}

type EventStake struct {
	ValAddr common.Address
	PubKey  string
	Amount  *big.Int
}
