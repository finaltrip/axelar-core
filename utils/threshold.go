package utils

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// IsMet returns true if share > threshold * total
func (m Threshold) IsMet(share sdk.Int, total sdk.Int) bool {
	return share.MulRaw(m.Denominator).GT(total.MulRaw(m.Numerator))
}

// GT returns true if and only if threshold is greater than the given one
func (m Threshold) GT(t Threshold) bool {
	return sdk.NewInt(m.Numerator).MulRaw(t.Denominator).GT(sdk.NewInt(m.Denominator).MulRaw(t.Numerator))
}

// GTE returns true if and only if threshold is greater than or equal to the given one
func (m Threshold) GTE(t Threshold) bool {
	return sdk.NewInt(m.Numerator).MulRaw(t.Denominator).GTE(sdk.NewInt(m.Denominator).MulRaw(t.Numerator))
}

// LT returns true if and only if threshold is less than the given one
func (m Threshold) LT(t Threshold) bool {
	return sdk.NewInt(m.Numerator).MulRaw(t.Denominator).LT(sdk.NewInt(m.Denominator).MulRaw(t.Numerator))
}

// LTE returns true if and only if threshold is less than or equal to the given one
func (m Threshold) LTE(t Threshold) bool {
	return sdk.NewInt(m.Numerator).MulRaw(t.Denominator).LTE(sdk.NewInt(m.Denominator).MulRaw(t.Numerator))
}

// Validate returns an error if threshold is invalid
func (m Threshold) Validate() error {
	if m.Denominator == 0 {
		return fmt.Errorf("denominator must not be 0")
	}

	return nil
}
