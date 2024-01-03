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

func (r *menuRepository) Fetch(ctx context.Context, limit int32, offset int32, city int32) ([]*domain.Menu, error) {
	arg := db.ListMenusParams{
		Limit:    limit,
		Offset:   offset,
		CityCode: city,
	}

	results, err := r.query.ListMenus(ctx, arg)

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

func (r *menuRepository) GetByDate(ctx context.Context, offeredAt time.Time, city int32) (*domain.Menu, error) {
	arg := db.GetMenuByOfferedAtParams{
		OfferedAt: offeredAt,
		CityCode:  city,
	}

	result, err := r.query.GetMenuByOfferedAt(ctx, arg)

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

func (r *menuRepository) FetchByRangeDate(ctx context.Context, start, end time.Time, city int32, limit int32) ([]*domain.Menu, error) {
	arg := db.ListMenusByOfferedAtParams{
		StartOfferedAt: start,
		EndOfferedAt:   end,
		CityCode:       city,
		Limit:          limit,
	}

	results, err := r.query.ListMenusByOfferedAt(ctx, arg)

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

func (r *menuRepository) GetByIDWithDishes(ctx context.Context, id string, city int32) (*domain.MenuWithDishes, error) {
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

func (r *menuRepository) FetchWithDishes(ctx context.Context, limit int32, offset int32, city int32) ([]*domain.MenuWithDishes, error) {
	arg := db.ListMenuWithDishesParams{
		Limit:    limit,
		Offset:   offset,
		CityCode: city,
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

func (r *menuRepository) GetByDateWithDishes(ctx context.Context, offeredAt time.Time, city int32) (*domain.MenuWithDishes, error) {
	arg := db.GetMenuWithDishesByOfferedAtParams{
		OfferedAt: offeredAt,
		CityCode:  city,
	}

	result, err := r.query.GetMenuWithDishesByOfferedAt(ctx, arg)

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

func (r *menuRepository) FetchByRangeDateWithDishes(ctx context.Context, start, end time.Time, city int32, limit int32) ([]*domain.MenuWithDishes, error) {
	arg := db.ListMenuWithDishesByOfferedAtParams{
		StartOfferedAt: start,
		EndOfferedAt:   end,
		CityCode:       city,
		Limit:          limit,
	}

	results, err := r.query.ListMenuWithDishesByOfferedAt(ctx, arg)

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
