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
	FetchByCity(ctx context.Context, limit int32, offset int32, offered time.Time, city int32) ([]*Menu, error)
}

type MenuUsecase interface {
	Create(ctx context.Context, menu *Menu) error
	GetByID(ctx context.Context, id string, city int32) (*Menu, error)
	FetchByCity(ctx context.Context, limit int32, offset int32, offered time.Time, city int32) ([]*Menu, error)
}

/************************
 * MenuWithDishes
 ************************/

type MenuWithDishesRepository interface {
	GetByID(ctx context.Context, id string, city int32) (*MenuWithDishes, error)
	FetchByCity(ctx context.Context, limit int32, offset int32, offered time.Time, city int32) ([]*MenuWithDishes, error)
	Fetch(ctx context.Context, limit int32, offset int32, offered time.Time) ([]*MenuWithDishes, error)
}

type MenuWithDishesUsecase interface {
	GetByID(ctx context.Context, id string, city int32) (*MenuWithDishes, error)
	FetchByCity(ctx context.Context, limit int32, offset int32, offered time.Time, city int32) ([]*MenuWithDishes, error)
	Fetch(ctx context.Context, limit int32, offset int32, offered time.Time) ([]*MenuWithDishes, error)
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

func ReNewMenuWithDishes(
	id string,
	offeredAt time.Time,
	photoUrl sql.NullString,
	elementarySchoolCalories int32,
	juniorHighSchoolCalories int32,
	cityCode int32,
	dishes []*Dish,
) (*MenuWithDishes, error) {
	menu, err := ReNewMenu(
		id,
		offeredAt,
		photoUrl,
		elementarySchoolCalories,
		juniorHighSchoolCalories,
		cityCode,
	)

	if err != nil {
		return nil, err
	}

	return &MenuWithDishes{
		Menu:   *menu,
		Dishes: dishes,
	}, nil
}
