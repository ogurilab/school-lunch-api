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
	allergen := createRandomAllergen(t, name)

	createRandomDishesAllergens(t, dish.ID, allergen.ID)
}

func createRandomDishesAllergens(t *testing.T, dishID string, AllergenID int32) error {
	ctx := context.Background()

	arg := CreateDishesAllergensParams{
		DishID:     dishID,
		AllergenID: AllergenID,
	}

	err := testQuery.CreateDishesAllergens(ctx, arg)

	require.NoError(t, err)

	result, err := testQuery.getDishesAllergens(dishID, AllergenID)

	require.NoError(t, err)

	require.NotEmpty(t, result)

	require.Equal(t, arg.DishID, result.DishID)
	require.Equal(t, arg.AllergenID, result.AllergenID)

	return nil
}
