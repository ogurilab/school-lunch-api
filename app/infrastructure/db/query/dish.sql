-- name: CreateDish :exec
INSERT INTO dishes (id, name)
VALUES (sqlc.arg(id), sqlc.arg(name));

-- name: GetDish :one
SELECT dishes.id,
  dishes.name
FROM dishes
WHERE id = sqlc.arg(id)
LIMIT 1;

-- name: ListDishByMenuID :many
SELECT dishes.id,
  dishes.name
FROM dishes
WHERE id IN (
    SELECT dish_id
    FROM menu_dishes
    WHERE menu_id = sqlc.arg(menu_id)
  )
ORDER BY id;

-- name: ListDishByName :many
SELECT dishes.id,
  dishes.name
FROM dishes
WHERE name LIKE ?
ORDER BY id
LIMIT ? OFFSET ?;

-- name: ListDish :many
SELECT dishes.id,
  dishes.name
FROM dishes
ORDER BY id
LIMIT ? OFFSET ?;