-- name: CreateAllergen :exec
INSERT INTO allergens (name)
VALUES (sqlc.arg(name));

-- name: GetAllergenByName :one
SELECT id,
  name
FROM allergens
WHERE name = sqlc.arg(name);

-- name: ListAllergenByDishID :many
SELECT allergens.id,
  allergens.name
FROM allergens
  JOIN dishes_allergens ON allergens.id = dishes_allergens.allergen_id
WHERE dishes_allergens.dish_id = sqlc.arg(dish_id)
ORDER BY allergens.name;

-- name: ListAllergenInDish :many
SELECT allergens.id,
  allergens.name
FROM allergens
  JOIN dishes_allergens ON allergens.id = dishes_allergens.allergen_id
WHERE dishes_allergens.dish_id IN (sqlc.slice(dish_ids))
ORDER BY allergens.name;