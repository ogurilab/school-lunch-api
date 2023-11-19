package db

import (
	"context"
	"testing"

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
	require.Equal(t, menu1.RegionID, menu2.RegionID)
	require.Equal(t, menu1.PhotoUrl, menu2.PhotoUrl.String)
	require.Equal(t, menu1.WikimediaCommonsUrl, menu2.WikimediaCommonsUrl.String)
	require.Equal(t, menu1.ElementarySchoolCalories, menu2.ElementarySchoolCalories.Int32)
	require.Equal(t, menu1.JuniorHighSchoolCalories, menu2.JuniorHighSchoolCalories.Int32)
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
		require.NotEmpty(t, menu.RegionID)
		require.NotEmpty(t, menu.PhotoUrl.String)
		require.NotEmpty(t, menu.WikimediaCommonsUrl.String)
		require.NotEmpty(t, menu.ElementarySchoolCalories.Int32)
		require.NotEmpty(t, menu.JuniorHighSchoolCalories.Int32)
		require.NotEmpty(t, menu.CreatedAt)
	}

}

func createRandomMenu(t *testing.T) *domain.Menu {
	ID := util.RandomUlid()

	args := CreateMenuParams{
		ID:                       ID,
		OfferedAt:                util.RandomDate(),
		RegionID:                 util.RandomInt32(),
		PhotoUrl:                 util.RandomSqlNullURL(),
		WikimediaCommonsUrl:      util.RandomSqlNullURL(),
		ElementarySchoolCalories: util.RandomSqlNullInt32(),

		JuniorHighSchoolCalories: util.RandomSqlNullInt32(),
	}

	err := testQuery.CreateMenu(context.Background(), args)

	require.NoError(t, err)

	menu, err := testQuery.GetMenu(context.Background(), ID)

	require.NoError(t, err)
	require.NotEmpty(t, menu)

	require.Equal(t, args.ID, menu.ID)
	require.Equal(t, args.OfferedAt, menu.OfferedAt)
	require.Equal(t, args.RegionID, menu.RegionID)
	require.Equal(t, args.PhotoUrl, menu.PhotoUrl)
	require.Equal(t, args.WikimediaCommonsUrl, menu.WikimediaCommonsUrl)
	require.Equal(t, args.ElementarySchoolCalories, menu.ElementarySchoolCalories)
	require.Equal(t, args.JuniorHighSchoolCalories, menu.JuniorHighSchoolCalories)
	require.NotEmpty(t, menu.CreatedAt)

	return domain.NewMenu(
		menu.ID,
		menu.OfferedAt,
		menu.RegionID,
		menu.PhotoUrl.String,
		menu.WikimediaCommonsUrl.String,
		menu.ElementarySchoolCalories.Int32,
		menu.JuniorHighSchoolCalories.Int32,
	)
}
