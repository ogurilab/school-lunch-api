package domain

import (
	"database/sql"
	"testing"
	"time"

	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestNewMenu(t *testing.T) {

	now := time.Now()

	type input struct {
		OfferedAt                time.Time
		PhotoUrl                 sql.NullString
		ElementarySchoolCalories int32
		JuniorHighSchoolCalories int32
		CityCode                 int32
	}

	testCases := []struct {
		name       string
		createStub func() input
		check      func(*Menu, error)
	}{
		{
			name: "OK Valid PhotoUrl",
			createStub: func() input {
				return input{
					OfferedAt:                now,
					PhotoUrl:                 sql.NullString{String: "http://example.com", Valid: true},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					CityCode:                 1,
				}
			},
			check: func(m *Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, m)
				require.NotEmpty(t, m.ID)
				require.Equal(t, now, m.OfferedAt)
				require.Equal(t, "http://example.com", m.PhotoUrl.String)
				require.Equal(t, true, m.PhotoUrl.Valid)
				require.Equal(t, int32(100), m.ElementarySchoolCalories)
				require.Equal(t, int32(200), m.JuniorHighSchoolCalories)
				require.Equal(t, int32(1), m.CityCode)
			},
		},
		{
			name: "OK Invalid PhotoUrl",
			createStub: func() input {
				return input{
					OfferedAt:                now,
					PhotoUrl:                 sql.NullString{String: "", Valid: false},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					CityCode:                 1,
				}
			},
			check: func(m *Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, m)
				require.NotEmpty(t, m.ID)
				require.Equal(t, now, m.OfferedAt)
				require.Equal(t, "", m.PhotoUrl.String)
				require.Equal(t, false, m.PhotoUrl.Valid)
				require.Equal(t, int32(100), m.ElementarySchoolCalories)
				require.Equal(t, int32(200), m.JuniorHighSchoolCalories)
				require.Equal(t, int32(1), m.CityCode)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			stub := tc.createStub()

			m, err := NewMenu(
				stub.OfferedAt,
				stub.PhotoUrl,
				stub.ElementarySchoolCalories,
				stub.JuniorHighSchoolCalories,
				stub.CityCode,
			)

			tc.check(m, err)
		})
	}
}

func TestReNewMenu(t *testing.T) {
	id := util.NewUlid()
	now := time.Now()

	testCases := []struct {
		name       string
		createStub func() Menu
		check      func(*Menu, error)
	}{
		{
			name: "OK Valid PhotoUrl",
			createStub: func() Menu {
				return Menu{
					ID:                       id,
					OfferedAt:                now,
					PhotoUrl:                 sql.NullString{String: "http://example.com", Valid: true},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					CityCode:                 1,
				}
			},
			check: func(m *Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, m)
				require.Equal(t, id, m.ID)
				require.Equal(t, now, m.OfferedAt)
				require.Equal(t, "http://example.com", m.PhotoUrl.String)
				require.Equal(t, true, m.PhotoUrl.Valid)
				require.Equal(t, int32(100), m.ElementarySchoolCalories)
				require.Equal(t, int32(200), m.JuniorHighSchoolCalories)
				require.Equal(t, int32(1), m.CityCode)
			},
		},
		{
			name: "OK Invalid PhotoUrl",
			createStub: func() Menu {
				return Menu{
					ID:                       id,
					OfferedAt:                now,
					PhotoUrl:                 sql.NullString{String: "", Valid: false},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					CityCode:                 1,
				}
			},
			check: func(m *Menu, err error) {
				require.NoError(t, err)
				require.NotNil(t, m)
				require.Equal(t, id, m.ID)
				require.Equal(t, now, m.OfferedAt)
				require.Equal(t, "", m.PhotoUrl.String)
				require.Equal(t, false, m.PhotoUrl.Valid)
				require.Equal(t, int32(100), m.ElementarySchoolCalories)
				require.Equal(t, int32(200), m.JuniorHighSchoolCalories)
				require.Equal(t, int32(1), m.CityCode)
			},
		},
		{
			name: "NG Invalid ID",
			createStub: func() Menu {
				return Menu{
					ID:                       "invalid",
					OfferedAt:                now,
					PhotoUrl:                 sql.NullString{String: "", Valid: false},
					ElementarySchoolCalories: 100,
					JuniorHighSchoolCalories: 200,
					CityCode:                 1,
				}
			},
			check: func(m *Menu, err error) {
				require.Error(t, err)
				require.Nil(t, m)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			stub := tc.createStub()

			m, err := ReNewMenu(
				stub.ID,
				stub.OfferedAt,
				stub.PhotoUrl,
				stub.ElementarySchoolCalories,
				stub.JuniorHighSchoolCalories,
				stub.CityCode,
			)

			tc.check(m, err)
		})
	}

}

func TestReNewMenuWithDishes(t *testing.T) {
	now := time.Now()
	var dishes []*Dish

	for i := 0; i < 3; i++ {
		dishes = append(dishes, &Dish{
			ID:     util.NewUlid(),
			Name:   "dish",
			MenuID: util.NewUlid(),
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
