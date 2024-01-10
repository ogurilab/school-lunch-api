package db

import (
	"context"
	"testing"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreateDish(t *testing.T) {
	menu := createRandomMenu(t)
	createRandomDish(t, menu.ID)
}

func TestGetDish(t *testing.T) {
	menu := createRandomMenu(t)
	dish1 := createRandomDish(t, menu.ID)
	dish2, err := testQuery.GetDish(context.Background(), dish1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, dish2)

	require.Equal(t, dish1.ID, dish2.ID)
	require.Equal(t, dish1.MenuID, dish2.MenuID)
	require.Equal(t, dish1.Name, dish2.Name)

}

func TestFetchDishByMenuID(t *testing.T) {
	menu := createRandomMenu(t)
	for i := 0; i < 10; i++ {
		createRandomDish(t, menu.ID)
	}

	dishes, err := testQuery.ListDishByMenuID(context.Background(), menu.ID)

	require.NoError(t, err)
	require.Len(t, dishes, 10)

	for _, dish := range dishes {
		require.Equal(t, menu.ID, dish.MenuID)
	}
}

func TestFetchDishesByName(t *testing.T) {
	menu := createRandomMenu(t)

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
		require.Equal(t, mockDishes[0].MenuID, dish.MenuID)
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
		require.NotEmpty(t, dish.MenuID)
		require.NotEmpty(t, dish.Name)

	}
}

func createRandomDish(t *testing.T, menuID string) *domain.Dish {
	arg := CreateDishParams{
		ID:     util.RandomUlid(),
		MenuID: menuID,
		Name:   util.RandomString(10),
	}

	err := testQuery.CreateDish(context.Background(), arg)

	require.NoError(t, err)

	dish, err := testQuery.GetDish(context.Background(), arg.ID)

	require.NoError(t, err)
	require.NotEmpty(t, dish)

	require.Equal(t, arg.ID, dish.ID)
	require.Equal(t, arg.MenuID, dish.MenuID)
	require.Equal(t, arg.Name, dish.Name)

	result, err := domain.ReNewDish(dish.ID, dish.MenuID, dish.Name)

	require.NoError(t, err)

	return result
}
