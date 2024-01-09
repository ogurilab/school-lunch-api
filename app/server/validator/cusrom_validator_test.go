package validator

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestValidateWithEcho(t *testing.T) {
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
