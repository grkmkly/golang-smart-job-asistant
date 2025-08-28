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
		protected.GET("/profile", s.UserHandler.GetProfile()) // GET PROFILE

		protected.GET("/announcements", s.AnnouncementHandler.GetAnnouncements()) // GET ANNOUNCEMENTS

		protected.GET("/jobposts", s.JobPostHandler.ListJobPosts()) // GET ACTIVE JOB

		protected.GET("jobquestions/:jobpostID", s.JobQuestionHandler.GetQuestionUserForPost()) // GET QUESTION FOR JOB

		protected.POST("/jobposts/:jobpostID/applications", s.ApplicationHandler.SubmitApplication()) // POST APPLICATION

		protected.GET("/me/applications/status", s.ApplicationHandler.GetUserApplications()) // GET APPLICATION STATUS FOR USER

		admin := protected.Group("admin")
		admin.Use(auth.AuthAdminMiddleWare())
		{
			admin.POST("/announcements", s.AnnouncementHandler.CreateAnnouncement()) // POST ANNOUNCEMENTS

			admin.POST("/jobposts", s.JobPostHandler.CreateNewJobPost())                      // POST NEW JOB
			admin.GET("/jobposts", s.JobPostHandler.ListJobPostsForAdmin())                   // LIST JOB POST FOR ADMIN (WITH CRITERIA)
			admin.GET("/jobposts/:jobpostID", s.JobQuestionHandler.GetQuestionAdminForPost()) // GET QUESTION FOR POST

			admin.POST("/questions", s.QuestionHandler.CreateQuestion())          // POST NEW QUESTION
			admin.GET("/questions", s.QuestionHandler.GetQuestionWithOption())    // GET QUESTION WITH OPTION
			admin.POST("/jobquestions", s.JobQuestionHandler.CreateJobQuestion()) // JOB POST NEW QUESTION

			admin.GET("/jobposts/:jobpostID/applications", s.ApplicationHandler.GetApplicationsByPostID()) // GET APPLICATION
			admin.PUT("/applications/:applicationID/status", s.ApplicationHandler.UpdateStatus())          // UPDATE STATUS
		}
	}
}
