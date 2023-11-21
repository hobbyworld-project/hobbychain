package types

import (
	"fmt"
	stdmath "math"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMinter returns a new Minter object with the given inflation and annual
// provisions values.
func NewMinter(inflation, annualProvisions sdk.Dec) Minter {
	return Minter{
		Inflation:        inflation,
		AnnualProvisions: annualProvisions,
	}
}

// InitialMinter returns an initial Minter object with a given inflation value.
func InitialMinter(inflation sdk.Dec) Minter {
	return NewMinter(
		inflation,
		math.LegacyNewDec(0),
	)
}

// DefaultInitialMinter returns a default initial Minter object for a new chain
// which uses an inflation rate of 13%.
func DefaultInitialMinter() Minter {
	return InitialMinter(
		sdk.NewDecWithPrec(13, 2),
	)
}

// ValidateMinter does a basic validation on minter.
func ValidateMinter(minter Minter) error {
	if minter.Inflation.IsNegative() {
		return fmt.Errorf("mint parameter Inflation should be positive, is %s",
			minter.Inflation.String())
	}
	return nil
}

// NextInflationRate returns the new inflation rate for the next hour.
func (m Minter) NextInflationRate(params Params, bondedRatio sdk.Dec) math.LegacyDec {
	// The target annual inflation rate is recalculated for each previsions cycle. The
	// inflation is also subject to a rate change (positive or negative) depending on
	// the distance from the desired ratio (67%). The maximum rate change possible is
	// defined to be 13% per year, however the annual inflation is capped as between
	// 7% and 20%.

	// (1 - bondedRatio/GoalBonded) * InflationRateChange
	inflationRateChangePerYear := math.LegacyOneDec().
		Sub(bondedRatio.Quo(params.GoalBonded)).
		Mul(params.InflationRateChange)
	inflationRateChange := inflationRateChangePerYear.Quo(math.LegacyNewDec(int64(params.BlocksPerYear)))

	// adjust the new annual inflation for this next cycle
	inflation := m.Inflation.Add(inflationRateChange) // note inflationRateChange may be negative
	if inflation.GT(params.InflationMax) {
		inflation = params.InflationMax
	}
	if inflation.LT(params.InflationMin) {
		inflation = params.InflationMin
	}

	return inflation
}

// NextAnnualProvisions returns the annual provisions based on current total
// supply and inflation rate.
func (m Minter) NextAnnualProvisions(_ Params, totalSupply math.Int) math.LegacyDec {
	return m.Inflation.MulInt(totalSupply)
}

// BlockProvision returns the provisions for a block based on the annual
// provisions rate.
func (m Minter) BlockProvision(params Params) sdk.Coin {
	provisionAmt := m.AnnualProvisions.QuoInt(sdk.NewInt(int64(params.BlocksPerYear)))
	return sdk.NewCoin(params.MintDenom, provisionAmt.TruncateInt())
}

// BlockReward returns the rewards for a block based on the reduction
func (m Minter) BlockProvisionReduction(height int64, params Params) sdk.Coin {
	var pos, last int
	var reward = math.LegacyNewDec(0)
	var prevHeight, curHeight uint64
	var reduction = params.Reduction
	var blockEpoch = uint64(height)
	var zero = math.LegacyNewDec(0).TruncateInt()
	if blockEpoch == 0 {
		return sdk.NewCoin(params.MintDenom, zero)
	}

	last = len(reduction.Heights) - 1

	// if the block epoch is greater than last reduction height and left is not 0, just return left amount as reward
	if blockEpoch > reduction.Heights[last] {
		return sdk.NewCoin(params.MintDenom, zero)
	}

	for i, h := range reduction.Heights {
		if blockEpoch <= h {
			pos = i
			curHeight = h
			if i > 0 {
				prevHeight = reduction.Heights[i-1]
			}
			break
		}
	}

	pow := int64(stdmath.Pow(2, float64(pos+1)))
	totalProvisions := reduction.TotalProvisions
	periodProvisions := totalProvisions.QuoInt64(pow)

	epochs := curHeight - prevHeight
	reward = periodProvisions.QuoInt64(int64(epochs))
	return sdk.NewCoin(params.MintDenom, reward.TruncateInt())
}
