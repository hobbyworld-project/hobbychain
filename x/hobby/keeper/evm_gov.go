package keeper

import (
	"cosmossdk.io/math"
	"encoding/hex"
	"encoding/json"
	"fmt"
	abci "github.com/cometbft/cometbft/abci/types"
	amino "github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v15/utils"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
	"math/big"
	"time"
)

func (k Keeper) GetGenesisNft(ctx sdk.Context, tokenId *sdk.Int) (types.GenesisNft, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GenesisNftPrefix)
	var card types.GenesisNft
	bz := store.Get(tokenId.BigInt().Bytes())
	if len(bz) == 0 {
		return types.GenesisNft{}, false
	}
	k.cdc.MustUnmarshal(bz, &card)
	return card, true
}

func (k Keeper) SetGenesisNft(ctx sdk.Context, card types.GenesisNft) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GenesisNftPrefix)
	bz := k.cdc.MustMarshal(&card)
	store.Set(card.TokenId.BigInt().Bytes(), bz)
}

func (k Keeper) DeleteGenesisNft(ctx sdk.Context, tokenId *sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GenesisNftPrefix)
	store.Delete(tokenId.BigInt().Bytes())
}

func (k Keeper) GetAllGenesisNfts(ctx sdk.Context) (cards []types.GenesisNft) {
	iter := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.GenesisNftPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		bz := iter.Value()
		var card types.GenesisNft
		k.cdc.MustUnmarshal(bz, &card)
		cards = append(cards, card)
	}
	return
}

// GetGovContractAddr get official gov contract address from kv store
func (k Keeper) GetGovContractAddr(ctx sdk.Context) (common.Address, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GovContractAddrPrefix)
	bs := store.Get(types.GovContractAddrKey)
	if len(bs) == 0 {
		return common.Address{}, false
	}
	return common.BytesToAddress(bs), true
}

// SetGovContractAddr set official gov contract address to kv store
func (k Keeper) SetGovContractAddr(ctx sdk.Context, addr common.Address) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GovContractAddrPrefix)
	store.Set(types.GovContractAddrKey, addr.Bytes())
}

// GetSwapContractAddr get official token swap contract address from kv store
func (k Keeper) GetSwapContractAddr(ctx sdk.Context) (common.Address, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SwapContractAddrPrefix)
	bs := store.Get(types.SwapContractAddrKey)
	if len(bs) == 0 {
		return common.Address{}, false
	}
	return common.BytesToAddress(bs), true
}

// SetSwapContractAddr set official token swap contract address to kv store
func (k Keeper) SetSwapContractAddr(ctx sdk.Context, addr common.Address) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SwapContractAddrPrefix)
	store.Set(types.SwapContractAddrKey, addr.Bytes())
}

func (k Keeper) SetCandidate(ctx sdk.Context, candidate types.Candidate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CandidatePrefix)
	data, err := json.Marshal(candidate)
	if err != nil {
		ctx.Logger().Error("marshal candidate", "error", err.Error())
		return
	}
	store.Set(candidate.ValAddr.Bytes(), data)
}

func (k Keeper) GetCandidate(ctx sdk.Context, valAddr sdk.ValAddress) (types.Candidate, bool) {
	var candidate types.Candidate
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CandidatePrefix)
	bz := store.Get(valAddr.Bytes())
	if len(bz) == 0 {
		return types.Candidate{}, false
	}
	_ = json.Unmarshal(bz, &candidate)
	return candidate, true
}

func (k Keeper) DeleteCandidate(ctx sdk.Context, valAddr sdk.ValAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.CandidatePrefix)
	store.Delete(valAddr.Bytes())
}

func (k Keeper) SetCandidateVoter(ctx sdk.Context, valAddr sdk.ValAddress, voter types.Voter) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetCandidatePrefix(valAddr))
	key := voter.Address.Bytes()
	data, err := json.Marshal(voter)
	if err != nil {
		ctx.Logger().Error("marshal voter", "error", err.Error())
		return
	}
	store.Set(key, data)
}

func (k Keeper) GetCandidateVoter(ctx sdk.Context, valAddr sdk.ValAddress, voterAddr sdk.AccAddress) (types.Voter, bool) {
	var voter types.Voter
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetCandidatePrefix(valAddr))
	key := voter.Address.Bytes()
	bz := store.Get(key)
	if len(bz) == 0 {
		ctx.Logger().Error("get candidate not found", "val-addr", valAddr.String(), "voter-addr", voterAddr.String())
		return types.Voter{}, false
	}
	_ = json.Unmarshal(bz, &voter)
	return voter, true
}

func (k Keeper) DeleteCandidateVoter(ctx sdk.Context, valAddr sdk.ValAddress, voterAddr sdk.AccAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetCandidatePrefix(valAddr))
	store.Delete(voterAddr.Bytes())
}

func (k Keeper) GetCandidateVoters(ctx sdk.Context, valAddr sdk.ValAddress) (voters []types.Voter) {
	iter := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.GetCandidatePrefix(valAddr))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var voter types.Voter
		_ = json.Unmarshal(iter.Value(), &voter)
		voters = append(voters, voter)
	}
	return
}

func (k Keeper) GetCandidateVoterCount(ctx sdk.Context, valAddr sdk.ValAddress) (cnt int) {
	iter := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.GetCandidatePrefix(valAddr))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		cnt++
	}
	return
}

func (k Keeper) DeleteCandidateVoters(ctx sdk.Context, valAddr sdk.ValAddress) {
	var keys [][]byte
	iter := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.GetCandidatePrefix(valAddr))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		keys = append(keys, iter.Key())
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.GetCandidatePrefix(valAddr))
	for _, key := range keys {
		store.Delete(key)
	}
}

func (k Keeper) CheckValidatorVotes(ctx sdk.Context, validators []abci.Validator) {
	var logger = ctx.Logger()
	var err error
	var minVotes, heightDiff *big.Int
	start := time.Now()
	minVotes, err = k.MinVote(ctx)
	if err != nil {
		logger.Error("[CheckValidatorVotes] call evm contract failed", "error", err.Error())
		return
	}
	heightDiff, err = k.HeightDiff(ctx)
	if err != nil {
		logger.Error("[CheckValidatorVotes] call evm contract failed", "error", err.Error())
		return
	}
	logger.Info("[CheckValidatorVotes]", "min-votes", minVotes, "height-diff", heightDiff)
	for _, v := range validators {
		validator := k.stakingKeeper.ValidatorByConsAddr(ctx, v.Address)
		valAddr := validator.GetOperator()
		delAddr := sdk.AccAddress(valAddr.Bytes())
		candidate, ok := k.GetCandidate(ctx, valAddr)
		if !ok {
			continue
		}
		voterCnt := k.GetCandidateVoterCount(ctx, validator.GetOperator())
		if int64(voterCnt) < minVotes.Int64() {
			if candidate.CheckPoint == 0 {
				candidate.CheckPoint = ctx.BlockHeight()
				k.SetCandidate(ctx, candidate)
			} else if ctx.BlockHeight()-candidate.CheckPoint > heightDiff.Int64() {
				//delete validator now
				err = k.Unbond(ctx, delAddr, valAddr)
				if err != nil {
					logger.Error("[CheckValidatorVotes] validator remove failed", "val-addr", valAddr.String(), "del-addr", delAddr.String(), "error", err.Error())
					return
				}
				callAddr := common.HexToAddress(hex.EncodeToString(valAddr.Bytes()))
				_, err = k.ContractCall(ctx, types.ContractMethodSetValidatorStatus, callAddr, big.NewInt(types.ValidatorStatusCanceled))
				if err != nil {
					logger.Error("[CheckValidatorVotes] evm update validator status failed", "val-addr", valAddr.String(), "del-addr", delAddr.String(), "error", err.Error())
					return
				}
			}
		} else if candidate.CheckPoint != 0 {
			candidate.CheckPoint = 0
			k.SetCandidate(ctx, candidate)
		}
	}
	end := time.Now()
	logger.Info("[CheckValidatorVotes] finished", "start-time", start, "end-time", end, "elapsed-ms", end.UnixMilli()-start.UnixMilli())
}

func (k Keeper) SettleVoterReward(ctx sdk.Context, validators []abci.Validator) {
	var err error
	var logger = ctx.Logger()
	var valCnt = len(validators)

	if valCnt == 0 {
		return
	}
	start := time.Now()
	var params = k.GetParams(ctx)
	bondDenom := k.stakingKeeper.BondDenom(ctx)
	feeCollector := k.accountKeeper.GetModuleAccount(ctx, k.feeCollectorName)
	balances := k.bankKeeper.GetAllBalances(ctx, feeCollector.GetAddress())
	valCoins := balances.QuoInt(sdk.NewInt(int64(valCnt)))
	logger.Info("[SettleVoterReward] fee collector module", "balances", balances)
	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, k.feeCollectorName, types.ModuleName, balances)
	if err != nil {
		panic(err)
	}

	// TODO: Consider parallelizing later
	for _, v := range validators {
		validator := k.stakingKeeper.ValidatorByConsAddr(ctx, v.Address)
		candidate, ok := k.GetCandidate(ctx, validator.GetOperator())
		if !ok {
			//genesis validator just burn the rewards
			err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, valCoins)
			if err != nil {
				panic(err)
			}
			continue
		}
		if candidate.ActiveHeight == 0 {
			logger.Error("[SettleVoterReward] validator active height must not be 0", "val-addr", validator.GetOperator().String())
			continue
		}
		voters := k.GetCandidateVoters(ctx, validator.GetOperator())
		voterCnt := len(voters)
		if voterCnt == 0 {
			logger.Info("[SettleVoterReward] no voter found", "validator", validator.GetOperator().String(), "burn-amount", valCoins)
			err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, valCoins)
			if err != nil {
				panic(err)
			}
			continue
		}
		var totalWeight = big.NewInt(0)
		for _, voter := range voters {
			totalWeight.SetInt64(voter.Weight.Int64() + totalWeight.Int64())
		}

		avgCoins := valCoins.QuoInt(sdk.NewInt(totalWeight.Int64()))
		logger.Info("[SettleVoterReward] voters", "validator", validator.GetOperator().String(), "count", voterCnt, "avg-coins", avgCoins)
		doSettle := ctx.BlockHeight()-candidate.LastSettleHeight >= params.GovErc721.SettleIntervalEpochs
		for _, c := range avgCoins {
			if !c.IsPositive() {
				continue
			}
			for _, voter := range voters {
				var weight sdk.Int
				weight = sdk.NewIntFromBigInt(voter.Weight)
				c.Amount = c.Amount.Mul(weight)
				c.Amount = c.Amount.Add(voter.ClaimableAmount)
				if bondDenom == c.Denom || doSettle {
					sendCoins := sdk.Coins{c}
					err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, voter.NativeAddress, sendCoins)
					if err != nil {
						logger.Error("[SettleVoterReward] send voter reward", "to", voter.NativeAddress.String(), "error", err.Error(), "coins", sendCoins)
						continue
					}
					if bondDenom != c.Denom {
						candidate.LastSettleHeight = ctx.BlockHeight()
						voter.ClaimableAmount = sdk.NewInt(0)
						voter.RewardedAmount = voter.RewardedAmount.Add(c.Amount)
						k.SetCandidateVoter(ctx, candidate.ValAddr, voter)
						_, err = k.ContractCall(ctx, types.ContractMethodSetRewardAmount, voter.Address, voter.RewardedAmount.BigInt())
						if err != nil {
							logger.Error("[SettleVoterReward] call evm method setRewardAmount", "address", voter.Address.String(), "error", err.Error())
							continue
						}
						logger.Info("[SettleVoterReward] send & update voter", "voter", voter)
					} else {
						logger.Info("[SettleVoterReward] send voter staking coins", "to", voter.NativeAddress.String(), "coins", sendCoins)
					}
				} else {
					voter.ClaimableAmount = voter.ClaimableAmount.Add(c.Amount)
					k.SetCandidateVoter(ctx, candidate.ValAddr, voter)
					logger.Info("[SettleVoterReward] accumulate & update voter", "voter", voter)
				}
			}
		}
		if doSettle {
			k.SetCandidate(ctx, candidate)
		}
	}
	end := time.Now()
	logger.Info("[SettleVoterReward] finished", "start-time", start, "end-time", end, "elapsed-ms", end.UnixMilli()-start.UnixMilli())
}

func (k Keeper) SettleNftVesting(ctx sdk.Context) {
	var err error
	blockHeight := ctx.BlockHeight()
	cards := k.GetAllGenesisNfts(ctx)
	var params = k.GetParams(ctx)
	erc721 := params.GovErc721
	denom := erc721.Denom
	quota := erc721.MintQuota
	logger := ctx.Logger()
	start := time.Now()

	logger.Info("[SettleNftVesting] genesis NFT cards", "count", len(cards))
	for _, card := range cards {
		var claimableAmt math.LegacyDec
		needSettle := (blockHeight - card.LastSettleHeight) >= erc721.SettleIntervalEpochs
		needRemove := blockHeight >= (card.ActiveHeight + card.VestingEpochs)
		claimableAmt = card.LinearAmount.Mul(math.LegacyNewDec(erc721.SettleIntervalEpochs))

		if needSettle || needRemove {
			//mint and send linear amount to card-holder on settle epoch
			coin := sdk.NewCoin(denom, claimableAmt.TruncateInt())
			if coin.Amount.GT(quota.TruncateInt()) {
				logger.Error("[SettleNftVesting] mint quota exceeded")
				return
			}
			err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{coin})
			if err != nil {
				logger.Error("[SettleNftVesting] mint genesis NFT card linear", "error", err.Error(), "token-id", card.TokenId, "card", card)
				return
			}
			var accAddr sdk.AccAddress
			accAddr, err = sdk.AccAddressFromBech32(card.Address)
			if err != nil {
				logger.Error("[SettleNftVesting] malformed account address", "error", err.Error(), "owner", "token-id", card.TokenId, "card", card)
				return
			}
			err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddr, sdk.Coins{coin})
			if err != nil {
				logger.Error("[SettleNftVesting] send coins to account address", "error", err.Error(), "token-id", card.TokenId, "card", card)
				return
			}

			released := card.ReleasedAmount.Add(claimableAmt)
			card.LastSettleHeight = blockHeight
			card.ReleasedAmount = &released
			//callback to evm and update genesis nft card released amount
			owner := common.HexToAddress(card.Owner)
			_, err = k.ContractCall(ctx, types.ContractMethodSetReleasedAmount, owner, released.BigInt())
			if err != nil {
				logger.Error("[SettleNftVesting] update released amount to evm contract failed", "error", err.Error())
				return
			}
			k.SetGenesisNft(ctx, card)
			logger.Info("[SettleNftVesting] settle claimable coins", "token-id", card.TokenId, "token-type", card.TokenType, "released", released, "claimable-coins", claimableAmt)
		}

		if needRemove {
			k.DeleteGenesisNft(ctx, card.TokenId) // all vesting tokens were released, remove it
			logger.Info("[SettleNftVesting] delete genesis NFT", "owner", card.Owner, "address", card.Address, "token-id", card.TokenId, "token-type", card.TokenType, "released", card.ReleasedAmount)
		}
	}
	end := time.Now()
	logger.Info("[SettleNftVesting] finished", "start-time", start, "end-time", end, "elapsed-ms", end.UnixMilli()-start.UnixMilli())
}

func (k Keeper) activeToken(ctx sdk.Context, st types.EventActiveToken) error {
	var err error
	var epochs int64
	var linearAmount sdk.Dec
	var vestingAmount *sdk.Dec
	var releasedAmount sdk.Dec
	var params types.Params

	logger := ctx.Logger()
	params = k.GetParams(ctx)
	gov721 := params.GovErc721
	blockHeight := ctx.BlockHeight()
	tokenId := sdk.NewInt(st.TokenId.Int64())
	tokenType := st.TokenType
	voteWeight := uint32(st.VoteWeight.Uint64())

	switch tokenType {
	case types.GenesisNftTypeMaster:
		epochs = gov721.MasterVestingEpochs
		vestingAmount = gov721.MasterVestingReward
		linearAmount = vestingAmount.QuoInt64(gov721.MasterVestingEpochs)
	case types.GenesisNftTypeSlave:
		epochs = gov721.SlaveVestingEpochs
		vestingAmount = gov721.SlaveVestingReward
		linearAmount = vestingAmount.QuoInt64(gov721.SlaveVestingEpochs)
	case types.GenesisNftTypeCommon:
		epochs = gov721.CommonVestingEpochs
		vestingAmount = gov721.CommonVestingReward
		linearAmount = vestingAmount.QuoInt64(gov721.CommonVestingEpochs)
	default:
		return nil
	}

	var acc sdk.AccAddress
	acc, err = utils.GetAccAddressFromHex(st.Owner.String())
	if err != nil {
		return fmt.Errorf("[activeToken] malformed address %s", st.Owner)
	}
	if tokenId.IsZero() {
		return fmt.Errorf("[activeToken] token id must not be 0")
	}
	if tokenType < types.NftTokenTypeMin || tokenType > types.NftTokenTypeMax {
		return fmt.Errorf("[activeToken] token id %v invalid token type %v", tokenId, tokenType)
	}
	if voteWeight < types.NftTokenWeightMin || voteWeight > types.NftTokenWeightMax {
		return fmt.Errorf("[activeToken] token id %v invalid token weight %v", tokenId, voteWeight)
	}
	if epochs <= 0 || vestingAmount.IsNegative() || linearAmount.IsNegative() {
		return fmt.Errorf("[activeToken] vesting info invalid epochs %v vesting %v linear %v", epochs, vestingAmount, linearAmount)
	}

	var card = types.GenesisNft{
		Owner:      st.Owner.String(),
		Address:    acc.String(),
		TokenId:    &tokenId,
		TokenType:  uint32(tokenType),
		VoteWeight: voteWeight,
	}
	_, ok := k.GetGenesisNft(ctx, card.TokenId)
	if ok {
		logger.Error("[activeToken] token id already exists", "token-id", tokenId)
		return nil // token id already exists
	}
	card.LinearAmount = &linearAmount
	card.VestingEpochs = epochs
	card.ActiveHeight = blockHeight
	card.LastSettleHeight = blockHeight
	card.VestingAmount = vestingAmount
	card.ReleasedAmount = &releasedAmount
	k.SetGenesisNft(ctx, card)
	return nil
}

// validatorUnbond validator unbond all tokens
func (k Keeper) validatorUnbond(ctx sdk.Context, st types.EventUnbond) error {
	delAddr := sdk.AccAddress(st.ValAddr.Bytes())
	valAddr := sdk.ValAddress(st.ValAddr.Bytes())
	return k.Unbond(ctx, delAddr, valAddr)
}

func (k Keeper) Unbond(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	logger := ctx.Logger()
	validator, ok := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !ok {
		// val address is a candidate, just return the delegated tokens
		candidate, exist := k.GetCandidate(ctx, valAddr)
		if !exist {
			return fmt.Errorf("[unbond] candidate address %s not found", valAddr)
		}
		err := k.bankKeeper.UndelegateCoinsFromModuleToAccount(ctx, stakingtypes.NotBondedPoolName, delAddr, sdk.Coins{candidate.StakingCoin})
		if err != nil {
			logger.Error("[unbond] undelegate staking coins failed", "del-addr", delAddr.String(), "val-addr", valAddr.String(), "error", err.Error())
			return err
		}
		logger.Info("[unbond] candidate unbond staking", "del-addr", delAddr.String(), "val-addr", valAddr.String(), "amount", candidate.StakingCoin)
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				stakingtypes.EventTypeCandidateUnbond,
				sdk.NewAttribute(stakingtypes.AttributeKeyValidator, valAddr.String()),
				sdk.NewAttribute(sdk.AttributeKeyAmount, candidate.StakingCoin.String()),
			),
		})
	} else {
		unbondAmount := validator.BondedTokens()
		shares, err := k.stakingKeeper.ValidateUnbondAmount(
			ctx, delAddr, valAddr, unbondAmount,
		)
		if err != nil {
			return fmt.Errorf("[unbond] get unbond amount from validator %s error: %s", valAddr.String(), err.Error())
		}
		var unbondTime time.Time
		unbondTime, err = k.stakingKeeper.Undelegate(ctx, delAddr, valAddr, shares)
		if err != nil {
			return fmt.Errorf("[validatorUnbond] undelegate validator %s error: %s", valAddr, err.Error())
		}
		logger.Info("[unbond] undelegate ok", "validator", valAddr, "unboned-time", unbondTime.Format(time.RFC3339))
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				stakingtypes.EventTypeUnbond,
				sdk.NewAttribute(stakingtypes.AttributeKeyValidator, valAddr.String()),
				sdk.NewAttribute(sdk.AttributeKeyAmount, unbondAmount.String()),
				sdk.NewAttribute(stakingtypes.AttributeKeyCompletionTime, unbondTime.Format(time.RFC3339)),
			),
		})
	}

	k.DeleteCandidate(ctx, valAddr)
	k.DeleteCandidateVoters(ctx, valAddr)
	return nil
}

// userVote voter delegate to a validator
func (k Keeper) userVote(ctx sdk.Context, st types.EventVote) error {
	valAddr := sdk.ValAddress(st.ValAddr.Bytes())
	candidate, ok := k.GetCandidate(ctx, valAddr)
	if !ok {
		return fmt.Errorf("[userVote] candidate %s not found", st.ValAddr)
	}
	var err error
	voterAddr := sdk.AccAddress(st.From.Bytes())
	if err != nil {
		return fmt.Errorf("[userVote] malformed from address %s", st.From)
	}
	if st.Weight.Int64() == 0 {
		return fmt.Errorf("[userVote] voter %s weight %v is invalid", st.From, st.Weight)
	}
	k.SetCandidateVoter(ctx, candidate.ValAddr, types.Voter{
		Address:       st.From,
		NativeAddress: voterAddr,
		Weight:        st.Weight,
	})
	return nil
}

// userUnvote voter undelegate  from a validator
func (k Keeper) userUnvote(ctx sdk.Context, st types.EventUnvote) error {
	voterAddr := sdk.AccAddress(st.From.Bytes())
	valAddr := sdk.ValAddress(st.ValAddr.Bytes())
	candidate, ok := k.GetCandidate(ctx, valAddr)
	if !ok {
		return fmt.Errorf("[userUnvote] candidate %s not found", st.ValAddr)
	}
	k.DeleteCandidateVoter(ctx, candidate.ValAddr, voterAddr)
	return nil
}

// createCandidate create candidate
func (k Keeper) createCandidate(ctx sdk.Context, st types.EventCreateCandidate) error {
	delAddr := sdk.AccAddress(st.ValAddr.Bytes())
	valAddr := sdk.ValAddress(st.ValAddr.Bytes())
	candidate, ok := k.GetCandidate(ctx, valAddr)
	if !ok {
		candidate = types.Candidate{
			AccAddr:     delAddr,
			ValAddr:     valAddr,
			Status:      0,
			StakingCoin: sdk.Coin{},
		}
	}
	candidate.StakingCoin = sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), sdk.NewIntFromBigInt(st.Amount))
	candidate.Status = types.ValidatorStatusCandidate
	k.DeleteCandidateVoters(ctx, candidate.ValAddr)
	k.SetCandidate(ctx, candidate)
	return nil
}

func (k Keeper) candidateStake(ctx sdk.Context, st types.EventStake) error {
	valAddr := sdk.ValAddress(st.ValAddr.Bytes())
	candidate, ok := k.GetCandidate(ctx, valAddr)
	if !ok {
		return fmt.Errorf("evm contract staking: no candidate found, val address: %v", valAddr.String())
	}
	amount := sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), sdk.NewIntFromBigInt(st.Amount))
	if !amount.Equal(candidate.StakingCoin) {
		return fmt.Errorf("evm contract staking: stake amount not match")
	}
	candidate.Status = types.ValidatorStatusVoting
	var pk cryptotypes.PubKey
	codec := amino.NewProtoCodec(codectypes.NewInterfaceRegistry())
	if err := codec.UnmarshalInterfaceJSON([]byte(st.PubKey), &pk); err != nil {
		return fmt.Errorf("evm contract staking: unmarshal interface JSON error: %v", err.Error())
	}

	msg, err := stakingtypes.NewMsgCreateValidator(
		valAddr, pk, amount, stakingtypes.Description{}, stakingtypes.CommissionRates{}, sdk.Int{},
	)
	if err != nil {
		return fmt.Errorf("evm contract staking: new create validator msg error: %v", err.Error())
	}
	_, err = k.stakingKeeper.CreateEvmStaking(ctx, msg)
	if err != nil {
		return fmt.Errorf("evm contract staking: create evm candidate staking error: %v", err.Error())
	}
	return nil
}

// voteFinish vote finished and candidate will become a real validator
func (k Keeper) voteFinish(ctx sdk.Context, st types.EventVoteFinished) error {
	valAddr := sdk.ValAddress(st.ValAddr.Bytes())
	candidate, ok := k.GetCandidate(ctx, valAddr)
	if !ok {
		return fmt.Errorf("[voteFinish] candidate %s not found", st.ValAddr)
	}
	candidate.ActiveHeight = ctx.BlockHeight()
	candidate.LastSettleHeight = ctx.BlockHeight()
	candidate.Status = types.ValidatorStatusOK
	if _, err := k.stakingKeeper.CreateEvmValidator(ctx, valAddr); err != nil {
		return fmt.Errorf("[voteFinish] create validator failed, error: %s", err.Error())
	}
	k.SetCandidate(ctx, candidate)
	return nil
}

// swapDeposit deposit the staking token to swap contract
func (k Keeper) swapDeposit(ctx sdk.Context, st types.EventSwapDeposit) error {
	logger := ctx.Logger()
	addr := sdk.AccAddress(st.Dst.Bytes())
	if st.Wad == nil {
		return fmt.Errorf("swap deposit: from [%s] amount illegal", st.Dst.String())
	}
	denom := k.stakingKeeper.BondDenom(ctx)
	n, ok := sdk.NewIntFromString(st.Wad.String())
	if !ok {
		return fmt.Errorf("swap deposit: from [%s] amount illegal", st.Dst.String())
	}
	amount := sdk.NewCoin(denom, n)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.SwapPoolName, sdk.NewCoins(amount))
	if err != nil {
		return fmt.Errorf("swap deposit: from [%s/%s] to [%s/%s] amount [%v] send error [%s]", st.Dst, addr, types.SwapPoolName, types.SwapPoolAddress, amount, err)
	}
	logger.Info("swap deposit", "eth-addr", st.Dst.String(), "hobby-addr", addr.String(), "to", types.SwapPoolAddress.String(), "amount", amount.String(), "end")
	return nil
}

// swapWithdraw withdraw the staking token from swap contract
func (k Keeper) swapWithdraw(ctx sdk.Context, st types.EventSwapWithdraw) error {
	logger := ctx.Logger()
	addr := sdk.AccAddress(st.Src.Bytes())
	if st.Wad == nil {
		return fmt.Errorf("swap withdraw: from [%s] amount illegal", st.Src.String())
	}
	denom := k.stakingKeeper.BondDenom(ctx)
	n, ok := sdk.NewIntFromString(st.Wad.String())
	if !ok {
		return fmt.Errorf("swap withdraw: from [%s] amount illegal", st.Src.String())
	}
	amount := sdk.NewCoin(denom, n)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.SwapPoolName, addr, sdk.NewCoins(amount))
	if err != nil {
		return fmt.Errorf("swap withdraw: from [%s/%s] to [%s/%s] amount [%v] send error [%s]", types.SwapPoolName, types.SwapPoolAddress, st.Src, addr, amount, err)
	}
	logger.Info("swap withdraw ", "eth-addr", st.Src.String(), "hobby-addr", addr.String(), "from", types.SwapPoolAddress.String(), "amount", amount.String())
	return nil
}

// voteFinish vote finished and candidate will become a real validator
func (k Keeper) deployContract(ctx sdk.Context, st types.EventDeploy, params types.Params) error {
	if params.GovErc721.AllowDeploy {
		k.SetGovContractAddr(ctx, st.Addr)
	}
	return nil
}

func (keeper Keeper) GovEventHanlder(ctx sdk.Context, e *sdk.GovEvent) error {
	switch e.Type {
	case sdk.GovEventCheckValidatorStatus:
		msg := e.Data.(*stakingtypes.MsgCreateValidator)
		return keeper.handleCheckValidatorStatus(ctx, msg)
	case sdk.GovEventSetValidatorStatus:
		msg := e.Data.(*stakingtypes.MsgCreateValidator)
		return keeper.handleSetValidatorStatus(ctx, msg)
	}
	return nil
}

func (keeper Keeper) handleCheckValidatorStatus(ctx sdk.Context, msg *stakingtypes.MsgCreateValidator) (err error) {
	var valAddr sdk.ValAddress
	valAddr, err = sdk.ValAddressFromBech32(msg.ValidatorAddress)
	candidate, ok := keeper.GetCandidate(ctx, valAddr)
	if !ok {
		return fmt.Errorf("candidate %s not found", valAddr)
	}
	if candidate.Status != types.ValidatorStatusCandidate {
		return fmt.Errorf("candidate %s has a invalid status %v", valAddr, candidate.Status)
	}
	if !candidate.StakingCoin.Equal(msg.Value) {
		return fmt.Errorf("candidate %s has a invalid staking amount, expect %v got %v", valAddr, candidate.StakingCoin, msg.Value)
	}
	return nil
}

func (keeper Keeper) handleSetValidatorStatus(ctx sdk.Context, msg *stakingtypes.MsgCreateValidator) error {
	status := big.NewInt(types.ValidatorStatusVoting)
	valAddr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	strHex := hex.EncodeToString(valAddr.Bytes())
	callAddr := common.HexToAddress(strHex)
	_, err = keeper.ContractCall(ctx, types.ContractMethodSetValidatorStatus, callAddr, uint8(status.Int64()))
	if err != nil {
		return fmt.Errorf("call evm method %s validator %s (%s) error %s", types.ContractMethodSetValidatorStatus, msg.ValidatorAddress, callAddr.String(), err.Error())
	}
	return nil
}
