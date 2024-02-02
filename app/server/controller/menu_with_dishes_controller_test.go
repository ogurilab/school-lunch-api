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

func TestGetMenuWithDishesByID(t *testing.T) {
	type req getMenuWithDishesRequest
	menu := randomMenuWithDishes(t)

	testCases := []struct {
		name      string
		req       req
		buildStub func(uc *mocks.MockMenuWithDishesUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.MenuWithDishes)
	}{
		{
			name: "OK",
			req: req{
				ID:       menu.ID,
				CityCode: menu.CityCode,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {

				uc.EXPECT().GetByID(gomock.Any(), gomock.Eq(menu.ID), gomock.Eq(menu.CityCode)).Times(1).Return(menu, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.MenuWithDishes) {
				require.Equal(t, 200, recorder.Code)
				requireBodyMatchMenuWithDishes(t, recorder.Body, menu)
			},
		},
		{
			name: "Bad Request - Invalid ID",
			req: req{
				ID:       "invalid-id",
				CityCode: menu.CityCode,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().GetByID(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid CityCode",
			req: req{
				ID:       menu.ID,
				CityCode: -1,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().GetByID(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "Not Found",
			req: req{
				ID:       menu.ID,
				CityCode: menu.CityCode,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().GetByID(gomock.Any(), gomock.Eq(menu.ID), gomock.Eq(menu.CityCode)).Times(1).Return(nil, sql.ErrNoRows)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.MenuWithDishes) {
				require.Equal(t, 404, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			req: req{
				ID:       menu.ID,
				CityCode: menu.CityCode,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().GetByID(gomock.Any(), gomock.Eq(menu.ID), gomock.Eq(menu.CityCode)).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.MenuWithDishes) {
				require.Equal(t, 500, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mocks.NewMockMenuWithDishesUsecase(ctrl)
		tc.buildStub(uc)

		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/cities/%d/menus/%s", tc.req.CityCode, tc.req.ID)
		req, err := http.NewRequest(http.MethodGet, url, nil)

		require.NoError(t, err)

		e := newSetUpTestServer()
		e.GET("/cities/:code/menus/:id", NewMenuWithDishesController(uc).GetByID)
		e.ServeHTTP(recorder, req)

		tc.check(t, recorder, menu)
	}
}

func TestFetchMenuWithDishes(t *testing.T) {
	limit := int32(10)
	offset := int32(0)
	var menus []*domain.MenuWithDishes

	for i := 0; i < int(limit); i++ {
		menus = append(menus, randomMenuWithDishes(t))
	}

	offered := menus[0].OfferedAt.Format("2006-01-02")

	type req struct {
		Limit   sql.NullInt32
		Offset  sql.NullInt32
		Offered string
	}

	testCases := []struct {
		name      string
		req       req
		buildStub func(uc *mocks.MockMenuWithDishesUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes)
	}{
		{
			name: "OK",
			req: req{
				Limit:   sql.NullInt32{Int32: limit, Valid: true},
				Offset:  sql.NullInt32{Int32: offset, Valid: true},
				Offered: offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				parsedOffered, err := util.ParseDate(offered)

				require.NoError(t, err)

				uc.EXPECT().Fetch(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(parsedOffered)).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 200, recorder.Code)
				requireBodyMatchMenuWithDishesList(t, recorder.Body, menus)
			},
		},
		{
			name: "Bad Request - Invalid Limit",
			req: req{
				Limit:   sql.NullInt32{Int32: -1, Valid: true},
				Offset:  sql.NullInt32{Int32: offset, Valid: true},
				Offered: offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid Offset",
			req: req{
				Limit:   sql.NullInt32{Int32: limit, Valid: true},
				Offset:  sql.NullInt32{Int32: -1, Valid: true},
				Offered: offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid Offered",
			req: req{
				Limit:   sql.NullInt32{Int32: limit, Valid: true},
				Offset:  sql.NullInt32{Int32: offset, Valid: true},
				Offered: "invalid-offered",
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().Fetch(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "If Offered is not specified, return 400 error",
			req: req{
				Limit:   sql.NullInt32{Int32: limit, Valid: true},
				Offset:  sql.NullInt32{Int32: offset, Valid: true},
				Offered: "",
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().Fetch(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Any()).Times(0).Return(menus, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "If Limit is not specified, it will be set to domain.DEFAULT_LIMIT",
			req: req{
				Limit:   sql.NullInt32{Int32: 0, Valid: false},
				Offset:  sql.NullInt32{Int32: offset, Valid: true},
				Offered: offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				parsedOffered, err := util.ParseDate(offered)

				require.NoError(t, err)
				uc.EXPECT().Fetch(gomock.Any(), gomock.Eq(domain.DEFAULT_LIMIT), gomock.Eq(offset), gomock.Eq(parsedOffered)).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 200, recorder.Code)
				requireBodyMatchMenuWithDishesList(t, recorder.Body, menus)
			},
		},
		{
			name: "If Offset is not specified, it will be set to domain.DEFAULT_OFFSET",
			req: req{
				Limit:   sql.NullInt32{Int32: limit, Valid: true},
				Offset:  sql.NullInt32{Int32: 0, Valid: false},
				Offered: offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				parsedOffered, err := util.ParseDate(offered)

				require.NoError(t, err)
				uc.EXPECT().Fetch(gomock.Any(), gomock.Eq(limit), gomock.Eq(domain.DEFAULT_OFFSET), gomock.Eq(parsedOffered)).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 200, recorder.Code)
				requireBodyMatchMenuWithDishesList(t, recorder.Body, menus)
			},
		},
		{
			name: "Internal Server Error",
			req: req{
				Limit:   sql.NullInt32{Int32: limit, Valid: true},
				Offset:  sql.NullInt32{Int32: offset, Valid: true},
				Offered: offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				parsedOffered, err := util.ParseDate(offered)

				require.NoError(t, err)
				uc.EXPECT().Fetch(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(parsedOffered)).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 500, recorder.Code)
			},
		},
		{
			name: "Empty Result",
			req: req{
				Limit:   sql.NullInt32{Int32: limit, Valid: true},
				Offset:  sql.NullInt32{Int32: offset, Valid: true},
				Offered: offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				parsedOffered, err := util.ParseDate(offered)

				require.NoError(t, err)
				uc.EXPECT().Fetch(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(parsedOffered)).Times(1).Return([]*domain.MenuWithDishes{}, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 200, recorder.Code)
				data, err := io.ReadAll(recorder.Body)

				require.NoError(t, err)

				var res fetchMenuWithDishesResponse
				err = json.Unmarshal(data, &res)

				require.NoError(t, err)

				var menuData []*domain.MenuWithDishes

				menuData = append(menuData, res.Menus...)

				require.Empty(t, menuData)
				require.Equal(t, res.Next, "")
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mocks.NewMockMenuWithDishesUsecase(ctrl)
		tc.buildStub(uc)

		q := make(url.Values)
		if tc.req.Limit.Valid {
			q.Set("limit", fmt.Sprintf("%d", tc.req.Limit.Int32))
		}

		if tc.req.Offset.Valid {
			q.Set("offset", fmt.Sprintf("%d", tc.req.Offset.Int32))
		}

		q.Set("offered", tc.req.Offered)

		url := fmt.Sprintf("/menus?%s", q.Encode())
		req, err := http.NewRequest(http.MethodGet, url, nil)

		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		e := newSetUpTestServer()
		e.GET("/menus", NewMenuWithDishesController(uc).Fetch)
		e.ServeHTTP(recorder, req)

		tc.check(t, recorder, menus)
	}
}

func TestFetchMenuWithDishesByCity(t *testing.T) {
	limit := int32(10)
	offset := int32(0)
	var menus []*domain.MenuWithDishes

	for i := 0; i < int(limit); i++ {
		menus = append(menus, randomMenuWithDishes(t))
	}

	offered := menus[0].OfferedAt.Format("2006-01-02")
	cityCode := menus[0].CityCode

	type req struct {
		CityCode int32
		Limit    sql.NullInt32
		Offset   sql.NullInt32
		Offered  string
	}

	testCases := []struct {
		name      string
		req       req
		buildStub func(uc *mocks.MockMenuWithDishesUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes)
	}{
		{
			name: "OK",
			req: req{
				CityCode: cityCode,
				Limit:    sql.NullInt32{Int32: limit, Valid: true},
				Offset:   sql.NullInt32{Int32: offset, Valid: true},
				Offered:  offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				parsedOffered, err := util.ParseDate(offered)

				require.NoError(t, err)

				uc.EXPECT().FetchByCity(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(parsedOffered), gomock.Eq(cityCode)).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 200, recorder.Code)
				requireBodyMatchMenuWithDishesList(t, recorder.Body, menus)
			},
		},
		{
			name: "Bad Request - Invalid CityCode",
			req: req{
				CityCode: -1,
				Limit:    sql.NullInt32{Int32: limit, Valid: true},
				Offset:   sql.NullInt32{Int32: offset, Valid: true},
				Offered:  offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().FetchByCity(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid Limit",
			req: req{
				CityCode: cityCode,
				Limit:    sql.NullInt32{Int32: -1, Valid: true},
				Offset:   sql.NullInt32{Int32: offset, Valid: true},
				Offered:  offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().FetchByCity(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid Offset",
			req: req{
				CityCode: cityCode,
				Limit:    sql.NullInt32{Int32: limit, Valid: true},
				Offset:   sql.NullInt32{Int32: -1, Valid: true},
				Offered:  offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().FetchByCity(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid Offered",
			req: req{
				CityCode: cityCode,
				Limit:    sql.NullInt32{Int32: limit, Valid: true},
				Offset:   sql.NullInt32{Int32: offset, Valid: true},
				Offered:  "invalid-offered",
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().FetchByCity(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "If Offered is not specified, return 400 error",
			req: req{
				CityCode: cityCode,
				Limit:    sql.NullInt32{Int32: limit, Valid: true},
				Offset:   sql.NullInt32{Int32: offset, Valid: true},
				Offered:  "",
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				uc.EXPECT().FetchByCity(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Any(), gomock.Eq(cityCode)).Times(0).Return(menus, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 400, recorder.Code)
			},
		},
		{
			name: "If Limit is not specified, it will be set to domain.DEFAULT_LIMIT",
			req: req{
				CityCode: cityCode,
				Limit:    sql.NullInt32{Int32: 0, Valid: false},
				Offset:   sql.NullInt32{Int32: offset, Valid: true},
				Offered:  offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				parsedOffered, err := util.ParseDate(offered)

				require.NoError(t, err)
				uc.EXPECT().FetchByCity(gomock.Any(), gomock.Eq(domain.DEFAULT_LIMIT), gomock.Eq(offset), gomock.Eq(parsedOffered), gomock.Eq(cityCode)).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 200, recorder.Code)
				requireBodyMatchMenuWithDishesList(t, recorder.Body, menus)
			},
		},
		{
			name: "If Offset is not specified, it will be set to domain.DEFAULT_OFFSET",
			req: req{
				CityCode: cityCode,
				Limit:    sql.NullInt32{Int32: limit, Valid: true},
				Offset:   sql.NullInt32{Int32: 0, Valid: false},
				Offered:  offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				parsedOffered, err := util.ParseDate(offered)

				require.NoError(t, err)
				uc.EXPECT().FetchByCity(gomock.Any(), gomock.Eq(limit), gomock.Eq(domain.DEFAULT_OFFSET), gomock.Eq(parsedOffered), gomock.Eq(cityCode)).Times(1).Return(menus, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 200, recorder.Code)
				requireBodyMatchMenuWithDishesList(t, recorder.Body, menus)
			},
		},
		{
			name: "Internal Server Error",
			req: req{
				CityCode: cityCode,
				Limit:    sql.NullInt32{Int32: limit, Valid: true},
				Offset:   sql.NullInt32{Int32: offset, Valid: true},
				Offered:  offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				parsedOffered, err := util.ParseDate(offered)

				require.NoError(t, err)
				uc.EXPECT().FetchByCity(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(parsedOffered), gomock.Eq(cityCode)).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 500, recorder.Code)
			},
		},
		{
			name: "Empty Result",
			req: req{
				CityCode: cityCode,
				Limit:    sql.NullInt32{Int32: limit, Valid: true},
				Offset:   sql.NullInt32{Int32: offset, Valid: true},
				Offered:  offered,
			},
			buildStub: func(uc *mocks.MockMenuWithDishesUsecase) {
				parsedOffered, err := util.ParseDate(offered)

				require.NoError(t, err)
				uc.EXPECT().FetchByCity(gomock.Any(), gomock.Eq(limit), gomock.Eq(offset), gomock.Eq(parsedOffered), gomock.Eq(cityCode)).Times(1).Return([]*domain.MenuWithDishes{}, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menus []*domain.MenuWithDishes) {
				require.Equal(t, 200, recorder.Code)
				data, err := io.ReadAll(recorder.Body)

				require.NoError(t, err)

				var res fetchMenuWithDishesResponse
				err = json.Unmarshal(data, &res)

				require.NoError(t, err)

				var menuData []*domain.MenuWithDishes

				menuData = append(menuData, res.Menus...)

				require.Empty(t, menuData)
				require.Equal(t, res.Next, "")
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mocks.NewMockMenuWithDishesUsecase(ctrl)
		tc.buildStub(uc)

		q := make(url.Values)
		if tc.req.Limit.Valid {
			q.Set("limit", fmt.Sprintf("%d", tc.req.Limit.Int32))
		}

		if tc.req.Offset.Valid {
			q.Set("offset", fmt.Sprintf("%d", tc.req.Offset.Int32))
		}

		q.Set("offered", tc.req.Offered)

		url := fmt.Sprintf("/cities/%d/menus?%s", tc.req.CityCode, q.Encode())
		req, err := http.NewRequest(http.MethodGet, url, nil)

		require.NoError(t, err)

		recorder := httptest.NewRecorder()
		e := newSetUpTestServer()
		e.GET("/cities/:code/menus", NewMenuWithDishesController(uc).FetchByCity)
		e.ServeHTTP(recorder, req)

		tc.check(t, recorder, menus)
	}
}

func requireBodyMatchMenuWithDishes(t *testing.T, body *bytes.Buffer, menu *domain.MenuWithDishes) {

	data, err := io.ReadAll(body)

	require.NoError(t, err)

	var res domain.MenuWithDishes
	err = json.Unmarshal(data, &res)
	require.NoError(t, err)

	require.NotNil(t, res.ID)

	require.Equal(t, menu.OfferedAt, res.OfferedAt)
	require.Equal(t, menu.PhotoUrl, res.PhotoUrl)
	require.Equal(t, menu.ElementarySchoolCalories, res.ElementarySchoolCalories)
	require.Equal(t, menu.JuniorHighSchoolCalories, res.JuniorHighSchoolCalories)
	require.Equal(t, menu.CityCode, res.CityCode)

	var dishes []*domain.Dish

	dishes = append(dishes, menu.Dishes...)

	require.Len(t, dishes, len(dishes))
}

func requireBodyMatchMenuWithDishesList(t *testing.T, body *bytes.Buffer, menus []*domain.MenuWithDishes) {
	data, err := io.ReadAll(body)

	require.NoError(t, err)

	var res fetchMenuWithDishesResponse
	err = json.Unmarshal(data, &res)

	require.NoError(t, err)

	var menuData []*domain.MenuWithDishes

	menuData = append(menuData, res.Menus...)

	require.Equal(t, menuData, res.Menus)
	require.Len(t, res.Menus, len(menus))
	require.Equal(t, res.Next, menus[len(menus)-1].OfferedAt.Format("2006-01-02"))
}

func randomMenuWithDishes(t *testing.T) *domain.MenuWithDishes {
	var dishes []*domain.Dish

	for i := 0; i < 10; i++ {
		dishes = append(dishes, randomDish(t))
	}
	menu, err := domain.ReNewMenuWithDishes(
		util.RandomUlid(),
		util.RandomDate(),
		util.RandomNullURL(),
		util.RandomInt32(),
		util.RandomInt32(),
		util.RandomCityCode(),
		dishes,
	)

	require.NoError(t, err)

	return menu
}
