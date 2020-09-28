package types


// Money представляет собой денежную сумму в минимальных единицах (центы, дирамы, копейки и т.д)
type Money int64

// Выбор Category
type Category string

// Status статус платежа
type Status string

// Currency предаставляет код валюты
type Currency string

// Код валюты
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN представляет информацию о карте
type PAN string

// Card представляет информацию о карте
type Card struct {
	ID			int
	PAN			PAN
	Balance		Money
	Currency	Currency
	Color 		string
	Name		string
	Active 		bool
	MinBalance	Money

}


// PaymentSource func
type PaymentSources struct {
	Type string //'card'
	Number string // '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
	Active bool

}

const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInprogress Status = "INPROGRESS"
)


// Payment представляет информацию о платеже
type Payment struct {
	ID int
	Amount Money
	Category Category
	Status Status

}
