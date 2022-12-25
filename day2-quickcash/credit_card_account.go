package day2_quickcash

type CreditCardAccount struct {
	balance float64
}

func (cca *CreditCardAccount) CanWithDraw(amount float64) bool {
	return amount <= cca.balance
}

func (cca *CreditCardAccount) WithDraw(amount float64) {
	cca.balance = cca.balance - amount
}

func (cca *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD_ACCOUNT"
}
