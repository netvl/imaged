package database

import (
    "github.com/dpx-infinity/imaged/common"
    "github.com/dpx-infinity/imaged/data"
    . "github.com/smartystreets/goconvey/convey"
    "os"
    "path/filepath"
    "testing"
)

func tempLogFile() string {
    return filepath.Join(os.TempDir(), "imaged_test.sqlite")
}

func TestInit(t *testing.T) {
    common.DisableLogs()

    tempfile := tempLogFile()

    Convey("Database file should be properly initialized", t, func() {
        os.Remove(tempfile)

        db, err := InitializeWithPath(tempfile)
        So(err, ShouldBeNil)
        defer db.Close()
    })
}

func TestInsert(t *testing.T) {
    common.DisableLogs()

    tempfile := tempLogFile()

    var tag *data.Tag
    var db *Database
    var err error

    Convey("Given an opened database", t, func() {
        os.Remove(tempfile)

        db, err = InitializeWithPath(tempfile)
        So(err, ShouldBeNil)

        Convey("When inserting a tag object", func() {
            tag = &data.Tag{
                Name:        "tag",
                Description: NullString("descr"),
                Type:        NullString("type"),
            }

            err = db.Tags().Insert(tag)

            Convey("Then there should be no errors", func() {
                So(err, ShouldBeNil)
            })

            Convey("And inserted tag could be read back", func() {
                tag2, err := db.Tags().FindById(tag.Id)

                So(err, ShouldBeNil)
                So(tag2, ShouldResemble, *tag)
            })

            Convey("Test1", func() {
            })

            Convey("Test2", func() {
            })

            Convey("Test3", func() {
            })
        })

        Reset(func() {
            db.Close()
        })
    })
}
