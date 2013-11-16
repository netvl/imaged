package data

import (
    "database/sql"
    "time"
)

type Tag struct {
    Id          int64 `db:"tag_id"`
    Name        string
    Description sql.NullString
    Type        sql.NullString
}

type File struct {
    Id           int64 `db:"file_id"`
    PrimaryTagId int64 `db:"primary_tag_id"`
    Hash         string
    Path         string
    DateAdded    time.Time `db:"date_added"`
    IsFavorite   bool      `db:"is_favorite"`
}

type GroupMapping struct {
    GroupId int64 `db:"group_id"`
    Indices []string
}

type Group struct {
    Id          int64         `db:"group_id"`
    TitleFileId sql.NullInt64 `db:"title_file_id"`
    Name        string
}
