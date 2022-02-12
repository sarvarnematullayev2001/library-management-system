package service

import (
	"context"
	"hw/prac/library/genproto/library_service"
	"hw/prac/library/pkg/helper"
	"hw/prac/library/pkg/logger"
	"hw/prac/library/storage"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type bookService struct {
	storage storage.StorageI
	logger  logger.Logger
	library_service.UnimplementedBookServiceServer
}

func NewBookService(logger logger.Logger, db *sqlx.DB) *bookService {
	return &bookService{
		logger:  logger,
		storage: storage.NewStoragePG(db),
	}
}

func (b *bookService) Create(ctx context.Context, req *library_service.Book) (*library_service.Msg, error) {
	resp, err := b.storage.Book().Create(req)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while creating book", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}

func (b *bookService) Get(ctx context.Context, req *library_service.GetBook) (*library_service.Book, error) {
	resp, err := b.storage.Book().Get(req)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while getting book", req, http.StatusInternalServerError)
	}

	return &library_service.Book{
		BookId:     resp.BookId,
		BookName:   resp.BookName,
		AuthorName: resp.AuthorName,
		NumBooks:   resp.NumBooks,
	}, nil
}

func (b *bookService) GetAll(ctx context.Context, req *library_service.GetAllBookRequest) (*library_service.GetAllBookResponse, error) {
	resp, err := b.storage.Book().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while getting all books", req, http.StatusInternalServerError)
	}

	return resp, nil
}

func (b *bookService) Update(ctx context.Context, req *library_service.Book) (*library_service.Msg, error) {
	resp, err := b.storage.Book().Update(req)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while updating book", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}

func (b *bookService) Delete(ctx context.Context, req *library_service.GetBook) (*library_service.Msg, error) {
	resp, err := b.storage.Book().Delete(req)
	if err != nil {
		return nil, helper.HandleError(b.logger, err, "error while deleting book", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}
