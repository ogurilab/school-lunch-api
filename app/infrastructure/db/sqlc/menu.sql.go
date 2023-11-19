// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: menu.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createMenu = `-- name: CreateMenu :exec
INSERT INTO menus (
    id,
    offered_at,
    region_id,
    photo_url,
    wikimedia_commons_url,
    elementary_school_calories,
    junior_high_school_calories
  )
VALUES (
    ?,
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
	RegionID                 int32          `json:"region_id"`
	PhotoUrl                 sql.NullString `json:"photo_url"`
	WikimediaCommonsUrl      sql.NullString `json:"wikimedia_commons_url"`
	ElementarySchoolCalories sql.NullInt32  `json:"elementary_school_calories"`
	JuniorHighSchoolCalories sql.NullInt32  `json:"junior_high_school_calories"`
}

func (q *Queries) CreateMenu(ctx context.Context, arg CreateMenuParams) error {
	_, err := q.db.ExecContext(ctx, createMenu,
		arg.ID,
		arg.OfferedAt,
		arg.RegionID,
		arg.PhotoUrl,
		arg.WikimediaCommonsUrl,
		arg.ElementarySchoolCalories,
		arg.JuniorHighSchoolCalories,
	)
	return err
}

const getMenu = `-- name: GetMenu :one
SELECT id, offered_at, region_id, photo_url, wikimedia_commons_url, created_at, elementary_school_calories, junior_high_school_calories
FROM menus
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetMenu(ctx context.Context, id string) (Menu, error) {
	row := q.db.QueryRowContext(ctx, getMenu, id)
	var i Menu
	err := row.Scan(
		&i.ID,
		&i.OfferedAt,
		&i.RegionID,
		&i.PhotoUrl,
		&i.WikimediaCommonsUrl,
		&i.CreatedAt,
		&i.ElementarySchoolCalories,
		&i.JuniorHighSchoolCalories,
	)
	return i, err
}

const listMenus = `-- name: ListMenus :many
SELECT id, offered_at, region_id, photo_url, wikimedia_commons_url, created_at, elementary_school_calories, junior_high_school_calories
FROM menus
ORDER BY offered_at
LIMIT ? OFFSET ?
`

type ListMenusParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListMenus(ctx context.Context, arg ListMenusParams) ([]Menu, error) {
	rows, err := q.db.QueryContext(ctx, listMenus, arg.Limit, arg.Offset)
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
			&i.RegionID,
			&i.PhotoUrl,
			&i.WikimediaCommonsUrl,
			&i.CreatedAt,
			&i.ElementarySchoolCalories,
			&i.JuniorHighSchoolCalories,
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