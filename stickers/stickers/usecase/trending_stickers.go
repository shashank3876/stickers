package usecase

import (
	"context"
	"fmt"
	"sort"
	"time"

	domain "shashank-gusain-backend-onboarding/domain"

	"github.com/spf13/viper"
)

type stickerUsecase struct {
	stickerRepo domain.StickerRepository
}

func NewStickerUsecase(s domain.StickerRepository) domain.StickerUsecase {

	return &stickerUsecase{

		stickerRepo: s,
	}
}

func (uc *stickerUsecase) GetTrendingStickers(c context.Context) ([]domain.Stickers, error) {
	sticker, err := uc.stickerRepo.GetStickers(c)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	returnLimit := viper.GetInt("returnLimit")

	var trendingStickers []domain.Stickers
	for _, stickers := range sticker {
		if stickers.AddedAt.After(time.Now().Add(-24 * time.Hour)) {
			trendingStickers = append(trendingStickers, stickers)
		}
	}

	sort.Slice(trendingStickers, func(i, j int) bool {

		return trendingStickers[i].Priority > trendingStickers[j].Priority
	})

	if len(trendingStickers) > returnLimit {
		trendingStickers = trendingStickers[:returnLimit]
	}

	return trendingStickers, nil
}
