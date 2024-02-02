package db

import (
	"context"

	"github.com/ogurilab/school-lunch-api/domain"
)

func (q *SQLQuery) CreateDishTx(ctx context.Context, dish *domain.Dish, menuID string) error {

	err := q.execTx(ctx, func(q *Queries) error {
		dishArgs := CreateDishParams{
			ID:   dish.ID,
			Name: dish.Name,
		}

		err := q.CreateDish(ctx, dishArgs)

		if err != nil {
			return err
		}

		menuDishArgs := CreateMenuDishParams{
			MenuID: menuID,
			DishID: dish.ID,
		}

		return q.CreateMenuDish(ctx, menuDishArgs)
	})

	return err
}
