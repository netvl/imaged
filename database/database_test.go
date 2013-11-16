package database

import (
    "github.com/dpx-infinity/imaged/common"
    . "github.com/smartystreets/goconvey/convey"
    "os"
    "path/filepath"
    "testing"
)

func TestInit(t *testing.T) {
    common.DisableLogs()

    tempfile := filepath.Join(os.TempDir(), "imaged_test.sqlite")

    Convey("Database file should be properly initialized", t, func() {
        os.Remove(tempfile)

        db, err := InitializeWithPath(tempfile)
        So(err, ShouldBeNil)
        defer db.Close()
    })
}
