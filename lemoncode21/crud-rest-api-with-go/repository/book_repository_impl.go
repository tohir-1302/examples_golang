package repository

import (
	"context"
	"database/sql"
	"errors"
	"rest-api-golang/helper"
	"rest-api-golang/model"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(Db *sql.DB) BookRepository {
	return &BookRepositoryImpl{
		Db: Db,
	}
}

func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "DELETE FROM books WHERE id = $1"

	_, errExec := tx.ExecContext(ctx, SQL, bookId)
	helper.PanicIfError(errExec)
}

func (b *BookRepositoryImpl) FindAll(ctx context.Context) ([]model.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name FROM books"

	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer result.Close()

	var books []model.Book

	for result.Next() {
		book := model.Book{}
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepositoryImpl) FindById(ctx context.Context, id int) (model.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name FROM books WHERE id = $1"

	result, errExec := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(errExec)
	defer result.Close()

	book := model.Book{}

	if result.Next() {
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("book not found")
	}

}

func (b *BookRepositoryImpl) Save(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "INSERT INTO books (name) VALUES ($1)"
	_, errExec := tx.QueryContext(ctx, SQL, book.Name)
	helper.PanicIfError(errExec)
}

func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "UPDATE books SET name = $1 WHERE id = $2"
	_, errExec := tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helper.PanicIfError(errExec)
}
