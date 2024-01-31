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

-- name: GetMenuWithDishes :many
SELECT m.*,
  d.id AS dish_id,
  d.name AS dish_name
FROM (
    SELECT *
    FROM menus
    WHERE menus.id = sqlc.arg(id)
      AND city_code = sqlc.arg(city_code)
  ) AS m
  INNER JOIN menu_dishes AS md ON m.id = md.menu_id
  INNER JOIN dishes AS d ON md.dish_id = d.id
ORDER BY d.id ASC;

-- name: ListMenuWithDishesByCity :many
SELECT m.*,
  d.id AS dish_id,
  d.name AS dish_name
FROM (
    SELECT *
    FROM menus AS m
    WHERE city_code = sqlc.arg(city_code)
      AND offered_at <= sqlc.arg(offered_at)
    ORDER BY offered_at DESC
    LIMIT ? OFFSET ?
  ) AS m
  INNER JOIN menu_dishes md ON m.id = md.menu_id
  INNER JOIN dishes d ON md.dish_id = d.id;

-- name: ListMenuWithDishes :many
SELECT m.*,
  d.id AS dish_id,
  d.name AS dish_name
FROM (
    SELECT *
    FROM menus AS m
    WHERE offered_at <= sqlc.arg(offered_at)
    ORDER BY offered_at DESC
    LIMIT ? OFFSET ?
  ) AS m
  INNER JOIN menu_dishes AS md ON m.id = md.menu_id
  INNER JOIN dishes AS d ON md.dish_id = d.id;