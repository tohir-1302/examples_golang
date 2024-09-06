package service

import (
	"context"
	"rest-api-golang/data/request"
	"rest-api-golang/data/response"
	"rest-api-golang/helper"
	"rest-api-golang/model"
	"rest-api-golang/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}

func (b *BookServiceImpl) Create(ctx context.Context, request request.BookCreateRequest) {
	book := model.Book{
		Name: request.Name,
	}

	b.BookRepository.Save(ctx, book)
}

func (b *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) {
	book := model.Book{
		Name: request.Name,
		Id:   request.Id,
	}
	b.BookRepository.Update(ctx, book)
}

func (b *BookServiceImpl) Delete(ctx context.Context, id int) {
	b.BookRepository.Delete(ctx, id)
}

func (b *BookServiceImpl) FindById(ctx context.Context, id int) response.BookResponse {
	res, err := b.BookRepository.FindById(ctx, id)
	helper.PanicIfError(err)
	return response.BookResponse(res)
}

func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	res, err := b.BookRepository.FindAll(ctx)
	helper.PanicIfError(err)

	var bookResp []response.BookResponse

	for _, value := range res {
		book := response.BookResponse{Id: value.Id, Name: value.Name}
		bookResp = append(bookResp, book)
	}

	return bookResp
}
