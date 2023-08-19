package server

import (
	"githab.com/dmmoody/go-microservices/internal/database"
	"githab.com/dmmoody/go-microservices/internal/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed to start: %s", err)
		return err
	}
	return s.echo.Start(":8080")
}

func (s *EchoServer) registerRoutes() {
	s.echo.GET("/health/readiness", s.Readiness)
	s.echo.GET("/health/liveness", s.Liveness)
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.HealthStatus{Status: "OK"})
	}
	return ctx.JSON(http.StatusInternalServerError, models.HealthStatus{Status: "FAILURE"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HealthStatus{Status: "OK"})
}
