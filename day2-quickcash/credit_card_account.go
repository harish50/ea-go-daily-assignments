package day2_quickcash

type CreditCardAccount struct {
	balance float64
}

func (sa *CreditCardAccount) CanWithDraw(amount float64) bool {
	return false
}

func (sa *CreditCardAccount) WithDraw(amount float64) {

}

func (sa *CreditCardAccount) GetIdentifier() string {
	return ""
}
