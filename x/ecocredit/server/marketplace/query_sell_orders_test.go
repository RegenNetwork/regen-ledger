package marketplace

import (
	"context"
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecocreditApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/marketplace"
)

var (
	classId    = "C01"
	batchDenom = "C01-20200101-20200201-001"
	start, end = timestamppb.Now(), timestamppb.Now()
	ask        = sdk.NewInt64Coin("ufoo", 10)
	creditType = core.CreditType{Name: "carbon", Abbreviation: "C", Unit: "tonnes", Precision: 6}
)

func TestSellOrders(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)
	_, _, addr2 := testdata.KeyTestPubAddr()

	order1 := insertSellOrder(t, s, s.addr, 1)
	order2 := insertSellOrder(t, s, addr2, 1)

	res, err := s.k.SellOrders(s.ctx, &marketplace.QuerySellOrdersRequest{
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.SellOrders))
	assertOrderEqual(t, s.ctx, s.k, res.SellOrders[0], order1)
	assertOrderEqual(t, s.ctx, s.k, res.SellOrders[1], order2)
	assert.Equal(t, uint64(2), res.Pagination.Total)
}

func TestSellOrdersByDenom(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	// make another batch
	otherDenom := "C01-19990101-20290101-001"
	assert.NilError(t, s.coreStore.BatchTable().Insert(s.ctx, &ecocreditApi.Batch{
		ProjectKey: 1,
		Denom:      otherDenom,
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	order1 := insertSellOrder(t, s, s.addr, 1)
	order2 := insertSellOrder(t, s, s.addr, 2)

	// query the first denom
	res, err := s.k.SellOrdersByBatchDenom(s.ctx, &marketplace.QuerySellOrdersByBatchDenomRequest{
		BatchDenom: batchDenom,
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))
	assertOrderEqual(t, s.ctx, s.k, res.SellOrders[0], order1)
	assert.Equal(t, uint64(1), res.Pagination.Total)

	// query the second denom
	res, err = s.k.SellOrdersByBatchDenom(s.ctx, &marketplace.QuerySellOrdersByBatchDenomRequest{
		BatchDenom: otherDenom,
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))
	assertOrderEqual(t, s.ctx, s.k, res.SellOrders[0], order2)
	assert.Equal(t, uint64(1), res.Pagination.Total)

	// bad denom should error
	res, err = s.k.SellOrdersByBatchDenom(s.ctx, &marketplace.QuerySellOrdersByBatchDenomRequest{
		BatchDenom: "yikes!",
		Pagination: nil,
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}

func TestSellOrdersByAddress(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	testSellSetup(t, s, batchDenom, ask.Denom, ask.Denom[1:], classId, start, end, creditType)

	_, _, otherAddr := testdata.KeyTestPubAddr()
	_, _, noOrdersAddr := testdata.KeyTestPubAddr()

	order1 := insertSellOrder(t, s, s.addr, 1)
	order2 := insertSellOrder(t, s, otherAddr, 1)

	res, err := s.k.SellOrdersByAddress(s.ctx, &marketplace.QuerySellOrdersByAddressRequest{
		Address:    s.addr.String(),
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))
	assertOrderEqual(t, s.ctx, s.k, res.SellOrders[0], order1)
	assert.Equal(t, uint64(1), res.Pagination.Total)

	res, err = s.k.SellOrdersByAddress(s.ctx, &marketplace.QuerySellOrdersByAddressRequest{
		Address:    otherAddr.String(),
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.SellOrders))
	assertOrderEqual(t, s.ctx, s.k, res.SellOrders[0], order2)
	assert.Equal(t, uint64(1), res.Pagination.Total)

	// addr with no sell orders should just return empty slice
	res, err = s.k.SellOrdersByAddress(s.ctx, &marketplace.QuerySellOrdersByAddressRequest{
		Address:    noOrdersAddr.String(),
		Pagination: &query.PageRequest{CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 0, len(res.SellOrders))
	assert.Equal(t, uint64(0), res.Pagination.Total)

	// bad address should fail
	res, err = s.k.SellOrdersByAddress(s.ctx, &marketplace.QuerySellOrdersByAddressRequest{
		Address:    "foobar1vlk23jrkl",
		Pagination: nil,
	})
	assert.ErrorContains(t, err, "decoding bech32 failed")
}

func insertSellOrder(t *testing.T, s *baseSuite, addr sdk.AccAddress, batchId uint64) *api.SellOrder {
	sellOrder := &api.SellOrder{
		Seller:            addr,
		BatchId:           batchId,
		Quantity:          "10",
		MarketId:          1,
		AskPrice:          "10",
		DisableAutoRetire: false,
		Expiration:        timestamppb.Now(),
		Maker:             false,
	}
	assert.NilError(t, s.marketStore.SellOrderTable().Insert(s.ctx, sellOrder))

	return sellOrder
}

func assertOrderEqual(t *testing.T, ctx context.Context, k Keeper, received *marketplace.SellOrderInfo, order *api.SellOrder) {
	seller := sdk.AccAddress(order.Seller)

	batch, err := k.coreStore.BatchTable().Get(ctx, order.BatchId)
	assert.NilError(t, err)

	market, err := k.stateStore.MarketTable().Get(ctx, order.MarketId)
	assert.NilError(t, err)

	info := marketplace.SellOrderInfo{
		Id:                order.Id,
		Seller:            seller.String(),
		BatchDenom:        batch.Denom,
		Quantity:          order.Quantity,
		AskDenom:          market.BankDenom,
		AskPrice:          order.AskPrice,
		DisableAutoRetire: order.DisableAutoRetire,
		Expiration:        types.ProtobufToGogoTimestamp(order.Expiration),
	}

	assert.DeepEqual(t, info, *received)
}
