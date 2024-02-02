package db

import (
	"context"

	"github.com/ogurilab/school-lunch-api/domain"
)

type bulkInsertDishQuery struct {
	query string
	args  []any
}

type bulkInsertMenuDishQuery struct {
	query string
	args  []any
}

func (q *SQLQuery) CreateDishesTx(ctx context.Context, dishes []*domain.Dish, menuID string) error {

	err := q.execTx(ctx, func(q *Queries) error {

		dishSQL := createBulkInsertDishQuery(dishes)

		_, err := q.db.ExecContext(ctx, dishSQL.query, dishSQL.args...)

		if err != nil {
			return err
		}

		menuDishSQL := createBulkInsertMenuDishQuery(dishes, menuID)

		_, err = q.db.ExecContext(ctx, menuDishSQL.query, menuDishSQL.args...)

		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func createBulkInsertDishQuery(dishes []*domain.Dish) bulkInsertDishQuery {

	insert := `INSERT INTO dishes (id, name) VALUES `

	values := make([]any, 0, len(dishes))

	for _, dish := range dishes {
		values = append(values, dish.ID, dish.Name)

		insert += "(?, ?),"

	}

	insert = insert[:len(insert)-1]

	return bulkInsertDishQuery{
		query: insert,
		args:  values,
	}
}

func createBulkInsertMenuDishQuery(dishes []*domain.Dish, menuID string) bulkInsertMenuDishQuery {

	insert := `INSERT INTO menu_dishes (menu_id, dish_id) VALUES `

	values := make([]any, 0, len(dishes))

	for _, dish := range dishes {
		values = append(values, menuID, dish.ID)

		insert += "(?, ?),"
	}

	insert = insert[:len(insert)-1]

	return bulkInsertMenuDishQuery{
		query: insert,
		args:  values,
	}
}
