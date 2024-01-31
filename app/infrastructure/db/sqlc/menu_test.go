package db

import (
	"context"
	"testing"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreateMenu(t *testing.T) {
	cityCode := util.RandomCityCode()
	createRandomMenu(t, cityCode)
}

func TestGetMenu(t *testing.T) {
	cityCode := util.RandomCityCode()
	menu1 := createRandomMenu(t, cityCode)

	arg := GetMenuParams{
		ID:       menu1.ID,
		CityCode: menu1.CityCode,
	}

	menu2, err := testQuery.GetMenu(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, menu2)

	require.Equal(t, menu1.ID, menu2.ID)
	require.Equal(t, menu1.OfferedAt, menu2.OfferedAt)

	require.Equal(t, menu1.PhotoUrl, menu2.PhotoUrl)
	require.Equal(t, menu1.ElementarySchoolCalories, menu2.ElementarySchoolCalories)
	require.Equal(t, menu1.JuniorHighSchoolCalories, menu2.JuniorHighSchoolCalories)
	require.NotEmpty(t, menu2.CreatedAt)
}

func TestFetchMenusByCity(t *testing.T) {
	cityCode := util.RandomCityCode()
	start := time.Now()
	for i := 0; i < 10; i++ {
		createRandomMenuFromStart(t, start, cityCode)
	}

	arg := ListMenuByCityParams{
		Limit:     5,
		Offset:    5,
		CityCode:  cityCode,
		OfferedAt: start,
	}

	menus, err := testQuery.ListMenuByCity(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, menus, 5)
}

func TestGetWithDishes(t *testing.T) {
	cityCode := util.RandomCityCode()
	menu := createRandomMenu(t, cityCode)

	var mockDishes []*domain.Dish

	for i := 0; i < 10; i++ {
		mockDishes = append(mockDishes, createRandomDish(t, menu.ID))
	}

	require.Len(t, mockDishes, 10)

	arg := GetMenuWithDishesParams{
		ID:       menu.ID,
		CityCode: menu.CityCode,
	}

	results, err := testQuery.GetMenuWithDishes(context.Background(), arg)

	require.NoError(t, err)

	var dishes []*domain.Dish

	for _, result := range results {
		dish, err := domain.ReNewDish(result.DishID, result.DishName)

		require.NoError(t, err)

		dishes = append(dishes, dish)
	}

	require.Len(t, dishes, 10)
	require.NoError(t, err)

	for i, dish := range dishes {
		require.Equal(t, mockDishes[i].ID, dish.ID)
		require.Equal(t, mockDishes[i].Name, dish.Name)
	}
}

func TestFetchMenuWithDishesByCity(t *testing.T) {
	cityCode := util.RandomCityCode()
	start := util.RandomDate()

	for i := 0; i < 10; i++ {
		menu := createRandomMenuFromStart(t, start, cityCode)
		for j := 0; j < 5; j++ {
			createRandomDish(t, menu.ID)
		}
	}

	arg := ListMenuWithDishesByCityParams{
		Limit:     5,
		Offset:    5,
		CityCode:  cityCode,
		OfferedAt: start,
	}

	results, err := testQuery.ListMenuWithDishesByCity(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, results, 25)

	require.NoError(t, err)

	var menus []*domain.MenuWithDishes

	menusMap := make(map[string]*domain.Menu)
	dishesMap := make(map[string][]*domain.Dish)

	for _, result := range results {
		if _, exists := menusMap[result.ID]; !exists {
			menu, err := domain.ReNewMenu(
				result.ID,
				result.OfferedAt,
				result.PhotoUrl,
				result.ElementarySchoolCalories,
				result.JuniorHighSchoolCalories,
				result.CityCode,
			)
			require.NoError(t, err)
			menusMap[result.ID] = menu
		}

		dish, err := domain.ReNewDish(
			result.DishID,
			result.DishName,
		)
		require.NoError(t, err)

		dishesMap[result.ID] = append(dishesMap[result.ID], dish)
	}

	for id, menu := range menusMap {
		withDishes, err := domain.ReNewMenuWithDishes(
			menu.ID,
			menu.OfferedAt,
			menu.PhotoUrl,
			menu.ElementarySchoolCalories,
			menu.JuniorHighSchoolCalories,
			menu.CityCode,
			dishesMap[id],
		)

		require.NoError(t, err)

		menus = append(menus, withDishes)
	}

	require.Len(t, menus, 5)

	for _, menu := range menus {
		require.Len(t, menu.Dishes, 5)
	}
}

func TestFetchMenuWithDishes(t *testing.T) {
	err := testQuery.truncateMenusTable()
	require.NoError(t, err)
	cityCode := util.RandomCityCode()
	start := util.RandomDate()

	for i := 0; i < 10; i++ {

		menu := createRandomMenuFromStart(t, start, cityCode)

		for j := 0; j < 5; j++ {
			createRandomDish(t, menu.ID)
		}

	}

	arg := ListMenuWithDishesParams{
		Limit:     5,
		Offset:    0,
		OfferedAt: start,
	}

	results, err := testQuery.ListMenuWithDishes(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, results, 25)

	mapMenus := make(map[string]*domain.Menu)
	mapDishes := make(map[string][]*domain.Dish)

	for _, result := range results {
		if _, exists := mapMenus[result.ID]; !exists {
			menu, err := domain.ReNewMenu(
				result.ID,
				result.OfferedAt,
				result.PhotoUrl,
				result.ElementarySchoolCalories,
				result.JuniorHighSchoolCalories,
				result.CityCode,
			)
			require.NoError(t, err)
			mapMenus[result.ID] = menu
		}

		dish, err := domain.ReNewDish(
			result.DishID,
			result.DishName,
		)
		require.NoError(t, err)

		mapDishes[result.ID] = append(mapDishes[result.ID], dish)
	}

	var menus []*domain.MenuWithDishes

	for id, menu := range mapMenus {
		withDishes, err := domain.ReNewMenuWithDishes(
			menu.ID,
			menu.OfferedAt,
			menu.PhotoUrl,
			menu.ElementarySchoolCalories,
			menu.JuniorHighSchoolCalories,
			menu.CityCode,
			mapDishes[id],
		)

		require.NoError(t, err)

		menus = append(menus, withDishes)
	}

	require.Len(t, menus, 5)

	for _, menu := range menus {
		require.Len(t, menu.Dishes, 5)
	}
}

func createRandomMenu(t *testing.T, cityCode int32) *domain.Menu {
	id := util.RandomUlid()

	args := CreateMenuParams{
		ID:                       id,
		OfferedAt:                util.RandomDate(),
		PhotoUrl:                 util.RandomNullURL(),
		ElementarySchoolCalories: util.RandomInt32(),
		JuniorHighSchoolCalories: util.RandomInt32(),
		CityCode:                 cityCode,
	}

	err := testQuery.CreateMenu(context.Background(), args)

	require.NoError(t, err)

	getArg := GetMenuParams{
		ID:       id,
		CityCode: cityCode,
	}

	menu, err := testQuery.GetMenu(context.Background(), getArg)

	require.NoError(t, err)
	require.NotEmpty(t, menu)

	require.Equal(t, args.ID, menu.ID)
	require.Equal(t, args.OfferedAt, menu.OfferedAt)
	require.Equal(t, args.PhotoUrl, menu.PhotoUrl)
	require.Equal(t, args.ElementarySchoolCalories, menu.ElementarySchoolCalories)
	require.Equal(t, args.JuniorHighSchoolCalories, menu.JuniorHighSchoolCalories)
	require.NotEmpty(t, menu.CreatedAt)

	result, err := domain.ReNewMenu(
		menu.ID,
		menu.OfferedAt,
		menu.PhotoUrl,
		menu.ElementarySchoolCalories,
		menu.JuniorHighSchoolCalories,
		menu.CityCode,
	)

	require.NoError(t, err)

	return result
}

func createRandomMenuFromStart(t *testing.T, start time.Time, cityCode int32) *domain.Menu {
	id := util.RandomUlid()

	args := CreateMenuParams{
		ID:                       id,
		OfferedAt:                util.RandomDateFromStart(start),
		PhotoUrl:                 util.RandomNullURL(),
		ElementarySchoolCalories: util.RandomInt32(),
		JuniorHighSchoolCalories: util.RandomInt32(),
		CityCode:                 cityCode,
	}

	err := testQuery.CreateMenu(context.Background(), args)

	require.NoError(t, err)

	getArg := GetMenuParams{
		ID:       id,
		CityCode: cityCode,
	}

	menu, err := testQuery.GetMenu(context.Background(), getArg)

	require.NoError(t, err)
	require.NotEmpty(t, menu)

	require.Equal(t, args.ID, menu.ID)
	require.Equal(t, args.OfferedAt, menu.OfferedAt)
	require.Equal(t, args.PhotoUrl, menu.PhotoUrl)
	require.Equal(t, args.ElementarySchoolCalories, menu.ElementarySchoolCalories)
	require.Equal(t, args.JuniorHighSchoolCalories, menu.JuniorHighSchoolCalories)
	require.NotEmpty(t, menu.CreatedAt)

	result, err := domain.ReNewMenu(
		menu.ID,
		menu.OfferedAt,
		menu.PhotoUrl,
		menu.ElementarySchoolCalories,
		menu.JuniorHighSchoolCalories,
		menu.CityCode,
	)

	require.NoError(t, err)

	return result
}
