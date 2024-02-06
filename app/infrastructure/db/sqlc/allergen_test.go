package db

import (
	"context"
	"testing"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAllergen(t *testing.T) {
	name := util.RandomString(50)
	createRandomAllergen(t, name)
}

func TestListAllergenByDishID(t *testing.T) {
	cityCode := util.RandomCityCode()
	menu := createRandomMenu(t, cityCode)
	dish := createRandomDish(t, menu.ID)

	allergens := createRandomAllergens(t, 10)

	for _, allergen := range allergens {
		createRandomDishesAllergens(t, dish.ID, allergen.ID)
	}

	res, err := testQuery.ListAllergenByDishID(context.Background(), dish.ID)

	require.NoError(t, err)

	require.Len(t, res, 10)

	resNames := make([]string, 0, 10)
	allergensNames := make([]string, 0, 10)

	for _, allergen := range allergens {
		allergensNames = append(allergensNames, allergen.Name)
	}

	for _, allergen := range res {
		resNames = append(resNames, allergen.Name)
	}

	require.ElementsMatch(t, allergensNames, resNames)
}

func TestListAllergenInDish(t *testing.T) {
	cityCode := util.RandomCityCode()
	menu := createRandomMenu(t, cityCode)

	dishLength := 3
	allergensLength := 5
	dishes := make([]*domain.Dish, 0, dishLength)
	allAllergens := make([]*domain.Allergen, 0, allergensLength*dishLength)

	// dishの作成
	for i := 0; i < dishLength; i++ {
		dish := createRandomDish(t, menu.ID)
		dishes = append(dishes, dish)
	}

	dishIDs := make([]string, 0, dishLength)

	for _, dish := range dishes {
		dishIDs = append(dishIDs, dish.ID)
	}

	// dishに対してそれぞれ、allergenを紐付ける
	for _, dishID := range dishIDs {
		allergens := createRandomAllergens(t, allergensLength)

		allAllergens = append(allAllergens, allergens...)

		for _, allergen := range allergens {
			createRandomDishesAllergens(t, dishID, allergen.ID)
		}

	}

	res, err := testQuery.ListAllergenInDish(context.Background(), dishIDs)

	require.NoError(t, err)

	// 指定したdishのidに紐づいたallergenが全て取得できているか
	require.Len(t, res, allergensLength*dishLength)

	resNames := make([]string, 0, allergensLength*dishLength)
	allergensNames := make([]string, 0, allergensLength*dishLength)

	for _, allergen := range res {
		resNames = append(resNames, allergen.Name)
	}

	for _, allergen := range allAllergens {
		allergensNames = append(allergensNames, allergen.Name)
	}

	require.ElementsMatch(t, allergensNames, resNames)
}

func createRandomAllergen(t *testing.T, name string) *domain.Allergen {

	ctx := context.Background()

	err := testQuery.CreateAllergen(ctx, name)

	require.NoError(t, err)

	res, err := testQuery.GetAllergenByName(ctx, name)

	require.NoError(t, err)

	require.NotEmpty(t, res)

	require.Equal(t, name, res.Name)

	return domain.ReNewAllergen(res.ID, res.Name)
}

func createRandomAllergens(t *testing.T, length int) []*domain.Allergen {
	allergens := make([]*domain.Allergen, 0, length)

	for i := 0; i < length; i++ {
		name := util.RandomString(50)
		allergen := createRandomAllergen(t, name)
		allergens = append(allergens, allergen)
	}

	return allergens
}
