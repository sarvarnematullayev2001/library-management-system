package models

type Student struct {
	StudentId        string            `json:"student_id"`
	StudentFirstName string            `json:"student_firstname"`
	StudentLastName  string            `json:"student_lastname"`
	StudentFaculty   string            `json:"student_faculty"`
	StudentCourse    uint32            `json:"student_course"`
	StudentPhone1    string            `json:"student_phone1"`
	StudentPhone2    string            `json:"student_phone2"`
	AllBookList      []*AllStuBookList `json:"all_booklist"`
}

type GetAllStudentResponse struct {
	Students []*Student `json:"students"`
	Count    uint32     `json:"count"`
}

type AllStuBookList struct {
	StudentId    string `json:"student_id"`
	BookListId   uint32 `json:"book_list_id"`
	Status       string `json:"status"`
	GivenDate    string `json:"given_date"`
	Deadline     string `json:"deadline"`
	BkName       string `json:"bk_name"`
	BkAuthorName string `json:"bk_authorname"`
	BkId         string `json:"bk_id"`
	BkNumsBook   uint32 `json:"bk_numsbook"`
}
