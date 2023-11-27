package domain

import (
	"context"
	"time"

	"github.com/ogurilab/school-lunch-api/util"
)

type Menu struct {
	ID                       string    `json:"id"`
	OfferedAt                time.Time `json:"offered_at"`
	PhotoUrl                 string    `json:"photo_url"`
	ElementarySchoolCalories int32     `json:"elementary_school_calories"`
	JuniorHighSchoolCalories int32     `json:"junior_high_school_calories"`
}

type MenuWithDishes struct {
	Menu
	Dishes []*Dish `json:"dishes"`
}

type MenuRepository interface {
	Create(ctx context.Context, menu *Menu) error
	GetByID(ctx context.Context, id string) (*Menu, error)
	Fetch(ctx context.Context, limit int32, offset int32) ([]*Menu, error)
	GetByDate(ctx context.Context, offeredAt time.Time) (*Menu, error)
	FetchByRangeDate(ctx context.Context, start, end time.Time) ([]*Menu, error)

	// MenuWithDishes
	GetByIDWithDishes(ctx context.Context, id string) (*MenuWithDishes, error)
	FetchWithDishes(ctx context.Context, limit int32, offset int32) ([]*MenuWithDishes, error)
	GetByDateWithDishes(ctx context.Context, offeredAt time.Time) (*MenuWithDishes, error)
	FetchByRangeDateWithDishes(ctx context.Context, start, end time.Time) ([]*MenuWithDishes, error)
}

type MenuUsecase interface {
	Create(ctx context.Context, menu *Menu) error
	GetByID(ctx context.Context, id string) (*Menu, error)
	Fetch(ctx context.Context, limit int32, offset int32) ([]*Menu, error)
	GetByDate(ctx context.Context, offeredAt time.Time) (*Menu, error)
	FetchByRangeDate(ctx context.Context, start, end time.Time) ([]*Menu, error)

	// MenuWithDishes
	GetByIDWithDishes(ctx context.Context, id string) (*MenuWithDishes, error)
	FetchWithDishes(ctx context.Context, limit int32, offset int32) ([]*MenuWithDishes, error)
	GetByDateWithDishes(ctx context.Context, offeredAt time.Time) (*MenuWithDishes, error)
	FetchByRangeDateWithDishes(ctx context.Context, start, end time.Time) ([]*MenuWithDishes, error)
}

func newMenu(
	id string,
	offeredAt time.Time,
	photoUrl string,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
) (*Menu, error) {

	if _, err := util.ParseUlid(id); err != nil {
		return nil, err
	}

	return &Menu{
		ID:                       id,
		OfferedAt:                offeredAt,
		PhotoUrl:                 photoUrl,
		ElementarySchoolCalories: elementarySchoolCalories,
		JuniorHighSchoolCalories: juniorHighSchoolCalories,
	}, nil
}

func ReNewMenu(
	id string,
	offeredAt time.Time,
	photoUrl string,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
) (*Menu, error) {
	return newMenu(
		id,
		offeredAt,
		photoUrl,
		elementarySchoolCalories,
		juniorHighSchoolCalories,
	)
}

func NewMenu(
	offeredAt time.Time,
	photoUrl string,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
) (*Menu, error) {
	id := util.NewUlid()
	return newMenu(
		id,
		offeredAt,
		photoUrl,
		elementarySchoolCalories,
		juniorHighSchoolCalories,
	)
}
