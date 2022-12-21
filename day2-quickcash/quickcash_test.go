package day2_quickcash

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
