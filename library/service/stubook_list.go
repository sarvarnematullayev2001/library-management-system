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

type StuBookListService struct {
	logger  logger.Logger
	storage storage.StorageI
	library_service.UnimplementedStuBookListServiceServer
}

func NewStuBookListService(logger logger.Logger, db *sqlx.DB) *StuBookListService {
	return &StuBookListService{
		logger:  logger,
		storage: storage.NewStoragePG(db),
	}
}

func (bl *StuBookListService) Create(ctx context.Context, req *library_service.StuBookList) (*library_service.Msg, error) {
	resp, err := bl.storage.StuBookList().Create(req)
	if err != nil {
		return nil, helper.HandleError(bl.logger, err, "error while creating student book list info", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}

func (bl *StuBookListService) GetStudent(ctx context.Context, req *library_service.GetStudentLibrary) (*library_service.GetStudentLibraryInfo, error) {
	resp, err := bl.storage.StuBookList().GetStudent(req)
	if err != nil {
		return nil, helper.HandleError(bl.logger, err, "error while getting student list info", req, http.StatusInternalServerError)
	}

	return resp, nil
}

func (bl *StuBookListService) GetAllStudent(ctx context.Context, req *library_service.GetAllStudentLibraryRequest) (*library_service.GetAllStudentLibraryResponse, error) {
	resp, err := bl.storage.StuBookList().GetAllStudent(req)
	if err != nil {
		return nil, helper.HandleError(bl.logger, err, "error while getting all student book list info", req, http.StatusInternalServerError)
	}

	return resp, nil
}

func (bl *StuBookListService) Return(ctx context.Context, req *library_service.ReturnBook) (*library_service.Msg, error) {
	resp, err := bl.storage.StuBookList().Return(req)

	if err != nil {
		return nil, helper.HandleError(bl.logger, err, "Error while updating student book info", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}
