package day2_quickcash

type Withdrawable interface {
	CanWithDraw(amount float64) bool
	WithDraw(amount float64)
	GetIdentifier() string
}
