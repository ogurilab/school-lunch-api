package db

import (
	"context"
	"testing"

	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreateDishesAllergens(t *testing.T) {
	city := createRandomCity(t)
	menu := createRandomMenu(t, city.CityCode)
	dish := createRandomDish(t, menu.ID)
	name := util.RandomString(50)
	category := util.RandomInt32()
	allergen := createRandomAllergen(t, name, category)

	createRandomDishesAllergens(t, dish.ID, allergen.ID, category)
}

func createRandomDishesAllergens(t *testing.T, dishID string, allergenID int32, category int32) error {
	ctx := context.Background()

	arg := CreateDishesAllergensParams{
		DishID:     dishID,
		AllergenID: allergenID,
		Category:   category,
	}

	err := testQuery.CreateDishesAllergens(ctx, arg)

	require.NoError(t, err)

	result, err := testQuery.getDishesAllergens(dishID, allergenID, category)

	require.NoError(t, err)

	require.NotEmpty(t, result)

	require.Equal(t, arg.DishID, result.DishID)
	require.Equal(t, arg.AllergenID, result.AllergenID)

	return nil
}
