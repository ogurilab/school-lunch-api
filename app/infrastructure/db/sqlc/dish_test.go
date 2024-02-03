package db

import (
	"context"
	"testing"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreateDish(t *testing.T) {
	cityCode := util.RandomCityCode()
	menu := createRandomMenu(t, cityCode)
	createRandomDish(t, menu.ID)
}

func TestGetDish(t *testing.T) {
	cityCode := util.RandomCityCode()
	dish := createRandomDish(t, util.RandomUlid())
	testQuery.truncateMenuDishesTable()

	length := 10
	menuIDs := createMenuDishesByDishID(t, dish.ID, cityCode, length)

	getArg := GetDishParams{
		ID:     dish.ID,
		Limit:  int32(length),
		Offset: 0,
	}

	results, err := testQuery.GetDish(context.Background(), getArg)

	require.NoError(t, err)

	resMenuIDs := make([]string, 0, length)

	for _, result := range results {
		resMenuIDs = append(resMenuIDs, result.MenuID)
	}

	require.Len(t, results, length)
	require.ElementsMatch(t, menuIDs, resMenuIDs)

	for _, result := range results {
		require.NotEmpty(t, result.ID)
		require.NotEmpty(t, result.Name)

		require.Equal(t, dish.ID, result.ID)
		require.Equal(t, dish.Name, result.Name)
	}
}

func TestGetDishInCity(t *testing.T) {
	cityCode := util.RandomCityCode()
	dish := createRandomDish(t, util.RandomUlid())
	testQuery.truncateMenuDishesTable()

	length := 10
	menuIDs := createMenuDishesByDishID(t, dish.ID, cityCode, length)

	getArg := GetDishInCityParams{
		ID:       dish.ID,
		CityCode: cityCode,
		Limit:    int32(length),
		Offset:   0,
	}

	results, err := testQuery.GetDishInCity(context.Background(), getArg)

	require.NoError(t, err)

	resMenuIDs := make([]string, 0, length)

	for _, result := range results {
		resMenuIDs = append(resMenuIDs, result.MenuID)
	}

	require.Len(t, results, length)
	require.ElementsMatch(t, menuIDs, resMenuIDs)

	for _, result := range results {
		require.NotEmpty(t, result.ID)
		require.NotEmpty(t, result.Name)

		require.Equal(t, dish.ID, result.ID)
		require.Equal(t, dish.Name, result.Name)
	}
}

func TestFetchDishByMenuID(t *testing.T) {
	cityCode := util.RandomCityCode()
	menu := createRandomMenu(t, cityCode)
	for i := 0; i < 10; i++ {
		createRandomDish(t, menu.ID)
	}

	dishes, err := testQuery.ListDishByMenuID(context.Background(), menu.ID)

	require.NoError(t, err)
	require.Len(t, dishes, 10)
}

func TestFetchDishesByName(t *testing.T) {
	cityCode := util.RandomCityCode()
	menu := createRandomMenu(t, cityCode)

	var mockDishes []*domain.Dish

	for i := 0; i < 10; i++ {
		mockDishes = append(mockDishes, createRandomDish(t, menu.ID))

	}

	arg := ListDishByNameParams{
		Name:   mockDishes[0].Name,
		Limit:  5,
		Offset: 0,
	}

	dishes, err := testQuery.ListDishByName(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, dishes, 1)

	for _, dish := range dishes {
		require.Equal(t, mockDishes[0].Name, dish.Name)
		require.Equal(t, arg.Name, dish.Name)
		require.Equal(t, mockDishes[0].ID, dish.ID)
	}
}

func TestFetchDishes(t *testing.T) {

	var mockDishes []*domain.Dish

	for i := 0; i < 10; i++ {
		dish := createRandomDish(t, util.RandomUlid())
		mockDishes = append(mockDishes, dish)
	}

	require.Len(t, mockDishes, 10)

	arg := ListDishParams{
		Limit:  5,
		Offset: 5,
	}

	dishes, err := testQuery.ListDish(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, dishes, 5)

	for _, dish := range dishes {
		require.NotEmpty(t, dish.ID)
		require.NotEmpty(t, dish.Name)

	}
}

func createMenuDishesByDishID(t *testing.T, dishID string, cityCode int32, length int) []string {

	menus := make([]*domain.Menu, 0, length)

	for i := 0; i < length; i++ {
		menu := createRandomMenu(t, cityCode)
		menus = append(menus, menu)
	}

	for _, menu := range menus {
		relationArg := CreateMenuDishParams{
			MenuID: menu.ID,
			DishID: dishID,
		}

		err := testQuery.CreateMenuDish(context.Background(), relationArg)

		require.NoError(t, err)
	}

	menuIDs := make([]string, 0, length)

	for _, menu := range menus {
		menuIDs = append(menuIDs, menu.ID)
	}

	return menuIDs
}

func createRandomDish(t *testing.T, menuID string) *domain.Dish {
	arg := CreateDishParams{
		ID:   util.RandomUlid(),
		Name: util.RandomString(10),
	}

	err := testQuery.CreateDish(context.Background(), arg)

	require.NoError(t, err)

	relationArg := CreateMenuDishParams{
		MenuID: menuID,
		DishID: arg.ID,
	}

	err = testQuery.CreateMenuDish(context.Background(), relationArg)

	require.NoError(t, err)

	getArg := GetDishParams{
		ID:     arg.ID,
		Limit:  1,
		Offset: 0,
	}

	res, err := testQuery.GetDish(context.Background(), getArg)

	dish := res[0]
	require.NoError(t, err)
	require.NotEmpty(t, dish)

	require.Equal(t, arg.ID, dish.ID)
	require.Equal(t, relationArg.DishID, dish.ID)
	require.Equal(t, arg.Name, dish.Name)

	result, err := domain.ReNewDish(dish.ID, dish.Name)

	require.NoError(t, err)

	return result
}
