package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type recipeRepositoryImpl struct {
	DB *sql.DB
}

func NewRecipeRepository(db *sql.DB) RecipeRepository {
	return &recipeRepositoryImpl{DB: db}
}

func (repository *recipeRepositoryImpl) InsertRecipe(ctx context.Context, medicalRecord int64, recipe entity.Recipe) (int64, error) {
	script := "INSERT INTO recipe(medicalRecord) VALUES (?)"
	result, err := repository.DB.ExecContext(ctx, script, medicalRecord)
	if err != nil {
		return medicalRecord, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return medicalRecord, err
	}

	return int64(id), nil
}

func (repository *recipeRepositoryImpl) FindByRecipeId(ctx context.Context, recipeId int64) (entity.Recipe, error) {
	script := "SELECT recipeId, medicalRecord, claimStatus, pharmacyId FROM recipe WHERE recipeId = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, recipeId)

	recipe := entity.Recipe{}
	if err != nil {
		return recipe, err
	}

	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(
			&recipe.RecipeId,
			&recipe.MedicalRecord,
			&recipe.ClaimStatus,
			&recipe.PharmacyId,
		)

		return recipe, nil
	} else {
		// tidak ada
		return recipe, errors.New("Recipe " + strconv.Itoa(int(recipeId)) + " Not Found")
	}
}

func (repository *recipeRepositoryImpl) UpdateRecipe(ctx context.Context, recipeId int64, recipe entity.Recipe) error {
	script := `
		UPDATE recipe
		SET claimStatus = ?, pharmacyId = ?
		WHERE recipeId=?
	`

	_, err := repository.DB.ExecContext(ctx, script, recipe.ClaimStatus, recipe.PharmacyId, recipeId)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	return nil
}
