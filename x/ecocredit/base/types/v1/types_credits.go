package v1

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/RegenNetwork/regen-ledger/types/v2/math"
	"github.com/RegenNetwork/regen-ledger/x/ecocredit/v3/base"
)

// Validate checks if Credits is valid.
func (c *Credits) Validate() error {
	if err := base.ValidateBatchDenom(c.BatchDenom); err != nil {
		return sdkerrors.ErrInvalidRequest.Wrapf("batch denom: %s", err)
	}

	if c.Amount == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("amount cannot be empty")
	}

	if _, err := math.NewPositiveDecFromString(c.Amount); err != nil {
		return err
	}

	return nil
}
