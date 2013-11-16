package config

import (
    "os"
    "os/user"
    "path/filepath"
)

const (
    XDG_CONFIG_HOME = "XDG_CONFIG_HOME"
    CONFIG_DIR      = ".config"
    IMAGED_DIR      = "imaged"

    CONFIG     = "imaged.conf"
    LOG_CONFIG = "seelog.xml"

    DATABASE_FILE = "imaged.db"
)

var (
    home          = ""
    configFile    = ""
    logConfigFile = ""
)

func ReplaceConfigPaths(homePath string, configFilePath string, logConfigFilePath string) {
    home = homePath
    configFile = configFilePath
    logConfigFile = logConfigFilePath
}

func Home() string {
    if home == "" {
        configHome := os.Getenv(XDG_CONFIG_HOME)

        if configHome == "" {
            user, err := user.Current()
            if err != nil {
                panic(err)
            }
            configHome = filepath.Join(user.HomeDir, CONFIG_DIR)
        }

        home = filepath.Join(configHome, IMAGED_DIR)
    }
    return home
}

func File() string {
    if configFile == "" {
        configFile = filepath.Join(Home(), CONFIG)
    }
    return configFile
}

func Log() string {
    if logConfigFile == "" {
        logConfigFile = filepath.Join(Home(), LOG_CONFIG)
    }
    return logConfigFile
}
