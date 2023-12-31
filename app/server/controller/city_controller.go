package controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/errors"
)

type CityController struct {
	CityUsecase domain.CityUsecase
}

func NewCityController(cu domain.CityUsecase) domain.CityController {
	return &CityController{
		CityUsecase: cu,
	}
}

type getCityRequest struct {
	CityCode int32 `param:"code" validate:"required,gt=0"`
}

func (cc *CityController) GetByCityCode(c echo.Context) error {
	var req getCityRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(errors.NewBadRequestError(err))
	}

	if err := c.Validate(&req); err != nil {

		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	city, err := cc.CityUsecase.GetByCityCode(ctx, req.CityCode)

	if err != nil {

		if err == sql.ErrNoRows {
			return c.JSON(errors.NewNotFoundError(err))
		}

		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, domain.NewResponse(city))
}

type fetchCityRequest struct {
	Search string `query:"search" validate:"omitempty"`
	Limit  int32  `query:"limit" validate:"gt=0"`
	Offset int32  `query:"offset" validate:"gte=0"`
}

func (cc *CityController) Fetch(c echo.Context) error {
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

	cities, err := cc.CityUsecase.Fetch(ctx, req.Limit, req.Offset, req.Search)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, domain.NewResponse(cities))

}

type fetchCityByPrefectureCodeRequest struct {
	PrefectureCode int32 `param:"code" validate:"required,gt=0"`
	Limit          int32 `query:"limit" validate:"gt=0"`
	Offset         int32 `query:"offset" validate:"gte=0"`
}

func (cc *CityController) FetchByPrefectureCode(c echo.Context) error {
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

	cities, err := cc.CityUsecase.FetchByPrefectureCode(ctx, req.Limit, req.Offset, req.PrefectureCode)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, domain.NewResponse(cities))

}
