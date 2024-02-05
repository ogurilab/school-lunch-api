package repository

import (
	"context"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
)

type menuRepository struct {
	query db.Query
}

func NewMenuRepository(query db.Query) domain.MenuRepository {
	return &menuRepository{
		query: query,
	}
}

func (r *menuRepository) Create(ctx context.Context, menu *domain.Menu) error {
	arg := db.CreateMenuParams{
		ID:                       menu.ID,
		OfferedAt:                menu.OfferedAt,
		CityCode:                 menu.CityCode,
		PhotoUrl:                 menu.PhotoUrl,
		ElementarySchoolCalories: menu.ElementarySchoolCalories,
		JuniorHighSchoolCalories: menu.JuniorHighSchoolCalories,
	}

	return r.query.CreateMenu(ctx, arg)
}

func (r *menuRepository) GetByID(ctx context.Context, id string, city int32) (*domain.Menu, error) {
	arg := db.GetMenuParams{
		ID:       id,
		CityCode: city,
	}

	result, err := r.query.GetMenu(ctx, arg)

	if err != nil {
		return nil, err
	}

	menu, err := domain.ReNewMenu(
		result.ID,
		result.OfferedAt,
		result.PhotoUrl,
		result.ElementarySchoolCalories,
		result.JuniorHighSchoolCalories,
		result.CityCode,
	)

	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (r *menuRepository) FetchByCity(ctx context.Context, limit int32, offset int32, offered time.Time, city int32) ([]*domain.Menu, error) {
	arg := db.ListMenuByCityParams{
		Limit:     limit,
		Offset:    offset,
		OfferedAt: offered,
		CityCode:  city,
	}

	results, err := r.query.ListMenuByCity(ctx, arg)

	if err != nil {
		return nil, err
	}

	menus := make([]*domain.Menu, 0, len(results))

	for _, result := range results {
		menu, err := domain.ReNewMenu(
			result.ID,
			result.OfferedAt,
			result.PhotoUrl,
			result.ElementarySchoolCalories,
			result.JuniorHighSchoolCalories,
			result.CityCode,
		)

		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}

func (r *menuRepository) Fetch(ctx context.Context, limit int32, offset int32, offered time.Time) ([]*domain.Menu, error) {
	arg := db.ListMenuParams{
		Limit:     limit,
		Offset:    offset,
		OfferedAt: offered,
	}

	results, err := r.query.ListMenu(ctx, arg)

	if err != nil {
		return nil, err
	}

	menus := make([]*domain.Menu, 0, len(results))

	for _, result := range results {
		menu, err := domain.ReNewMenu(
			result.ID,
			result.OfferedAt,
			result.PhotoUrl,
			result.ElementarySchoolCalories,
			result.JuniorHighSchoolCalories,
			result.CityCode,
		)

		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}

func (r *menuRepository) FetchByIDs(ctx context.Context, limit int32, offset int32, offered time.Time, ids []string) ([]*domain.Menu, error) {
	arg := db.ListMenuByIDsParams{
		Limit:     limit,
		Offset:    offset,
		OfferedAt: offered,
		IDs:       ids,
	}

	results, err := r.query.ListMenuByIDs(ctx, arg)

	if err != nil {
		return nil, err
	}

	menus := make([]*domain.Menu, 0, len(results))

	for _, result := range results {
		menu, err := domain.ReNewMenu(
			result.ID,
			result.OfferedAt,
			result.PhotoUrl,
			result.ElementarySchoolCalories,
			result.JuniorHighSchoolCalories,
			result.CityCode,
		)

		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}
