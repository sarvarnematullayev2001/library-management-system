package v1

import (
	"context"
	"hw/prac/library_api_gateway/api/models"
	"hw/prac/library_api_gateway/genproto/library_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Professor godoc
// @ID create-professor
// @Router /v1/professor [POST]
// @Summary create professor
// @Description Create Professor
// @Tags professor
// @Accept json
// @Produce json
// @Param professor body models.Professor true "professor"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateProfessor(c *gin.Context) {
	var professor models.Professor

	if err := c.BindJSON(&professor); err != nil {
		h.handleErrorResponse(c, 400, "error while binding json", err)
		return
	}

	resp, err := h.services.ProfessorService().Create(
		context.Background(),
		&library_service.Professor{
			ProfessorId:        professor.ProfessorId,
			ProfessorFirstname: professor.ProfessorFirstName,
			ProfessorLastname:  professor.ProfessorLastName,
			ProfessorPhone1:    professor.ProfessorPhone1,
			ProfessorPhone2:    professor.ProfessorPhone2,
		},
	)

	if !handleError(h.log, c, err, "error while creating professor") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get Professor godoc
// @ID get-professor
// @Router /v1/professor [GET]
// @Summary get professor
// @Description Get Professor
// @Tags professor
// @Accept json
// @Produce json
// @Param professor_id query string true "professor_id"
// @Success 200 {object} models.ResponseModel{data=models.Professor} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetProfessor(c *gin.Context) {
	var professor models.Professor

	resp, err := h.services.ProfessorService().Get(
		context.Background(),
		&library_service.GetProfessor{
			ProfessorId: c.Query("professor_id"),
		},
	)

	if !handleError(h.log, c, err, "error while getting professor") {
		return
	}

	err = ParseToStruct(&professor, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", professor)
}

// Get Professors godoc
// @ID get-professors
// @Router /v1/professors [GET]
// @Summary get professors
// @Description Get Professors
// @Tags professor
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.ResponseModel{data=models.GetAllProfessorResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllProfessor(c *gin.Context) {
	var professors models.GetAllProfessorResponse

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	limit, err := h.ParseQueryParam(c, "limit", "100")
	if err != nil {
		return
	}

	resp, err := h.services.ProfessorService().GetAll(
		context.Background(),
		&library_service.GetAllProfessorRequest{
			Offset: uint32(offset),
			Limit:  uint32(limit),
		},
	)

	if !handleError(h.log, c, err, "error while getting professors") {
		return
	}

	err = ParseToStruct(&professors, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", professors)
}

// Update Professor godoc
// @ID update-professor
// @Router /v1/professor [PUT]
// @Summary update professor
// @Description Update Professor
// @Tags professor
// @Accept json
// @Produce json
// @Param professor body models.Professor true "professor"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateProfessor(c *gin.Context) {
	var professor models.Professor

	if err := c.BindJSON(&professor); err != nil {
		h.handleErrorResponse(c, 400, "error while binding json", err)
		return
	}

	resp, err := h.services.ProfessorService().Update(
		context.Background(),
		&library_service.Professor{
			ProfessorId:        professor.ProfessorId,
			ProfessorFirstname: professor.ProfessorFirstName,
			ProfessorLastname:  professor.ProfessorLastName,
			ProfessorPhone1:    professor.ProfessorPhone1,
			ProfessorPhone2:    professor.ProfessorPhone2,
		},
	)

	if !handleError(h.log, c, err, "error while updating professor") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Delete Professor godoc
// @ID delete-professor
// @Router /v1/professor [DELETE]
// @Summary delete professor
// @Description Delete Professor
// @Tags professor
// @Accept json
// @Produce json
// @Param professor_id query string true "professor_id"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteProfessor(c *gin.Context) {

	resp, err := h.services.ProfessorService().Delete(
		context.Background(),
		&library_service.GetProfessor{
			ProfessorId: c.Query("professor_id"),
		},
	)

	if !handleError(h.log, c, err, "error while deleting professor") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
