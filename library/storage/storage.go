package storage

import (
	"hw/prac/library/storage/postgres"
	"hw/prac/library/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StoragePG struct {
	db           *sqlx.DB
	book         repo.BookRepoI
	student      repo.StudentRepoI
	professor    repo.ProfessorRepoI
	stubook_list repo.StuBookListRepoI
	probook_list repo.ProBookListRepoI
}

type StorageI interface {
	Book() repo.BookRepoI
	Student() repo.StudentRepoI
	Professor() repo.ProfessorRepoI
	StuBookList() repo.StuBookListRepoI
	ProBookList() repo.ProBookListRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return &StoragePG{
		db:           db,
		book:         postgres.NewBookRepo(db),
		student:      postgres.NewStudentRepo(db),
		professor:    postgres.NewProfessorRepo(db),
		stubook_list: postgres.NewStuBookListRepo(db),
		probook_list: postgres.NewProBookListRepo(db),
	}
}

func (s *StoragePG) Book() repo.BookRepoI {
	return s.book
}

func (s *StoragePG) Student() repo.StudentRepoI {
	return s.student
}

func (s *StoragePG) Professor() repo.ProfessorRepoI {
	return s.professor
}

func (s *StoragePG) StuBookList() repo.StuBookListRepoI {
	return s.stubook_list
}

func (s *StoragePG) ProBookList() repo.ProBookListRepoI {
	return s.probook_list
}
