package repo

import "hw/prac/library/genproto/library_service"

type BookRepoI interface {
	Create(req *library_service.Book) (string, error)
	Get(req *library_service.GetBook) (*library_service.Book, error)
	GetAll(req *library_service.GetAllBookRequest) (*library_service.GetAllBookResponse, error)
	Update(req *library_service.Book) (string, error)
	Delete(req *library_service.GetBook) (string, error)
}
