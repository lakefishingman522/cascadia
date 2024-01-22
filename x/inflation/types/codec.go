package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

var amino = codec.NewLegacyAmino()

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&InflationCreateControlParamsProposal{}, "inflation/InflationCreateControlParamsProposal", nil)
	cdc.RegisterConcrete(&InflationUpdateControlParamsProposal{}, "inflation/InflationUpdateControlParamsProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&InflationCreateControlParamsProposal{},
		&InflationUpdateControlParamsProposal{},
	)

	// this line is used by starport scaffolding # 3
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
