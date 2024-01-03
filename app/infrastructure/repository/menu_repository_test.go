package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc/mocks"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	menu := &domain.Menu{
		ID:                       util.NewUlid(),
		OfferedAt:                util.RandomDate(),
		CityCode:                 util.RandomCityCode(),
		PhotoUrl:                 util.RandomNullURL(),
		ElementarySchoolCalories: util.RandomInt32(),
		JuniorHighSchoolCalories: util.RandomInt32(),
	}

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
func TestFetchMenu(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name      string
		input     db.ListMenusParams
		buildStub func(query *mocks.MockQuery)
		check     func(t *testing.T, menus []*domain.Menu, err error)
	}{
		{
			name: "OK",
			input: db.ListMenusParams{
				Limit:    10,
				Offset:   0,
				CityCode: 1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.ListMenusParams{
					Limit:    10,
					Offset:   0,
					CityCode: 1,
				}
				results := randomResults(10)
				query.EXPECT().ListMenus(ctx, arg).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)
				require.Len(t, menus, 10)
			},
		},
		{
			name: "Bad Limit",
			input: db.ListMenusParams{
				Limit:    -1,
				Offset:   0,
				CityCode: 1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.ListMenusParams{
					Limit:    -1,
					Offset:   0,
					CityCode: 1,
				}

				query.EXPECT().ListMenus(ctx, arg).Times(1).Return([]db.Menu{}, sql.ErrNoRows)
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

			menus, err := repo.Fetch(ctx, tc.input.Limit, tc.input.Offset, tc.input.CityCode)

			tc.check(t, menus, err)
		})
	}
}

func TestGetByDate(t *testing.T) {

	date := util.RandomDate()

	testCases := []struct {
		name      string
		input     db.GetMenuByOfferedAtParams
		buildStub func(query *mocks.MockQuery)
		check     func(t *testing.T, menu *domain.Menu, err error)
	}{
		{
			name: "OK",
			input: db.GetMenuByOfferedAtParams{
				OfferedAt: date,
				CityCode:  1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.GetMenuByOfferedAtParams{
					OfferedAt: date,
					CityCode:  1,
				}
				result := db.Menu{
					ID:        util.NewUlid(),
					OfferedAt: date,
					CityCode:  arg.CityCode,
					PhotoUrl: sql.NullString{
						String: "https://example.com",
						Valid:  true,
					},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
				}
				query.EXPECT().GetMenuByOfferedAt(context.Background(), arg).Times(1).Return(result, nil)
			},
			check: func(t *testing.T, menu *domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menu)
				require.Equal(t, int32(1), menu.CityCode)
				require.Equal(t, date, menu.OfferedAt)
				require.Equal(t, "https://example.com", menu.PhotoUrl.String)
				require.Equal(t, true, menu.PhotoUrl.Valid)
				require.Equal(t, int32(100), menu.ElementarySchoolCalories)
				require.Equal(t, int32(200), menu.JuniorHighSchoolCalories)

			},
		},
		{
			name: "Bad OfferedAt",
			input: db.GetMenuByOfferedAtParams{
				OfferedAt: time.Time{},
				CityCode:  1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.GetMenuByOfferedAtParams{
					OfferedAt: time.Time{},
					CityCode:  1,
				}
				query.EXPECT().GetMenuByOfferedAt(context.Background(), arg).Times(1).Return(db.Menu{}, sql.ErrNoRows)
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

			menu, err := repo.GetByDate(context.Background(), tc.input.OfferedAt, tc.input.CityCode)

			tc.check(t, menu, err)

		})
	}
}

func TestFetchByRangeDate(t *testing.T) {
	ctx := context.Background()
	startDate := util.RandomDate()
	endDate := startDate.AddDate(0, 0, 10)

	testCases := []struct {
		name      string
		input     db.ListMenusByOfferedAtParams
		buildStub func(query *mocks.MockQuery)
		check     func(t *testing.T, menus []*domain.Menu, err error)
	}{
		{
			name: "OK",
			input: db.ListMenusByOfferedAtParams{
				StartOfferedAt: startDate,
				EndOfferedAt:   endDate,
				CityCode:       1,
				Limit:          10,
			},

			buildStub: func(query *mocks.MockQuery) {
				arg := db.ListMenusByOfferedAtParams{
					StartOfferedAt: startDate,
					EndOfferedAt:   endDate,
					CityCode:       1,
					Limit:          10,
				}
				results := randomResults(10)
				query.EXPECT().ListMenusByOfferedAt(ctx, arg).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)
				require.Len(t, menus, 10)
			},
		},
		{
			name: "Bad Limit",
			input: db.ListMenusByOfferedAtParams{
				StartOfferedAt: startDate,
				EndOfferedAt:   endDate,
				CityCode:       1,
				Limit:          -1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.ListMenusByOfferedAtParams{
					StartOfferedAt: startDate,
					EndOfferedAt:   endDate,
					CityCode:       1,
					Limit:          -1,
				}
				query.EXPECT().ListMenusByOfferedAt(ctx, arg).Times(1).Return([]db.Menu{}, sql.ErrNoRows)
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

			menus, err := repo.FetchByRangeDate(ctx, tc.input.StartOfferedAt, tc.input.EndOfferedAt, tc.input.CityCode, tc.input.Limit)

			tc.check(t, menus, err)
		})
	}
}

func TestGetByIDWithDishes(t *testing.T) {
	id := util.NewUlid()
	date := util.RandomDate()
	dishID := util.NewUlid()

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

				dishes := randomDish(dishID)

				result := db.GetMenuWithDishesRow{
					ID:        arg.ID,
					OfferedAt: date,
					CityCode:  arg.CityCode,
					PhotoUrl: sql.NullString{
						String: "https://example.com",
						Valid:  true,
					},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					Dishes:                   dishes,
				}
				query.EXPECT().GetMenuWithDishes(context.Background(), arg).Times(1).Return(result, nil)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menu)
				require.Equal(t, int32(1), menu.CityCode)
				require.Equal(t, date, menu.OfferedAt)
				require.Equal(t, "https://example.com", menu.PhotoUrl.String)
				require.Equal(t, true, menu.PhotoUrl.Valid)
				require.Equal(t, int32(100), menu.ElementarySchoolCalories)
				require.Equal(t, int32(200), menu.JuniorHighSchoolCalories)
				require.NotNil(t, menu.Dishes)
				require.Len(t, menu.Dishes, 1)

				for _, dish := range menu.Dishes {
					require.Equal(t, dishID, dish.ID)
					require.Equal(t, "dish", dish.Name)
					require.Equal(t, "1", dish.MenuID)
				}
			},
		},
		{
			name: "Bad Dishes",
			input: db.GetMenuWithDishesParams{
				ID:       id,
				CityCode: 1,
			},

			buildStub: func(query *mocks.MockQuery) {
				arg := db.GetMenuWithDishesParams{
					ID:       id,
					CityCode: 1,
				}

				result := db.GetMenuWithDishesRow{
					ID:        arg.ID,
					OfferedAt: date,
					CityCode:  arg.CityCode,
					PhotoUrl: sql.NullString{
						String: "https://example.com",
						Valid:  true,
					},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					Dishes:                   json.RawMessage(""),
				}
				query.EXPECT().GetMenuWithDishes(context.Background(), arg).Times(1).Return(result, nil)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.Nil(t, menu)
			},
		},
		{
			name: "No Dishes",
			input: db.GetMenuWithDishesParams{
				ID:       id,
				CityCode: 1,
			},
			buildStub: func(query *mocks.MockQuery) {
				arg := db.GetMenuWithDishesParams{
					ID:       id,
					CityCode: 1,
				}

				result := db.GetMenuWithDishesRow{
					ID:        arg.ID,
					OfferedAt: date,
					CityCode:  arg.CityCode,
					PhotoUrl: sql.NullString{
						String: "https://example.com",
						Valid:  true,
					},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					Dishes:                   json.RawMessage("[]"),
				}
				query.EXPECT().GetMenuWithDishes(context.Background(), arg).Times(1).Return(result, nil)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menu)
				require.Equal(t, int32(1), menu.CityCode)
				require.Equal(t, date, menu.OfferedAt)
				require.Equal(t, "https://example.com", menu.PhotoUrl.String)
				require.Equal(t, true, menu.PhotoUrl.Valid)
				require.Equal(t, int32(100), menu.ElementarySchoolCalories)
				require.Equal(t, int32(200), menu.JuniorHighSchoolCalories)
				require.NotNil(t, menu.Dishes)
				require.Len(t, menu.Dishes, 0)
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
				query.EXPECT().GetMenuWithDishes(context.Background(), arg).Times(1).Return(db.GetMenuWithDishesRow{}, sql.ErrNoRows)
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

			repo := NewMenuRepository(query)

			menu, err := repo.GetByIDWithDishes(context.Background(), tc.input.ID, tc.input.CityCode)

			tc.check(t, menu, err)
		})
	}
}

func TestFetchWithDishes(t *testing.T) {

	testCases := []struct {
		name  string
		input db.ListMenuWithDishesParams
		build func(query *mocks.MockQuery)
		check func(t *testing.T, menus []*domain.MenuWithDishes, err error)
	}{
		{
			name: "OK",
			input: db.ListMenuWithDishesParams{
				Limit:    10,
				Offset:   0,
				CityCode: 1,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesParams{
					Limit:    10,
					Offset:   0,
					CityCode: 1,
				}
				results := randomWithDishesResult(10)
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
				Limit:    -1,
				Offset:   0,
				CityCode: 1,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesParams{
					Limit:    -1,
					Offset:   0,
					CityCode: 1,
				}
				query.EXPECT().ListMenuWithDishes(context.Background(), arg).Times(1).Return([]db.ListMenuWithDishesRow{}, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrNoRows)
				require.Nil(t, menus)
			},
		},
		{
			name: "Bad Dishes",
			input: db.ListMenuWithDishesParams{
				Limit:    10,
				Offset:   0,
				CityCode: 1,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesParams{
					Limit:    10,
					Offset:   0,
					CityCode: 1,
				}
				results := randomWithDishesResult(10)
				results[0].Dishes = json.RawMessage("")

				query.EXPECT().ListMenuWithDishes(context.Background(), arg).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.Nil(t, menus)
			},
		},
		{
			name: "No Dishes",
			input: db.ListMenuWithDishesParams{
				Limit:    10,
				Offset:   0,
				CityCode: 1,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesParams{
					Limit:    10,
					Offset:   0,
					CityCode: 1,
				}
				var results []db.ListMenuWithDishesRow

				for i := 0; i < 10; i++ {
					results = append(results, db.ListMenuWithDishesRow{
						ID:                       util.NewUlid(),
						OfferedAt:                util.RandomDate(),
						CityCode:                 util.RandomCityCode(),
						PhotoUrl:                 util.RandomNullURL(),
						ElementarySchoolCalories: util.RandomInt32(),
						JuniorHighSchoolCalories: util.RandomInt32(),
						Dishes:                   json.RawMessage("[]"),
					})
				}
				query.EXPECT().ListMenuWithDishes(context.Background(), arg).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)
				require.Len(t, menus, 10)
				for _, menu := range menus {
					require.Empty(t, menu.Dishes)
					require.Len(t, menu.Dishes, 0)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.build(query)

			repo := NewMenuRepository(query)

			menus, err := repo.FetchWithDishes(context.Background(), tc.input.Limit, tc.input.Offset, tc.input.CityCode)

			tc.check(t, menus, err)
		})
	}
}

func TestGetByDateWithDishes(t *testing.T) {
	date := util.RandomDate()
	dishID := util.NewUlid()

	testCases := []struct {
		name  string
		input db.GetMenuWithDishesByOfferedAtParams
		build func(query *mocks.MockQuery)
		check func(t *testing.T, menu *domain.MenuWithDishes, err error)
	}{
		{
			name: "OK",
			input: db.GetMenuWithDishesByOfferedAtParams{
				OfferedAt: date,
				CityCode:  1,
			},
			build: func(query *mocks.MockQuery) {

				arg := db.GetMenuWithDishesByOfferedAtParams{
					OfferedAt: date,
					CityCode:  1,
				}

				dishes := randomDish(dishID)

				result := db.GetMenuWithDishesByOfferedAtRow{
					ID:        util.NewUlid(),
					OfferedAt: date,
					CityCode:  arg.CityCode,
					PhotoUrl: sql.NullString{
						String: "https://example.com",
						Valid:  true,
					},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					Dishes:                   dishes,
				}
				query.EXPECT().GetMenuWithDishesByOfferedAt(context.Background(), arg).Times(1).Return(result, nil)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menu)
				require.Equal(t, int32(1), menu.CityCode)
				require.Equal(t, date, menu.OfferedAt)
				require.Equal(t, "https://example.com", menu.PhotoUrl.String)
				require.Equal(t, true, menu.PhotoUrl.Valid)
				require.Equal(t, int32(100), menu.ElementarySchoolCalories)
				require.Equal(t, int32(200), menu.JuniorHighSchoolCalories)
				require.NotNil(t, menu.Dishes)
				require.Len(t, menu.Dishes, 1)

				for _, dish := range menu.Dishes {
					require.Equal(t, dishID, dish.ID)
					require.Equal(t, "dish", dish.Name)
					require.Equal(t, "1", dish.MenuID)
				}
			},
		},
		{
			name: "Bad Dishes",
			input: db.GetMenuWithDishesByOfferedAtParams{
				OfferedAt: date,
				CityCode:  1,
			},
			build: func(query *mocks.MockQuery) {

				arg := db.GetMenuWithDishesByOfferedAtParams{
					OfferedAt: date,
					CityCode:  1,
				}

				result := db.GetMenuWithDishesByOfferedAtRow{
					ID:        util.NewUlid(),
					OfferedAt: date,
					CityCode:  arg.CityCode,
					PhotoUrl: sql.NullString{
						String: "https://example.com",
						Valid:  true,
					},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					Dishes:                   json.RawMessage(""),
				}
				query.EXPECT().GetMenuWithDishesByOfferedAt(context.Background(), arg).Times(1).Return(result, nil)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.Nil(t, menu)
			},
		},
		{
			name: "No Dishes",
			input: db.GetMenuWithDishesByOfferedAtParams{
				OfferedAt: date,
				CityCode:  1,
			},
			build: func(query *mocks.MockQuery) {

				arg := db.GetMenuWithDishesByOfferedAtParams{
					OfferedAt: date,
					CityCode:  1,
				}

				result := db.GetMenuWithDishesByOfferedAtRow{
					ID:        util.NewUlid(),
					OfferedAt: date,
					CityCode:  arg.CityCode,
					PhotoUrl: sql.NullString{
						String: "https://example.com",
						Valid:  true,
					},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					Dishes:                   json.RawMessage("[]"),
				}
				query.EXPECT().GetMenuWithDishesByOfferedAt(context.Background(), arg).Times(1).Return(result, nil)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menu)
				require.Equal(t, int32(1), menu.CityCode)
				require.Equal(t, date, menu.OfferedAt)
				require.Equal(t, "https://example.com", menu.PhotoUrl.String)
				require.Equal(t, true, menu.PhotoUrl.Valid)
				require.Equal(t, int32(100), menu.ElementarySchoolCalories)
				require.Equal(t, int32(200), menu.JuniorHighSchoolCalories)
				require.NotNil(t, menu.Dishes)
				require.Len(t, menu.Dishes, 0)
			},
		},
		{
			name: "Bad OfferedAt",
			input: db.GetMenuWithDishesByOfferedAtParams{
				OfferedAt: time.Time{},
				CityCode:  1,
			},
			build: func(query *mocks.MockQuery) {

				arg := db.GetMenuWithDishesByOfferedAtParams{
					OfferedAt: time.Time{},
					CityCode:  1,
				}

				query.EXPECT().GetMenuWithDishesByOfferedAt(context.Background(), arg).Times(1).Return(db.GetMenuWithDishesByOfferedAtRow{}, sql.ErrNoRows)
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
			tc.build(query)

			repo := NewMenuRepository(query)

			menu, err := repo.GetByDateWithDishes(context.Background(), tc.input.OfferedAt, tc.input.CityCode)

			tc.check(t, menu, err)
		})
	}
}

func TestFetchByRangeDateWithDishes(t *testing.T) {
	ctx := context.Background()
	startDate := util.RandomDate()
	endDate := startDate.AddDate(0, 0, 10)

	testCases := []struct {
		name  string
		input db.ListMenuWithDishesByOfferedAtParams
		build func(query *mocks.MockQuery)
		check func(t *testing.T, menus []*domain.MenuWithDishes, err error)
	}{
		{
			name: "OK",
			input: db.ListMenuWithDishesByOfferedAtParams{
				StartOfferedAt: startDate,
				EndOfferedAt:   endDate,
				CityCode:       1,
				Limit:          10,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesByOfferedAtParams{
					StartOfferedAt: startDate,
					EndOfferedAt:   endDate,
					CityCode:       1,
					Limit:          10,
				}
				results := randomWithDishesResultByOfferedAt(10)
				query.EXPECT().ListMenuWithDishesByOfferedAt(ctx, arg).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)
				require.Len(t, menus, 10)
			},
		},
		{
			name: "Bad Limit",
			input: db.ListMenuWithDishesByOfferedAtParams{
				StartOfferedAt: startDate,
				EndOfferedAt:   endDate,
				CityCode:       1,
				Limit:          -1,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesByOfferedAtParams{
					StartOfferedAt: startDate,
					EndOfferedAt:   endDate,
					CityCode:       1,
					Limit:          -1,
				}
				query.EXPECT().ListMenuWithDishesByOfferedAt(ctx, arg).Times(1).Return([]db.ListMenuWithDishesByOfferedAtRow{}, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrNoRows)
				require.Nil(t, menus)
			},
		},
		{
			name: "Bad Dishes",
			input: db.ListMenuWithDishesByOfferedAtParams{
				StartOfferedAt: startDate,
				EndOfferedAt:   endDate,
				CityCode:       1,
				Limit:          10,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesByOfferedAtParams{
					StartOfferedAt: startDate,
					EndOfferedAt:   endDate,
					CityCode:       1,
					Limit:          10,
				}
				var results []db.ListMenuWithDishesByOfferedAtRow

				for i := 0; i < 10; i++ {
					results = append(results, db.ListMenuWithDishesByOfferedAtRow{
						ID:                       util.NewUlid(),
						OfferedAt:                util.RandomDate(),
						CityCode:                 util.RandomCityCode(),
						PhotoUrl:                 util.RandomNullURL(),
						ElementarySchoolCalories: util.RandomInt32(),
						JuniorHighSchoolCalories: util.RandomInt32(),
						Dishes:                   json.RawMessage(""),
					})
				}

				query.EXPECT().ListMenuWithDishesByOfferedAt(ctx, arg).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.Nil(t, menus)
			},
		},
		{
			name: "No Dishes",

			input: db.ListMenuWithDishesByOfferedAtParams{
				StartOfferedAt: startDate,
				EndOfferedAt:   endDate,
				CityCode:       1,
				Limit:          10,
			},
			build: func(query *mocks.MockQuery) {
				arg := db.ListMenuWithDishesByOfferedAtParams{
					StartOfferedAt: startDate,
					EndOfferedAt:   endDate,
					CityCode:       1,
					Limit:          10,
				}
				var results []db.ListMenuWithDishesByOfferedAtRow

				for i := 0; i < 10; i++ {
					results = append(results, db.ListMenuWithDishesByOfferedAtRow{
						ID:                       util.NewUlid(),
						OfferedAt:                util.RandomDate(),
						CityCode:                 util.RandomCityCode(),
						PhotoUrl:                 util.RandomNullURL(),
						ElementarySchoolCalories: util.RandomInt32(),
						JuniorHighSchoolCalories: util.RandomInt32(),
						Dishes:                   json.RawMessage("[]"),
					})
				}
				query.EXPECT().ListMenuWithDishesByOfferedAt(ctx, arg).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)
				require.Len(t, menus, 10)
				for _, menu := range menus {
					require.Empty(t, menu.Dishes)
					require.Len(t, menu.Dishes, 0)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.build(query)

			repo := NewMenuRepository(query)

			menus, err := repo.FetchByRangeDateWithDishes(ctx, tc.input.StartOfferedAt, tc.input.EndOfferedAt, tc.input.CityCode, tc.input.Limit)

			tc.check(t, menus, err)
		})
	}
}

func randomResults(length int) []db.Menu {
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

func randomDish(id string) json.RawMessage {
	dishes := fmt.Sprintf(`[{"id":"%s","name":"dish","menu_id":"%d"}]`, id, 1)
	return json.RawMessage(dishes)

}

func randomWithDishesResult(length int) []db.ListMenuWithDishesRow {
	var results []db.ListMenuWithDishesRow

	for i := 0; i < length; i++ {
		results = append(results, db.ListMenuWithDishesRow{
			ID:                       util.NewUlid(),
			OfferedAt:                util.RandomDate(),
			CityCode:                 util.RandomCityCode(),
			PhotoUrl:                 util.RandomNullURL(),
			ElementarySchoolCalories: util.RandomInt32(),
			JuniorHighSchoolCalories: util.RandomInt32(),
			Dishes:                   randomDish(util.NewUlid()),
		})
	}

	return results
}

func randomWithDishesResultByOfferedAt(length int) []db.ListMenuWithDishesByOfferedAtRow {
	var results []db.ListMenuWithDishesByOfferedAtRow

	for i := 0; i < length; i++ {
		results = append(results, db.ListMenuWithDishesByOfferedAtRow{
			ID:                       util.NewUlid(),
			OfferedAt:                util.RandomDate(),
			CityCode:                 util.RandomCityCode(),
			PhotoUrl:                 util.RandomNullURL(),
			ElementarySchoolCalories: util.RandomInt32(),
			JuniorHighSchoolCalories: util.RandomInt32(),
			Dishes:                   randomDish(util.NewUlid()),
		})
	}

	return results
}
