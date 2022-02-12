package models

type Book struct {
	BookId     string `json:"book_id"`
	BookName   string `json:"book_name"`
	AuthorName string `json:"author_name"`
	NumsBook   uint32 `json:"nums_book" default:"100"`
}

type GetBook struct {
	BookId string `json:"book_id"`
}

type Msg struct {
	Msg string `json:"msg"`
}

type GetAllBookRequest struct {
	Offset   uint32 `json:"offset"`
	Limit    uint32 `json:"limit"`
	BookName string `json:"book_name"`
}

type GetAllBookResponse struct {
	Books []*Book `json:"books"`
	Count uint32  `json:"count"`
}
