package core

import (
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// Validate performs basic validation of the ClassSequence state type
func (m *ClassSequence) Validate() error {
	if err := ValidateCreditTypeAbbreviation(m.CreditTypeAbbrev); err != nil {
		return err // returns parse error
	}

	if m.NextSequence == 0 {
		return ecocredit.ErrParseFailure.Wrap("next sequence cannot be zero")
	}

	return nil
}