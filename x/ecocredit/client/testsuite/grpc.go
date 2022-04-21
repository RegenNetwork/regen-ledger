package testsuite

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	types2 "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

const (
	marketplaceRoute = "/regen/ecocredit/marketplace/v1/"
	basketRoute      = "/regen/ecocredit/basket/v1/"
	coreRoute        = "/regen/ecocredit/v1/"
)

func (s *IntegrationTestSuite) TestQueryClasses() {
	val := s.network.Validators[0]
	ctx := val.ClientCtx
	_, err := s.createClass(ctx, &core.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String()},
		Metadata:         "m",
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &core.DefaultParams().CreditClassFee[0],
	})
	s.Require().NoError(err)

	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid query",
			fmt.Sprintf("%s%sclasses", val.APIAddress, coreRoute),
			false,
		},
		{
			"valid query pagination",
			fmt.Sprintf("%s%sclasses?pagination.limit=1", val.APIAddress, coreRoute),
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
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
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
	val := s.network.Validators[0]
	classId, err := s.createClass(val.ClientCtx, &core.MsgCreateClass{
		Admin:            val.Address.String(),
		Issuers:          []string{val.Address.String()},
		Metadata:         "m",
		CreditTypeAbbrev: validCreditTypeAbbrev,
		Fee:              &core.DefaultParams().CreditClassFee[0],
	})
	s.Require().NoError(err)

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid class-id",
			fmt.Sprintf("%s%sclasses/%s", val.APIAddress, coreRoute, classId),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var class core.QueryClassInfoResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &class)
			require.NoError(err)
			require.NotNil(class.Info)
			require.Contains(class.Info.Id, classId)

		})
	}
}

func (s *IntegrationTestSuite) TestQueryBatches() {
	val := s.network.Validators[0]
	_, pid, _ := s.createClassProjectBatch(val.ClientCtx, val.Address.String())

	testCases := []struct {
		name      string
		url       string
		paginated bool
	}{
		{
			"valid request",
			fmt.Sprintf("%s%sprojects/%s/batches", val.APIAddress, coreRoute, pid),
			false,
		},
		{
			"valid request with pagination",
			fmt.Sprintf("%s%sprojects/%s/batches?pagination.limit=2", val.APIAddress, coreRoute, pid),
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
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
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
	val := s.network.Validators[0]
	_, _, batchDenom := s.createClassProjectBatch(val.ClientCtx, val.Address.String())

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid request",
			fmt.Sprintf("%s%sbatches/%s", val.APIAddress, coreRoute, batchDenom),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var batch core.QueryBatchInfoResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &batch)
			require.NoError(err)
			require.NotNil(batch.Info)
			require.Equal(batch.Info.Denom, batchDenom)
		})
	}
}

func (s *IntegrationTestSuite) TestCreditTypes() {
	require := s.Require()
	val := s.network.Validators[0]
	creditTypes := core.DefaultParams().CreditTypes

	url := fmt.Sprintf("%s%scredit-types", val.APIAddress, coreRoute)
	resp, err := rest.GetRequest(url)
	require.NoError(err)

	var res core.QueryCreditTypesResponse
	err = val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
	require.NoError(err)
	require.Equal(creditTypes, res.CreditTypes)

}

func (s *IntegrationTestSuite) TestQueryBalance() {
	val := s.network.Validators[0]
	noBalAddr := s.network.Validators[1]
	_, _, batchDenom := s.createClassProjectBatch(val.ClientCtx, val.Address.String())

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid request",
			fmt.Sprintf("%s%sbatches/%s/balance/%s", val.APIAddress, coreRoute, batchDenom, val.Address.String()),
		},
		{
			"valid request - no balance",
			fmt.Sprintf("%s%sbatches/%s/balance/%s", val.APIAddress, coreRoute, batchDenom, noBalAddr.Address.String()),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QueryBalanceResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res)
			require.NotEmpty(res.Balance.Tradable)
			require.NotEmpty(res.Balance.Retired)
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySupply() {
	val := s.network.Validators[0]
	_, _, batchDenom := s.createClassProjectBatch(val.ClientCtx, val.Address.String())

	testCases := []struct {
		name string
		url  string
	}{
		{
			"valid request",
			fmt.Sprintf("%s%sbatches/%s/supply", val.APIAddress, coreRoute, batchDenom),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res core.QuerySupplyResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res)
			require.NotEmpty(res.RetiredSupply)
			require.NotEmpty(res.TradableSupply)

		})
	}
}

func (s *IntegrationTestSuite) TestGRPCQueryParams() {
	val := s.network.Validators[0]
	require := s.Require()

	resp, err := rest.GetRequest(fmt.Sprintf("%s%sparams", val.APIAddress, coreRoute))
	require.NoError(err)

	var res core.QueryParamsResponse
	require.NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, &res))
	s.Require().Equal(core.DefaultParams(), *res.Params)
}

func (s *IntegrationTestSuite) TestQuerySellOrder() {
	val := s.network.Validators[0]
	_, _, batchDenom := s.createClassProjectBatch(val.ClientCtx, val.Address.String())
	validAsk := types.NewInt64Coin(core.DefaultParams().AllowedAskDenoms[0].Denom, 10)
	expiration, err := types2.ParseDate("expiration", "2090-10-10")
	s.Require().NoError(err)
	orderIds, err := s.createSellOrder(val.ClientCtx, &marketplace.MsgSell{
		Owner: val.Address.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
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
			fmt.Sprintf("%s%ssell-orders/%d", val.APIAddress, marketplaceRoute, orderId),
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)

			var res marketplace.QuerySellOrderResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
			require.NoError(err)
			require.NotNil(res.SellOrder)
			require.Equal(res.SellOrder.Quantity, "10")
		})
	}
}

func (s *IntegrationTestSuite) TestQuerySellOrders() {
	val := s.network.Validators[0]
	_, _, batchDenom := s.createClassProjectBatch(val.ClientCtx, val.Address.String())
	validAsk := types.NewInt64Coin(core.DefaultParams().AllowedAskDenoms[0].Denom, 10)
	expiration, err := types2.ParseDate("expiration", "2090-10-10")
	s.Require().NoError(err)
	_, err = s.createSellOrder(val.ClientCtx, &marketplace.MsgSell{
		Owner: val.Address.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
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
			fmt.Sprintf("%s%ssell-orders", val.APIAddress, marketplaceRoute),
			-1,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s%ssell-orders?pagination.limit=2", val.APIAddress, marketplaceRoute),
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
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
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
	val := s.network.Validators[0]
	_, _, batchDenom := s.createClassProjectBatch(val.ClientCtx, val.Address.String())
	validAsk := types.NewInt64Coin(core.DefaultParams().AllowedAskDenoms[0].Denom, 10)
	expiration, err := types2.ParseDate("expiration", "2090-10-10")
	s.Require().NoError(err)
	_, err = s.createSellOrder(val.ClientCtx, &marketplace.MsgSell{
		Owner: val.Address.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
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
			fmt.Sprintf("%s%ssell-orders/batch-denom/%s", val.APIAddress, marketplaceRoute, batchDenom),
			-1,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s%ssell-orders/batch-denom/%s?pagination.limit=2", val.APIAddress, marketplaceRoute, batchDenom),
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
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
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

func (s *IntegrationTestSuite) TestQuerySellOrdersByAddress() {
	val := s.network.Validators[0]
	_, _, batchDenom := s.createClassProjectBatch(val.ClientCtx, val.Address.String())
	validAsk := types.NewInt64Coin(core.DefaultParams().AllowedAskDenoms[0].Denom, 10)
	expiration, err := types2.ParseDate("expiration", "2090-10-10")
	s.Require().NoError(err)
	_, err = s.createSellOrder(val.ClientCtx, &marketplace.MsgSell{
		Owner: val.Address.String(),
		Orders: []*marketplace.MsgSell_Order{
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
			{BatchDenom: batchDenom, Quantity: "10", AskPrice: &validAsk, Expiration: &expiration},
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
			fmt.Sprintf("%s%ssell-orders/address/%s", val.APIAddress, marketplaceRoute, val.Address.String()),
			-1,
		},
		{
			"valid request pagination",
			fmt.Sprintf("%s%ssell-orders/address/%s?pagination.limit=2", val.APIAddress, marketplaceRoute, val.Address.String()),
			2,
		},
	}

	require := s.Require()
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			resp, err := rest.GetRequest(tc.url)
			require.NoError(err)
			var res marketplace.QuerySellOrdersByAddressResponse
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, &res)
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
