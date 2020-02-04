package utils

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"regexp"
)

var (
	mobileRegex *regexp.Regexp
	emailRegex  *regexp.Regexp
)

func init() {
	// from https://github.com/VincentSit/ChinaMobilePhoneNumberRegex
	mobileRegex = regexp.MustCompile(`^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[01356789]\d{2}|6[2567]\d{2}|4[579]\d{2})\d{6}$`)
	// from https://github.com/badoux/checkmail/blob/master/checkmail.go
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
}

func CheckNullString(s string) sql.NullString {
	if len(s) < 1 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func IskMySQLError(err error, errNum int) bool {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return ok
	}
	return mysqlErr.Number == uint16(errNum)
}

func NoErrors(errs ...error) error {
	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	return nil
}

func ValidateMobile(num string) bool {
	return mobileRegex.MatchString(num)
}

func ValidateEmailFormat(addr string) bool {
	return emailRegex.MatchString(addr)
}
