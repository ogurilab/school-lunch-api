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

	var menus []*domain.Menu

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

/********************
 * MenuWithDishesRepository
 ********************/

type menuWithDishesRepository struct {
	query db.Query
}

func NewMenuWithDishesRepository(query db.Query) domain.MenuWithDishesRepository {
	return &menuWithDishesRepository{
		query: query,
	}
}

func (r *menuWithDishesRepository) GetByID(ctx context.Context, id string, city int32) (*domain.MenuWithDishes, error) {
	arg := db.GetMenuWithDishesParams{
		ID:       id,
		CityCode: city,
	}

	result, err := r.query.GetMenuWithDishes(ctx, arg)

	if err != nil {
		return nil, err
	}

	dishes, err := domain.NewDishesFromJson(result.Dishes)

	if err != nil {
		return nil, err
	}

	menu, err := domain.ReNewMenuWithDishes(
		result.ID,
		result.OfferedAt,
		result.PhotoUrl,
		result.ElementarySchoolCalories,
		result.JuniorHighSchoolCalories,
		result.CityCode,
		dishes,
	)

	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (r *menuWithDishesRepository) FetchByCity(ctx context.Context, limit int32, offset int32, offered time.Time, city int32) ([]*domain.MenuWithDishes, error) {
	arg := db.ListMenuWithDishesByCityParams{
		Limit:     limit,
		Offset:    offset,
		OfferedAt: offered,
		CityCode:  city,
	}

	results, err := r.query.ListMenuWithDishesByCity(ctx, arg)

	if err != nil {
		return nil, err
	}

	var menus []*domain.MenuWithDishes

	for _, result := range results {
		dishes, err := domain.NewDishesFromJson(result.Dishes)

		if err != nil {
			return nil, err
		}

		menu, err := domain.ReNewMenuWithDishes(
			result.ID,

			result.OfferedAt,
			result.PhotoUrl,
			result.ElementarySchoolCalories,
			result.JuniorHighSchoolCalories,
			result.CityCode,
			dishes,
		)

		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}
func (r *menuWithDishesRepository) Fetch(ctx context.Context, limit int32, offset int32, offered time.Time) ([]*domain.MenuWithDishes, error) {
	arg := db.ListMenuWithDishesParams{
		Limit:     limit,
		Offset:    offset,
		OfferedAt: offered,
	}

	results, err := r.query.ListMenuWithDishes(ctx, arg)

	if err != nil {
		return nil, err
	}

	var menus []*domain.MenuWithDishes

	for _, result := range results {
		dishes, err := domain.NewDishesFromJson(result.Dishes)

		if err != nil {
			return nil, err
		}

		menu, err := domain.ReNewMenuWithDishes(
			result.ID,

			result.OfferedAt,
			result.PhotoUrl,
			result.ElementarySchoolCalories,
			result.JuniorHighSchoolCalories,
			result.CityCode,
			dishes,
		)

		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}
