package http_test

import (
	"net/http"
	"net/http/httptest"

	"testing"
	"time"

	domain "shashank-gusain-backend-onboarding/domain"

	"shashank-gusain-backend-onboarding/domain/mocks"
	https "shashank-gusain-backend-onboarding/stickers/delivery"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetStickersHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/trendingStickers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockStickerUsecase := &mocks.StickerUsecase{}
	mockStickerUsecase.On("GetTrendingStickers", mock.Anything).Return([]domain.Stickers{
		{
			ID: 1, Name: "Sticker1", Priority: 1, AddedAt: time.Time{},
		},
	}, nil)

	handler := &https.StickerHandler{
		SUsecase: mockStickerUsecase,
	}

	if assert.NoError(t, handler.GetStickers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Sticker1")
	}
}
