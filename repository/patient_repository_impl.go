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

func (repository *patientRepositoryImpl) GetAllPatient(ctx context.Context) ([]entity.Patient, error) {
	script := `SELECT * FROM patient`
	rows, err := repository.DB.QueryContext(ctx, script)
	patients := []entity.Patient{}
	if err != nil {
		return patients, err
	}
	defer rows.Close()

	for rows.Next() {
		// ada
		patient := entity.Patient{}
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
		patients = append(patients, patient)
	}

	return patients, nil
}

func (repository *patientRepositoryImpl) Insert(ctx context.Context, patien entity.Patient) (entity.Patient, error) {
	script := "INSERT INTO patient(medicalRecord, bpjsNumber, name, hospital, diagnosis, bithdate, address, phoneNumber) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := repository.DB.ExecContext(ctx, script, patien.MedicalRecord, patien.BpjsNumber, patien.Name, patien.Hospital, patien.Diagnosis, patien.Bithdate, patien.Address, patien.PhoneNumber)
	if err != nil {
		return patien, err
	}

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
