package types

import (
	otypes "github.com/cascadiafoundation/cascadia/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InflationCalculationFn defines the function required to calculate inflation rate during
// BeginBlock. It receives the minter and params stored in the keeper, along with the current
// bondedRatio and returns the newly calculated inflation rate.
// It can be used to specify a _Default inflation calculation logic, instead of relying on the
// default logic provided by the sdk.
type InflationCalculationFn func(ctx sdk.Context, minter Minter, params Params, bondedRatio sdk.Dec, cascadiacoinPrice otypes.Price) sdk.Dec

// InflationCalculationFn_ implemented customized Inflaion Calculation Logic
type InflationCalculationFn_ func(ctx sdk.Context, minter Minter, params Params, bondedRatio sdk.Dec, priceStatistics otypes.PriceStatistics, inflationControlParams InflationControlParams) sdk.Dec

// DefaultInflationCalculationFn is the default function used to calculate inflation.
func DefaultInflationCalculationFn(_ sdk.Context, minter Minter, params Params, bondedRatio sdk.Dec, cascadiacoinPrice otypes.Price) sdk.Dec {
	return minter.NextInflationRate(params, bondedRatio, cascadiacoinPrice)
}

// DefaultInflationCalculationFn_ is the customized function used to calculation inflation
func DefaultInflationCalculationFn_(ctx sdk.Context, minter Minter, params Params, bondedRatio sdk.Dec, priceStatistics otypes.PriceStatistics, inflationControlParams InflationControlParams) sdk.Dec {
	return minter._NextInflationRate(params, bondedRatio, priceStatistics, inflationControlParams)
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(minter Minter, params Params, inflation_control_params InflationControlParams) *GenesisState {
	return &GenesisState{
		Minter:                 minter,
		Params:                 params,
		InflationControlParams: inflation_control_params,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Minter:                 DefaultInitialMinter(),
		Params:                 DefaultParams(),
		InflationControlParams: DefaultInflationControlParams(),
	}
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}

	return ValidateMinter(data.Minter)
}

// default inflation control related parameters
func DefaultInflationControlParams() InflationControlParams {
	return InflationControlParams{
		Lambda: sdk.NewDecWithPrec(1, 1),  // 1.000000
		W360:   sdk.NewDecWithPrec(5, 2),  // 0.05
		W180:   sdk.NewDecWithPrec(5, 2),  // 0.05
		W90:    sdk.NewDecWithPrec(5, 2),  // 0.05
		W30:    sdk.NewDecWithPrec(5, 2),  // 0.05
		W14:    sdk.NewDecWithPrec(5, 2),  // 0.05
		W7:     sdk.NewDecWithPrec(15, 2), // 0.15
		W1:     sdk.NewDecWithPrec(60, 2), // 0.6
	}
}
