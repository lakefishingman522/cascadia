package inflation

import (
	"time"

	"github.com/cascadiafoundation/cascadia/x/inflation/keeper"
	"github.com/cascadiafoundation/cascadia/x/inflation/types"
	otypes "github.com/cascadiafoundation/cascadia/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, ic types.InflationCalculationFn) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	btcPrice, found := k.GetLatestPriceFromAssetAndSource(ctx, otypes.BTC, otypes.BAND)
	if found {
		btcPrice = otypes.Price{
			Asset:     otypes.BTC,
			Price:     sdk.NewDec(0),
			Source:    otypes.BAND,
			Provider:  "",
			Timestamp: 0,
		}
	}

	// recalculate inflation rate
	totalStakingSupply := k.StakingTokenSupply(ctx)
	bondedRatio := k.BondedRatio(ctx)
	minter.Inflation = ic(ctx, minter, params, bondedRatio, btcPrice)
	minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply)
	k.SetMinter(ctx, minter)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params)
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	k.AddContractIncentives(ctx, mintedCoin, 1)
	k.AddContractIncentives(ctx, mintedCoin, 2)

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx)
	if err != nil {
		panic(err)
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyBondedRatio, bondedRatio.String()),
			sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
			sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
