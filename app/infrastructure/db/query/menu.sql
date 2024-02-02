-- name: CreateMenu :exec
INSERT INTO menus (
    id,
    offered_at,
    photo_url,
    elementary_school_calories,
    junior_high_school_calories,
    city_code
  )
VALUES (
    sqlc.arg(id),
    sqlc.arg(offered_at),
    sqlc.arg(photo_url),
    sqlc.arg(elementary_school_calories),
    sqlc.arg(junior_high_school_calories),
    sqlc.arg(city_code)
  );

-- name: GetMenu :one
SELECT *
FROM menus
WHERE id = sqlc.arg(id)
  AND city_code = sqlc.arg(city_code);

-- name: ListMenuByCity :many
SELECT *
FROM menus AS m
WHERE city_code = sqlc.arg(city_code)
  AND offered_at <= sqlc.arg(offered_at)
ORDER BY offered_at DESC
LIMIT ? OFFSET ?;