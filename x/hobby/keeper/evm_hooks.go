// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package keeper

import (
	"bytes"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmcommon "github.com/evmos/evmos/v15/precompiles/common"
	evmtypes "github.com/evmos/evmos/v15/x/evm/types"
	"github.com/hobbyworld-project/hobbychain/contracts"
	types "github.com/hobbyworld-project/hobbychain/x/hobby/types"
)

var _ evmtypes.EvmHooks = Hooks{}

// Hooks wrapper struct for erc20 keeper
type Hooks struct {
	k Keeper
}

// Return the wrapper struct
func (k Keeper) EvmHooks() Hooks {
	return Hooks{k}
}

// PostTxProcessing is a wrapper for calling the EVM PostTxProcessing hook on
// the module keeper
func (h Hooks) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	return h.k.PostTxProcessing(ctx, msg, receipt)
}

// PostTxProcessing implements EvmHooks.PostTxProcessing.
func (k Keeper) PostTxProcessing(
	ctx sdk.Context,
	msg core.Message,
	receipt *ethtypes.Receipt,
) error {
	logger := ctx.Logger()
	params := k.GetParams(ctx)
	if !params.GovErc721.EnableEvm {
		// no error is returned to avoid reverting the tx and allow for other post
		// processing txs to pass and
		return nil
	}
	logger.Info("[PostTxProcessing]", "gov-module-address", types.ModuleAddress.String())
	if len(msg.Data()) == 0 {
		//this is a native token transfer msg
		return nil
	}
	govAddr, ok := k.GetGovContractAddr(ctx)
	if !ok {
		logger.Error("no gov contract address found for evm tx event processing")
		return nil
	}
	swapAddr, ok := k.GetSwapContractAddr(ctx)
	if !ok {
		logger.Error("no swap contract address found for evm tx event processing")
		return nil
	}

	swapABI := contracts.WHBYContract.ABI
	erc721 := contracts.ERC721DelegateContract.ABI

	for _, log := range receipt.Logs {

		eventID := log.Topics[0]
		event, err := erc721.EventByID(eventID)
		if err != nil {
			event, err = swapABI.EventByID(eventID)
			if err != nil {
				logger.Info("[EvmHook] event not found", "event-id", eventID.String(), "error", err.Error())
				return nil
			}
		}

		switch event.Name {
		case types.ContractEventNameDeploy: //contract test only
			err = k.handleContractEventDeploy(ctx, log, event, erc721, params)
		case types.ContractEventNameTransfer:
			err = k.handleContractEventTransfer(ctx, log, event, erc721, govAddr)
		case types.ContractEventNameCreateCandidate:
			err = k.handleContractEventCreateCandidate(ctx, log, event, erc721, govAddr)
		case types.ContractEventNameVoteFinish:
			err = k.handleContractEventVoteFinish(ctx, log, event, erc721, govAddr)
		case types.ContractEventNameVote:
			err = k.handleContractEventVote(ctx, log, event, erc721, govAddr)
		case types.ContractEventNameUnvote:
			err = k.handleContractEventUnvote(ctx, log, event, erc721, govAddr)
		case types.ContractEventNameUnbond:
			err = k.handleContractEventUnbond(ctx, log, event, erc721, govAddr)
		case types.ContractEventNameActiveToken:
			err = k.handleContractEventActiveToken(ctx, log, event, erc721, govAddr)
		case types.ContractEventNameDeposit:
			err = k.handleContractEventDeposit(ctx, log, event, swapABI, swapAddr)
		case types.ContractEventNameWithdrawal:
			err = k.handleContractEventWithdraw(ctx, log, event, swapABI, swapAddr)
		default:
			logger.Info("[EvmHook] can not handle event", "name", event.Name, "receipt-addr", log.Address.String())
		}
		if err != nil {
			logger.Error("[EvmHook] handle event failed", "event", event.Name, "error", err.Error())
			return err
		}
	}
	return nil
}

func (k Keeper) handleContractEventDeploy(ctx sdk.Context, log *ethtypes.Log, event *abi.Event, contactABI abi.ABI, params types.Params) error {
	logger := ctx.Logger()
	contractAddr := log.Address
	var st types.EventDeploy

	err := evmcommon.UnpackLog(contactABI, &st, event.Name, *log)
	if err != nil {
		return fmt.Errorf("[EvmHook] contract %s event %s unpack log error: %s", contractAddr, event.Name, err.Error())
	}
	logger.Info("[EvmHook]", "contract-address", contractAddr, "event", event.Name, "unpack", st)
	return k.deployContract(ctx, st, params)
}

func (k Keeper) handleContractEventCreateCandidate(ctx sdk.Context, log *ethtypes.Log, event *abi.Event, contactABI abi.ABI, govAddr common.Address) error {
	logger := ctx.Logger()
	contractAddr := log.Address
	if !bytes.Equal(contractAddr.Bytes(), govAddr.Bytes()) {
		logger.Debug("[EvmHook] contract address is not equal to gov address")
		return nil
	}

	var st types.EventCreateCandidate

	err := evmcommon.UnpackLog(contactABI, &st, event.Name, *log)
	if err != nil {
		return fmt.Errorf("[EvmHook] contract %s event %s unpack log error: %s", contractAddr, event.Name, err.Error())
	}
	logger.Info("[EvmHook]", "contract-address", contractAddr, "event", event.Name, "unpack", st)
	return k.createCandidate(ctx, st)
}

func (k Keeper) handleContractEventVoteFinish(ctx sdk.Context, log *ethtypes.Log, event *abi.Event, contactABI abi.ABI, govAddr common.Address) error {
	logger := ctx.Logger()
	contractAddr := log.Address
	if !bytes.Equal(contractAddr.Bytes(), govAddr.Bytes()) {
		logger.Debug("[EvmHook] contract address is not equal to gov address")
		return nil
	}

	var st types.EventVoteFinished

	err := evmcommon.UnpackLog(contactABI, &st, event.Name, *log)
	if err != nil {
		return fmt.Errorf("[EvmHook] contract %s event %s unpack log error: %s", contractAddr, event.Name, err.Error())
	}
	logger.Info("[EvmHook]", "contract-address", contractAddr, "event", event.Name, "unpack", st)
	return k.voteFinish(ctx, st)
}

func (k Keeper) handleContractEventVote(ctx sdk.Context, log *ethtypes.Log, event *abi.Event, contactABI abi.ABI, govAddr common.Address) error {
	logger := ctx.Logger()
	contractAddr := log.Address
	if !bytes.Equal(contractAddr.Bytes(), govAddr.Bytes()) {
		logger.Debug("[EvmHook] contract address is not equal to gov address")
		return nil
	}
	var st types.EventVote
	err := evmcommon.UnpackLog(contactABI, &st, event.Name, *log)
	if err != nil {
		return fmt.Errorf("[EvmHook] contract %s event %s unpack log error: %s", contractAddr, event.Name, err.Error())
	}
	logger.Info("[EvmHook]", "contract-address", contractAddr, "event", event.Name, "unpack", st)
	return k.userVote(ctx, st)
}

func (k Keeper) handleContractEventUnvote(ctx sdk.Context, log *ethtypes.Log, event *abi.Event, contactABI abi.ABI, govAddr common.Address) error {
	logger := ctx.Logger()
	contractAddr := log.Address
	if !bytes.Equal(contractAddr.Bytes(), govAddr.Bytes()) {
		logger.Debug("[EvmHook] contract address is not equal to gov address")
		return nil
	}
	var st types.EventUnvote

	err := evmcommon.UnpackLog(contactABI, &st, event.Name, *log)
	if err != nil {
		return fmt.Errorf("[EvmHook] contract %s event %s unpack log error: %s", contractAddr, event.Name, err.Error())
	}
	logger.Info("[EvmHook]", "contract-address", contractAddr, "event", event.Name, "unpack", st)
	return k.userUnvote(ctx, st)
}

func (k Keeper) handleContractEventUnbond(ctx sdk.Context, log *ethtypes.Log, event *abi.Event, contactABI abi.ABI, govAddr common.Address) error {
	logger := ctx.Logger()
	contractAddr := log.Address
	if !bytes.Equal(contractAddr.Bytes(), govAddr.Bytes()) {
		logger.Debug("[EvmHook] contract address is not equal to gov address")
		return nil
	}
	var st types.EventUnbond

	err := evmcommon.UnpackLog(contactABI, &st, event.Name, *log)
	if err != nil {
		return fmt.Errorf("[EvmHook] contract %s event %s unpack log error: %s", contractAddr, event.Name, err.Error())
	}
	logger.Info("[EvmHook]", "contract-address", contractAddr, "event", event.Name, "unpack", st)
	return k.validatorUnbond(ctx, st)
}

func (k Keeper) handleContractEventActiveToken(ctx sdk.Context, log *ethtypes.Log, event *abi.Event, contactABI abi.ABI, govAddr common.Address) error {
	logger := ctx.Logger()
	contractAddr := log.Address
	if !bytes.Equal(contractAddr.Bytes(), govAddr.Bytes()) {
		logger.Debug("[EvmHook] contract address is not equal to gov address")
		return nil
	}
	var st types.EventActiveToken

	err := evmcommon.UnpackLog(contactABI, &st, event.Name, *log)
	if err != nil {
		return fmt.Errorf("[EvmHook] contract %s event %s unpack log error: %s", contractAddr, event.Name, err.Error())
	}
	logger.Info("[EvmHook]", "contract-address", contractAddr, "event", event.Name, "unpack", st)
	return k.activeToken(ctx, st)
}

func (k Keeper) handleContractEventTransfer(ctx sdk.Context, log *ethtypes.Log, event *abi.Event, contactABI abi.ABI, govAddr common.Address) error {
	logger := ctx.Logger()
	contractAddr := log.Address
	if !bytes.Equal(contractAddr.Bytes(), govAddr.Bytes()) {
		logger.Debug("[EvmHook] contract address is not equal to gov address")
		return nil
	}
	var st types.EventTransfer

	err := evmcommon.UnpackLog(contactABI, &st, event.Name, *log)
	if err != nil {
		logger.Info("[EvmHook] unpack log failed", "contract-address", contractAddr, "event", event.Name, "error", err.Error())
		return nil //do not return error
	}
	logger.Info("[EvmHook]", "contract-address", contractAddr, "event", event.Name, "unpack", st)
	return nil
}

func (k Keeper) handleContractEventDeposit(ctx sdk.Context, log *ethtypes.Log, event *abi.Event, contactABI abi.ABI, swapAddr common.Address) error {
	logger := ctx.Logger()
	contractAddr := log.Address
	if !bytes.Equal(contractAddr.Bytes(), swapAddr.Bytes()) {
		logger.Debug("[EvmHook] contract address is not equal to swap address")
		return nil
	}
	var st types.EventSwapDeposit

	err := evmcommon.UnpackLog(contactABI, &st, event.Name, *log)
	if err != nil {
		return fmt.Errorf("[EvmHook] contract %s event %s unpack log error: %s", contractAddr, event.Name, err.Error())
	}
	logger.Info("[EvmHook]", "contract-address", contractAddr, "event", event.Name, "unpack", st)
	return k.swapDeposit(ctx, st)
}

func (k Keeper) handleContractEventWithdraw(ctx sdk.Context, log *ethtypes.Log, event *abi.Event, contactABI abi.ABI, swapAddr common.Address) error {
	logger := ctx.Logger()
	contractAddr := log.Address
	if !bytes.Equal(contractAddr.Bytes(), swapAddr.Bytes()) {
		logger.Debug("[EvmHook] contract address is not equal to swap address")
		return nil
	}
	var st types.EventSwapWithdraw

	err := evmcommon.UnpackLog(contactABI, &st, event.Name, *log)
	if err != nil {
		return fmt.Errorf("[EvmHook] contract %s event %s unpack log error: %s", contractAddr, event.Name, err.Error())
	}
	logger.Info("[EvmHook]", "contract-address", contractAddr, "event", event.Name, "unpack", st)
	return k.swapWithdraw(ctx, st)
}
