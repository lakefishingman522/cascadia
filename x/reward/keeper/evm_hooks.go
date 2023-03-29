package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	evmtypes "github.com/cascadiafoundation/cascadia/x/evm/types"
)

var _ evmtypes.EvmHooks = Hooks{}

// Hooks wrapper struct for fees keeper
type Hooks struct {
	k Keeper
}

// Hooks return the wrapper hooks struct for the Keeper
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// PostTxProcessing is a wrapper for calling the EVM PostTxProcessing hook on
// the module keeper
func (h Hooks) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	return h.k.PostTxProcessing(ctx, msg, receipt)
}

// PostTxProcessing implements EvmHooks.PostTxProcessing. After each successful
// interaction with a registered contract, the contract deployer (or, if set,
// the withdraw address) receives a share from the transaction fees paid by the
// transaction sender.
func (k Keeper) PostTxProcessing(
	ctx sdk.Context,
	msg core.Message,
	receipt *ethtypes.Receipt,
) error {
	// check if the fees are globally enabled
	params := k.GetParams(ctx)

	txFee := sdk.NewIntFromUint64(receipt.GasUsed).Mul(sdk.NewIntFromBigInt(msg.GasPrice()))
	evmDenom := k.evmKeeper.GetParams(ctx).EvmDenom

	// distribute the fees to the VE contract
	feeDistrContract, found := k.GetRewardContract(ctx, 0)
	if found {
		veContractDist := sdk.NewDecFromInt(txFee).Mul(params.GasFeeDistribution.VecontractRewards).TruncateInt()
		veContractFees := sdk.Coins{{Denom: evmDenom, Amount: veContractDist}}

		veContractAddress, err := sdk.AccAddressFromHexUnsafe(feeDistrContract.Address[2:])
		if err != nil {
			return err
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			k.feeCollectorName,
			veContractAddress,
			veContractFees,
		)

		if err != nil {
			return err
		}
	}

	// distribute the fees to the nProtocol
	nProtocol, found := k.GetRewardContract(ctx, 1)

	if found {
		nProtocolDist := sdk.NewDecFromInt(txFee).Mul(params.GasFeeDistribution.PotocolRewards).TruncateInt()
		nProtocolFees := sdk.Coins{{Denom: evmDenom, Amount: nProtocolDist}}

		nProtocolAddress, err := sdk.AccAddressFromHexUnsafe(nProtocol.Address[2:])
		if err != nil {
			return nil
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			k.feeCollectorName,
			nProtocolAddress,
			nProtocolFees,
		)

		return err
	}

	return nil
}
