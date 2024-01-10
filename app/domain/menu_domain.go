package domain

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
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

type MenuController interface {
	Create(c echo.Context) error
	GetByID(c echo.Context) error
	FetchByCity(c echo.Context) error
}

func (m *Menu) MarshalJSON() ([]byte, error) {
	type Alias Menu

	type Date struct {
		ID        string  `json:"id"`
		OfferedAt string  `json:"offered_at"`
		PhotoUrl  *string `json:"photo_url"`
		*Alias
	}

	return json.Marshal(
		&Date{
			ID:        m.ID,
			OfferedAt: m.OfferedAt.Format("2006-01-02"),
			PhotoUrl:  util.NullStringToPointer(m.PhotoUrl),
			Alias:     (*Alias)(m),
		},
	)
}

func (m *Menu) UnmarshalJSON(data []byte) error {

	type Alias Menu
	aux := &struct {
		OfferedAt string  `json:"offered_at"`
		PhotoUrl  *string `json:"photo_url"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	offeredAt, err := time.Parse("2006-01-02", aux.OfferedAt)

	if err != nil {
		return err
	}

	m.OfferedAt = offeredAt
	m.PhotoUrl = util.PointerToNullString(aux.PhotoUrl)

	return nil
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

type MenuWithDishesController interface {
	GetByID(c echo.Context) error
	FetchByCity(c echo.Context) error
	Fetch(c echo.Context) error
}

func (m *MenuWithDishes) MarshalJSON() ([]byte, error) {

	type Date struct {
		ID                       string  `json:"id"`
		OfferedAt                string  `json:"offered_at"`
		PhotoUrl                 *string `json:"photo_url"`
		ElementarySchoolCalories int32   `json:"elementary_school_calories"`
		JuniorHighSchoolCalories int32   `json:"junior_high_school_calories"`
		CityCode                 int32   `json:"city_code"`
		Dishes                   []*Dish `json:"dishes"`
	}

	if m.Dishes == nil {
		m.Dishes = []*Dish{}
	}

	return json.Marshal(&Date{
		Dishes:                   m.Dishes,
		OfferedAt:                m.OfferedAt.Format("2006-01-02"),
		ID:                       m.ID,
		PhotoUrl:                 util.NullStringToPointer(m.PhotoUrl),
		ElementarySchoolCalories: m.ElementarySchoolCalories,
		JuniorHighSchoolCalories: m.JuniorHighSchoolCalories,
		CityCode:                 m.CityCode,
	})
}

func (m *MenuWithDishes) UnmarshalJSON(data []byte) error {

	type MenuAlias Menu

	aux := &struct {
		Dishes    []*Dish `json:"dishes"`
		OfferedAt string  `json:"offered_at"`
		PhotoUrl  *string `json:"photo_url"`
		*MenuAlias
	}{
		Dishes:    m.Dishes,
		OfferedAt: m.OfferedAt.Format("2006-01-02"),
		PhotoUrl:  util.NullStringToPointer(m.PhotoUrl),
		MenuAlias: (*MenuAlias)(&m.Menu),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	offeredAt, err := time.Parse("2006-01-02", aux.OfferedAt)

	if err != nil {
		return err
	}

	m.OfferedAt = offeredAt
	m.PhotoUrl = util.PointerToNullString(aux.PhotoUrl)

	if aux.Dishes != nil {
		m.Dishes = aux.Dishes
	} else {
		m.Dishes = []*Dish{}
	}

	return nil
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
