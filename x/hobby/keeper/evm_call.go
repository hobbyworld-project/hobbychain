// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package keeper

import (
	"encoding/json"
	"fmt"
	"github.com/hobbyworld-project/hobbychain/x/hobby/types"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/evmos/v15/server/config"
	evmtypes "github.com/evmos/evmos/v15/x/evm/types"

	"github.com/hobbyworld-project/hobbychain/contracts"
)

var abi721 = contracts.ERC721DelegateContract.ABI
var bin721 = contracts.ERC721DelegateContract.Bin
var abiSwap = contracts.WHBYContract.ABI
var binSwap = contracts.WHBYContract.Bin

// DeployGovContract creates and deploys an ERC721 contract on the EVM with the module account as owner.
func (k Keeper) DeployGovContract(
	ctx sdk.Context,
) (common.Address, error) {
	params := k.GetParams(ctx)

	ctx.Logger().Info("contract", "admin address", params.GovErc721.ContractAdminAddr)
	adminAddr := common.HexToAddress(params.GovErc721.ContractAdminAddr)
	ctorArgs, err := abi721.Pack(
		"",
		adminAddr,
		types.ModuleAddress,
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(types.ErrABIPack, "delegate contract abi pack error: %s", err.Error())
	}

	data := make([]byte, len(bin721)+len(ctorArgs))
	copy(data[:len(bin721)], bin721)
	copy(data[len(bin721):], ctorArgs)

	nonce, err := k.accountKeeper.GetSequence(ctx, types.ModuleAddress.Bytes())
	if err != nil {
		return common.Address{}, err
	}

	contractAddr := crypto.CreateAddress(types.ModuleAddress, nonce)
	_, err = k.CallEVMWithData(ctx, types.ModuleAddress, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy delegate contract")
	}
	ctx.Logger().Info("delegate contract deployed", "contract-address", contractAddr.String(), "gov-module-addr", types.ModuleAddress.String())
	return contractAddr, nil
}

// DeploySwapContract creates and deploys token swap contract on the EVM with the module account as owner.
func (k Keeper) DeploySwapContract(
	ctx sdk.Context,
) (common.Address, error) {

	ctorArgs, err := abiSwap.Pack(
		"",
	)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(types.ErrABIPack, "swap contract abi pack error: %s", err.Error())
	}

	data := make([]byte, len(binSwap)+len(ctorArgs))
	copy(data[:len(binSwap)], binSwap)
	copy(data[len(binSwap):], ctorArgs)

	nonce, err := k.accountKeeper.GetSequence(ctx, types.ModuleAddress.Bytes())
	if err != nil {
		return common.Address{}, err
	}

	contractAddr := crypto.CreateAddress(types.ModuleAddress, nonce)
	_, err = k.CallEVMWithData(ctx, types.ModuleAddress, nil, data, true)
	if err != nil {
		return common.Address{}, errorsmod.Wrapf(err, "failed to deploy swap contract")
	}
	ctx.Logger().Info("swap contract deployed", "contract-address", contractAddr.String(), "gov-module-addr", types.ModuleAddress.String())
	return contractAddr, nil
}

func (k Keeper) ContractQuery(ctx sdk.Context, method string, args ...interface{}) (res *evmtypes.MsgEthereumTxResponse, err error) {
	contractAddr, exist := k.GetGovContractAddr(ctx)
	if !exist {
		return nil, fmt.Errorf("no gov contract address found")
	}
	res, err = k.CallEVM(ctx, abi721, types.ModuleAddress, contractAddr, false, method, args...)
	if err != nil {
		return nil, fmt.Errorf("query evm method %s error %s", method, err.Error())
	}
	return res, nil
}

func (k Keeper) ContractCall(ctx sdk.Context, method string, args ...interface{}) (res *evmtypes.MsgEthereumTxResponse, err error) {
	contractAddr, exist := k.GetGovContractAddr(ctx)
	if !exist {
		return nil, fmt.Errorf("no gov contract address found")
	}
	res, err = k.CallEVM(ctx, abi721, types.ModuleAddress, contractAddr, true, method, args...)
	if err != nil {
		return nil, fmt.Errorf("call evm method %s error %s", method, err.Error())
	}
	return res, nil
}

// CallEVM performs a smart contract method call using given args
func (k Keeper) CallEVM(
	ctx sdk.Context,
	abi abi.ABI,
	from, contract common.Address,
	commit bool,
	method string,
	args ...interface{},
) (*evmtypes.MsgEthereumTxResponse, error) {
	data, err := abi.Pack(method, args...)
	if err != nil {
		return nil, errorsmod.Wrap(
			types.ErrABIPack,
			errorsmod.Wrap(err, "failed to create transaction data").Error(),
		)
	}

	resp, err := k.CallEVMWithData(ctx, from, &contract, data, commit)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "contract call failed: method '%s', contract '%s'", method, contract)
	}
	return resp, nil
}

// CallEVMWithData performs a smart contract method call using contract data
func (k Keeper) CallEVMWithData(
	ctx sdk.Context,
	from common.Address,
	contract *common.Address,
	data []byte,
	commit bool,
) (*evmtypes.MsgEthereumTxResponse, error) {
	nonce, err := k.accountKeeper.GetSequence(ctx, from.Bytes())
	if err != nil {
		return nil, err
	}

	gasCap := config.DefaultGasCap
	if commit {
		args, err := json.Marshal(evmtypes.TransactionArgs{
			From: &from,
			To:   contract,
			Data: (*hexutil.Bytes)(&data),
		})
		if err != nil {
			return nil, errorsmod.Wrapf(errortypes.ErrJSONMarshal, "failed to marshal tx args: %s", err.Error())
		}

		gasRes, err := k.evmKeeper.EstimateGas(sdk.WrapSDKContext(ctx), &evmtypes.EthCallRequest{
			Args:   args,
			GasCap: config.DefaultGasCap,
		})
		if err != nil {
			return nil, err
		}
		gasCap = gasRes.Gas
	}

	msg := ethtypes.NewMessage(
		from,
		contract,
		nonce,
		big.NewInt(0), // amount
		gasCap,        // gasLimit
		big.NewInt(0), // gasFeeCap
		big.NewInt(0), // gasTipCap
		big.NewInt(0), // gasPrice
		data,
		ethtypes.AccessList{}, // AccessList
		!commit,               // isFake
	)

	res, err := k.evmKeeper.ApplyMessage(ctx, msg, evmtypes.NewNoOpTracer(), commit)
	if err != nil {
		return nil, err
	}

	if res.Failed() {
		return nil, errorsmod.Wrap(evmtypes.ErrVMExecution, res.VmError)
	}

	return res, nil
}

func (k Keeper) MaxVote(ctx sdk.Context) (*big.Int, error) {
	return k.evmPublicBigInt(ctx, types.ContractMethodMaxVote)
}

func (k Keeper) MinVote(ctx sdk.Context) (*big.Int, error) {
	return k.evmPublicBigInt(ctx, types.ContractMethodMinVote)
}

func (k Keeper) HeightDiff(ctx sdk.Context) (*big.Int, error) {
	return k.evmPublicBigInt(ctx, types.ContractMethodHeightDiff)
}

func (k Keeper) evmPublicBigInt(ctx sdk.Context, method string) (*big.Int, error) {
	var err error
	var res *evmtypes.MsgEthereumTxResponse
	res, err = k.ContractQuery(ctx, method)
	if err != nil {
		return nil, err
	}
	var unpacked []interface{}
	unpacked, err = abi721.Unpack(method, res.Ret)
	if err != nil || len(unpacked) == 0 {
		return nil, err
	}
	val, ok := unpacked[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("method %s return value type dismatch, expect uint256", method)
	}
	return val, nil
}
