package util

import "database/sql"

func NullStringToPointer(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	} else {
		return nil
	}
}

func PointerToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{
			Valid: false,
		}
	} else {
		return sql.NullString{
			Valid:  true,
			String: *s,
		}
	}
}
