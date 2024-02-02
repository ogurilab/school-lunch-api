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

func TestCreate(t *testing.T) {
	ctx := context.Background()

	menu := randomMenu(t)

	testCases := []struct {
		name      string
		input     *domain.Menu
		buildStub func(query *mocks.MockQuery)
		check     func(t *testing.T, err error)
	}{
		{
			name:  "OK",
			input: menu,
			buildStub: func(query *mocks.MockQuery) {
				arg := db.CreateMenuParams{
					ID:                       menu.ID,
					OfferedAt:                menu.OfferedAt,
					CityCode:                 menu.CityCode,
					PhotoUrl:                 menu.PhotoUrl,
					ElementarySchoolCalories: menu.ElementarySchoolCalories,
					JuniorHighSchoolCalories: menu.JuniorHighSchoolCalories,
				}
				query.EXPECT().CreateMenu(ctx, arg).Times(1).Return(nil)
			},
			check: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.buildStub(query)

			repo := NewMenuRepository(query)

			err := repo.Create(ctx, tc.input)

			tc.check(t, err)
		})
	}
}

func TestGetByID(t *testing.T) {

	ctx := context.Background()
	id := util.NewUlid()

	testCases := []struct {
		name      string
		input     db.GetMenuParams
		buildStub func(query *mocks.MockQuery)
		check     func(t *testing.T, menu *domain.Menu, err error)
	}{
		{
			name: "OK",
			input: db.GetMenuParams{
				ID:       id,
				CityCode: 1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.GetMenuParams{
					ID:       id,
					CityCode: 1,
				}
				result := db.Menu{
					ID:                       arg.ID,
					OfferedAt:                util.RandomDate(),
					CityCode:                 arg.CityCode,
					PhotoUrl:                 util.RandomNullURL(),
					ElementarySchoolCalories: util.RandomInt32(),
					JuniorHighSchoolCalories: util.RandomInt32(),
				}
				query.EXPECT().GetMenu(ctx, arg).Times(1).Return(result, nil)
			},
			check: func(t *testing.T, menu *domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menu)
				require.Equal(t, int32(1), menu.CityCode)
				require.NotEmpty(t, menu.OfferedAt)
				require.NotEmpty(t, menu.PhotoUrl)
				require.NotEmpty(t, menu.ElementarySchoolCalories)
				require.NotEmpty(t, menu.JuniorHighSchoolCalories)

			},
		},
		{
			name: "Bad ID",
			input: db.GetMenuParams{
				ID:       "bad_id",
				CityCode: 1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.GetMenuParams{
					ID:       "bad_id",
					CityCode: 1,
				}
				query.EXPECT().GetMenu(ctx, arg).Times(1).Return(db.Menu{}, sql.ErrNoRows)
			},
			check: func(t *testing.T, menu *domain.Menu, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrNoRows)
				require.Nil(t, menu)
			},
		},
		{
			name: "Bad CityCode",
			input: db.GetMenuParams{
				ID:       id,
				CityCode: -1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.GetMenuParams{
					ID:       id,
					CityCode: -1,
				}
				query.EXPECT().GetMenu(ctx, arg).Times(1).Return(db.Menu{}, sql.ErrNoRows)
			},
			check: func(t *testing.T, menu *domain.Menu, err error) {
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

			repo := NewMenuRepository(query)

			menu, err := repo.GetByID(ctx, tc.input.ID, tc.input.CityCode)

			tc.check(t, menu, err)
		})
	}
}
func TestFetchMenuByCity(t *testing.T) {
	ctx := context.Background()
	offered := util.RandomDate()

	testCases := []struct {
		name      string
		input     db.ListMenuByCityParams
		buildStub func(query *mocks.MockQuery)
		check     func(t *testing.T, menus []*domain.Menu, err error)
	}{
		{
			name: "OK",
			input: db.ListMenuByCityParams{
				Limit:     10,
				Offset:    0,
				OfferedAt: offered,
				CityCode:  1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.ListMenuByCityParams{
					Limit:     10,
					Offset:    0,
					OfferedAt: offered,
					CityCode:  1,
				}
				results := randomMenuResults(10)
				query.EXPECT().ListMenuByCity(ctx, arg).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)
				require.Len(t, menus, 10)
			},
		},
		{
			name: "Bad Limit",
			input: db.ListMenuByCityParams{
				Limit:     -1,
				Offset:    0,
				OfferedAt: offered,
				CityCode:  1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.ListMenuByCityParams{
					Limit:     -1,
					Offset:    0,
					OfferedAt: offered,
					CityCode:  1,
				}

				query.EXPECT().ListMenuByCity(ctx, arg).Times(1).Return([]db.Menu{}, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
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

			tc.buildStub(query)

			repo := NewMenuRepository(query)

			menus, err := repo.FetchByCity(ctx, tc.input.Limit, tc.input.Offset, tc.input.OfferedAt, tc.input.CityCode)

			tc.check(t, menus, err)
		})
	}
}

func randomMenuResults(length int) []db.Menu {
	var menus []db.Menu
	for i := 0; i < length; i++ {
		menus = append(menus, db.Menu{
			ID:                       util.NewUlid(),
			OfferedAt:                util.RandomDate(),
			CityCode:                 util.RandomCityCode(),
			PhotoUrl:                 util.RandomNullURL(),
			ElementarySchoolCalories: util.RandomInt32(),
			JuniorHighSchoolCalories: util.RandomInt32(),
		})
	}

	return menus
}
