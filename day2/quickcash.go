package day2

type QuickCash struct {
	PrimaryAccount   Withdrawable
	SecondaryAccount Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	if qc.PrimaryAccount.CanWithDraw(amount) {
		qc.PrimaryAccount.WithDraw(amount)
		return amount, qc.PrimaryAccount.GetIdentifier()
	}
	if qc.SecondaryAccount.CanWithDraw(amount) {
		qc.SecondaryAccount.WithDraw(amount)
		return amount, qc.SecondaryAccount.GetIdentifier()
	}
	return amount, "NO_ACCOUNT"
}
