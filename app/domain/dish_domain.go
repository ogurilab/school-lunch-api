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

type DishRepository interface {
	Create(ctx context.Context, dish *Dish) error
	GetByID(ctx context.Context, id string) (*Dish, error)
	FetchByMenuID(ctx context.Context, menuID string) ([]*Dish, error)
	FetchByName(ctx context.Context, search string, limit int32, offset int32) ([]*Dish, error)
	Fetch(ctx context.Context, limit int32, offset int32) ([]*Dish, error)
}

type DishUsecase interface {
	Create(ctx context.Context, dish *Dish) error
	GetByID(ctx context.Context, id string) (*Dish, error)
	FetchByMenuID(ctx context.Context, menuID string) ([]*Dish, error)
	Fetch(ctx context.Context, search string, limit int32, offset int32) ([]*Dish, error)
}

type DishController interface {
	GetByID(c echo.Context) error
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
