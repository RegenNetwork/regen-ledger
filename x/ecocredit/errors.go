package ecocredit

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrParseFailure        = sdkerrors.Register(ModuleName, 2, "parse error")
)
