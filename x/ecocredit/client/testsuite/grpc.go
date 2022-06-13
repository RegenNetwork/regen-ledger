package testsuite

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	types2 "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

const (
	marketplaceRoute = "regen/ecocredit/marketplace/v1"
	basketRoute      = "regen/ecocredit/basket/v1"
	coreRoute        = "regen/ecocredit/v1"
)

func (s *IntegrationTestSuite) TestQueryClasses() {
	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/classes", s.val.APIAddress, coreRoute),
			false,
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/classes?pagination.limit=1", s.val.APIAddress, coreRoute),
			true,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QueryClassesResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Classes)
			require.True(len(res.Classes) > 0)
			if tc.paginated {
				require.NotNil(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryClass() {
	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/class/%s", s.val.APIAddress, coreRoute, s.classId),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/classes/%s", s.val.APIAddress, coreRoute, s.classId),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QueryClassResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Class)
			require.Equal(res.Class.Id, s.classId)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProject() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/project/%s", s.val.APIAddress, coreRoute, s.projectId),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/projects/%s", s.val.APIAddress, coreRoute, s.projectId),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res core.QueryProjectResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Project)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProjects() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/projects", s.val.APIAddress, coreRoute),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/projects?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				coreRoute,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res core.QueryProjectsResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Projects)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Projects, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			} else {
				require.Empty(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProjectsByClass() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/projects-by-class/%s", s.val.APIAddress, coreRoute, s.classId),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/projects-by-class/%s?pagination.countTotal=true",
				// TODO: #1113
				// "%s/%s/projects-by-class/%s?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				coreRoute,
				s.classId,
			),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/projects/class/%s", s.val.APIAddress, coreRoute, s.classId),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/classes/%s/projects", s.val.APIAddress, coreRoute, s.classId),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res core.QueryProjectsByClassResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Projects)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Projects, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			} else {
				require.Empty(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryProjectsByReferenceId() {
	require := s.Require()

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf(
				"%s/%s/projects-by-reference-id/%s",
				s.val.APIAddress,
				coreRoute,
				s.projectReferenceId,
			),
		},
		{
			"valid with pagination",
			fmt.Sprintf(
				"%s/%s/projects-by-reference-id/%s?pagination.limit=1&pagination.countTotal=true",
				s.val.APIAddress,
				coreRoute,
				s.projectReferenceId,
			),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/projects/reference-id/%s",
				s.val.APIAddress,
				coreRoute,
				s.projectReferenceId,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			bz, err := rest.GetRequest(tc.url)
			require.NoError(err)
			require.NotContains(string(bz), "code")

			var res core.QueryProjectsByReferenceIdResponse
			require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(bz, &res))
			require.NotEmpty(res.Projects)

			if strings.Contains(tc.name, "pagination") {
				require.Len(res.Projects, 1)
				require.NotEmpty(res.Pagination)
				require.NotEmpty(res.Pagination.Total)
			} else {
				require.Empty(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatches() {
	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/batches", s.val.APIAddress, coreRoute),
			false,
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/batches?pagination.limit=2", s.val.APIAddress, coreRoute),
			true,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QueryBatchesResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Batches)
			require.Greater(len(res.Batches), 0)
			if tc.paginated {
				require.NotNil(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesByIssuer() {
	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/batches-by-issuer/%s", s.val.APIAddress, coreRoute, s.addr1),
			false,
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/batches-by-issuer/%s?pagination.limit=2", s.val.APIAddress, coreRoute, s.addr1),
			true,
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/issuer/%s", s.val.APIAddress, coreRoute, s.addr1),
			false,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QueryBatchesByIssuerResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Batches)
			require.Greater(len(res.Batches), 0)
			if tc.paginated {
				require.NotNil(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesByClass() {
	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/batches-by-class/%s", s.val.APIAddress, coreRoute, s.classId),
			false,
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/batches-by-class/%s?pagination.limit=2", s.val.APIAddress, coreRoute, s.classId),
			true,
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/class/%s", s.val.APIAddress, coreRoute, s.classId),
			false,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QueryBatchesByClassResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Batches)
			require.Greater(len(res.Batches), 0)
			if tc.paginated {
				require.NotNil(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatchesByProject() {
	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/batches-by-project/%s", s.val.APIAddress, coreRoute, s.projectId),
			false,
		},
		{
			"valid with pagination",
			fmt.Sprintf("%s/%s/batches-by-project/%s?pagination.limit=2", s.val.APIAddress, coreRoute, s.projectId),
			true,
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/project/%s", s.val.APIAddress, coreRoute, s.projectId),
			false,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QueryBatchesResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Batches)
			require.Greater(len(res.Batches), 0)
			if tc.paginated {
				require.NotNil(res.Pagination)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatch() {
	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/batch/%s", s.val.APIAddress, coreRoute, s.batchDenom),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/%s", s.val.APIAddress, coreRoute, s.batchDenom),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QueryBatchResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.Batch)
			require.Equal(res.Batch.Denom, s.batchDenom)
		})
	}
}

func (s *IntegrationTestSuite) TestCreditTypes() {
	require := s.Require()

	url := fmt.Sprintf("%s/%s/credit-types", s.val.APIAddress, coreRoute)
	resp, err := rest.GetRequest(url)
	require.NoError(err)

	var res core.QueryCreditTypesResponse
	err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
	require.NoError(err)
	require.Greater(len(res.CreditTypes), 0)
}

func (s *IntegrationTestSuite) TestQueryBalance() {
	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/balance/%s/%s", s.val.APIAddress, coreRoute, s.batchDenom, s.addr1),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/%s/balance/%s", s.val.APIAddress, coreRoute, s.batchDenom, s.addr1),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/balances/%s/batch/%s", s.val.APIAddress, coreRoute, s.addr1, s.batchDenom),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QueryBalanceResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res)
			require.NotEmpty(res.Balance.TradableAmount)
			require.NotEmpty(res.Balance.RetiredAmount)
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySupply() {
	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid",
			fmt.Sprintf("%s/%s/supply/%s", s.val.APIAddress, coreRoute, s.batchDenom),
		},
		{
			"valid alternative",
			fmt.Sprintf("%s/%s/batches/%s/supply", s.val.APIAddress, coreRoute, s.batchDenom),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QuerySupplyResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res)
			require.NotEmpty(res.RetiredSupply)
			require.NotEmpty(res.TradableSupply)
		})
	}
}

func (s *IntegrationTestSuite) TestQueryParams() {
	require := s.Require()

	resp, err := rest.GetRequest(fmt.Sprintf("%s/%s/params", s.val.APIAddress, coreRoute))
	require.NoError(err)

	var res core.QueryParamsResponse
	require.NoError(s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res))
	s.Require().Equal(core.DefaultParams(), *res.Params)
}

func (s *IntegrationTestSuite) TestQuerySellOrder() {
	validAsk := types.NewInt64Coin(types.DefaultBondDenom, 10)
	expiration, err := types2.ParseDate("expiration", "2090-10-10")
	s.Require().NoError(err)
	orderIds, err := s.createSellOrder(s.val.ClientCtx, &marketplace.MsgSell{
		Seller: s.addr1.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: s.batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
		},
	})
	s.Require().NoError(err)
	orderId := orderIds[0]

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid request",
			fmt.Sprintf("%s/%s/sell-orders/%d", s.val.APIAddress, marketplaceRoute, orderId),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res marketplace.QuerySellOrderResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.SellOrder)
			require.Equal(res.SellOrder.Quantity, "10")
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrders() {
	validAsk := types.NewInt64Coin(types.DefaultBondDenom, 10)
	expiration, err := types2.ParseDate("expiration", "2090-10-10")
	s.Require().NoError(err)
	_, err = s.createSellOrder(s.val.ClientCtx, &marketplace.MsgSell{
		Seller: s.addr1.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: s.batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: s.batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: s.batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
		},
	})
	s.Require().NoError(err)

	testCases := []struct {
		name     string
		url      string
		expItems int
	}{
		{
			"valid request",
			fmt.Sprintf("%s/%s/sell-orders", s.val.APIAddress, marketplaceRoute),
			-1,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/%s/sell-orders?pagination.limit=2", s.val.APIAddress, marketplaceRoute),
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res marketplace.QuerySellOrdersResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.SellOrders)
			if tc.expItems > 0 {
				require.Len(res.SellOrders, tc.expItems)
			} else {
				require.Greater(len(res.SellOrders), 0)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersByBatchDenom() {
	validAsk := types.NewInt64Coin(types.DefaultBondDenom, 10)
	expiration, err := types2.ParseDate("expiration", "2090-10-10")
	s.Require().NoError(err)
	_, err = s.createSellOrder(s.val.ClientCtx, &marketplace.MsgSell{
		Seller: s.addr1.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: s.batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: s.batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: s.batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
		},
	})
	s.Require().NoError(err)

	testCases := []struct {
		name     string
		url      string
		expItems int
	}{
		{
			"valid request",
			fmt.Sprintf("%s/%s/sell-orders/batch-denom/%s", s.val.APIAddress, marketplaceRoute, s.batchDenom),
			-1,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/%s/sell-orders/batch-denom/%s?pagination.limit=2", s.val.APIAddress, marketplaceRoute, s.batchDenom),
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err, err)
			var res marketplace.QuerySellOrdersByBatchDenomResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err, err)
			require.NotNil(res.SellOrders)
			if tc.expItems > 0 {
				require.Len(res.SellOrders, tc.expItems)
			} else {
				require.Greater(len(res.SellOrders), 0)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrdersBySeller() {
	validAsk := types.NewInt64Coin(types.DefaultBondDenom, 10)
	expiration, err := types2.ParseDate("expiration", "2090-10-10")
	s.Require().NoError(err)
	_, err = s.createSellOrder(s.val.ClientCtx, &marketplace.MsgSell{
		Seller: s.addr1.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: s.batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: s.batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: s.batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
		},
	})
	s.Require().NoError(err)

	testCases := []struct {
		name     string
		url      string
		expItems int
	}{
		{
			"valid request",
			fmt.Sprintf("%s/%s/sell-orders/seller/%s", s.val.APIAddress, marketplaceRoute, s.addr1),
			-1,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s/%s/sell-orders/seller/%s?pagination.limit=2", s.val.APIAddress, marketplaceRoute, s.addr1),
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)
			var res marketplace.QuerySellOrdersBySellerResponse
			err = s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.SellOrders)

			if tc.expItems > 0 {
				require.Len(res.SellOrders, tc.expItems)
			} else {
				require.Greater(len(res.SellOrders), 0)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestQueryAllowedDenoms() {
	url := fmt.Sprintf("%s/%s/allowed-denoms?pagination.count_total=true", s.val.APIAddress, marketplaceRoute)
	resp, err := rest.GetRequest(url)
	s.Require().NoError(err)

	var res marketplace.QueryAllowedDenomsResponse
	s.Require().NoError(s.val.ClientCtx.Codec.UnmarshalJSON(resp, &res))
	s.Require().Equal(uint64(len(s.allowedDenoms)), res.Pagination.Total)
	for _, d := range res.AllowedDenoms {
		s.Require().Contains(s.allowedDenoms, d.BankDenom)
	}
}
