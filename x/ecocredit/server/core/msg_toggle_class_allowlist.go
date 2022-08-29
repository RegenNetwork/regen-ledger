package core

import (
	"context"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ecocreditv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func (k Keeper) ToggleCreditClassAllowlist(ctx context.Context, req *core.MsgToggleCreditClassAllowlist) (*core.MsgToggleCreditClassAllowlistResponse, error) {
	if k.authority.String() != req.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf("invalid authority: expected %s, got %s", k.authority, req.Authority)
	}

	if err := k.stateStore.AllowListEnabledTable().Save(ctx, &ecocreditv1.AllowListEnabled{
		Enabled: req.Enabled,
	}); err != nil {
		return nil, err
	}

	return &core.MsgToggleCreditClassAllowlistResponse{}, nil
}