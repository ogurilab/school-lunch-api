package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ogurilab/school-lunch-api/domain"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc/mocks"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetByIDWithDishes(t *testing.T) {
	id := util.NewUlid()
	queryResults := randomMenuWithDishesRows()

	testCases := []struct {
		name      string
		input     db.GetMenuWithDishesParams
		buildStub func(query *mocks.MockQuery)
		check     func(t *testing.T, menu *domain.MenuWithDishes, err error)
	}{
		{
			name: "OK",
			input: db.GetMenuWithDishesParams{
				ID:       id,
				CityCode: 1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.GetMenuWithDishesParams{
					ID:       id,
					CityCode: 1,
				}
				query.EXPECT().GetMenuWithDishes(context.Background(), arg).Times(1).Return(queryResults, nil)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				menuData := queryResults[0]

				require.NoError(t, err)
				require.NotNil(t, menu)

				require.Equal(t, menuData.ID, menu.ID)
				require.Equal(t, menuData.OfferedAt, menu.OfferedAt)
				require.Equal(t, menuData.PhotoUrl, menu.PhotoUrl)
				require.Equal(t, menuData.ElementarySchoolCalories, menu.ElementarySchoolCalories)
				require.Equal(t, menuData.JuniorHighSchoolCalories, menu.JuniorHighSchoolCalories)
				require.Equal(t, menuData.CityCode, menu.CityCode)

				require.NotNil(t, menu.Dishes)
				require.Len(t, menu.Dishes, len(queryResults))

				for i, dish := range menu.Dishes {
					require.Equal(t, queryResults[i].DishID, dish.ID)
					require.Equal(t, queryResults[i].DishName, dish.Name)
				}

			},
		},
		{
			name: "Bad ID",
			input: db.GetMenuWithDishesParams{
				ID:       "bad_id",
				CityCode: 1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.GetMenuWithDishesParams{
					ID:       "bad_id",
					CityCode: 1,
				}
				query.EXPECT().GetMenuWithDishes(context.Background(), arg).Times(1).Return([]db.GetMenuWithDishesRow{}, sql.ErrNoRows)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrNoRows)
				require.Nil(t, menu)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.buildStub(query)

			repo := NewMenuWithDishesRepository(query)

			menu, err := repo.GetByID(context.Background(), tc.input.ID, tc.input.CityCode)

			tc.check(t, menu, err)
		})
	}
}

func TestFetchByCityWithDishes(t *testing.T) {
	offered := util.RandomDate()

	testCases := []struct {
		name  string
		input db.ListMenuWithDishesByCityParams
		build func(query *mocks.MockQuery)
		check func(t *testing.T, menus []*domain.MenuWithDishes, err error)
	}{
		{
			name: "OK",
			input: db.ListMenuWithDishesByCityParams{
				CityCode:  1,
				Limit:     10,
				Offset:    0,
				OfferedAt: offered,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesByCityParams{
					CityCode:  1,
					Limit:     10,
					Offset:    0,
					OfferedAt: offered,
				}
				results := randomWithDishesByCityResults(int(arg.Limit))

				query.EXPECT().ListMenuWithDishesByCity(context.Background(), arg).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)
				require.Len(t, menus, 10)
			},
		},
		{
			name: "Bad Limit",
			input: db.ListMenuWithDishesByCityParams{
				CityCode:  1,
				Limit:     -1,
				Offset:    0,
				OfferedAt: offered,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesByCityParams{
					CityCode:  1,
					Limit:     -1,
					Offset:    0,
					OfferedAt: offered,
				}
				query.EXPECT().ListMenuWithDishesByCity(context.Background(), arg).Times(1).Return([]db.ListMenuWithDishesByCityRow{}, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrNoRows)
				require.Nil(t, menus)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.build(query)

			repo := NewMenuWithDishesRepository(query)

			menus, err := repo.FetchByCity(context.Background(), tc.input.Limit, tc.input.Offset, tc.input.OfferedAt, tc.input.CityCode)

			tc.check(t, menus, err)
		})
	}
}

func TestFetchWithDishes(t *testing.T) {
	offered := util.RandomDate()

	testCases := []struct {
		name  string
		input db.ListMenuWithDishesParams
		build func(query *mocks.MockQuery)
		check func(t *testing.T, menus []*domain.MenuWithDishes, err error)
	}{
		{
			name: "OK",
			input: db.ListMenuWithDishesParams{
				Limit:     10,
				Offset:    0,
				OfferedAt: offered,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesParams{
					Limit:     10,
					Offset:    0,
					OfferedAt: offered,
				}
				results := randomWithDishesResults(int(arg.Limit))
				query.EXPECT().ListMenuWithDishes(context.Background(), arg).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)
				require.Len(t, menus, 10)
			},
		},
		{
			name: "Bad Limit",
			input: db.ListMenuWithDishesParams{
				Limit:     -1,
				Offset:    0,
				OfferedAt: offered,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesParams{
					Limit:     -1,
					Offset:    0,
					OfferedAt: offered,
				}
				query.EXPECT().ListMenuWithDishes(context.Background(), arg).Times(1).Return([]db.ListMenuWithDishesRow{}, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrNoRows)
				require.Nil(t, menus)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.build(query)

			repo := NewMenuWithDishesRepository(query)

			menus, err := repo.Fetch(context.Background(), tc.input.Limit, tc.input.Offset, tc.input.OfferedAt)

			tc.check(t, menus, err)
		})
	}
}

func randomWithDishesResults(length int) []db.ListMenuWithDishesRow {

	results := make([]db.ListMenuWithDishesRow, 0, length)

	for i := 0; i < length; i++ {
		results = append(results, db.ListMenuWithDishesRow{
			ID:                       util.NewUlid(),
			OfferedAt:                util.RandomDate(),
			CityCode:                 util.RandomCityCode(),
			PhotoUrl:                 util.RandomNullURL(),
			ElementarySchoolCalories: util.RandomInt32(),
			JuniorHighSchoolCalories: util.RandomInt32(),
			DishID:                   util.NewUlid(),
			DishName:                 "dish",
		})
	}

	return results
}

func randomWithDishesByCityResults(length int) []db.ListMenuWithDishesByCityRow {

	results := make([]db.ListMenuWithDishesByCityRow, 0, length)

	for i := 0; i < length; i++ {
		data := db.ListMenuWithDishesByCityRow{
			ID:                       util.NewUlid(),
			OfferedAt:                util.RandomDate(),
			CityCode:                 util.RandomCityCode(),
			PhotoUrl:                 util.RandomNullURL(),
			ElementarySchoolCalories: util.RandomInt32(),
			JuniorHighSchoolCalories: util.RandomInt32(),
			DishID:                   util.NewUlid(),
			DishName:                 "dish",
		}

		results = append(results, data)
	}

	return results
}

func randomMenuWithDishesRows() []db.GetMenuWithDishesRow {
	n := 5

	results := make([]db.GetMenuWithDishesRow, n)

	for i := 0; i < n; i++ {
		results[i] = db.GetMenuWithDishesRow{
			ID:                       util.NewUlid(),
			OfferedAt:                util.RandomDate(),
			CityCode:                 util.RandomCityCode(),
			PhotoUrl:                 util.RandomNullURL(),
			ElementarySchoolCalories: util.RandomInt32(),
			JuniorHighSchoolCalories: util.RandomInt32(),
			DishID:                   util.NewUlid(),
			DishName:                 "dish",
		}
	}

	return results
}
