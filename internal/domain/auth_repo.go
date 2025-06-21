package domain

type AuthRepository interface {
	SaveMapping(telegramID int64, traderID string) error
	GetTraderID(telegramID int64) (string, error)
	GetTelegramIDsByTraderID(traderID string) ([]int64, error)
}