package crypto_helpers

import (
	"github.com/Mixilino/helpers/errs"
	"github.com/Mixilino/logger_helper/logger"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(pass string) (string, *errs.RestErr) {
	password, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Error crypting password"+pass, err)
		return "", errs.NewInternalServerError()
	}
	return string(password), nil
}

func CompareTwoPasswords(passInput string, passInDb string) *errs.RestErr {
	if err := bcrypt.CompareHashAndPassword([]byte(passInDb), []byte(passInput)); err != nil {
		logger.Error("Wrong password input", err)
		return errs.NewBadRequestError("Wrong password")
	}
	return nil
}
