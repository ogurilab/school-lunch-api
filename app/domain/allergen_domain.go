package domain

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Allergen struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Category int32  `json:"category"`
}

type AllergenRepository interface {
	FetchByDishID(ctx context.Context, dishID string) ([]*Allergen, error)
	FetchInDish(ctx context.Context, dishIDs []string) ([]*Allergen, error)
}

type AllergenUsecase interface {
	FetchByDishID(ctx context.Context, dishID string) ([]*Allergen, error)
	FetchByMenuID(ctx context.Context, menuID string) ([]*Allergen, error)
}

type AllergenController interface {
	FetchByDishID(c echo.Context) error
	FetchByMenuID(c echo.Context) error
}

func newAllergen(id int32, name string, category int32) *Allergen {
	return &Allergen{
		ID:       id,
		Name:     name,
		Category: category,
	}
}

func ReNewAllergen(id int32, name string, category int32) *Allergen {
	return newAllergen(id, name, category)
}
