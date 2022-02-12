package postgres

import (
	"errors"
	"hw/prac/library/genproto/library_service"
	"hw/prac/library/storage/repo"

	"github.com/jmoiron/sqlx"
)

type ProfessorRepo struct {
	db *sqlx.DB
}

func NewProfessorRepo(db *sqlx.DB) repo.ProfessorRepoI {
	return &ProfessorRepo{
		db: db,
	}
}

func (p *ProfessorRepo) Create(req *library_service.Professor) (string, error) {
	tx, err := p.db.Begin()
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

	query := `INSERT INTO professor (professor_id, professor_firstname, professor_lastname,
			professor_phone1, professor_phone2) VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.Exec(query, req.ProfessorId, req.ProfessorFirstname,
		req.ProfessorLastname, req.ProfessorPhone1, req.ProfessorPhone2)

	if err != nil {
		return "", err
	}

	return "created", nil
}

func (p *ProfessorRepo) Get(req *library_service.GetProfessor) (*library_service.Professor, error) {
	var professor library_service.Professor
	query := `SELECT professor_id, professor_firstname, professor_lastname, professor_phone1, professor_phone2 FROM professor WHERE professor_id = $1`
	row := p.db.QueryRow(query, req.ProfessorId)
	err := row.Scan(&professor.ProfessorId, &professor.ProfessorFirstname, &professor.ProfessorLastname, &professor.ProfessorPhone1, &professor.ProfessorPhone2)
	if err != nil {
		return nil, err
	}

	return &professor, nil
}

func (p *ProfessorRepo) GetAll(req *library_service.GetAllProfessorRequest) (*library_service.GetAllProfessorResponse, error) {
	var (
		filter     string
		professors []*library_service.Professor
		count      uint32
	)

	countQuery := `SELECT count(1) FROM professor`
	row, err := p.db.Query(countQuery)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err := row.Scan(
			&count,
		)
		if err != nil {
			return nil, err
		}
	}

	filter += ` OFFSET $1 LIMIT $2 `

	query := `SELECT
				professor_id, professor_firstname, professor_lastname, 
				professor_phone1, professor_phone2
			  FROM 
			  	professor 
			  WHERE true ` + filter
	rows, err := p.db.Query(query, req.Offset, req.Limit)
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
					bk_name, bk_authorname, bk_id, bk_numsbook, professor_id
				  FROM 
				  	probook_list
				  WHERE
				  	professor_id = $1`
		rows, err := p.db.Query(query, professor.ProfessorId)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var booklist library_service.AllProBookList
			err = rows.Scan(
				&booklist.BookListId,
				&booklist.Status,
				&booklist.GivenDate,
				&booklist.Deadline,
				&booklist.BkName,
				&booklist.BkAuthorname,
				&booklist.BkId,
				&booklist.BkNumsbook,
				&booklist.ProfessorId,
			)

			if err != nil {
				return nil, err
			}

			professor.AllBooklist = append(professor.AllBooklist, &booklist)
		}

		professors = append(professors, &professor)
	}

	return &library_service.GetAllProfessorResponse{
		Professors: professors,
		Count:      count,
	}, nil
}

func (p *ProfessorRepo) Update(req *library_service.Professor) (string, error) {
	query := `UPDATE professor SET professor_firstname = $1, professor_lastname = $2, professor_phone1 = $3, professor_phone2 = $4 WHERE professor_id = $5`
	result, err := p.db.Exec(query, req.ProfessorFirstname, req.ProfessorLastname, req.ProfessorPhone1, req.ProfessorPhone2, req.ProfessorId)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	return "updated", nil
}

func (p *ProfessorRepo) Delete(req *library_service.GetProfessor) (string, error) {
	query := `DELETE FROM professor WHERE professor_id = $1`
	result, err := p.db.Exec(query, req.ProfessorId)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return "", errors.New("NOT FOUND")
	}

	return "deleted", nil
}
