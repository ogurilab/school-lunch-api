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
	dish := randomDishWithMenuIDs(t)

	type req struct {
		id     string
		limit  sql.NullInt32
		offset sql.NullInt32
	}

	testCases := []struct {
		name      string
		req       req
		buildStub func(du *mocks.MockDishUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs)
	}{
		{
			name: "OK",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: 10, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
				}
				du.EXPECT().GetByID(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32)).Times(1).Return(dish, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishWithMenuIDs(t, recorder.Body, dish)
			},
		},
		{
			name: "Bad Request",
			req: req{
				id:     "invalid",
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     "invalid",
					limit:  sql.NullInt32{Int32: 10, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
				}
				du.EXPECT().GetByID(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32)).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request - Limit",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: -1, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: -1, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
				}
				du.EXPECT().GetByID(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32)).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request - Offset",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: -1, Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: 10, Valid: true},
					offset: sql.NullInt32{Int32: -1, Valid: true},
				}
				du.EXPECT().GetByID(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32)).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "If limit is not set, it will be set to domain.DEFAULT_LIMIT",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: 0, Valid: false},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},

			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: domain.DEFAULT_LIMIT, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
				}
				du.EXPECT().GetByID(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32)).Times(1).Return(dish, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishWithMenuIDs(t, recorder.Body, dish)
			},
		},
		{
			name: "If offset is not set, it will be set to domain.DEFAULT_OFFSET",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: false},
			},
			buildStub: func(du *mocks.MockDishUsecase) {

				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: 10, Valid: true},
					offset: sql.NullInt32{Int32: domain.DEFAULT_OFFSET, Valid: true},
				}
				du.EXPECT().GetByID(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32)).Times(1).Return(dish, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishWithMenuIDs(t, recorder.Body, dish)
			},
		},
		{
			name: "Max Limit Error",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: domain.MAX_LIMIT + 1, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: domain.MAX_LIMIT + 1, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
				}
				du.EXPECT().GetByID(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32)).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: 10, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
				}
				du.EXPECT().GetByID(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32)).Times(1).Return(dish, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
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

		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/dishes/%s?%s", tc.req.id, q.Encode())
		req, err := http.NewRequest(http.MethodGet, url, nil)

		require.NoError(t, err)

		e := newSetUpTestServer()
		e.GET("/dishes/:id", NewDishController(du).GetByID)

		e.ServeHTTP(recorder, req)

		tc.check(t, recorder, dish)
	}
}

func TestGetDishByIDInCity(t *testing.T) {
	dish := randomDishWithMenuIDs(t)
	cityCode := util.RandomInt32()

	type req struct {
		id     string
		limit  sql.NullInt32
		offset sql.NullInt32
		city   int32
	}

	testCases := []struct {
		name      string
		req       req
		buildStub func(du *mocks.MockDishUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs)
	}{
		{
			name: "OK",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				city:   cityCode,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: 10, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
					city:   cityCode,
				}
				du.EXPECT().GetByIdInCity(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32), gomock.Eq(arg.city)).Times(1).Return(dish, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishWithMenuIDs(t, recorder.Body, dish)
			},
		},
		{
			name: "Bad Request",
			req: req{
				id:     "invalid",
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				city:   cityCode,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     "invalid",
					limit:  sql.NullInt32{Int32: 10, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
					city:   cityCode,
				}
				du.EXPECT().GetByIdInCity(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32), gomock.Eq(arg.city)).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request - Limit",
			req: req{
				id: dish.ID,

				limit:  sql.NullInt32{Int32: -1, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				city:   cityCode,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: -1, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
					city:   cityCode,
				}
				du.EXPECT().GetByIdInCity(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32), gomock.Eq(arg.city)).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request - Offset",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: -1, Valid: true},
				city:   cityCode,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: 10, Valid: true},
					offset: sql.NullInt32{Int32: -1, Valid: true},
					city:   cityCode,
				}
				du.EXPECT().GetByIdInCity(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32), gomock.Eq(arg.city)).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
		{
			name: "If limit is not set, it will be set to domain.DEFAULT_LIMIT",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: 0, Valid: false},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				city:   cityCode,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: domain.DEFAULT_LIMIT, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
					city:   cityCode,
				}
				du.EXPECT().GetByIdInCity(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32), gomock.Eq(arg.city)).Times(1).Return(dish, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishWithMenuIDs(t, recorder.Body, dish)
			},
		},
		{
			name: "If offset is not set, it will be set to domain.DEFAULT_OFFSET",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: false},
				city:   cityCode,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: 10, Valid: true},
					offset: sql.NullInt32{Int32: domain.DEFAULT_OFFSET, Valid: true},
					city:   cityCode,
				}
				du.EXPECT().GetByIdInCity(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32), gomock.Eq(arg.city)).Times(1).Return(dish, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchDishWithMenuIDs(t, recorder.Body, dish)
			},
		},
		{
			name: "Max Limit Error",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: domain.MAX_LIMIT + 1, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				city:   cityCode,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: domain.MAX_LIMIT + 1, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
					city:   cityCode,
				}
				du.EXPECT().GetByIdInCity(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32), gomock.Eq(arg.city)).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			req: req{
				id:     dish.ID,
				limit:  sql.NullInt32{Int32: 10, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				city:   cityCode,
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				arg := req{
					id:     dish.ID,
					limit:  sql.NullInt32{Int32: 10, Valid: true},
					offset: sql.NullInt32{Int32: 0, Valid: true},
					city:   cityCode,
				}
				du.EXPECT().GetByIdInCity(gomock.Any(), gomock.Eq(arg.id), gomock.Eq(arg.limit.Int32), gomock.Eq(arg.offset.Int32), gomock.Eq(arg.city)).Times(1).Return(dish, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, dish *domain.DishWithMenuIDs) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
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

		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/cities/%d/dishes/%s?%s", tc.req.city, tc.req.id, q.Encode())
		req, err := http.NewRequest(http.MethodGet, url, nil)

		require.NoError(t, err)

		e := newSetUpTestServer()
		e.GET("/cities/:code/dishes/:id", NewDishController(du).GetByIdInCity)

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
			name: "Max Limit Error",
			req: req{
				limit:  sql.NullInt32{Int32: domain.MAX_LIMIT + 1, Valid: true},
				offset: sql.NullInt32{Int32: 0, Valid: true},
				search: sql.NullString{String: "dish", Valid: true},
			},
			buildStub: func(du *mocks.MockDishUsecase) {
				du.EXPECT().Fetch(gomock.Any(), gomock.Eq("dish"), gomock.Eq(int32(domain.MAX_LIMIT+1)), gomock.Eq(int32(0))).Times(0)
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

func requireBodyMatchDishWithMenuIDs(t *testing.T, body *bytes.Buffer, dish *domain.DishWithMenuIDs) {
	data, err := io.ReadAll(body)

	require.NoError(t, err)

	var dishData domain.DishWithMenuIDs

	err = json.Unmarshal(data, &dishData)

	require.NoError(t, err)

	require.Equal(t, dish.ID, dishData.ID)
	require.Equal(t, dish.Name, dishData.Name)
	require.ElementsMatch(t, dish.MenuIDs, dishData.MenuIDs)
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

func randomDishWithMenuIDs(t *testing.T) *domain.DishWithMenuIDs {
	dish := randomDish(t)

	n := 10

	menuIDs := make([]string, 0, n)

	for i := 0; i < n; i++ {
		menuIDs = append(menuIDs, util.RandomUlid())
	}

	dishes, err := domain.ReNewDishWithMenuIDs(
		dish.ID,
		dish.Name,
		menuIDs,
	)

	require.NoError(t, err)

	return dishes
}
