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

func TestGetByCityCode(t *testing.T) {
	code := util.RandomCityCode()

	type input struct {
		code int32
		ctx  context.Context
	}

	testCases := []struct {
		name      string
		input     input
		buildStub func(query *mocks.MockQuery)
		check     func(t *testing.T, city *domain.City, err error)
	}{
		{
			name: "OK",
			input: input{
				code: code,
				ctx:  context.Background(),
			},
			buildStub: func(query *mocks.MockQuery) {

				city := db.City{
					CityCode:                 code,
					CityName:                 "city_name",
					PrefectureCode:           1,
					PrefectureName:           "prefecture_name",
					SchoolLunchInfoAvailable: false,
				}

				query.EXPECT().GetCity(gomock.Any(), gomock.Eq(code)).Times(1).Return(city, nil)
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
			buildStub: func(query *mocks.MockQuery) {
				query.EXPECT().GetCity(gomock.Any(), gomock.Any()).Return(db.City{}, sql.ErrNoRows)
			},
			check: func(t *testing.T, city *domain.City, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrNoRows)
				require.Nil(t, city)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.buildStub(query)

			repo := NewCityRepository(query)

			city, err := repo.GetByCityCode(tc.input.ctx, tc.input.code)
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
		buildStub func(query *mocks.MockQuery)
		check     func(t *testing.T, cities []*domain.City, err error)
	}{
		{
			name: "OK By Search",
			input: input{
				limit:  limit,
				offset: offset,
				search: search,
				ctx:    context.Background(),
			},
			buildStub: func(query *mocks.MockQuery) {

				cities := []db.City{
					{
						CityCode:                 util.RandomCityCode(),
						CityName:                 "city_name",
						PrefectureCode:           1,
						PrefectureName:           "prefecture_name",
						SchoolLunchInfoAvailable: false,
					},
					{
						CityCode:                 util.RandomCityCode(),
						CityName:                 "city_name",
						PrefectureCode:           1,
						PrefectureName:           "prefecture_name",
						SchoolLunchInfoAvailable: false,
					},
				}

				query.EXPECT().ListCitiesByName(gomock.Any(), gomock.Any()).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				require.NotNil(t, cities)

				require.Len(t, cities, 2)

				for _, city := range cities {
					require.Equal(t, "city_name", city.CityName)
					require.Equal(t, int32(1), city.PrefectureCode)
					require.Equal(t, "prefecture_name", city.PrefectureName)
				}
			},
		},
		{
			name: "Ok By No Search",
			input: input{
				limit:  limit,
				offset: offset,
				search: "",
				ctx:    context.Background(),
			},
			buildStub: func(query *mocks.MockQuery) {

				cities := []db.City{
					{
						CityCode:                 util.RandomCityCode(),
						CityName:                 "city_name",
						PrefectureCode:           1,
						PrefectureName:           "prefecture_name",
						SchoolLunchInfoAvailable: false,
					},
					{
						CityCode:                 util.RandomCityCode(),
						CityName:                 "city_name",
						PrefectureCode:           1,
						PrefectureName:           "prefecture_name",
						SchoolLunchInfoAvailable: false,
					},
				}

				query.EXPECT().ListCities(gomock.Any(), gomock.Any()).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				require.NotNil(t, cities)

				require.Len(t, cities, 2)

				for _, city := range cities {
					require.Equal(t, "city_name", city.CityName)
					require.Equal(t, int32(1), city.PrefectureCode)
					require.Equal(t, "prefecture_name", city.PrefectureName)
				}
			},
		},
		{
			name: "Bad Search",
			input: input{
				limit:  limit,
				offset: offset,
				search: search,
				ctx:    context.Background(),
			},
			buildStub: func(query *mocks.MockQuery) {

				query.EXPECT().ListCitiesByName(gomock.Any(), gomock.Any()).Return([]db.City{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrConnDone)
				require.Nil(t, cities)
			},
		},
		{
			name: "Bad No Search",
			input: input{
				limit:  limit,
				offset: offset,
				search: "",
				ctx:    context.Background(),
			},
			buildStub: func(query *mocks.MockQuery) {

				query.EXPECT().ListCities(gomock.Any(), gomock.Any()).Return([]db.City{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrConnDone)
				require.Nil(t, cities)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.buildStub(query)

			repo := NewCityRepository(query)

			cities, err := repo.Fetch(tc.input.ctx, tc.input.limit, tc.input.offset, tc.input.search)
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
		buildStub func(query *mocks.MockQuery)
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
			buildStub: func(query *mocks.MockQuery) {

				cities := []db.City{
					{
						CityCode:                 util.RandomCityCode(),
						CityName:                 "city_name",
						PrefectureCode:           prefectureCode,
						PrefectureName:           "prefecture_name",
						SchoolLunchInfoAvailable: false,
					},
					{
						CityCode:                 util.RandomCityCode(),
						CityName:                 "city_name",
						PrefectureCode:           prefectureCode,
						PrefectureName:           "prefecture_name",
						SchoolLunchInfoAvailable: false,
					},
				}

				query.EXPECT().ListCitiesByPrefecture(gomock.Any(), gomock.Any()).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.NoError(t, err)
				require.NotNil(t, cities)

				require.Len(t, cities, 2)

				for _, city := range cities {
					require.Equal(t, "city_name", city.CityName)
					require.Equal(t, prefectureCode, city.PrefectureCode)
					require.Equal(t, "prefecture_name", city.PrefectureName)
				}
			},
		},
		{
			name: "Bad",
			input: input{
				limit:          limit,
				offset:         offset,
				prefectureCode: prefectureCode,
				ctx:            context.Background(),
			},
			buildStub: func(query *mocks.MockQuery) {

				query.EXPECT().ListCitiesByPrefecture(gomock.Any(), gomock.Any()).Return([]db.City{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, cities []*domain.City, err error) {
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrConnDone)
				require.Nil(t, cities)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.buildStub(query)

			repo := NewCityRepository(query)

			cities, err := repo.FetchByPrefectureCode(tc.input.ctx, tc.input.limit, tc.input.offset, tc.input.prefectureCode)
			tc.check(t, cities, err)
		})
	}
}
