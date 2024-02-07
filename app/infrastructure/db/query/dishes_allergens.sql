-- name: CreateDishesAllergens :exec
INSERT INTO dishes_allergens (dish_id, allergen_id, category)
VALUES (
    sqlc.arg(dish_id),
    sqlc.arg(allergen_id),
    sqlc.arg(category)
  );