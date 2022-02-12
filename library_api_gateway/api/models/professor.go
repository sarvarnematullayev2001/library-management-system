package models

type Professor struct {
	ProfessorId        string            `json:"professor_id"`
	ProfessorFirstName string            `json:"professor_firstname"`
	ProfessorLastName  string            `json:"professor_lastname"`
	ProfessorPhone1    string            `json:"professor_phone1"`
	ProfessorPhone2    string            `json:"professor_phone2"`
	AllBookList        []*AllProBookList `json:"all_booklist"`
}

type GetProfessor struct {
	ProfessorId string `json:"professor_id"`
}

type GetAllProfessorRequest struct {
	Offset uint32 `json:"offset"`
	Limit  uint32 `json:"limit"`
}

type GetAllProfessorResponse struct {
	Professors []*Professor `json:"professors"`
	Count      uint32       `json:"count"`
}

type AllProBookList struct {
	ProfessorId  string `json:"professor_id"`
	BookListId   uint32 `json:"book_list_id"`
	Status       string `json:"status"`
	GivenDate    string `json:"given_date"`
	Deadline     string `json:"deadline"`
	BkName       string `json:"bk_name"`
	BkAuthorName string `json:"bk_authorname"`
	BkId         string `json:"bk_id"`
	BkNumsBook   uint32 `json:"bk_numsbook"`
}
