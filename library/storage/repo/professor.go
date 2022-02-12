package repo

import "hw/prac/library/genproto/library_service"

type ProfessorRepoI interface {
	Create(req *library_service.Professor) (string, error)
	Get(req *library_service.GetProfessor) (*library_service.Professor, error)
	GetAll(req *library_service.GetAllProfessorRequest) (*library_service.GetAllProfessorResponse, error)
	Update(req *library_service.Professor) (string, error)
	Delete(req *library_service.GetProfessor) (string, error)
}
