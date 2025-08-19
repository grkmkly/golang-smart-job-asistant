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
		protected.GET("jobquestions/:jobpostID", s.JobQuestionHandler.GetQuestionUserForPost())

		protected.POST("/applications/:jobpostID", s.ApplicationHandler.SubmitApplication())

		admin := protected.Group("admin")
		admin.Use(auth.AuthAdminMiddleWare())
		{
			admin.POST("/announcements", s.AnnouncementHandler.CreateAnnouncement())

			admin.POST("/jobposts", s.JobPostHandler.CreateNewJobPost())
			admin.GET("/jobposts", s.JobPostHandler.ListJobPostsForAdmin())
			admin.GET("/jobposts/:jobpostID", s.JobQuestionHandler.GetQuestionAdminForPost())

			admin.POST("/questions", s.QuestionHandler.CreateQuestion())
			admin.GET("/questions", s.QuestionHandler.GetQuestionWithOption())
			admin.POST("/jobquestions", s.JobQuestionHandler.CreateJobQuestion())
			admin.GET("/applications/:post_id", s.ApplicationHandler.GetApplicationsByPostID())
		}
	}
}
