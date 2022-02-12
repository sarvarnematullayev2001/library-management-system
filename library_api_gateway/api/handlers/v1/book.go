package v1

import (
	"context"
	"hw/prac/library_api_gateway/api/models"
	"hw/prac/library_api_gateway/genproto/library_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Book godoc
// @ID create-book
// @Router /v1/book [POST]
// @Summary create book
// @Description Create Book
// @Tags book
// @Accept json
// @Produce json
// @Param book body models.Book true "book"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.BindJSON(&book); err != nil {
		h.handleErrorResponse(c, 400, "error while binding json", err)
		return
	}

	resp, err := h.services.BookService().Create(
		context.Background(),
		&library_service.Book{
			BookId:     book.BookId,
			BookName:   book.BookName,
			AuthorName: book.AuthorName,
			NumBooks:   book.NumsBook,
		},
	)

	if !handleError(h.log, c, err, "error while creating book") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get Book godoc
// @ID get-book
// @Router /v1/book [GET]
// @Summary get book
// @Description Get Book
// @Tags book
// @Accept json
// @Produce json
// @Param book_id query string true "book_id"
// @Success 200 {object} models.ResponseModel{data=models.Book} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetBook(c *gin.Context) {
	var book models.Book

	resp, err := h.services.BookService().Get(
		context.Background(),
		&library_service.GetBook{
			BookId: c.Query("book_id"),
		},
	)

	if !handleError(h.log, c, err, "error while getting book") {
		return
	}

	err = ParseToStruct(&book, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get Books godoc
// @ID get-books
// @Router /v1/books [GET]
// @Summary get books
// @Description Get Books
// @Tags book
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param book_name query string false "book_name"
// @Success 200 {object} models.ResponseModel{data=models.GetAllBookResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllBook(c *gin.Context) {
	var books models.GetAllBookResponse

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	limit, err := h.ParseQueryParam(c, "limit", "100")
	if err != nil {
		return
	}

	resp, err := h.services.BookService().GetAll(
		context.Background(),
		&library_service.GetAllBookRequest{
			Offset:   uint32(offset),
			Limit:    uint32(limit),
			BookName: c.Query("book_name"),
		},
	)

	if !handleError(h.log, c, err, "error while getting books") {
		return
	}

	err = ParseToStruct(&books, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Update Book godoc
// @ID update-book
// @Router /v1/book [PUT]
// @Summary update book
// @Description Update Book
// @Tags book
// @Accept json
// @Produce json
// @Param book body models.Book true "book"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateBook(c *gin.Context) {
	var book models.Book

	if err := c.BindJSON(&book); err != nil {
		h.handleErrorResponse(c, 400, "error while binding json", err)
		return
	}

	resp, err := h.services.BookService().Update(
		context.Background(),
		&library_service.Book{
			BookId:     book.BookId,
			BookName:   book.BookName,
			AuthorName: book.AuthorName,
			NumBooks:   book.NumsBook,
		},
	)

	if !handleError(h.log, c, err, "error while updating book") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Delete Book godoc
// @ID delete-book
// @Router /v1/book [DELETE]
// @Summary delete book
// @Description Delete Book
// @Tags book
// @Accept json
// @Produce json
// @Param book_id query string true "book_id"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteBook(c *gin.Context) {

	resp, err := h.services.BookService().Delete(
		context.Background(),
		&library_service.GetBook{
			BookId: c.Query("book_id"),
		},
	)

	if !handleError(h.log, c, err, "error while deleting book") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
