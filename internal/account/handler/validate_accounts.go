package handler

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo"
	"github.com/shotokan/ocr-for-text-file/internal/account/dto"
)

type AccountValidator interface {
	ValidateTextFromFile(file multipart.File) ([]dto.Account, error)
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

		accountsValidated, err := va.AccountValidator.ValidateTextFromFile(src)
		if err != nil {
			return c.JSON(http.StatusBadRequest, struct{
				message string
			}{message: err.Error()})
		}

		return c.JSON(http.StatusOK, accountsValidated)
	}
	
}