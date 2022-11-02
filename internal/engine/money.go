package engine

import (
	"fmt"
	"math"
)

// Money represents money object
type Money struct {
	Amount   uint64
	Currency string
}

func (m Money) GetAmount() uint64 {
	return m.Amount
}

func (m Money) LessThan(money Money) bool {
	return m.Amount < money.Amount
}

func (m Money) LargerThan(money Money) bool {
	return m.Amount > money.Amount
}

func (m Money) GetPercentageDifference(money Money) float64 {
	return math.Abs(float64(money.Amount)-float64(m.Amount)) / float64(m.Amount) * 100

}

func (m Money) Sub(money Money) (Money, error) {
	if m.Amount < money.Amount {
		return Money{}, fmt.Errorf("cannot subtract larger money value")
	}
	return Money{
		Amount:   m.Amount - money.Amount,
		Currency: m.Currency,
	}, nil
}
