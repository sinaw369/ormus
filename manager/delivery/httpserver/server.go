package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/ormushq/ormus/manager"
	"github.com/ormushq/ormus/manager/delivery/httpserver/sourcehandler"
	"github.com/ormushq/ormus/manager/delivery/httpserver/userhandler"
)

type SetupServicesResponse struct {
	UserHandler *userhandler.Handler
}

type Server struct {
	config        manager.Config
	userHandler   userhandler.Handler
	sourceHandler sourcehandler.Handler
	Router        *echo.Echo
}

func New(cfg manager.Config, setupSvc SetupServicesResponse) *Server {
	return &Server{
		config:      cfg,
		userHandler: *setupSvc.UserHandler,
		Router:      echo.New(),
	}
}

func (s *Server) Server() {
	e := echo.New()

	s.userHandler.SetUserRoute(e)
	s.sourceHandler.SetSourceRoute(e)

	e.GET("/health-check", s.healthCheck)

	e.Logger.Fatal(e.Start(":8080"))
}
