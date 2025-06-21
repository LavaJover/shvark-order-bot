package postgres

import "gorm.io/gorm"

type DefaultAuthRepository struct {
	db *gorm.DB
}

func NewDefaultAuthRepository(db *gorm.DB) *DefaultAuthRepository {
	return &DefaultAuthRepository{
		db: db,
	}
}

func (r *DefaultAuthRepository)SaveMapping (telegramID int64, traderID string) error {
	bindingModel := TelegramBinding{
		TelegramID: telegramID,
		TraderID: traderID,
	}
	if err := r.db.Save(&bindingModel).Error; err != nil {
		return err
	}
	return nil
}

func (r *DefaultAuthRepository) GetTraderID (telegramID int64) (string, error) {
	var bindingModel TelegramBinding
	if err := r.db.Model(&TelegramBinding{}).Where("telegram_id = ?", telegramID).First(&bindingModel).Error; err != nil {
		return "", err
	}
	return bindingModel.TraderID, nil
}

func (r *DefaultAuthRepository) GetTelegramIDsByTraderID(traderID string) ([]int64, error) {
	var bindings []TelegramBinding
	if err := r.db.Model(&TelegramBinding{}).Where("trader_id = ?", traderID).Find(&bindings).Error; err != nil {
		return nil, err
	}
	telegramIds := make([]int64, len(bindings))
	for i, binding := range bindings {
		telegramIds[i] = binding.TelegramID
	}
	return telegramIds, nil
}