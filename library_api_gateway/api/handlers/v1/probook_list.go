package v1

import (
	"context"
	"hw/prac/library_api_gateway/api/models"
	"hw/prac/library_api_gateway/genproto/library_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Professor BookList godoc
// @ID create-professor-booklist
// @Router /v1/probooklist [POST]
// @Summary create professor booklist
// @Description Create Professor BookList
// @Tags professorbooklist
// @Accept json
// @Produce json
// @Param probook_list body models.ProBookList true "probook_list"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateProfessorBookList(c *gin.Context) {
	var booklist models.ProBookList

	if err := c.BindJSON(&booklist); err != nil {
		h.handleErrorResponse(c, 400, "error while binding json", err)
		return
	}

	resp, err := h.services.ProBookListService().Create(
		context.Background(),
		&library_service.ProBookList{
			BkId:        booklist.BkId,
			Deadline:    booklist.Deadline,
			ProfessorId: booklist.ProfessorId,
		},
	)

	if !handleError(h.log, c, err, "error while creating student book list") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get Professor BookList godoc
// @ID get-professor-booklist
// @Router /v1/probooklist [GET]
// @Summary get professor booklist
// @Description Get Professor BookList
// @Tags professorbooklist
// @Accept json
// @Produce json
// @Param professor_id query string false "professor_id"
// @Param professor_firstname query string false "professor_firstname"
// @Param professor_lastname  query string false "professor_lastname"
// @Success 200 {object} models.ResponseModel{data=models.GetProfessorLibraryInfo} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetProfessorBookList(c *gin.Context) {
	var booklist models.GetProfessorLibraryInfo

	resp, err := h.services.ProBookListService().GetProfessor(
		context.Background(),
		&library_service.GetProfessorLibrary{
			ProfessorId:        c.Query("professor_id"),
			ProfessorFirstname: c.Query("professor_firstname"),
			ProfessorLastname:  c.Query("professor_lastname"),
		},
	)

	if !handleError(h.log, c, err, "error while getting book") {
		return
	}

	err = ParseToStruct(&booklist, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", booklist)
}

// Get Professor BookLists godoc
// @ID get-professor-booklists
// @Router /v1/probooklists [GET]
// @Summary get professor booklists
// @Description Get Professor BookLists
// @Tags professorbooklist
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.ResponseModel{data=models.GetAllProfessorLibraryResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllProfessorBookList(c *gin.Context) {
	var booklist models.GetAllProfessorLibraryResponse

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	limit, err := h.ParseQueryParam(c, "limit", "100")
	if err != nil {
		return
	}

	resp, err := h.services.ProBookListService().GetAllProfessor(
		context.Background(),
		&library_service.GetAllProfessorLibraryRequest{
			Offset: uint32(offset),
			Limit:  uint32(limit),
		},
	)

	if !handleError(h.log, c, err, "error while getting all professor booklist") {
		return
	}

	err = ParseToStruct(&booklist, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", booklist)
}

// Return Professor BookList godoc
// @ID Return-probook
// @Router /v1/probooklist [PUT]
// @Summary Return probook
// @Description Return ProBook
// @Tags professorbooklist
// @Accept json
// @Produce json
// @Param bk_id query string true "bk_id"
// @Param probook_list_id query int true "probook_list_id"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) ReturnProBook(c *gin.Context) {

	booklistId, err := strconv.Atoi(c.Query("probook_list_id"))
	if !handleError(h.log, c, err, "error while converting string to int") {
		return
	}
	resp, err := h.services.ProBookListService().Return(
		context.Background(),
		&library_service.ReturnBook{
			BkId:       c.Query("bk_id"),
			BooklistId: uint32(booklistId),
		},
	)
	if !handleError(h.log, c, err, "error while returning book") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
