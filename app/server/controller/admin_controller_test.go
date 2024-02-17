package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/bootstrap"
	"github.com/ogurilab/school-lunch-api/domain/mocks"
	"github.com/ogurilab/school-lunch-api/server/validator"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func createValidAdminKey(t *testing.T, env bootstrap.Env, req *http.Request) {
	req.Header.Set("X-Admin-Key", env.ADMIN_KEY)
}

func TestCreateMenu(t *testing.T) {
	menu := randomMenu(t)
	offered := menu.OfferedAt.Format("2006-01-02")
	type body createMenuRequest

	testCases := []struct {
		name      string
		body      body
		setUpKey  func(t *testing.T, env bootstrap.Env, req *http.Request)
		buildStub func(uc *mocks.MockMenuUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: body{
				OfferedAt:                offered,
				PhotoUrl:                 menu.PhotoUrl.String,
				ElementarySchoolCalories: menu.ElementarySchoolCalories,
				JuniorHighSchoolCalories: menu.JuniorHighSchoolCalories,
				CityCode:                 menu.CityCode,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockMenuUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid OfferedAt",
			body: body{
				OfferedAt:                "1000-01-01:00:00:00",
				PhotoUrl:                 menu.PhotoUrl.String,
				ElementarySchoolCalories: menu.ElementarySchoolCalories,
				JuniorHighSchoolCalories: menu.JuniorHighSchoolCalories,
				CityCode:                 menu.CityCode,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockMenuUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)

			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid PhotoUrl",
			body: body{
				OfferedAt:                offered,
				PhotoUrl:                 "invalid-url",
				ElementarySchoolCalories: menu.ElementarySchoolCalories,
				JuniorHighSchoolCalories: menu.JuniorHighSchoolCalories,
				CityCode:                 menu.CityCode,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockMenuUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid ElementarySchoolCalories",
			body: body{
				OfferedAt:                offered,
				PhotoUrl:                 menu.PhotoUrl.String,
				ElementarySchoolCalories: -1,
				JuniorHighSchoolCalories: menu.JuniorHighSchoolCalories,
				CityCode:                 menu.CityCode,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockMenuUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid JuniorHighSchoolCalories",
			body: body{
				OfferedAt:                offered,
				PhotoUrl:                 menu.PhotoUrl.String,
				ElementarySchoolCalories: menu.ElementarySchoolCalories,
				JuniorHighSchoolCalories: -1,
				CityCode:                 menu.CityCode,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockMenuUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Bad Request - Invalid CityCode",
			body: body{
				OfferedAt:                offered,
				PhotoUrl:                 menu.PhotoUrl.String,
				ElementarySchoolCalories: menu.ElementarySchoolCalories,
				JuniorHighSchoolCalories: menu.JuniorHighSchoolCalories,
				CityCode:                 -1,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockMenuUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			body: body{
				OfferedAt:                offered,
				PhotoUrl:                 menu.PhotoUrl.String,
				ElementarySchoolCalories: menu.ElementarySchoolCalories,
				JuniorHighSchoolCalories: menu.JuniorHighSchoolCalories,
				CityCode:                 menu.CityCode,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockMenuUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any()).Times(1).Return(sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, 500, recorder.Code)
			},
		},
		{
			name: "Bad Admin Key",
			body: body{
				OfferedAt:                offered,
				PhotoUrl:                 menu.PhotoUrl.String,
				ElementarySchoolCalories: menu.ElementarySchoolCalories,
				JuniorHighSchoolCalories: menu.JuniorHighSchoolCalories,
				CityCode:                 menu.CityCode,
			},
			setUpKey: func(t *testing.T, env bootstrap.Env, req *http.Request) {
				req.Header.Set("X-Admin-Key", "invalid")
			},
			buildStub: func(uc *mocks.MockMenuUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "No Admin Key",
			body: body{
				OfferedAt:                offered,
				PhotoUrl:                 menu.PhotoUrl.String,
				ElementarySchoolCalories: menu.ElementarySchoolCalories,
				JuniorHighSchoolCalories: menu.JuniorHighSchoolCalories,
				CityCode:                 menu.CityCode,
			},
			setUpKey: func(t *testing.T, env bootstrap.Env, req *http.Request) {
				// do nothing
			},
			buildStub: func(uc *mocks.MockMenuUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mocks.NewMockMenuUsecase(ctrl)
		tc.buildStub(uc)

		recorder := httptest.NewRecorder()

		jsonData, err := json.Marshal(tc.body)
		reqBody := bytes.NewBuffer(jsonData)

		require.NotNil(t, reqBody)
		require.NoError(t, err)

		url := "/admin/menus"
		req, err := http.NewRequest(http.MethodPost, url, reqBody)

		require.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		e, env := newSetupAdminTestServer(t)
		tc.setUpKey(t, env, req)

		e.POST(url, NewAdminController(uc, nil).CreateMenu)
		e.ServeHTTP(recorder, req)

		tc.check(t, recorder)
	}
}

func TestCreateDish(t *testing.T) {
	menu := randomMenu(t)
	dish := randomDish(t)

	type body struct {
		Name string `json:"name" validate:"required"`
	}

	testCases := []struct {
		name      string
		menuID    string
		body      body
		setUpKey  func(t *testing.T, env bootstrap.Env, req *http.Request)
		buildStub func(uc *mocks.MockDishUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			menuID: menu.ID,
			body: body{
				Name: dish.Name,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name:   "Bad Request - Invalid MenuID",
			menuID: "invalid",
			body: body{
				Name: dish.Name,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "Bad Request - Invalid Name",
			menuID: menu.ID,
			body: body{
				Name: "",
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "Internal Server Error",
			menuID: menu.ID,
			body: body{
				Name: dish.Name,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "Bad Admin Key",
			menuID: menu.ID,
			body: body{
				Name: dish.Name,
			},
			setUpKey: func(t *testing.T, env bootstrap.Env, req *http.Request) {
				req.Header.Set("X-Admin-Key", "invalid")
			},
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:   "No Admin Key",
			menuID: menu.ID,
			body: body{
				Name: dish.Name,
			},
			setUpKey: func(t *testing.T, env bootstrap.Env, req *http.Request) {
				// do nothing
			},
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		uc := mocks.NewMockDishUsecase(ctrl)
		tc.buildStub(uc)

		recorder := httptest.NewRecorder()

		jsonData, err := json.Marshal(tc.body)
		reqBody := bytes.NewBuffer(jsonData)

		require.NotNil(t, reqBody)
		require.NoError(t, err)

		url := fmt.Sprintf("/admin/menus/%s/dishes", tc.menuID)
		req, err := http.NewRequest(http.MethodPost, url, reqBody)

		require.NoError(t, err)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		e, env := newSetupAdminTestServer(t)
		tc.setUpKey(t, env, req)

		e.POST("/admin/menus/:id/dishes", NewAdminController(nil, uc).CreateDish)
		e.ServeHTTP(recorder, req)

		tc.check(t, recorder)
	}
}

func TestCreateDishes(t *testing.T) {
	menu := randomMenu(t)

	dishes := make([]validator.Dish, 0, 3)
	badDishes := make([]validator.Dish, 0, 3)

	for i := 0; i < 3; i++ {
		d := randomDish(t)
		dishes = append(dishes, validator.Dish{
			Name: d.Name,
		})

		if i == 0 {
			d.Name = ""
		}

		badDishes = append(badDishes, validator.Dish{
			Name: d.Name,
		})
	}

	type body struct {
		Dishes []validator.Dish `json:"dishes" validate:"required,dishes"`
	}

	testCases := []struct {
		name      string
		menuID    string
		body      body
		setUpKey  func(t *testing.T, env bootstrap.Env, req *http.Request)
		buildStub func(uc *mocks.MockDishUsecase)
		check     func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			menuID: menu.ID,
			body: body{
				Dishes: dishes,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockDishUsecase) {

				uc.EXPECT().CreateMany(gomock.Any(), gomock.Any(), gomock.Eq(menu.ID)).Times(1).Return(nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name:   "Bad Request - Invalid MenuID",
			menuID: "invalid",
			body: body{
				Dishes: dishes,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().CreateMany(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "Bad Request - Invalid Dishes",
			menuID: menu.ID,
			body: body{
				Dishes: badDishes,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().CreateMany(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "Internal Server Error",
			menuID: menu.ID,
			body: body{
				Dishes: dishes,
			},
			setUpKey: createValidAdminKey,
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().CreateMany(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "Bad Admin Key",
			menuID: menu.ID,
			body: body{
				Dishes: dishes,
			},
			setUpKey: func(t *testing.T, env bootstrap.Env, req *http.Request) {
				req.Header.Set("X-Admin-Key", "invalid")
			},
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().CreateMany(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:   "No Admin Key",
			menuID: menu.ID,
			body: body{
				Dishes: dishes,
			},
			setUpKey: func(t *testing.T, env bootstrap.Env, req *http.Request) {
				// do nothing
			},
			buildStub: func(uc *mocks.MockDishUsecase) {
				uc.EXPECT().CreateMany(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mocks.NewMockDishUsecase(ctrl)
			tc.buildStub(uc)

			recorder := httptest.NewRecorder()

			jsonData, err := json.Marshal(tc.body)
			reqBody := bytes.NewBuffer(jsonData)

			require.NotNil(t, reqBody)
			require.NoError(t, err)

			url := fmt.Sprintf("/admin/menus/%s/dishes/bulk", tc.menuID)
			req, err := http.NewRequest(http.MethodPost, url, reqBody)

			require.NoError(t, err)

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e, env := newSetupAdminTestServer(t)
			tc.setUpKey(t, env, req)

			e.POST("/admin/menus/:id/dishes/bulk", NewAdminController(nil, uc).CreateDishes)
			e.ServeHTTP(recorder, req)

			tc.check(t, recorder)
		})
	}
}
