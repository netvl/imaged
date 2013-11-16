package database

import (
    "github.com/dpx-infinity/imaged/data"
    _ "github.com/jmoiron/sqlx"
)

type TagOps struct {
    *Database
}

func (db *Database) Tags() TagOps {
    return TagOps{db}
}

const query_create_tags = `
    create table tags (
        tag_id      integer primary key,
        name        text not null,
        description text,
        type        text
    )
`

func (db TagOps) createTable() error {
    _, err := db.db.Exec(query_create_tags)
    return err
}

const query_find_tag_by_id = `
    select * from tags where tag_id = ?
`

func (db TagOps) FindById(id int64) (tag data.Tag, err error) {
    err = db.db.Select(&tag, query_find_tag_by_id, id)
    return
}

const query_save_tag = `
    insert into tags (name, description, type) values (?, ?, ?)
`

func (db TagOps) Save(tag data.Tag) error {
    _, err := db.db.Exec(query_save_tag, tag.Name, tag.Description, tag.Type)
    return err
}
