package usecase

import (
	"context"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
)

type dishUsecase struct {
	dishRepo       domain.DishRepository
	contextTimeout time.Duration
}

func NewDishUsecase(dr domain.DishRepository, timeout time.Duration) domain.DishUsecase {
	return &dishUsecase{
		dishRepo:       dr,
		contextTimeout: timeout,
	}
}

func (du *dishUsecase) Create(ctx context.Context, dish *domain.Dish, menuID string) error {
	ctx, cancel := context.WithTimeout(ctx, du.contextTimeout)
	defer cancel()

	return du.dishRepo.Create(ctx, dish, menuID)
}

func (du *dishUsecase) GetByID(ctx context.Context, id string) (*domain.Dish, error) {
	ctx, cancel := context.WithTimeout(ctx, du.contextTimeout)
	defer cancel()

	return du.dishRepo.GetByID(ctx, id)
}

func (du *dishUsecase) FetchByMenuID(ctx context.Context, menuID string) ([]*domain.Dish, error) {
	ctx, cancel := context.WithTimeout(ctx, du.contextTimeout)
	defer cancel()

	dishes, err := du.dishRepo.FetchByMenuID(ctx, menuID)

	if err != nil {
		return nil, err
	}

	if len(dishes) == 0 {
		return []*domain.Dish{}, nil
	}

	return dishes, nil
}

func (du *dishUsecase) Fetch(ctx context.Context, search string, limit int32, offset int32) ([]*domain.Dish, error) {
	ctx, cancel := context.WithTimeout(ctx, du.contextTimeout)
	defer cancel()

	if search != "" {
		like := "%" + search + "%"
		dishes, err := du.dishRepo.FetchByName(ctx, like, limit, offset)

		if err != nil {
			return nil, err
		}

		if len(dishes) == 0 {
			return []*domain.Dish{}, nil
		}

		return dishes, nil
	}

	dishes, err := du.dishRepo.Fetch(ctx, limit, offset)

	if err != nil {
		return nil, err
	}

	if len(dishes) == 0 {
		return []*domain.Dish{}, nil
	}

	return dishes, nil
}
