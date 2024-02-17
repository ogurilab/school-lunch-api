package validator

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestValidateDateWithEcho(t *testing.T) {
	e := echo.New()
	e.Validator = NewCustomValidator()
	ctx := e.NewContext(nil, nil)

	type input struct {
		Date string `validate:"required,YYYY-MM-DD"`
		Name string `validate:"required"`
	}

	testCases := []struct {
		name  string
		input input
		check func(output *input, err error)
	}{
		{
			name: "valid input data",
			input: input{
				Date: "2021-01-01",
				Name: "test",
			},
			check: func(output *input, err error) {
				require.NoError(t, err)
				require.Equal(t, "2021-01-01", output.Date)
				require.Equal(t, "test", output.Name)

			},
		},
		{
			name: "invalid input date",
			input: input{
				Date: "2021-01-32",
				Name: "test",
			},
			check: func(output *input, err error) {
				require.Error(t, err)
				require.Equal(t, "2021-01-32", output.Date)
				require.Equal(t, "test", output.Name)
			},
		},
		{
			name: "invalid input name",
			input: input{
				Date: "2021-01-01",
				Name: "",
			},
			check: func(output *input, err error) {
				require.Error(t, err)
				require.Equal(t, "2021-01-01", output.Date)
				require.Equal(t, "", output.Name)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := ctx.Validate(&tc.input)
			tc.check(&tc.input, err)
		})
	}
}

func TestDateFormat(t *testing.T) {
	validator := NewCustomValidator()

	type input struct {
		Date string `validate:"required,YYYY-MM-DD"`
	}

	testCases := []struct {
		name  string
		input input
		check func(err error)
	}{
		{
			name: "valid date",
			input: input{
				Date: "2021-01-01",
			},
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "invalid date",
			input: input{
				Date: "2021-01-32",
			},
			check: func(err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.Validate(tc.input)

			tc.check(err)
		})
	}
}

func TestValidateMultipleULIDWithEcho(t *testing.T) {
	e := echo.New()
	e.Validator = NewCustomValidator()
	ctx := e.NewContext(nil, nil)

	type input struct {
		Limit int32    `validate:"required"`
		IDs   []string `validate:"multipleULID"`
	}

	testCases := []struct {
		name  string
		input input
		check func(output *input, err error)
	}{
		{
			name: "valid ids",
			input: input{
				Limit: domain.DEFAULT_LIMIT,
				IDs:   []string{util.NewUlid(), util.NewUlid()},
			},
			check: func(output *input, err error) {
				require.NoError(t, err)
				require.Equal(t, domain.DEFAULT_LIMIT, output.Limit)
				require.Len(t, output.IDs, 2)
			},
		},
		{
			name: "empty ids",
			input: input{
				Limit: domain.DEFAULT_LIMIT,
				IDs:   []string{},
			},
			check: func(output *input, err error) {
				require.NoError(t, err)
				require.Equal(t, domain.DEFAULT_LIMIT, output.Limit)
				require.Len(t, output.IDs, 0)
			},
		},
		{
			name: "invalid ids",
			input: input{
				Limit: domain.DEFAULT_LIMIT,
				IDs:   []string{"invalid"},
			},
			check: func(output *input, err error) {
				require.Error(t, err)
				require.Equal(t, domain.DEFAULT_LIMIT, output.Limit)
				require.Len(t, output.IDs, 1)
			},
		},
		{
			name: "if ids length is greater than limit then return error",
			input: input{
				Limit: 1,
				IDs:   []string{util.NewUlid(), util.NewUlid()},
			},
			check: func(output *input, err error) {
				require.Error(t, err)
				require.Equal(t, int32(1), output.Limit)
				require.Len(t, output.IDs, 2)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := ctx.Validate(&tc.input)
			tc.check(&tc.input, err)
		})
	}
}

func TestMultipleULID(t *testing.T) {
	validator := NewCustomValidator()

	type input struct {
		Limit int32    `validate:"required"`
		IDs   []string `validate:"multipleULID"`
	}

	testCases := []struct {
		name       string
		buildInput func() input
		check      func(err error)
	}{
		{
			name: "valid ids",
			buildInput: func() input {
				validIds := make([]string, 0, 2)

				for i := 0; i < 2; i++ {
					ulid := util.NewUlid()
					validIds = append(validIds, ulid)
				}

				return input{
					Limit: domain.DEFAULT_LIMIT,
					IDs:   validIds,
				}
			},
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "empty ids",
			buildInput: func() input {
				return input{
					Limit: domain.DEFAULT_LIMIT,
					IDs:   []string{},
				}
			},
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "valid ids",
			buildInput: func() input {
				validIds := make([]string, 0, 2)

				for i := 0; i < 2; i++ {
					ulid := util.NewUlid()
					validIds = append(validIds, ulid)
				}

				inValidIds := "invalid"

				return input{
					Limit: domain.DEFAULT_LIMIT,
					IDs:   append(validIds, inValidIds),
				}
			},
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			name: "if ids length is greater than limit then return error",
			buildInput: func() input {
				validIds := make([]string, 0, 2)

				for i := 0; i < 2; i++ {
					ulid := util.NewUlid()
					validIds = append(validIds, ulid)
				}

				return input{
					Limit: 1,
					IDs:   validIds,
				}
			},
			check: func(err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			input := tc.buildInput()
			err := validator.Validate(input)

			tc.check(err)
		})
	}
}

func TestDishesWithEcho(t *testing.T) {
	e := echo.New()
	e.Validator = NewCustomValidator()
	ctx := e.NewContext(nil, nil)

	type input struct {
		Dishes []Dish `validate:"dishes"`
	}

	testCases := []struct {
		name  string
		input input
		check func(output *input, err error)
	}{
		{
			name: "valid dishes",
			input: input{
				Dishes: []Dish{
					{
						Name: "valid",
					},
					{
						Name: "valid",
					},
				},
			},
			check: func(output *input, err error) {
				require.NoError(t, err)
				require.Len(t, output.Dishes, 2)
			},
		},
		{
			name: "empty dishes",
			input: input{
				Dishes: []Dish{},
			},
			check: func(output *input, err error) {
				require.Error(t, err)
				require.Len(t, output.Dishes, 0)
			},
		},
		{
			name: "invalid value dishes",
			input: input{
				Dishes: []Dish{{Name: "valid"}, {Name: ""}},
			},
			check: func(output *input, err error) {
				require.Error(t, err)
				require.Len(t, output.Dishes, 2)
			},
		},
		{
			name: "max value dishes",
			input: input{
				Dishes: []Dish{{Name: "valid"}, {Name: util.RandomString(256)}},
			},
			check: func(output *input, err error) {
				require.Error(t, err)
				require.Len(t, output.Dishes, 2)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := ctx.Validate(&tc.input)
			tc.check(&tc.input, err)
		})
	}
}

func TestDishes(t *testing.T) {
	validator := NewCustomValidator()

	type input struct {
		Dishes []Dish `validate:"dishes"`
	}

	testCases := []struct {
		name  string
		input input
		check func(err error)
	}{
		{
			name: "valid dishes",
			input: input{
				Dishes: []Dish{{Name: "test"}},
			},
			check: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			name: "empty dishes",
			input: input{
				Dishes: []Dish{},
			},
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			name: "invalid value dishes",
			input: input{
				Dishes: []Dish{{Name: "valid"}, {Name: ""}},
			},
			check: func(err error) {
				require.Error(t, err)
			},
		},
		{
			name: "max value dishes",
			input: input{
				Dishes: []Dish{{Name: "valid"}, {Name: util.RandomString(256)}},
			},
			check: func(err error) {
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := validator.Validate(tc.input)

			tc.check(err)
		})
	}
}
