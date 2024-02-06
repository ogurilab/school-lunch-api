package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/mocks"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestFetchAllergensByDishID(t *testing.T) {

	menu := randomMenu(t)
	allergens := randomAllergens(t, 10)

	testCases := []struct {
		name       string
		req        fetchAllergenByMenuIDRequest
		buildStubs func(u *mocks.MockAllergenUsecase)
		check      func(t *testing.T, recorder *httptest.ResponseRecorder, allergens []*domain.Allergen)
	}{
		{
			name: "OK",
			req: fetchAllergenByMenuIDRequest{
				MenuID: menu.ID,
			},
			buildStubs: func(au *mocks.MockAllergenUsecase) {
				au.EXPECT().FetchByDishID(gomock.Any(), gomock.Eq(menu.ID)).Times(1).Return(allergens, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, allergens []*domain.Allergen) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAllergens(t, recorder.Body, allergens)
			},
		},
		{
			name: "Bad Request",
			req: fetchAllergenByMenuIDRequest{
				MenuID: "invalid",
			},
			buildStubs: func(au *mocks.MockAllergenUsecase) {
				au.EXPECT().FetchByDishID(gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, allergens []*domain.Allergen) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			req: fetchAllergenByMenuIDRequest{
				MenuID: menu.ID,
			},
			buildStubs: func(au *mocks.MockAllergenUsecase) {
				au.EXPECT().FetchByDishID(gomock.Any(), gomock.Eq(menu.ID)).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, allergens []*domain.Allergen) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			au := mocks.NewMockAllergenUsecase(ctrl)
			tc.buildStubs(au)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/dishes/%s/allergens", tc.req.MenuID)

			req, err := http.NewRequest("GET", url, nil)

			require.NoError(t, err)

			e := newSetUpTestServer()
			e.GET("/dishes/:id/allergens", NewAllergenController(au).FetchByDishID)

			e.ServeHTTP(recorder, req)

			tc.check(t, recorder, allergens)
		})
	}
}

func TestFetchAllergensByMenuID(t *testing.T) {
	dish := randomDish(t)
	allergens := randomAllergens(t, 10)

	testCases := []struct {
		name       string
		req        fetchAllergenByDishIDRequest
		buildStubs func(u *mocks.MockAllergenUsecase)
		check      func(t *testing.T, recorder *httptest.ResponseRecorder, allergens []*domain.Allergen)
	}{
		{
			name: "OK",
			req: fetchAllergenByDishIDRequest{
				DishID: dish.ID,
			},
			buildStubs: func(au *mocks.MockAllergenUsecase) {
				au.EXPECT().FetchByMenuID(gomock.Any(), gomock.Eq(dish.ID)).Times(1).Return(allergens, nil)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, allergens []*domain.Allergen) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAllergens(t, recorder.Body, allergens)
			},
		},
		{
			name: "Bad Request",
			req: fetchAllergenByDishIDRequest{
				DishID: "invalid",
			},
			buildStubs: func(au *mocks.MockAllergenUsecase) {
				au.EXPECT().FetchByMenuID(gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, allergens []*domain.Allergen) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			req: fetchAllergenByDishIDRequest{
				DishID: dish.ID,
			},
			buildStubs: func(au *mocks.MockAllergenUsecase) {
				au.EXPECT().FetchByMenuID(gomock.Any(), gomock.Eq(dish.ID)).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, recorder *httptest.ResponseRecorder, allergens []*domain.Allergen) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			au := mocks.NewMockAllergenUsecase(ctrl)
			tc.buildStubs(au)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/menus/%s/allergens", tc.req.DishID)

			req, err := http.NewRequest("GET", url, nil)

			require.NoError(t, err)

			e := newSetUpTestServer()
			e.GET("/menus/:id/allergens", NewAllergenController(au).FetchByMenuID)

			e.ServeHTTP(recorder, req)

			tc.check(t, recorder, allergens)
		})
	}
}

func randomAllergen(t *testing.T) *domain.Allergen {
	return &domain.Allergen{
		ID:   util.RandomInt32(),
		Name: util.RandomString(50),
	}
}
func randomAllergens(t *testing.T, length int) []*domain.Allergen {

	allergens := make([]*domain.Allergen, 0, length)

	for i := 0; i < length; i++ {
		allergens = append(allergens, randomAllergen(t))
	}

	return allergens
}

func requireBodyMatchAllergens(t *testing.T, body *bytes.Buffer, allergens []*domain.Allergen) {
	data, err := io.ReadAll(body)

	require.NoError(t, err)

	var allergensData []*domain.Allergen

	err = json.Unmarshal(data, &allergensData)

	require.NoError(t, err)

	require.Equal(t, len(allergens), len(allergensData))

	require.ElementsMatch(t, allergens, allergensData)
}
