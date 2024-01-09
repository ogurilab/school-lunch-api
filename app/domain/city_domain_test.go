package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCity(t *testing.T) {

	type input struct {
		cityCode       int32
		cityName       string
		prefectureCode int32
		prefectureName string
	}
	testCases := []struct {
		name       string
		createStub func() input
		check      func(*City, error)
	}{
		{
			name: "OK",
			createStub: func() input {
				return input{
					cityCode:       1,
					cityName:       "cityName",
					prefectureCode: 1,
					prefectureName: "prefectureName",
				}
			},
			check: func(c *City, err error) {
				require.NoError(t, err)
				require.NotNil(t, c)
				require.Equal(t, int32(1), c.CityCode)
				require.Equal(t, "cityName", c.CityName)
				require.Equal(t, int32(1), c.PrefectureCode)
				require.Equal(t, "prefectureName", c.PrefectureName)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.createStub()
			c := NewCity(input.cityCode, input.cityName, input.prefectureCode, input.prefectureName)
			tc.check(c, nil)
		})
	}

}
