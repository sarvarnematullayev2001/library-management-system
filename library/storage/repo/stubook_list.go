package repo

import "hw/prac/library/genproto/library_service"

type StuBookListRepoI interface {
	Create(req *library_service.StuBookList) (string, error)
	GetAllStudent(req *library_service.GetAllStudentLibraryRequest) (*library_service.GetAllStudentLibraryResponse, error)
	GetStudent(req *library_service.GetStudentLibrary) (*library_service.GetStudentLibraryInfo, error)
	Return(req *library_service.ReturnBook) (string, error)
}
