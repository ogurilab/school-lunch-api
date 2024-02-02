// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: menu_with_dishes.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getMenuWithDishes = `-- name: GetMenuWithDishes :many
SELECT m.id, m.offered_at, m.photo_url, m.created_at, m.elementary_school_calories, m.junior_high_school_calories, m.city_code,
  d.id AS dish_id,
  d.name AS dish_name
FROM (
    SELECT id, offered_at, photo_url, created_at, elementary_school_calories, junior_high_school_calories, city_code
    FROM menus
    WHERE menus.id = ?
      AND city_code = ?
  ) AS m
  INNER JOIN menu_dishes AS md ON m.id = md.menu_id
  INNER JOIN dishes AS d ON md.dish_id = d.id
ORDER BY d.id ASC
`

type GetMenuWithDishesParams struct {
	ID       string `json:"id"`
	CityCode int32  `json:"city_code"`
}

type GetMenuWithDishesRow struct {
	ID                       string         `json:"id"`
	OfferedAt                time.Time      `json:"offered_at"`
	PhotoUrl                 sql.NullString `json:"photo_url"`
	CreatedAt                time.Time      `json:"created_at"`
	ElementarySchoolCalories int32          `json:"elementary_school_calories"`
	JuniorHighSchoolCalories int32          `json:"junior_high_school_calories"`
	CityCode                 int32          `json:"city_code"`
	DishID                   string         `json:"dish_id"`
	DishName                 string         `json:"dish_name"`
}

func (q *Queries) GetMenuWithDishes(ctx context.Context, arg GetMenuWithDishesParams) ([]GetMenuWithDishesRow, error) {
	rows, err := q.db.QueryContext(ctx, getMenuWithDishes, arg.ID, arg.CityCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetMenuWithDishesRow{}
	for rows.Next() {
		var i GetMenuWithDishesRow
		if err := rows.Scan(
			&i.ID,
			&i.OfferedAt,
			&i.PhotoUrl,
			&i.CreatedAt,
			&i.ElementarySchoolCalories,
			&i.JuniorHighSchoolCalories,
			&i.CityCode,
			&i.DishID,
			&i.DishName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMenuWithDishes = `-- name: ListMenuWithDishes :many
SELECT m.id, m.offered_at, m.photo_url, m.created_at, m.elementary_school_calories, m.junior_high_school_calories, m.city_code,
  d.id AS dish_id,
  d.name AS dish_name
FROM (
    SELECT id, offered_at, photo_url, created_at, elementary_school_calories, junior_high_school_calories, city_code
    FROM menus AS m
    WHERE offered_at <= ?
    ORDER BY offered_at DESC
    LIMIT ? OFFSET ?
  ) AS m
  INNER JOIN menu_dishes AS md ON m.id = md.menu_id
  INNER JOIN dishes AS d ON md.dish_id = d.id
`

type ListMenuWithDishesParams struct {
	OfferedAt time.Time `json:"offered_at"`
	Limit     int32     `json:"limit"`
	Offset    int32     `json:"offset"`
}

type ListMenuWithDishesRow struct {
	ID                       string         `json:"id"`
	OfferedAt                time.Time      `json:"offered_at"`
	PhotoUrl                 sql.NullString `json:"photo_url"`
	CreatedAt                time.Time      `json:"created_at"`
	ElementarySchoolCalories int32          `json:"elementary_school_calories"`
	JuniorHighSchoolCalories int32          `json:"junior_high_school_calories"`
	CityCode                 int32          `json:"city_code"`
	DishID                   string         `json:"dish_id"`
	DishName                 string         `json:"dish_name"`
}

func (q *Queries) ListMenuWithDishes(ctx context.Context, arg ListMenuWithDishesParams) ([]ListMenuWithDishesRow, error) {
	rows, err := q.db.QueryContext(ctx, listMenuWithDishes, arg.OfferedAt, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListMenuWithDishesRow{}
	for rows.Next() {
		var i ListMenuWithDishesRow
		if err := rows.Scan(
			&i.ID,
			&i.OfferedAt,
			&i.PhotoUrl,
			&i.CreatedAt,
			&i.ElementarySchoolCalories,
			&i.JuniorHighSchoolCalories,
			&i.CityCode,
			&i.DishID,
			&i.DishName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMenuWithDishesByCity = `-- name: ListMenuWithDishesByCity :many
SELECT m.id, m.offered_at, m.photo_url, m.created_at, m.elementary_school_calories, m.junior_high_school_calories, m.city_code,
  d.id AS dish_id,
  d.name AS dish_name
FROM (
    SELECT id, offered_at, photo_url, created_at, elementary_school_calories, junior_high_school_calories, city_code
    FROM menus AS m
    WHERE city_code = ?
      AND offered_at <= ?
    ORDER BY offered_at DESC
    LIMIT ? OFFSET ?
  ) AS m
  INNER JOIN menu_dishes md ON m.id = md.menu_id
  INNER JOIN dishes d ON md.dish_id = d.id
`

type ListMenuWithDishesByCityParams struct {
	CityCode  int32     `json:"city_code"`
	OfferedAt time.Time `json:"offered_at"`
	Limit     int32     `json:"limit"`
	Offset    int32     `json:"offset"`
}

type ListMenuWithDishesByCityRow struct {
	ID                       string         `json:"id"`
	OfferedAt                time.Time      `json:"offered_at"`
	PhotoUrl                 sql.NullString `json:"photo_url"`
	CreatedAt                time.Time      `json:"created_at"`
	ElementarySchoolCalories int32          `json:"elementary_school_calories"`
	JuniorHighSchoolCalories int32          `json:"junior_high_school_calories"`
	CityCode                 int32          `json:"city_code"`
	DishID                   string         `json:"dish_id"`
	DishName                 string         `json:"dish_name"`
}

func (q *Queries) ListMenuWithDishesByCity(ctx context.Context, arg ListMenuWithDishesByCityParams) ([]ListMenuWithDishesByCityRow, error) {
	rows, err := q.db.QueryContext(ctx, listMenuWithDishesByCity,
		arg.CityCode,
		arg.OfferedAt,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListMenuWithDishesByCityRow{}
	for rows.Next() {
		var i ListMenuWithDishesByCityRow
		if err := rows.Scan(
			&i.ID,
			&i.OfferedAt,
			&i.PhotoUrl,
			&i.CreatedAt,
			&i.ElementarySchoolCalories,
			&i.JuniorHighSchoolCalories,
			&i.CityCode,
			&i.DishID,
			&i.DishName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
