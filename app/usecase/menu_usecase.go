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

func (mu *menuUsecase) FetchByCity(ctx context.Context, limit int32, offset int32, offered time.Time, city int32) ([]*domain.Menu, error) {
	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()

	r, err := mu.menuRepo.FetchByCity(ctx, limit, offset, offered, city)

	if err != nil {
		return nil, err
	}

	if len(r) == 0 {
		return []*domain.Menu{}, nil
	}

	return r, nil
}

func (mu *menuUsecase) Fetch(ctx context.Context, limit int32, offset int32, offered time.Time, ids []string) ([]*domain.Menu, error) {

	ctx, cancel := context.WithTimeout(ctx, mu.contextTimeout)
	defer cancel()
	isId := len(ids) > 0

	if isId {
		menus, err := mu.menuRepo.FetchByIDs(ctx, limit, offset, offered, ids)

		if err != nil {
			return nil, err
		}

		if len(menus) == 0 {
			return []*domain.Menu{}, nil
		}

		return menus, nil
	}

	menus, err := mu.menuRepo.Fetch(ctx, limit, offset, offered)

	if err != nil {
		return nil, err
	}

	if len(menus) == 0 {
		return []*domain.Menu{}, nil
	}

	return menus, nil
}
