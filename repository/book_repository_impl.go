package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type bookRepositoryImpl struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepositoryImpl{DB: db}
}

func (repository *bookRepositoryImpl) InsertBook(ctx context.Context, medicalRecordNumber int64, time int) (int64, error) {
	script := "INSERT INTO book(medicalRecord, checkDate) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, medicalRecordNumber, time)
	if err != nil {
		return medicalRecordNumber, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return medicalRecordNumber, err
	}

	return id, nil
}

func (repository *bookRepositoryImpl) FindBookByMedicalRecordNumber(ctx context.Context, medicalRecordNumber int64) ([]entity.Book, error) {
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

	if len(books) == 0 {
		return books, errors.New("No Books belongs medical record number " + strconv.Itoa(int(medicalRecordNumber)))
	}

	return books, nil
}

func (repository *bookRepositoryImpl) FindBookById(ctx context.Context, bookId int32, medicalRecordNumber int64) error {
	var book entity.Book

	script := "SELECT bookId FROM book WHERE bookId = ? AND medicalRecord = ?"
	rows, err := repository.DB.QueryContext(ctx, script, bookId, medicalRecordNumber)
	if err != nil {
		return errors.New("Book Id " + strconv.Itoa(int(bookId)) + "Not found")
	}
	defer rows.Close()

	if rows.Next() {
		// ada
		rows.Scan(
			&book.BookId,
		)
	} else {
		return errors.New("Book Id " + strconv.Itoa(int(bookId)) + " on medical record number " + strconv.Itoa(int(medicalRecordNumber)) + " Not found")
	}

	return nil
}

func (repository *bookRepositoryImpl) UpdateBook(ctx context.Context, book entity.Book, medicalRecordNumber int64, bookid int32) (int32, error) {
	script := `
		UPDATE book
		SET doctorName = ?, medicalStatus = ?, note = ?
		WHERE medicalRecord=? AND bookId=?
	`

	result, err := repository.DB.ExecContext(ctx, script, book.DoctorName, book.MedicalStatus, book.Note, medicalRecordNumber, bookid)
	if err != nil {
		return bookid, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return bookid, err
	}

	return int32(id), nil
}
