package repository

import (
	"context"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type BookRepository interface {
	InsertBook(ctx context.Context, medicalRecordNumber int64, time int) (int64, error)
	FindBookByMedicalRecordNumber(ctx context.Context, medicalRecordNumber int64) ([]entity.Book, error)
	FindBookById(ctx context.Context, bookId int32, medicalRecordNumber int64) error
	UpdateBook(ctx context.Context, book entity.Book, medicalRecordNumber int64, bookid int32) (int32, error)
}
