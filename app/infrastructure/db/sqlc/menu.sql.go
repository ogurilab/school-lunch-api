// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: menu.sql

package db

import (
	"context"
	"database/sql"
	"strings"
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

const listMenuByCity = `-- name: ListMenuByCity :many
SELECT id, offered_at, photo_url, created_at, elementary_school_calories, junior_high_school_calories, city_code
FROM menus AS m
WHERE city_code = ?
  AND offered_at <= ?
ORDER BY offered_at DESC
LIMIT ? OFFSET ?
`

type ListMenuByCityParams struct {
	CityCode  int32     `json:"city_code"`
	OfferedAt time.Time `json:"offered_at"`
	Limit     int32     `json:"limit"`
	Offset    int32     `json:"offset"`
}

func (q *Queries) ListMenuByCity(ctx context.Context, arg ListMenuByCityParams) ([]Menu, error) {
	rows, err := q.db.QueryContext(ctx, listMenuByCity,
		arg.CityCode,
		arg.OfferedAt,
		arg.Limit,
		arg.Offset,
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

const listMenuInIds = `-- name: ListMenuInIds :many
SELECT id, offered_at, photo_url, created_at, elementary_school_calories, junior_high_school_calories, city_code
FROM menus
WHERE id IN (/*SLICE:ids*/?)
  AND offered_at <= ?
ORDER BY offered_at DESC
LIMIT ? OFFSET ?
`

type ListMenuInIdsParams struct {
	Ids       []string  `json:"ids"`
	OfferedAt time.Time `json:"offered_at"`
	Limit     int32     `json:"limit"`
	Offset    int32     `json:"offset"`
}

func (q *Queries) ListMenuInIds(ctx context.Context, arg ListMenuInIdsParams) ([]Menu, error) {
	query := listMenuInIds
	var queryParams []interface{}
	if len(arg.Ids) > 0 {
		for _, v := range arg.Ids {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:ids*/?", strings.Repeat(",?", len(arg.Ids))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:ids*/?", "NULL", 1)
	}
	queryParams = append(queryParams, arg.OfferedAt)
	queryParams = append(queryParams, arg.Limit)
	queryParams = append(queryParams, arg.Offset)
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
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
