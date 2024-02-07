package usecase

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/domain/mocks"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestFetchAllergensByDishID(t *testing.T) {

	dish := randomDish(t)
	timeout := time.Second * 10
	ctx := context.Background()
	results := randomAllergens(t, 10)

	testCases := []struct {
		name       string
		dishID     string
		buildStubs func(r *mocks.MockAllergenRepository)
		check      func(t *testing.T, allergens []*domain.Allergen, err error)
	}{
		{
			name:   "OK",
			dishID: dish.ID,
			buildStubs: func(r *mocks.MockAllergenRepository) {
				r.EXPECT().FetchByDishID(gomock.Any(), gomock.Eq(dish.ID)).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {

				requireAllergensResult(t, allergens, err)
			},
		},
		{
			name:   "NG",
			dishID: dish.ID,
			buildStubs: func(r *mocks.MockAllergenRepository) {
				r.EXPECT().FetchByDishID(gomock.Any(), gomock.Any()).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				require.Error(t, err)
				require.Nil(t, allergens)
			},
		},
		{
			name:   "Empty Result",
			dishID: dish.ID,
			buildStubs: func(r *mocks.MockAllergenRepository) {
				r.EXPECT().FetchByDishID(gomock.Any(), gomock.Eq(dish.ID)).Times(1).Return([]*domain.Allergen{}, nil)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				require.NoError(t, err)
				require.Len(t, allergens, 0)
				require.Empty(t, allergens)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockAllergenRepository(ctrl)
			tc.buildStubs(repo)

			au := NewAllergenUsecase(repo, nil, timeout)

			allergens, err := au.FetchByDishID(ctx, tc.dishID)

			tc.check(t, allergens, err)
		})
	}
}

func TestFetchAllergensByMenuID(t *testing.T) {
	menu := randomMenu(t)

	timeout := time.Second * 10
	ctx := context.Background()

	dishes := make([]*domain.Dish, 0, 10)
	for i := 0; i < 10; i++ {
		dish := randomDish(t)
		dishes = append(dishes, dish)
	}
	results := randomAllergens(t, 10)

	testCases := []struct {
		name       string
		menuID     string
		buildStubs func(dr *mocks.MockDishRepository, ar *mocks.MockAllergenRepository)
		check      func(t *testing.T, allergens []*domain.Allergen, err error)
	}{
		{
			name:   "OK",
			menuID: menu.ID,
			buildStubs: func(dr *mocks.MockDishRepository, ar *mocks.MockAllergenRepository) {
				dr.EXPECT().FetchByMenuID(gomock.Any(), gomock.Eq(menu.ID)).Times(1).Return(dishes, nil)

				dishIDs := createDishIds(dishes)

				ar.EXPECT().FetchInDish(gomock.Any(), gomock.Eq(dishIDs)).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				requireAllergensResult(t, allergens, err)
			},
		},
		{
			name:   "NG - FetchByMenuID",
			menuID: menu.ID,
			buildStubs: func(dr *mocks.MockDishRepository, ar *mocks.MockAllergenRepository) {
				dr.EXPECT().FetchByMenuID(gomock.Any(), gomock.Any()).Times(1).Return(nil, sql.ErrConnDone)

				ar.EXPECT().FetchInDish(gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				require.Error(t, err)
				require.Nil(t, allergens)
			},
		},
		{
			name:   "NG - FetchInDish",
			menuID: menu.ID,
			buildStubs: func(dr *mocks.MockDishRepository, ar *mocks.MockAllergenRepository) {
				dr.EXPECT().FetchByMenuID(gomock.Any(), gomock.Eq(menu.ID)).Times(1).Return(dishes, nil)

				dishIDs := createDishIds(dishes)

				ar.EXPECT().FetchInDish(gomock.Any(), gomock.Eq(dishIDs)).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				require.Error(t, err)
				require.Nil(t, allergens)
			},
		},
		{
			name:   "Empty Result - FetchByMenuID",
			menuID: menu.ID,
			buildStubs: func(dr *mocks.MockDishRepository, ar *mocks.MockAllergenRepository) {
				dr.EXPECT().FetchByMenuID(gomock.Any(), gomock.Eq(menu.ID)).Times(1).Return([]*domain.Dish{}, nil)

				ar.EXPECT().FetchInDish(gomock.Any(), gomock.Any()).Times(0)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				require.NoError(t, err)
				require.Len(t, allergens, 0)
				require.Empty(t, allergens)
			},
		},
		{
			name:   "Empty Result - FetchInDish",
			menuID: menu.ID,
			buildStubs: func(dr *mocks.MockDishRepository, ar *mocks.MockAllergenRepository) {
				dr.EXPECT().FetchByMenuID(gomock.Any(), gomock.Eq(menu.ID)).Times(1).Return(dishes, nil)

				dishIDs := createDishIds(dishes)

				ar.EXPECT().FetchInDish(gomock.Any(), gomock.Eq(dishIDs)).Times(1).Return([]*domain.Allergen{}, nil)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				require.NoError(t, err)
				require.Len(t, allergens, 0)
				require.Empty(t, allergens)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			dr := mocks.NewMockDishRepository(ctrl)
			ar := mocks.NewMockAllergenRepository(ctrl)
			tc.buildStubs(dr, ar)

			au := NewAllergenUsecase(ar, dr, timeout)

			allergens, err := au.FetchByMenuID(ctx, tc.menuID)

			tc.check(t, allergens, err)
		})
	}
}

func randomAllergens(t *testing.T, length int) []*domain.Allergen {

	allergens := make([]*domain.Allergen, 0, length)

	for i := 0; i < length; i++ {
		allergen := domain.ReNewAllergen(
			util.RandomInt32(),
			util.RandomString(10),
			util.RandomInt32(),
		)

		allergens = append(allergens, allergen)
	}

	return allergens
}

func requireAllergensResult(t *testing.T, allergens []*domain.Allergen, err error) {
	require.NoError(t, err)
	require.Len(t, allergens, len(allergens))

	require.ElementsMatch(t, allergens, allergens)
}

func createDishIds(dishes []*domain.Dish) []string {
	dishIds := make([]string, 0, len(dishes))

	for _, dish := range dishes {
		dishIds = append(dishIds, dish.ID)
	}

	return dishIds
}
