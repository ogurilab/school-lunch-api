-- name: CreateDish :exec
INSERT INTO dishes (id, menu_id, name)
VALUES (
    sqlc.arg(id),
    sqlc.arg(menu_id),
    sqlc.arg(name)
  );

-- name: GetDish :one
SELECT *
FROM dishes
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListDishByMenuID :many
SELECT *
FROM dishes
WHERE menu_id = sqlc.arg(menu_id)
ORDER BY id;

-- name: ListDishByName :many
SELECT *
FROM dishes
WHERE name LIKE ?
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListDish :many
SELECT *
FROM dishes
ORDER BY id
LIMIT ? OFFSET ?;