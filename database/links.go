package database

const (
    query_create_files_tags_map = `
        create table files_tags_map (
            file_id integer,
            tag_id  integer,
            constraint files_tags_map_pkey         
                primary key (file_id, tag_id),
            constraint files_tags_map_file_id_fkey 
                foreign key (file_id) references files (file_id) 
                on delete cascade,
            constraint files_tags_map_tag_id_fkey  
                foreign key (tag_id)  references tags (tag_id) 
                on delete cascade
        )
    `
    query_create_files_groups_map = `
        create table files_groups_map (
            file_id  integer,
            group_id integer,
            indices  text default '',
            constraint files_groups_map_pkey          
                primary key (file_id, group_id),
            constraint files_groups_map_file_id_fkey  
                foreign key (file_id)  references files (file_id)
                on delete cascade,
            constraint files_groups_map_group_id_fkey 
                foreign key (group_id) references groups (group_id)
                on delete cascade
        )
    `
)

func (db *Database) establishKeys() error {
    _, err := db.Exec(query_create_files_tags_map)
    if err != nil {
        return err
    }

    _, err = db.Exec(query_create_files_groups_map)
    return err
}
