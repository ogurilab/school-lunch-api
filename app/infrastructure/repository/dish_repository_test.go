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
	menuID := util.RandomString(10)

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
				arg, err := domain.ReNewDish(dish.ID, dish.Name)
				require.NoError(t, err)
				query.EXPECT().CreateDishTx(gomock.Any(), gomock.Eq(arg), gomock.Eq(menuID)).Times(1).Return(nil)
			},
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name:  "NG",
			input: &domain.Dish{},
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().CreateDishTx(gomock.Any(), gomock.Any(), gomock.Eq(menuID)).Times(1).Return(sql.ErrConnDone)
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

			err := repo.Create(ctx, tc.input, menuID)

			tc.check(err)
		})
	}
}

func TestGetDishByID(t *testing.T) {
	results := randomGetDishRowResults(t, 10)
	ctx := context.Background()

	testCases := []struct {
		name       string
		arg        db.GetDishParams
		buildStubs func(query *mocks.MockQuery)
		check      func(t *testing.T, dishes *domain.DishWithMenuIDs, err error)
	}{
		{
			name: "OK",
			arg: db.GetDishParams{
				ID:     results[0].ID,
				Limit:  10,
				Offset: 0,
			},
			buildStubs: func(query *mocks.MockQuery) {

				arg := db.GetDishParams{
					ID:     results[0].ID,
					Limit:  10,
					Offset: 0,
				}

				query.EXPECT().GetDish(gomock.Any(), gomock.Eq(arg)).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, dishes *domain.DishWithMenuIDs, err error) {
				require.NoError(t, err)
				requireGetDishWithMenuIDsResult(t, results, dishes)

			},
		},
		{
			name: "NG",
			arg: db.GetDishParams{
				ID:     results[0].ID,
				Limit:  10,
				Offset: 0,
			},
			buildStubs: func(query *mocks.MockQuery) {

				arg := db.GetDishParams{
					ID:     results[0].ID,
					Limit:  10,
					Offset: 0,
				}

				query.EXPECT().GetDish(gomock.Any(), gomock.Eq(arg)).Times(1).Return([]db.GetDishRow{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, dish *domain.DishWithMenuIDs, err error) {
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

			result, err := repo.GetByID(ctx, tc.arg.ID, tc.arg.Limit, tc.arg.Offset)

			tc.check(t, result, err)
		})
	}
}

func TestGetDishByIDInCity(t *testing.T) {
	cityCode := util.RandomInt32()
	results := randomGetDishInCityRowResults(t, 10)
	ctx := context.Background()

	testCases := []struct {
		name       string
		arg        db.GetDishInCityParams
		buildStubs func(query *mocks.MockQuery)
		check      func(t *testing.T, dishes *domain.DishWithMenuIDs, err error)
	}{
		{
			name: "OK",
			arg: db.GetDishInCityParams{
				ID:       results[0].ID,
				Limit:    10,
				Offset:   0,
				CityCode: cityCode,
			},
			buildStubs: func(query *mocks.MockQuery) {

				arg := db.GetDishInCityParams{
					ID:       results[0].ID,
					Limit:    10,
					Offset:   0,
					CityCode: cityCode,
				}

				query.EXPECT().GetDishInCity(gomock.Any(), gomock.Eq(arg)).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, dishes *domain.DishWithMenuIDs, err error) {
				require.NoError(t, err)

				requireGetInCityDishWithMenuIDsResult(t, results, dishes)
			},
		},
		{
			name: "NG",
			arg: db.GetDishInCityParams{
				ID:       results[0].ID,
				Limit:    10,
				Offset:   0,
				CityCode: cityCode,
			},
			buildStubs: func(query *mocks.MockQuery) {

				arg := db.GetDishInCityParams{
					ID:       results[0].ID,
					Limit:    10,
					Offset:   0,
					CityCode: cityCode,
				}

				query.EXPECT().GetDishInCity(gomock.Any(), gomock.Eq(arg)).Times(1).Return([]db.GetDishInCityRow{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, dish *domain.DishWithMenuIDs, err error) {
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

			result, err := repo.GetByIdInCity(ctx, tc.arg.ID, tc.arg.Limit, tc.arg.Offset, tc.arg.CityCode)

			tc.check(t, result, err)
		})

	}
}

func TestFetchDishByMenuID(t *testing.T) {
	dishes := randomListDishByMenuIDRow(t, 10)
	menu := randomMenu(t)
	ctx := context.Background()

	testCases := []struct {
		name       string
		menuID     string
		buildStubs func(query *mocks.MockQuery)
		check      func(t *testing.T, dishes []*domain.Dish, err error)
	}{
		{
			name:   "OK",
			menuID: menu.ID,
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().ListDishByMenuID(gomock.Any(), gomock.Eq(menu.ID)).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, dishes []*domain.Dish, err error) {
				require.NoError(t, err)
				require.Len(t, dishes, len(dishes))
			},
		},
		{
			name:   "NG",
			menuID: menu.ID,
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().ListDishByMenuID(gomock.Any(), gomock.Any()).Times(1).Return([]db.ListDishByMenuIDRow{}, sql.ErrConnDone)
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

	dishes := randomListDishByNameRow(t, 10)
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
				query.EXPECT().ListDishByName(gomock.Any(), gomock.Any()).Times(1).Return([]db.ListDishByNameRow{}, sql.ErrConnDone)
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

	dishes := randomListDishRow(t, 10)
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
				query.EXPECT().ListDish(gomock.Any(), gomock.Any()).Times(1).Return([]db.ListDishRow{}, sql.ErrConnDone)
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
		ID:   dish.ID,
		Name: dish.Name,
	}
}

func randomListDishByNameRow(t *testing.T, length int) []db.ListDishByNameRow {

	dishes := make([]db.ListDishByNameRow, 0, length)

	for i := 0; i < length; i++ {
		d := randomDishResult(t)

		data := db.ListDishByNameRow{
			ID:   d.ID,
			Name: d.Name,
		}

		dishes = append(dishes, data)
	}

	return dishes
}

func randomListDishByMenuIDRow(t *testing.T, length int) []db.ListDishByMenuIDRow {

	dishes := make([]db.ListDishByMenuIDRow, 0, length)

	for i := 0; i < length; i++ {
		d := randomDishResult(t)

		data := db.ListDishByMenuIDRow{
			ID:   d.ID,
			Name: d.Name,
		}

		dishes = append(dishes, data)
	}

	return dishes
}

func randomListDishRow(t *testing.T, length int) []db.ListDishRow {

	dishes := make([]db.ListDishRow, 0, length)

	for i := 0; i < length; i++ {
		d := randomDishResult(t)

		data := db.ListDishRow{
			ID:   d.ID,
			Name: d.Name,
		}

		dishes = append(dishes, data)
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

	dish, err := domain.NewDish(
		util.RandomString(10),
	)

	require.NoError(t, err)

	return dish
}

func randomGetDishRowResults(t *testing.T, length int) []db.GetDishRow {
	dish := randomDishResult(t)

	results := make([]db.GetDishRow, 0, length)

	for i := 0; i < length; i++ {
		menu := randomMenu(t)
		results = append(results, db.GetDishRow{
			ID:     dish.ID,
			Name:   dish.Name,
			MenuID: menu.ID,
		})
	}

	return results
}

func randomGetDishInCityRowResults(t *testing.T, length int) []db.GetDishInCityRow {
	dish := randomDishResult(t)

	results := make([]db.GetDishInCityRow, 0, length)

	for i := 0; i < length; i++ {
		menu := randomMenu(t)
		results = append(results, db.GetDishInCityRow{
			ID:     dish.ID,
			Name:   dish.Name,
			MenuID: menu.ID,
		})
	}

	return results
}

func requireGetDishWithMenuIDsResult(t *testing.T, expected []db.GetDishRow, actual *domain.DishWithMenuIDs) {

	menuIDs := make([]string, 0, len(expected))

	for _, result := range expected {
		menuIDs = append(menuIDs, result.MenuID)
	}

	require.Equal(t, expected[0].ID, actual.ID)
	require.Equal(t, expected[0].Name, actual.Name)

	require.ElementsMatch(t, menuIDs, actual.MenuIDs)
}

func requireGetInCityDishWithMenuIDsResult(t *testing.T, expected []db.GetDishInCityRow, actual *domain.DishWithMenuIDs) {

	menuIDs := make([]string, 0, len(expected))

	for _, result := range expected {
		menuIDs = append(menuIDs, result.MenuID)
	}

	require.Equal(t, expected[0].ID, actual.ID)
	require.Equal(t, expected[0].Name, actual.Name)

	require.ElementsMatch(t, menuIDs, actual.MenuIDs)
}
