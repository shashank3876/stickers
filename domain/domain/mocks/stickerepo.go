package mocks

import (
	context "context"

	domain "shashank-gusain-backend-onboarding/domain"

	mock "github.com/stretchr/testify/mock"
)

type StickerRepository struct {
	mock.Mock
}

func (_m *StickerRepository) GetStickers(ctx context.Context) ([]domain.Stickers, error) {

	ret := _m.Called(ctx)

	return ret.Get(0).([]domain.Stickers), ret.Error(1)
}
