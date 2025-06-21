package domain

type AuthUsecase interface {
	Authorize(telegramID int64, token string) (string, error)
	GetTraderIDByTelegramID(telegramID int64) (string, error)
	GetTelegramIDsByTraderID(traderID string) ([]int64, error)
}