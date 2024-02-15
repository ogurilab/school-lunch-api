package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/bootstrap"
	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/mocks"
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
		check     func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.Menu)
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
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.Menu) {
				require.Equal(t, 200, recorder.Code)
				requireBodyMatchMenu(t, recorder.Body, menu)
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
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.Menu) {
				require.Equal(t, 400, recorder.Code)
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
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.Menu) {
				require.Equal(t, 400, recorder.Code)
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
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.Menu) {
				require.Equal(t, 400, recorder.Code)
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
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.Menu) {
				require.Equal(t, 400, recorder.Code)
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
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.Menu) {
				require.Equal(t, 400, recorder.Code)
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
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.Menu) {
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
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.Menu) {
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
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, menu *domain.Menu) {
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

		tc.check(t, recorder, menu)
	}
}
