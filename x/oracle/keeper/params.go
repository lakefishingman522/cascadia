package keeper

import (
	"fmt"
	"reflect"

	"github.com/cascadiafoundation/cascadia/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {

	defaultParams := types.DefaultParams()
	// k.paramstore.GetParamSetIfExists(ctx, &params)
	for _, pair := range defaultParams.ParamSetPairs() {
		// pair.Field is a pointer to the field, so indirecting the ptr.
		// go-amino automatically handles it but just for sure,
		// since SetStruct is meant to be used in InitGenesis
		// so this method will not be called frequently
		v := reflect.Indirect(reflect.ValueOf(pair.Value)).Interface()

		if err := pair.ValidatorFn(v); err != nil {
			panic(fmt.Sprintf("value from ParamSetPair is invalid: %s", err))
		}
		exist := k.paramstore.Has(ctx, pair.Key)
		// fmt.Println(exist, v)
		if exist != true {
			k.paramstore.Set(ctx, pair.Key, v)
		}
	}

	k.paramstore.GetParamSetIfExists(ctx, &params)

	return
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
