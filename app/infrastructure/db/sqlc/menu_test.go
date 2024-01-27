package db

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

var cityCode = util.RandomCityCode()

func TestCreateMenu(t *testing.T) {
	createRandomMenu(t)
}

func TestGetMenu(t *testing.T) {
	menu1 := createRandomMenu(t)
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
	start := time.Now()
	for i := 0; i < 10; i++ {
		createRandomMenuFromStart(t, start)
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
	menu := createRandomMenu(t)

	var mockDishes []*domain.Dish

	for i := 0; i < 10; i++ {
		mockDishes = append(mockDishes, createRandomDish(t, menu.ID))
	}

	require.Len(t, mockDishes, 10)

	arg := GetMenuWithDishesParams{
		ID:       menu.ID,
		CityCode: menu.CityCode,
	}

	result, err := testQuery.GetMenuWithDishes(context.Background(), arg)

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

func TestFetchMenuWithDishesByCity(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		menu := createRandomMenuFromStart(t, start)

		for j := 0; j < 10; j++ {
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
	require.Len(t, results, 5)

	var dishes []*domain.Dish

	for _, result := range results {
		err := json.Unmarshal([]byte(result.Dishes), &dishes)
		require.NoError(t, err)
	}
}
func TestFetchMenuWithDishes(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		menu := createRandomMenuFromStart(t, start)

		for j := 0; j < 10; j++ {
			createRandomDish(t, menu.ID)
		}
	}

	arg := ListMenuWithDishesParams{
		Limit:     5,
		Offset:    5,
		OfferedAt: start,
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

func createRandomMenu(t *testing.T) *domain.Menu {
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
func createRandomMenuFromStart(t *testing.T, start time.Time) *domain.Menu {
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
