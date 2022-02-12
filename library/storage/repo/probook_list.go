package repo

import "hw/prac/library/genproto/library_service"

type ProBookListRepoI interface {
	Create(req *library_service.ProBookList) (string, error)
	GetAllProfessor(req *library_service.GetAllProfessorLibraryRequest) (*library_service.GetAllProfessorLibraryResponse, error)
	GetProfessor(req *library_service.GetProfessorLibrary) (*library_service.GetProfessorLibraryInfo, error)
	Return(req *library_service.ReturnBook) (string, error)
}
