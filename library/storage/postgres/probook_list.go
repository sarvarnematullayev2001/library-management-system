package postgres

import (
	"errors"
	"hw/prac/library/genproto/library_service"
	"hw/prac/library/storage/repo"

	"github.com/jmoiron/sqlx"
)

type ProBookListRepo struct {
	db *sqlx.DB
}

func NewProBookListRepo(db *sqlx.DB) repo.ProBookListRepoI {
	return &ProBookListRepo{
		db: db,
	}
}

func (bl *ProBookListRepo) Create(req *library_service.ProBookList) (string, error) {
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
	query = `INSERT INTO probook_list (bk_name, bk_authorname, bk_id, bk_numsbook)
		(SELECT book_name, author_name, book_id, num_books FROM book WHERE book_id = $1)`
	result, err = tx.Exec(query, req.BkId)
	if err != nil {
		return "", err
	}

	rowsAffected, err = result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	query = `UPDATE probook_list SET professor_id = $1 WHERE bk_id = $2`
	_, err = tx.Exec(query, req.ProfessorId, req.BkId)
	if err != nil {
		return "", err
	}

	query = `UPDATE probook_list SET status = 'Not Given', deadline = $1 WHERE bk_id = $2`
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

func (bl *ProBookListRepo) GetProfessor(req *library_service.GetProfessorLibrary) (*library_service.GetProfessorLibraryInfo, error) {
	var (
		args       = make(map[string]interface{})
		filter     string
		professors []*library_service.Professor
	)

	if req.ProfessorFirstname != "" {
		filter += ` AND professor_firstname = :professor_firstname `
		args["professor_firstname"] = req.ProfessorFirstname
	}

	if req.ProfessorLastname != "" {
		filter += ` AND professor_lastname = :professor_lastname `
		args["professor_lastname"] = req.ProfessorLastname
	}

	if req.ProfessorId != "" {
		filter += ` AND professor.professor_id = :professor_id `
		args["professor_id"] = req.ProfessorId
	}

	query := `SELECT 
				professor_id, professor_firstname, professor_lastname,
				professor_phone1, professor_phone2
			  FROM
			  	professor
			  WHERE true ` + filter
	row, err := bl.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	if row.Next() {
		var professor library_service.Professor
		err = row.Scan(
			&professor.ProfessorId,
			&professor.ProfessorFirstname,
			&professor.ProfessorLastname,
			&professor.ProfessorPhone1,
			&professor.ProfessorPhone2,
		)

		if err != nil {
			return nil, err
		}

		query = `SELECT
					probook_list_id, status, given_date, deadline,
					bk_name, bk_authorname, bk_id, bk_numsbook
			     FROM
				  	probook_list
				 JOIN
				 	professor
				 ON
					probook_list.professor_id = professor.professor_id 
				 WHERE 
				 	true ` + filter
		rows, err := bl.db.NamedQuery(query, args)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var probooklist library_service.AllProBookList
			err = rows.Scan(
				&probooklist.BookListId,
				&probooklist.Status,
				&probooklist.GivenDate,
				&probooklist.Deadline,
				&probooklist.BkName,
				&probooklist.BkAuthorname,
				&probooklist.BkId,
				&probooklist.BkNumsbook,
			)

			if err != nil {
				return nil, err
			}

			professor.AllBooklist = append(professor.AllBooklist, &probooklist)

			defer rows.Close()
		}
		professors = append(professors, &professor)
	}

	return &library_service.GetProfessorLibraryInfo{
		ProfessorLibrary: professors,
	}, nil
}

func (bl *ProBookListRepo) GetAllProfessor(req *library_service.GetAllProfessorLibraryRequest) (*library_service.GetAllProfessorLibraryResponse, error) {
	var (
		professors []*library_service.Professor
		args       = make(map[string]interface{})
		count      uint32
		filter     string
	)

	countQuery := `SELECT count(1) FROM probook_list WHERE true ` + filter
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
				professor_id, professor_firstname, professor_lastname, 
				professor_phone1, professor_phone2
			  FROM
			  	professor
			  WHERE true ` + filter
	rows, err := bl.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var professor library_service.Professor
		err := rows.Scan(
			&professor.ProfessorId,
			&professor.ProfessorFirstname,
			&professor.ProfessorLastname,
			&professor.ProfessorPhone1,
			&professor.ProfessorPhone2,
		)
		if err != nil {
			return nil, err
		}

		query := `SELECT
					probook_list_id, status, given_date, deadline,
					bk_name, bk_authorname, bk_id, bk_numsbook
			      FROM
				  	probook_list WHERE professor_id = $1`
		rows, err := bl.db.Query(query, professor.ProfessorId)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var probooklist library_service.AllProBookList
			err := rows.Scan(
				&probooklist.BookListId,
				&probooklist.Status,
				&probooklist.GivenDate,
				&probooklist.Deadline,
				&probooklist.BkName,
				&probooklist.BkAuthorname,
				&probooklist.BkId,
				&probooklist.BkNumsbook,
			)
			if err != nil {
				return nil, err
			}
			professor.AllBooklist = append(professor.AllBooklist, &probooklist)
		}
		professors = append(professors, &professor)

	}
	return &library_service.GetAllProfessorLibraryResponse{
		ProfessorLibrary: professors,
		Count:            count,
	}, nil
}

func (bl *ProBookListRepo) Return(req *library_service.ReturnBook) (string, error) {
	query := `UPDATE book SET num_books=num_books + 1 WHERE book_id = $1`
	result, err := bl.db.Exec(query, req.BkId)
	if err != nil {
		return "", err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	query = `UPDATE probook_list SET status='Given', bk_numsbook = (SELECT num_books FROM book WHERE book_id = $1) 
		WHERE bk_id = $1 AND probook_list_id = $2`
	result, err = bl.db.Exec(query, req.BkId, req.BooklistId)
	if err != nil {
		return "", err
	}
	rowsAffected, err = result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	return "Book returned", nil
}
