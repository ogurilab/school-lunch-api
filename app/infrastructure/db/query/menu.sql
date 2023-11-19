-- name: CreateMenu :exec
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
    sqlc.arg(id),
    sqlc.arg(offered_at),
    sqlc.arg(region_id),
    sqlc.arg(photo_url),
    sqlc.arg(wikimedia_commons_url),
    sqlc.arg(elementary_school_calories),
    sqlc.arg(junior_high_school_calories)
  );

-- name: GetMenu :one
SELECT *
FROM menus
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListMenus :many
SELECT *
FROM menus
ORDER BY offered_at
LIMIT ? OFFSET ?;