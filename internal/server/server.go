package server

import (
	"log"
	"smartjob/internal/database"
	"smartjob/internal/handlers"
	"smartjob/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	//Dependicies
	DB     *gorm.DB
	Router *gin.Engine
	//Handlers
	UserHandler         *handlers.UserHandler
	AuthHandler         *handlers.AuthHandler
	AnnouncementHandler *handlers.AnnouncementHandler
	JobPostHandler      *handlers.JobPostHandler
	QuestionHandler     *handlers.QuestionHandler
	JobQuestionHandler  *handlers.JobQuestionHandler
	ApplicationHandler  *handlers.ApplicationHandler
}

func NewServer() *Server {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env Error", err)
	}
	router := gin.Default()

	database.Connect()
	database.AutoMigrate()

	db := database.DB

	//Services
	userService := *services.NewUserService(db)
	tokenService := *services.NewTokenService(db, &userService)
	authService := *services.NewAuthService(db, &userService, &tokenService)
	announcementService := *services.NewAnnouncementService(db)
	jobpostService := *services.NewJobPostService(db)
	questionService := *services.NewQuestionService(db)
	jobQuestionService := *services.NewJobQuestionService(db)
	applicationService := *services.NewApplicationService(db)

	//Handlers
	authHandler := *handlers.NewAuthHandler(&authService)
	userHandler := *handlers.NewUserHandler(&userService)
	announcementHandler := *handlers.NewAnnouncementHandler(&announcementService)
	jobpostHandler := *handlers.NewJobPostHandler(&jobpostService)
	questionHandler := *handlers.NewQuestionHandler(&questionService)
	jobQuestionHandler := *handlers.NewJobQuestioHandler(&jobQuestionService)
	applicationHandler := *handlers.NewApplicationHandler(&applicationService)

	return &Server{
		DB:                  db,
		Router:              router,
		UserHandler:         &userHandler,
		AuthHandler:         &authHandler,
		AnnouncementHandler: &announcementHandler,
		JobPostHandler:      &jobpostHandler,
		QuestionHandler:     &questionHandler,
		JobQuestionHandler:  &jobQuestionHandler,
		ApplicationHandler:  &applicationHandler,
	}
}
func (s *Server) Run() {

	s.SetupRoutes()

	s.Router.Run(":8080")
}
