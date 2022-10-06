package repository

import (
	"context"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type RecipeRepository interface {
	InsertRecipe(ctx context.Context, medicalRecord int64, recipe entity.Recipe) (int64, error)
	FindByRecipeId(ctx context.Context, recipeId int64) (entity.Recipe, error)
	UpdateRecipe(ctx context.Context, recipeId int64, recipe entity.Recipe) error
}
