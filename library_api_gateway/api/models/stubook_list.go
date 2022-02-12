package models

type StuBookList struct {
	BkId      string `json:"bk_id"`
	Deadline  string `json:"deadline"`
	StudentId string `json:"student_id"`
}

type StuBookListId struct {
	BookListId uint32 `json:"stubook_list_id"`
}

type GetAllStudentLibraryRequest struct {
	Offset uint32 `json:"offset"`
	Limit  uint32 `json:"limit"`
}

type GetAllStudentLibraryResponse struct {
	StudentLibrary []*Student `json:"student_library"`
	Count          uint32     `json:"count"`
}

type GetStudentLibrary struct {
	StudentId        string `json:"student_id"`
	StudentFirstName string `json:"student_firstname"`
	StudentLastName  string `json:"student_lastname"`
}

type GetStudentLibraryInfo struct {
	StudentLibrary []*Student `json:"student_library"`
}

type ReturnBook struct {
	BkId       string `json:"bk_id"`
	BookListId uint32
}
