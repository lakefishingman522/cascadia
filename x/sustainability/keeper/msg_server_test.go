package keeper_test

import (
	"testing"

	"github.com/cascadiafoundation/cascadia/x/sustainability/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *KeeperTestSuite) TestUpdatePenaltyAccount() {
	// Setting initial penalty account
	initialAcc := s.app.StakingKeeper.GetPenaltyAccount(s.ctx)

	tests := []struct {
		name     string
		creator  sdk.AccAddress
		newAddr  sdk.AccAddress
		expected sdk.AccAddress
		hasError bool
	}{
		{
			name:     "successful update",
			creator:  initialAcc,
			newAddr:  sdk.AccAddress([]byte("newAcc")),
			expected: sdk.AccAddress([]byte("newAcc")),
			hasError: false,
		},
		{
			name:     "non-matching creator",
			creator:  sdk.AccAddress([]byte("nonmatching")),
			newAddr:  sdk.AccAddress([]byte("anotherNew")),
			expected: sdk.AccAddress([]byte("newAcc")), // the address should not change from the previous test
			hasError: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			msg := &types.MsgUpdatePenaltyAccountRequest{
				Creator:    tt.creator.String(),
				NewAddress: tt.newAddr.String(),
			}

			_, err := s.app.PenaltyKeeper.UpdatePenaltyAccount(s.ctx, msg)
			if tt.hasError {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
			}

			// Check penalty account
			penaltyAcc := s.app.StakingKeeper.GetPenaltyAccount(s.ctx)
			s.Require().Equal(tt.expected, penaltyAcc)
		})
	}
}
