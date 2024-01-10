package domain

import (
	"context"
	"encoding/json"

	"github.com/ogurilab/school-lunch-api/util"
)

type Dish struct {
	ID     string `json:"id"`
	MenuID string `json:"menu_id"`
	Name   string `json:"name"`
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

func newDish(id string, menuID string, name string) (*Dish, error) {

	if _, err := util.ParseUlid(id); err != nil {
		return nil, err
	}

	return &Dish{
		ID:     id,
		MenuID: menuID,
		Name:   name,
	}, nil
}

func ReNewDish(id string, menuID string, name string) (*Dish, error) {

	return newDish(id, menuID, name)
}
func NewDish(menuID string, name string) (*Dish, error) {
	id := util.NewUlid()

	return newDish(id, menuID, name)
}

func NewDishesFromJson(
	input json.RawMessage,
) ([]*Dish, error) {
	var dishes []*Dish

	if err := json.Unmarshal(input, &dishes); err != nil {
		return nil, err
	}

	return dishes, nil
}
