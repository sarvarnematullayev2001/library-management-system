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

type ProBookListService struct {
	logger  logger.Logger
	storage storage.StorageI
	library_service.UnimplementedProBookListServiceServer
}

func NewProBookListService(logger logger.Logger, db *sqlx.DB) *ProBookListService {
	return &ProBookListService{
		logger:  logger,
		storage: storage.NewStoragePG(db),
	}
}

func (bl *ProBookListService) Create(ctx context.Context, req *library_service.ProBookList) (*library_service.Msg, error) {
	resp, err := bl.storage.ProBookList().Create(req)
	if err != nil {
		return nil, helper.HandleError(bl.logger, err, "error while creating Professor book list info", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}

func (bl *ProBookListService) GetProfessor(ctx context.Context, req *library_service.GetProfessorLibrary) (*library_service.GetProfessorLibraryInfo, error) {
	resp, err := bl.storage.ProBookList().GetProfessor(req)
	if err != nil {
		return nil, helper.HandleError(bl.logger, err, "error while getting Professor list info", req, http.StatusInternalServerError)
	}

	return resp, nil
}

func (bl *ProBookListService) GetAllProfessor(ctx context.Context, req *library_service.GetAllProfessorLibraryRequest) (*library_service.GetAllProfessorLibraryResponse, error) {
	resp, err := bl.storage.ProBookList().GetAllProfessor(req)
	if err != nil {
		return nil, helper.HandleError(bl.logger, err, "error while getting all Professor book list info", req, http.StatusInternalServerError)
	}

	return resp, nil
}

func (bl *ProBookListService) Return(ctx context.Context, req *library_service.ReturnBook) (*library_service.Msg, error) {
	resp, err := bl.storage.ProBookList().Return(req)

	if err != nil {
		return nil, helper.HandleError(bl.logger, err, "Error while updating Professor book info", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}
