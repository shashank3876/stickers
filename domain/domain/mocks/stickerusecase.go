package mocks

import (
	context "context"

	domain "shashank-gusain-backend-onboarding/domain"

	mock "github.com/stretchr/testify/mock"
)

type StickerUsecase struct {
	mock.Mock
}

func (_m *StickerUsecase) GetTrendingStickers(ctx context.Context) ([]domain.Stickers, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Stickers
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Stickers); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Stickers)
		}
	}

	return r0, nil
}

func (_m *StickerUsecase) GetStickers(ctx context.Context) ([]domain.Stickers, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Stickers
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Stickers); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Stickers)
		}
	}

	return r0, nil
}
