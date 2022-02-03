package server

import (
	"fmt"
	"log"

	"github.com/shotokan/ocr-for-text-file/internal/account/handler"

	"github.com/labstack/echo"
)

type Server struct {
	httpAddr string
	engine *echo.Echo
	// deps
	validateAccountsHdlr handler.ValidateAccountsHandler
}

func New(host string, port uint, validateAccountsHdlr handler.ValidateAccountsHandler) Server{
	e := echo.New()
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine: e,
		validateAccountsHdlr: validateAccountsHdlr,
	}
	srv.registerUserRoutes()
	return srv
}

func (s Server) Run() error{
	log.Println("Server running on", s.httpAddr)
	return s.engine.Start(s.httpAddr)
}

func (s Server) registerUserRoutes() {
	s.engine.POST("/accounts/validate", s.validateAccountsHdlr.ValidateAccountsTextHandler())
}