package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCashFromSavingsAccount(t *testing.T) {
	pa := &SavingsAccount{balance: 400}
	sa := &CreditCardAccount{balance: 500}

	qc := QuickCash{
		pa,
		sa,
	}

	amt, accType := qc.getCash(400)
	assert.Equal(t, float64(400), amt)
	assert.Equal(t, pa.GetIdentifier(), accType)
}

func TestGetCashFromSecondaryAccount(t *testing.T) {
	pa := &SavingsAccount{balance: 400}
	sa := &CreditCardAccount{balance: 500}

	qc := QuickCash{
		pa,
		sa,
	}

	amt, accType := qc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, sa.GetIdentifier(), accType)
}
