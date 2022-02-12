package repo

import "hw/prac/library/genproto/library_service"

type StudentRepoI interface {
	Create(req *library_service.Student) (string, error)
	GetAll(req *library_service.GetAllStudentRequest) (*library_service.GetAllStudentResponse, error)
	Get(req *library_service.StudentId) (*library_service.Student, error)
	Update(req *library_service.Student) (string, error)
	Delete(req *library_service.StudentId) (string, error)
}
