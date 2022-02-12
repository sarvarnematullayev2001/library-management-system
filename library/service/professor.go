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

type ProfessorService struct {
	logger  logger.Logger
	storage storage.StorageI
	library_service.UnimplementedProfessorServiceServer
}

func NewProfessorService(logger logger.Logger, db *sqlx.DB) *ProfessorService {
	return &ProfessorService{
		logger:  logger,
		storage: storage.NewStoragePG(db),
	}
}

func (p *ProfessorService) Create(ctx context.Context, req *library_service.Professor) (*library_service.Msg, error) {
	resp, err := p.storage.Professor().Create(req)
	if err != nil {
		return nil, helper.HandleError(p.logger, err, "error while creating professor info", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}

func (p *ProfessorService) Get(ctx context.Context, req *library_service.GetProfessor) (*library_service.Professor, error) {
	resp, err := p.storage.Professor().Get(req)
	if err != nil {
		return nil, helper.HandleError(p.logger, err, "error while getting professor info", req, http.StatusInternalServerError)
	}

	return &library_service.Professor{
		ProfessorId:        resp.ProfessorId,
		ProfessorFirstname: resp.ProfessorFirstname,
		ProfessorLastname:  resp.ProfessorLastname,
		ProfessorPhone1:    resp.ProfessorPhone1,
		ProfessorPhone2:    resp.ProfessorPhone2,
	}, nil
}

func (p *ProfessorService) GetAll(ctx context.Context, req *library_service.GetAllProfessorRequest) (*library_service.GetAllProfessorResponse, error) {
	resp, err := p.storage.Professor().GetAll(req)
	if err != nil {
		return nil, helper.HandleError(p.logger, err, "error while getting all professor info", req, http.StatusInternalServerError)
	}

	return resp, nil
}

func (p *ProfessorService) Update(ctx context.Context, req *library_service.Professor) (*library_service.Msg, error) {
	resp, err := p.storage.Professor().Update(req)
	if err != nil {
		return nil, helper.HandleError(p.logger, err, "error while updating professor info", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}

func (p *ProfessorService) Delete(ctx context.Context, req *library_service.GetProfessor) (*library_service.Msg, error) {
	resp, err := p.storage.Professor().Delete(req)
	if err != nil {
		return nil, helper.HandleError(p.logger, err, "error while deleting professor info", req, http.StatusInternalServerError)
	}

	return &library_service.Msg{
		Msg: resp,
	}, nil
}
