package usecase

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/mocks"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateDish(t *testing.T) {
	dish := randomDish(t)
	menu := randomMenu(t)
	timeout := time.Second * 10
	ctx := context.Background()

	testCases := []struct {
		name       string
		dish       *domain.Dish
		buildStubs func(r *mocks.MockDishRepository)
		check      func(err error)
	}{
		{
			name: "OK",
			dish: dish,
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().Create(gomock.Any(), gomock.Eq(dish), gomock.Eq(menu.ID)).Times(1).Return(nil)
			},
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "NG",
			dish: dish,
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().Create(gomock.Any(), gomock.Eq(dish), gomock.Eq(menu.ID)).Times(1).Return(sql.ErrNoRows)
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

			repo := mocks.NewMockDishRepository(ctrl)
			tc.buildStubs(repo)

			du := NewDishUsecase(repo, timeout)

			err := du.Create(ctx, tc.dish, menu.ID)

			tc.check(err)
		})
	}
}

func TestCreateManyDish(t *testing.T) {
	var dishes []*domain.Dish
	menu := randomMenu(t)

	for i := 0; i < 10; i++ {
		dishes = append(dishes, randomDish(t))
	}

	timeout := time.Second * 10
	ctx := context.Background()

	testCases := []struct {
		name       string
		dishes     []*domain.Dish
		buildStubs func(r *mocks.MockDishRepository)
		check      func(err error)
	}{
		{
			name:   "OK",
			dishes: dishes,
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().CreateMany(gomock.Any(), gomock.Eq(dishes), gomock.Eq(menu.ID)).Times(1).Return(nil)
			},
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name:   "NG",
			dishes: dishes,
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().CreateMany(gomock.Any(), gomock.Eq(dishes), gomock.Eq(menu.ID)).Times(1).Return(sql.ErrNoRows)
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

			repo := mocks.NewMockDishRepository(ctrl)
			tc.buildStubs(repo)

			du := NewDishUsecase(repo, timeout)

			err := du.CreateMany(ctx, tc.dishes, menu.ID)

			tc.check(err)
		})
	}
}

func TestGetDishByID(t *testing.T) {
	dish := randomDishWithMenuIDs(t)
	timeout := time.Second * 10
	ctx := context.Background()

	type input struct {
		limit  int32
		offset int32
		id     string
	}

	testCases := []struct {
		name       string
		input      input
		buildStubs func(r *mocks.MockDishRepository)
		check      func(result *domain.DishWithMenuIDs, err error)
	}{
		{
			name: "OK",
			input: input{
				limit:  10,
				offset: 0,
				id:     dish.ID,
			},
			buildStubs: func(r *mocks.MockDishRepository) {

				arg := input{
					limit:  10,
					offset: 0,
					id:     dish.ID,
				}

				r.EXPECT().GetByID(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit), gomock.Eq(arg.offset)).Times(1).Return(dish, nil)
			},
			check: func(result *domain.DishWithMenuIDs, err error) {
				require.NoError(t, err)
				requireDishWithMenuIDsResult(t, dish, result)
			},
		},
		{
			name: "NG",
			input: input{
				limit:  10,
				offset: 0,
				id:     dish.ID,
			},
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().GetByID(gomock.Any(), gomock.Eq(dish.ID), gomock.Eq(int32(10)), gomock.Eq(int32(0))).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(result *domain.DishWithMenuIDs, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockDishRepository(ctrl)
			tc.buildStubs(repo)

			du := NewDishUsecase(repo, timeout)

			result, err := du.GetByID(ctx, tc.input.id, tc.input.limit, tc.input.offset)

			tc.check(result, err)
		})
	}
}

func TestGetDishByIdInCity(t *testing.T) {
	cityCode := util.RandomCityCode()

	dish := randomDishWithMenuIDs(t)
	timeout := time.Second * 10
	ctx := context.Background()

	type input struct {
		limit  int32
		offset int32
		id     string
		city   int32
	}

	testCases := []struct {
		name       string
		input      input
		buildStubs func(r *mocks.MockDishRepository)
		check      func(result *domain.DishWithMenuIDs, err error)
	}{
		{
			name: "OK",
			input: input{
				limit:  10,
				offset: 0,
				id:     dish.ID,
				city:   cityCode,
			},
			buildStubs: func(r *mocks.MockDishRepository) {

				arg := input{
					limit:  10,
					offset: 0,
					id:     dish.ID,
					city:   cityCode,
				}

				r.EXPECT().GetByIdInCity(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit), gomock.Eq(arg.offset), gomock.Eq(arg.city)).Times(1).Return(dish, nil)
			},

			check: func(result *domain.DishWithMenuIDs, err error) {
				require.NoError(t, err)
				requireDishWithMenuIDsResult(t, dish, result)
			},
		},
		{
			name: "NG",
			input: input{
				limit:  10,
				offset: 0,
				id:     dish.ID,
				city:   cityCode,
			},
			buildStubs: func(r *mocks.MockDishRepository) {
				arg := input{
					limit:  10,
					offset: 0,
					id:     dish.ID,
					city:   cityCode,
				}
				r.EXPECT().GetByIdInCity(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit), gomock.Eq(arg.offset), gomock.Eq(arg.city)).Times(1).Return(nil, sql.ErrNoRows)
			},

			check: func(result *domain.DishWithMenuIDs, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockDishRepository(ctrl)
			tc.buildStubs(repo)

			du := NewDishUsecase(repo, timeout)

			result, err := du.GetByIdInCity(ctx, tc.input.id, tc.input.limit, tc.input.offset, tc.input.city)

			tc.check(result, err)
		})
	}

}

func TestFetchDishByMenuID(t *testing.T) {
	menuID := util.NewUlid()
	var dishes []*domain.Dish

	for i := 0; i < 10; i++ {
		dishes = append(dishes, randomDish(t))
	}

	timeout := time.Second * 10
	ctx := context.Background()

	testCases := []struct {
		name       string
		menuID     string
		buildStubs func(r *mocks.MockDishRepository)
		check      func(t *testing.T, result []*domain.Dish, err error)
	}{
		{
			name:   "OK",
			menuID: menuID,
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().FetchByMenuID(gomock.Any(), gomock.Eq(menuID)).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, result []*domain.Dish, err error) {
				require.NoError(t, err)
				require.Equal(t, len(dishes), len(result))
			},
		},
		{
			name:   "NG",
			menuID: menuID,
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().FetchByMenuID(gomock.Any(), gomock.Eq(menuID)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, result []*domain.Dish, err error) {
				require.Error(t, err)
			},
		},
		{
			name:   "Empty",
			menuID: menuID,
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().FetchByMenuID(gomock.Any(), gomock.Eq(menuID)).Times(1).Return([]*domain.Dish{}, nil)
			},
			check: func(t *testing.T, result []*domain.Dish, err error) {
				require.NoError(t, err)
				require.Equal(t, 0, len(result))
				require.Empty(t, result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockDishRepository(ctrl)
			tc.buildStubs(repo)

			du := NewDishUsecase(repo, timeout)

			result, err := du.FetchByMenuID(ctx, tc.menuID)

			tc.check(t, result, err)
		})
	}
}

func TestFetchDish(t *testing.T) {
	var dishes []*domain.Dish

	for i := 0; i < 10; i++ {
		dishes = append(dishes, randomDish(t))
	}

	timeout := time.Second * 10
	ctx := context.Background()

	type input struct {
		search string
		limit  int32
		offset int32
	}

	testCases := []struct {
		name       string
		input      input
		buildStubs func(r *mocks.MockDishRepository)
		check      func(t *testing.T, result []*domain.Dish, err error)
	}{
		{
			name: "OK with search",
			input: input{
				search: dishes[0].Name,
				limit:  10,
				offset: 0,
			},
			buildStubs: func(r *mocks.MockDishRepository) {
				like := "%" + dishes[0].Name + "%"
				r.EXPECT().FetchByName(gomock.Any(), gomock.Eq(like), gomock.Eq(int32(10)), gomock.Eq(int32(0))).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, result []*domain.Dish, err error) {
				require.NoError(t, err)
				require.Equal(t, len(dishes), len(result))
			},
		},
		{
			name: "OK without search",
			input: input{
				search: "",
				limit:  10,
				offset: 0,
			},
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0))).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, result []*domain.Dish, err error) {
				require.NoError(t, err)
				require.Equal(t, len(dishes), len(result))
			},
		},
		{
			name: "NG with search",
			input: input{
				search: dishes[0].Name,
				limit:  10,
				offset: 0,
			},
			buildStubs: func(r *mocks.MockDishRepository) {
				like := "%" + dishes[0].Name + "%"
				r.EXPECT().FetchByName(gomock.Any(), gomock.Eq(like), gomock.Eq(int32(10)), gomock.Eq(int32(0))).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, result []*domain.Dish, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "NG without search",
			input: input{
				search: "",
				limit:  10,
				offset: 0,
			},
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0))).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, result []*domain.Dish, err error) {
				require.Error(t, err)
			},
		},
		{
			name: "Empty with search",
			input: input{
				search: dishes[0].Name,
				limit:  10,
				offset: 0,
			},
			buildStubs: func(r *mocks.MockDishRepository) {
				like := "%" + dishes[0].Name + "%"
				r.EXPECT().FetchByName(gomock.Any(), gomock.Eq(like), gomock.Eq(int32(10)), gomock.Eq(int32(0))).Times(1).Return([]*domain.Dish{}, nil)
			},
			check: func(t *testing.T, result []*domain.Dish, err error) {
				require.NoError(t, err)
				require.Equal(t, 0, len(result))
				require.Empty(t, result)
			},
		},
		{
			name: "Empty without search",
			input: input{
				search: "",
				limit:  10,
				offset: 0,
			},
			buildStubs: func(r *mocks.MockDishRepository) {
				r.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0))).Times(1).Return([]*domain.Dish{}, nil)
			},
			check: func(t *testing.T, result []*domain.Dish, err error) {
				require.NoError(t, err)
				require.Equal(t, 0, len(result))
				require.Empty(t, result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockDishRepository(ctrl)
			tc.buildStubs(repo)

			du := NewDishUsecase(repo, timeout)

			result, err := du.Fetch(ctx, tc.input.search, tc.input.limit, tc.input.offset)

			tc.check(t, result, err)
		})
	}
}

func randomDish(t *testing.T) *domain.Dish {
	dish, err := domain.NewDish(
		util.RandomString(10),
	)

	require.NoError(t, err)

	return dish
}

func randomDishWithMenuIDs(t *testing.T) *domain.DishWithMenuIDs {
	dish := randomDish(t)

	n := 10

	menuIDs := make([]string, 0, n)

	for i := 0; i < n; i++ {
		menuIDs = append(menuIDs, util.NewUlid())
	}

	dishWithMenuIDs, err := domain.ReNewDishWithMenuIDs(
		dish.ID,
		dish.Name,
		menuIDs,
	)

	require.NoError(t, err)

	return dishWithMenuIDs
}

func requireDishWithMenuIDsResult(t *testing.T, expected *domain.DishWithMenuIDs, actual *domain.DishWithMenuIDs) {
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Name, actual.Name)
	require.ElementsMatch(t, expected.MenuIDs, actual.MenuIDs)
}
