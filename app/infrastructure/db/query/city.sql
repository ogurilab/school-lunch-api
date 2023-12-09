-- name: CreateCity :exec
INSERT INTO cities (
    city_code,
    city_name,
    prefecture_code,
    prefecture_name
  )
VALUES (
    sqlc.arg(city_code),
    sqlc.arg(city_name),
    sqlc.arg(prefecture_code),
    sqlc.arg(prefecture_name)
  );

-- name: GetCity :one
SELECT *
FROM cities
WHERE city_code = sqlc.arg(city_code)
LIMIT 1;

-- name: ListCities :many
SELECT *
FROM cities
ORDER BY city_code
LIMIT ? OFFSET ?;

-- name: ListCitiesByName :many
SELECT *
FROM cities
WHERE city_name LIKE ?
ORDER BY city_code
LIMIT ? OFFSET ?;

-- name: ListCitiesByPrefecture :many
SELECT *
FROM cities
WHERE prefecture_code = ?
ORDER BY city_code
LIMIT ? OFFSET ?;