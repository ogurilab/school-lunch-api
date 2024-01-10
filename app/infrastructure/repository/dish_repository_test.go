package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ogurilab/school-lunch-api/domain"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	"github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc/mocks"
	"go.uber.org/mock/gomock"

	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreateDish(t *testing.T) {
	dish := randomDish(t)
	ctx := context.Background()

	testCases := []struct {
		name       string
		input      *domain.Dish
		buildStubs func(query *mocks.MockQuery)
		check      func(err error)
	}{
		{
			name:  "OK",
			input: dish,
			buildStubs: func(query *mocks.MockQuery) {
				arg := db.CreateDishParams{
					ID:     dish.ID,
					MenuID: dish.MenuID,
					Name:   dish.Name,
				}
				query.EXPECT().CreateDish(gomock.Any(), gomock.Eq(arg)).Times(1).Return(nil)
			},
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name:  "NG",
			input: &domain.Dish{},
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().CreateDish(gomock.Any(), gomock.Any()).Times(1).Return(sql.ErrConnDone)
			},
			check: func(err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)

			tc.buildStubs(query)

			repo := NewDishRepository(query)

			err := repo.Create(ctx, tc.input)

			tc.check(err)
		})
	}
}

func TestGetDishByID(t *testing.T) {
	dish := randomDishResult(t)
	ctx := context.Background()

	testCases := []struct {
		name       string
		id         string
		buildStubs func(query *mocks.MockQuery)
		check      func(t *testing.T, dish *domain.Dish, err error)
	}{
		{
			name: "OK",
			id:   dish.ID,
			buildStubs: func(query *mocks.MockQuery) {

				query.EXPECT().GetDish(gomock.Any(), gomock.Eq(dish.ID)).Times(1).Return(dish, nil)
			},
			check: func(t *testing.T, dish *domain.Dish, err error) {
				require.NoError(t, err)
				requireDishResult(t, dish, dish)
			},
		},
		{
			name: "NG",
			id:   dish.ID,
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().GetDish(gomock.Any(), gomock.Any()).Times(1).Return(db.Dish{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, dish *domain.Dish, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)

			tc.buildStubs(query)

			repo := NewDishRepository(query)

			result, err := repo.GetByID(ctx, tc.id)

			tc.check(t, result, err)
		})
	}
}

func TestFetchDishByMenuID(t *testing.T) {

	dishes := radomDishResults(t)
	ctx := context.Background()

	testCases := []struct {
		name       string
		menuID     string
		buildStubs func(query *mocks.MockQuery)
		check      func(t *testing.T, dishes []*domain.Dish, err error)
	}{
		{
			name:   "OK",
			menuID: dishes[0].MenuID,
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().ListDishByMenuID(gomock.Any(), gomock.Eq(dishes[0].MenuID)).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, dishes []*domain.Dish, err error) {
				require.NoError(t, err)
				require.Len(t, dishes, len(dishes))
			},
		},
		{
			name:   "NG",
			menuID: dishes[0].MenuID,
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().ListDishByMenuID(gomock.Any(), gomock.Any()).Times(1).Return([]db.Dish{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, dishes []*domain.Dish, err error) {
				require.Error(t, err)
				require.Nil(t, dishes)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)

			tc.buildStubs(query)

			repo := NewDishRepository(query)

			result, err := repo.FetchByMenuID(ctx, tc.menuID)

			tc.check(t, result, err)
		})
	}
}

func TestFetchDishByName(t *testing.T) {

	dishes := radomDishResults(t)
	ctx := context.Background()

	type input struct {
		search string
		limit  int32
		offset int32
	}

	testCases := []struct {
		name       string
		input      input
		buildStubs func(query *mocks.MockQuery)
		check      func(t *testing.T, dishes []*domain.Dish, err error)
	}{
		{
			name: "OK",
			input: input{
				search: dishes[0].Name,
				limit:  10,
				offset: 0,
			},
			buildStubs: func(query *mocks.MockQuery) {
				arg := db.ListDishByNameParams{
					Name:   dishes[0].Name,
					Limit:  10,
					Offset: 0,
				}
				query.EXPECT().ListDishByName(gomock.Any(), gomock.Eq(arg)).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, dishes []*domain.Dish, err error) {
				require.NoError(t, err)
				require.Len(t, dishes, len(dishes))
			},
		},
		{
			name: "NG",
			input: input{
				search: dishes[0].Name,
				limit:  10,
				offset: 0,
			},
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().ListDishByName(gomock.Any(), gomock.Any()).Times(1).Return([]db.Dish{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, dishes []*domain.Dish, err error) {
				require.Error(t, err)
				require.Nil(t, dishes)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)

			tc.buildStubs(query)

			repo := NewDishRepository(query)

			result, err := repo.FetchByName(ctx, tc.input.search, tc.input.limit, tc.input.offset)

			tc.check(t, result, err)
		})
	}
}

func TestFetchDish(t *testing.T) {

	dishes := radomDishResults(t)
	ctx := context.Background()

	type input struct {
		limit  int32
		offset int32
	}

	testCases := []struct {
		name       string
		input      input
		buildStubs func(query *mocks.MockQuery)
		check      func(t *testing.T, dishes []*domain.Dish, err error)
	}{
		{
			name: "OK",
			input: input{
				limit:  10,
				offset: 0,
			},
			buildStubs: func(query *mocks.MockQuery) {
				arg := db.ListDishParams{
					Limit:  10,
					Offset: 0,
				}
				query.EXPECT().ListDish(gomock.Any(), gomock.Eq(arg)).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, dishes []*domain.Dish, err error) {
				require.NoError(t, err)
				require.Len(t, dishes, len(dishes))
			},
		},
		{
			name: "NG",
			input: input{
				limit:  10,
				offset: 0,
			},
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().ListDish(gomock.Any(), gomock.Any()).Times(1).Return([]db.Dish{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, dishes []*domain.Dish, err error) {
				require.Error(t, err)
				require.Nil(t, dishes)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)

			tc.buildStubs(query)

			repo := NewDishRepository(query)

			result, err := repo.Fetch(ctx, tc.input.limit, tc.input.offset)

			tc.check(t, result, err)
		})
	}
}

func randomDishResult(t *testing.T) db.Dish {
	dish := randomDish(t)

	return db.Dish{
		ID:     dish.ID,
		MenuID: dish.MenuID,
		Name:   dish.Name,
	}
}

func radomDishResults(t *testing.T) []db.Dish {
	var dishes []db.Dish
	for i := 0; i < 10; i++ {
		dishes = append(dishes, randomDishResult(t))
	}

	return dishes
}

func randomMenu(t *testing.T) *domain.Menu {
	menu, err := domain.NewMenu(
		util.RandomDate(),
		util.RandomNullURL(),
		util.RandomInt32(),
		util.RandomInt32(),
		util.RandomInt32(),
	)

	require.NoError(t, err)

	return menu
}

func randomDish(t *testing.T) *domain.Dish {
	menu := randomMenu(t)

	dish, err := domain.NewDish(
		menu.ID,
		util.RandomString(10),
	)

	require.NoError(t, err)

	return dish
}

func requireDishResult(t *testing.T, expected *domain.Dish, actual *domain.Dish) {
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.MenuID, actual.MenuID)
	require.Equal(t, expected.Name, actual.Name)
}
