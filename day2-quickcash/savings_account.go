package day2_quickcash

type SavingsAccount struct {
	balance float64
}

func (sa *SavingsAccount) CanWithDraw(amount float64) bool {
	if sa.balance < amount {
		return false
	}
	return true
}

func (sa *SavingsAccount) WithDraw(amount float64) {
	sa.balance = sa.balance - amount
}

func (sa *SavingsAccount) GetIdentifier() string {
	return "SAVINGS_ACCOUNT"
}
