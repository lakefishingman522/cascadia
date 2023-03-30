package types

import (
	"errors"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyGasFeeDistribution = []byte("GasFeeDistribution")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(gasFeeDistribution GasFeeDistribution) Params {
	return Params{
		GasFeeDistribution: gasFeeDistribution,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(GasFeeDistribution{
		StakingRewards:    sdk.NewDecWithPrec(333333333, 9),
		VecontractRewards: sdk.NewDecWithPrec(333333333, 9),
		NprotocolRewards:  sdk.NewDecWithPrec(333333334, 9),
	})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyGasFeeDistribution, &p.GasFeeDistribution, validateGasFeeDistribution),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateGasFeeDistribution(p.GasFeeDistribution); err != nil {
		return err
	}

	return nil
}

func validateGasFeeDistribution(i interface{}) error {
	v, ok := i.(GasFeeDistribution)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.StakingRewards.IsNegative() {
		return errors.New("staking distribution ratio must not be negative")
	}

	if v.VecontractRewards.IsNegative() {
		return errors.New("vecontract distribution ratio must not be negative")
	}

	if v.NprotocolRewards.IsNegative() {
		return errors.New("nprotocol distribution ratio must not be negative")
	}

	totalProportions := v.StakingRewards.Add(v.VecontractRewards).Add(v.NprotocolRewards)
	if !totalProportions.Equal(sdk.NewDec(1)) {
		return errors.New("total distributions ratio should be 1")
	}

	return nil
}
