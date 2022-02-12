package v1

import (
	"context"
	"hw/prac/library_api_gateway/api/models"
	"hw/prac/library_api_gateway/genproto/library_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Student BookList godoc
// @ID create-student-booklist
// @Router /v1/stubooklist [POST]
// @Summary create student booklist
// @Description Create Student BookList
// @Tags studentbooklist
// @Accept json
// @Produce json
// @Param stubook_list body models.StuBookList true "stubook_list"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateStudentBookList(c *gin.Context) {
	var booklist models.StuBookList

	if err := c.BindJSON(&booklist); err != nil {
		h.handleErrorResponse(c, 400, "error while binding json", err)
		return
	}

	resp, err := h.services.StuBookListService().Create(
		context.Background(),
		&library_service.StuBookList{
			BkId:      booklist.BkId,
			Deadline:  booklist.Deadline,
			StudentId: booklist.StudentId,
		},
	)

	if !handleError(h.log, c, err, "error while creating student book list") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get Student BookList godoc
// @ID get-student-booklist
// @Router /v1/stubooklist [GET]
// @Summary get student booklist
// @Description Get Student BookList
// @Tags studentbooklist
// @Accept json
// @Produce json
// @Param student_id query string false "student_id"
// @Param student_firstname query string false "student_firstname"
// @Param student_lastname  query string false "student_lastname"
// @Success 200 {object} models.ResponseModel{data=models.GetStudentLibraryInfo} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetStudentBookList(c *gin.Context) {
	var booklist models.GetStudentLibraryInfo

	resp, err := h.services.StuBookListService().GetStudent(
		context.Background(),
		&library_service.GetStudentLibrary{
			StudentId:        c.Query("student_id"),
			StudentFirstname: c.Query("student_firstname"),
			StudentLastname:  c.Query("student_lastname"),
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

// Get Student BookLists godoc
// @ID get-student-booklists
// @Router /v1/stubooklists [GET]
// @Summary get student booklists
// @Description Get Student BookLists
// @Tags studentbooklist
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.ResponseModel{data=models.GetAllStudentLibraryResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllStudentBookList(c *gin.Context) {
	var booklist models.GetAllStudentLibraryResponse

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	limit, err := h.ParseQueryParam(c, "limit", "100")
	if err != nil {
		return
	}

	resp, err := h.services.StuBookListService().GetAllStudent(
		context.Background(),
		&library_service.GetAllStudentLibraryRequest{
			Offset: uint32(offset),
			Limit:  uint32(limit),
		},
	)

	if !handleError(h.log, c, err, "error while getting all student booklist") {
		return
	}

	err = ParseToStruct(&booklist, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", booklist)
}

// Return Student BookList godoc
// @ID return-student-booklist
// @Router /v1/stubooklist [PUT]
// @Summary return student booklist
// @Description Return Student BookList
// @Tags studentbooklist
// @Accept json
// @Produce json
// @Param bk_id query string true "bk_id"
// @Param stubooklist_id query int true "stubooklist_id"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) ReturnStuBook(c *gin.Context) {
	booklist_id, err := strconv.Atoi(c.Query("stubooklist_id"))
	if !handleError(h.log, c, err, "error while converting string to int") {
		return
	}

	resp, err := h.services.StuBookListService().Return(
		context.Background(),
		&library_service.ReturnBook{
			BkId:       c.Query("bk_id"),
			BooklistId: uint32(booklist_id),
		},
	)

	if !handleError(h.log, c, err, "error while returning student book") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
