package repository

import (
	"context"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type MedicineRepository interface {
	FindByRecipeId(ctx context.Context, recipeId int64) ([]entity.Medicine, error)
}
