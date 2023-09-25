package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/RegenNetwork/regen-ledger/api/v2/regen/ecocredit/v1"
	regenerrors "github.com/RegenNetwork/regen-ledger/types/v2/errors"
	"github.com/RegenNetwork/regen-ledger/types/v2/ormutil"
	types "github.com/RegenNetwork/regen-ledger/x/ecocredit/v3/base/types/v1"
)

// ClassIssuers returns a list of addresses that are allowed to issue batches from the given class.
func (k Keeper) ClassIssuers(ctx context.Context, request *types.QueryClassIssuersRequest) (*types.QueryClassIssuersResponse, error) {
	pg, err := ormutil.GogoPageReqToPulsarPageReq(request.Pagination)
	if err != nil {
		return nil, regenerrors.ErrInvalidArgument.Wrap(err.Error())
	}

	classInfo, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, regenerrors.ErrNotFound.Wrapf("could not get class with id %s: %s", request.ClassId, err.Error())
	}

	it, err := k.stateStore.ClassIssuerTable().List(ctx, api.ClassIssuerClassKeyIssuerIndexKey{}.WithClassKey(classInfo.Key), ormlist.Paginate(pg))
	if err != nil {
		return nil, err
	}
	defer it.Close()

	issuers := make([]string, 0)
	for it.Next() {
		val, err := it.Value()
		if err != nil {
			return nil, err
		}
		issuers = append(issuers, sdk.AccAddress(val.Issuer).String())
	}
	pr, err := ormutil.PulsarPageResToGogoPageRes(it.PageResponse())
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrap(err.Error())
	}
	return &types.QueryClassIssuersResponse{
		Issuers:    issuers,
		Pagination: pr,
	}, nil
}
