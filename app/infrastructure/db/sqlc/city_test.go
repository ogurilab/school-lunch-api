package db

import (
	"context"
	"testing"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

func TestCreateCity(t *testing.T) {
	createRandomCity(t)
}

func TestUpdateAvailable(t *testing.T) {
	city := createRandomCity(t)

	err := testQuery.UpdateAvailable(context.Background(), city.CityCode)

	require.NoError(t, err)

	city2, err := testQuery.GetCity(context.Background(), city.CityCode)

	require.NoError(t, err)
	require.NotEmpty(t, city2)

	require.Equal(t, city.CityCode, city2.CityCode)
	require.Equal(t, city.CityName, city2.CityName)
	require.Equal(t, city.PrefectureCode, city2.PrefectureCode)
	require.Equal(t, city.PrefectureName, city2.PrefectureName)
	require.Equal(t, city2.SchoolLunchInfoAvailable, true)
}

func TestGetCity(t *testing.T) {
	city := createRandomCity(t)
	city2, err := testQuery.GetCity(context.Background(), city.CityCode)

	require.NoError(t, err)
	require.NotEmpty(t, city2)

	require.Equal(t, city.CityCode, city2.CityCode)
	require.Equal(t, city.CityName, city2.CityName)
	require.Equal(t, city.PrefectureCode, city2.PrefectureCode)
	require.Equal(t, city.PrefectureName, city2.PrefectureName)
	require.Equal(t, city.SchoolLunchInfoAvailable, city2.SchoolLunchInfoAvailable)
}

func TestListCities(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCity(t)
	}

	arg := ListCitiesParams{
		Limit:  5,
		Offset: 5,
	}

	cities, err := testQuery.ListCities(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, cities, 5)

	for _, city := range cities {
		require.NotEmpty(t, city)
	}
}

func TestListCitiesByName(t *testing.T) {

	var names []string

	for i := 0; i < 10; i++ {
		city := createRandomCity(t)
		names = append(names, city.CityName)
	}

	arg := ListCitiesByNameParams{
		Limit:    5,
		Offset:   0,
		CityName: names[0],
	}

	cities, err := testQuery.ListCitiesByName(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, cities, 1)

	for _, city := range cities {
		require.NotEmpty(t, city)
		require.Equal(t, city.CityName, names[0])
	}
}

func TestListCitiesByPrefecture(t *testing.T) {

	var prefectureCodes []int32

	for i := 0; i < 10; i++ {
		city := createRandomCity(t)
		prefectureCodes = append(prefectureCodes, city.PrefectureCode)
	}

	arg := ListCitiesByPrefectureParams{
		Limit:          5,
		Offset:         0,
		PrefectureCode: prefectureCodes[0],
	}

	cities, err := testQuery.ListCitiesByPrefecture(context.Background(), arg)

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(cities), 1)

	for _, city := range cities {
		require.NotEmpty(t, city)
		require.Equal(t, city.PrefectureCode, prefectureCodes[0])
	}
}

func createRandomCity(t *testing.T) *domain.City {

	cityCode := util.RandomCityCode()

	arg := CreateCityParams{
		CityCode:       cityCode,
		CityName:       util.RandomString(10),
		PrefectureCode: util.RandomInt32(),
	}

	err := testQuery.CreateCity(context.Background(), arg)

	require.NoError(t, err)

	city, err := testQuery.GetCity(context.Background(), cityCode)

	require.NoError(t, err)
	require.NotEmpty(t, city)

	require.Equal(t, arg.CityCode, city.CityCode)
	require.Equal(t, arg.CityName, city.CityName)
	require.Equal(t, arg.PrefectureCode, city.PrefectureCode)
	require.Equal(t, city.SchoolLunchInfoAvailable, false)

	return domain.NewCity(
		city.CityCode,
		city.CityName,
		city.PrefectureCode,
		city.PrefectureName,
	)
}
