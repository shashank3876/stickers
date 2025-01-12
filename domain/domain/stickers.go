package mock

import (
	"context"
	"time"
)

type Stickers struct {
	ID       int       `json:"id"`
	Name     string    `json:"name" validate:"required"`
	Priority int       `json:"priority" validate:"required"`
	AddedAt  time.Time `json:"added_at"`
}

type StickerUsecase interface {
	GetTrendingStickers(ctx context.Context) ([]Stickers, error)
}

type StickerRepository interface {
	GetStickers(ctx context.Context) ([]Stickers, error)
}
