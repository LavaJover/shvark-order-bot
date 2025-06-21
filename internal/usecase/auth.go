package usecase

import (
	"fmt"

	"github.com/LavaJover/shvark-order-bot/internal/domain"
	"github.com/LavaJover/shvark-order-bot/internal/grpcapi"
)

type AuthUsecase struct {
	authRepo domain.AuthRepository
	ssoClient *grpcapi.SSOClient
}

func NewAuthUsecase(repo domain.AuthRepository, ssoClient *grpcapi.SSOClient) *AuthUsecase {
	return &AuthUsecase{authRepo: repo, ssoClient: ssoClient}
}

func (uc *AuthUsecase) Authorize(telegramID int64, token string) (string, error) {
	// check if given token is valid
	valid, traderID, err := uc.ssoClient.ValidateToken(token)
	if err != nil || !valid{
		return "", fmt.Errorf("invalid token")
	}
	// if token is valid => we can map telegramID and traderID
	err = uc.authRepo.SaveMapping(telegramID, traderID)
	if err != nil {
		return "", err
	}
	return traderID, nil
}

func (uc *AuthUsecase) GetTraderIDByTelegramID(telegramID int64) (string, error) {
	return uc.authRepo.GetTraderID(telegramID)
}

func (uc *AuthUsecase) GetTelegramIDsByTraderID(traderID string) ([]int64, error) {
	return uc.authRepo.GetTelegramIDsByTraderID(traderID)
}