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

func (r *dishRepository) Create(ctx context.Context, dish *domain.Dish, menuID string) error {

	return r.query.CreateDishTx(ctx, dish, menuID)
}

func (r *dishRepository) GetByID(ctx context.Context, id string) (*domain.Dish, error) {

	result, err := r.query.GetDish(ctx, id)

	if err != nil {
		return nil, err
	}

	return domain.ReNewDish(
		result.ID,
		result.Name,
	)
}

func (r *dishRepository) FetchByMenuID(ctx context.Context, menuID string) ([]*domain.Dish, error) {

	results, err := r.query.ListDishByMenuID(ctx, menuID)

	if err != nil {
		return nil, err
	}

	dishes := make([]*domain.Dish, 0, len(results))

	for _, result := range results {
		dish, err := domain.ReNewDish(
			result.ID,
			result.Name,
		)

		if err != nil {
			return nil, err
		}

		dishes = append(dishes, dish)
	}

	return dishes, nil
}

func (r *dishRepository) FetchByName(ctx context.Context, search string, limit int32, offset int32) ([]*domain.Dish, error) {
	arg := db.ListDishByNameParams{
		Name:   search,
		Limit:  limit,
		Offset: offset,
	}

	results, err := r.query.ListDishByName(ctx, arg)

	if err != nil {
		return nil, err
	}

	dishes := make([]*domain.Dish, 0, len(results))

	for _, result := range results {
		dish, err := domain.ReNewDish(
			result.ID,
			result.Name,
		)

		if err != nil {
			return nil, err
		}

		dishes = append(dishes, dish)
	}

	return dishes, nil
}

func (r *dishRepository) Fetch(ctx context.Context, limit int32, offset int32) ([]*domain.Dish, error) {
	arg := db.ListDishParams{
		Limit:  limit,
		Offset: offset,
	}

	results, err := r.query.ListDish(ctx, arg)

	if err != nil {
		return nil, err
	}

	dishes := make([]*domain.Dish, 0, len(results))

	for _, result := range results {
		dish, err := domain.ReNewDish(
			result.ID,
			result.Name,
		)

		if err != nil {
			return nil, err
		}

		dishes = append(dishes, dish)
	}

	return dishes, nil
}
