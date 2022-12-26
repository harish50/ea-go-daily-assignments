package day2

type QuickCash struct {
	Accounts []Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	for _, acc := range qc.Accounts {
		if acc.CanWithDraw(amount) {
			acc.WithDraw(amount)
			return amount, acc.GetIdentifier()
		}
	}
	return amount, "NO_ACCOUNT"
}
