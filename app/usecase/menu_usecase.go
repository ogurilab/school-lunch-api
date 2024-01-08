package usecase

import (
	"context"
	"time"

	"github.com/ogurilab/school-lunch-api/domain"
)

type menuUsecase struct {
	menuRepo       domain.MenuRepository
	contextTimeout time.Duration
}

func NewMenuUsecase(mr domain.MenuRepository, timeout time.Duration) domain.MenuUsecase {
	return &menuUsecase{
		menuRepo:       mr,
		contextTimeout: timeout,
	}
}

func (mu *menuUsecase) Create(ctx context.Context, menu *domain.Menu) error {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	return mu.menuRepo.Create(ctx, menu)
}

func (mu *menuUsecase) GetByID(ctx context.Context, id string, city int32) (*domain.Menu, error) {

	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	return mu.menuRepo.GetByID(ctx, id, city)
}

func (mu *menuUsecase) Fetch(ctx context.Context, limit int32, offset int32, city int32) ([]*domain.Menu, error) {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	r, err := mu.menuRepo.Fetch(ctx, limit, offset, city)

	if err != nil {
		return nil, err
	}

	if len(r) == 0 {
		return []*domain.Menu{}, nil
	}

	return r, nil
}

func (mu *menuUsecase) GetByDate(ctx context.Context, offeredAt time.Time, city int32) (*domain.Menu, error) {

	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	return mu.menuRepo.GetByDate(ctx, offeredAt, city)
}

func (mu *menuUsecase) FetchByRangeDate(ctx context.Context, start, end time.Time, city int32, limit int32) ([]*domain.Menu, error) {

	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	r, err := mu.menuRepo.FetchByRangeDate(ctx, start, end, city, limit)

	if err != nil {
		return nil, err
	}

	if len(r) == 0 {
		return []*domain.Menu{}, nil
	}

	return r, nil
}

type menuWithDishesUsecase struct {
	menuRepo       domain.MenuWithDishesRepository
	contextTimeout time.Duration
}

func NewMenuWithDishesUsecase(mr domain.MenuWithDishesRepository, timeout time.Duration) domain.MenuWithDishesUsecase {
	return &menuWithDishesUsecase{
		menuRepo:       mr,
		contextTimeout: timeout,
	}
}

func (mu *menuWithDishesUsecase) GetByID(ctx context.Context, id string, city int32) (*domain.MenuWithDishes, error) {

	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	return mu.menuRepo.GetByID(ctx, id, city)
}

func (mu *menuWithDishesUsecase) Fetch(ctx context.Context, limit int32, offset int32, city int32) ([]*domain.MenuWithDishes, error) {

	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	r, err := mu.menuRepo.Fetch(ctx, limit, offset, city)

	if err != nil {
		return nil, err
	}

	if len(r) == 0 {
		return []*domain.MenuWithDishes{}, nil
	}

	return r, nil
}

func (mu *menuWithDishesUsecase) GetByDate(ctx context.Context, offeredAt time.Time, city int32) (*domain.MenuWithDishes, error) {

	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	return mu.menuRepo.GetByDate(ctx, offeredAt, city)
}

func (mu *menuWithDishesUsecase) FetchByRangeDate(ctx context.Context, start, end time.Time, city int32, limit int32) ([]*domain.MenuWithDishes, error) {

	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	r, err := mu.menuRepo.FetchByRangeDate(ctx, start, end, city, limit)

	if err != nil {
		return nil, err
	}

	if len(r) == 0 {
		return []*domain.MenuWithDishes{}, nil
	}

	return r, nil
}
