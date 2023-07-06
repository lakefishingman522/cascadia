package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"


)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	HasAccount(sdk.Context, sdk.AccAddress) bool
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}


type StakingKeeper interface{
	SetPenaltyAccount(ctx sdk.Context, penaltyAccount PenaltyAccount)
	GetPenaltyAccount(ctx sdk.Context) (val PenaltyAccount, found bool)
	RemovePenaltyAccount(ctx sdk.Context)
}
