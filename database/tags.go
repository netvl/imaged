package database

import (
    log "github.com/cihub/seelog"
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
    log.Debug("Creating tags table")
    _, err := db.db.Exec(query_create_tags)
    return err
}

func (db TagOps) prepareTable() {
    db.db.AddTable(data.Tag{}, "tags").SetKeys(true, "tag_id")
}

func (db TagOps) FindById(id int64) (tag data.Tag, err error) {
    err = db.db.Get(&tag, id)
    return
}

func (db TagOps) Insert(tag *data.Tag) error {
    return db.db.Insert(tag)
}
