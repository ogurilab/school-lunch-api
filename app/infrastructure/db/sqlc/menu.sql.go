// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: menu.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const createMenu = `-- name: CreateMenu :exec
INSERT INTO menus (
    id,
    offered_at,
    photo_url,
    elementary_school_calories,
    junior_high_school_calories,
    city_code
  )
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
  )
`

type CreateMenuParams struct {
	ID                       string         `json:"id"`
	OfferedAt                time.Time      `json:"offered_at"`
	PhotoUrl                 sql.NullString `json:"photo_url"`
	ElementarySchoolCalories int32          `json:"elementary_school_calories"`
	JuniorHighSchoolCalories int32          `json:"junior_high_school_calories"`
	CityCode                 int32          `json:"city_code"`
}

func (q *Queries) CreateMenu(ctx context.Context, arg CreateMenuParams) error {
	_, err := q.db.ExecContext(ctx, createMenu,
		arg.ID,
		arg.OfferedAt,
		arg.PhotoUrl,
		arg.ElementarySchoolCalories,
		arg.JuniorHighSchoolCalories,
		arg.CityCode,
	)
	return err
}

const getMenu = `-- name: GetMenu :one
SELECT id, offered_at, photo_url, created_at, elementary_school_calories, junior_high_school_calories, city_code
FROM menus
WHERE id = ?
  AND city_code = ?
`

type GetMenuParams struct {
	ID       string `json:"id"`
	CityCode int32  `json:"city_code"`
}

func (q *Queries) GetMenu(ctx context.Context, arg GetMenuParams) (Menu, error) {
	row := q.db.QueryRowContext(ctx, getMenu, arg.ID, arg.CityCode)
	var i Menu
	err := row.Scan(
		&i.ID,
		&i.OfferedAt,
		&i.PhotoUrl,
		&i.CreatedAt,
		&i.ElementarySchoolCalories,
		&i.JuniorHighSchoolCalories,
		&i.CityCode,
	)
	return i, err
}

const getMenuByOfferedAt = `-- name: GetMenuByOfferedAt :one
SELECT id, offered_at, photo_url, created_at, elementary_school_calories, junior_high_school_calories, city_code
FROM menus
WHERE offered_at = ?
  AND city_code = ?
`

type GetMenuByOfferedAtParams struct {
	OfferedAt time.Time `json:"offered_at"`
	CityCode  int32     `json:"city_code"`
}

func (q *Queries) GetMenuByOfferedAt(ctx context.Context, arg GetMenuByOfferedAtParams) (Menu, error) {
	row := q.db.QueryRowContext(ctx, getMenuByOfferedAt, arg.OfferedAt, arg.CityCode)
	var i Menu
	err := row.Scan(
		&i.ID,
		&i.OfferedAt,
		&i.PhotoUrl,
		&i.CreatedAt,
		&i.ElementarySchoolCalories,
		&i.JuniorHighSchoolCalories,
		&i.CityCode,
	)
	return i, err
}

const getMenuWithDishes = `-- name: GetMenuWithDishes :one
SELECT m.id, m.offered_at, m.photo_url, m.created_at, m.elementary_school_calories, m.junior_high_school_calories, m.city_code,
  JSON_ARRAYAGG(
    JSON_OBJECT(
      'id',
      d.id,
      'name',
      d.name,
      'menu_id',
      d.menu_id
    )
  ) AS dishes
FROM menus AS m
  LEFT JOIN dishes AS d ON m.id = d.menu_id
WHERE m.id = ?
  AND m.city_code = ?
GROUP BY m.id
`

type GetMenuWithDishesParams struct {
	ID       string `json:"id"`
	CityCode int32  `json:"city_code"`
}

type GetMenuWithDishesRow struct {
	ID                       string          `json:"id"`
	OfferedAt                time.Time       `json:"offered_at"`
	PhotoUrl                 sql.NullString  `json:"photo_url"`
	CreatedAt                time.Time       `json:"created_at"`
	ElementarySchoolCalories int32           `json:"elementary_school_calories"`
	JuniorHighSchoolCalories int32           `json:"junior_high_school_calories"`
	CityCode                 int32           `json:"city_code"`
	Dishes                   json.RawMessage `json:"dishes"`
}

func (q *Queries) GetMenuWithDishes(ctx context.Context, arg GetMenuWithDishesParams) (GetMenuWithDishesRow, error) {
	row := q.db.QueryRowContext(ctx, getMenuWithDishes, arg.ID, arg.CityCode)
	var i GetMenuWithDishesRow
	err := row.Scan(
		&i.ID,
		&i.OfferedAt,
		&i.PhotoUrl,
		&i.CreatedAt,
		&i.ElementarySchoolCalories,
		&i.JuniorHighSchoolCalories,
		&i.CityCode,
		&i.Dishes,
	)
	return i, err
}

const getMenuWithDishesByOfferedAt = `-- name: GetMenuWithDishesByOfferedAt :one
SELECT m.id, m.offered_at, m.photo_url, m.created_at, m.elementary_school_calories, m.junior_high_school_calories, m.city_code,
  JSON_ARRAYAGG(
    JSON_OBJECT(
      'id',
      d.id,
      'name',
      d.name,
      'menu_id',
      d.menu_id
    )
  ) AS dishes
FROM menus AS m
  LEFT JOIN dishes AS d ON m.id = d.menu_id
WHERE m.offered_at = ?
  AND m.city_code = ?
GROUP BY m.id
`

type GetMenuWithDishesByOfferedAtParams struct {
	OfferedAt time.Time `json:"offered_at"`
	CityCode  int32     `json:"city_code"`
}

type GetMenuWithDishesByOfferedAtRow struct {
	ID                       string          `json:"id"`
	OfferedAt                time.Time       `json:"offered_at"`
	PhotoUrl                 sql.NullString  `json:"photo_url"`
	CreatedAt                time.Time       `json:"created_at"`
	ElementarySchoolCalories int32           `json:"elementary_school_calories"`
	JuniorHighSchoolCalories int32           `json:"junior_high_school_calories"`
	CityCode                 int32           `json:"city_code"`
	Dishes                   json.RawMessage `json:"dishes"`
}

func (q *Queries) GetMenuWithDishesByOfferedAt(ctx context.Context, arg GetMenuWithDishesByOfferedAtParams) (GetMenuWithDishesByOfferedAtRow, error) {
	row := q.db.QueryRowContext(ctx, getMenuWithDishesByOfferedAt, arg.OfferedAt, arg.CityCode)
	var i GetMenuWithDishesByOfferedAtRow
	err := row.Scan(
		&i.ID,
		&i.OfferedAt,
		&i.PhotoUrl,
		&i.CreatedAt,
		&i.ElementarySchoolCalories,
		&i.JuniorHighSchoolCalories,
		&i.CityCode,
		&i.Dishes,
	)
	return i, err
}

const listMenuWithDishes = `-- name: ListMenuWithDishes :many
SELECT m.id, m.offered_at, m.photo_url, m.created_at, m.elementary_school_calories, m.junior_high_school_calories, m.city_code,
  JSON_ARRAYAGG(
    JSON_OBJECT(
      'id',
      d.id,
      'name',
      d.name,
      'menu_id',
      d.menu_id
    )
  ) AS dishes
FROM menus AS m
  LEFT JOIN dishes AS d ON m.id = d.menu_id
WHERE m.city_code = ?
GROUP BY m.id
ORDER BY offered_at
LIMIT ? OFFSET ?
`

type ListMenuWithDishesParams struct {
	CityCode int32 `json:"city_code"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

type ListMenuWithDishesRow struct {
	ID                       string          `json:"id"`
	OfferedAt                time.Time       `json:"offered_at"`
	PhotoUrl                 sql.NullString  `json:"photo_url"`
	CreatedAt                time.Time       `json:"created_at"`
	ElementarySchoolCalories int32           `json:"elementary_school_calories"`
	JuniorHighSchoolCalories int32           `json:"junior_high_school_calories"`
	CityCode                 int32           `json:"city_code"`
	Dishes                   json.RawMessage `json:"dishes"`
}

func (q *Queries) ListMenuWithDishes(ctx context.Context, arg ListMenuWithDishesParams) ([]ListMenuWithDishesRow, error) {
	rows, err := q.db.QueryContext(ctx, listMenuWithDishes, arg.CityCode, arg.Limit, arg.Offset)
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
			&i.Dishes,
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

const listMenuWithDishesByOfferedAt = `-- name: ListMenuWithDishesByOfferedAt :many
SELECT m.id, m.offered_at, m.photo_url, m.created_at, m.elementary_school_calories, m.junior_high_school_calories, m.city_code,
  JSON_ARRAYAGG(
    JSON_OBJECT(
      'id',
      d.id,
      'name',
      d.name,
      'menu_id',
      d.menu_id
    )
  ) AS dishes
FROM menus AS m
  LEFT JOIN dishes AS d ON m.id = d.menu_id
WHERE m.offered_at >= ?
  AND m.offered_at <= ?
  AND m.city_code = ?
GROUP BY m.id
ORDER BY offered_at
LIMIT ?
`

type ListMenuWithDishesByOfferedAtParams struct {
	StartOfferedAt time.Time `json:"start_offered_at"`
	EndOfferedAt   time.Time `json:"end_offered_at"`
	CityCode       int32     `json:"city_code"`
	Limit          int32     `json:"limit"`
}

type ListMenuWithDishesByOfferedAtRow struct {
	ID                       string          `json:"id"`
	OfferedAt                time.Time       `json:"offered_at"`
	PhotoUrl                 sql.NullString  `json:"photo_url"`
	CreatedAt                time.Time       `json:"created_at"`
	ElementarySchoolCalories int32           `json:"elementary_school_calories"`
	JuniorHighSchoolCalories int32           `json:"junior_high_school_calories"`
	CityCode                 int32           `json:"city_code"`
	Dishes                   json.RawMessage `json:"dishes"`
}

func (q *Queries) ListMenuWithDishesByOfferedAt(ctx context.Context, arg ListMenuWithDishesByOfferedAtParams) ([]ListMenuWithDishesByOfferedAtRow, error) {
	rows, err := q.db.QueryContext(ctx, listMenuWithDishesByOfferedAt,
		arg.StartOfferedAt,
		arg.EndOfferedAt,
		arg.CityCode,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListMenuWithDishesByOfferedAtRow{}
	for rows.Next() {
		var i ListMenuWithDishesByOfferedAtRow
		if err := rows.Scan(
			&i.ID,
			&i.OfferedAt,
			&i.PhotoUrl,
			&i.CreatedAt,
			&i.ElementarySchoolCalories,
			&i.JuniorHighSchoolCalories,
			&i.CityCode,
			&i.Dishes,
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

const listMenus = `-- name: ListMenus :many
SELECT id, offered_at, photo_url, created_at, elementary_school_calories, junior_high_school_calories, city_code
FROM menus AS m
WHERE city_code = ?
ORDER BY offered_at
LIMIT ? OFFSET ?
`

type ListMenusParams struct {
	CityCode int32 `json:"city_code"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

func (q *Queries) ListMenus(ctx context.Context, arg ListMenusParams) ([]Menu, error) {
	rows, err := q.db.QueryContext(ctx, listMenus, arg.CityCode, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Menu{}
	for rows.Next() {
		var i Menu
		if err := rows.Scan(
			&i.ID,
			&i.OfferedAt,
			&i.PhotoUrl,
			&i.CreatedAt,
			&i.ElementarySchoolCalories,
			&i.JuniorHighSchoolCalories,
			&i.CityCode,
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

const listMenusByOfferedAt = `-- name: ListMenusByOfferedAt :many
SELECT id, offered_at, photo_url, created_at, elementary_school_calories, junior_high_school_calories, city_code
FROM menus
WHERE offered_at >= ?
  AND offered_at < ?
  AND city_code = ?
ORDER BY offered_at
LIMIT ?
`

type ListMenusByOfferedAtParams struct {
	StartOfferedAt time.Time `json:"start_offered_at"`
	EndOfferedAt   time.Time `json:"end_offered_at"`
	CityCode       int32     `json:"city_code"`
	Limit          int32     `json:"limit"`
}

func (q *Queries) ListMenusByOfferedAt(ctx context.Context, arg ListMenusByOfferedAtParams) ([]Menu, error) {
	rows, err := q.db.QueryContext(ctx, listMenusByOfferedAt,
		arg.StartOfferedAt,
		arg.EndOfferedAt,
		arg.CityCode,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Menu{}
	for rows.Next() {
		var i Menu
		if err := rows.Scan(
			&i.ID,
			&i.OfferedAt,
			&i.PhotoUrl,
			&i.CreatedAt,
			&i.ElementarySchoolCalories,
			&i.JuniorHighSchoolCalories,
			&i.CityCode,
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
