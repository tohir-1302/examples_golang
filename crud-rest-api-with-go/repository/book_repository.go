package repository

import (
	"context"
	"rest-api-golang/model"
)

type BookRepository interface {
	Save(ctx context.Context, book model.Book)
	Update(ctx context.Context, book model.Book)
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) (model.Book, error)
	FindAll(ctx context.Context) ([]model.Book, error)
}
