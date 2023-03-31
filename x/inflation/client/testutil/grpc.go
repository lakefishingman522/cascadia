package testutil

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"

	"github.com/gogo/protobuf/proto"

	inflationtypes "github.com/cascadiafoundation/cascadia/x/inflation/types"
)

func (s *IntegrationTestSuite) TestQueryGRPC() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress
	testCases := []struct {
		name     string
		url      string
		headers  map[string]string
		respType proto.Message
		expected proto.Message
	}{
		{
			"gRPC request params",
			fmt.Sprintf("%s/inflation/v1/params", baseURL),
			map[string]string{},
			&inflationtypes.QueryParamsResponse{},
			&inflationtypes.QueryParamsResponse{
				Params: inflationtypes.NewParams("stake", sdk.NewDecWithPrec(13, 2), sdk.NewDecWithPrec(100, 2),
					sdk.NewDec(1), sdk.NewDecWithPrec(67, 2), (60 * 60 * 8766 / 5), inflationtypes.InflationDistribution{
						StakingRewards:    sdk.NewDecWithPrec(333333333, 9), // 33%
						VecontractRewards: sdk.NewDecWithPrec(333333333, 9), // 33%
						NprotocolRewards:  sdk.NewDecWithPrec(333333334, 9), // 33%
					}),
			},
		},
		{
			"gRPC request inflation",
			fmt.Sprintf("%s/inflation/v1/inflation", baseURL),
			map[string]string{},
			&inflationtypes.QueryInflationRateResponse{},
			&inflationtypes.QueryInflationRateResponse{
				InflationRate: sdk.NewDec(1),
			},
		},
		{
			"gRPC request annual provisions",
			fmt.Sprintf("%s/inflation/v1/annual_provisions", baseURL),
			map[string]string{
				grpctypes.GRPCBlockHeightHeader: "1",
			},
			&inflationtypes.QueryAnnualProvisionsResponse{},
			&inflationtypes.QueryAnnualProvisionsResponse{
				AnnualProvisions: sdk.NewDec(500000000),
			},
		},
	}
	for _, tc := range testCases {
		resp, err := testutil.GetRequestWithHeaders(tc.url, tc.headers)
		s.Run(tc.name, func() {
			s.Require().NoError(err)
			s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, tc.respType))
			s.Require().Equal(tc.expected.String(), tc.respType.String())
		})
	}
}
