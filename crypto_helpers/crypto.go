package crypto_helpers


import (
	"golang.org/x/crypto/bcrypt"
	"users_microservice/logger"
	"helpers/errs"
)

func GeneratePassword(pass string) (string, *errs.RestErr) {
	password, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Error crypting password" + pass, err)
		return "", errs.NewInternalServerError()
	}
	return string(password), nil
}

func CompareTwoPasswords(passInput string, passInDb string) *errs.RestErr{
	if err := bcrypt.CompareHashAndPassword([]byte(passInDb), []byte(passInput)); err!=nil{
		logger.Error("Wrong password input", err)
		return errs.NewBadRequestError("Wrong password")
	}
	return nil
}
