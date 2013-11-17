package database

import (
    "database/sql"
)

func NullString(s string) sql.NullString {
    return sql.NullString{s, true}
}
