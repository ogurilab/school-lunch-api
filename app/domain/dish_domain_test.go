package domain

import (
	"encoding/json"
	"fmt"
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
				require.Equal(t, "dish1", dish.MenuID)
				require.Equal(t, "dish1", dish.Name)
				require.NotEmpty(t, dish.ID)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dish, err := NewDish(tc.menuID, tc.dishName)
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
				require.Equal(t, "dish1", dish.MenuID)
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
			dish, err := ReNewDish(tc.id, tc.menuID, tc.dishName)
			tc.check(dish, err)
		})
	}
}

func TestNewDishesFromJson(t *testing.T) {
	id := util.NewUlid()
	testCases := []struct {
		name        string
		createStubs func() json.RawMessage
		check       func([]*Dish, error)
	}{
		{
			name: "OK",
			createStubs: func() json.RawMessage {
				data := fmt.Sprintf(`[{"id": "%s", "menu_id": "menu1", "name": "dish1"}]`, id)

				return json.RawMessage(data)
			},
			check: func(dishes []*Dish, err error) {
				require.NoError(t, err)
				require.NotNil(t, dishes)
				require.Len(t, dishes, 1)
				require.Equal(t, "menu1", dishes[0].MenuID)
				require.Equal(t, "dish1", dishes[0].Name)
				require.Equal(t, id, dishes[0].ID)
			},
		},

		{
			name: "Multiple Dishes",
			createStubs: func() json.RawMessage {
				data := fmt.Sprintf(`[{"id": "%s", "menu_id": "menu1", "name": "dish1"},{"id": "%s", "menu_id": "menu1", "name": "dish2"}]`, id, id)

				return json.RawMessage(data)

			},
			check: func(dishes []*Dish, err error) {
				require.NoError(t, err)
				require.NotNil(t, dishes)
				require.Len(t, dishes, 2)
			},
		},
		{
			name: "Invalid JSON",
			createStubs: func() json.RawMessage {
				data := fmt.Sprintf(`[{"id": "%s", "menu_id": "menu1", "name": "dish1"}`, id)

				return json.RawMessage(data)
			},
			check: func(dishes []*Dish, err error) {
				require.Error(t, err)
				require.Nil(t, dishes)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stub := tc.createStubs()
			dishes, err := NewDishesFromJson(stub)
			tc.check(dishes, err)
		})

	}
}
