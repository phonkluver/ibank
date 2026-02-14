package types

type Money int64

// Валюта денег
type Currency string

// Код валюты
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

type PAN string

type Card struct {
	id       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    string
	Name     string
	Active   bool
}

type Payment struct {
	id     int
	Amount Money
}

type Category string
