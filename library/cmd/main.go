package main

import (
	"fmt"
	"hw/prac/library/config"
	"hw/prac/library/genproto/library_service"
	"hw/prac/library/pkg/logger"
	"hw/prac/library/service"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.Load()

	log := logger.New(c.Environment, "library_service")
	defer logger.Cleanup(log)

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDB,
		"disable",
	)

	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		log.Error("error while connecting database", logger.Error(err))
		return
	}

	listen, err := net.Listen("tcp", c.RPCPort)
	if err != nil {
		log.Error("Error while listening port: %v", logger.Error(err))
		return
	}

	bookService := service.NewBookService(log, db)
	professorService := service.NewProfessorService(log, db)
	studentService := service.NewStudentService(log, db)
	stuBookListService := service.NewStuBookListService(log, db)
	proBookListService := service.NewProBookListService(log, db)

	server := grpc.NewServer()
	reflection.Register(server)

	library_service.RegisterBookServiceServer(server, bookService)
	library_service.RegisterProfessorServiceServer(server, professorService)
	library_service.RegisterStudentServiceServer(server, studentService)
	library_service.RegisterStuBookListServiceServer(server, stuBookListService)
	library_service.RegisterProBookListServiceServer(server, proBookListService)

	log.Info("main: server running", logger.String("port", c.RPCPort))

	if err := server.Serve(listen); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
	}

}
