package domain

import (
	"fmt"
)

type OrderNotification struct {
	OrderID 	string
	TraderID 	string
	Status 		string
	Amount 		float64
	Currency	string
	BankName    string  
	Phone  		string  
	CardNumber  string  
	Owner       string  
}

func (n OrderNotification) String() string {
	if n.Phone != "" {
		return fmt.Sprintf("%s\n%s\n%d %s %s / %s / %s", n.Status, n.OrderID, int(n.Amount), n.Currency, n.Phone, n.BankName, n.Owner)
	}
	return fmt.Sprintf("%s\n%s\n%d %s %s / %s / %s", n.Status, n.OrderID, int(n.Amount), n.Currency, n.CardNumber, n.BankName, n.Owner)
}