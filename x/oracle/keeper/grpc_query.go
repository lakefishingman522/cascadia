package keeper

import (
	"github.com/cascadiafoundation/cascadia/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
