package domain

import "fmt"

type OrderNotification struct {
	OrderID 	string
	TraderID 	string
	Status 		string
	Amount 		float64
	Currency	string
}

func (n OrderNotification) String() string {
	return fmt.Sprintf("Сделка %s\nСтатус: %s\nСумма: %f %s", n.OrderID, n.Status, n.Amount, n.Currency)
}