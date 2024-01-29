package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/mocks"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestFetchDishByMenuID(t *testing.T) {
	menuID := util.RandomUlid()
	var dishes []*domain.Dish

	for i := 0; i < 10; i++ {
		dishes = append(dishes, randomDish(t))
	}

	type req fetchDishByMenuIDRequest

	testCases := []struct {
		name      string
		req       req
		buildStub func(du *mocks.MockDishUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish)
	}{
		{
			name: "OK",
			req: req{
				MenuID: menuID,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().FetchByMenuID(gomock.Any(), menuID).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishes(t, recorder.Body, dishes)
			},
		},
		{
			name: "Bad Request",
			req: req{
				MenuID: "invalid",
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().FetchByMenuID(gomock.Any(), menuID).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			req: req{
				MenuID: menuID,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().FetchByMenuID(gomock.Any(), menuID).Times(1).Return([]*domain.Dish{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		du := mocks.NewMockDishUsecase(ctrl)
		tc.buildStub(du)

		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/menus/%s/dishes", tc.req.MenuID)
		req, err := http.NewRequest(http.MethodGet, url, nil)

		require.NoError(t, err)

		e := newSetUpTestServer()
		e.GET("/menus/:menuID/dishes", NewDishController(du).FetchByMenuID)

		e.ServeHTTP(recorder, req)

		tc.check(t, recorder, dishes)
	}
}

func TestGetDishByID(t *testing.T) {
	dish := randomDish(t)

	type req getDishRequest

	testCases := []struct {
		name      string
		req       req
		buildStub func(du *mocks.MockDishUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.Dish)
	}{
		{
			name: "OK",
			req: req{
				ID: dish.ID,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().GetByID(gomock.Any(), dish.ID).Times(1).Return(dish, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.Dish) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDish(t, recorder.Body, dish)
			},
		},
		{
			name: "Bad Request",
			req: req{
				ID: "invalid",
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().GetByID(gomock.Any(), dish.ID).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.Dish) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			req: req{
				ID: dish.ID,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().GetByID(gomock.Any(), dish.ID).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.Dish) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		du := mocks.NewMockDishUsecase(ctrl)
		tc.buildStub(du)

		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/dishes/%s", tc.req.ID)
		req, err := http.NewRequest(http.MethodGet, url, nil)

		require.NoError(t, err)

		e := newSetUpTestServer()
		e.GET("/dishes/:id", NewDishController(du).GetByID)

		e.ServeHTTP(recorder, req)

		tc.check(t, recorder, dish)
	}
}

func TestFetchDish(t *testing.T) {
	var dishes []*domain.Dish

	for i := 0; i < 10; i++ {
		dishes = append(dishes, randomDish(t))
	}

	type req struct {
		limit  sql.NullInt32
		offset sql.NullInt32
		search sql.NullString
	}

	testCases := []struct {
		name      string
		req       req
		buildStub func(du *mocks.MockDishUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish)
	}{
		{
			name: "OK",
			req: req{
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				search: sql.NullString{String: "dish", Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().Fetch(gomock.Any(), gomock.Eq("dish"), gomock.Eq(int32(10)), gomock.Eq(int32(0))).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishes(t, recorder.Body, dishes)
			},
		},
		{
			name: "If limit is not set, it will be set to domain.DEFAULT_LIMIT",
			req: req{
				limit:  sql.NullInt32{Int32: 0, Valid: false},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				search: sql.NullString{String: "dish", Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().Fetch(gomock.Any(), gomock.Eq("dish"), gomock.Eq(domain.DEFAULT_LIMIT), gomock.Eq(int32(0))).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishes(t, recorder.Body, dishes)
			},
		},
		{
			name: "If offset is not set, it will be set to domain.DEFAULT_OFFSET",
			req: req{
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: false},
				search: sql.NullString{String: "dish", Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().Fetch(gomock.Any(), gomock.Eq("dish"), gomock.Eq(int32(10)), gomock.Eq(domain.DEFAULT_OFFSET)).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishes(t, recorder.Body, dishes)
			},
		},
		{
			name: "If search is not set, it will be set to empty string",
			req: req{
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				search: sql.NullString{String: "", Valid: false},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().Fetch(gomock.Any(), gomock.Eq(""), gomock.Eq(int32(10)), gomock.Eq(int32(0))).Times(1).Return(dishes, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishes(t, recorder.Body, dishes)
			},
		},
		{
			name: "Bad Request - Limit",
			req: req{
				limit:  sql.NullInt32{Int32: -1, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				search: sql.NullString{String: "dish", Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().Fetch(gomock.Any(), gomock.Eq("dish"), gomock.Eq(int32(-1)), gomock.Eq(int32(0))).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request - Offset",
			req: req{
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: -1, Valid: true},
				search: sql.NullString{String: "dish", Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().Fetch(gomock.Any(), gomock.Eq("dish"), gomock.Eq(int32(-1)), gomock.Eq(int32(0))).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			req: req{
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				search: sql.NullString{String: "dish", Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().Fetch(gomock.Any(), gomock.Eq("dish"), gomock.Eq(int32(10)), gomock.Eq(int32(0))).Times(1).Return([]*domain.Dish{}, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dishes []*domain.Dish) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			du := mocks.NewMockDishUsecase(ctrl)
			tc.buildStub(du)

			q := make(url.Values)

			if tc.req.limit.Valid {
				q.Set("limit", fmt.Sprintf("%d", tc.req.limit.Int32))
			}

			if tc.req.offset.Valid {
				q.Set("offset", fmt.Sprintf("%d", tc.req.offset.Int32))
			}

			if tc.req.search.Valid {
				q.Set("search", tc.req.search.String)
			}

			url := fmt.Sprintf("/dishes?%s", q.Encode())
			req, err := http.NewRequest(http.MethodGet, url, nil)

			require.NoError(t, err)

			recorder := httptest.NewRecorder()
			e := newSetUpTestServer()
			e.GET("/dishes", NewDishController(du).Fetch)
			e.ServeHTTP(recorder, req)

			tc.check(t, recorder, dishes)
		})
	}
}

func requireBodyMatchDish(t *testing.T, body *bytes.Buffer, dish *domain.Dish) {
	data, err := io.ReadAll(body)

	require.NoError(t, err)

	var dishData domain.Dish

	err = json.Unmarshal(data, &dishData)

	require.NoError(t, err)

	require.Equal(t, dish.ID, dishData.ID)
	require.Equal(t, dish.Name, dishData.Name)
}

func requireBodyMatchDishes(t *testing.T, body *bytes.Buffer, dishes []*domain.Dish) {
	data, err := io.ReadAll(body)

	require.NoError(t, err)

	var dishesData []*domain.Dish

	err = json.Unmarshal(data, &dishesData)

	require.NoError(t, err)

	require.Equal(t, len(dishes), len(dishesData))
}

func randomDish(t *testing.T) *domain.Dish {
	dish, err := domain.NewDish(
		util.RandomUlid(),
	)

	require.NoError(t, err)

	return dish
}
