package service

import (
	"log"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/shotokan/ocr-for-text-file/internal/account/domain"
	"github.com/shotokan/ocr-for-text-file/internal/account/dto"
	"github.com/shotokan/ocr-for-text-file/pkg/ocr"
)

type AccountValidatorService struct{
	// deps as repositories
}

func NewAccountValidatorService() AccountValidatorService{
	return AccountValidatorService{}
}

func (avs AccountValidatorService) ValidateTextFromFile(file multipart.File) ([]dto.Account, error) {
	accountsValidated := ocr.Recognize(file)
	accountsChecksum := make([]dto.Account, 0, len(accountsValidated))
	for _, account := range accountsValidated {
		account, err := domain.NewAccount(uuid.NewString(), account)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		account.Validate()
		accountsChecksum = append(accountsChecksum, dto.Account{
			ID: account.ID(),
			Checksum: account.CheckSum(),
			Number: account.Number(),
		})
	}

	return accountsChecksum, nil
}