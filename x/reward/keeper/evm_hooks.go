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

// PostTxProcessing implements EvmHooks.PostTxProcessing.
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
	feeDistrContract, found := k.GetRewardContract(ctx, 1)
	if found {
		feeDistrAmount := sdk.NewDecFromInt(txFee).Mul(params.GasFeeDistribution.VecontractRewards).TruncateInt()
		feeDistFees := sdk.Coins{{Denom: evmDenom, Amount: feeDistrAmount}}

		feeDistAddress, err := sdk.AccAddressFromHexUnsafe(feeDistrContract.Address[2:])
		if err != nil {
			return err
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			k.feeCollectorName,
			feeDistAddress,
			feeDistFees,
		)

		if err != nil {
			return err
		}
	}

	// distribute the fees to the nProtocol
	nProtocol, found := k.GetRewardContract(ctx, 2)

	if found {
		nProtocolDist := sdk.NewDecFromInt(txFee).Mul(params.GasFeeDistribution.NprotocolRewards).TruncateInt()
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
