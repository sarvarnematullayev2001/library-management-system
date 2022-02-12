package models

type ProBookList struct {
	BkId        string `json:"bk_id"`
	Deadline    string `json:"deadline"`
	ProfessorId string `json:"professor_id"`
}

type ProBookListId struct {
	ProBookListId uint32 `json:"probook_list_id"`
}

type GetAllProfessorLibraryRequest struct {
	Offset uint32 `json:"offset"`
	Limit  uint32 `json:"limit"`
}

type GetAllProfessorLibraryResponse struct {
	ProfessorLibrary []*Professor `json:"professor_library"`
	Count            uint32       `json:"count"`
}

type GetProfessorLibrary struct {
	ProfessorId        string `json:"professor_id"`
	ProfessorFirstName string `json:"professor_firstname"`
	ProfessorLastName  string `json:"professor_lastname"`
}

type GetProfessorLibraryInfo struct {
	ProfessorLibrary []*Professor `json:"professor_library"`
}
