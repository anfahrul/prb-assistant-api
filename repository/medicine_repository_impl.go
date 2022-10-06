package repository

import (
	"context"
	"database/sql"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type medicineRepositoryImpl struct {
	DB *sql.DB
}

func NewMedicineRepository(db *sql.DB) MedicineRepository {
	return &medicineRepositoryImpl{DB: db}
}

func (repository *medicineRepositoryImpl) FindByRecipeId(ctx context.Context, recipeId int64) ([]entity.Medicine, error) {
	script := `SELECT * FROM medicine WHERE recipeId = ?`
	rows, err := repository.DB.QueryContext(ctx, script, recipeId)
	medicine := []entity.Medicine{}
	if err != nil {
		return medicine, err
	}
	defer rows.Close()

	for rows.Next() {
		// ada
		drug := entity.Medicine{}
		rows.Scan(
			&drug.MedicineId,
			&drug.RecipeId,
			&drug.Name,
			&drug.Quantity,
			&drug.Portion,
		)
		medicine = append(medicine, drug)
	}

	return medicine, nil
}
