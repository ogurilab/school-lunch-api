package repository

import (
	"context"
	"database/sql"
	"sort"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
)

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

	results, err := r.query.GetMenuWithDishes(ctx, arg)

	if err != nil {
		return nil, err
	}

	dishes := make([]*domain.Dish, 0, len(results))

	for _, result := range results {
		dish, err := domain.ReNewDish(
			result.DishID,
			result.DishName,
		)

		if err != nil {
			return nil, err
		}

		dishes = append(dishes, dish)
	}

	menuData := results[0]

	return domain.ReNewMenuWithDishes(
		menuData.ID,
		menuData.OfferedAt,
		menuData.PhotoUrl,
		menuData.ElementarySchoolCalories,
		menuData.JuniorHighSchoolCalories,
		menuData.CityCode,
		dishes,
	)
}

type mapKey struct {
	id      string
	offered time.Time
}

type processMenuDishesMapInput struct {
	id                       string
	offered                  time.Time
	photoUrl                 sql.NullString
	elementarySchoolCalories int32
	juniorHighSchoolCalories int32
	cityCode                 int32
	dishID                   string
	dishName                 string
	key                      mapKey
	menuMap                  map[mapKey]*domain.Menu
	dishesMap                map[mapKey][]*domain.Dish
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

	menusMap := make(map[mapKey]*domain.Menu)
	dishesMap := make(map[mapKey][]*domain.Dish)

	for _, result := range results {
		key := mapKey{id: result.ID, offered: result.OfferedAt}

		err := processMenuDishesMap(processMenuDishesMapInput{
			id:                       result.ID,
			offered:                  result.OfferedAt,
			photoUrl:                 result.PhotoUrl,
			elementarySchoolCalories: result.ElementarySchoolCalories,
			juniorHighSchoolCalories: result.JuniorHighSchoolCalories,
			cityCode:                 result.CityCode,
			dishID:                   result.DishID,
			dishName:                 result.DishName,
			key:                      key,
			menuMap:                  menusMap,
			dishesMap:                dishesMap,
		})

		if err != nil {
			return nil, err
		}
	}

	return processMenuWithDishesResults(menusMap, dishesMap, len(menusMap))
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

	menusMap := make(map[mapKey]*domain.Menu)
	dishesMap := make(map[mapKey][]*domain.Dish)

	for _, result := range results {
		key := mapKey{id: result.ID, offered: result.OfferedAt}

		err := processMenuDishesMap(processMenuDishesMapInput{
			id:                       result.ID,
			offered:                  result.OfferedAt,
			photoUrl:                 result.PhotoUrl,
			elementarySchoolCalories: result.ElementarySchoolCalories,
			juniorHighSchoolCalories: result.JuniorHighSchoolCalories,
			cityCode:                 result.CityCode,
			dishID:                   result.DishID,
			dishName:                 result.DishName,
			key:                      key,
			menuMap:                  menusMap,
			dishesMap:                dishesMap,
		})

		if err != nil {
			return nil, err
		}
	}

	return processMenuWithDishesResults(menusMap, dishesMap, len(menusMap))
}

func processMenuDishesMap(
	input processMenuDishesMapInput,
) error {
	menusMap := input.menuMap
	dishesMap := input.dishesMap

	key := input.key
	if _, exists := menusMap[key]; !exists {
		menu, err := domain.ReNewMenu(
			input.id,
			input.offered,
			input.photoUrl,
			input.elementarySchoolCalories,
			input.juniorHighSchoolCalories,
			input.cityCode,
		)
		if err != nil {

			return err
		}

		menusMap[key] = menu
	}

	dish, err := domain.ReNewDish(
		input.dishID,
		input.dishName,
	)
	if err != nil {
		return err
	}

	dishesMap[key] = append(dishesMap[key], dish)

	return nil
}

func processMenuWithDishesResults(menuMap map[mapKey]*domain.Menu, dishesMap map[mapKey][]*domain.Dish, length int) ([]*domain.MenuWithDishes, error) {

	keys := make([]mapKey, 0, length)

	for key := range menuMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].offered.After(keys[j].offered)
	})

	menus := make([]*domain.MenuWithDishes, 0, length)

	for _, key := range keys {
		menu := menuMap[key]
		dishes := dishesMap[key]

		menuWithDishes, err := domain.ReNewMenuWithDishes(
			menu.ID,
			menu.OfferedAt,
			menu.PhotoUrl,
			menu.ElementarySchoolCalories,
			menu.JuniorHighSchoolCalories,
			menu.CityCode,
			dishes,
		)

		if err != nil {
			return nil, err
		}

		menus = append(menus, menuWithDishes)
	}

	return menus, nil
}
