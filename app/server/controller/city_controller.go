package controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/errors"
)

type cityController struct {
	cityUsecase domain.CityUsecase
}

func NewCityController(cu domain.CityUsecase) domain.CityController {
	return &cityController{
		cityUsecase: cu,
	}
}

type getCityRequest struct {
	CityCode int32 `param:"code" validate:"required,gt=0"`
}

func (cc *cityController) GetByCityCode(c echo.Context) error {
	var req getCityRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(errors.NewBadRequestError(err))
	}

	if err := c.Validate(&req); err != nil {

		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	city, err := cc.cityUsecase.GetByCityCode(ctx, req.CityCode)

	if err != nil {

		if err == sql.ErrNoRows {
			return c.JSON(errors.NewNotFoundError(err))
		}

		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, city)
}

type fetchCityRequest struct {
	Search string `query:"search" validate:"omitempty"`
	Limit  int32  `query:"limit" validate:"gt=0"`
	Offset int32  `query:"offset" validate:"gte=0"`
}

func (cc *cityController) Fetch(c echo.Context) error {
	var req fetchCityRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if req.Limit == 0 {
		req.Limit = domain.DEFAULT_LIMIT
	}

	if req.Offset == 0 {
		req.Offset = domain.DEFAULT_OFFSET
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	cities, err := cc.cityUsecase.Fetch(ctx, req.Limit, req.Offset, req.Search)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, cities)

}

type fetchCityByPrefectureCodeRequest struct {
	PrefectureCode int32 `param:"code" validate:"required,gt=0"`
	Limit          int32 `query:"limit" validate:"gt=0"`
	Offset         int32 `query:"offset" validate:"gte=0"`
}

func (cc *cityController) FetchByPrefectureCode(c echo.Context) error {
	var req fetchCityByPrefectureCodeRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if req.Limit == 0 {
		req.Limit = domain.DEFAULT_LIMIT
	}

	if req.Offset == 0 {
		req.Offset = domain.DEFAULT_OFFSET
	}

	if err := c.Validate(&req); err != nil {

		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	cities, err := cc.cityUsecase.FetchByPrefectureCode(ctx, req.Limit, req.Offset, req.PrefectureCode)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, cities)

}
