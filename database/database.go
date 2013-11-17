package database

import (
    "database/sql"
    log "github.com/cihub/seelog"
    "github.com/dpx-infinity/imaged/common"
    "github.com/dpx-infinity/imaged/config"
    "github.com/jmoiron/modl"
    _ "github.com/mattn/go-sqlite3"
    "os"
)

type Database struct {
    db *modl.DbMap
}

func Initialize(conf *config.Config) (*Database, error) {
    return InitializeWithPath(conf.Paths.DatabaseFile())
}

func InitializeWithPath(databaseFile string) (*Database, error) {
    dbExists, err := checkDatabaseExists(databaseFile)
    if err != nil {
        return nil, common.NewError("Cannot access database file", err)
    }

    db, err := sql.Open("sqlite3", databaseFile)
    if err != nil {
        return nil, common.NewError("Cannot open database", err)
    }
    dbmap := modl.NewDbMap(db, modl.SqliteDialect{})

    database := &Database{dbmap}

    if !dbExists {
        log.Info("Creating tables")
        if err = database.createTables(); err != nil {
            return nil, common.NewError("Cannot create tables", err)
        }
    }

    database.prepareTables()

    return database, nil
}

func checkDatabaseExists(databaseFile string) (bool, error) {
    if _, err := os.Stat(databaseFile); err != nil {
        if os.IsNotExist(err) {
            log.Infof("Database file does not exist: %s", databaseFile)
            return false, nil
        } else {
            return false, err
        }
    }
    log.Infof("Found database file: %s", databaseFile)
    return true, nil
}

func (db *Database) prepareTables() {
    db.Tags().prepareTable()
}

func (db *Database) createTables() error {
    if err := db.Tags().createTable(); err != nil {
        log.Debugf("Failed: %s", err)
        return err
    }

    if err := db.Files().createTable(); err != nil {
        log.Debugf("Failed: %s", err)
        return err
    }

    if err := db.Groups().createTable(); err != nil {
        log.Debugf("Failed: %s", err)
        return err
    }

    return db.establishKeys()
}

func (db *Database) Close() {
    err := db.db.Db.Close()
    if err != nil {
        log.Warnf("Error closing database: %?", err)
    }
}
