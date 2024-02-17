package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/errors"
)

type dishController struct {
	du domain.DishUsecase
}

func NewDishController(du domain.DishUsecase) domain.DishController {
	return &dishController{
		du: du,
	}
}

type fetchDishByMenuIDRequest struct {
	MenuID string `param:"menuID" validate:"required,ulid"`
}

func (dc *dishController) FetchByMenuID(c echo.Context) error {
	var req fetchDishByMenuIDRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	dishes, err := dc.du.FetchByMenuID(ctx, req.MenuID)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, dishes)
}

type getDishRequest struct {
	ID     string `param:"id" validate:"required,ulid"`
	Limit  int32  `query:"limit" validate:"gt=0"`
	Offset int32  `query:"offset" validate:"gte=0"`
}

func (dc *dishController) GetByID(c echo.Context) error {
	var req getDishRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if req.Limit > domain.MAX_LIMIT {
		return c.JSON(errors.NewMaxLimitError())
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

	dish, err := dc.du.GetByID(ctx, req.ID, req.Limit, req.Offset)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, dish)
}

type getDishInCityRequest struct {
	ID       string `param:"id" validate:"required,ulid"`
	Limit    int32  `query:"limit" validate:"gt=0"`
	Offset   int32  `query:"offset" validate:"gte=0"`
	CityCode int32  `param:"code" validate:"required,gte=0"`
}

func (dc *dishController) GetByIdInCity(c echo.Context) error {
	var req getDishInCityRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if req.Limit > domain.MAX_LIMIT {
		return c.JSON(errors.NewMaxLimitError())
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

	dish, err := dc.du.GetByIdInCity(ctx, req.ID, req.Limit, req.Offset, req.CityCode)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, dish)
}

type fetchDishRequest struct {
	Limit  int32  `query:"limit" validate:"gt=0"`
	Offset int32  `query:"offset" validate:"gte=0"`
	Search string `query:"search"`
}

func (dc *dishController) Fetch(c echo.Context) error {
	var req fetchDishRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if req.Limit > domain.MAX_LIMIT {
		return c.JSON(errors.NewMaxLimitError())
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

	dishes, err := dc.du.Fetch(ctx, req.Search, req.Limit, req.Offset)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, dishes)
}
