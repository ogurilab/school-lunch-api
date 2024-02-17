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

func (r *dishRepository) CreateMany(ctx context.Context, dishes []*domain.Dish, menuID string) error {
	return r.query.CreateDishesTx(ctx, dishes, menuID)
}

func (r *dishRepository) GetByID(ctx context.Context, id string, limit int32, offset int32) (*domain.DishWithMenuIDs, error) {

	arg := db.GetDishParams{
		ID:     id,
		Limit:  limit,
		Offset: offset,
	}

	results, err := r.query.GetDish(ctx, arg)

	if err != nil {
		return nil, err
	}

	menuIDs := make([]string, 0, len(results))

	for _, result := range results {
		menuIDs = append(menuIDs, result.MenuID)
	}

	firstResult := results[0]

	return domain.ReNewDishWithMenuIDs(
		firstResult.ID,
		firstResult.Name,
		menuIDs,
	)
}

func (r *dishRepository) GetByIdInCity(ctx context.Context, id string, limit int32, offset int32, city int32) (*domain.DishWithMenuIDs, error) {

	arg := db.GetDishInCityParams{
		ID:       id,
		Limit:    limit,
		Offset:   offset,
		CityCode: city,
	}

	results, err := r.query.GetDishInCity(ctx, arg)

	if err != nil {
		return nil, err
	}

	menuIDs := make([]string, 0, len(results))

	for _, result := range results {
		menuIDs = append(menuIDs, result.MenuID)
	}

	firstResult := results[0]

	return domain.ReNewDishWithMenuIDs(
		firstResult.ID,
		firstResult.Name,
		menuIDs,
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
