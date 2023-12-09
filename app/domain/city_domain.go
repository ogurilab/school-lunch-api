package domain

import (
	"context"
)

type City struct {
	CityCode                 int32  `json:"city_code"`
	CityName                 string `json:"city_name"`
	PrefectureCode           int32  `json:"prefecture_code"`
	PrefectureName           string `json:"prefecture_name"`
	SchoolLunchInfoAvailable bool   `json:"school_lunch_info_available"`
}

type CityRepository interface {
	GetByCityCode(ctx context.Context, code int32) (*City, error)
	Fetch(ctx context.Context, limit int32, offset int32, search string) ([]*City, error)
	FetchByPrefectureCode(ctx context.Context, limit int32, offset int32, prefectureCode int32) ([]*City, error)
}

type CityUsecase interface {
	GetByCityCode(ctx context.Context, code int32) (*City, error)
	Fetch(ctx context.Context, limit int32, offset int32, search string) ([]*City, error)
	FetchByPrefectureCode(ctx context.Context, limit int32, offset int32, prefectureCode int32) ([]*City, error)
}

func NewCity(
	cityCode int32,
	cityName string,
	prefectureCode int32,
	prefectureName string,
) *City {
	return &City{
		CityCode:       cityCode,
		CityName:       cityName,
		PrefectureCode: prefectureCode,
		PrefectureName: prefectureName,
	}
}
