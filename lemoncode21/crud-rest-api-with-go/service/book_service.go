package service

import (
	"context"
	"rest-api-golang/data/request"
	"rest-api-golang/data/response"
)

type BookService interface {
	Create(ctx context.Context, request request.BookCreateRequest)
	Update(ctx context.Context, request request.BookUpdateRequest)
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) response.BookResponse
	FindAll(ctx context.Context) []response.BookResponse
}
