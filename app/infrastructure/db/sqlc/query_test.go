package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/ogurilab/school-lunch-api/domain"
	"github.com/ogurilab/school-lunch-api/util"
	"github.com/stretchr/testify/require"
)

type createDishResult struct {
	dish     *domain.Dish
	menuDish MenuDish
}

type createDishesResult struct {
	dishes     []*domain.Dish
	menuDishes []MenuDish
}

func TestCreateDishTx(t *testing.T) {
	cityCode := util.RandomCityCode()
	menu := createRandomMenu(t, cityCode)

	// transactionなので、goroutineを使って並列処理を行う
	n := 10
	errs := make(chan error)
	results := make(chan createDishResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			dish, err := domain.NewDish(util.RandomString(10))
			require.NoError(t, err)

			err = testQuery.CreateDishTx(ctx, dish, menu.ID)
			errs <- err
			results <- createDishResult{
				dish:     dish,
				menuDish: MenuDish{MenuID: menu.ID, DishID: dish.ID},
			}
		}()
	}

	for i := 0; i < n; i++ {
		ctx := context.Background()
		err := <-errs
		require.NoError(t, err)

		result := <-results

		requireCreatedDishAndMenuDish(t, ctx, menu, result)
	}
}

func TestCreateDishTxRollback(t *testing.T) {
	testQuery.truncateMenuDishesTable()
	cityCode := util.RandomCityCode()
	menu := createRandomMenu(t, cityCode)

	// transactionなので、goroutineを使って並列処理を行う
	n := 5
	errs := make(chan error)
	results := make(chan createDishResult)

	duplicateDish, err := domain.NewDish(util.RandomString(10))

	require.NoError(t, err)

	for i := 0; i < n; i++ {
		go func() {

			ctx := context.Background()

			err = testQuery.CreateDishTx(ctx, duplicateDish, menu.ID)

			errs <- err
			results <- createDishResult{
				dish:     duplicateDish,
				menuDish: MenuDish{MenuID: menu.ID, DishID: duplicateDish.ID},
			}
		}()
	}

	for i := 0; i < n; i++ {
		ctx := context.Background()
		err := <-errs

		if i == 0 {
			// 1回目の処理は成功する
			require.NoError(t, err)

			result := <-results

			requireCreatedDishAndMenuDish(t, ctx, menu, result)
		} else {
			// 2回目以降は失敗する
			require.Error(t, err)
			expectError := fmt.Sprintf("Error 1062 (23000): Duplicate entry '%s' for key 'dishes.PRIMARY'", duplicateDish.ID)

			require.EqualError(t, err, expectError)

			// MenuDishが保存されていないことを確認する
			count, err := testQuery.getMenuDishesAllCount(menu.ID)

			require.NoError(t, err)
			require.Equal(t, 1, count)
		}
	}
}

func TestCreateDishesTx(t *testing.T) {
	cityCode := util.RandomCityCode()
	menu := createRandomMenu(t, cityCode)

	// transactionなので、goroutineを使って並列処理を行う
	n := 5

	errs := make(chan error)
	results := make(chan createDishesResult)

	for i := 0; i < n; i++ {
		go func() {

			var dishes []*domain.Dish
			for i := 0; i < n; i++ {
				dish, err := domain.NewDish(util.RandomString(10))
				require.NoError(t, err)

				dishes = append(dishes, dish)
			}

			ctx := context.Background()

			err := testQuery.CreateDishesTx(ctx, dishes, menu.ID)

			errs <- err

			menuDishes := make([]MenuDish, 0, n)

			for _, dish := range dishes {
				menuDishes = append(menuDishes, MenuDish{MenuID: menu.ID, DishID: dish.ID})
			}

			results <- createDishesResult{
				dishes:     dishes,
				menuDishes: menuDishes,
			}
		}()
	}

	for i := 0; i < n; i++ {
		ctx := context.Background()
		err := <-errs
		require.NoError(t, err)

		result := <-results

		requireCreatedDishesAndMenuDishes(t, ctx, menu, result)
	}
}

func TestCreateDishesTxRollback(t *testing.T) {
	testQuery.truncateMenuDishesTable()
	cityCode := util.RandomCityCode()
	menu := createRandomMenu(t, cityCode)

	// transactionなので、goroutineを使って並列処理を行う
	n := 5
	errs := make(chan error)
	results := make(chan createDishesResult)

	duplicateDish, err := domain.NewDish(util.RandomString(10))

	require.NoError(t, err)

	for i := 0; i < n; i++ {
		go func() {

			var dishes []*domain.Dish
			for i := 0; i < n; i++ {
				dish, err := domain.NewDish(util.RandomString(10))
				require.NoError(t, err)

				dishes = append(dishes, dish)
			}

			dishes = append(dishes, duplicateDish)

			ctx := context.Background()

			err = testQuery.CreateDishesTx(ctx, dishes, menu.ID)

			errs <- err

			menuDishes := make([]MenuDish, 0, n)

			for _, dish := range dishes {
				menuDishes = append(menuDishes, MenuDish{MenuID: menu.ID, DishID: dish.ID})
			}

			results <- createDishesResult{
				dishes:     dishes,
				menuDishes: menuDishes,
			}
		}()
	}

	for i := 0; i < n; i++ {
		ctx := context.Background()
		err := <-errs

		if i == 0 {
			// 1回目の処理は成功する
			require.NoError(t, err)

			result := <-results

			requireCreatedDishesAndMenuDishes(t, ctx, menu, result)
		} else {
			// 2回目以降は失敗する
			require.Error(t, err)
			expectError := fmt.Sprintf("Error 1062 (23000): Duplicate entry '%s' for key 'dishes.PRIMARY'", duplicateDish.ID)

			require.EqualError(t, err, expectError)

			// MenuDishが保存されていないことを確認する
			count, err := testQuery.getMenuDishesAllCount(menu.ID)

			require.NoError(t, err)
			require.Equal(t, n+1, count)
		}
	}
}

func requireCreatedDishAndMenuDish(t *testing.T, ctx context.Context, menu *domain.Menu, result createDishResult) {
	// 正常に処理が完了しているか確認する
	require.NotEmpty(t, result.dish.ID)
	require.NotEmpty(t, result.menuDish.MenuID)
	require.NotEmpty(t, result.menuDish.DishID)
	require.Equal(t, result.dish.ID, result.menuDish.DishID)
	require.Equal(t, result.menuDish.MenuID, menu.ID)

	// MenuDishを取得して、データが正しく保存されているか確認する
	menuDish, err := testQuery.getMenuDishes(result.menuDish.MenuID, result.menuDish.DishID)

	require.NoError(t, err)

	require.NotEmpty(t, menuDish.MenuID)
	require.NotEmpty(t, menuDish.DishID)
	require.Equal(t, menuDish.MenuID, result.menuDish.MenuID)
	require.Equal(t, menuDish.DishID, result.menuDish.DishID)

	// Dishを取得して、データが正しく保存されているか確認する
	dish, err := testQuery.GetDish(ctx, result.dish.ID)

	require.NoError(t, err)

	require.NotEmpty(t, dish.ID)
	require.NotEmpty(t, dish.Name)
	require.Equal(t, dish.ID, result.dish.ID)
	require.Equal(t, dish.Name, result.dish.Name)
}

func requireCreatedDishesAndMenuDishes(t *testing.T, ctx context.Context, menu *domain.Menu, result createDishesResult) {
	// 正常に処理が完了しているか確認する
	for i, dish := range result.dishes {
		require.NotEmpty(t, dish.ID)
		require.NotEmpty(t, result.menuDishes[i].MenuID)
		require.NotEmpty(t, result.menuDishes[i].DishID)
		require.Equal(t, dish.ID, result.menuDishes[i].DishID)
		require.Equal(t, result.menuDishes[i].MenuID, menu.ID)
	}

	// MenuDishを取得して、データが正しく保存されているか確認する
	for _, menuDish := range result.menuDishes {
		menuDishRes, err := testQuery.getMenuDishes(menuDish.MenuID, menuDish.DishID)

		require.NoError(t, err)
		require.NotEmpty(t, menuDish.MenuID)
		require.NotEmpty(t, menuDish.DishID)
		require.Equal(t, menuDish.MenuID, menuDishRes.MenuID)
		require.Equal(t, menuDish.DishID, menuDishRes.DishID)
	}

	// Dishを取得して、データが正しく保存されているか確認する
	for _, dish := range result.dishes {
		res, err := testQuery.GetDish(ctx, dish.ID)

		require.NoError(t, err)

		require.NotEmpty(t, dish.ID)
		require.NotEmpty(t, dish.Name)
		require.Equal(t, dish.ID, res.ID)
		require.Equal(t, dish.Name, res.Name)
	}
}
