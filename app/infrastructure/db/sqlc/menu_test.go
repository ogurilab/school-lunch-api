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

func TestFetchMenus(t *testing.T) {
	cityCode := util.RandomCityCode()
	start := time.Now()
	for i := 0; i < 10; i++ {
		createRandomMenuFromStart(t, start, cityCode)
	}

	arg := ListMenuParams{
		Limit:     5,
		Offset:    5,
		OfferedAt: start,
	}

	menus, err := testQuery.ListMenu(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, menus, 5)
}

func TestFetchMenusInId(t *testing.T) {
	cityCode := util.RandomCityCode()
	start := time.Now()
	menus := make([]*domain.Menu, 0)

	for i := 0; i < 10; i++ {
		menu := createRandomMenuFromStart(t, start, cityCode)
		menus = append(menus, menu)
	}

	ids := make([]string, 0, 5)

	for i := 0; i < 5; i++ {
		ids = append(ids, menus[i].ID)
	}

	arg := ListMenuInIdsParams{
		Ids:       ids,
		OfferedAt: start,
		Limit:     domain.DEFAULT_LIMIT,
		Offset:    domain.DEFAULT_OFFSET,
	}

	resMenus, err := testQuery.ListMenuInIds(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, resMenus, 5)

	resMenuIds := make([]string, 0, 5)

	for _, menu := range resMenus {
		resMenuIds = append(resMenuIds, menu.ID)
	}

	require.ElementsMatch(t, ids, resMenuIds)

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
