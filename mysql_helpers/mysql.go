package mysql_helpers

import (
	"github.com/Mixilino/helpers/errs"
	"github.com/Mixilino/logger_helper/logger"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errs.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			logger.Error("No record matching given informations", err)
			return errs.NewNotFoundError("No record matching given informations")
		}
		logger.Error("Error parsing sql error", err)
		return errs.NewInternalServerError()
	}
	switch sqlErr.Number {
	case 1062:
		logger.Error("Email already exist", err)
		return errs.NewBadRequestError("Email already exists.")
	}
	logger.Error("Error when trying to proccess request", err)
	return errs.NewInternalServerError()
}
