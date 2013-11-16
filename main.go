package main

import (
    "flag"
    log "github.com/cihub/seelog"
    "github.com/dpx-infinity/imaged/config"
)

func main() {
    defer log.Flush()

    configHome := flag.String("home", "", "config home, defaults to ~/.config/imaged")
    configFile := flag.String("config", "", "config file, defaults to ${config home}/imaged.conf")
    logConfigFile := flag.String("logConfig", "", "log config file, defaults to ${config home}/seelog.xml")
    stageDir := flag.String("stage", "", "stage directory, overrides one set in the config")
    storageDir := flag.String("storage", "", "storage directory, overrides one set in the config")
    flag.Parse()

    config.ReplaceConfigPaths(*configHome, *configFile, *logConfigFile)
    config.OverrideDataPaths(*stageDir, *storageDir)

    if err := config.FullLoad(); err != nil {
        log.Critical("Failed to load configuration:", err)
        panic(err)
    }

}
