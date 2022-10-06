package repository

import (
	"context"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type MedicineRepository interface {
	FindByRecipeId(ctx context.Context, recipeId int64) ([]entity.Medicine, error)
	InsertMedicine(ctx context.Context, recipeId int32, medicine entity.Medicine) (int32, error)
}
