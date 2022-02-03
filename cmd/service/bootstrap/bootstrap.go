package bootstrap

import (
	"github.com/shotokan/ocr-for-text-file/internal/account/handler"
	"github.com/shotokan/ocr-for-text-file/internal/platform/server"
)

func Run() error {
	valAccountsHandler := handler.ValidateAccountsHandler{}
	srv := server.New("localhost", uint(8080), valAccountsHandler)
	
	return srv.Run()
}