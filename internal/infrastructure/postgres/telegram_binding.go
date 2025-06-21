package postgres

type TelegramBinding struct {
	ID 		   uint  `gorm:"primaryKey"`
	TelegramID int64 `gorm:"uniqueIndex"`
	TraderID   string
}