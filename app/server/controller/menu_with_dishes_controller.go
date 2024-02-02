package controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/errors"
	"github.com/ogurilab/school-lunch-api/util"
)

/************************
 * MenuWithDishesController
 ************************/

type menuWithDishesController struct {
	mu domain.MenuWithDishesUsecase
}

func NewMenuWithDishesController(mu domain.MenuWithDishesUsecase) domain.MenuWithDishesController {
	return &menuWithDishesController{
		mu: mu,
	}
}

type getMenuWithDishesRequest struct {
	ID       string `param:"id" validate:"required,ulid"`
	CityCode int32  `param:"code" validate:"required,gt=0"`
}

func (mc *menuWithDishesController) GetByID(c echo.Context) error {
	var req getMenuWithDishesRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	menu, err := mc.mu.GetByID(ctx, req.ID, req.CityCode)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(errors.NewNotFoundError(err))
		}

		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, menu)
}

type fetchMenuWithDishesByCityRequest struct {
	CityCode int32  `param:"code" validate:"required,gt=0"`
	Limit    int32  `query:"limit" validate:"gt=0"`
	Offset   int32  `query:"offset" validate:"gte=0"`
	Offered  string `query:"offered" validate:"YYYY-MM-DD,required"`
}

type fetchMenuWithDishesResponse struct {
	Menus []*domain.MenuWithDishes `json:"menus"`
	Next  string                   `json:"next"`
}

func (mc *menuWithDishesController) FetchByCity(c echo.Context) error {
	var req fetchMenuWithDishesByCityRequest

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

	parsedDate, err := util.ParseDate(req.Offered)

	if err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	menus, err := mc.mu.FetchByCity(
		ctx,
		req.Limit,
		req.Offset,
		parsedDate,
		req.CityCode,
	)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	if len(menus) == 0 {
		return c.JSON(200, fetchMenuResponse{Menus: []*domain.Menu{}, Next: ""})
	}

	next := util.FormatDate(menus[len(menus)-1].OfferedAt)

	res := fetchMenuWithDishesResponse{
		Menus: menus,
		Next:  next,
	}

	return c.JSON(200, res)
}

type fetchMenuWithDishesRequest struct {
	Limit   int32  `query:"limit" validate:"gt=0"`
	Offset  int32  `query:"offset" validate:"gte=0"`
	Offered string `query:"offered" validate:"YYYY-MM-DD,required"`
}

func (mc *menuWithDishesController) Fetch(c echo.Context) error {
	var req fetchMenuWithDishesRequest

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

	parsedDate, err := util.ParseDate(req.Offered)

	if err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	menus, err := mc.mu.Fetch(
		ctx,
		req.Limit,
		req.Offset,
		parsedDate,
	)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	if len(menus) == 0 {
		return c.JSON(200, fetchMenuResponse{Menus: []*domain.Menu{}, Next: ""})
	}

	next := util.FormatDate(menus[len(menus)-1].OfferedAt)

	res := fetchMenuWithDishesResponse{
		Menus: menus,
		Next:  next,
	}
	return c.JSON(200, res)
}