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

func TestFetchAllegerByDishID(t *testing.T) {
	dish := randomDish(t)
	results := randomDbAllergens(t, 10)

	testCases := []struct {
		name       string
		dishID     string
		buildStubs func(query *mocks.MockQuery)
		check      func(t *testing.T, allergens []*domain.Allergen, err error)
	}{
		{
			name:   "OK",
			dishID: dish.ID,
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().ListAllergenByDishID(gomock.Any(), gomock.Eq(dish.ID)).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				require.NoError(t, err)
				require.Len(t, allergens, len(results))
			},
		},
		{
			name: "NG",
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().ListAllergenByDishID(gomock.Any(), gomock.Any()).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				require.Error(t, err)
				require.Nil(t, allergens)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.buildStubs(query)

			repo := NewAllergenRepository(query)

			allergens, err := repo.FetchByDishID(context.Background(), tc.dishID)

			tc.check(t, allergens, err)
		})
	}
}

func TestFetchInDish(t *testing.T) {
	length := 10
	dishIds := make([]string, 0, length)

	for i := 0; i < length; i++ {
		dish := randomDish(t)
		dishIds = append(dishIds, dish.ID)
	}

	results := randomDbAllergens(t, 10)

	testCases := []struct {
		name       string
		dishIDs    []string
		buildStubs func(query *mocks.MockQuery)
		check      func(t *testing.T, allergens []*domain.Allergen, err error)
	}{
		{
			name:    "OK",
			dishIDs: dishIds,
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().ListAllergenInDish(gomock.Any(), gomock.Eq(dishIds)).Times(1).Return(results, nil)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				require.NoError(t, err)
				require.Len(t, allergens, len(results))
			},
		},
		{
			name:    "NG",
			dishIDs: dishIds,
			buildStubs: func(query *mocks.MockQuery) {
				query.EXPECT().ListAllergenInDish(gomock.Any(), gomock.Eq(dishIds)).Times(1).Return(nil, sql.ErrConnDone)
			},
			check: func(t *testing.T, allergens []*domain.Allergen, err error) {
				require.Error(t, err)
				require.Nil(t, allergens)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			query := mocks.NewMockQuery(ctrl)
			tc.buildStubs(query)

			repo := NewAllergenRepository(query)

			allergens, err := repo.FetchInDish(context.Background(), tc.dishIDs)

			tc.check(t, allergens, err)
		})
	}
}

func randomDbAllergens(t *testing.T, length int) []db.Allergen {

	allergens := make([]db.Allergen, 0, length)

	for i := 0; i < length; i++ {
		allergen := db.Allergen{
			ID:   util.RandomInt32(),
			Name: util.RandomString(50),
		}

		allergens = append(allergens, allergen)
	}

	return allergens

}
