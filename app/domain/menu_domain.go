package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/ogurilab/school-lunch-api/util"
)

type Menu struct {
	ID                       string         `json:"id"`
	OfferedAt                time.Time      `json:"offered_at"`
	PhotoUrl                 sql.NullString `json:"photo_url"`
	ElementarySchoolCalories int32          `json:"elementary_school_calories"`
	JuniorHighSchoolCalories int32          `json:"junior_high_school_calories"`
	CityCode                 int32          `json:"city_code"`
}

type MenuWithDishes struct {
	Menu
	Dishes []*Dish `json:"dishes"`
}

type MenuRepository interface {
	Create(ctx context.Context, menu *Menu) error
	GetByID(ctx context.Context, id string, city int32) (*Menu, error)
	Fetch(ctx context.Context, limit int32, offset int32, city int32) ([]*Menu, error)
	GetByDate(ctx context.Context, offeredAt time.Time, city int32) (*Menu, error)
	FetchByRangeDate(ctx context.Context, start, end time.Time, city int32) ([]*Menu, error)

	// MenuWithDishes
	GetByIDWithDishes(ctx context.Context, id string, city int32) (*MenuWithDishes, error)
	FetchWithDishes(ctx context.Context, limit int32, offset int32, city int32) ([]*MenuWithDishes, error)
	GetByDateWithDishes(ctx context.Context, offeredAt time.Time, city int32) (*MenuWithDishes, error)
	FetchByRangeDateWithDishes(ctx context.Context, start, end time.Time, city int32) ([]*MenuWithDishes, error)
}

type MenuUsecase interface {
	Create(ctx context.Context, menu *Menu) error
	GetByID(ctx context.Context, id string, city int32) (*Menu, error)
	Fetch(ctx context.Context, limit int32, offset int32, city int32) ([]*Menu, error)
	GetByDate(ctx context.Context, offeredAt time.Time, city int32) (*Menu, error)
	FetchByRangeDate(ctx context.Context, start, end time.Time, city int32) ([]*Menu, error)

	// MenuWithDishes
	GetByIDWithDishes(ctx context.Context, id string, city int32) (*MenuWithDishes, error)
	FetchWithDishes(ctx context.Context, limit int32, offset int32, city int32) ([]*MenuWithDishes, error)
	GetByDateWithDishes(ctx context.Context, offeredAt time.Time, city int32) (*MenuWithDishes, error)
	FetchByRangeDateWithDishes(ctx context.Context, start, end time.Time, city int32) ([]*MenuWithDishes, error)
}

func newMenu(
	id string,
	offeredAt time.Time,
	photoUrl sql.NullString,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
	cityCode int32,
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
		CityCode:                 cityCode,
	}, nil
}

func ReNewMenu(
	id string,
	offeredAt time.Time,
	photoUrl sql.NullString,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
	cityCode int32,
) (*Menu, error) {
	return newMenu(
		id,
		offeredAt,
		photoUrl,
		elementarySchoolCalories,
		juniorHighSchoolCalories,
		cityCode,
	)
}

func NewMenu(
	offeredAt time.Time,
	photoUrl sql.NullString,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
	cityCode int32,
) (*Menu, error) {
	id := util.NewUlid()
	return newMenu(
		id,
		offeredAt,
		photoUrl,
		elementarySchoolCalories,
		juniorHighSchoolCalories,
		cityCode,
	)
}
