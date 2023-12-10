package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ogurilab/school-lunch-api/util"
)

func TestGetByCityCode(t *testing.T) {
	code := util.RandomCityCode()

	type input struct {
		code int32
		ctx  context.Context
	}

	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockCityRepository)
		check     func(t *testing.T, city *domain.City, err error)
	}{
		{
			name: "OK",
			input: input{
				code: code,
				ctx:  context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				city := domain.NewCity(
					code,
					"city_name",
					1,
					"prefecture_name",
				)

				repo.EXPECT().GetByCityCode(gomock.Any(), gomock.Eq(code)).Times(1).Return(city, nil)
			},
			check: func(t *testing.T, city *domain.City, err error) {
				require.NoError(t, err)
				require.NotNil(t, city)

				require.Equal(t, code, city.CityCode)
				require.Equal(t, "city_name", city.CityName)
				require.Equal(t, int32(1), city.PrefectureCode)
				require.Equal(t, "prefecture_name", city.PrefectureName)
			},
		},
		{
			name: "Bad Code",
			input: input{
				code: -1,
				ctx:  context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				repo.EXPECT().GetByCityCode(gomock.Any(), gomock.Eq(int32(-1))).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, city *domain.City, err error) {
				require.Error(t, err)
				require.Nil(t, city)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockCityRepository(ctrl)
			tc.buildStub(repo)

			uc := NewCityUsecase(repo, 0)

			city, err := uc.GetByCityCode(tc.input.ctx, tc.input.code)

			tc.check(t, city, err)
		})
	}
}

func TestFetch(t *testing.T) {
	limit := util.RandomInt32()
	offset := util.RandomInt32()
	search := util.RandomString(10)

	type input struct {
		limit  int32
		offset int32
		search string
		ctx    context.Context
	}

	testCases := []struct {
		name      string
		input     input
		buildStub func(repo *mocks.MockCityRepository)
		check     func(t *testing.T, cities []*domain.City, err error)
	}{
		{
			name: "OK",
			input: input{
				limit:  limit,
				offset: offset,
				search: search,
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				cities := []*domain.City{
					domain.NewCity(
						1,
						"city_name",
						1,
						"prefecture_name",
					),
					domain.NewCity(
						2,
						"city_name",
						1,
						"prefecture_name",
					),
				}

				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(search)).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				require.NotNil(t, cities)

				require.Equal(t, 2, len(cities))

				require.Equal(t, int32(1), cities[0].CityCode)
				require.Equal(t, "city_name", cities[0].CityName)
				require.Equal(t, int32(1), cities[0].PrefectureCode)
				require.Equal(t, "prefecture_name", cities[0].PrefectureName)

				require.Equal(t, int32(2), cities[1].CityCode)
				require.Equal(t, "city_name", cities[1].CityName)
				require.Equal(t, int32(1), cities[1].PrefectureCode)
				require.Equal(t, "prefecture_name", cities[1].PrefectureName)
			},
		},
		{
			name: "Bad Code",
			input: input{
				limit:  -1,
				offset: -1,
				search: search,
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(-1)), gomock.Eq(int32(-1)), gomock.Eq(search)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.Error(t, err)
				require.Nil(t, cities)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockCityRepository(ctrl)
			tc.buildStub(repo)

			uc := NewCityUsecase(repo, 0)

			cities, err := uc.Fetch(tc.input.ctx, tc.input.limit, tc.input.offset, tc.input.search)

			tc.check(t, cities, err)
		})
	}
}

func TestFetchByPrefectureCode(t *testing.T) {
	limit := util.RandomInt32()
	offset := util.RandomInt32()
	prefectureCode := util.RandomInt32()

	type input struct {
		limit          int32
		offset         int32
		prefectureCode int32
		ctx            context.Context
	}

	testCases := []struct {
		name      string
		input     input
		buildStub func(repo *mocks.MockCityRepository)
		check     func(t *testing.T, cities []*domain.City, err error)
	}{
		{
			name: "OK",
			input: input{
				limit:          limit,
				offset:         offset,
				prefectureCode: prefectureCode,
				ctx:            context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				cities := []*domain.City{
					domain.NewCity(
						1,
						"city_name",
						1,
						"prefecture_name",
					),
					domain.NewCity(
						2,
						"city_name",
						1,
						"prefecture_name",
					),
				}

				repo.EXPECT().FetchByPrefectureCode(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(prefectureCode)).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				require.NotNil(t, cities)

				require.Equal(t, 2, len(cities))

				require.Equal(t, int32(1), cities[0].CityCode)
				require.Equal(t, "city_name", cities[0].CityName)
				require.Equal(t, int32(1), cities[0].PrefectureCode)
				require.Equal(t, "prefecture_name", cities[0].PrefectureName)

				require.Equal(t, int32(2), cities[1].CityCode)
				require.Equal(t, "city_name", cities[1].CityName)
				require.Equal(t, int32(1), cities[1].PrefectureCode)
				require.Equal(t, "prefecture_name", cities[1].PrefectureName)
			},
		},
		{
			name: "Bad Code",
			input: input{
				limit:          -1,
				offset:         -1,
				prefectureCode: prefectureCode,
				ctx:            context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				repo.EXPECT().FetchByPrefectureCode(gomock.Any(), gomock.Eq(int32(-1)), gomock.Eq(int32(-1)), gomock.Eq(prefectureCode)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.Error(t, err)
				require.Nil(t, cities)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockCityRepository(ctrl)
			tc.buildStub(repo)

			uc := NewCityUsecase(repo, 0)

			cities, err := uc.FetchByPrefectureCode(tc.input.ctx, tc.input.limit, tc.input.offset, tc.input.prefectureCode)

			tc.check(t, cities, err)
		})
	}
}
