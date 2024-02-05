package domain

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/util"
)

type Dish struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DishWithMenuIDs struct {
	Dish
	MenuIDs []string `json:"menu_ids"`
}

type DishRepository interface {
	Create(ctx context.Context, dish *Dish, menuID string) error
	GetByID(ctx context.Context, id string, limit int32, offset int32) (*DishWithMenuIDs, error)
	GetByIdInCity(ctx context.Context, id string, limit int32, offset int32, city int32) (*DishWithMenuIDs, error)
	FetchByMenuID(ctx context.Context, menuID string) ([]*Dish, error)
	FetchByName(ctx context.Context, search string, limit int32, offset int32) ([]*Dish, error)
	Fetch(ctx context.Context, limit int32, offset int32) ([]*Dish, error)
}

type DishUsecase interface {
	Create(ctx context.Context, dish *Dish, menuID string) error
	GetByID(ctx context.Context, id string, limit int32, offset int32) (*DishWithMenuIDs, error)
	GetByIdInCity(ctx context.Context, id string, limit int32, offset int32, city int32) (*DishWithMenuIDs, error)
	FetchByMenuID(ctx context.Context, menuID string) ([]*Dish, error)
	Fetch(ctx context.Context, search string, limit int32, offset int32) ([]*Dish, error)
}

type DishController interface {
	GetByID(c echo.Context) error
	GetByIdInCity(c echo.Context) error
	FetchByMenuID(c echo.Context) error
	Fetch(c echo.Context) error
}

func newDish(id string, name string) (*Dish, error) {

	if _, err := util.ParseUlid(id); err != nil {
		return nil, err
	}

	return &Dish{
		ID:   id,
		Name: name,
	}, nil
}

func ReNewDish(id string, name string) (*Dish, error) {

	return newDish(id, name)
}

func NewDish(name string) (*Dish, error) {
	id := util.NewUlid()

	return newDish(id, name)
}

func ReNewDishWithMenuIDs(id string, name string, menuIDs []string) (*DishWithMenuIDs, error) {

	dish, err := newDish(id, name)

	if err != nil {
		return nil, err
	}

	return &DishWithMenuIDs{
		Dish:    *dish,
		MenuIDs: menuIDs,
	}, nil
}
