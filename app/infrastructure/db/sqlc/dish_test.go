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

func TestFetchMenuByID(t *testing.T) {
	menu := createRandomMenu(t)
	for i := 0; i < 10; i++ {
		createRandomDish(t, menu.ID)
	}

	dishes, err := testQuery.ListDishes(context.Background(), menu.ID)

	require.NoError(t, err)
	require.Len(t, dishes, 5)

	for _, dish := range dishes {
		require.NotEmpty(t, dish)
		require.NotEmpty(t, dish.ID)
		require.NotEmpty(t, dish.MenuID)
		require.NotEmpty(t, dish.Name)
		require.NotEmpty(t, dish.CreatedAt)
	}
}

func TestFetchByNames(t *testing.T) {
	menu := createRandomMenu(t)

	var names []string

	for i := 0; i < 10; i++ {
		dish := createRandomDish(t, menu.ID)
		names = append(names, dish.Name)
	}

	arg := GetDishByNamesParams{
		Names:  names,
		Limit:  5,
		Offset: 5,
	}

	dishes, err := testQuery.GetDishByNames(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, dishes, 5)

	containsNames := names[5:]
	notContainsNames := names[:5]

	for _, dish := range dishes {
		require.NotEmpty(t, dish)
		require.NotEmpty(t, dish.ID)
		require.NotEmpty(t, dish.MenuID)
		require.NotEmpty(t, dish.Name)
		require.NotEmpty(t, dish.CreatedAt)

		require.Contains(t, containsNames, dish.Name)
		require.NotContains(t, notContainsNames, dish.Name)
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
