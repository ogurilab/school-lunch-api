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

func TestGetCityByCityCode(t *testing.T) {
	code := util.RandomCityCode()

	city := randomCity()

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
				repo.EXPECT().GetByCityCode(gomock.Any(), gomock.Eq(code)).Times(1).Return(city, nil)
			},
			check: func(t *testing.T, city *domain.City, err error) {
				require.NoError(t, err)
				requireCityResult(t, city, city)
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

func TestFetchCity(t *testing.T) {
	limit := util.RandomInt32()
	offset := util.RandomInt32()
	search := util.RandomString(10)

	var cities []*domain.City

	for i := 0; i < 10; i++ {
		cities = append(cities, randomCity())
	}

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
			name: "OK No Search",
			input: input{
				limit:  limit,
				offset: offset,
				search: "",
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset)).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				requireCityResults(t, cities, cities)
			},
		},
		{
			name: "Bad Code No Search",
			input: input{
				limit:  -1,
				offset: -1,
				search: "",
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(int32(-1)), gomock.Eq(int32(-1))).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.Error(t, err)
				require.Nil(t, cities)
			},
		},
		{
			name: "Empty No Search",
			input: input{
				limit:  limit,
				offset: offset,
				search: "",
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				repo.EXPECT().Fetch(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset)).Times(1).Return([]*domain.City{}, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				require.NotNil(t, cities)

				require.Equal(t, 0, len(cities))
				require.Empty(t, cities)
			},
		},
		{
			name: "OK With Search",
			input: input{
				limit:  limit,
				offset: offset,
				search: search,
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				like := "%" + search + "%"

				repo.EXPECT().FetchByName(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(like)).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				requireCityResults(t, cities, cities)
			},
		},
		{
			name: "Bad Code With Search",
			input: input{
				limit:  -1,
				offset: -1,
				search: search,
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				like := "%" + search + "%"

				repo.EXPECT().FetchByName(gomock.Any(), gomock.Eq(int32(-1)), gomock.Eq(int32(-1)), gomock.Eq(like)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.Error(t, err)
				require.Nil(t, cities)
			},
		},
		{
			name: "Empty With Search",
			input: input{
				limit:  limit,
				offset: offset,
				search: search,
				ctx:    context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				like := "%" + search + "%"

				repo.EXPECT().FetchByName(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(like)).Times(1).Return([]*domain.City{}, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				require.NotNil(t, cities)

				require.Equal(t, 0, len(cities))
				require.Empty(t, cities)
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

func TestFetchCityByPrefectureCode(t *testing.T) {
	limit := util.RandomInt32()
	offset := util.RandomInt32()
	prefectureCode := util.RandomInt32()

	var cities []*domain.City

	for i := 0; i < 10; i++ {
		cities = append(cities, randomCity())
	}

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

				repo.EXPECT().FetchByPrefectureCode(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(prefectureCode)).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				requireCityResults(t, cities, cities)
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
		{
			name: "Empty",
			input: input{
				limit:          limit,
				offset:         offset,
				prefectureCode: prefectureCode,
				ctx:            context.Background(),
			},
			buildStub: func(repo *mocks.MockCityRepository) {

				repo.EXPECT().FetchByPrefectureCode(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(prefectureCode)).Times(1).Return([]*domain.City{}, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				require.NotNil(t, cities)

				require.Equal(t, 0, len(cities))
				require.Empty(t, cities)
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

func requireCityResults(t *testing.T, cities, mockData []*domain.City) {
	require.NotNil(t, cities)
	require.Equal(t, len(mockData), len(cities))

}

func requireCityResult(t *testing.T, cities, mockData *domain.City) {
	require.NotNil(t, cities)

	require.Equal(t, mockData.CityCode, cities.CityCode)
	require.Equal(t, mockData.CityName, cities.CityName)
	require.Equal(t, mockData.PrefectureCode, cities.PrefectureCode)
	require.Equal(t, mockData.PrefectureName, cities.PrefectureName)
}

func randomCity() *domain.City {
	return domain.NewCity(
		util.RandomCityCode(),
		util.RandomString(10),
		util.RandomInt32(),
		util.RandomString(10),
	)
}
