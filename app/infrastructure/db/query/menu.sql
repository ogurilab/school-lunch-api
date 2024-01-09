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

-- name: ListMenuNyCity :many
SELECT *
FROM menus AS m
WHERE city_code = sqlc.arg(city_code)
  AND offered_at >= sqlc.arg(offered_at)
ORDER BY offered_at
LIMIT ? OFFSET ?;

-- name: GetMenuWithDishes :one
SELECT m.*,
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
WHERE m.id = sqlc.arg(id)
  AND m.city_code = sqlc.arg(city_code)
GROUP BY m.id;

-- name: ListMenuWithDishesByCity :many
SELECT m.*,
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
WHERE m.city_code = sqlc.arg(city_code)
  AND m.offered_at >= sqlc.arg(offered_at)
GROUP BY m.id
ORDER BY offered_at
LIMIT ? OFFSET ?;

-- name: ListMenuWithDishes :many
SELECT m.*,
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
WHERE m.offered_at >= sqlc.arg(offered_at)
GROUP BY m.id
ORDER BY offered_at
LIMIT ? OFFSET ?;