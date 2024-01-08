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

func TestMenuCreate(t *testing.T) {
	time := time.Duration(10 * time.Second)
	type input struct {
		menu *domain.Menu
		ctx  context.Context
	}

	menu := randomMenu()
	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuRepository)
		check     func(t *testing.T, err error)
	}{
		{
			name: "OK",
			input: input{
				menu: menu,
				ctx:  context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				repo.EXPECT().Create(gomock.Any(), gomock.Eq(menu)).Times(1).Return(nil)
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

			repo := mocks.NewMockMenuRepository(ctrl)

			tc.buildStub(repo)

			uc := NewMenuUsecase(repo, time)

			err := uc.Create(tc.input.ctx, tc.input.menu)

			tc.check(t, err)
		})
	}
}

func TestMenuGetByID(t *testing.T) {
	time := time.Duration(10 * time.Second)
	type input struct {
		id   string
		city int32
		ctx  context.Context
	}

	menu := randomMenu()
	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuRepository)
		check     func(t *testing.T, menu *domain.Menu, err error)
	}{
		{
			name: "OK",
			input: input{
				id:   menu.ID,
				city: menu.CityCode,
				ctx:  context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				repo.EXPECT().GetByID(gomock.Any(), gomock.Eq(menu.ID), gomock.Eq(menu.CityCode)).Times(1).Return(menu, nil)
			},
			check: func(t *testing.T, menu *domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menu)

				require.Equal(t, menu.ID, menu.ID)
				require.Equal(t, menu.CityCode, menu.CityCode)
				require.Equal(t, menu.OfferedAt, menu.OfferedAt)
				require.Equal(t, menu.PhotoUrl, menu.PhotoUrl)
				require.Equal(t, menu.ElementarySchoolCalories, menu.ElementarySchoolCalories)
				require.Equal(t, menu.JuniorHighSchoolCalories, menu.JuniorHighSchoolCalories)
			},
		},
		{
			name: "Bad ID",
			input: input{
				id:   "bad_id",
				city: menu.CityCode,
				ctx:  context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				repo.EXPECT().GetByID(gomock.Any(), gomock.Eq("bad_id"), gomock.Eq(menu.CityCode)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menu *domain.Menu, err error) {
				require.Error(t, err)
				require.Nil(t, menu)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockMenuRepository(ctrl)

			tc.buildStub(repo)

			uc := NewMenuUsecase(repo, time)

			menu, err := uc.GetByID(tc.input.ctx, tc.input.id, tc.input.city)

			tc.check(t, menu, err)
		})
	}
}

func TestMenuFetch(t *testing.T) {
	time := time.Duration(10 * time.Second)
	type input struct {
		limit  int32
		offset int32
		city   int32
		ctx    context.Context
	}

	menu := randomMenu()
	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuRepository)
		check     func(t *testing.T, menus []*domain.Menu, err error)
	}{
		{
			name: "OK",
			input: input{
				limit:  10,
				offset: 0,
				city:   menu.CityCode,
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {

				var menus []*domain.Menu

				for i := 0; i < 10; i++ {
					menus = append(menus, menu)
				}

				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0)), gomock.Eq(menu.CityCode)).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 10)
			},
		},
		{
			name: "Bad City Code",
			input: input{
				limit:  10,
				offset: 0,
				city:   -1,
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0)), gomock.Eq(int32(-1))).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.Error(t, err)
				require.Nil(t, menus)
			},
		},
		{
			name: "Empty Result",
			input: input{
				limit:  10,
				offset: 0,
				city:   menu.CityCode,
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0)), gomock.Eq(menu.CityCode)).Times(1).Return([]*domain.Menu{}, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 0)
				require.Empty(t, menus)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockMenuRepository(ctrl)

			tc.buildStub(repo)

			uc := NewMenuUsecase(repo, time)

			menus, err := uc.Fetch(tc.input.ctx, tc.input.limit, tc.input.offset, tc.input.city)

			tc.check(t, menus, err)
		})
	}
}

func TestMenuGetByDate(t *testing.T) {
	ctxTime := time.Duration(10 * time.Second)

	type input struct {
		offeredAt time.Time
		city      int32
		ctx       context.Context
	}

	menu := randomMenu()
	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuRepository)
		check     func(t *testing.T, menu *domain.Menu, err error)
	}{
		{
			name: "OK",
			input: input{
				offeredAt: menu.OfferedAt,
				city:      menu.CityCode,
				ctx:       context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				repo.EXPECT().GetByDate(gomock.Any(), gomock.Eq(menu.OfferedAt), gomock.Eq(menu.CityCode)).Times(1).Return(menu, nil)
			},
			check: func(t *testing.T, menu *domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menu)

				require.Equal(t, menu.ID, menu.ID)
				require.Equal(t, menu.CityCode, menu.CityCode)
				require.Equal(t, menu.OfferedAt, menu.OfferedAt)
				require.Equal(t, menu.PhotoUrl, menu.PhotoUrl)
				require.Equal(t, menu.ElementarySchoolCalories, menu.ElementarySchoolCalories)
				require.Equal(t, menu.JuniorHighSchoolCalories, menu.JuniorHighSchoolCalories)
			},
		},
		{
			name: "Bad OfferedAt",
			input: input{
				offeredAt: time.Time{},
				city:      menu.CityCode,
				ctx:       context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				repo.EXPECT().GetByDate(gomock.Any(), gomock.Eq(time.Time{}), gomock.Eq(menu.CityCode)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menu *domain.Menu, err error) {
				require.Error(t, err)
				require.Nil(t, menu)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockMenuRepository(ctrl)

			tc.buildStub(repo)

			uc := NewMenuUsecase(repo, ctxTime)

			menu, err := uc.GetByDate(tc.input.ctx, tc.input.offeredAt, tc.input.city)

			tc.check(t, menu, err)
		})
	}
}

func TestMenuFetchByRangeDate(t *testing.T) {
	ctxTime := time.Duration(10 * time.Second)
	start := time.Now()
	end := time.Now().AddDate(0, 0, 1)

	type input struct {
		start time.Time
		end   time.Time
		city  int32
		limit int32
		ctx   context.Context
	}

	menu := randomMenu()
	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuRepository)
		check     func(t *testing.T, menus []*domain.Menu, err error)
	}{
		{
			name: "OK",
			input: input{
				start: start,
				end:   end,
				city:  menu.CityCode,
				limit: 10,
				ctx:   context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {

				var menus []*domain.Menu

				for i := 0; i < 10; i++ {
					menus = append(menus, menu)
				}

				repo.EXPECT().FetchByRangeDate(gomock.Any(), gomock.Eq(start), gomock.Eq(end), gomock.Eq(menu.CityCode), gomock.Eq(int32(10))).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 10)
			},
		},
		{
			name: "Bad City Code",
			input: input{
				start: start,
				end:   end,
				city:  -1,
				limit: 10,
				ctx:   context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				repo.EXPECT().FetchByRangeDate(gomock.Any(), gomock.Eq(start), gomock.Eq(end), gomock.Eq(int32(-1)), gomock.Eq(int32(10))).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.Error(t, err)
				require.Nil(t, menus)
			},
		},
		{
			name: "Empty Result",
			input: input{
				start: start,
				end:   end,
				city:  menu.CityCode,
				limit: 10,
				ctx:   context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {

				repo.EXPECT().FetchByRangeDate(gomock.Any(), gomock.Eq(start), gomock.Eq(end), gomock.Eq(menu.CityCode), gomock.Eq(int32(10))).Times(1).Return([]*domain.Menu{}, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 0)
				require.Empty(t, menus)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockMenuRepository(ctrl)

			tc.buildStub(repo)

			uc := NewMenuUsecase(repo, ctxTime)

			menus, err := uc.FetchByRangeDate(tc.input.ctx, tc.input.start, tc.input.end, tc.input.city, tc.input.limit)

			tc.check(t, menus, err)
		})
	}
}

/******************
 * MenuWithDishes *
 ******************/

func TestMenuWithDishesGetByID(t *testing.T) {
	ctxTime := time.Duration(10 * time.Second)
	type input struct {
		id   string
		city int32
		ctx  context.Context
	}

	menu := randomMenuWithDishes()
	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuWithDishesRepository)
		check     func(t *testing.T, menu *domain.MenuWithDishes, err error)
	}{
		{
			name: "OK",
			input: input{
				id:   menu.ID,
				city: menu.CityCode,
				ctx:  context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuWithDishesRepository) {
				repo.EXPECT().GetByID(gomock.Any(), gomock.Eq(menu.ID), gomock.Eq(menu.CityCode)).Times(1).Return(menu, nil)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menu)

				require.Equal(t, menu.ID, menu.ID)
				require.Equal(t, menu.CityCode, menu.CityCode)
				require.Equal(t, menu.OfferedAt, menu.OfferedAt)
				require.Equal(t, menu.PhotoUrl, menu.PhotoUrl)
				require.Equal(t, menu.ElementarySchoolCalories, menu.ElementarySchoolCalories)
				require.Equal(t, menu.JuniorHighSchoolCalories, menu.JuniorHighSchoolCalories)
			},
		},
		{
			name: "Bad ID",
			input: input{
				id:   "bad_id",
				city: menu.CityCode,
				ctx:  context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuWithDishesRepository) {
				repo.EXPECT().GetByID(gomock.Any(), gomock.Eq("bad_id"), gomock.Eq(menu.CityCode)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.Nil(t, menu)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockMenuWithDishesRepository(ctrl)

			tc.buildStub(repo)

			uc := NewMenuWithDishesUsecase(repo, ctxTime)

			menu, err := uc.GetByID(tc.input.ctx, tc.input.id, tc.input.city)

			tc.check(t, menu, err)
		})
	}
}

func TestMenuWithDishesFetch(t *testing.T) {
	ctxTime := time.Duration(10 * time.Second)
	type input struct {
		limit  int32
		offset int32
		city   int32
		ctx    context.Context
	}

	menu := randomMenuWithDishes()
	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuWithDishesRepository)
		check     func(t *testing.T, menus []*domain.MenuWithDishes, err error)
	}{
		{
			name: "OK",
			input: input{
				limit:  10,
				offset: 0,
				city:   menu.CityCode,
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuWithDishesRepository) {

				var menus []*domain.MenuWithDishes

				for i := 0; i < 10; i++ {
					menus = append(menus, menu)
				}

				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0)), gomock.Eq(menu.CityCode)).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 10)
			},
		},
		{
			name: "Bad City Code",
			input: input{
				limit:  10,
				offset: 0,
				city:   -1,
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuWithDishesRepository) {
				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0)), gomock.Eq(int32(-1))).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.Nil(t, menus)
			},
		},
		{
			name: "Empty Result",
			input: input{
				limit: 10,
				city:  menu.CityCode,
				ctx:   context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuWithDishesRepository) {
				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0)), gomock.Eq(menu.CityCode)).Times(1).Return([]*domain.MenuWithDishes{}, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 0)
				require.Empty(t, menus)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockMenuWithDishesRepository(ctrl)

			tc.buildStub(repo)

			uc := NewMenuWithDishesUsecase(repo, ctxTime)

			menus, err := uc.Fetch(tc.input.ctx, tc.input.limit, tc.input.offset, tc.input.city)

			tc.check(t, menus, err)
		})
	}
}

func TestMenuWithDishesGetByDate(t *testing.T) {
	ctxTime := time.Duration(10 * time.Second)
	type input struct {
		offeredAt time.Time
		city      int32
		ctx       context.Context
	}

	menu := randomMenuWithDishes()
	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuWithDishesRepository)
		check     func(t *testing.T, menu *domain.MenuWithDishes, err error)
	}{
		{
			name: "OK",
			input: input{
				offeredAt: menu.OfferedAt,
				city:      menu.CityCode,
				ctx:       context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuWithDishesRepository) {
				repo.EXPECT().GetByDate(gomock.Any(), gomock.Eq(menu.OfferedAt), gomock.Eq(menu.CityCode)).Times(1).Return(menu, nil)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menu)

				require.Equal(t, menu.ID, menu.ID)
				require.Equal(t, menu.CityCode, menu.CityCode)
				require.Equal(t, menu.OfferedAt, menu.OfferedAt)
				require.Equal(t, menu.PhotoUrl, menu.PhotoUrl)
				require.Equal(t, menu.ElementarySchoolCalories, menu.ElementarySchoolCalories)
				require.Equal(t, menu.JuniorHighSchoolCalories, menu.JuniorHighSchoolCalories)
			},
		},
		{
			name: "Bad OfferedAt",
			input: input{
				offeredAt: time.Time{},
				city:      menu.CityCode,
				ctx:       context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuWithDishesRepository) {
				repo.EXPECT().GetByDate(gomock.Any(), gomock.Eq(time.Time{}), gomock.Eq(menu.CityCode)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menu *domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.Nil(t, menu)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockMenuWithDishesRepository(ctrl)

			tc.buildStub(repo)

			uc := NewMenuWithDishesUsecase(repo, ctxTime)

			menu, err := uc.GetByDate(tc.input.ctx, tc.input.offeredAt, tc.input.city)

			tc.check(t, menu, err)
		})
	}
}

func TestMenuWithDishesFetchByRangeDate(t *testing.T) {
	ctxTime := time.Duration(10 * time.Second)
	start := time.Now()
	end := time.Now().AddDate(0, 0, 1)

	type input struct {
		start time.Time
		end   time.Time
		city  int32
		limit int32
		ctx   context.Context
	}

	menu := randomMenuWithDishes()
	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuWithDishesRepository)
		check     func(t *testing.T, menus []*domain.MenuWithDishes, err error)
	}{
		{
			name: "OK",
			input: input{
				start: start,
				end:   end,
				city:  menu.CityCode,
				limit: 10,
				ctx:   context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuWithDishesRepository) {

				var menus []*domain.MenuWithDishes

				for i := 0; i < 10; i++ {
					menus = append(menus, menu)
				}

				repo.EXPECT().FetchByRangeDate(gomock.Any(), gomock.Eq(start), gomock.Eq(end), gomock.Eq(menu.CityCode), gomock.Eq(int32(10))).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 10)
			},
		},
		{
			name: "Bad City Code",
			input: input{
				start: start,
				end:   end,
				city:  -1,
				limit: 10,
				ctx:   context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuWithDishesRepository) {
				repo.EXPECT().FetchByRangeDate(gomock.Any(), gomock.Eq(start), gomock.Eq(end), gomock.Eq(int32(-1)), gomock.Eq(int32(10))).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.Error(t, err)
				require.Nil(t, menus)
			},
		},
		{
			name: "Empty Result",
			input: input{
				start: start,
				end:   end,
				city:  menu.CityCode,
				limit: 10,
				ctx:   context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuWithDishesRepository) {

				repo.EXPECT().FetchByRangeDate(gomock.Any(), gomock.Eq(start), gomock.Eq(end), gomock.Eq(menu.CityCode), gomock.Eq(int32(10))).Times(1).Return([]*domain.MenuWithDishes{}, nil)
			},
			check: func(t *testing.T, menus []*domain.MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 0)
				require.Empty(t, menus)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockMenuWithDishesRepository(ctrl)

			tc.buildStub(repo)

			uc := NewMenuWithDishesUsecase(repo, ctxTime)

			menus, err := uc.FetchByRangeDate(tc.input.ctx, tc.input.start, tc.input.end, tc.input.city, tc.input.limit)

			tc.check(t, menus, err)
		})
	}
}

func randomMenu() *domain.Menu {
	menu, _ := domain.NewMenu(
		util.RandomDate(),
		util.RandomNullURL(),
		util.RandomInt32(),
		util.RandomInt32(),
		util.RandomCityCode(),
	)

	return menu
}

func randomDish() *domain.Dish {
	dish, _ := domain.NewDish(
		util.RandomUlid(),
		util.RandomString(10),
	)

	return dish
}

func randomMenuWithDishes() *domain.MenuWithDishes {
	var dishes []*domain.Dish

	for i := 0; i < 10; i++ {
		dishes = append(dishes, randomDish())
	}
	menu, _ := domain.ReNewMenuWithDishes(
		util.RandomUlid(),
		util.RandomDate(),
		util.RandomNullURL(),
		util.RandomInt32(),
		util.RandomInt32(),
		util.RandomCityCode(),
		dishes,
	)

	return menu
}
