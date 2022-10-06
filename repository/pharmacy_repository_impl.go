package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type pharmacyRepositoryImpl struct {
	DB *sql.DB
}

func NewPharmacyRepository(db *sql.DB) PharmacyRepository {
	return &pharmacyRepositoryImpl{DB: db}
}

func (repository *pharmacyRepositoryImpl) FindByPharmacyId(ctx context.Context, pharmacyId int64) (entity.Pharmacy, error) {
	script := "SELECT pharmacyId, name, address FROM pharmacy WHERE pharmacyId = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, pharmacyId)

	pharmacy := entity.Pharmacy{}
	if err != nil {
		return entity.Pharmacy{}, err
	}

	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(
			&pharmacy.PharmacyId,
			&pharmacy.Name,
			&pharmacy.Address,
		)

		return pharmacy, nil
	} else {
		// tidak ada
		return pharmacy, errors.New("Pharmacy " + strconv.Itoa(int(pharmacyId)) + " Not Found")
	}
}

func (repository *pharmacyRepositoryImpl) GetPharmacy(ctx context.Context) ([]entity.Pharmacy, error) {
	var pharmacies []entity.Pharmacy

	script := "SELECT pharmacyId, name, address FROM pharmacy"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return pharmacies, err
	}
	defer rows.Close()

	for rows.Next() {
		// ada
		pharmacy := entity.Pharmacy{}
		rows.Scan(
			&pharmacy.PharmacyId,
			&pharmacy.Name,
			&pharmacy.Address,
		)
		pharmacies = append(pharmacies, pharmacy)
	}

	return pharmacies, nil
}
