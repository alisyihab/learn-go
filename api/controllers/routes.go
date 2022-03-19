package controllers

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		// Login Route
		v1.POST("/login", s.Login)
	}
}