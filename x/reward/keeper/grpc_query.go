package keeper

import (
	"github.com/cascadiafoundation/cascadia/x/reward/types"
)

var _ types.QueryServer = Keeper{}
