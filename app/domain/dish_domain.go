package domain

import "github.com/ogurilab/school-lunch-api/util"

type Dish struct {
	ID     string `json:"id"`
	MenuID string `json:"menu_id"`
	Name   string `json:"name"`
}

type DishRepository interface {
	Create(dish *Dish) error
	GetByID(id string) (*Dish, error)
	GetByMenuID(menuID string) ([]*Dish, error)
	GetByNames(names []string) ([]*Dish, error)
}

type DishUsecase interface {
	Create(dish *Dish) error
	GetByID(id string) (*Dish, error)
	GetByMenuID(menuID string) ([]*Dish, error)
	GetByNames(names []string) ([]*Dish, error)
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
