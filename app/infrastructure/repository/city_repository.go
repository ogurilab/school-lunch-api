package repository

import (
	"context"

	"github.com/ogurilab/school-lunch-api/domain"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
)

type cityRepository struct {
	query db.Query
}

func NewCityRepository(query db.Query) domain.CityRepository {
	return &cityRepository{
		query: query,
	}
}

func (r *cityRepository) GetByCityCode(ctx context.Context, code int32) (*domain.City, error) {

	result, err := r.query.GetCity(ctx, code)

	if err != nil {
		return nil, err
	}

	city := domain.NewCity(
		result.CityCode,
		result.CityName,
		result.PrefectureCode,
		result.PrefectureName,
	)

	return city, nil

}

func (r *cityRepository) Fetch(ctx context.Context, limit int32, offset int32, search string) ([]*domain.City, error) {

	if search != "" {
		arg := db.ListCitiesByNameParams{
			Limit:    limit,
			Offset:   offset,
			CityName: search,
		}

		result, err := r.query.ListCitiesByName(ctx, arg)
		if err != nil {
			return nil, err
		}

		var cities []*domain.City

		for _, city := range result {
			cities = append(cities, domain.NewCity(
				city.CityCode,
				city.CityName,
				city.PrefectureCode,
				city.PrefectureName,
			))
		}

		return cities, nil
	}

	arg := db.ListCitiesParams{
		Limit:  limit,
		Offset: offset,
	}

	result, err := r.query.ListCities(ctx, arg)
	if err != nil {
		return nil, err
	}

	var cities []*domain.City

	for _, city := range result {
		cities = append(cities, domain.NewCity(
			city.CityCode,
			city.CityName,
			city.PrefectureCode,
			city.PrefectureName,
		))
	}

	return cities, nil

}

func (r *cityRepository) FetchByPrefectureCode(ctx context.Context, limit int32, offset int32, prefectureCode int32) ([]*domain.City, error) {
	arg := db.ListCitiesByPrefectureParams{
		PrefectureCode: prefectureCode,
		Limit:          limit,
		Offset:         offset,
	}

	result, err := r.query.ListCitiesByPrefecture(ctx, arg)

	if err != nil {
		return nil, err
	}

	var cities []*domain.City

	for _, city := range result {
		cities = append(cities, domain.NewCity(
			city.CityCode,
			city.CityName,
			city.PrefectureCode,
			city.PrefectureName,
		))
	}

	return cities, nil
}
