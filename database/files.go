package database

type FileOps struct {
    *Database
}

func (db *Database) Files() FileOps {
    return FileOps{db}
}

const query_create_files = `
    create table files (
        file_id        integer primary key,
        primary_tag_id integer not null,
        hash           text not null,
        path           text not null,
        date_added     integer not null,
        is_favorite    integer not null,
        constraint tags_primary_tag_id_fkey 
            foreign key (primary_tag_id) references tags (tag_id)
            on delete restrict
    )
`

func (db FileOps) createTable() error {
    _, err := db.db.Exec(query_create_files)
    return err
}
