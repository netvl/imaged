package config

import (
    "bytes"
    . "github.com/smartystreets/goconvey/convey"
    "os"
    "testing"
)

const (
    CONFIG_FULL = `
        [paths]
        stage-dir = "/stage/dir"
        storage-dir = "/storage/dir"

        [interface.socket]
        enabled = true
        path = "/socket/path"
        
        [interface.network]
        enabled = true
        host = "abcd"
        port = 11111
    `
)

func TestFullLoad(t *testing.T) {
    var cw *ConfigWrapper
    var reader *bytes.Buffer

    Convey("Given a buffer with a full config file", t, func() {
        reader = bytes.NewBufferString(CONFIG_FULL)

        Convey("When loading it", func() {
            cw = new(ConfigWrapper)
            err := cw.Load(reader)

            Convey("Then there should be no errors", func() {
                So(err, ShouldBeNil)
            })

            Convey("And validation should pass successfully", func() {
                So(cw.Validate(), ShouldBeNil)
            })

            Convey("And config struct should contain loaded values", func() {
                So(cw.Paths.StageDir, ShouldEqual, "/stage/dir")
                So(cw.Paths.StorageDir, ShouldEqual, "/storage/dir")

                So(cw.Interface.Socket.Enabled, ShouldBeTrue)
                So(cw.Interface.Socket.Path, ShouldEqual, "/socket/path")

                So(cw.Interface.Network.Enabled, ShouldBeTrue)
                So(cw.Interface.Network.Host, ShouldEqual, "abcd")
                So(cw.Interface.Network.Port, ShouldEqual, 11111)
            })
        })
    })
}

const (
    CONFIG_MINIMAL = `
        [interface.socket]
        enabled = true
        path = "/some/where"
    `
)

func TestDefaults(t *testing.T) {
    var cw *ConfigWrapper
    var reader *bytes.Buffer

    Convey("Given minimal config file and data path overrides", t, func() {
        cw = new(ConfigWrapper)

        reader = bytes.NewBufferString(CONFIG_MINIMAL)
        cw.OverrideDataPaths("/stage/dir", "/storage/dir")

        Convey("When loading it and applying fixes", func() {
            err := cw.Load(reader)
            cw.Fix()

            Convey("Then there should be no errors", func() {
                So(err, ShouldBeNil)
            })

            Convey("And validation should pass successfully", func() {
                So(cw.Validate(), ShouldBeNil)
            })

            Convey("And config struct should be filled with sensible defaults", func() {
                So(cw.Paths.StageDir, ShouldEqual, "/stage/dir")
                So(cw.Paths.StorageDir, ShouldEqual, "/storage/dir")

                So(cw.Interface.Socket.Enabled, ShouldBeTrue)
                So(cw.Interface.Socket.Path, ShouldEqual, "/some/where")

                So(cw.Interface.Network.Enabled, ShouldBeFalse)
                So(cw.Interface.Network.Host, ShouldBeBlank)
                So(cw.Interface.Network.Port, ShouldEqual, 0)
            })
        })
    })
}

const (
    CONFIG_INTERPOLATION = `
        [paths]
        stage-dir = "$VAR1/a/b"
        storage-dir = "$VAR2/c/d"

        [interface.socket]
        path = "/e/$VAR3/f"

        [interface.network]
        host = "/g/$VAR4/h"
    `
)

func TestInterpolation(t *testing.T) {
    var cw *ConfigWrapper
    var reader *bytes.Buffer

    os.Setenv("VAR1", "varval1")
    os.Setenv("VAR2", "varval2")
    os.Setenv("VAR3", "varval3")
    os.Setenv("VAR4", "varval4")

    Convey("Given config file with references to environment variables", t, func() {
        cw = new(ConfigWrapper)
        reader = bytes.NewBufferString(CONFIG_INTERPOLATION)

        Convey("When loading it and performing interpolation", func() {
            err := cw.Load(reader)
            cw.Interpolate()

            Convey("Then there should be no errors", func() {
                So(err, ShouldBeNil)
            })

            Convey("Then stage dir should be interpolated", func() {
                So(cw.Paths.StageDir, ShouldEqual, "varval1/a/b")
            })

            Convey("And storage dir should be interpolated", func() {
                So(cw.Paths.StorageDir, ShouldEqual, "varval2/c/d")
            })

            Convey("And socket path should be interpolated", func() {
                So(cw.Interface.Socket.Path, ShouldEqual, "/e/varval3/f")
            })

            Convey("And network host should be interpolated", func() {
                So(cw.Interface.Network.Host, ShouldEqual, "/g/varval4/h")
            })
        })
    })
}

const (
    CONFIG_VALIDATION_NO_STAGE_DIR = `
        [paths]
        storage-dir = "/some/where"

        [interface.socket]
        enabled = true
        path = "/a/b"
    `

    CONFIG_VALIDATION_NO_STORAGE_DIR = `
        [paths]
        stage-dir = "/some/where"

        [interface.socket]
        enabled = true
        path = "/a/b"
    `

    CONFIG_VALIDATION_NO_INTERFACES = `
        [paths]
        stage-dir = "/some/where"
        storage-dir = "/some/where/else"
    `
)

func TestValidation(t *testing.T) {
    var cw *ConfigWrapper
    var reader *bytes.Buffer

    Convey("While testing validation", t, func() {
        Convey("Given not loaded config structure", func() {
            cw = new(ConfigWrapper)

            Convey("When validating it", func() {
                verr := cw.Validate()

                Convey("Then there should be validation error", func() {
                    So(verr, ShouldNotBeNil)
                    So(verr, ShouldHaveSameTypeAs, ValidationError(""))
                })
            })
        })

        Convey("Given config file without stage dir and no stage dir overrides", func() {
            cw = new(ConfigWrapper)
            reader = bytes.NewBufferString(CONFIG_VALIDATION_NO_STAGE_DIR)

            Convey("When loading and validating it", func() {
                err := cw.Load(reader)
                verr := cw.Validate()

                Convey("Then there should be no loading errors", func() {
                    So(err, ShouldBeNil)
                })

                Convey("But there should be validation error", func() {
                    So(verr, ShouldNotBeNil)
                    So(verr, ShouldHaveSameTypeAs, ValidationError(""))
                })
            })
        })

        Convey("Given config file without storage dir and no storage dir overrides", func() {
            cw = new(ConfigWrapper)
            reader = bytes.NewBufferString(CONFIG_VALIDATION_NO_STORAGE_DIR)

            Convey("When loading and validating it", func() {
                err := cw.Load(reader)
                verr := cw.Validate()

                Convey("Then there should be no loading errors", func() {
                    So(err, ShouldBeNil)
                })

                Convey("But there should be validation error", func() {
                    So(verr, ShouldNotBeNil)
                    So(verr, ShouldHaveSameTypeAs, ValidationError(""))
                })
            })
        })

        Convey("Given config file without any interfaces configured", func() {
            cw = new(ConfigWrapper)
            reader = bytes.NewBufferString(CONFIG_VALIDATION_NO_INTERFACES)

            Convey("When loading and validating it", func() {
                err := cw.Load(reader)
                verr := cw.Validate()

                Convey("Then there should be no loading errors", func() {
                    So(err, ShouldBeNil)
                })

                Convey("But there should be validation error", func() {
                    So(verr, ShouldNotBeNil)
                    So(verr, ShouldHaveSameTypeAs, ValidationError(""))
                })
            })
        })
    })
}
