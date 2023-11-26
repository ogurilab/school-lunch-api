package db

import (
	"context"
	"encoding/json"
	"sort"
	"testing"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreateMenu(t *testing.T) {
	createRandomMenu(t)
}

func TestGetMenu(t *testing.T) {
	menu1 := createRandomMenu(t)
	menu2, err := testQuery.GetMenu(context.Background(), menu1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, menu2)

	require.Equal(t, menu1.ID, menu2.ID)
	require.Equal(t, menu1.OfferedAt, menu2.OfferedAt)

	require.Equal(t, menu1.PhotoUrl, menu2.PhotoUrl)
	require.Equal(t, menu1.ElementarySchoolCalories, menu2.ElementarySchoolCalories)
	require.Equal(t, menu1.JuniorHighSchoolCalories, menu2.JuniorHighSchoolCalories)
	require.NotEmpty(t, menu2.CreatedAt)
}

func TestListMenus(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomMenu(t)
	}

	arg := ListMenusParams{
		Limit:  5,
		Offset: 5,
	}

	menus, err := testQuery.ListMenus(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, menus, 5)

	for _, menu := range menus {
		require.NotEmpty(t, menu)
		require.NotEmpty(t, menu.ID)
		require.NotEmpty(t, menu.OfferedAt)
		require.NotEmpty(t, menu.PhotoUrl)
		require.NotEmpty(t, menu.ElementarySchoolCalories)
		require.NotEmpty(t, menu.JuniorHighSchoolCalories)
		require.NotEmpty(t, menu.CreatedAt)
	}
}

func TestGetMenuByOfferedAt(t *testing.T) {
	menu := createRandomMenu(t)

	result, err := testQuery.GetMenuByOfferedAt(context.Background(), menu.OfferedAt)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, menu.ID, result.ID)
	require.Equal(t, menu.OfferedAt, result.OfferedAt)
	require.Equal(t, menu.PhotoUrl, result.PhotoUrl)
	require.Equal(t, menu.ElementarySchoolCalories, result.ElementarySchoolCalories)
	require.Equal(t, menu.JuniorHighSchoolCalories, result.JuniorHighSchoolCalories)
	require.NotEmpty(t, result.CreatedAt)
}

func TestListMenusByOfferedAt(t *testing.T) {

	var offeredAts []time.Time

	for i := 0; i < 10; i++ {
		menu := createRandomMenu(t)
		offeredAts = append(offeredAts, menu.OfferedAt)
	}

	sort.Slice(offeredAts, func(i, j int) bool {
		return offeredAts[i].Before(offeredAts[j])
	})

	arg := ListMenusByOfferedAtParams{
		StartOfferedAt: offeredAts[0],
		EndOfferedAt:   offeredAts[5],
		Limit:          5,
	}

	menus, err := testQuery.ListMenusByOfferedAt(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, menus, 5)

	for _, menu := range menus {
		require.NotEmpty(t, menu)
		require.NotEmpty(t, menu.ID)
		require.NotEmpty(t, menu.OfferedAt)
		require.NotEmpty(t, menu.PhotoUrl)
		require.NotEmpty(t, menu.ElementarySchoolCalories)
		require.NotEmpty(t, menu.JuniorHighSchoolCalories)
		require.NotEmpty(t, menu.CreatedAt)

		require.True(t, menu.OfferedAt.After(arg.StartOfferedAt) || menu.OfferedAt.Equal(arg.StartOfferedAt))

		require.True(t, menu.OfferedAt.Before(arg.EndOfferedAt))
	}
}

func TestGetWithDishes(t *testing.T) {
	menu := createRandomMenu(t)

	var mockDishes []*domain.Dish

	for i := 0; i < 10; i++ {
		mockDishes = append(mockDishes, createRandomDish(t, menu.ID))
	}

	require.Len(t, mockDishes, 10)

	result, err := testQuery.GetMenuWithDishes(context.Background(), menu.ID)

	require.NoError(t, err)

	var dishes []*domain.Dish

	err = json.Unmarshal([]byte(result.Dishes), &dishes)

	require.NoError(t, err)

	for _, dish := range dishes {

		require.NotEmpty(t, dish.ID)
		require.NotEmpty(t, dish.MenuID)
		require.Equal(t, menu.ID, dish.MenuID)
		require.NotEmpty(t, dish.Name)

	}

}

func TestListMenuWithDishes(t *testing.T) {
	for i := 0; i < 10; i++ {
		menu := createRandomMenu(t)

		for j := 0; j < 10; j++ {
			createRandomDish(t, menu.ID)
		}
	}

	arg := ListMenuWithDishesParams{
		Limit:  5,
		Offset: 5,
	}

	results, err := testQuery.ListMenuWithDishes(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, results, 5)

	var dishes []*domain.Dish

	for _, result := range results {
		err := json.Unmarshal([]byte(result.Dishes), &dishes)
		require.NoError(t, err)
	}
}

func TestGetMenuWithDishesByOfferedAt(t *testing.T) {
	menu := createRandomMenu(t)

	var mockDishes []*domain.Dish

	for i := 0; i < 10; i++ {
		mockDishes = append(mockDishes, createRandomDish(t, menu.ID))
	}

	require.Len(t, mockDishes, 10)

	result, err := testQuery.GetMenuWithDishesByOfferedAt(context.Background(), menu.OfferedAt)

	require.NoError(t, err)

	var dishes []*domain.Dish

	err = json.Unmarshal([]byte(result.Dishes), &dishes)

	require.NoError(t, err)

	for _, dish := range dishes {

		require.NotEmpty(t, dish.ID)
		require.NotEmpty(t, dish.MenuID)
		require.Equal(t, menu.ID, dish.MenuID)
		require.NotEmpty(t, dish.Name)

	}

}

func TestListMenuWithDishesByOfferedAt(t *testing.T) {
	for i := 0; i < 10; i++ {
		menu := createRandomMenu(t)

		for j := 0; j < 10; j++ {
			createRandomDish(t, menu.ID)
		}
	}

	var offeredAts []time.Time

	for i := 0; i < 10; i++ {
		menu := createRandomMenu(t)
		offeredAts = append(offeredAts, menu.OfferedAt)
	}

	sort.Slice(offeredAts, func(i, j int) bool {
		return offeredAts[i].Before(offeredAts[j])
	})

	arg := ListMenuWithDishesByOfferedAtParams{
		StartOfferedAt: offeredAts[0],
		EndOfferedAt:   offeredAts[5],
		Limit:          5,
	}

	results, err := testQuery.ListMenuWithDishesByOfferedAt(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, results, 5)

	var dishes []*domain.Dish

	for _, result := range results {
		err := json.Unmarshal([]byte(result.Dishes), &dishes)
		require.NoError(t, err)
	}
}

func createRandomMenu(t *testing.T) *domain.Menu {
	ID := util.RandomUlid()

	args := CreateMenuParams{
		ID:                       ID,
		OfferedAt:                util.RandomDate(),
		PhotoUrl:                 util.RandomURL(),
		ElementarySchoolCalories: util.RandomInt32(),
		JuniorHighSchoolCalories: util.RandomInt32(),
	}

	err := testQuery.CreateMenu(context.Background(), args)

	require.NoError(t, err)

	menu, err := testQuery.GetMenu(context.Background(), ID)

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
	)

	require.NoError(t, err)

	return result
}
