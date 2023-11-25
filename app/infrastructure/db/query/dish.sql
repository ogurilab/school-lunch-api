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

-- name: ListDishes :many
SELECT *
FROM dishes
WHERE menu_id = sqlc.arg(menu_id)
ORDER BY id
LIMIT ? OFFSET ?;

-- name: GetDishByNames :many
SELECT *
FROM dishes
WHERE name IN (sqlc.slice(names))
ORDER BY id
LIMIT ? OFFSET ?;