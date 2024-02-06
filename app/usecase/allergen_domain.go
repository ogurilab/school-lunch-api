package usecase

import (
	"context"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
)

type allergenUsecase struct {
	allergenRepo   domain.AllergenRepository
	dishRepo       domain.DishRepository
	contextTimeout time.Duration
}

func NewAllergenUsecase(ar domain.AllergenRepository, dr domain.DishRepository, timeout time.Duration) domain.AllergenUsecase {
	return &allergenUsecase{
		allergenRepo:   ar,
		dishRepo:       dr,
		contextTimeout: timeout,
	}
}

func (au *allergenUsecase) FetchByDishID(ctx context.Context, dishID string) ([]*domain.Allergen, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	allergens, err := au.allergenRepo.FetchByDishID(ctx, dishID)

	if err != nil {
		return nil, err
	}

	if len(allergens) == 0 {
		return []*domain.Allergen{}, nil
	}

	return allergens, nil
}

func (au *allergenUsecase) FetchByMenuID(ctx context.Context, menuID string) ([]*domain.Allergen, error) {
	dishes, err := au.dishRepo.FetchByMenuID(ctx, menuID)

	if err != nil {
		return nil, err
	}

	if len(dishes) == 0 {
		return []*domain.Allergen{}, nil
	}

	dishIDs := make([]string, 0, len(dishes))

	for _, dish := range dishes {
		dishIDs = append(dishIDs, dish.ID)
	}

	allergens, err := au.allergenRepo.FetchInDish(ctx, dishIDs)

	if err != nil {
		return nil, err
	}

	if len(allergens) == 0 {
		return []*domain.Allergen{}, nil
	}

	return allergens, nil
}
