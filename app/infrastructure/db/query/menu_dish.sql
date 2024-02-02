-- name: CreateMenuDish :exec
INSERT INTO menu_dishes (menu_id, dish_id)
VALUES (sqlc.arg("menu_id"), sqlc.arg("dish_id"));