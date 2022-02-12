package api

import (
	v1 "hw/prac/library_api_gateway/api/handlers/v1"
	"hw/prac/library_api_gateway/config"
	"hw/prac/library_api_gateway/pkg/logger"
	"hw/prac/library_api_gateway/services"

	_ "hw/prac/library_api_gateway/api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// @Summary 登录
	// @Description 登录
	// @Produce json
	// @Param body body controllers.LoginParams true "body参数"
	// @Success 200 {string} string "ok" "返回用户信息"
	// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
	// @Failure 401 {string} string "err_code：10001 登录失败"
	// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
	// @Router /user/person/login [post]
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	router.GET("/config", handlerV1.GetConfig)

	apiV1 := router.Group("/v1")
	apiV1.GET("/ping", handlerV1.Ping)

	// book
	apiV1.POST("/book", handlerV1.CreateBook)
	apiV1.GET("/book", handlerV1.GetBook)
	apiV1.GET("/books", handlerV1.GetAllBook)
	apiV1.PUT("/book", handlerV1.UpdateBook)
	apiV1.DELETE("/book", handlerV1.DeleteBook)

	// professor
	apiV1.POST("/professor", handlerV1.CreateProfessor)
	apiV1.GET("/professor", handlerV1.GetProfessor)
	apiV1.GET("/professors", handlerV1.GetAllProfessor)
	apiV1.PUT("/professor", handlerV1.UpdateProfessor)
	apiV1.DELETE("/professor", handlerV1.DeleteProfessor)

	//stubooklist
	apiV1.POST("/stubooklist", handlerV1.CreateStudentBookList)
	apiV1.GET("/stubooklist", handlerV1.GetStudentBookList)
	apiV1.GET("/stubooklists", handlerV1.GetAllStudentBookList)
	apiV1.PUT("/stubooklist", handlerV1.ReturnStuBook)

	//probooklist
	apiV1.POST("/probooklist", handlerV1.CreateProfessorBookList)
	apiV1.GET("/probooklist", handlerV1.GetProfessorBookList)
	apiV1.GET("/probooklists", handlerV1.GetAllProfessorBookList)
	apiV1.PUT("/probooklist", handlerV1.ReturnProBook)

	// student
	apiV1.POST("/student", handlerV1.CreateStudent)
	apiV1.GET("/students", handlerV1.GetAllStudents)
	apiV1.GET("/student", handlerV1.GetStudent)
	apiV1.PUT("/student", handlerV1.UpdateStudent)
	apiV1.DELETE("/student", handlerV1.DeleteStudent)

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
