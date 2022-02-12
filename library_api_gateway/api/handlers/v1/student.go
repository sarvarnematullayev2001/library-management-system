package v1

import (
	"context"
	"hw/prac/library_api_gateway/api/models"
	"hw/prac/library_api_gateway/genproto/library_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Student godoc
// @ID create_student
// @Router /v1/student [POST]
// @Summary Create Student
// @Description Create Student
// @Tags student
// @Accept json
// @Produce json
// @Param student body models.Student true "student"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateStudent(c *gin.Context) {
	var student models.Student

	err := c.BindJSON(&student)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.StudentService().Create(
		context.Background(),
		&library_service.Student{
			StudentId:        student.StudentId,
			StudentFirstname: student.StudentFirstName,
			StudentLastname:  student.StudentLastName,
			StudentFaculty:   student.StudentFaculty,
			StudentCourse:    student.StudentCourse,
			StudentPhone1:    student.StudentPhone1,
			StudentPhone2:    student.StudentPhone2,
		},
	)
	if !handleError(h.log, c, err, "error while creating student") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get All Student godoc
// @ID get_all_student
// @Router /v1/students [GET]
// @Summary Get All Student
// @Description Get All Student
// @Tags student
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} models.ResponseModel{data=models.GetAllStudentResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllStudents(c *gin.Context) {
	var students models.GetAllStudentResponse

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	resp, err := h.services.StudentService().GetAll(
		context.Background(),
		&library_service.GetAllStudentRequest{
			Offset: uint32(offset),
			Limit:  uint32(limit),
		},
	)

	if !handleError(h.log, c, err, "error while getting all students") {
		return
	}

	err = ParseToStruct(&students, resp)
	if !handleError(h.log, c, err, "error while parse to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", students)
}

// Get Student godoc
// @ID get_student
// @Router /v1/student [GET]
// @Summary Get Student
// @Description Get Student
// @Tags student
// @Accept json
// @Produce json
// @Param student_id query string true "student_id"
// @Success 200 {Object} models.ResponseModel{data=models.GetAllStudentResponse} "desc"
// @Response 400 {Object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetStudent(c *gin.Context) {
	var student models.Student

	resp, err := h.services.StudentService().Get(
		context.Background(),
		&library_service.StudentId{
			StudentId: c.Query("student_id"),
		},
	)
	if !handleError(h.log, c, err, "error while getting student") {
		return
	}

	err = ParseToStruct(&student, resp)
	if !handleError(h.log, c, err, "error while parsing to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", student)
}

// Update Student godoc
// @ID update_student
// @Router /v1/student [PUT]
// @Summary Update Student
// @Description Update Student by ID
// @Tags student
// @Accept json
// @Produce json
// @Param student body models.Student true "student"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateStudent(c *gin.Context) {
	var student models.Student

	if err := c.BindJSON(&student); err != nil {
		h.handleErrorResponse(c, 400, "error while binding json", err)
		return
	}
	resp, err := h.services.StudentService().Update(
		context.Background(),
		&library_service.Student{
			StudentId:        student.StudentId,
			StudentFirstname: student.StudentFirstName,
			StudentLastname:  student.StudentLastName,
			StudentFaculty:   student.StudentFaculty,
			StudentCourse:    student.StudentCourse,
			StudentPhone1:    student.StudentPhone1,
			StudentPhone2:    student.StudentPhone2,
		},
	)
	if !handleError(h.log, c, err, "error while updating student") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Delete Student godoc
// @ID delete_student
// @Router /v1/student [DELETE]
// @Summary Delete Student
// @Description Delete Student by given ID
// @Tags student
// @Accept json
// @Produce json
// @Param student_id query string true "student_id"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteStudent(c *gin.Context) {
	resp, err := h.services.StudentService().Delete(
		context.Background(),
		&library_service.StudentId{
			StudentId: c.Query("student_id"),
		},
	)

	if !handleError(h.log, c, err, "error while getting student") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
