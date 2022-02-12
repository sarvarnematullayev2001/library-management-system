package postgres

import (
	"errors"
	"hw/prac/library/genproto/library_service"
	"hw/prac/library/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StuBookListRepo struct {
	db *sqlx.DB
}

func NewStuBookListRepo(db *sqlx.DB) repo.StuBookListRepoI {
	return &StuBookListRepo{
		db: db,
	}
}

func (bl *StuBookListRepo) Create(req *library_service.StuBookList) (string, error) {
	tx, err := bl.db.Begin()

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

	query := `UPDATE book SET num_books = num_books - 1 WHERE book_id = $1 AND num_books > 0`
	result, err := tx.Exec(query, req.BkId)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	query = `INSERT INTO stubook_list (bk_name, bk_authorname, bk_id, bk_numsbook)
		(SELECT book_name, author_name, book_id, num_books FROM book WHERE book_id = $1)`
	result, err = tx.Exec(query, req.BkId)
	if err != nil {
		return "", err
	}

	rowsAffected, err = result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	query = `UPDATE stubook_list SET student_id = $1 WHERE bk_id = $2`
	result, err = tx.Exec(query, req.StudentId, req.BkId)
	if err != nil {
		return "", err
	}

	rowsAffected, err = result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	query = `UPDATE stubook_list SET status = 'Not Given', deadline = $1 WHERE bk_id = $2`
	result, err = tx.Exec(query, req.Deadline, req.BkId)
	if err != nil {
		return "", err
	}

	rowsAffected, err = result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	return "created", nil
}

func (bl *StuBookListRepo) GetStudent(req *library_service.GetStudentLibrary) (*library_service.GetStudentLibraryInfo, error) {
	var (
		args    = make(map[string]interface{})
		filter  string
		student []*library_service.Student
	)

	if req.StudentFirstname != "" {
		filter += ` AND student_firstname = :student_firstname `
		args["student_firstname"] = req.StudentFirstname
	}

	if req.StudentLastname != "" {
		filter += ` AND student_lastname = :student_lastname `
		args["student_lastname"] = req.StudentLastname
	}

	if req.StudentId != "" {
		filter += ` AND student.student_id = :student_id `
		args["student_id"] = req.StudentId
	}

	query := `SELECT 
				student_id, student_firstname, student_lastname, student_faculty,
	 			student_course, student_phone1, student_phone2
			  FROM
			  	student
			  WHERE true ` + filter
	row, err := bl.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	if row.Next() {
		var s library_service.Student
		err = row.Scan(
			&s.StudentId,
			&s.StudentFirstname,
			&s.StudentLastname,
			&s.StudentFaculty,
			&s.StudentCourse,
			&s.StudentPhone1,
			&s.StudentPhone2,
		)

		if err != nil {
			return nil, err
		}

		query = `SELECT
					stubook_list_id, status, given_date, deadline,
					bk_name, bk_authorname, bk_id, bk_numsbook
			     FROM
				  	stubook_list
				 JOIN
				 	student
				 ON
					stubook_list.student_id = student.student_id 
				 WHERE 
				 	true ` + filter
		rows, err := bl.db.NamedQuery(query, args)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var stubooklist library_service.AllStuBookList
			err = rows.Scan(
				&stubooklist.BookListId,
				&stubooklist.Status,
				&stubooklist.GivenDate,
				&stubooklist.Deadline,
				&stubooklist.BkName,
				&stubooklist.BkAuthorname,
				&stubooklist.BkId,
				&stubooklist.BkNumsbook,
			)

			if err != nil {
				return nil, err
			}

			s.AllBooklist = append(s.AllBooklist, &stubooklist)

			defer rows.Close()
		}
		student = append(student, &s)
	}

	return &library_service.GetStudentLibraryInfo{
		StudentLibrary: student,
	}, nil
}

func (bl *StuBookListRepo) GetAllStudent(req *library_service.GetAllStudentLibraryRequest) (*library_service.GetAllStudentLibraryResponse, error) {
	var (
		students []*library_service.Student
		args     = make(map[string]interface{})
		count    uint32
		filter   string
	)

	countQuery := `SELECT count(1) FROM stubook_list`
	row, err := bl.db.NamedQuery(countQuery, args)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err = row.Scan(
			&count,
		)
		if err != nil {
			return nil, err
		}
	}

	filter += " OFFSET :offset LIMIT :limit "
	args["offset"] = req.Offset
	args["limit"] = req.Limit

	query := `SELECT
				student_id, student_firstname, student_lastname, 
				student_faculty, student_phone1, student_phone2, student_course
			  FROM
			  	student
			  WHERE true ` + filter
	rows, err := bl.db.NamedQuery(query, args)
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
			&student.StudentPhone1,
			&student.StudentPhone2,
			&student.StudentCourse,
		)
		if err != nil {
			return nil, err
		}

		query := `SELECT
					stubook_list_id, status, given_date, deadline,
					bk_name, bk_authorname, bk_id, bk_numsbook
			      FROM
				  	stubook_list WHERE student_id = $1`
		rows, err := bl.db.Query(query, student.StudentId)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var stubooklist library_service.AllStuBookList
			err := rows.Scan(
				&stubooklist.BookListId,
				&stubooklist.Status,
				&stubooklist.GivenDate,
				&stubooklist.Deadline,
				&stubooklist.BkName,
				&stubooklist.BkAuthorname,
				&stubooklist.BkId,
				&stubooklist.BkNumsbook,
			)
			if err != nil {
				return nil, err
			}
			student.AllBooklist = append(student.AllBooklist, &stubooklist)
		}
		students = append(students, &student)

	}
	return &library_service.GetAllStudentLibraryResponse{
		StudentLibrary: students,
		Count:          count,
	}, nil
}

func (bl *StuBookListRepo) Return(req *library_service.ReturnBook) (string, error) {
	query := `UPDATE book SET num_books=num_books + 1 WHERE book_id = $1`
	result, err := bl.db.Exec(query, req.BkId)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	query = `UPDATE stubook_list SET status='Given', bk_numsbook = (SELECT num_books FROM book WHERE book_id = $1) 
		WHERE bk_id = $1 AND stubook_list_id = $2`
	result, err = bl.db.Exec(query, req.BkId, req.BooklistId)
	if err != nil {
		return "", err
	}

	rowsAffected, err = result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	return "book returned", nil
}
