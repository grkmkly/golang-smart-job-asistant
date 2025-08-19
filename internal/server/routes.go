package server

import "smartjob/internal/auth"

func (s *Server) SetupRoutes() {

	userRoutes := s.Router.Group("/public")
	{
		userRoutes.POST("/register", s.AuthHandler.RegisterUser())
		userRoutes.POST("/login", s.AuthHandler.LoginUser())
		userRoutes.POST("/refresh", s.AuthHandler.RefreshToken())
	}
	protected := s.Router.Group("/api")
	protected.Use(auth.AuthMiddleWare())
	{
		protected.GET("/profile", s.UserHandler.GetProfile())

		protected.GET("/announcements", s.AnnouncementHandler.GetAnnouncements())

		protected.GET("/jobposts", s.JobPostHandler.ListJobPosts())
		protected.GET("jobquestions/:jobpostID", s.JobQuestionHandler.GetQuestionForPost())

		protected.POST("/applications", s.ApplicationHandler.SubmitApplication())

		admin := protected.Group("admin")
		admin.Use(auth.AuthAdminMiddleWare())
		{
			admin.POST("/announcements", s.AnnouncementHandler.CreateAnnouncement())

			admin.POST("/jobposts", s.JobPostHandler.CreateNewJobPost())

			admin.POST("/questions", s.QuestionHandler.CreateQuestion())
			admin.GET("/questions", s.QuestionHandler.GetQuestionWithOption())
			admin.POST("/jobquestions", s.JobQuestionHandler.CreateJobQuestion())
		}
	}
}
