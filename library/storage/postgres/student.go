package postgres

import (
	"errors"
	"hw/prac/library/genproto/library_service"
	"hw/prac/library/storage/repo"

	"github.com/jmoiron/sqlx"
)

type studentRepo struct {
	db *sqlx.DB
}

func NewStudentRepo(db *sqlx.DB) repo.StudentRepoI {
	return &studentRepo{
		db: db,
	}
}

func (s *studentRepo) Create(req *library_service.Student) (string, error) {
	tx, err := s.db.Begin()

	if err != nil {
		return "", err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `INSERT INTO student (student_id, student_firstname, student_lastname, student_faculty,student_course,student_phone1,student_phone2) VALUES ($1, $2, $3, $4,$5,$6,$7)`

	_, err = tx.Exec(query, req.StudentId, req.StudentFirstname, req.StudentLastname, req.StudentFaculty, req.StudentCourse, req.StudentPhone1, req.StudentPhone2)
	if err != nil {
		return "", err
	}

	return "created", nil
}

func (s *studentRepo) GetAll(req *library_service.GetAllStudentRequest) (*library_service.GetAllStudentResponse, error) {
	var (
		students []*library_service.Student
		filter   string
		count    uint32
	)

	countQuery := `SELECT count(1) FROM student`
	rows, err := s.db.Query(countQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&count,
		)
		if err != nil {
			return nil, err
		}
	}

	filter += " OFFSET $1 LIMIT $2 "

	query := `SELECT
				student_id, student_firstname, student_lastname, 
				student_faculty,student_course,student_phone1,
				student_phone2 
			  FROM 
			    student 
			  WHERE true ` + filter
	rows, err = s.db.Query(query, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var student library_service.Student
		err := rows.Scan(
			&student.StudentId,
			&student.StudentFirstname,
			&student.StudentLastname,
			&student.StudentFaculty,
			&student.StudentCourse,
			&student.StudentPhone1,
			&student.StudentPhone2,
		)
		if err != nil {
			return nil, err
		}

		query := `SELECT
					stubook_list_id, status, given_date, deadline, 
					bk_name, bk_authorname, bk_id, bk_numsbook, student_id
				  FROM 
				  	stubook_list
				  WHERE
				  student_id = $1`
		rows, err := s.db.Query(query, student.StudentId)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var booklist library_service.AllStuBookList
			err = rows.Scan(
				&booklist.BookListId,
				&booklist.Status,
				&booklist.GivenDate,
				&booklist.Deadline,
				&booklist.BkName,
				&booklist.BkAuthorname,
				&booklist.BkId,
				&booklist.BkNumsbook,
				&booklist.StudentId,
			)
			if err != nil {
				return nil, err
			}

			student.AllBooklist = append(student.AllBooklist, &booklist)
		}

		students = append(students, &student)
	}

	return &library_service.GetAllStudentResponse{
		Students: students,
		Count:    count,
	}, nil
}

func (s *studentRepo) Get(req *library_service.StudentId) (*library_service.Student, error) {
	var student library_service.Student
	query := `SELECT student_id, student_firstname, student_lastname, student_faculty,student_course,student_phone1,student_phone2 FROM student WHERE student_id = $1`
	row := s.db.QueryRow(query, req.StudentId)
	err := row.Scan(&student.StudentId, &student.StudentFirstname, &student.StudentLastname, &student.StudentFaculty, &student.StudentCourse, &student.StudentPhone1, &student.StudentPhone2)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *studentRepo) Update(req *library_service.Student) (string, error) {
	query := `UPDATE student SET student_firstname = $1, student_lastname = $2, student_faculty = $3,student_course=$4,student_phone1=$5,student_phone2=$6 WHERE student_id = $7`
	result, err := s.db.Exec(query, req.StudentFirstname, req.StudentLastname, req.StudentFaculty, req.StudentCourse, req.StudentPhone1, req.StudentPhone2, req.StudentId)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	return "updated", nil
}

func (s *studentRepo) Delete(req *library_service.StudentId) (string, error) {
	query := `DELETE FROM student WHERE student_id = $1`
	result, err := s.db.Exec(query, req.StudentId)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	return "deleted", nil
}
