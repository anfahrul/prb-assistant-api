package repository

import (
	"context"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type PharmacyRepository interface {
	FindByPharmacyId(ctx context.Context, pharmacyId int64) (entity.Pharmacy, error)
}
