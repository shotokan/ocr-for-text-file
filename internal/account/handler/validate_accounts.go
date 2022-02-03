package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/shotokan/ocr-for-text-file/pkg/ocr"
)

type AccountValidator interface {
	ValidateText() error
}

type ValidateAccountsHandler struct {
	AccountValidator
}


func (va ValidateAccountsHandler) ValidateAccountsTextHandler() echo.HandlerFunc{
	return func(c echo.Context) error {
		name := c.FormValue("name")
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		log.Println(fmt.Sprintf("name: %v, file size: %v", name, file.Size))
 
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		
		accountsValidated := ocr.Recognize(src)
		return c.JSON(http.StatusOK, accountsValidated)
	}
	
}