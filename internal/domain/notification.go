package domain

import "fmt"

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
		return fmt.Sprintf("Сделка %s\nСтатус: %s\nСумма: %f %s\nРеквизит: %s / %s / %s", n.OrderID, n.Status, n.Amount, n.Currency, n.Phone, n.BankName, n.Owner)
	}
	return fmt.Sprintf("Сделка %s\nСтатус: %s\nСумма: %f %s\nРеквизит: %s / %s / %s", n.OrderID, n.Status, n.Amount, n.Currency, n.CardNumber, n.BankName, n.Owner)
}