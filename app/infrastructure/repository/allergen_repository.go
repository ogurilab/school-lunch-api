package repository

import (
	"context"

	"github.com/ogurilab/school-lunch-api/domain"
	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
)

type allergenRepository struct {
	query db.Query
}

func NewAllergenRepository(query db.Query) domain.AllergenRepository {
	return &allergenRepository{
		query: query,
	}
}

func (r *allergenRepository) FetchByDishID(ctx context.Context, dishID string) ([]*domain.Allergen, error) {

	results, err := r.query.ListAllergenByDishID(ctx, dishID)

	if err != nil {
		return nil, err
	}

	allergens := make([]*domain.Allergen, 0, len(results))

	for _, result := range results {
		allergen := domain.ReNewAllergen(result.ID, result.Name)

		allergens = append(allergens, allergen)
	}

	return allergens, nil
}

func (r *allergenRepository) FetchInDish(ctx context.Context, dishIDs []string) ([]*domain.Allergen, error) {

	results, err := r.query.ListAllergenInDish(ctx, dishIDs)

	if err != nil {
		return nil, err
	}

	allergens := make([]*domain.Allergen, 0, len(results))

	for _, result := range results {
		allergen := domain.ReNewAllergen(result.ID, result.Name)

		allergens = append(allergens, allergen)
	}

	return allergens, nil
}
