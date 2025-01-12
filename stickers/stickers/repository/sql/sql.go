package sql

import (
	"context"

	domain "shashank-gusain-backend-onboarding/domain"

	"gorm.io/gorm"
)

type mysqlStickerRepository struct {
	Conn *gorm.DB
}

func NewSQLStickerRepository(conn *gorm.DB) domain.StickerRepository {

	return &mysqlStickerRepository{conn}
}

func (repo *mysqlStickerRepository) GetStickers(ctx context.Context) ([]domain.Stickers, error) {

	var stickers []domain.Stickers
	err := repo.Conn.Find(&stickers).Error
	if err != nil {

		return nil, err
	}
	return stickers, nil
}
