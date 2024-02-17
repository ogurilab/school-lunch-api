package controller

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/errors"
	"github.com/ogurilab/school-lunch-api/server/validator"
	"github.com/ogurilab/school-lunch-api/util"
)

type adminController struct {
	mu domain.MenuUsecase
	du domain.DishUsecase
}

func NewAdminController(mu domain.MenuUsecase, du domain.DishUsecase) domain.AdminController {
	return &adminController{
		mu: mu,
		du: du,
	}
}

type createMenuRequest struct {
	OfferedAt                string `json:"offered_at" validate:"required,YYYY-MM-DD"`
	PhotoUrl                 string `json:"photo_url" validate:"omitempty,url"`
	ElementarySchoolCalories int32  `json:"elementary_school_calories" validate:"gt=0"`
	JuniorHighSchoolCalories int32  `json:"junior_high_school_calories" validate:"gt=0"`
	CityCode                 int32  `param:"code" validate:"required,gt=0"`
}

func (ac *adminController) CreateMenu(c echo.Context) error {
	var req createMenuRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(errors.NewBadRequestError(err))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	offeredAt, err := util.ParseDate(req.OfferedAt)

	if err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	menu, err := domain.NewMenu(
		offeredAt,
		sql.NullString{String: req.PhotoUrl, Valid: req.PhotoUrl != ""},
		req.ElementarySchoolCalories,
		req.JuniorHighSchoolCalories,
		req.CityCode,
	)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	if err := ac.mu.Create(ctx, menu); err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.NoContent(http.StatusCreated)
}

type createDishRequest struct {
	MenuID string `param:"id" validate:"required,ulid"`
	Name   string `json:"name" validate:"required"`
}

func (ac *adminController) CreateDish(c echo.Context) error {
	var req createDishRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	dish, err := domain.NewDish(req.Name)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	if err := ac.du.Create(ctx, dish, req.MenuID); err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.NoContent(http.StatusCreated)
}

type createDishesRequest struct {
	MenuID string           `param:"id" validate:"required,ulid"`
	Dishes []validator.Dish `json:"dishes" validate:"required,dishes"`
}

func (ac *adminController) CreateDishes(c echo.Context) error {
	var req createDishesRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	dishes := make([]*domain.Dish, 0, len(req.Dishes))

	for _, d := range req.Dishes {
		dish, err := domain.NewDish(d.Name)

		if err != nil {
			return c.JSON(errors.NewInternalServerError(err))
		}

		dishes = append(dishes, dish)
	}

	if err := ac.du.CreateMany(ctx, dishes, req.MenuID); err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.NoContent(http.StatusCreated)
}
