package repository

import (
	"context"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type PatientRepository interface {
	Insert(ctx context.Context, patient entity.Patient) (entity.Patient, error)
	FindByMedicalRecordNumber(ctx context.Context, medicalRecordNumber int32) (entity.Patient, error)
}
