package domain

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestReNewMenuWithDishes(t *testing.T) {
	now := time.Now()
	var dishes []*Dish

	for i := 0; i < 3; i++ {
		dishes = append(dishes, &Dish{
			ID:   util.NewUlid(),
			Name: "dish",
		})
	}

	type input struct {
		id                       string
		OfferedAt                time.Time
		PhotoUrl                 sql.NullString
		ElementarySchoolCalories int32
		JuniorHighSchoolCalories int32
		CityCode                 int32
		Dishes                   []*Dish
	}

	testCases := []struct {
		name       string
		createStub func() input
		check      func(*MenuWithDishes, error)
	}{
		{
			name: "OK",
			createStub: func() input {
				return input{
					id:                       util.NewUlid(),
					OfferedAt:                now,
					PhotoUrl:                 sql.NullString{String: "http://example.com", Valid: true},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					CityCode:                 1,
					Dishes:                   dishes,
				}
			},
			check: func(m *MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, m)
				require.NotEmpty(t, m.ID)
				require.Equal(t, now, m.OfferedAt)
				require.Equal(t, "http://example.com", m.PhotoUrl.String)
				require.Equal(t, true, m.PhotoUrl.Valid)
				require.Equal(t, int32(100), m.ElementarySchoolCalories)
				require.Equal(t, int32(200), m.JuniorHighSchoolCalories)
				require.Equal(t, int32(1), m.CityCode)
				require.NotNil(t, m.Dishes)
				require.Len(t, m.Dishes, 3)
			},
		},
		{
			name: "Empty Dishes",
			createStub: func() input {
				return input{
					id:                       util.NewUlid(),
					OfferedAt:                now,
					PhotoUrl:                 sql.NullString{String: "http://example.com", Valid: true},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					CityCode:                 1,
					Dishes:                   nil,
				}
			},
			check: func(m *MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, m)
				require.NotEmpty(t, m.ID)
				require.Equal(t, now, m.OfferedAt)
				require.Equal(t, "http://example.com", m.PhotoUrl.String)
				require.Equal(t, true, m.PhotoUrl.Valid)
				require.Equal(t, int32(100), m.ElementarySchoolCalories)
				require.Equal(t, int32(200), m.JuniorHighSchoolCalories)
				require.Equal(t, int32(1), m.CityCode)
				require.Empty(t, m.Dishes)
				require.Len(t, m.Dishes, 0)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			stub := tc.createStub()

			m, err := ReNewMenuWithDishes(
				stub.id,
				stub.OfferedAt,
				stub.PhotoUrl,
				stub.ElementarySchoolCalories,
				stub.JuniorHighSchoolCalories,
				stub.CityCode,
				stub.Dishes,
			)

			tc.check(m, err)
		})
	}
}

func TestMenuWithDishesMarshalJSON(t *testing.T) {

	validPhotoUrlMenu := randomMenuWithDishes(t, true, false)

	invalidPhotoUrlMenu := randomMenuWithDishes(t, false, false)

	emptyDishesMenu := randomMenuWithDishes(t, true, true)

	testCases := []struct {
		name       string
		createStub func() MenuWithDishes
		check      func([]byte, error)
	}{
		{
			name: "OK Valid PhotoUrl",
			createStub: func() MenuWithDishes {
				return validPhotoUrlMenu
			},
			check: func(b []byte, err error) {
				require.NoError(t, err)
				require.NotNil(t, b)
				require.NotEmpty(t, b)
				requireEqualMenuWithDishesJSON(t, validPhotoUrlMenu, b)
			},
		},
		{
			name: "OK Invalid PhotoUrl",
			createStub: func() MenuWithDishes {
				return invalidPhotoUrlMenu
			},
			check: func(b []byte, err error) {
				require.NoError(t, err)
				require.NotNil(t, b)
				require.NotEmpty(t, b)
				requireEqualMenuWithDishesJSON(t, invalidPhotoUrlMenu, b)
			},
		},
		{
			name: "OK Empty Dishes",
			createStub: func() MenuWithDishes {
				return emptyDishesMenu
			},
			check: func(b []byte, err error) {
				require.NoError(t, err)
				require.NotNil(t, b)
				requireEqualMenuWithEmptyDishesJSON(t, emptyDishesMenu, b)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			stub := tc.createStub()

			b, err := stub.MarshalJSON()

			tc.check(b, err)
		})
	}
}

func TestMenuWithDishesUnmarshalJSON(t *testing.T) {

	validPhotoUrlMenu := randomMenuWithDishes(t, true, false)

	invalidPhotoUrlMenu := randomMenuWithDishes(t, false, false)

	emptyDishesMenu := randomMenuWithDishes(t, true, true)

	testCases := []struct {
		name       string
		createStub func() MenuWithDishes
		check      func(*MenuWithDishes, error)
	}{
		{
			name: "OK Valid PhotoUrl",
			createStub: func() MenuWithDishes {
				return validPhotoUrlMenu
			},
			check: func(m *MenuWithDishes, err error) {

				require.NoError(t, err)
				require.NotNil(t, m)
				requireEqualMenuWithDishes(t, validPhotoUrlMenu, *m)
			},
		},
		{
			name: "OK Invalid PhotoUrl",
			createStub: func() MenuWithDishes {
				return invalidPhotoUrlMenu
			},
			check: func(m *MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, m)

				requireEqualMenuWithDishes(t, invalidPhotoUrlMenu, *m)
			},
		},
		{
			name: "OK Empty Dishes",
			createStub: func() MenuWithDishes {
				return emptyDishesMenu
			},
			check: func(m *MenuWithDishes, err error) {
				require.NoError(t, err)
				require.NotNil(t, m)
				requireEqualMenuWithDishes(t, emptyDishesMenu, *m)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			stub := tc.createStub()

			b, err := stub.MarshalJSON()
			require.NoError(t, err)

			m := &MenuWithDishes{}
			err = m.UnmarshalJSON(b)

			tc.check(m, err)
		})
	}
}

func requireEqualMenuWithDishesJSON(t *testing.T, m MenuWithDishes, actual []byte) {

	var photoUrlStr string
	if m.PhotoUrl.Valid {
		photoUrlStr = fmt.Sprintf(`"%s"`, m.PhotoUrl.String)
	} else {
		photoUrlStr = "null"
	}

	expect := fmt.Sprintf(`{"id":"%s","offered_at":"%s","photo_url":%s,"elementary_school_calories":%d,"junior_high_school_calories":%d,"city_code":%d,"dishes":[{"id":"%s","name":"%s"}]}`,
		m.ID,
		m.OfferedAt.Format("2006-01-02"),
		photoUrlStr,
		m.ElementarySchoolCalories,
		m.JuniorHighSchoolCalories,
		m.CityCode,
		m.Dishes[0].ID,
		m.Dishes[0].Name,
	)

	require.Equal(t, expect, string(actual))
}

func requireEqualMenuWithEmptyDishesJSON(t *testing.T, m MenuWithDishes, actual []byte) {
	var photoUrlStr string
	if m.PhotoUrl.Valid {
		photoUrlStr = fmt.Sprintf(`"%s"`, m.PhotoUrl.String)
	} else {
		photoUrlStr = "null"
	}

	expect := fmt.Sprintf(`{"id":"%s","offered_at":"%s","photo_url":%s,"elementary_school_calories":%d,"junior_high_school_calories":%d,"city_code":%d,"dishes":[]}`,
		m.ID,
		m.OfferedAt.Format("2006-01-02"),
		photoUrlStr,
		m.ElementarySchoolCalories,
		m.JuniorHighSchoolCalories,
		m.CityCode,
	)

	require.Equal(t, expect, string(actual))
}

func requireEqualMenuWithDishes(t *testing.T, expect MenuWithDishes, actual MenuWithDishes) {
	requireEqualMenu(t, expect.Menu, actual.Menu)
	require.Len(t, actual.Dishes, len(expect.Dishes))
}

func randomMenuWithDishes(t *testing.T, valid bool, empty bool) MenuWithDishes {
	menu := randomMenu(t, valid)

	if empty {
		return MenuWithDishes{
			Menu:   menu,
			Dishes: nil,
		}
	}

	var dishes []*Dish

	for i := 0; i < 1; i++ {
		dishes = append(dishes, &Dish{
			ID:   "dish",
			Name: "dish",
		})
	}

	return MenuWithDishes{
		Menu:   menu,
		Dishes: dishes,
	}
}
