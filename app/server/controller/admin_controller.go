package controller

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/errors"
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

	return c.JSON(200, menu)
}
