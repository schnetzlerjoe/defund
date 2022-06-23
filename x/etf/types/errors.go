package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/etf module sentinel errors
var (
	ErrWrongBroker          = sdkerrors.Register(ModuleName, 0, "Invalid broker")
	ErrFundNotFound         = sdkerrors.Register(ModuleName, 1, "Fund not found")
	ErrNextSequenceNotFound = sdkerrors.Register(ModuleName, 2, "Next Sequence not found")
	ErrInvalidDenom         = sdkerrors.Register(ModuleName, 3, "Denom is currenly not supported")
	ErrInvalidPool          = sdkerrors.Register(ModuleName, 4, "Pool is currenly not supported")
	ErrSymbolExists         = sdkerrors.Register(ModuleName, 5, "Symbol already exists")
	ErrWrongBaseDenom       = sdkerrors.Register(ModuleName, 6, "Invalid base denom")
	ErrPercentComp          = sdkerrors.Register(ModuleName, 7, "Invalid Percent Composition")
	ErrInvalidPools         = sdkerrors.Register(ModuleName, 8, "No pools found in store")
	ErrMarshallingError     = sdkerrors.Register(ModuleName, 9, "Marshalling error")
)
