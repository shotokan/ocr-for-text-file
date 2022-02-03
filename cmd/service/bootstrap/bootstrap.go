package bootstrap

import (
	"github.com/shotokan/ocr-for-text-file/internal/account/handler"
	"github.com/shotokan/ocr-for-text-file/internal/account/service"
	"github.com/shotokan/ocr-for-text-file/internal/platform/server"
)

func Run() error {
	accountValidatorService := service.NewAccountValidatorService()
	valAccountsHandler := handler.ValidateAccountsHandler{
		AccountValidator: accountValidatorService,
	}
	srv := server.New("localhost", uint(8080), valAccountsHandler)
	
	return srv.Run()
}