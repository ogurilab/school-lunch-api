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
