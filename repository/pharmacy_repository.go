package repository

import (
	"context"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type PharmacyRepository interface {
	GetPharmacy(ctx context.Context) ([]entity.Pharmacy, error)
	FindByPharmacyId(ctx context.Context, pharmacyId int64) (entity.Pharmacy, error)
}
