package repository

import (
	"context"

	"github.com/anfahrul/prb-assistant-api/entity"
)

type BookRepository interface {
	InsertBook(ctx context.Context, medicalRecordNumber int32) (int32, error)
	FindBookByMedicalRecordNumber(ctx context.Context, medicalRecordNumber int32) ([]entity.Book, error)
	UpdateBook(ctx context.Context, book entity.Book, medicalRecordNumber int32, bookid int32) (int32, error)
}
