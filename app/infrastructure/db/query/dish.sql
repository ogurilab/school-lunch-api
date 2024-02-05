-- name: CreateDish :exec
INSERT INTO dishes (id, name)
VALUES (sqlc.arg(id), sqlc.arg(name));

-- name: GetDish :many
SELECT dishes.id,
  dishes.name,
  md.menu_id AS menu_id
FROM dishes
  INNER JOIN menu_dishes AS md ON dishes.id = md.dish_id
WHERE dishes.id = sqlc.arg(id)
ORDER BY dishes.id
LIMIT ? OFFSET ?;

-- name: GetDishInCity :many
SELECT dishes.id,
  dishes.name,
  md.menu_id AS menu_id
FROM dishes
  INNER JOIN menu_dishes AS md ON dishes.id = md.dish_id
  INNER JOIN menus AS m ON md.menu_id = m.id
WHERE dishes.id = sqlc.arg(id)
  AND m.city_code = sqlc.arg(city_code)
ORDER BY dishes.id
LIMIT ? OFFSET ?;

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