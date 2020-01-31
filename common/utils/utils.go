package utils

import "database/sql"

func CheckNullString(s string) sql.NullString {
	if len(s) < 1 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
