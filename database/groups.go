package database

type GroupOps struct {
    *Database
}

func (db *Database) Groups() GroupOps {
    return GroupOps{db}
}

const query_create_groups = `
    create table groups (
        group_id      integer primary key,
        title_file_id integer null,
        name          text not null,
        constraint files_title_file_id_fkey
            foreign key (title_file_id) references files (file_id)
            on delete set null
    )
`

func (db GroupOps) createTable() error {
    _, err := db.Exec(query_create_groups)
    return err
}
