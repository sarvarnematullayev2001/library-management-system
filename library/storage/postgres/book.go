package postgres

import (
	"errors"
	"hw/prac/library/genproto/library_service"
	"hw/prac/library/storage/repo"

	"github.com/jmoiron/sqlx"
)

type bookRepo struct {
	db *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) repo.BookRepoI {
	return &bookRepo{
		db: db,
	}
}

func (b *bookRepo) Create(req *library_service.Book) (string, error) {
	tx, err := b.db.Begin()

	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `INSERT INTO book (book_id, book_name, author_name, num_books) VALUES ($1, $2, $3, $4)`

	_, err = tx.Exec(query, req.BookId, req.BookName, req.AuthorName, req.NumBooks)
	if err != nil {
		return "", err
	}

	return "created", nil
}

func (b *bookRepo) Get(req *library_service.GetBook) (*library_service.Book, error) {
	var book library_service.Book

	query := `SELECT book_id, book_name, author_name, num_books FROM book WHERE book_id = $1`
	row := b.db.QueryRow(query, req.BookId)
	err := row.Scan(&book.BookId, &book.BookName, &book.AuthorName, &book.NumBooks)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *bookRepo) GetAll(req *library_service.GetAllBookRequest) (*library_service.GetAllBookResponse, error) {
	var (
		books  []*library_service.Book
		args   = make(map[string]interface{})
		filter string
		count  uint32
	)

	if req.BookName != "" {
		filter += ` AND book_name ilike '%' || :book_name || '%' `
		args["book_name"] = req.BookName
	}

	countQuery := `SELECT count(1) FROM book WHERE true ` + filter
	rows, err := b.db.NamedQuery(countQuery, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&count,
		)
		if err != nil {
			return nil, err
		}
	}

	filter += " OFFSET :offset LIMIT :limit "
	args["offset"] = req.Offset
	args["limit"] = req.Limit

	query := `SELECT book_id, book_name, author_name, num_books FROM book WHERE true ` + filter
	rows, err = b.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book library_service.Book
		err := rows.Scan(
			&book.BookId,
			&book.BookName,
			&book.AuthorName,
			&book.NumBooks,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return &library_service.GetAllBookResponse{
		Books: books,
		Count: count,
	}, nil
}

func (b *bookRepo) Update(req *library_service.Book) (string, error) {
	query := `UPDATE book SET book_name = $1, author_name = $2, num_books = $3 WHERE book_id = $4`
	result, err := b.db.Exec(query, req.BookName, req.AuthorName, req.NumBooks, req.BookId)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	return "updated", nil
}

func (b *bookRepo) Delete(req *library_service.GetBook) (string, error) {

	query := `DELETE FROM book WHERE book_id = $1`
	result, err := b.db.Exec(query, req.BookId)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	return "deleted", nil
}
