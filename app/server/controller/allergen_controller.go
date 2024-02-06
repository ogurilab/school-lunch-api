package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/errors"
)

type allergenController struct {
	au domain.AllergenUsecase
}

func NewAllergenController(au domain.AllergenUsecase) domain.AllergenController {
	return &allergenController{
		au: au,
	}
}

type fetchAllergenByDishIDRequest struct {
	DishID string `param:"id" validate:"required,ulid"`
}

func (ac *allergenController) FetchByDishID(c echo.Context) error {
	var req fetchAllergenByDishIDRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	allergens, err := ac.au.FetchByDishID(ctx, req.DishID)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, allergens)
}

type fetchAllergenByMenuIDRequest struct {
	MenuID string `param:"id" validate:"required,ulid"`
}

func (ac *allergenController) FetchByMenuID(c echo.Context) error {
	var req fetchAllergenByMenuIDRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(errors.NewBadRequestError(err))
	}

	ctx := c.Request().Context()

	allergens, err := ac.au.FetchByMenuID(ctx, req.MenuID)

	if err != nil {
		return c.JSON(errors.NewInternalServerError(err))
	}

	return c.JSON(200, allergens)
}
