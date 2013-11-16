package config

import (
    "github.com/BurntSushi/toml"
    "io"
    "os"
    "path/filepath"
)

func (c *ConfigWrapper) OverrideDataPaths(stageDirPath string, storageDirPath string) {
    c.Overrides.Paths.StageDir = stageDirPath
    c.Overrides.Paths.StorageDir = storageDirPath
}

func (c *ConfigWrapper) FullLoad(path string) error {
    if err := c.LoadFile(path); err != nil {
        return err
    }
    c.Fix()
    c.Interpolate()
    return c.Validate()
}

func (c *ConfigWrapper) Load(reader io.Reader) error {
    // We completely replace config object here instead of in-place loading
    // This is needed for safe runtime config reload
    var conf Config
    if _, err := toml.DecodeReader(reader, &conf); err != nil {
        return err
    }
    c.Config = &conf
    return nil
}

func (c *ConfigWrapper) LoadFile(path string) error {
    reader, err := os.Open(path)
    if err != nil {
        return err
    }
    defer reader.Close()

    return c.Load(reader)
}

func (c *ConfigWrapper) Fix() {
    if c.Config != nil {
        if c.Paths.StageDir == "" {
            c.Paths.StageDir = c.Overrides.Paths.StageDir
        }
        if c.Paths.StorageDir == "" {
            c.Paths.StorageDir = c.Overrides.Paths.StorageDir
        }
    }
}

func (c *ConfigWrapper) Interpolate() {
    if c.Config != nil {
        c.Interface.Network.Host = os.ExpandEnv(c.Interface.Network.Host)
        c.Interface.Socket.Path = os.ExpandEnv(c.Interface.Socket.Path)

        c.Paths.StageDir = os.ExpandEnv(c.Paths.StageDir)
        c.Paths.StorageDir = os.ExpandEnv(c.Paths.StorageDir)
    }
}

func (c *ConfigWrapper) Validate() error {
    if c.Config == nil {
        return ValidationError("Configuration is not loaded")
    }

    if c.Paths.StageDir == "" {
        return ValidationError("Stage directory is not configured and override is not present")
    }

    if c.Paths.StorageDir == "" {
        return ValidationError("Storage directory is not configured and override is not present")
    }

    if !c.Interface.Network.Enabled && !c.Interface.Socket.Enabled {
        return ValidationError("Both socket and network interfaces are disabled")
    }

    // TODO: more validations?

    return nil
}

func (p *Paths) DatabaseFile() string {
    return filepath.Join(p.StorageDir, DATABASE_FILE)
}
