package repository

import (
	"context"
	"database/sql"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type bookRepositoryImpl struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepositoryImpl{DB: db}
}

func (repository *bookRepositoryImpl) InsertBook(ctx context.Context, medicalRecordNumber int32) (int32, error) {
	script := "INSERT INTO book(medicalRecord) VALUES (?)"
	result, err := repository.DB.ExecContext(ctx, script, medicalRecordNumber)
	if err != nil {
		return medicalRecordNumber, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return medicalRecordNumber, err
	}

	return int32(id), nil
}

func (repository *bookRepositoryImpl) FindBookByMedicalRecordNumber(ctx context.Context, medicalRecordNumber int32) ([]entity.Book, error) {
	var books []entity.Book

	script := "SELECT bookId, medicalRecord, checkDate, doctorName, medicalStatus, note FROM book WHERE medicalRecord = ?"
	rows, err := repository.DB.QueryContext(ctx, script, medicalRecordNumber)
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		// ada
		book := entity.Book{}
		rows.Scan(
			&book.BookId,
			&book.MedicalRecord,
			&book.CheckDate,
			&book.DoctorName,
			&book.MedicalStatus,
			&book.Note,
		)
		books = append(books, book)
	}

	return books, nil
}

func (repository *bookRepositoryImpl) UpdateBook(ctx context.Context, book entity.Book, medicalRecordNumber int32, bookid int32) (int32, error) {
	script := `
		UPDATE book
		SET checkDate = ?, doctorName = ?, medicalStatus = ?, note = ?
		WHERE medicalRecord=? AND bookId=?
	`

	result, err := repository.DB.ExecContext(ctx, script, book.CheckDate, book.DoctorName, book.MedicalStatus, book.Note, medicalRecordNumber, bookid)
	if err != nil {
		return bookid, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return bookid, err
	}

	return int32(id), nil
}
