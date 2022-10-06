package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type patientRepositoryImpl struct {
	DB *sql.DB
}

func NewPatientRepository(db *sql.DB) PatientRepository {
	return &patientRepositoryImpl{DB: db}
}

func (repository *patientRepositoryImpl) Insert(ctx context.Context, patien entity.Patient) (entity.Patient, error) {
	script := "INSERT INTO patient(bpjsNumber, name, hospital, diagnosis, bithdate, address, phoneNumber) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, patien.BpjsNumber, patien.Name, patien.Hospital, patien.Diagnosis, patien.Bithdate, patien.Address, patien.PhoneNumber)
	if err != nil {
		return patien, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return patien, err
	}
	patien.MedicalRecord = int32(id)
	return patien, nil
}

func (repository *patientRepositoryImpl) FindByMedicalRecordNumber(ctx context.Context, medicalRecordNumber int32) (entity.Patient, error) {
	script := "SELECT medicalRecord, bpjsNumber, name, hospital, diagnosis, bithdate, address, phoneNumber FROM patient WHERE medicalRecord = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, medicalRecordNumber)
	patient := entity.Patient{}
	if err != nil {
		return patient, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(
			&patient.MedicalRecord,
			&patient.BpjsNumber,
			&patient.Name,
			&patient.Hospital,
			&patient.Diagnosis,
			&patient.Bithdate,
			&patient.Address,
			&patient.PhoneNumber,
		)
		return patient, nil
	} else {
		// tidak ada
		return patient, errors.New("medicalRecord " + strconv.Itoa(int(medicalRecordNumber)) + " Not Found")
	}
}
