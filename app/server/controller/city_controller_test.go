package controller

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/mocks"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetByCityCode(t *testing.T) {

	ctx := echo.New().NewContext(nil, nil)
	city := randomCity()

	type args struct {
		c    echo.Context
		code int32
	}

	tests := []struct {
		name      string
		args      args
		buildStub func(uc *mocks.MockCityUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			args: args{
				c:    ctx,
				code: city.CityCode,
			},
			buildStub: func(uc *mocks.MockCityUsecase) {

				uc.EXPECT().GetByCityCode(context.Background(), city.CityCode).Return(city, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCity(t, recorder.Body, city)
			},
		},
		{
			name: "Bad Request",
			args: args{
				c:    ctx,
				code: int32(-1),
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().GetByCityCode(context.Background(), -1).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Not Found",
			args: args{
				c:    ctx,
				code: city.CityCode,
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().GetByCityCode(context.Background(), city.CityCode).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			args: args{
				c:    ctx,
				code: city.CityCode,
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().GetByCityCode(context.Background(), city.CityCode).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mocks.NewMockCityUsecase(ctrl)
			tc.buildStub(uc)

			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/cities/%d", tc.args.code)
			req, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			e := newSetUpTestServer()
			e.GET("/cities/:code", NewCityController(uc).GetByCityCode)
			e.ServeHTTP(recorder, req)

			tc.check(t, recorder)
		})
	}

}

func TestFetch(t *testing.T) {
	ctx := echo.New().NewContext(nil, nil)
	var cities []*domain.City

	limit := 10
	for i := 0; i < limit; i++ {
		cities = append(cities, randomCity())
	}

	type query struct {
		limit  sql.NullInt32
		offset sql.NullInt32
		search sql.NullString
	}

	tests := []struct {
		name      string
		ctx       echo.Context
		query     query
		buildStub func(uc *mocks.MockCityUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK No Search",
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: int32(limit), Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				search: sql.NullString{String: "", Valid: false},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().Fetch(context.Background(), gomock.Eq(int32(limit)), gomock.Eq(int32(0)), gomock.Eq("")).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCities(t, recorder.Body, cities)
			},
		},
		{
			name: "OK No Search No query params",
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: 0, Valid: false},
				offset: sql.NullInt32{Int32: 0, Valid: false},
				search: sql.NullString{String: "", Valid: false},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().Fetch(context.Background(), gomock.Eq(domain.DEFAULT_LIMIT), gomock.Eq(domain.DEFAULT_OFFSET), gomock.Eq("")).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCities(t, recorder.Body, cities)
			},
		},
		{
			name: "OK With Search",
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: int32(limit), Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				search: sql.NullString{String: cities[0].CityName, Valid: true},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {

				city := cities[0]
				uc.EXPECT().Fetch(context.Background(), gomock.Eq(int32(limit)), gomock.Eq(int32(0)), gomock.Eq(city.CityName)).Times(1).Return([]*domain.City{city}, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCities(t, recorder.Body, []*domain.City{cities[0]})
			},
		},
		{
			name: "OK With Search No query params",
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: 0, Valid: false},
				offset: sql.NullInt32{Int32: 0, Valid: false},
				search: sql.NullString{String: cities[0].CityName, Valid: true},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {

				city := cities[0]
				uc.EXPECT().Fetch(context.Background(), gomock.Eq(domain.DEFAULT_LIMIT), gomock.Eq(domain.DEFAULT_OFFSET), gomock.Eq(city.CityName)).Times(1).Return([]*domain.City{city}, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCities(t, recorder.Body, []*domain.City{cities[0]})
			},
		},
		{
			name: "Bad Request Invalid Offset",
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: int32(limit), Valid: true},
				offset: sql.NullInt32{Int32: int32(-1), Valid: true},
				search: sql.NullString{String: "", Valid: false},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().Fetch(context.Background(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request Invalid Limit",
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: int32(-1), Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				search: sql.NullString{String: "", Valid: false},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().Fetch(context.Background(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Not Found is Empty Array",
			ctx:  ctx,
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().Fetch(context.Background(), gomock.Eq(int32(limit)), gomock.Eq(int32(0)), gomock.Eq("")).Times(1).Return([]*domain.City{}, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCities(t, recorder.Body, []*domain.City{})
			},
		},
		{
			name: "Max Limit Error",
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: int32(domain.MAX_LIMIT + 1), Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				search: sql.NullString{String: "", Valid: false},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().Fetch(context.Background(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			ctx:  ctx,
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().Fetch(context.Background(), gomock.Eq(int32(limit)), gomock.Eq(int32(0)), gomock.Eq("")).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mocks.NewMockCityUsecase(ctrl)
			tc.buildStub(uc)

			recorder := httptest.NewRecorder()

			q := make(url.Values)
			if tc.query.limit.Valid {
				q.Set("limit", fmt.Sprintf("%d", tc.query.limit.Int32))
			}

			if tc.query.offset.Valid {
				q.Set("offset", fmt.Sprintf("%d", tc.query.offset.Int32))
			}

			if tc.query.search.Valid {
				q.Set("search", tc.query.search.String)
			}

			req, err := http.NewRequest(http.MethodGet, "/cities?"+q.Encode(), nil)

			require.NoError(t, err)

			e := newSetUpTestServer()
			e.GET("/cities", NewCityController(uc).Fetch)

			e.ServeHTTP(recorder, req)

			tc.check(t, recorder)
		})
	}

}

func TestFetchByPrefectureCode(t *testing.T) {
	ctx := echo.New().NewContext(nil, nil)
	var cities []*domain.City

	limit := 10
	for i := 0; i < limit; i++ {
		cities = append(cities, randomCity())
	}

	type query struct {
		limit  sql.NullInt32
		offset sql.NullInt32
	}

	tests := []struct {
		name      string
		code      int32
		ctx       echo.Context
		query     query
		buildStub func(uc *mocks.MockCityUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK No query params",
			code: cities[0].PrefectureCode,
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: 0, Valid: false},
				offset: sql.NullInt32{Int32: 0, Valid: false},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().FetchByPrefectureCode(context.Background(), gomock.Eq(domain.DEFAULT_LIMIT), gomock.Eq(domain.DEFAULT_OFFSET), gomock.Eq(cities[0].PrefectureCode)).Times(1).Return(cities, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCities(t, recorder.Body, cities)
			},
		},
		{
			name: "OK",
			code: cities[0].PrefectureCode,
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: int32(limit), Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {

				city := cities[0]
				uc.EXPECT().FetchByPrefectureCode(context.Background(), gomock.Eq(int32(limit)), gomock.Eq(int32(0)), gomock.Eq(city.PrefectureCode)).Times(1).Return([]*domain.City{city}, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCities(t, recorder.Body, []*domain.City{cities[0]})
			},
		},
		{
			name: "Bad Request Invalid Offset",
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: int32(limit), Valid: true},
				offset: sql.NullInt32{Int32: int32(-1), Valid: true},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().FetchByPrefectureCode(context.Background(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request Invalid Limit",
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: int32(-1), Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().FetchByPrefectureCode(context.Background(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Not Found is Empty Array",
			ctx:  ctx,
			code: cities[0].PrefectureCode,
			query: query{
				limit:  sql.NullInt32{Int32: int32(limit), Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().FetchByPrefectureCode(context.Background(), gomock.Eq(int32(limit)), gomock.Eq(int32(0)), gomock.Eq(cities[0].PrefectureCode)).Times(1).Return([]*domain.City{}, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCities(t, recorder.Body, []*domain.City{})
			},
		},
		{
			name: "Max Limit Error",
			ctx:  ctx,
			code: cities[0].PrefectureCode,
			query: query{
				limit:  sql.NullInt32{Int32: int32(domain.MAX_LIMIT + 1), Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},
			buildStub: func(uc *mocks.MockCityUsecase) {
				uc.EXPECT().FetchByPrefectureCode(context.Background(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			ctx:  ctx,
			query: query{
				limit:  sql.NullInt32{Int32: int32(limit), Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},
			code: cities[0].PrefectureCode,
			buildStub: func(uc *mocks.MockCityUsecase) {

				uc.EXPECT().FetchByPrefectureCode(context.Background(), gomock.Eq(int32(limit)), gomock.Eq(int32(0)), gomock.Eq(cities[0].PrefectureCode)).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mocks.NewMockCityUsecase(ctrl)
			tc.buildStub(uc)

			recorder := httptest.NewRecorder()

			q := make(url.Values)
			if tc.query.limit.Valid {
				q.Set("limit", fmt.Sprintf("%d", tc.query.limit.Int32))
			}

			if tc.query.offset.Valid {
				q.Set("offset", fmt.Sprintf("%d", tc.query.offset.Int32))
			}

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/cities/prefectures/%d", tc.code)+"?"+q.Encode(), nil)

			require.NoError(t, err)

			e := newSetUpTestServer()
			e.GET("/cities/prefectures/:code", NewCityController(uc).FetchByPrefectureCode)

			e.ServeHTTP(recorder, req)

			tc.check(t, recorder)
		})
	}

}

func randomCity() *domain.City {
	return domain.NewCity(
		util.RandomCityCode(),
		util.RandomString(10),
		util.RandomInt32(),
		util.RandomString(10),
	)
}

func requireBodyMatchCities(t *testing.T, body *bytes.Buffer, cities []*domain.City) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var citiesData []*domain.City

	err = json.Unmarshal(data, &citiesData)
	require.NoError(t, err)
	require.Equal(t, cities, citiesData)
	require.Equal(t, len(cities), len(citiesData))
}

func requireBodyMatchCity(t *testing.T, body *bytes.Buffer, city *domain.City) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var cityData domain.City

	err = json.Unmarshal(data, &cityData)
	require.NoError(t, err)
	require.Equal(t, city.CityCode, cityData.CityCode)
	require.Equal(t, city.CityName, cityData.CityName)
	require.Equal(t, city.PrefectureCode, cityData.PrefectureCode)
	require.Equal(t, city.PrefectureName, cityData.PrefectureName)

}
