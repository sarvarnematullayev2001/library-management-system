package services

import (
	"fmt"
	"hw/prac/library_api_gateway/config"
	"hw/prac/library_api_gateway/genproto/library_service"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	StudentService() library_service.StudentServiceClient
	ProfessorService() library_service.ProfessorServiceClient
	BookService() library_service.BookServiceClient
	StuBookListService() library_service.StuBookListServiceClient
	ProBookListService() library_service.ProBookListServiceClient
}

type grpcClients struct {
	studentService     library_service.StudentServiceClient
	professorService   library_service.ProfessorServiceClient
	bookService        library_service.BookServiceClient
	stuBookListService library_service.StuBookListServiceClient
	proBookListService library_service.ProBookListServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connLibraryService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.LibraryServiceHost, conf.LibraryServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		studentService:     library_service.NewStudentServiceClient(connLibraryService),
		professorService:   library_service.NewProfessorServiceClient(connLibraryService),
		bookService:        library_service.NewBookServiceClient(connLibraryService),
		stuBookListService: library_service.NewStuBookListServiceClient(connLibraryService),
		proBookListService: library_service.NewProBookListServiceClient(connLibraryService),
	}, nil

}

func (g *grpcClients) StudentService() library_service.StudentServiceClient {
	return g.studentService
}

func (g *grpcClients) ProfessorService() library_service.ProfessorServiceClient {
	return g.professorService
}

func (g *grpcClients) BookService() library_service.BookServiceClient {
	return g.bookService
}

func (g *grpcClients) StuBookListService() library_service.StuBookListServiceClient {
	return g.stuBookListService
}

func (g *grpcClients) ProBookListService() library_service.ProBookListServiceClient {
	return g.proBookListService
}
