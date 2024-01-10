package repository

import (
	"context"

	"github.com/ogurilab/school-lunch-api/domain"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
)

type dishRepository struct {
	query db.Query
}

func NewDishRepository(query db.Query) domain.DishRepository {
	return &dishRepository{
		query: query,
	}
}

func (r *dishRepository) Create(ctx context.Context, dish *domain.Dish) error {
	arg := db.CreateDishParams{
		ID:     dish.ID,
		MenuID: dish.MenuID,
		Name:   dish.Name,
	}

	return r.query.CreateDish(ctx, arg)
}

func (r *dishRepository) GetByID(ctx context.Context, id string) (*domain.Dish, error) {

	result, err := r.query.GetDish(ctx, id)

	if err != nil {
		return nil, err
	}

	return domain.ReNewDish(
		result.ID,
		result.MenuID,
		result.Name,
	)

}

func (r *dishRepository) FetchByMenuID(ctx context.Context, menuID string) ([]*domain.Dish, error) {

	result, err := r.query.ListDishByMenuID(ctx, menuID)

	if err != nil {
		return nil, err
	}

	var dishes []*domain.Dish

	for _, dish := range result {
		d, err := domain.ReNewDish(
			dish.ID,
			dish.MenuID,
			dish.Name,
		)

		if err != nil {
			return nil, err
		}

		dishes = append(dishes, d)
	}

	return dishes, nil
}

func (r *dishRepository) FetchByName(ctx context.Context, search string, limit int32, offset int32) ([]*domain.Dish, error) {
	arg := db.ListDishByNameParams{
		Name:   search,
		Limit:  limit,
		Offset: offset,
	}

	result, err := r.query.ListDishByName(ctx, arg)

	if err != nil {
		return nil, err
	}

	var dishes []*domain.Dish

	for _, dish := range result {
		d, err := domain.ReNewDish(
			dish.ID,
			dish.MenuID,
			dish.Name,
		)

		if err != nil {
			return nil, err
		}

		dishes = append(dishes, d)
	}

	return dishes, nil
}

func (r *dishRepository) Fetch(ctx context.Context, limit int32, offset int32) ([]*domain.Dish, error) {
	arg := db.ListDishParams{
		Limit:  limit,
		Offset: offset,
	}

	result, err := r.query.ListDish(ctx, arg)

	if err != nil {
		return nil, err
	}

	var dishes []*domain.Dish

	for _, dish := range result {
		d, err := domain.ReNewDish(
			dish.ID,
			dish.MenuID,
			dish.Name,
		)

		if err != nil {
			return nil, err
		}

		dishes = append(dishes, d)
	}

	return dishes, nil
}
