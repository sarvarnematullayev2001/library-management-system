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

type studentService struct {
	storage storage.StorageI
	logger  logger.Logger
	library_service.UnimplementedStudentServiceServer
}

func NewStudentService(logger logger.Logger, db *sqlx.DB) *studentService {
	return &studentService{
		logger:  logger,
		storage: storage.NewStoragePG(db),
	}
}

func (s *studentService) Create(ctx context.Context, req *library_service.Student) (*library_service.Msg, error) {
	resp, err := s.storage.Student().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating student", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}

func (s *studentService) GetAll(ctx context.Context, req *library_service.GetAllStudentRequest) (*library_service.GetAllStudentResponse, error) {
	resp, err := s.storage.Student().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all students", req, http.StatusInternalServerError)
	}

	return resp, nil
}

func (s *studentService) Get(ctx context.Context, req *library_service.StudentId) (*library_service.Student, error) {
	resp, err := s.storage.Student().Get(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting student", req, http.StatusInternalServerError)
	}
	return &library_service.Student{
		StudentId:        req.StudentId,
		StudentFirstname: resp.StudentFirstname,
		StudentLastname:  resp.StudentLastname,
		StudentFaculty:   resp.StudentFaculty,
		StudentCourse:    resp.StudentCourse,
		StudentPhone1:    resp.StudentPhone1,
		StudentPhone2:    resp.StudentPhone2,
	}, nil
}

func (s *studentService) Update(ctx context.Context, req *library_service.Student) (*library_service.Msg, error) {
	resp, err := s.storage.Student().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating student", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}
func (s *studentService) Delete(ctx context.Context, req *library_service.StudentId) (*library_service.Msg, error) {
	resp, err := s.storage.Student().Delete(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting StudentS", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}
