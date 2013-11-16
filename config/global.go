package config

import (
    "sync"
)

var (
    config ConfigWrapper
    lock   sync.RWMutex
)

func OverrideDataPaths(stageDirPath string, storageDirPath string) {
    lock.Lock()
    defer lock.Unlock()
    config.OverrideDataPaths(stageDirPath, storageDirPath)
}

func Object() *Config {
    lock.RLock()
    defer lock.RUnlock()
    return config.Config
}

func FullLoad() error {
    lock.Lock()
    defer lock.Unlock()
    return config.FullLoad(File())
}
