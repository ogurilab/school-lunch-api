// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"database/sql"
	"time"
)

type Allergen struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Dish struct {
	ID        string         `json:"id"`
	MenuID    string         `json:"menu_id"`
	Name      sql.NullString `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
}

type DishesAllergen struct {
	AllergenID int32  `json:"allergen_id"`
	DishID     string `json:"dish_id"`
}

type Menu struct {
	ID string `json:"id"`
	// 給食の提供日
	OfferedAt                time.Time      `json:"offered_at"`
	RegionID                 int32          `json:"region_id"`
	PhotoUrl                 sql.NullString `json:"photo_url"`
	WikimediaCommonsUrl      sql.NullString `json:"wikimedia_commons_url"`
	CreatedAt                time.Time      `json:"created_at"`
	ElementarySchoolCalories sql.NullInt32  `json:"elementary_school_calories"`
	JuniorHighSchoolCalories sql.NullInt32  `json:"junior_high_school_calories"`
}

type Region struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
