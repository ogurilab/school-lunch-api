package domain

import (
	"context"

	"github.com/labstack/echo/v4"
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
	FetchByName(ctx context.Context, limit int32, offset int32, search string) ([]*City, error)
	Fetch(ctx context.Context, limit int32, offset int32) ([]*City, error)
	FetchByPrefectureCode(ctx context.Context, limit int32, offset int32, prefectureCode int32) ([]*City, error)
}

type CityUsecase interface {
	GetByCityCode(ctx context.Context, code int32) (*City, error)
	Fetch(ctx context.Context, limit int32, offset int32, search string) ([]*City, error)
	FetchByPrefectureCode(ctx context.Context, limit int32, offset int32, prefectureCode int32) ([]*City, error)
}

type CityController interface {
	GetByCityCode(c echo.Context) error
	Fetch(c echo.Context) error
	FetchByPrefectureCode(c echo.Context) error
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
