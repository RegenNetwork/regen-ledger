package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	regenerrors "github.com/regen-network/regen-ledger/errors"
	regentypes "github.com/regen-network/regen-ledger/types"
	types "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

// ClassFee queries credit class creation fees.
func (k Keeper) ClassFee(ctx context.Context, request *types.QueryClassFeeRequest) (*types.QueryClassFeeResponse, error) {
	classFee, err := k.stateStore.ClassFeeTable().Get(ctx)
	if err != nil {
		return nil, regenerrors.ErrInternal.Wrapf("failed to get class fee: %s", err.Error())
	}

	var fee sdk.Coin
	if classFee.Fee != nil {
		var ok bool
		fee, ok = regentypes.ProtoCoinToCoin(classFee.Fee)
		if !ok {
			return nil, regenerrors.ErrInternal.Wrapf("failed to parse class fee")
		}
	}

	return &types.QueryClassFeeResponse{
		Fee: &fee,
	}, nil
}
