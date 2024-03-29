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

func TestCreateMenu(t *testing.T) {
	time := time.Duration(10 * time.Second)
	type input struct {
		menu *domain.Menu
		ctx  context.Context
	}

	menu := randomMenu(t)
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

func TestGetMenuByID(t *testing.T) {
	time := time.Duration(10 * time.Second)
	type input struct {
		id   string
		city int32
		ctx  context.Context
	}

	menu := randomMenu(t)
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

func TestFetchMenuByCity(t *testing.T) {
	ctxTime := time.Duration(10 * time.Second)

	type input struct {
		limit   int32
		offset  int32
		offered time.Time
		city    int32
		ctx     context.Context
	}

	menu := randomMenu(t)
	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuRepository)
		check     func(t *testing.T, menus []*domain.Menu, err error)
	}{
		{
			name: "OK",
			input: input{
				limit:   10,
				offset:  0,
				offered: menu.OfferedAt,
				city:    menu.CityCode,
				ctx:     context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {

				var menus []*domain.Menu

				for i := 0; i < 10; i++ {
					menus = append(menus, menu)
				}

				repo.EXPECT().FetchByCity(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0)), gomock.Eq(menu.OfferedAt), gomock.Eq(menu.CityCode)).Times(1).Return(menus, nil)
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
				limit:   10,
				offset:  0,
				offered: menu.OfferedAt,
				city:    -1,
				ctx:     context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				repo.EXPECT().FetchByCity(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0)), gomock.Eq(menu.OfferedAt), gomock.Eq(int32(-1))).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.Error(t, err)
				require.Nil(t, menus)
			},
		},
		{
			name: "Empty Result",
			input: input{
				limit:   10,
				offset:  0,
				offered: menu.OfferedAt,
				city:    menu.CityCode,
				ctx:     context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				repo.EXPECT().FetchByCity(gomock.Any(), gomock.Eq(int32(10)), gomock.Eq(int32(0)), gomock.Eq(menu.OfferedAt), gomock.Eq(menu.CityCode)).Times(1).Return([]*domain.Menu{}, nil)
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

			menus, err := uc.FetchByCity(tc.input.ctx, tc.input.limit, tc.input.offset, tc.input.offered, tc.input.city)

			tc.check(t, menus, err)
		})
	}
}

func TestFetchMenu(t *testing.T) {
	ctxTime := time.Duration(10 * time.Second)

	type input struct {
		limit   int32
		offset  int32
		offered time.Time
		ids     []string
		ctx     context.Context
	}

	menus := make([]*domain.Menu, 0, 10)

	for i := 0; i < 10; i++ {
		menus = append(menus, randomMenu(t))
	}

	ids := make([]string, 0, 10)

	for _, menu := range menus {
		ids = append(ids, menu.ID)
	}

	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockMenuRepository)
		check     func(t *testing.T, menus []*domain.Menu, err error)
	}{
		{
			name: "OK",
			input: input{
				limit:   10,
				offset:  0,
				offered: menus[0].OfferedAt,
				ids:     ids,
				ctx:     context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {

				arg := input{
					limit:   10,
					offset:  0,
					offered: menus[0].OfferedAt,
					ids:     ids,
					ctx:     context.Background(),
				}

				repo.EXPECT().FetchByIDs(gomock.Any(), gomock.Eq(arg.limit), gomock.Eq(arg.offset), gomock.Eq(arg.offered), gomock.Eq(arg.ids)).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 10)
			},
		},
		{
			name: "if ids is empty, call Fetch",
			input: input{
				limit:   10,
				offset:  0,
				offered: menus[0].OfferedAt,
				ids:     []string{},
				ctx:     context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				arg := input{
					limit:   10,
					offset:  0,
					offered: menus[0].OfferedAt,
					ids:     []string{},
					ctx:     context.Background(),
				}

				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(arg.limit), gomock.Eq(arg.offset), gomock.Eq(arg.offered)).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {},
		},
		{
			name: "Empty Result By FetchByIDs",
			input: input{
				limit:   10,
				offset:  0,
				offered: menus[0].OfferedAt,
				ids:     ids,
				ctx:     context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				arg := input{
					limit:   10,
					offset:  0,
					offered: menus[0].OfferedAt,
					ids:     ids,
					ctx:     context.Background(),
				}

				repo.EXPECT().FetchByIDs(gomock.Any(), gomock.Eq(arg.limit), gomock.Eq(arg.offset), gomock.Eq(arg.offered), gomock.Eq(arg.ids)).Times(1).Return([]*domain.Menu{}, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 0)
				require.Empty(t, menus)
			},
		},
		{
			name: "Empty Result By Fetch",
			input: input{
				limit:   10,
				offset:  0,
				offered: menus[0].OfferedAt,
				ids:     []string{},
				ctx:     context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				arg := input{
					limit:   10,
					offset:  0,
					offered: menus[0].OfferedAt,
					ids:     []string{},
					ctx:     context.Background(),
				}

				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(arg.limit), gomock.Eq(arg.offset), gomock.Eq(arg.offered)).Times(1).Return([]*domain.Menu{}, nil)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, menus)

				require.Len(t, menus, 0)
				require.Empty(t, menus)
			},
		},
		{

			name: "NG FetchByIDs",
			input: input{
				limit:   10,
				offset:  0,
				offered: menus[0].OfferedAt,
				ids:     ids,
				ctx:     context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				arg := input{
					limit:   10,
					offset:  0,
					offered: menus[0].OfferedAt,
					ids:     ids,
					ctx:     context.Background(),
				}

				repo.EXPECT().FetchByIDs(gomock.Any(), gomock.Eq(arg.limit), gomock.Eq(arg.offset), gomock.Eq(arg.offered), gomock.Eq(arg.ids)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.Error(t, err)
				require.Nil(t, menus)
			},
		},
		{
			name: "NG Fetch",
			input: input{
				limit:   10,
				offset:  0,
				offered: menus[0].OfferedAt,
				ids:     []string{},
				ctx:     context.Background(),
			},
			buildStub: func(repo *mocks.MockMenuRepository) {
				arg := input{
					limit:   10,
					offset:  0,
					offered: menus[0].OfferedAt,
					ids:     []string{},
					ctx:     context.Background(),
				}

				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(arg.limit), gomock.Eq(arg.offset), gomock.Eq(arg.offered)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, menus []*domain.Menu, err error) {
				require.Error(t, err)
				require.Nil(t, menus)
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

			menus, err := uc.Fetch(tc.input.ctx, tc.input.limit, tc.input.offset, tc.input.offered, tc.input.ids)

			tc.check(t, menus, err)
		})
	}

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
