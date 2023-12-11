package usecase

import (
	"context"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
)

type cityUsecase struct {
	cityRepo       domain.CityRepository
	contextTimeout time.Duration
}

func NewCityUsecase(cr domain.CityRepository, timeout time.Duration) domain.CityUsecase {

	return &cityUsecase{
		cityRepo:       cr,
		contextTimeout: timeout,
	}

}

func (cu *cityUsecase) GetByCityCode(ctx context.Context, code int32) (*domain.City, error) {

	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	return cu.cityRepo.GetByCityCode(ctx, code)
}

func (cu *cityUsecase) Fetch(ctx context.Context, limit int32, offset int32, search string) ([]*domain.City, error) {

	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	r, err := cu.cityRepo.Fetch(ctx, limit, offset, search)

	if err != nil {
		return nil, err
	}

	if len(r) == 0 {
		return []*domain.City{}, nil
	}

	return r, nil
}

func (cu *cityUsecase) FetchByPrefectureCode(ctx context.Context, limit int32, offset int32, prefectureCode int32) ([]*domain.City, error) {

	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	r, err := cu.cityRepo.FetchByPrefectureCode(ctx, limit, offset, prefectureCode)

	if err != nil {
		return nil, err
	}

	if len(r) == 0 {
		return []*domain.City{}, nil
	}

	return r, nil
}
