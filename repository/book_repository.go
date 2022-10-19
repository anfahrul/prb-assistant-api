package repository

import (
	"context"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type BookRepository interface {
	InsertBook(ctx context.Context, medicalRecordNumber int32, time int) (int32, error)
	FindBookByMedicalRecordNumber(ctx context.Context, medicalRecordNumber int32) ([]entity.Book, error)
	FindBookById(ctx context.Context, bookId int32, medicalRecordNumber int32) error
	UpdateBook(ctx context.Context, book entity.Book, medicalRecordNumber int32, bookid int32) (int32, error)
}
