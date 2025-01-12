package http

import (
	"fmt"

	"net/http"

	domain "shashank-gusain-backend-onboarding/domain"

	errors "shashank-gusain-backend-onboarding/domain"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

type StickerHandler struct {
	SUsecase domain.StickerUsecase
}

func NewStickerHandler(e *echo.Echo, us domain.StickerUsecase) {
	handler := &StickerHandler{
		SUsecase: us,
	}

	e.GET("/v1/trendingStickers", handler.GetStickers)

}

func (s *StickerHandler) GetStickers(c echo.Context) error {
	
	ctx := c.Request().Context()

	stkr, err := s.SUsecase.GetTrendingStickers(ctx)

	fmt.Println(stkr)

	if err != nil {
		fmt.Println(stkr)
		return c.JSON(GetStatusCode(err), errors.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, stkr)
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case errors.ErrInternalServerError:
		return http.StatusInternalServerError
	case errors.ErrNotFound:
		return http.StatusNotFound
	case errors.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
