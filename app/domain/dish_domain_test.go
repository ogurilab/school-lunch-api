package domain

import (
	"testing"

	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestNewDish(t *testing.T) {

	testCases := []struct {
		name     string
		menuID   string
		dishName string
		check    func(*Dish, error)
	}{
		{
			name:     "OK",
			menuID:   "dish1",
			dishName: "dish1",
			check: func(dish *Dish, err error) {
				require.NoError(t, err)
				require.NotNil(t, dish)
				require.Equal(t, "dish1", dish.Name)
				require.NotEmpty(t, dish.ID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dish, err := NewDish(tc.dishName)
			tc.check(dish, err)
		})
	}

}

func TestReNewDish(t *testing.T) {
	id := util.NewUlid()
	testCases := []struct {
		name     string
		id       string
		menuID   string
		dishName string
		check    func(*Dish, error)
	}{
		{
			name:     "OK",
			id:       id,
			menuID:   "dish1",
			dishName: "dish1",
			check: func(dish *Dish, err error) {
				require.NoError(t, err)
				require.NotNil(t, dish)
				require.Equal(t, "dish1", dish.Name)
				require.Equal(t, id, dish.ID)
			},
		},

		{
			name:     "Invalid Ulid",
			id:       "invalid",
			menuID:   "dish1",
			dishName: "dish1",
			check: func(dish *Dish, err error) {
				require.Error(t, err)
				require.Nil(t, dish)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dish, err := ReNewDish(tc.id, tc.dishName)
			tc.check(dish, err)
		})
	}
}
