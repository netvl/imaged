package common

import (
    log "github.com/cihub/seelog"
)

const (
    DEFAULT_LOG_CONFIG = `
        <seelog>
            <outputs>
                <console/>
            </outputs>
        </seelog>
    `

    DEFAULT_SYNC_LOG_CONFIG = `
        <seelog type="sync">
            <outputs>
                <console/>
            </outputs>
        </seelog>
    `

    DEFAULT_DISABLE_LOGS_CONFIG = `
        <seelog type="sync" levels="off" />
    `
)

func InitDefaultLogConfig() {
    logger, err := log.LoggerFromConfigAsString(DEFAULT_LOG_CONFIG)
    if err != nil {
        panic(err)
    }
    log.ReplaceLogger(logger)
}

func InitDefaultSyncLogConfig() {
    logger, err := log.LoggerFromConfigAsString(DEFAULT_SYNC_LOG_CONFIG)
    if err != nil {
        panic(err)
    }
    log.ReplaceLogger(logger)
}

func DisableLogs() {
    logger, err := log.LoggerFromConfigAsString(DEFAULT_DISABLE_LOGS_CONFIG)
    if err != nil {
        panic(err)
    }
    log.ReplaceLogger(logger)
}

func LogFlush() {
    log.Flush()
}
